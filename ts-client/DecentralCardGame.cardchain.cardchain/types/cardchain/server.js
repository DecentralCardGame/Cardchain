"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Server = exports.protobufPackage = void 0;
/* eslint-disable */
const long_1 = require("long");
const minimal_1 = require("protobufjs/minimal");
exports.protobufPackage = "DecentralCardGame.cardchain.cardchain";
function createBaseServer() {
    return { reporter: "", invalidReports: 0, validReports: 0 };
}
exports.Server = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.reporter !== "") {
            writer.uint32(10).string(message.reporter);
        }
        if (message.invalidReports !== 0) {
            writer.uint32(16).uint64(message.invalidReports);
        }
        if (message.validReports !== 0) {
            writer.uint32(24).uint64(message.validReports);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseServer();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.reporter = reader.string();
                    break;
                case 2:
                    message.invalidReports = longToNumber(reader.uint64());
                    break;
                case 3:
                    message.validReports = longToNumber(reader.uint64());
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
            reporter: isSet(object.reporter) ? String(object.reporter) : "",
            invalidReports: isSet(object.invalidReports) ? Number(object.invalidReports) : 0,
            validReports: isSet(object.validReports) ? Number(object.validReports) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        message.reporter !== undefined && (obj.reporter = message.reporter);
        message.invalidReports !== undefined && (obj.invalidReports = Math.round(message.invalidReports));
        message.validReports !== undefined && (obj.validReports = Math.round(message.validReports));
        return obj;
    },
    fromPartial(object) {
        const message = createBaseServer();
        message.reporter = object.reporter ?? "";
        message.invalidReports = object.invalidReports ?? 0;
        message.validReports = object.validReports ?? 0;
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
