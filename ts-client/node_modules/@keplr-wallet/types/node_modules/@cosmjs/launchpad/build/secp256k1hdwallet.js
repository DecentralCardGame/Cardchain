"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Secp256k1HdWallet = exports.extractKdfConfiguration = void 0;
const crypto_1 = require("@cosmjs/crypto");
const encoding_1 = require("@cosmjs/encoding");
const utils_1 = require("@cosmjs/utils");
const address_1 = require("./address");
const encoding_2 = require("./encoding");
const paths_1 = require("./paths");
const signature_1 = require("./signature");
const wallet_1 = require("./wallet");
const serializationTypeV1 = "secp256k1wallet-v1";
/**
 * A KDF configuration that is not very strong but can be used on the main thread.
 * It takes about 1 second in Node.js 12.15 and should have similar runtimes in other modern Wasm hosts.
 */
const basicPasswordHashingOptions = {
    algorithm: "argon2id",
    params: {
        outputLength: 32,
        opsLimit: 20,
        memLimitKib: 12 * 1024,
    },
};
function isDerivationJson(thing) {
    if (!utils_1.isNonNullObject(thing))
        return false;
    if (typeof thing.hdPath !== "string")
        return false;
    if (typeof thing.prefix !== "string")
        return false;
    return true;
}
function extractKdfConfigurationV1(doc) {
    return doc.kdf;
}
function extractKdfConfiguration(serialization) {
    const root = JSON.parse(serialization);
    if (!utils_1.isNonNullObject(root))
        throw new Error("Root document is not an object.");
    switch (root.type) {
        case serializationTypeV1:
            return extractKdfConfigurationV1(root);
        default:
            throw new Error("Unsupported serialization type");
    }
}
exports.extractKdfConfiguration = extractKdfConfiguration;
class Secp256k1HdWallet {
    constructor(mnemonic, hdPath, privkey, pubkey, prefix) {
        this.secret = mnemonic;
        this.accounts = [
            {
                hdPath: hdPath,
                prefix: prefix,
            },
        ];
        this.privkey = privkey;
        this.pubkey = pubkey;
    }
    /**
     * Restores a wallet from the given BIP39 mnemonic.
     *
     * @param mnemonic Any valid English mnemonic.
     * @param hdPath The BIP-32/SLIP-10 derivation path. Defaults to the Cosmos Hub/ATOM path `m/44'/118'/0'/0/0`.
     * @param prefix The bech32 address prefix (human readable part). Defaults to "cosmos".
     */
    static async fromMnemonic(mnemonic, hdPath = paths_1.makeCosmoshubPath(0), prefix = "cosmos") {
        const mnemonicChecked = new crypto_1.EnglishMnemonic(mnemonic);
        const seed = await crypto_1.Bip39.mnemonicToSeed(mnemonicChecked);
        const { privkey } = crypto_1.Slip10.derivePath(crypto_1.Slip10Curve.Secp256k1, seed, hdPath);
        const uncompressed = (await crypto_1.Secp256k1.makeKeypair(privkey)).pubkey;
        return new Secp256k1HdWallet(mnemonicChecked, hdPath, privkey, crypto_1.Secp256k1.compressPubkey(uncompressed), prefix);
    }
    /**
     * Generates a new wallet with a BIP39 mnemonic of the given length.
     *
     * @param length The number of words in the mnemonic (12, 15, 18, 21 or 24).
     * @param hdPath The BIP-32/SLIP-10 derivation path. Defaults to the Cosmos Hub/ATOM path `m/44'/118'/0'/0/0`.
     * @param prefix The bech32 address prefix (human readable part). Defaults to "cosmos".
     */
    static async generate(length = 12, hdPath = paths_1.makeCosmoshubPath(0), prefix = "cosmos") {
        const entropyLength = 4 * Math.floor((11 * length) / 33);
        const entropy = crypto_1.Random.getBytes(entropyLength);
        const mnemonic = crypto_1.Bip39.encode(entropy);
        return Secp256k1HdWallet.fromMnemonic(mnemonic.toString(), hdPath, prefix);
    }
    /**
     * Restores a wallet from an encrypted serialization.
     *
     * @param password The user provided password used to generate an encryption key via a KDF.
     *                 This is not normalized internally (see "Unicode normalization" to learn more).
     */
    static async deserialize(serialization, password) {
        const root = JSON.parse(serialization);
        if (!utils_1.isNonNullObject(root))
            throw new Error("Root document is not an object.");
        switch (root.type) {
            case serializationTypeV1:
                return Secp256k1HdWallet.deserializeTypeV1(serialization, password);
            default:
                throw new Error("Unsupported serialization type");
        }
    }
    /**
     * Restores a wallet from an encrypted serialization.
     *
     * This is an advanced alternative to calling `deserialize(serialization, password)` directly, which allows
     * you to offload the KDF execution to a non-UI thread (e.g. in a WebWorker).
     *
     * The caller is responsible for ensuring the key was derived with the given KDF configuration. This can be
     * done using `extractKdfConfiguration(serialization)` and `executeKdf(password, kdfConfiguration)` from this package.
     */
    static async deserializeWithEncryptionKey(serialization, encryptionKey) {
        const root = JSON.parse(serialization);
        if (!utils_1.isNonNullObject(root))
            throw new Error("Root document is not an object.");
        const untypedRoot = root;
        switch (untypedRoot.type) {
            case serializationTypeV1: {
                const decryptedBytes = await wallet_1.decrypt(encoding_1.fromBase64(untypedRoot.data), encryptionKey, untypedRoot.encryption);
                const decryptedDocument = JSON.parse(encoding_1.fromUtf8(decryptedBytes));
                const { mnemonic, accounts } = decryptedDocument;
                utils_1.assert(typeof mnemonic === "string");
                if (!Array.isArray(accounts))
                    throw new Error("Property 'accounts' is not an array");
                if (accounts.length !== 1)
                    throw new Error("Property 'accounts' only supports one entry");
                const account = accounts[0];
                if (!isDerivationJson(account))
                    throw new Error("Account is not in the correct format.");
                return Secp256k1HdWallet.fromMnemonic(mnemonic, crypto_1.stringToPath(account.hdPath), account.prefix);
            }
            default:
                throw new Error("Unsupported serialization type");
        }
    }
    static async deserializeTypeV1(serialization, password) {
        const root = JSON.parse(serialization);
        if (!utils_1.isNonNullObject(root))
            throw new Error("Root document is not an object.");
        const encryptionKey = await wallet_1.executeKdf(password, root.kdf);
        return Secp256k1HdWallet.deserializeWithEncryptionKey(serialization, encryptionKey);
    }
    get mnemonic() {
        return this.secret.toString();
    }
    get address() {
        return address_1.rawSecp256k1PubkeyToAddress(this.pubkey, this.accounts[0].prefix);
    }
    async getAccounts() {
        return [
            {
                algo: "secp256k1",
                address: this.address,
                pubkey: this.pubkey,
            },
        ];
    }
    async signAmino(signerAddress, signDoc) {
        if (signerAddress !== this.address) {
            throw new Error(`Address ${signerAddress} not found in wallet`);
        }
        const message = crypto_1.sha256(encoding_2.serializeSignDoc(signDoc));
        const signature = await crypto_1.Secp256k1.createSignature(message, this.privkey);
        const signatureBytes = new Uint8Array([...signature.r(32), ...signature.s(32)]);
        return {
            signed: signDoc,
            signature: signature_1.encodeSecp256k1Signature(this.pubkey, signatureBytes),
        };
    }
    /**
     * Generates an encrypted serialization of this wallet.
     *
     * @param password The user provided password used to generate an encryption key via a KDF.
     *                 This is not normalized internally (see "Unicode normalization" to learn more).
     */
    async serialize(password) {
        const kdfConfiguration = basicPasswordHashingOptions;
        const encryptionKey = await wallet_1.executeKdf(password, kdfConfiguration);
        return this.serializeWithEncryptionKey(encryptionKey, kdfConfiguration);
    }
    /**
     * Generates an encrypted serialization of this wallet.
     *
     * This is an advanced alternative to calling `serialize(password)` directly, which allows you to
     * offload the KDF execution to a non-UI thread (e.g. in a WebWorker).
     *
     * The caller is responsible for ensuring the key was derived with the given KDF options. If this
     * is not the case, the wallet cannot be restored with the original password.
     */
    async serializeWithEncryptionKey(encryptionKey, kdfConfiguration) {
        const dataToEncrypt = {
            mnemonic: this.mnemonic,
            accounts: this.accounts.map((account) => ({
                hdPath: crypto_1.pathToString(account.hdPath),
                prefix: account.prefix,
            })),
        };
        const dataToEncryptRaw = encoding_1.toUtf8(JSON.stringify(dataToEncrypt));
        const encryptionConfiguration = {
            algorithm: wallet_1.supportedAlgorithms.xchacha20poly1305Ietf,
        };
        const encryptedData = await wallet_1.encrypt(dataToEncryptRaw, encryptionKey, encryptionConfiguration);
        const out = {
            type: serializationTypeV1,
            kdf: kdfConfiguration,
            encryption: encryptionConfiguration,
            data: encoding_1.toBase64(encryptedData),
        };
        return JSON.stringify(out);
    }
}
exports.Secp256k1HdWallet = Secp256k1HdWallet;
//# sourceMappingURL=secp256k1hdwallet.js.map