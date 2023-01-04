"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.AuxSignerData = exports.Tip = exports.Fee = exports.ModeInfo_Multi = exports.ModeInfo_Single = exports.ModeInfo = exports.SignerInfo = exports.AuthInfo = exports.TxBody = exports.SignDocDirectAux = exports.SignDoc = exports.TxRaw = exports.Tx = exports.protobufPackage = void 0;
/* eslint-disable */
const long_1 = require("long");
const minimal_1 = require("protobufjs/minimal");
const any_1 = require("../../../google/protobuf/any");
const coin_1 = require("../../base/v1beta1/coin");
const multisig_1 = require("../../crypto/multisig/v1beta1/multisig");
const signing_1 = require("../signing/v1beta1/signing");
exports.protobufPackage = "cosmos.tx.v1beta1";
function createBaseTx() {
    return { body: undefined, authInfo: undefined, signatures: [] };
}
exports.Tx = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.body !== undefined) {
            exports.TxBody.encode(message.body, writer.uint32(10).fork()).ldelim();
        }
        if (message.authInfo !== undefined) {
            exports.AuthInfo.encode(message.authInfo, writer.uint32(18).fork()).ldelim();
        }
        for (const v of message.signatures) {
            writer.uint32(26).bytes(v);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseTx();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.body = exports.TxBody.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.authInfo = exports.AuthInfo.decode(reader, reader.uint32());
                    break;
                case 3:
                    message.signatures.push(reader.bytes());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return {
            body: isSet(object.body) ? exports.TxBody.fromJSON(object.body) : undefined,
            authInfo: isSet(object.authInfo) ? exports.AuthInfo.fromJSON(object.authInfo) : undefined,
            signatures: Array.isArray(object?.signatures) ? object.signatures.map((e) => bytesFromBase64(e)) : [],
        };
    },
    toJSON(message) {
        const obj = {};
        message.body !== undefined && (obj.body = message.body ? exports.TxBody.toJSON(message.body) : undefined);
        message.authInfo !== undefined && (obj.authInfo = message.authInfo ? exports.AuthInfo.toJSON(message.authInfo) : undefined);
        if (message.signatures) {
            obj.signatures = message.signatures.map((e) => base64FromBytes(e !== undefined ? e : new Uint8Array()));
        }
        else {
            obj.signatures = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = createBaseTx();
        message.body = (object.body !== undefined && object.body !== null) ? exports.TxBody.fromPartial(object.body) : undefined;
        message.authInfo = (object.authInfo !== undefined && object.authInfo !== null)
            ? exports.AuthInfo.fromPartial(object.authInfo)
            : undefined;
        message.signatures = object.signatures?.map((e) => e) || [];
        return message;
    },
};
function createBaseTxRaw() {
    return { bodyBytes: new Uint8Array(), authInfoBytes: new Uint8Array(), signatures: [] };
}
exports.TxRaw = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.bodyBytes.length !== 0) {
            writer.uint32(10).bytes(message.bodyBytes);
        }
        if (message.authInfoBytes.length !== 0) {
            writer.uint32(18).bytes(message.authInfoBytes);
        }
        for (const v of message.signatures) {
            writer.uint32(26).bytes(v);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseTxRaw();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.bodyBytes = reader.bytes();
                    break;
                case 2:
                    message.authInfoBytes = reader.bytes();
                    break;
                case 3:
                    message.signatures.push(reader.bytes());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return {
            bodyBytes: isSet(object.bodyBytes) ? bytesFromBase64(object.bodyBytes) : new Uint8Array(),
            authInfoBytes: isSet(object.authInfoBytes) ? bytesFromBase64(object.authInfoBytes) : new Uint8Array(),
            signatures: Array.isArray(object?.signatures) ? object.signatures.map((e) => bytesFromBase64(e)) : [],
        };
    },
    toJSON(message) {
        const obj = {};
        message.bodyBytes !== undefined
            && (obj.bodyBytes = base64FromBytes(message.bodyBytes !== undefined ? message.bodyBytes : new Uint8Array()));
        message.authInfoBytes !== undefined
            && (obj.authInfoBytes = base64FromBytes(message.authInfoBytes !== undefined ? message.authInfoBytes : new Uint8Array()));
        if (message.signatures) {
            obj.signatures = message.signatures.map((e) => base64FromBytes(e !== undefined ? e : new Uint8Array()));
        }
        else {
            obj.signatures = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = createBaseTxRaw();
        message.bodyBytes = object.bodyBytes ?? new Uint8Array();
        message.authInfoBytes = object.authInfoBytes ?? new Uint8Array();
        message.signatures = object.signatures?.map((e) => e) || [];
        return message;
    },
};
function createBaseSignDoc() {
    return { bodyBytes: new Uint8Array(), authInfoBytes: new Uint8Array(), chainId: "", accountNumber: 0 };
}
exports.SignDoc = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.bodyBytes.length !== 0) {
            writer.uint32(10).bytes(message.bodyBytes);
        }
        if (message.authInfoBytes.length !== 0) {
            writer.uint32(18).bytes(message.authInfoBytes);
        }
        if (message.chainId !== "") {
            writer.uint32(26).string(message.chainId);
        }
        if (message.accountNumber !== 0) {
            writer.uint32(32).uint64(message.accountNumber);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseSignDoc();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.bodyBytes = reader.bytes();
                    break;
                case 2:
                    message.authInfoBytes = reader.bytes();
                    break;
                case 3:
                    message.chainId = reader.string();
                    break;
                case 4:
                    message.accountNumber = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return {
            bodyBytes: isSet(object.bodyBytes) ? bytesFromBase64(object.bodyBytes) : new Uint8Array(),
            authInfoBytes: isSet(object.authInfoBytes) ? bytesFromBase64(object.authInfoBytes) : new Uint8Array(),
            chainId: isSet(object.chainId) ? String(object.chainId) : "",
            accountNumber: isSet(object.accountNumber) ? Number(object.accountNumber) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        message.bodyBytes !== undefined
            && (obj.bodyBytes = base64FromBytes(message.bodyBytes !== undefined ? message.bodyBytes : new Uint8Array()));
        message.authInfoBytes !== undefined
            && (obj.authInfoBytes = base64FromBytes(message.authInfoBytes !== undefined ? message.authInfoBytes : new Uint8Array()));
        message.chainId !== undefined && (obj.chainId = message.chainId);
        message.accountNumber !== undefined && (obj.accountNumber = Math.round(message.accountNumber));
        return obj;
    },
    fromPartial(object) {
        const message = createBaseSignDoc();
        message.bodyBytes = object.bodyBytes ?? new Uint8Array();
        message.authInfoBytes = object.authInfoBytes ?? new Uint8Array();
        message.chainId = object.chainId ?? "";
        message.accountNumber = object.accountNumber ?? 0;
        return message;
    },
};
function createBaseSignDocDirectAux() {
    return {
        bodyBytes: new Uint8Array(),
        publicKey: undefined,
        chainId: "",
        accountNumber: 0,
        sequence: 0,
        tip: undefined,
    };
}
exports.SignDocDirectAux = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.bodyBytes.length !== 0) {
            writer.uint32(10).bytes(message.bodyBytes);
        }
        if (message.publicKey !== undefined) {
            any_1.Any.encode(message.publicKey, writer.uint32(18).fork()).ldelim();
        }
        if (message.chainId !== "") {
            writer.uint32(26).string(message.chainId);
        }
        if (message.accountNumber !== 0) {
            writer.uint32(32).uint64(message.accountNumber);
        }
        if (message.sequence !== 0) {
            writer.uint32(40).uint64(message.sequence);
        }
        if (message.tip !== undefined) {
            exports.Tip.encode(message.tip, writer.uint32(50).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseSignDocDirectAux();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.bodyBytes = reader.bytes();
                    break;
                case 2:
                    message.publicKey = any_1.Any.decode(reader, reader.uint32());
                    break;
                case 3:
                    message.chainId = reader.string();
                    break;
                case 4:
                    message.accountNumber = longToNumber(reader.uint64());
                    break;
                case 5:
                    message.sequence = longToNumber(reader.uint64());
                    break;
                case 6:
                    message.tip = exports.Tip.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return {
            bodyBytes: isSet(object.bodyBytes) ? bytesFromBase64(object.bodyBytes) : new Uint8Array(),
            publicKey: isSet(object.publicKey) ? any_1.Any.fromJSON(object.publicKey) : undefined,
            chainId: isSet(object.chainId) ? String(object.chainId) : "",
            accountNumber: isSet(object.accountNumber) ? Number(object.accountNumber) : 0,
            sequence: isSet(object.sequence) ? Number(object.sequence) : 0,
            tip: isSet(object.tip) ? exports.Tip.fromJSON(object.tip) : undefined,
        };
    },
    toJSON(message) {
        const obj = {};
        message.bodyBytes !== undefined
            && (obj.bodyBytes = base64FromBytes(message.bodyBytes !== undefined ? message.bodyBytes : new Uint8Array()));
        message.publicKey !== undefined && (obj.publicKey = message.publicKey ? any_1.Any.toJSON(message.publicKey) : undefined);
        message.chainId !== undefined && (obj.chainId = message.chainId);
        message.accountNumber !== undefined && (obj.accountNumber = Math.round(message.accountNumber));
        message.sequence !== undefined && (obj.sequence = Math.round(message.sequence));
        message.tip !== undefined && (obj.tip = message.tip ? exports.Tip.toJSON(message.tip) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseSignDocDirectAux();
        message.bodyBytes = object.bodyBytes ?? new Uint8Array();
        message.publicKey = (object.publicKey !== undefined && object.publicKey !== null)
            ? any_1.Any.fromPartial(object.publicKey)
            : undefined;
        message.chainId = object.chainId ?? "";
        message.accountNumber = object.accountNumber ?? 0;
        message.sequence = object.sequence ?? 0;
        message.tip = (object.tip !== undefined && object.tip !== null) ? exports.Tip.fromPartial(object.tip) : undefined;
        return message;
    },
};
function createBaseTxBody() {
    return { messages: [], memo: "", timeoutHeight: 0, extensionOptions: [], nonCriticalExtensionOptions: [] };
}
exports.TxBody = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        for (const v of message.messages) {
            any_1.Any.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.memo !== "") {
            writer.uint32(18).string(message.memo);
        }
        if (message.timeoutHeight !== 0) {
            writer.uint32(24).uint64(message.timeoutHeight);
        }
        for (const v of message.extensionOptions) {
            any_1.Any.encode(v, writer.uint32(8186).fork()).ldelim();
        }
        for (const v of message.nonCriticalExtensionOptions) {
            any_1.Any.encode(v, writer.uint32(16378).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseTxBody();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.messages.push(any_1.Any.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.memo = reader.string();
                    break;
                case 3:
                    message.timeoutHeight = longToNumber(reader.uint64());
                    break;
                case 1023:
                    message.extensionOptions.push(any_1.Any.decode(reader, reader.uint32()));
                    break;
                case 2047:
                    message.nonCriticalExtensionOptions.push(any_1.Any.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return {
            messages: Array.isArray(object?.messages) ? object.messages.map((e) => any_1.Any.fromJSON(e)) : [],
            memo: isSet(object.memo) ? String(object.memo) : "",
            timeoutHeight: isSet(object.timeoutHeight) ? Number(object.timeoutHeight) : 0,
            extensionOptions: Array.isArray(object?.extensionOptions)
                ? object.extensionOptions.map((e) => any_1.Any.fromJSON(e))
                : [],
            nonCriticalExtensionOptions: Array.isArray(object?.nonCriticalExtensionOptions)
                ? object.nonCriticalExtensionOptions.map((e) => any_1.Any.fromJSON(e))
                : [],
        };
    },
    toJSON(message) {
        const obj = {};
        if (message.messages) {
            obj.messages = message.messages.map((e) => e ? any_1.Any.toJSON(e) : undefined);
        }
        else {
            obj.messages = [];
        }
        message.memo !== undefined && (obj.memo = message.memo);
        message.timeoutHeight !== undefined && (obj.timeoutHeight = Math.round(message.timeoutHeight));
        if (message.extensionOptions) {
            obj.extensionOptions = message.extensionOptions.map((e) => e ? any_1.Any.toJSON(e) : undefined);
        }
        else {
            obj.extensionOptions = [];
        }
        if (message.nonCriticalExtensionOptions) {
            obj.nonCriticalExtensionOptions = message.nonCriticalExtensionOptions.map((e) => e ? any_1.Any.toJSON(e) : undefined);
        }
        else {
            obj.nonCriticalExtensionOptions = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = createBaseTxBody();
        message.messages = object.messages?.map((e) => any_1.Any.fromPartial(e)) || [];
        message.memo = object.memo ?? "";
        message.timeoutHeight = object.timeoutHeight ?? 0;
        message.extensionOptions = object.extensionOptions?.map((e) => any_1.Any.fromPartial(e)) || [];
        message.nonCriticalExtensionOptions = object.nonCriticalExtensionOptions?.map((e) => any_1.Any.fromPartial(e)) || [];
        return message;
    },
};
function createBaseAuthInfo() {
    return { signerInfos: [], fee: undefined, tip: undefined };
}
exports.AuthInfo = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        for (const v of message.signerInfos) {
            exports.SignerInfo.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.fee !== undefined) {
            exports.Fee.encode(message.fee, writer.uint32(18).fork()).ldelim();
        }
        if (message.tip !== undefined) {
            exports.Tip.encode(message.tip, writer.uint32(26).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseAuthInfo();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.signerInfos.push(exports.SignerInfo.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.fee = exports.Fee.decode(reader, reader.uint32());
                    break;
                case 3:
                    message.tip = exports.Tip.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return {
            signerInfos: Array.isArray(object?.signerInfos) ? object.signerInfos.map((e) => exports.SignerInfo.fromJSON(e)) : [],
            fee: isSet(object.fee) ? exports.Fee.fromJSON(object.fee) : undefined,
            tip: isSet(object.tip) ? exports.Tip.fromJSON(object.tip) : undefined,
        };
    },
    toJSON(message) {
        const obj = {};
        if (message.signerInfos) {
            obj.signerInfos = message.signerInfos.map((e) => e ? exports.SignerInfo.toJSON(e) : undefined);
        }
        else {
            obj.signerInfos = [];
        }
        message.fee !== undefined && (obj.fee = message.fee ? exports.Fee.toJSON(message.fee) : undefined);
        message.tip !== undefined && (obj.tip = message.tip ? exports.Tip.toJSON(message.tip) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseAuthInfo();
        message.signerInfos = object.signerInfos?.map((e) => exports.SignerInfo.fromPartial(e)) || [];
        message.fee = (object.fee !== undefined && object.fee !== null) ? exports.Fee.fromPartial(object.fee) : undefined;
        message.tip = (object.tip !== undefined && object.tip !== null) ? exports.Tip.fromPartial(object.tip) : undefined;
        return message;
    },
};
function createBaseSignerInfo() {
    return { publicKey: undefined, modeInfo: undefined, sequence: 0 };
}
exports.SignerInfo = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.publicKey !== undefined) {
            any_1.Any.encode(message.publicKey, writer.uint32(10).fork()).ldelim();
        }
        if (message.modeInfo !== undefined) {
            exports.ModeInfo.encode(message.modeInfo, writer.uint32(18).fork()).ldelim();
        }
        if (message.sequence !== 0) {
            writer.uint32(24).uint64(message.sequence);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseSignerInfo();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.publicKey = any_1.Any.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.modeInfo = exports.ModeInfo.decode(reader, reader.uint32());
                    break;
                case 3:
                    message.sequence = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return {
            publicKey: isSet(object.publicKey) ? any_1.Any.fromJSON(object.publicKey) : undefined,
            modeInfo: isSet(object.modeInfo) ? exports.ModeInfo.fromJSON(object.modeInfo) : undefined,
            sequence: isSet(object.sequence) ? Number(object.sequence) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        message.publicKey !== undefined && (obj.publicKey = message.publicKey ? any_1.Any.toJSON(message.publicKey) : undefined);
        message.modeInfo !== undefined && (obj.modeInfo = message.modeInfo ? exports.ModeInfo.toJSON(message.modeInfo) : undefined);
        message.sequence !== undefined && (obj.sequence = Math.round(message.sequence));
        return obj;
    },
    fromPartial(object) {
        const message = createBaseSignerInfo();
        message.publicKey = (object.publicKey !== undefined && object.publicKey !== null)
            ? any_1.Any.fromPartial(object.publicKey)
            : undefined;
        message.modeInfo = (object.modeInfo !== undefined && object.modeInfo !== null)
            ? exports.ModeInfo.fromPartial(object.modeInfo)
            : undefined;
        message.sequence = object.sequence ?? 0;
        return message;
    },
};
function createBaseModeInfo() {
    return { single: undefined, multi: undefined };
}
exports.ModeInfo = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.single !== undefined) {
            exports.ModeInfo_Single.encode(message.single, writer.uint32(10).fork()).ldelim();
        }
        if (message.multi !== undefined) {
            exports.ModeInfo_Multi.encode(message.multi, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseModeInfo();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.single = exports.ModeInfo_Single.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.multi = exports.ModeInfo_Multi.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return {
            single: isSet(object.single) ? exports.ModeInfo_Single.fromJSON(object.single) : undefined,
            multi: isSet(object.multi) ? exports.ModeInfo_Multi.fromJSON(object.multi) : undefined,
        };
    },
    toJSON(message) {
        const obj = {};
        message.single !== undefined && (obj.single = message.single ? exports.ModeInfo_Single.toJSON(message.single) : undefined);
        message.multi !== undefined && (obj.multi = message.multi ? exports.ModeInfo_Multi.toJSON(message.multi) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseModeInfo();
        message.single = (object.single !== undefined && object.single !== null)
            ? exports.ModeInfo_Single.fromPartial(object.single)
            : undefined;
        message.multi = (object.multi !== undefined && object.multi !== null)
            ? exports.ModeInfo_Multi.fromPartial(object.multi)
            : undefined;
        return message;
    },
};
function createBaseModeInfo_Single() {
    return { mode: 0 };
}
exports.ModeInfo_Single = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.mode !== 0) {
            writer.uint32(8).int32(message.mode);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseModeInfo_Single();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.mode = reader.int32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return { mode: isSet(object.mode) ? (0, signing_1.signModeFromJSON)(object.mode) : 0 };
    },
    toJSON(message) {
        const obj = {};
        message.mode !== undefined && (obj.mode = (0, signing_1.signModeToJSON)(message.mode));
        return obj;
    },
    fromPartial(object) {
        const message = createBaseModeInfo_Single();
        message.mode = object.mode ?? 0;
        return message;
    },
};
function createBaseModeInfo_Multi() {
    return { bitarray: undefined, modeInfos: [] };
}
exports.ModeInfo_Multi = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.bitarray !== undefined) {
            multisig_1.CompactBitArray.encode(message.bitarray, writer.uint32(10).fork()).ldelim();
        }
        for (const v of message.modeInfos) {
            exports.ModeInfo.encode(v, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseModeInfo_Multi();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.bitarray = multisig_1.CompactBitArray.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.modeInfos.push(exports.ModeInfo.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return {
            bitarray: isSet(object.bitarray) ? multisig_1.CompactBitArray.fromJSON(object.bitarray) : undefined,
            modeInfos: Array.isArray(object?.modeInfos) ? object.modeInfos.map((e) => exports.ModeInfo.fromJSON(e)) : [],
        };
    },
    toJSON(message) {
        const obj = {};
        message.bitarray !== undefined
            && (obj.bitarray = message.bitarray ? multisig_1.CompactBitArray.toJSON(message.bitarray) : undefined);
        if (message.modeInfos) {
            obj.modeInfos = message.modeInfos.map((e) => e ? exports.ModeInfo.toJSON(e) : undefined);
        }
        else {
            obj.modeInfos = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = createBaseModeInfo_Multi();
        message.bitarray = (object.bitarray !== undefined && object.bitarray !== null)
            ? multisig_1.CompactBitArray.fromPartial(object.bitarray)
            : undefined;
        message.modeInfos = object.modeInfos?.map((e) => exports.ModeInfo.fromPartial(e)) || [];
        return message;
    },
};
function createBaseFee() {
    return { amount: [], gasLimit: 0, payer: "", granter: "" };
}
exports.Fee = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        for (const v of message.amount) {
            coin_1.Coin.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.gasLimit !== 0) {
            writer.uint32(16).uint64(message.gasLimit);
        }
        if (message.payer !== "") {
            writer.uint32(26).string(message.payer);
        }
        if (message.granter !== "") {
            writer.uint32(34).string(message.granter);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseFee();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.amount.push(coin_1.Coin.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.gasLimit = longToNumber(reader.uint64());
                    break;
                case 3:
                    message.payer = reader.string();
                    break;
                case 4:
                    message.granter = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return {
            amount: Array.isArray(object?.amount) ? object.amount.map((e) => coin_1.Coin.fromJSON(e)) : [],
            gasLimit: isSet(object.gasLimit) ? Number(object.gasLimit) : 0,
            payer: isSet(object.payer) ? String(object.payer) : "",
            granter: isSet(object.granter) ? String(object.granter) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        if (message.amount) {
            obj.amount = message.amount.map((e) => e ? coin_1.Coin.toJSON(e) : undefined);
        }
        else {
            obj.amount = [];
        }
        message.gasLimit !== undefined && (obj.gasLimit = Math.round(message.gasLimit));
        message.payer !== undefined && (obj.payer = message.payer);
        message.granter !== undefined && (obj.granter = message.granter);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseFee();
        message.amount = object.amount?.map((e) => coin_1.Coin.fromPartial(e)) || [];
        message.gasLimit = object.gasLimit ?? 0;
        message.payer = object.payer ?? "";
        message.granter = object.granter ?? "";
        return message;
    },
};
function createBaseTip() {
    return { amount: [], tipper: "" };
}
exports.Tip = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        for (const v of message.amount) {
            coin_1.Coin.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.tipper !== "") {
            writer.uint32(18).string(message.tipper);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseTip();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.amount.push(coin_1.Coin.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.tipper = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return {
            amount: Array.isArray(object?.amount) ? object.amount.map((e) => coin_1.Coin.fromJSON(e)) : [],
            tipper: isSet(object.tipper) ? String(object.tipper) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        if (message.amount) {
            obj.amount = message.amount.map((e) => e ? coin_1.Coin.toJSON(e) : undefined);
        }
        else {
            obj.amount = [];
        }
        message.tipper !== undefined && (obj.tipper = message.tipper);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseTip();
        message.amount = object.amount?.map((e) => coin_1.Coin.fromPartial(e)) || [];
        message.tipper = object.tipper ?? "";
        return message;
    },
};
function createBaseAuxSignerData() {
    return { address: "", signDoc: undefined, mode: 0, sig: new Uint8Array() };
}
exports.AuxSignerData = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.address !== "") {
            writer.uint32(10).string(message.address);
        }
        if (message.signDoc !== undefined) {
            exports.SignDocDirectAux.encode(message.signDoc, writer.uint32(18).fork()).ldelim();
        }
        if (message.mode !== 0) {
            writer.uint32(24).int32(message.mode);
        }
        if (message.sig.length !== 0) {
            writer.uint32(34).bytes(message.sig);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseAuxSignerData();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.address = reader.string();
                    break;
                case 2:
                    message.signDoc = exports.SignDocDirectAux.decode(reader, reader.uint32());
                    break;
                case 3:
                    message.mode = reader.int32();
                    break;
                case 4:
                    message.sig = reader.bytes();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return {
            address: isSet(object.address) ? String(object.address) : "",
            signDoc: isSet(object.signDoc) ? exports.SignDocDirectAux.fromJSON(object.signDoc) : undefined,
            mode: isSet(object.mode) ? (0, signing_1.signModeFromJSON)(object.mode) : 0,
            sig: isSet(object.sig) ? bytesFromBase64(object.sig) : new Uint8Array(),
        };
    },
    toJSON(message) {
        const obj = {};
        message.address !== undefined && (obj.address = message.address);
        message.signDoc !== undefined
            && (obj.signDoc = message.signDoc ? exports.SignDocDirectAux.toJSON(message.signDoc) : undefined);
        message.mode !== undefined && (obj.mode = (0, signing_1.signModeToJSON)(message.mode));
        message.sig !== undefined
            && (obj.sig = base64FromBytes(message.sig !== undefined ? message.sig : new Uint8Array()));
        return obj;
    },
    fromPartial(object) {
        const message = createBaseAuxSignerData();
        message.address = object.address ?? "";
        message.signDoc = (object.signDoc !== undefined && object.signDoc !== null)
            ? exports.SignDocDirectAux.fromPartial(object.signDoc)
            : undefined;
        message.mode = object.mode ?? 0;
        message.sig = object.sig ?? new Uint8Array();
        return message;
    },
};
var globalThis = (() => {
    if (typeof globalThis !== "undefined") {
        return globalThis;
    }
    if (typeof self !== "undefined") {
        return self;
    }
    if (typeof window !== "undefined") {
        return window;
    }
    if (typeof global !== "undefined") {
        return global;
    }
    throw "Unable to locate global object";
})();
function bytesFromBase64(b64) {
    if (globalThis.Buffer) {
        return Uint8Array.from(globalThis.Buffer.from(b64, "base64"));
    }
    else {
        const bin = globalThis.atob(b64);
        const arr = new Uint8Array(bin.length);
        for (let i = 0; i < bin.length; ++i) {
            arr[i] = bin.charCodeAt(i);
        }
        return arr;
    }
}
function base64FromBytes(arr) {
    if (globalThis.Buffer) {
        return globalThis.Buffer.from(arr).toString("base64");
    }
    else {
        const bin = [];
        arr.forEach((byte) => {
            bin.push(String.fromCharCode(byte));
        });
        return globalThis.btoa(bin.join(""));
    }
}
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
    }
    return long.toNumber();
}
if (minimal_1.default.util.Long !== long_1.default) {
    minimal_1.default.util.Long = long_1.default;
    minimal_1.default.configure();
}
function isSet(value) {
    return value !== null && value !== undefined;
}
