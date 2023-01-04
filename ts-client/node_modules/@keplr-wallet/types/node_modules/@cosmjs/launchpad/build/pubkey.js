"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.encodeBech32Pubkey = exports.encodeAminoPubkey = exports.decodeBech32Pubkey = exports.decodeAminoPubkey = exports.encodeSecp256k1Pubkey = void 0;
const encoding_1 = require("@cosmjs/encoding");
const utils_1 = require("@cosmjs/utils");
const types_1 = require("./types");
function encodeSecp256k1Pubkey(pubkey) {
    if (pubkey.length !== 33 || (pubkey[0] !== 0x02 && pubkey[0] !== 0x03)) {
        throw new Error("Public key must be compressed secp256k1, i.e. 33 bytes starting with 0x02 or 0x03");
    }
    return {
        type: types_1.pubkeyType.secp256k1,
        value: encoding_1.toBase64(pubkey),
    };
}
exports.encodeSecp256k1Pubkey = encodeSecp256k1Pubkey;
// As discussed in https://github.com/binance-chain/javascript-sdk/issues/163
// Prefixes listed here: https://github.com/tendermint/tendermint/blob/d419fffe18531317c28c29a292ad7d253f6cafdf/docs/spec/blockchain/encoding.md#public-key-cryptography
// Last bytes is varint-encoded length prefix
const pubkeyAminoPrefixSecp256k1 = encoding_1.fromHex("eb5ae98721");
const pubkeyAminoPrefixEd25519 = encoding_1.fromHex("1624de6420");
const pubkeyAminoPrefixSr25519 = encoding_1.fromHex("0dfb1005");
const pubkeyAminoPrefixLength = pubkeyAminoPrefixSecp256k1.length;
/**
 * Decodes a pubkey in the Amino binary format to a type/value object.
 */
function decodeAminoPubkey(data) {
    const aminoPrefix = data.slice(0, pubkeyAminoPrefixLength);
    const rest = data.slice(pubkeyAminoPrefixLength);
    if (utils_1.arrayContentEquals(aminoPrefix, pubkeyAminoPrefixSecp256k1)) {
        if (rest.length !== 33) {
            throw new Error("Invalid rest data length. Expected 33 bytes (compressed secp256k1 pubkey).");
        }
        return {
            type: types_1.pubkeyType.secp256k1,
            value: encoding_1.toBase64(rest),
        };
    }
    else if (utils_1.arrayContentEquals(aminoPrefix, pubkeyAminoPrefixEd25519)) {
        if (rest.length !== 32) {
            throw new Error("Invalid rest data length. Expected 32 bytes (Ed25519 pubkey).");
        }
        return {
            type: types_1.pubkeyType.ed25519,
            value: encoding_1.toBase64(rest),
        };
    }
    else if (utils_1.arrayContentEquals(aminoPrefix, pubkeyAminoPrefixSr25519)) {
        if (rest.length !== 32) {
            throw new Error("Invalid rest data length. Expected 32 bytes (Sr25519 pubkey).");
        }
        return {
            type: types_1.pubkeyType.sr25519,
            value: encoding_1.toBase64(rest),
        };
    }
    else {
        throw new Error("Unsupported Pubkey type. Amino prefix: " + encoding_1.toHex(aminoPrefix));
    }
}
exports.decodeAminoPubkey = decodeAminoPubkey;
/**
 * Decodes a bech32 pubkey to Amino binary, which is then decoded to a type/value object.
 * The bech32 prefix is ignored and discareded.
 *
 * @param bechEncoded the bech32 encoded pubkey
 */
function decodeBech32Pubkey(bechEncoded) {
    const { data } = encoding_1.Bech32.decode(bechEncoded);
    return decodeAminoPubkey(data);
}
exports.decodeBech32Pubkey = decodeBech32Pubkey;
/**
 * Encodes a public key to binary Amino.
 */
function encodeAminoPubkey(pubkey) {
    let aminoPrefix;
    switch (pubkey.type) {
        // Note: please don't add cases here without writing additional unit tests
        case types_1.pubkeyType.secp256k1:
            aminoPrefix = pubkeyAminoPrefixSecp256k1;
            break;
        case types_1.pubkeyType.ed25519:
            aminoPrefix = pubkeyAminoPrefixEd25519;
            break;
        default:
            throw new Error("Unsupported pubkey type");
    }
    return new Uint8Array([...aminoPrefix, ...encoding_1.fromBase64(pubkey.value)]);
}
exports.encodeAminoPubkey = encodeAminoPubkey;
/**
 * Encodes a public key to binary Amino and then to bech32.
 *
 * @param pubkey the public key to encode
 * @param prefix the bech32 prefix (human readable part)
 */
function encodeBech32Pubkey(pubkey, prefix) {
    return encoding_1.Bech32.encode(prefix, encodeAminoPubkey(pubkey));
}
exports.encodeBech32Pubkey = encodeBech32Pubkey;
//# sourceMappingURL=pubkey.js.map