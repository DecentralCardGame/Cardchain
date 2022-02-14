/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { VoteRight } from "../cardchain/vote_right";
export const protobufPackage = "DecentralCardGame.cardchain.cardchain";
export var CouncilStatus;
(function (CouncilStatus) {
    CouncilStatus[CouncilStatus["available"] = 0] = "available";
    CouncilStatus[CouncilStatus["unavailable"] = 1] = "unavailable";
    CouncilStatus[CouncilStatus["openCouncil"] = 2] = "openCouncil";
    CouncilStatus[CouncilStatus["startedCouncil"] = 3] = "startedCouncil";
    CouncilStatus[CouncilStatus["UNRECOGNIZED"] = -1] = "UNRECOGNIZED";
})(CouncilStatus || (CouncilStatus = {}));
export function councilStatusFromJSON(object) {
    switch (object) {
        case 0:
        case "available":
            return CouncilStatus.available;
        case 1:
        case "unavailable":
            return CouncilStatus.unavailable;
        case 2:
        case "openCouncil":
            return CouncilStatus.openCouncil;
        case 3:
        case "startedCouncil":
            return CouncilStatus.startedCouncil;
        case -1:
        case "UNRECOGNIZED":
        default:
            return CouncilStatus.UNRECOGNIZED;
    }
}
export function councilStatusToJSON(object) {
    switch (object) {
        case CouncilStatus.available:
            return "available";
        case CouncilStatus.unavailable:
            return "unavailable";
        case CouncilStatus.openCouncil:
            return "openCouncil";
        case CouncilStatus.startedCouncil:
            return "startedCouncil";
        default:
            return "UNKNOWN";
    }
}
const baseUser = {
    alias: "",
    ownedCardSchemes: 0,
    ownedPrototypes: 0,
    cards: 0,
    CouncilStatus: 0,
    ReportMatches: false,
};
export const User = {
    encode(message, writer = Writer.create()) {
        if (message.alias !== "") {
            writer.uint32(10).string(message.alias);
        }
        writer.uint32(18).fork();
        for (const v of message.ownedCardSchemes) {
            writer.uint64(v);
        }
        writer.ldelim();
        writer.uint32(26).fork();
        for (const v of message.ownedPrototypes) {
            writer.uint64(v);
        }
        writer.ldelim();
        writer.uint32(34).fork();
        for (const v of message.cards) {
            writer.uint64(v);
        }
        writer.ldelim();
        for (const v of message.voteRights) {
            VoteRight.encode(v, writer.uint32(42).fork()).ldelim();
        }
        if (message.CouncilStatus !== 0) {
            writer.uint32(48).int32(message.CouncilStatus);
        }
        if (message.ReportMatches === true) {
            writer.uint32(56).bool(message.ReportMatches);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseUser };
        message.ownedCardSchemes = [];
        message.ownedPrototypes = [];
        message.cards = [];
        message.voteRights = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.alias = reader.string();
                    break;
                case 2:
                    if ((tag & 7) === 2) {
                        const end2 = reader.uint32() + reader.pos;
                        while (reader.pos < end2) {
                            message.ownedCardSchemes.push(longToNumber(reader.uint64()));
                        }
                    }
                    else {
                        message.ownedCardSchemes.push(longToNumber(reader.uint64()));
                    }
                    break;
                case 3:
                    if ((tag & 7) === 2) {
                        const end2 = reader.uint32() + reader.pos;
                        while (reader.pos < end2) {
                            message.ownedPrototypes.push(longToNumber(reader.uint64()));
                        }
                    }
                    else {
                        message.ownedPrototypes.push(longToNumber(reader.uint64()));
                    }
                    break;
                case 4:
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
                case 5:
                    message.voteRights.push(VoteRight.decode(reader, reader.uint32()));
                    break;
                case 6:
                    message.CouncilStatus = reader.int32();
                    break;
                case 7:
                    message.ReportMatches = reader.bool();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseUser };
        message.ownedCardSchemes = [];
        message.ownedPrototypes = [];
        message.cards = [];
        message.voteRights = [];
        if (object.alias !== undefined && object.alias !== null) {
            message.alias = String(object.alias);
        }
        else {
            message.alias = "";
        }
        if (object.ownedCardSchemes !== undefined &&
            object.ownedCardSchemes !== null) {
            for (const e of object.ownedCardSchemes) {
                message.ownedCardSchemes.push(Number(e));
            }
        }
        if (object.ownedPrototypes !== undefined &&
            object.ownedPrototypes !== null) {
            for (const e of object.ownedPrototypes) {
                message.ownedPrototypes.push(Number(e));
            }
        }
        if (object.cards !== undefined && object.cards !== null) {
            for (const e of object.cards) {
                message.cards.push(Number(e));
            }
        }
        if (object.voteRights !== undefined && object.voteRights !== null) {
            for (const e of object.voteRights) {
                message.voteRights.push(VoteRight.fromJSON(e));
            }
        }
        if (object.CouncilStatus !== undefined && object.CouncilStatus !== null) {
            message.CouncilStatus = councilStatusFromJSON(object.CouncilStatus);
        }
        else {
            message.CouncilStatus = 0;
        }
        if (object.ReportMatches !== undefined && object.ReportMatches !== null) {
            message.ReportMatches = Boolean(object.ReportMatches);
        }
        else {
            message.ReportMatches = false;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.alias !== undefined && (obj.alias = message.alias);
        if (message.ownedCardSchemes) {
            obj.ownedCardSchemes = message.ownedCardSchemes.map((e) => e);
        }
        else {
            obj.ownedCardSchemes = [];
        }
        if (message.ownedPrototypes) {
            obj.ownedPrototypes = message.ownedPrototypes.map((e) => e);
        }
        else {
            obj.ownedPrototypes = [];
        }
        if (message.cards) {
            obj.cards = message.cards.map((e) => e);
        }
        else {
            obj.cards = [];
        }
        if (message.voteRights) {
            obj.voteRights = message.voteRights.map((e) => e ? VoteRight.toJSON(e) : undefined);
        }
        else {
            obj.voteRights = [];
        }
        message.CouncilStatus !== undefined &&
            (obj.CouncilStatus = councilStatusToJSON(message.CouncilStatus));
        message.ReportMatches !== undefined &&
            (obj.ReportMatches = message.ReportMatches);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseUser };
        message.ownedCardSchemes = [];
        message.ownedPrototypes = [];
        message.cards = [];
        message.voteRights = [];
        if (object.alias !== undefined && object.alias !== null) {
            message.alias = object.alias;
        }
        else {
            message.alias = "";
        }
        if (object.ownedCardSchemes !== undefined &&
            object.ownedCardSchemes !== null) {
            for (const e of object.ownedCardSchemes) {
                message.ownedCardSchemes.push(e);
            }
        }
        if (object.ownedPrototypes !== undefined &&
            object.ownedPrototypes !== null) {
            for (const e of object.ownedPrototypes) {
                message.ownedPrototypes.push(e);
            }
        }
        if (object.cards !== undefined && object.cards !== null) {
            for (const e of object.cards) {
                message.cards.push(e);
            }
        }
        if (object.voteRights !== undefined && object.voteRights !== null) {
            for (const e of object.voteRights) {
                message.voteRights.push(VoteRight.fromPartial(e));
            }
        }
        if (object.CouncilStatus !== undefined && object.CouncilStatus !== null) {
            message.CouncilStatus = object.CouncilStatus;
        }
        else {
            message.CouncilStatus = 0;
        }
        if (object.ReportMatches !== undefined && object.ReportMatches !== null) {
            message.ReportMatches = object.ReportMatches;
        }
        else {
            message.ReportMatches = false;
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
