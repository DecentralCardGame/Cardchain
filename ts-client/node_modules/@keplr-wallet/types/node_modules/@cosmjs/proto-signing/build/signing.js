"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.makeSignBytes = exports.makeSignDoc = exports.makeAuthInfoBytes = void 0;
/* eslint-disable @typescript-eslint/naming-convention */
const long_1 = __importDefault(require("long"));
const signing_1 = require("./codec/cosmos/tx/signing/v1beta1/signing");
const tx_1 = require("./codec/cosmos/tx/v1beta1/tx");
/**
 * Creates and serializes an AuthInfo document using SIGN_MODE_DIRECT.
 */
function makeAuthInfoBytes(pubkeys, feeAmount, gasLimit, sequence, signMode = signing_1.SignMode.SIGN_MODE_DIRECT) {
    const authInfo = {
        signerInfos: pubkeys.map((pubkey) => ({
            publicKey: pubkey,
            modeInfo: {
                single: { mode: signMode },
            },
            sequence: long_1.default.fromNumber(sequence),
        })),
        fee: {
            amount: [...feeAmount],
            gasLimit: long_1.default.fromNumber(gasLimit),
        },
    };
    return tx_1.AuthInfo.encode(tx_1.AuthInfo.fromPartial(authInfo)).finish();
}
exports.makeAuthInfoBytes = makeAuthInfoBytes;
function makeSignDoc(bodyBytes, authInfoBytes, chainId, accountNumber) {
    return {
        bodyBytes: bodyBytes,
        authInfoBytes: authInfoBytes,
        chainId: chainId,
        accountNumber: long_1.default.fromNumber(accountNumber),
    };
}
exports.makeSignDoc = makeSignDoc;
function makeSignBytes({ accountNumber, authInfoBytes, bodyBytes, chainId }) {
    const signDoc = tx_1.SignDoc.fromPartial({
        accountNumber: accountNumber,
        authInfoBytes: authInfoBytes,
        bodyBytes: bodyBytes,
        chainId: chainId,
    });
    return tx_1.SignDoc.encode(signDoc).finish();
}
exports.makeSignBytes = makeSignBytes;
//# sourceMappingURL=signing.js.map