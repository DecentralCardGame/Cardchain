/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "DecentralCardGame.cardchain.cardchain";
const baseCardNoB64 = {
    owner: "",
    content: "",
    image: "",
    notes: "",
    fullArt: false,
    status: "",
    votePool: "",
    fairEnoughVotes: 0,
    overpoweredVotes: 0,
    underpoweredVotes: 0,
    inappropriateVotes: 0,
    nerflevel: 0,
};
export const CardNoB64 = {
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
        if (message.notes !== "") {
            writer.uint32(34).string(message.notes);
        }
        if (message.fullArt === true) {
            writer.uint32(40).bool(message.fullArt);
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
        const message = { ...baseCardNoB64 };
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
                    message.notes = reader.string();
                    break;
                case 5:
                    message.fullArt = reader.bool();
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
        const message = { ...baseCardNoB64 };
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
        if (object.notes !== undefined && object.notes !== null) {
            message.notes = String(object.notes);
        }
        else {
            message.notes = "";
        }
        if (object.fullArt !== undefined && object.fullArt !== null) {
            message.fullArt = Boolean(object.fullArt);
        }
        else {
            message.fullArt = false;
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
        message.notes !== undefined && (obj.notes = message.notes);
        message.fullArt !== undefined && (obj.fullArt = message.fullArt);
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
        const message = { ...baseCardNoB64 };
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
        if (object.notes !== undefined && object.notes !== null) {
            message.notes = object.notes;
        }
        else {
            message.notes = "";
        }
        if (object.fullArt !== undefined && object.fullArt !== null) {
            message.fullArt = object.fullArt;
        }
        else {
            message.fullArt = false;
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
