"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Params = exports.ModuleAccount = exports.BaseAccount = exports.protobufPackage = void 0;
/* eslint-disable */
const long_1 = require("long");
const minimal_1 = require("protobufjs/minimal");
const any_1 = require("../../../google/protobuf/any");
exports.protobufPackage = "cosmos.auth.v1beta1";
function createBaseBaseAccount() {
    return { address: "", pubKey: undefined, accountNumber: 0, sequence: 0 };
}
exports.BaseAccount = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.address !== "") {
            writer.uint32(10).string(message.address);
        }
        if (message.pubKey !== undefined) {
            any_1.Any.encode(message.pubKey, writer.uint32(18).fork()).ldelim();
        }
        if (message.accountNumber !== 0) {
            writer.uint32(24).uint64(message.accountNumber);
        }
        if (message.sequence !== 0) {
            writer.uint32(32).uint64(message.sequence);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseBaseAccount();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.address = reader.string();
                    break;
                case 2:
                    message.pubKey = any_1.Any.decode(reader, reader.uint32());
                    break;
                case 3:
                    message.accountNumber = longToNumber(reader.uint64());
                    break;
                case 4:
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
            address: isSet(object.address) ? String(object.address) : "",
            pubKey: isSet(object.pubKey) ? any_1.Any.fromJSON(object.pubKey) : undefined,
            accountNumber: isSet(object.accountNumber) ? Number(object.accountNumber) : 0,
            sequence: isSet(object.sequence) ? Number(object.sequence) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        message.address !== undefined && (obj.address = message.address);
        message.pubKey !== undefined && (obj.pubKey = message.pubKey ? any_1.Any.toJSON(message.pubKey) : undefined);
        message.accountNumber !== undefined && (obj.accountNumber = Math.round(message.accountNumber));
        message.sequence !== undefined && (obj.sequence = Math.round(message.sequence));
        return obj;
    },
    fromPartial(object) {
        const message = createBaseBaseAccount();
        message.address = object.address ?? "";
        message.pubKey = (object.pubKey !== undefined && object.pubKey !== null)
            ? any_1.Any.fromPartial(object.pubKey)
            : undefined;
        message.accountNumber = object.accountNumber ?? 0;
        message.sequence = object.sequence ?? 0;
        return message;
    },
};
function createBaseModuleAccount() {
    return { baseAccount: undefined, name: "", permissions: [] };
}
exports.ModuleAccount = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.baseAccount !== undefined) {
            exports.BaseAccount.encode(message.baseAccount, writer.uint32(10).fork()).ldelim();
        }
        if (message.name !== "") {
            writer.uint32(18).string(message.name);
        }
        for (const v of message.permissions) {
            writer.uint32(26).string(v);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseModuleAccount();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.baseAccount = exports.BaseAccount.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.name = reader.string();
                    break;
                case 3:
                    message.permissions.push(reader.string());
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
            baseAccount: isSet(object.baseAccount) ? exports.BaseAccount.fromJSON(object.baseAccount) : undefined,
            name: isSet(object.name) ? String(object.name) : "",
            permissions: Array.isArray(object?.permissions) ? object.permissions.map((e) => String(e)) : [],
        };
    },
    toJSON(message) {
        const obj = {};
        message.baseAccount !== undefined
            && (obj.baseAccount = message.baseAccount ? exports.BaseAccount.toJSON(message.baseAccount) : undefined);
        message.name !== undefined && (obj.name = message.name);
        if (message.permissions) {
            obj.permissions = message.permissions.map((e) => e);
        }
        else {
            obj.permissions = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = createBaseModuleAccount();
        message.baseAccount = (object.baseAccount !== undefined && object.baseAccount !== null)
            ? exports.BaseAccount.fromPartial(object.baseAccount)
            : undefined;
        message.name = object.name ?? "";
        message.permissions = object.permissions?.map((e) => e) || [];
        return message;
    },
};
function createBaseParams() {
    return {
        maxMemoCharacters: 0,
        txSigLimit: 0,
        txSizeCostPerByte: 0,
        sigVerifyCostEd25519: 0,
        sigVerifyCostSecp256k1: 0,
    };
}
exports.Params = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.maxMemoCharacters !== 0) {
            writer.uint32(8).uint64(message.maxMemoCharacters);
        }
        if (message.txSigLimit !== 0) {
            writer.uint32(16).uint64(message.txSigLimit);
        }
        if (message.txSizeCostPerByte !== 0) {
            writer.uint32(24).uint64(message.txSizeCostPerByte);
        }
        if (message.sigVerifyCostEd25519 !== 0) {
            writer.uint32(32).uint64(message.sigVerifyCostEd25519);
        }
        if (message.sigVerifyCostSecp256k1 !== 0) {
            writer.uint32(40).uint64(message.sigVerifyCostSecp256k1);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseParams();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.maxMemoCharacters = longToNumber(reader.uint64());
                    break;
                case 2:
                    message.txSigLimit = longToNumber(reader.uint64());
                    break;
                case 3:
                    message.txSizeCostPerByte = longToNumber(reader.uint64());
                    break;
                case 4:
                    message.sigVerifyCostEd25519 = longToNumber(reader.uint64());
                    break;
                case 5:
                    message.sigVerifyCostSecp256k1 = longToNumber(reader.uint64());
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
            maxMemoCharacters: isSet(object.maxMemoCharacters) ? Number(object.maxMemoCharacters) : 0,
            txSigLimit: isSet(object.txSigLimit) ? Number(object.txSigLimit) : 0,
            txSizeCostPerByte: isSet(object.txSizeCostPerByte) ? Number(object.txSizeCostPerByte) : 0,
            sigVerifyCostEd25519: isSet(object.sigVerifyCostEd25519) ? Number(object.sigVerifyCostEd25519) : 0,
            sigVerifyCostSecp256k1: isSet(object.sigVerifyCostSecp256k1) ? Number(object.sigVerifyCostSecp256k1) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        message.maxMemoCharacters !== undefined && (obj.maxMemoCharacters = Math.round(message.maxMemoCharacters));
        message.txSigLimit !== undefined && (obj.txSigLimit = Math.round(message.txSigLimit));
        message.txSizeCostPerByte !== undefined && (obj.txSizeCostPerByte = Math.round(message.txSizeCostPerByte));
        message.sigVerifyCostEd25519 !== undefined && (obj.sigVerifyCostEd25519 = Math.round(message.sigVerifyCostEd25519));
        message.sigVerifyCostSecp256k1 !== undefined
            && (obj.sigVerifyCostSecp256k1 = Math.round(message.sigVerifyCostSecp256k1));
        return obj;
    },
    fromPartial(object) {
        const message = createBaseParams();
        message.maxMemoCharacters = object.maxMemoCharacters ?? 0;
        message.txSigLimit = object.txSigLimit ?? 0;
        message.txSizeCostPerByte = object.txSizeCostPerByte ?? 0;
        message.sigVerifyCostEd25519 = object.sigVerifyCostEd25519 ?? 0;
        message.sigVerifyCostSecp256k1 = object.sigVerifyCostSecp256k1 ?? 0;
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
