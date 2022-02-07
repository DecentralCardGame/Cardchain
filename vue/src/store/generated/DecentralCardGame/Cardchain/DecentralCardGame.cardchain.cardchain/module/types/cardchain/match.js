/* eslint-disable */
import { outcomeFromJSON, outcomeToJSON } from "../cardchain/tx";
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "DecentralCardGame.cardchain.cardchain";
const baseMatch = {
    timestamp: 0,
    reporter: "",
    playerA: "",
    playerB: "",
    outcome: 0,
};
export const Match = {
    encode(message, writer = Writer.create()) {
        if (message.timestamp !== 0) {
            writer.uint32(8).uint64(message.timestamp);
        }
        if (message.reporter !== "") {
            writer.uint32(18).string(message.reporter);
        }
        if (message.playerA !== "") {
            writer.uint32(26).string(message.playerA);
        }
        if (message.playerB !== "") {
            writer.uint32(34).string(message.playerB);
        }
        if (message.outcome !== 0) {
            writer.uint32(40).int32(message.outcome);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMatch };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.timestamp = longToNumber(reader.uint64());
                    break;
                case 2:
                    message.reporter = reader.string();
                    break;
                case 3:
                    message.playerA = reader.string();
                    break;
                case 4:
                    message.playerB = reader.string();
                    break;
                case 5:
                    message.outcome = reader.int32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMatch };
        if (object.timestamp !== undefined && object.timestamp !== null) {
            message.timestamp = Number(object.timestamp);
        }
        else {
            message.timestamp = 0;
        }
        if (object.reporter !== undefined && object.reporter !== null) {
            message.reporter = String(object.reporter);
        }
        else {
            message.reporter = "";
        }
        if (object.playerA !== undefined && object.playerA !== null) {
            message.playerA = String(object.playerA);
        }
        else {
            message.playerA = "";
        }
        if (object.playerB !== undefined && object.playerB !== null) {
            message.playerB = String(object.playerB);
        }
        else {
            message.playerB = "";
        }
        if (object.outcome !== undefined && object.outcome !== null) {
            message.outcome = outcomeFromJSON(object.outcome);
        }
        else {
            message.outcome = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.timestamp !== undefined && (obj.timestamp = message.timestamp);
        message.reporter !== undefined && (obj.reporter = message.reporter);
        message.playerA !== undefined && (obj.playerA = message.playerA);
        message.playerB !== undefined && (obj.playerB = message.playerB);
        message.outcome !== undefined &&
            (obj.outcome = outcomeToJSON(message.outcome));
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMatch };
        if (object.timestamp !== undefined && object.timestamp !== null) {
            message.timestamp = object.timestamp;
        }
        else {
            message.timestamp = 0;
        }
        if (object.reporter !== undefined && object.reporter !== null) {
            message.reporter = object.reporter;
        }
        else {
            message.reporter = "";
        }
        if (object.playerA !== undefined && object.playerA !== null) {
            message.playerA = object.playerA;
        }
        else {
            message.playerA = "";
        }
        if (object.playerB !== undefined && object.playerB !== null) {
            message.playerB = object.playerB;
        }
        else {
            message.playerB = "";
        }
        if (object.outcome !== undefined && object.outcome !== null) {
            message.outcome = object.outcome;
        }
        else {
            message.outcome = 0;
        }
        return message;
    },
};
var globalThis = (() => {
    if (typeof globalThis !== "undefined")
        return globalThis;
    if (typeof self !== "undefined")
        return self;
    if (typeof window !== "undefined")
        return window;
    if (typeof global !== "undefined")
        return global;
    throw "Unable to locate global object";
})();
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}
