/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
export const protobufPackage = "DecentralCardGame.cardchain.cardchain";
export var Outcome;
(function (Outcome) {
    Outcome[Outcome["AWon"] = 0] = "AWon";
    Outcome[Outcome["BWon"] = 1] = "BWon";
    Outcome[Outcome["Draw"] = 2] = "Draw";
    Outcome[Outcome["Aborted"] = 3] = "Aborted";
    Outcome[Outcome["UNRECOGNIZED"] = -1] = "UNRECOGNIZED";
})(Outcome || (Outcome = {}));
export function outcomeFromJSON(object) {
    switch (object) {
        case 0:
        case "AWon":
            return Outcome.AWon;
        case 1:
        case "BWon":
            return Outcome.BWon;
        case 2:
        case "Draw":
            return Outcome.Draw;
        case 3:
        case "Aborted":
            return Outcome.Aborted;
        case -1:
        case "UNRECOGNIZED":
        default:
            return Outcome.UNRECOGNIZED;
    }
}
export function outcomeToJSON(object) {
    switch (object) {
        case Outcome.AWon:
            return "AWon";
        case Outcome.BWon:
            return "BWon";
        case Outcome.Draw:
            return "Draw";
        case Outcome.Aborted:
            return "Aborted";
        default:
            return "UNKNOWN";
    }
}
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
const baseMsgSaveCardContent = {
    creator: "",
    cardId: 0,
    notes: "",
    owner: "",
    artist: "",
};
export const MsgSaveCardContent = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.cardId !== 0) {
            writer.uint32(16).uint64(message.cardId);
        }
        if (message.content.length !== 0) {
            writer.uint32(26).bytes(message.content);
        }
        if (message.notes !== "") {
            writer.uint32(34).string(message.notes);
        }
        if (message.owner !== "") {
            writer.uint32(42).string(message.owner);
        }
        if (message.artist !== "") {
            writer.uint32(50).string(message.artist);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgSaveCardContent };
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
                    message.content = reader.bytes();
                    break;
                case 4:
                    message.notes = reader.string();
                    break;
                case 5:
                    message.owner = reader.string();
                    break;
                case 6:
                    message.artist = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgSaveCardContent };
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
        if (object.content !== undefined && object.content !== null) {
            message.content = bytesFromBase64(object.content);
        }
        if (object.notes !== undefined && object.notes !== null) {
            message.notes = String(object.notes);
        }
        else {
            message.notes = "";
        }
        if (object.owner !== undefined && object.owner !== null) {
            message.owner = String(object.owner);
        }
        else {
            message.owner = "";
        }
        if (object.artist !== undefined && object.artist !== null) {
            message.artist = String(object.artist);
        }
        else {
            message.artist = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.cardId !== undefined && (obj.cardId = message.cardId);
        message.content !== undefined &&
            (obj.content = base64FromBytes(message.content !== undefined ? message.content : new Uint8Array()));
        message.notes !== undefined && (obj.notes = message.notes);
        message.owner !== undefined && (obj.owner = message.owner);
        message.artist !== undefined && (obj.artist = message.artist);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgSaveCardContent };
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
        if (object.content !== undefined && object.content !== null) {
            message.content = object.content;
        }
        else {
            message.content = new Uint8Array();
        }
        if (object.notes !== undefined && object.notes !== null) {
            message.notes = object.notes;
        }
        else {
            message.notes = "";
        }
        if (object.owner !== undefined && object.owner !== null) {
            message.owner = object.owner;
        }
        else {
            message.owner = "";
        }
        if (object.artist !== undefined && object.artist !== null) {
            message.artist = object.artist;
        }
        else {
            message.artist = "";
        }
        return message;
    },
};
const baseMsgSaveCardContentResponse = {};
export const MsgSaveCardContentResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgSaveCardContentResponse,
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
            ...baseMsgSaveCardContentResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgSaveCardContentResponse,
        };
        return message;
    },
};
const baseMsgTransferCard = { creator: "", cardId: 0, receiver: "" };
export const MsgTransferCard = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.cardId !== 0) {
            writer.uint32(16).uint64(message.cardId);
        }
        if (message.receiver !== "") {
            writer.uint32(34).string(message.receiver);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgTransferCard };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.cardId = longToNumber(reader.uint64());
                    break;
                case 4:
                    message.receiver = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgTransferCard };
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
        if (object.receiver !== undefined && object.receiver !== null) {
            message.receiver = String(object.receiver);
        }
        else {
            message.receiver = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.cardId !== undefined && (obj.cardId = message.cardId);
        message.receiver !== undefined && (obj.receiver = message.receiver);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgTransferCard };
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
        if (object.receiver !== undefined && object.receiver !== null) {
            message.receiver = object.receiver;
        }
        else {
            message.receiver = "";
        }
        return message;
    },
};
const baseMsgTransferCardResponse = {};
export const MsgTransferCardResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgTransferCardResponse,
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
            ...baseMsgTransferCardResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgTransferCardResponse,
        };
        return message;
    },
};
const baseMsgDonateToCard = {
    creator: "",
    cardId: 0,
    donator: "",
    amount: "",
};
export const MsgDonateToCard = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.cardId !== 0) {
            writer.uint32(16).uint64(message.cardId);
        }
        if (message.donator !== "") {
            writer.uint32(26).string(message.donator);
        }
        if (message.amount !== "") {
            writer.uint32(34).string(message.amount);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgDonateToCard };
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
                    message.donator = reader.string();
                    break;
                case 4:
                    message.amount = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgDonateToCard };
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
        if (object.donator !== undefined && object.donator !== null) {
            message.donator = String(object.donator);
        }
        else {
            message.donator = "";
        }
        if (object.amount !== undefined && object.amount !== null) {
            message.amount = String(object.amount);
        }
        else {
            message.amount = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.cardId !== undefined && (obj.cardId = message.cardId);
        message.donator !== undefined && (obj.donator = message.donator);
        message.amount !== undefined && (obj.amount = message.amount);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgDonateToCard };
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
        if (object.donator !== undefined && object.donator !== null) {
            message.donator = object.donator;
        }
        else {
            message.donator = "";
        }
        if (object.amount !== undefined && object.amount !== null) {
            message.amount = object.amount;
        }
        else {
            message.amount = "";
        }
        return message;
    },
};
const baseMsgDonateToCardResponse = {};
export const MsgDonateToCardResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgDonateToCardResponse,
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
            ...baseMsgDonateToCardResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgDonateToCardResponse,
        };
        return message;
    },
};
const baseMsgAddArtwork = { creator: "", cardId: 0, fullArt: false };
export const MsgAddArtwork = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.cardId !== 0) {
            writer.uint32(16).uint64(message.cardId);
        }
        if (message.image.length !== 0) {
            writer.uint32(26).bytes(message.image);
        }
        if (message.fullArt === true) {
            writer.uint32(32).bool(message.fullArt);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgAddArtwork };
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
                    message.image = reader.bytes();
                    break;
                case 4:
                    message.fullArt = reader.bool();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgAddArtwork };
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
        if (object.image !== undefined && object.image !== null) {
            message.image = bytesFromBase64(object.image);
        }
        if (object.fullArt !== undefined && object.fullArt !== null) {
            message.fullArt = Boolean(object.fullArt);
        }
        else {
            message.fullArt = false;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.cardId !== undefined && (obj.cardId = message.cardId);
        message.image !== undefined &&
            (obj.image = base64FromBytes(message.image !== undefined ? message.image : new Uint8Array()));
        message.fullArt !== undefined && (obj.fullArt = message.fullArt);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgAddArtwork };
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
        if (object.image !== undefined && object.image !== null) {
            message.image = object.image;
        }
        else {
            message.image = new Uint8Array();
        }
        if (object.fullArt !== undefined && object.fullArt !== null) {
            message.fullArt = object.fullArt;
        }
        else {
            message.fullArt = false;
        }
        return message;
    },
};
const baseMsgAddArtworkResponse = {};
export const MsgAddArtworkResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgAddArtworkResponse };
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
        const message = { ...baseMsgAddArtworkResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgAddArtworkResponse };
        return message;
    },
};
const baseMsgSubmitCopyrightProposal = {
    creator: "",
    cardId: 0,
    description: "",
    link: "",
};
export const MsgSubmitCopyrightProposal = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.cardId !== 0) {
            writer.uint32(16).uint64(message.cardId);
        }
        if (message.description !== "") {
            writer.uint32(26).string(message.description);
        }
        if (message.link !== "") {
            writer.uint32(34).string(message.link);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgSubmitCopyrightProposal,
        };
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
                    message.description = reader.string();
                    break;
                case 4:
                    message.link = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseMsgSubmitCopyrightProposal,
        };
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
        if (object.description !== undefined && object.description !== null) {
            message.description = String(object.description);
        }
        else {
            message.description = "";
        }
        if (object.link !== undefined && object.link !== null) {
            message.link = String(object.link);
        }
        else {
            message.link = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.cardId !== undefined && (obj.cardId = message.cardId);
        message.description !== undefined &&
            (obj.description = message.description);
        message.link !== undefined && (obj.link = message.link);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseMsgSubmitCopyrightProposal,
        };
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
        if (object.description !== undefined && object.description !== null) {
            message.description = object.description;
        }
        else {
            message.description = "";
        }
        if (object.link !== undefined && object.link !== null) {
            message.link = object.link;
        }
        else {
            message.link = "";
        }
        return message;
    },
};
const baseMsgSubmitCopyrightProposalResponse = {};
export const MsgSubmitCopyrightProposalResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgSubmitCopyrightProposalResponse,
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
            ...baseMsgSubmitCopyrightProposalResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgSubmitCopyrightProposalResponse,
        };
        return message;
    },
};
const baseMsgChangeArtist = { creator: "", cardID: 0, artist: "" };
export const MsgChangeArtist = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.cardID !== 0) {
            writer.uint32(16).uint64(message.cardID);
        }
        if (message.artist !== "") {
            writer.uint32(26).string(message.artist);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgChangeArtist };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.cardID = longToNumber(reader.uint64());
                    break;
                case 3:
                    message.artist = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgChangeArtist };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.cardID !== undefined && object.cardID !== null) {
            message.cardID = Number(object.cardID);
        }
        else {
            message.cardID = 0;
        }
        if (object.artist !== undefined && object.artist !== null) {
            message.artist = String(object.artist);
        }
        else {
            message.artist = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.cardID !== undefined && (obj.cardID = message.cardID);
        message.artist !== undefined && (obj.artist = message.artist);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgChangeArtist };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.cardID !== undefined && object.cardID !== null) {
            message.cardID = object.cardID;
        }
        else {
            message.cardID = 0;
        }
        if (object.artist !== undefined && object.artist !== null) {
            message.artist = object.artist;
        }
        else {
            message.artist = "";
        }
        return message;
    },
};
const baseMsgChangeArtistResponse = {};
export const MsgChangeArtistResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgChangeArtistResponse,
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
            ...baseMsgChangeArtistResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgChangeArtistResponse,
        };
        return message;
    },
};
const baseMsgRegisterForCouncil = { creator: "" };
export const MsgRegisterForCouncil = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgRegisterForCouncil };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgRegisterForCouncil };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgRegisterForCouncil };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        return message;
    },
};
const baseMsgRegisterForCouncilResponse = {};
export const MsgRegisterForCouncilResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgRegisterForCouncilResponse,
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
            ...baseMsgRegisterForCouncilResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgRegisterForCouncilResponse,
        };
        return message;
    },
};
const baseMsgReportMatch = {
    creator: "",
    playerA: "",
    playerB: "",
    cardsA: 0,
    cardsB: 0,
    outcome: 0,
};
export const MsgReportMatch = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.playerA !== "") {
            writer.uint32(18).string(message.playerA);
        }
        if (message.playerB !== "") {
            writer.uint32(26).string(message.playerB);
        }
        writer.uint32(42).fork();
        for (const v of message.cardsA) {
            writer.uint64(v);
        }
        writer.ldelim();
        writer.uint32(50).fork();
        for (const v of message.cardsB) {
            writer.uint64(v);
        }
        writer.ldelim();
        if (message.outcome !== 0) {
            writer.uint32(56).int32(message.outcome);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgReportMatch };
        message.cardsA = [];
        message.cardsB = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.playerA = reader.string();
                    break;
                case 3:
                    message.playerB = reader.string();
                    break;
                case 5:
                    if ((tag & 7) === 2) {
                        const end2 = reader.uint32() + reader.pos;
                        while (reader.pos < end2) {
                            message.cardsA.push(longToNumber(reader.uint64()));
                        }
                    }
                    else {
                        message.cardsA.push(longToNumber(reader.uint64()));
                    }
                    break;
                case 6:
                    if ((tag & 7) === 2) {
                        const end2 = reader.uint32() + reader.pos;
                        while (reader.pos < end2) {
                            message.cardsB.push(longToNumber(reader.uint64()));
                        }
                    }
                    else {
                        message.cardsB.push(longToNumber(reader.uint64()));
                    }
                    break;
                case 7:
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
        const message = { ...baseMsgReportMatch };
        message.cardsA = [];
        message.cardsB = [];
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
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
        if (object.cardsA !== undefined && object.cardsA !== null) {
            for (const e of object.cardsA) {
                message.cardsA.push(Number(e));
            }
        }
        if (object.cardsB !== undefined && object.cardsB !== null) {
            for (const e of object.cardsB) {
                message.cardsB.push(Number(e));
            }
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
        message.creator !== undefined && (obj.creator = message.creator);
        message.playerA !== undefined && (obj.playerA = message.playerA);
        message.playerB !== undefined && (obj.playerB = message.playerB);
        if (message.cardsA) {
            obj.cardsA = message.cardsA.map((e) => e);
        }
        else {
            obj.cardsA = [];
        }
        if (message.cardsB) {
            obj.cardsB = message.cardsB.map((e) => e);
        }
        else {
            obj.cardsB = [];
        }
        message.outcome !== undefined &&
            (obj.outcome = outcomeToJSON(message.outcome));
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgReportMatch };
        message.cardsA = [];
        message.cardsB = [];
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
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
        if (object.cardsA !== undefined && object.cardsA !== null) {
            for (const e of object.cardsA) {
                message.cardsA.push(e);
            }
        }
        if (object.cardsB !== undefined && object.cardsB !== null) {
            for (const e of object.cardsB) {
                message.cardsB.push(e);
            }
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
const baseMsgReportMatchResponse = { matchId: 0 };
export const MsgReportMatchResponse = {
    encode(message, writer = Writer.create()) {
        if (message.matchId !== 0) {
            writer.uint32(8).uint64(message.matchId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgReportMatchResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.matchId = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgReportMatchResponse };
        if (object.matchId !== undefined && object.matchId !== null) {
            message.matchId = Number(object.matchId);
        }
        else {
            message.matchId = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.matchId !== undefined && (obj.matchId = message.matchId);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgReportMatchResponse };
        if (object.matchId !== undefined && object.matchId !== null) {
            message.matchId = object.matchId;
        }
        else {
            message.matchId = 0;
        }
        return message;
    },
};
const baseMsgSubmitMatchReporterProposal = {
    creator: "",
    reporter: "",
    deposit: "",
    description: "",
};
export const MsgSubmitMatchReporterProposal = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.reporter !== "") {
            writer.uint32(18).string(message.reporter);
        }
        if (message.deposit !== "") {
            writer.uint32(26).string(message.deposit);
        }
        if (message.description !== "") {
            writer.uint32(34).string(message.description);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgSubmitMatchReporterProposal,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.reporter = reader.string();
                    break;
                case 3:
                    message.deposit = reader.string();
                    break;
                case 4:
                    message.description = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseMsgSubmitMatchReporterProposal,
        };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.reporter !== undefined && object.reporter !== null) {
            message.reporter = String(object.reporter);
        }
        else {
            message.reporter = "";
        }
        if (object.deposit !== undefined && object.deposit !== null) {
            message.deposit = String(object.deposit);
        }
        else {
            message.deposit = "";
        }
        if (object.description !== undefined && object.description !== null) {
            message.description = String(object.description);
        }
        else {
            message.description = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.reporter !== undefined && (obj.reporter = message.reporter);
        message.deposit !== undefined && (obj.deposit = message.deposit);
        message.description !== undefined &&
            (obj.description = message.description);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseMsgSubmitMatchReporterProposal,
        };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.reporter !== undefined && object.reporter !== null) {
            message.reporter = object.reporter;
        }
        else {
            message.reporter = "";
        }
        if (object.deposit !== undefined && object.deposit !== null) {
            message.deposit = object.deposit;
        }
        else {
            message.deposit = "";
        }
        if (object.description !== undefined && object.description !== null) {
            message.description = object.description;
        }
        else {
            message.description = "";
        }
        return message;
    },
};
const baseMsgSubmitMatchReporterProposalResponse = {};
export const MsgSubmitMatchReporterProposalResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgSubmitMatchReporterProposalResponse,
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
            ...baseMsgSubmitMatchReporterProposalResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgSubmitMatchReporterProposalResponse,
        };
        return message;
    },
};
const baseMsgApointMatchReporter = { creator: "", reporter: "" };
export const MsgApointMatchReporter = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.reporter !== "") {
            writer.uint32(18).string(message.reporter);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgApointMatchReporter };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.reporter = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgApointMatchReporter };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.reporter !== undefined && object.reporter !== null) {
            message.reporter = String(object.reporter);
        }
        else {
            message.reporter = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.reporter !== undefined && (obj.reporter = message.reporter);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgApointMatchReporter };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.reporter !== undefined && object.reporter !== null) {
            message.reporter = object.reporter;
        }
        else {
            message.reporter = "";
        }
        return message;
    },
};
const baseMsgApointMatchReporterResponse = {};
export const MsgApointMatchReporterResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgApointMatchReporterResponse,
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
            ...baseMsgApointMatchReporterResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgApointMatchReporterResponse,
        };
        return message;
    },
};
const baseMsgCreateCollection = {
    creator: "",
    name: "",
    contributors: "",
    story: "",
};
export const MsgCreateCollection = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.name !== "") {
            writer.uint32(18).string(message.name);
        }
        for (const v of message.contributors) {
            writer.uint32(26).string(v);
        }
        if (message.story !== "") {
            writer.uint32(34).string(message.story);
        }
        if (message.artwork.length !== 0) {
            writer.uint32(42).bytes(message.artwork);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCreateCollection };
        message.contributors = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.name = reader.string();
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
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgCreateCollection };
        message.contributors = [];
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.name !== undefined && object.name !== null) {
            message.name = String(object.name);
        }
        else {
            message.name = "";
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
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.name !== undefined && (obj.name = message.name);
        if (message.contributors) {
            obj.contributors = message.contributors.map((e) => e);
        }
        else {
            obj.contributors = [];
        }
        message.story !== undefined && (obj.story = message.story);
        message.artwork !== undefined &&
            (obj.artwork = base64FromBytes(message.artwork !== undefined ? message.artwork : new Uint8Array()));
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgCreateCollection };
        message.contributors = [];
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.name !== undefined && object.name !== null) {
            message.name = object.name;
        }
        else {
            message.name = "";
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
        return message;
    },
};
const baseMsgCreateCollectionResponse = {};
export const MsgCreateCollectionResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgCreateCollectionResponse,
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
            ...baseMsgCreateCollectionResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgCreateCollectionResponse,
        };
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
    SaveCardContent(request) {
        const data = MsgSaveCardContent.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "SaveCardContent", data);
        return promise.then((data) => MsgSaveCardContentResponse.decode(new Reader(data)));
    }
    TransferCard(request) {
        const data = MsgTransferCard.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "TransferCard", data);
        return promise.then((data) => MsgTransferCardResponse.decode(new Reader(data)));
    }
    DonateToCard(request) {
        const data = MsgDonateToCard.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "DonateToCard", data);
        return promise.then((data) => MsgDonateToCardResponse.decode(new Reader(data)));
    }
    AddArtwork(request) {
        const data = MsgAddArtwork.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "AddArtwork", data);
        return promise.then((data) => MsgAddArtworkResponse.decode(new Reader(data)));
    }
    SubmitCopyrightProposal(request) {
        const data = MsgSubmitCopyrightProposal.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "SubmitCopyrightProposal", data);
        return promise.then((data) => MsgSubmitCopyrightProposalResponse.decode(new Reader(data)));
    }
    ChangeArtist(request) {
        const data = MsgChangeArtist.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "ChangeArtist", data);
        return promise.then((data) => MsgChangeArtistResponse.decode(new Reader(data)));
    }
    RegisterForCouncil(request) {
        const data = MsgRegisterForCouncil.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "RegisterForCouncil", data);
        return promise.then((data) => MsgRegisterForCouncilResponse.decode(new Reader(data)));
    }
    ReportMatch(request) {
        const data = MsgReportMatch.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "ReportMatch", data);
        return promise.then((data) => MsgReportMatchResponse.decode(new Reader(data)));
    }
    SubmitMatchReporterProposal(request) {
        const data = MsgSubmitMatchReporterProposal.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "SubmitMatchReporterProposal", data);
        return promise.then((data) => MsgSubmitMatchReporterProposalResponse.decode(new Reader(data)));
    }
    ApointMatchReporter(request) {
        const data = MsgApointMatchReporter.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "ApointMatchReporter", data);
        return promise.then((data) => MsgApointMatchReporterResponse.decode(new Reader(data)));
    }
    CreateCollection(request) {
        const data = MsgCreateCollection.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "CreateCollection", data);
        return promise.then((data) => MsgCreateCollectionResponse.decode(new Reader(data)));
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
