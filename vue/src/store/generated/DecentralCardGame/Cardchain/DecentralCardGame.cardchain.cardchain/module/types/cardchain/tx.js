/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
export const protobufPackage = "DecentralCardGame.cardchain.cardchain";
const baseMsgCreateuser = { creator: "", newUser: "", alias: "" };
export const MsgCreateuser = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.newUser !== "") {
            writer.uint32(18).string(message.newUser);
        }
        if (message.alias !== "") {
            writer.uint32(26).string(message.alias);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCreateuser };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.newUser = reader.string();
                    break;
                case 3:
                    message.alias = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgCreateuser };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.newUser !== undefined && object.newUser !== null) {
            message.newUser = String(object.newUser);
        }
        else {
            message.newUser = "";
        }
        if (object.alias !== undefined && object.alias !== null) {
            message.alias = String(object.alias);
        }
        else {
            message.alias = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.newUser !== undefined && (obj.newUser = message.newUser);
        message.alias !== undefined && (obj.alias = message.alias);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgCreateuser };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.newUser !== undefined && object.newUser !== null) {
            message.newUser = object.newUser;
        }
        else {
            message.newUser = "";
        }
        if (object.alias !== undefined && object.alias !== null) {
            message.alias = object.alias;
        }
        else {
            message.alias = "";
        }
        return message;
    },
};
const baseMsgCreateuserResponse = {};
export const MsgCreateuserResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCreateuserResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        const message = { ...baseMsgCreateuserResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgCreateuserResponse };
        return message;
    },
};
const baseMsgBuyCardScheme = { creator: "", bid: "", buyer: "" };
export const MsgBuyCardScheme = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.bid !== "") {
            writer.uint32(18).string(message.bid);
        }
        if (message.buyer !== "") {
            writer.uint32(26).string(message.buyer);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgBuyCardScheme };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.bid = reader.string();
                    break;
                case 3:
                    message.buyer = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgBuyCardScheme };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.bid !== undefined && object.bid !== null) {
            message.bid = String(object.bid);
        }
        else {
            message.bid = "";
        }
        if (object.buyer !== undefined && object.buyer !== null) {
            message.buyer = String(object.buyer);
        }
        else {
            message.buyer = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.bid !== undefined && (obj.bid = message.bid);
        message.buyer !== undefined && (obj.buyer = message.buyer);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgBuyCardScheme };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.bid !== undefined && object.bid !== null) {
            message.bid = object.bid;
        }
        else {
            message.bid = "";
        }
        if (object.buyer !== undefined && object.buyer !== null) {
            message.buyer = object.buyer;
        }
        else {
            message.buyer = "";
        }
        return message;
    },
};
const baseMsgBuyCardSchemeResponse = {};
export const MsgBuyCardSchemeResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgBuyCardSchemeResponse,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        const message = {
            ...baseMsgBuyCardSchemeResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgBuyCardSchemeResponse,
        };
        return message;
    },
};
const baseMsgVoteCard = {
    creator: "",
    cardId: 0,
    voteType: "",
    voter: "",
};
export const MsgVoteCard = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.cardId !== 0) {
            writer.uint32(16).uint64(message.cardId);
        }
        if (message.voteType !== "") {
            writer.uint32(26).string(message.voteType);
        }
        if (message.voter !== "") {
            writer.uint32(34).string(message.voter);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgVoteCard };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.cardId = longToNumber(reader.uint64());
                    break;
                case 3:
                    message.voteType = reader.string();
                    break;
                case 4:
                    message.voter = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgVoteCard };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.cardId !== undefined && object.cardId !== null) {
            message.cardId = Number(object.cardId);
        }
        else {
            message.cardId = 0;
        }
        if (object.voteType !== undefined && object.voteType !== null) {
            message.voteType = String(object.voteType);
        }
        else {
            message.voteType = "";
        }
        if (object.voter !== undefined && object.voter !== null) {
            message.voter = String(object.voter);
        }
        else {
            message.voter = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.cardId !== undefined && (obj.cardId = message.cardId);
        message.voteType !== undefined && (obj.voteType = message.voteType);
        message.voter !== undefined && (obj.voter = message.voter);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgVoteCard };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.cardId !== undefined && object.cardId !== null) {
            message.cardId = object.cardId;
        }
        else {
            message.cardId = 0;
        }
        if (object.voteType !== undefined && object.voteType !== null) {
            message.voteType = object.voteType;
        }
        else {
            message.voteType = "";
        }
        if (object.voter !== undefined && object.voter !== null) {
            message.voter = object.voter;
        }
        else {
            message.voter = "";
        }
        return message;
    },
};
const baseMsgVoteCardResponse = {};
export const MsgVoteCardResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgVoteCardResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        const message = { ...baseMsgVoteCardResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgVoteCardResponse };
        return message;
    },
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    Createuser(request) {
        const data = MsgCreateuser.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "Createuser", data);
        return promise.then((data) => MsgCreateuserResponse.decode(new Reader(data)));
    }
    BuyCardScheme(request) {
        const data = MsgBuyCardScheme.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "BuyCardScheme", data);
        return promise.then((data) => MsgBuyCardSchemeResponse.decode(new Reader(data)));
    }
    VoteCard(request) {
        const data = MsgVoteCard.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "VoteCard", data);
        return promise.then((data) => MsgVoteCardResponse.decode(new Reader(data)));
    }
}
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
