"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.decodePubkey = exports.encodePubkey = void 0;
/* eslint-disable @typescript-eslint/naming-convention */
const encoding_1 = require("@cosmjs/encoding");
const launchpad_1 = require("@cosmjs/launchpad");
const keys_1 = require("./codec/cosmos/crypto/secp256k1/keys");
const any_1 = require("./codec/google/protobuf/any");
function encodePubkey(pubkey) {
    switch (pubkey.type) {
        case "tendermint/PubKeySecp256k1": {
            const pubkeyProto = keys_1.PubKey.fromPartial({
                key: encoding_1.fromBase64(pubkey.value),
            });
            return any_1.Any.fromPartial({
                typeUrl: "/cosmos.crypto.secp256k1.PubKey",
                value: Uint8Array.from(keys_1.PubKey.encode(pubkeyProto).finish()),
            });
        }
        default:
            throw new Error(`Pubkey type ${pubkey.type} not recognized`);
    }
}
exports.encodePubkey = encodePubkey;
function decodePubkey(pubkey) {
    if (!pubkey || !pubkey.value) {
        return null;
    }
    switch (pubkey.typeUrl) {
        case "/cosmos.crypto.secp256k1.PubKey": {
            const { key } = keys_1.PubKey.decode(pubkey.value);
            return launchpad_1.encodeSecp256k1Pubkey(key);
        }
        default:
            throw new Error(`Pubkey type_url ${pubkey.typeUrl} not recognized`);
    }
}
exports.decodePubkey = decodePubkey;
//# sourceMappingURL=pubkey.js.map