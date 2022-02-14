/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "DecentralCardGame.cardchain.cardchain";
export var CStatus;
(function (CStatus) {
    CStatus[CStatus["design"] = 0] = "design";
    CStatus[CStatus["finalized"] = 1] = "finalized";
    CStatus[CStatus["active"] = 2] = "active";
    CStatus[CStatus["archived"] = 3] = "archived";
    CStatus[CStatus["UNRECOGNIZED"] = -1] = "UNRECOGNIZED";
})(CStatus || (CStatus = {}));
export function cStatusFromJSON(object) {
    switch (object) {
        case 0:
        case "design":
            return CStatus.design;
        case 1:
        case "finalized":
            return CStatus.finalized;
        case 2:
        case "active":
            return CStatus.active;
        case 3:
        case "archived":
            return CStatus.archived;
        case -1:
        case "UNRECOGNIZED":
        default:
            return CStatus.UNRECOGNIZED;
    }
}
export function cStatusToJSON(object) {
    switch (object) {
        case CStatus.design:
            return "design";
        case CStatus.finalized:
            return "finalized";
        case CStatus.active:
            return "active";
        case CStatus.archived:
            return "archived";
        default:
            return "UNKNOWN";
    }
}
const baseCollection = {
    name: "",
    cards: 0,
    contributors: "",
    story: "",
    status: 0,
    timeStamp: 0,
};
export const Collection = {
    encode(message, writer = Writer.create()) {
        if (message.name !== "") {
            writer.uint32(10).string(message.name);
        }
        writer.uint32(18).fork();
        for (const v of message.cards) {
            writer.uint64(v);
        }
        writer.ldelim();
        for (const v of message.contributors) {
            writer.uint32(26).string(v);
        }
        if (message.story !== "") {
            writer.uint32(34).string(message.story);
        }
        if (message.artwork.length !== 0) {
            writer.uint32(42).bytes(message.artwork);
        }
        if (message.status !== 0) {
            writer.uint32(48).int32(message.status);
        }
        if (message.timeStamp !== 0) {
            writer.uint32(56).int64(message.timeStamp);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseCollection };
        message.cards = [];
        message.contributors = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.name = reader.string();
                    break;
                case 2:
                    if ((tag & 7) === 2) {
                        const end2 = reader.uint32() + reader.pos;
                        while (reader.pos < end2) {
                            message.cards.push(longToNumber(reader.uint64()));
                        }
                    }
                    else {
                        message.cards.push(longToNumber(reader.uint64()));
                    }
                    break;
                case 3:
                    message.contributors.push(reader.string());
                    break;
                case 4:
                    message.story = reader.string();
                    break;
                case 5:
                    message.artwork = reader.bytes();
                    break;
                case 6:
                    message.status = reader.int32();
                    break;
                case 7:
                    message.timeStamp = longToNumber(reader.int64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseCollection };
        message.cards = [];
        message.contributors = [];
        if (object.name !== undefined && object.name !== null) {
            message.name = String(object.name);
        }
        else {
            message.name = "";
        }
        if (object.cards !== undefined && object.cards !== null) {
            for (const e of object.cards) {
                message.cards.push(Number(e));
            }
        }
        if (object.contributors !== undefined && object.contributors !== null) {
            for (const e of object.contributors) {
                message.contributors.push(String(e));
            }
        }
        if (object.story !== undefined && object.story !== null) {
            message.story = String(object.story);
        }
        else {
            message.story = "";
        }
        if (object.artwork !== undefined && object.artwork !== null) {
            message.artwork = bytesFromBase64(object.artwork);
        }
        if (object.status !== undefined && object.status !== null) {
            message.status = cStatusFromJSON(object.status);
        }
        else {
            message.status = 0;
        }
        if (object.timeStamp !== undefined && object.timeStamp !== null) {
            message.timeStamp = Number(object.timeStamp);
        }
        else {
            message.timeStamp = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.name !== undefined && (obj.name = message.name);
        if (message.cards) {
            obj.cards = message.cards.map((e) => e);
        }
        else {
            obj.cards = [];
        }
        if (message.contributors) {
            obj.contributors = message.contributors.map((e) => e);
        }
        else {
            obj.contributors = [];
        }
        message.story !== undefined && (obj.story = message.story);
        message.artwork !== undefined &&
            (obj.artwork = base64FromBytes(message.artwork !== undefined ? message.artwork : new Uint8Array()));
        message.status !== undefined &&
            (obj.status = cStatusToJSON(message.status));
        message.timeStamp !== undefined && (obj.timeStamp = message.timeStamp);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseCollection };
        message.cards = [];
        message.contributors = [];
        if (object.name !== undefined && object.name !== null) {
            message.name = object.name;
        }
        else {
            message.name = "";
        }
        if (object.cards !== undefined && object.cards !== null) {
            for (const e of object.cards) {
                message.cards.push(e);
            }
        }
        if (object.contributors !== undefined && object.contributors !== null) {
            for (const e of object.contributors) {
                message.contributors.push(e);
            }
        }
        if (object.story !== undefined && object.story !== null) {
            message.story = object.story;
        }
        else {
            message.story = "";
        }
        if (object.artwork !== undefined && object.artwork !== null) {
            message.artwork = object.artwork;
        }
        else {
            message.artwork = new Uint8Array();
        }
        if (object.status !== undefined && object.status !== null) {
            message.status = object.status;
        }
        else {
            message.status = 0;
        }
        if (object.timeStamp !== undefined && object.timeStamp !== null) {
            message.timeStamp = object.timeStamp;
        }
        else {
            message.timeStamp = 0;
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
const atob = globalThis.atob ||
    ((b64) => globalThis.Buffer.from(b64, "base64").toString("binary"));
function bytesFromBase64(b64) {
    const bin = atob(b64);
    const arr = new Uint8Array(bin.length);
    for (let i = 0; i < bin.length; ++i) {
        arr[i] = bin.charCodeAt(i);
    }
    return arr;
}
const btoa = globalThis.btoa ||
    ((bin) => globalThis.Buffer.from(bin, "binary").toString("base64"));
function base64FromBytes(arr) {
    const bin = [];
    for (let i = 0; i < arr.byteLength; ++i) {
        bin.push(String.fromCharCode(arr[i]));
    }
    return btoa(bin.join(""));
}
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
