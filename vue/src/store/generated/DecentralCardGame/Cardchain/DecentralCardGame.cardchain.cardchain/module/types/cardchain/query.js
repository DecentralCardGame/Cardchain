/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Params } from "../cardchain/params";
import { VoteRight } from "../cardchain/vote_right";
export const protobufPackage = "DecentralCardGame.cardchain.cardchain";
const baseQueryParamsRequest = {};
export const QueryParamsRequest = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryParamsRequest };
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
        const message = { ...baseQueryParamsRequest };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseQueryParamsRequest };
        return message;
    },
};
const baseQueryParamsResponse = {};
export const QueryParamsResponse = {
    encode(message, writer = Writer.create()) {
        if (message.params !== undefined) {
            Params.encode(message.params, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryParamsResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.params = Params.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryParamsResponse };
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromJSON(object.params);
        }
        else {
            message.params = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.params !== undefined &&
            (obj.params = message.params ? Params.toJSON(message.params) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryParamsResponse };
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromPartial(object.params);
        }
        else {
            message.params = undefined;
        }
        return message;
    },
};
const baseQueryQCardRequest = { cardId: "" };
export const QueryQCardRequest = {
    encode(message, writer = Writer.create()) {
        if (message.cardId !== "") {
            writer.uint32(10).string(message.cardId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryQCardRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.cardId = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryQCardRequest };
        if (object.cardId !== undefined && object.cardId !== null) {
            message.cardId = String(object.cardId);
        }
        else {
            message.cardId = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.cardId !== undefined && (obj.cardId = message.cardId);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryQCardRequest };
        if (object.cardId !== undefined && object.cardId !== null) {
            message.cardId = object.cardId;
        }
        else {
            message.cardId = "";
        }
        return message;
    },
};
const baseQueryQCardResponse = {
    owner: "",
    content: "",
    image: "",
    fullArt: "",
    notes: "",
    status: "",
    votePool: "",
    fairEnoughVotes: 0,
    overpoweredVotes: 0,
    underpoweredVotes: 0,
    inappropriateVotes: 0,
    nerflevel: 0,
};
export const QueryQCardResponse = {
    encode(message, writer = Writer.create()) {
        if (message.owner !== "") {
            writer.uint32(10).string(message.owner);
        }
        if (message.content !== "") {
            writer.uint32(18).string(message.content);
        }
        if (message.image !== "") {
            writer.uint32(26).string(message.image);
        }
        if (message.fullArt !== "") {
            writer.uint32(34).string(message.fullArt);
        }
        if (message.notes !== "") {
            writer.uint32(42).string(message.notes);
        }
        if (message.status !== "") {
            writer.uint32(50).string(message.status);
        }
        if (message.votePool !== "") {
            writer.uint32(58).string(message.votePool);
        }
        if (message.fairEnoughVotes !== 0) {
            writer.uint32(64).uint64(message.fairEnoughVotes);
        }
        if (message.overpoweredVotes !== 0) {
            writer.uint32(72).uint64(message.overpoweredVotes);
        }
        if (message.underpoweredVotes !== 0) {
            writer.uint32(80).uint64(message.underpoweredVotes);
        }
        if (message.inappropriateVotes !== 0) {
            writer.uint32(88).uint64(message.inappropriateVotes);
        }
        if (message.nerflevel !== 0) {
            writer.uint32(96).int64(message.nerflevel);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryQCardResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.owner = reader.string();
                    break;
                case 2:
                    message.content = reader.string();
                    break;
                case 3:
                    message.image = reader.string();
                    break;
                case 4:
                    message.fullArt = reader.string();
                    break;
                case 5:
                    message.notes = reader.string();
                    break;
                case 6:
                    message.status = reader.string();
                    break;
                case 7:
                    message.votePool = reader.string();
                    break;
                case 8:
                    message.fairEnoughVotes = longToNumber(reader.uint64());
                    break;
                case 9:
                    message.overpoweredVotes = longToNumber(reader.uint64());
                    break;
                case 10:
                    message.underpoweredVotes = longToNumber(reader.uint64());
                    break;
                case 11:
                    message.inappropriateVotes = longToNumber(reader.uint64());
                    break;
                case 12:
                    message.nerflevel = longToNumber(reader.int64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryQCardResponse };
        if (object.owner !== undefined && object.owner !== null) {
            message.owner = String(object.owner);
        }
        else {
            message.owner = "";
        }
        if (object.content !== undefined && object.content !== null) {
            message.content = String(object.content);
        }
        else {
            message.content = "";
        }
        if (object.image !== undefined && object.image !== null) {
            message.image = String(object.image);
        }
        else {
            message.image = "";
        }
        if (object.fullArt !== undefined && object.fullArt !== null) {
            message.fullArt = String(object.fullArt);
        }
        else {
            message.fullArt = "";
        }
        if (object.notes !== undefined && object.notes !== null) {
            message.notes = String(object.notes);
        }
        else {
            message.notes = "";
        }
        if (object.status !== undefined && object.status !== null) {
            message.status = String(object.status);
        }
        else {
            message.status = "";
        }
        if (object.votePool !== undefined && object.votePool !== null) {
            message.votePool = String(object.votePool);
        }
        else {
            message.votePool = "";
        }
        if (object.fairEnoughVotes !== undefined &&
            object.fairEnoughVotes !== null) {
            message.fairEnoughVotes = Number(object.fairEnoughVotes);
        }
        else {
            message.fairEnoughVotes = 0;
        }
        if (object.overpoweredVotes !== undefined &&
            object.overpoweredVotes !== null) {
            message.overpoweredVotes = Number(object.overpoweredVotes);
        }
        else {
            message.overpoweredVotes = 0;
        }
        if (object.underpoweredVotes !== undefined &&
            object.underpoweredVotes !== null) {
            message.underpoweredVotes = Number(object.underpoweredVotes);
        }
        else {
            message.underpoweredVotes = 0;
        }
        if (object.inappropriateVotes !== undefined &&
            object.inappropriateVotes !== null) {
            message.inappropriateVotes = Number(object.inappropriateVotes);
        }
        else {
            message.inappropriateVotes = 0;
        }
        if (object.nerflevel !== undefined && object.nerflevel !== null) {
            message.nerflevel = Number(object.nerflevel);
        }
        else {
            message.nerflevel = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.owner !== undefined && (obj.owner = message.owner);
        message.content !== undefined && (obj.content = message.content);
        message.image !== undefined && (obj.image = message.image);
        message.fullArt !== undefined && (obj.fullArt = message.fullArt);
        message.notes !== undefined && (obj.notes = message.notes);
        message.status !== undefined && (obj.status = message.status);
        message.votePool !== undefined && (obj.votePool = message.votePool);
        message.fairEnoughVotes !== undefined &&
            (obj.fairEnoughVotes = message.fairEnoughVotes);
        message.overpoweredVotes !== undefined &&
            (obj.overpoweredVotes = message.overpoweredVotes);
        message.underpoweredVotes !== undefined &&
            (obj.underpoweredVotes = message.underpoweredVotes);
        message.inappropriateVotes !== undefined &&
            (obj.inappropriateVotes = message.inappropriateVotes);
        message.nerflevel !== undefined && (obj.nerflevel = message.nerflevel);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryQCardResponse };
        if (object.owner !== undefined && object.owner !== null) {
            message.owner = object.owner;
        }
        else {
            message.owner = "";
        }
        if (object.content !== undefined && object.content !== null) {
            message.content = object.content;
        }
        else {
            message.content = "";
        }
        if (object.image !== undefined && object.image !== null) {
            message.image = object.image;
        }
        else {
            message.image = "";
        }
        if (object.fullArt !== undefined && object.fullArt !== null) {
            message.fullArt = object.fullArt;
        }
        else {
            message.fullArt = "";
        }
        if (object.notes !== undefined && object.notes !== null) {
            message.notes = object.notes;
        }
        else {
            message.notes = "";
        }
        if (object.status !== undefined && object.status !== null) {
            message.status = object.status;
        }
        else {
            message.status = "";
        }
        if (object.votePool !== undefined && object.votePool !== null) {
            message.votePool = object.votePool;
        }
        else {
            message.votePool = "";
        }
        if (object.fairEnoughVotes !== undefined &&
            object.fairEnoughVotes !== null) {
            message.fairEnoughVotes = object.fairEnoughVotes;
        }
        else {
            message.fairEnoughVotes = 0;
        }
        if (object.overpoweredVotes !== undefined &&
            object.overpoweredVotes !== null) {
            message.overpoweredVotes = object.overpoweredVotes;
        }
        else {
            message.overpoweredVotes = 0;
        }
        if (object.underpoweredVotes !== undefined &&
            object.underpoweredVotes !== null) {
            message.underpoweredVotes = object.underpoweredVotes;
        }
        else {
            message.underpoweredVotes = 0;
        }
        if (object.inappropriateVotes !== undefined &&
            object.inappropriateVotes !== null) {
            message.inappropriateVotes = object.inappropriateVotes;
        }
        else {
            message.inappropriateVotes = 0;
        }
        if (object.nerflevel !== undefined && object.nerflevel !== null) {
            message.nerflevel = object.nerflevel;
        }
        else {
            message.nerflevel = 0;
        }
        return message;
    },
};
const baseQueryQCardContentRequest = { cardId: "" };
export const QueryQCardContentRequest = {
    encode(message, writer = Writer.create()) {
        if (message.cardId !== "") {
            writer.uint32(10).string(message.cardId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryQCardContentRequest,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.cardId = reader.string();
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
            ...baseQueryQCardContentRequest,
        };
        if (object.cardId !== undefined && object.cardId !== null) {
            message.cardId = String(object.cardId);
        }
        else {
            message.cardId = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.cardId !== undefined && (obj.cardId = message.cardId);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryQCardContentRequest,
        };
        if (object.cardId !== undefined && object.cardId !== null) {
            message.cardId = object.cardId;
        }
        else {
            message.cardId = "";
        }
        return message;
    },
};
const baseQueryQCardContentResponse = {};
export const QueryQCardContentResponse = {
    encode(message, writer = Writer.create()) {
        if (message.content.length !== 0) {
            writer.uint32(10).bytes(message.content);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryQCardContentResponse,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.content = reader.bytes();
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
            ...baseQueryQCardContentResponse,
        };
        if (object.content !== undefined && object.content !== null) {
            message.content = bytesFromBase64(object.content);
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.content !== undefined &&
            (obj.content = base64FromBytes(message.content !== undefined ? message.content : new Uint8Array()));
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryQCardContentResponse,
        };
        if (object.content !== undefined && object.content !== null) {
            message.content = object.content;
        }
        else {
            message.content = new Uint8Array();
        }
        return message;
    },
};
const baseQueryQUserRequest = { address: "" };
export const QueryQUserRequest = {
    encode(message, writer = Writer.create()) {
        if (message.address !== "") {
            writer.uint32(10).string(message.address);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryQUserRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.address = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryQUserRequest };
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.address !== undefined && (obj.address = message.address);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryQUserRequest };
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = "";
        }
        return message;
    },
};
const baseQueryQUserResponse = {
    alias: "",
    ownedCardSchemes: 0,
    ownedCards: 0,
};
export const QueryQUserResponse = {
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
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryQUserResponse };
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
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryQUserResponse };
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
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryQUserResponse };
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
        return message;
    },
};
export class QueryClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    Params(request) {
        const data = QueryParamsRequest.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "Params", data);
        return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
    }
    QCard(request) {
        const data = QueryQCardRequest.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QCard", data);
        return promise.then((data) => QueryQCardResponse.decode(new Reader(data)));
    }
    QCardContent(request) {
        const data = QueryQCardContentRequest.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QCardContent", data);
        return promise.then((data) => QueryQCardContentResponse.decode(new Reader(data)));
    }
    QUser(request) {
        const data = QueryQUserRequest.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QUser", data);
        return promise.then((data) => QueryQUserResponse.decode(new Reader(data)));
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
