"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.VoteRight = exports.protobufPackage = void 0;
/* eslint-disable */
const long_1 = require("long");
const minimal_1 = require("protobufjs/minimal");
exports.protobufPackage = "DecentralCardGame.cardchain.cardchain";
function createBaseVoteRight() {
    return { cardId: 0, expireBlock: 0 };
}
exports.VoteRight = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.cardId !== 0) {
            writer.uint32(8).uint64(message.cardId);
        }
        if (message.expireBlock !== 0) {
            writer.uint32(16).int64(message.expireBlock);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseVoteRight();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.cardId = longToNumber(reader.uint64());
                    break;
                case 2:
                    message.expireBlock = longToNumber(reader.int64());
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
            cardId: isSet(object.cardId) ? Number(object.cardId) : 0,
            expireBlock: isSet(object.expireBlock) ? Number(object.expireBlock) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        message.cardId !== undefined && (obj.cardId = Math.round(message.cardId));
        message.expireBlock !== undefined && (obj.expireBlock = Math.round(message.expireBlock));
        return obj;
    },
    fromPartial(object) {
        const message = createBaseVoteRight();
        message.cardId = object.cardId ?? 0;
        message.expireBlock = object.expireBlock ?? 0;
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
