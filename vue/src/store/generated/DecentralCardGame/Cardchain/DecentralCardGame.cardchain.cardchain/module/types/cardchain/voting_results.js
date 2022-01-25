/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { VotingResult } from "../cardchain/voting_result";
export const protobufPackage = "DecentralCardGame.cardchain.cardchain";
const baseVotingResults = {
    totalVotes: 0,
    totalFairEnoughVotes: 0,
    totalOverpoweredVotes: 0,
    totalUnderpoweredVotes: 0,
    totalInappropriateVotes: 0,
    notes: "",
};
export const VotingResults = {
    encode(message, writer = Writer.create()) {
        if (message.totalVotes !== 0) {
            writer.uint32(8).uint64(message.totalVotes);
        }
        if (message.totalFairEnoughVotes !== 0) {
            writer.uint32(16).uint64(message.totalFairEnoughVotes);
        }
        if (message.totalOverpoweredVotes !== 0) {
            writer.uint32(24).uint64(message.totalOverpoweredVotes);
        }
        if (message.totalUnderpoweredVotes !== 0) {
            writer.uint32(32).uint64(message.totalUnderpoweredVotes);
        }
        if (message.totalInappropriateVotes !== 0) {
            writer.uint32(40).uint64(message.totalInappropriateVotes);
        }
        for (const v of message.cardResults) {
            VotingResult.encode(v, writer.uint32(50).fork()).ldelim();
        }
        if (message.notes !== "") {
            writer.uint32(58).string(message.notes);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseVotingResults };
        message.cardResults = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.totalVotes = longToNumber(reader.uint64());
                    break;
                case 2:
                    message.totalFairEnoughVotes = longToNumber(reader.uint64());
                    break;
                case 3:
                    message.totalOverpoweredVotes = longToNumber(reader.uint64());
                    break;
                case 4:
                    message.totalUnderpoweredVotes = longToNumber(reader.uint64());
                    break;
                case 5:
                    message.totalInappropriateVotes = longToNumber(reader.uint64());
                    break;
                case 6:
                    message.cardResults.push(VotingResult.decode(reader, reader.uint32()));
                    break;
                case 7:
                    message.notes = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseVotingResults };
        message.cardResults = [];
        if (object.totalVotes !== undefined && object.totalVotes !== null) {
            message.totalVotes = Number(object.totalVotes);
        }
        else {
            message.totalVotes = 0;
        }
        if (object.totalFairEnoughVotes !== undefined &&
            object.totalFairEnoughVotes !== null) {
            message.totalFairEnoughVotes = Number(object.totalFairEnoughVotes);
        }
        else {
            message.totalFairEnoughVotes = 0;
        }
        if (object.totalOverpoweredVotes !== undefined &&
            object.totalOverpoweredVotes !== null) {
            message.totalOverpoweredVotes = Number(object.totalOverpoweredVotes);
        }
        else {
            message.totalOverpoweredVotes = 0;
        }
        if (object.totalUnderpoweredVotes !== undefined &&
            object.totalUnderpoweredVotes !== null) {
            message.totalUnderpoweredVotes = Number(object.totalUnderpoweredVotes);
        }
        else {
            message.totalUnderpoweredVotes = 0;
        }
        if (object.totalInappropriateVotes !== undefined &&
            object.totalInappropriateVotes !== null) {
            message.totalInappropriateVotes = Number(object.totalInappropriateVotes);
        }
        else {
            message.totalInappropriateVotes = 0;
        }
        if (object.cardResults !== undefined && object.cardResults !== null) {
            for (const e of object.cardResults) {
                message.cardResults.push(VotingResult.fromJSON(e));
            }
        }
        if (object.notes !== undefined && object.notes !== null) {
            message.notes = String(object.notes);
        }
        else {
            message.notes = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.totalVotes !== undefined && (obj.totalVotes = message.totalVotes);
        message.totalFairEnoughVotes !== undefined &&
            (obj.totalFairEnoughVotes = message.totalFairEnoughVotes);
        message.totalOverpoweredVotes !== undefined &&
            (obj.totalOverpoweredVotes = message.totalOverpoweredVotes);
        message.totalUnderpoweredVotes !== undefined &&
            (obj.totalUnderpoweredVotes = message.totalUnderpoweredVotes);
        message.totalInappropriateVotes !== undefined &&
            (obj.totalInappropriateVotes = message.totalInappropriateVotes);
        if (message.cardResults) {
            obj.cardResults = message.cardResults.map((e) => e ? VotingResult.toJSON(e) : undefined);
        }
        else {
            obj.cardResults = [];
        }
        message.notes !== undefined && (obj.notes = message.notes);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseVotingResults };
        message.cardResults = [];
        if (object.totalVotes !== undefined && object.totalVotes !== null) {
            message.totalVotes = object.totalVotes;
        }
        else {
            message.totalVotes = 0;
        }
        if (object.totalFairEnoughVotes !== undefined &&
            object.totalFairEnoughVotes !== null) {
            message.totalFairEnoughVotes = object.totalFairEnoughVotes;
        }
        else {
            message.totalFairEnoughVotes = 0;
        }
        if (object.totalOverpoweredVotes !== undefined &&
            object.totalOverpoweredVotes !== null) {
            message.totalOverpoweredVotes = object.totalOverpoweredVotes;
        }
        else {
            message.totalOverpoweredVotes = 0;
        }
        if (object.totalUnderpoweredVotes !== undefined &&
            object.totalUnderpoweredVotes !== null) {
            message.totalUnderpoweredVotes = object.totalUnderpoweredVotes;
        }
        else {
            message.totalUnderpoweredVotes = 0;
        }
        if (object.totalInappropriateVotes !== undefined &&
            object.totalInappropriateVotes !== null) {
            message.totalInappropriateVotes = object.totalInappropriateVotes;
        }
        else {
            message.totalInappropriateVotes = 0;
        }
        if (object.cardResults !== undefined && object.cardResults !== null) {
            for (const e of object.cardResults) {
                message.cardResults.push(VotingResult.fromPartial(e));
            }
        }
        if (object.notes !== undefined && object.notes !== null) {
            message.notes = object.notes;
        }
        else {
            message.notes = "";
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
