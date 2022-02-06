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
    ownedCards: 0,
    CouncilStatus: 0,
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
        for (const v of message.ownedCards) {
            writer.uint64(v);
        }
        writer.ldelim();
        for (const v of message.voteRights) {
            VoteRight.encode(v, writer.uint32(34).fork()).ldelim();
        }
        if (message.CouncilStatus !== 0) {
            writer.uint32(40).int32(message.CouncilStatus);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseUser };
        message.ownedCardSchemes = [];
        message.ownedCards = [];
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
                            message.ownedCards.push(longToNumber(reader.uint64()));
                        }
                    }
                    else {
                        message.ownedCards.push(longToNumber(reader.uint64()));
                    }
                    break;
                case 4:
                    message.voteRights.push(VoteRight.decode(reader, reader.uint32()));
                    break;
                case 5:
                    message.CouncilStatus = reader.int32();
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
        message.ownedCards = [];
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
        if (object.ownedCards !== undefined && object.ownedCards !== null) {
            for (const e of object.ownedCards) {
                message.ownedCards.push(Number(e));
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
        if (message.ownedCards) {
            obj.ownedCards = message.ownedCards.map((e) => e);
        }
        else {
            obj.ownedCards = [];
        }
        if (message.voteRights) {
            obj.voteRights = message.voteRights.map((e) => e ? VoteRight.toJSON(e) : undefined);
        }
        else {
            obj.voteRights = [];
        }
        message.CouncilStatus !== undefined &&
            (obj.CouncilStatus = councilStatusToJSON(message.CouncilStatus));
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseUser };
        message.ownedCardSchemes = [];
        message.ownedCards = [];
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
        if (object.ownedCards !== undefined && object.ownedCards !== null) {
            for (const e of object.ownedCards) {
                message.ownedCards.push(e);
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
