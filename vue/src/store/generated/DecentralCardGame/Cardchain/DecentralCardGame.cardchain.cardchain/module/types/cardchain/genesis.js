/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Params } from "../cardchain/params";
import { Card } from "../cardchain/card";
import { User } from "../cardchain/user";
import { Match } from "../cardchain/match";
export const protobufPackage = "DecentralCardGame.cardchain.cardchain";
const baseGenesisState = { addresses: "", lastCardSchemeId: 0 };
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        if (message.params !== undefined) {
            Params.encode(message.params, writer.uint32(10).fork()).ldelim();
        }
        for (const v of message.cardRecords) {
            Card.encode(v, writer.uint32(18).fork()).ldelim();
        }
        for (const v of message.users) {
            User.encode(v, writer.uint32(26).fork()).ldelim();
        }
        for (const v of message.addresses) {
            writer.uint32(34).string(v);
        }
        if (message.lastCardSchemeId !== 0) {
            writer.uint32(40).uint64(message.lastCardSchemeId);
        }
        for (const v of message.matches) {
            Match.encode(v, writer.uint32(50).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.cardRecords = [];
        message.users = [];
        message.addresses = [];
        message.matches = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.params = Params.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.cardRecords.push(Card.decode(reader, reader.uint32()));
                    break;
                case 3:
                    message.users.push(User.decode(reader, reader.uint32()));
                    break;
                case 4:
                    message.addresses.push(reader.string());
                    break;
                case 5:
                    message.lastCardSchemeId = longToNumber(reader.uint64());
                    break;
                case 6:
                    message.matches.push(Match.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseGenesisState };
        message.cardRecords = [];
        message.users = [];
        message.addresses = [];
        message.matches = [];
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromJSON(object.params);
        }
        else {
            message.params = undefined;
        }
        if (object.cardRecords !== undefined && object.cardRecords !== null) {
            for (const e of object.cardRecords) {
                message.cardRecords.push(Card.fromJSON(e));
            }
        }
        if (object.users !== undefined && object.users !== null) {
            for (const e of object.users) {
                message.users.push(User.fromJSON(e));
            }
        }
        if (object.addresses !== undefined && object.addresses !== null) {
            for (const e of object.addresses) {
                message.addresses.push(String(e));
            }
        }
        if (object.lastCardSchemeId !== undefined &&
            object.lastCardSchemeId !== null) {
            message.lastCardSchemeId = Number(object.lastCardSchemeId);
        }
        else {
            message.lastCardSchemeId = 0;
        }
        if (object.matches !== undefined && object.matches !== null) {
            for (const e of object.matches) {
                message.matches.push(Match.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.params !== undefined &&
            (obj.params = message.params ? Params.toJSON(message.params) : undefined);
        if (message.cardRecords) {
            obj.cardRecords = message.cardRecords.map((e) => e ? Card.toJSON(e) : undefined);
        }
        else {
            obj.cardRecords = [];
        }
        if (message.users) {
            obj.users = message.users.map((e) => (e ? User.toJSON(e) : undefined));
        }
        else {
            obj.users = [];
        }
        if (message.addresses) {
            obj.addresses = message.addresses.map((e) => e);
        }
        else {
            obj.addresses = [];
        }
        message.lastCardSchemeId !== undefined &&
            (obj.lastCardSchemeId = message.lastCardSchemeId);
        if (message.matches) {
            obj.matches = message.matches.map((e) => e ? Match.toJSON(e) : undefined);
        }
        else {
            obj.matches = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.cardRecords = [];
        message.users = [];
        message.addresses = [];
        message.matches = [];
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromPartial(object.params);
        }
        else {
            message.params = undefined;
        }
        if (object.cardRecords !== undefined && object.cardRecords !== null) {
            for (const e of object.cardRecords) {
                message.cardRecords.push(Card.fromPartial(e));
            }
        }
        if (object.users !== undefined && object.users !== null) {
            for (const e of object.users) {
                message.users.push(User.fromPartial(e));
            }
        }
        if (object.addresses !== undefined && object.addresses !== null) {
            for (const e of object.addresses) {
                message.addresses.push(e);
            }
        }
        if (object.lastCardSchemeId !== undefined &&
            object.lastCardSchemeId !== null) {
            message.lastCardSchemeId = object.lastCardSchemeId;
        }
        else {
            message.lastCardSchemeId = 0;
        }
        if (object.matches !== undefined && object.matches !== null) {
            for (const e of object.matches) {
                message.matches.push(Match.fromPartial(e));
            }
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
