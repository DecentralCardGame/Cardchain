/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Params } from "../cardchain/params";
import { Card } from "../cardchain/card";
import { User } from "../cardchain/user";
import { Match } from "../cardchain/match";
import { Collection } from "../cardchain/collection";
import { SellOffer } from "../cardchain/sell_offer";

export const protobufPackage = "DecentralCardGame.cardchain.cardchain";

/** GenesisState defines the cardchain module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  cardRecords: Card[];
  users: User[];
  addresses: string[];
  lastCardSchemeId: number;
  matches: Match[];
  collections: Collection[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  sellOffers: SellOffer[];
}

const baseGenesisState: object = { addresses: "", lastCardSchemeId: 0 };

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.cardRecords) {
      Card.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.users) {
      User.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.addresses) {
      writer.uint32(34).string(v!);
    }
    if (message.lastCardSchemeId !== 0) {
      writer.uint32(40).uint64(message.lastCardSchemeId);
    }
    for (const v of message.matches) {
      Match.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    for (const v of message.collections) {
      Collection.encode(v!, writer.uint32(58).fork()).ldelim();
    }
    for (const v of message.sellOffers) {
      SellOffer.encode(v!, writer.uint32(66).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.cardRecords = [];
    message.users = [];
    message.addresses = [];
    message.matches = [];
    message.collections = [];
    message.sellOffers = [];
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
          message.lastCardSchemeId = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.matches.push(Match.decode(reader, reader.uint32()));
          break;
        case 7:
          message.collections.push(Collection.decode(reader, reader.uint32()));
          break;
        case 8:
          message.sellOffers.push(SellOffer.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.cardRecords = [];
    message.users = [];
    message.addresses = [];
    message.matches = [];
    message.collections = [];
    message.sellOffers = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
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
    if (
      object.lastCardSchemeId !== undefined &&
      object.lastCardSchemeId !== null
    ) {
      message.lastCardSchemeId = Number(object.lastCardSchemeId);
    } else {
      message.lastCardSchemeId = 0;
    }
    if (object.matches !== undefined && object.matches !== null) {
      for (const e of object.matches) {
        message.matches.push(Match.fromJSON(e));
      }
    }
    if (object.collections !== undefined && object.collections !== null) {
      for (const e of object.collections) {
        message.collections.push(Collection.fromJSON(e));
      }
    }
    if (object.sellOffers !== undefined && object.sellOffers !== null) {
      for (const e of object.sellOffers) {
        message.sellOffers.push(SellOffer.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.cardRecords) {
      obj.cardRecords = message.cardRecords.map((e) =>
        e ? Card.toJSON(e) : undefined
      );
    } else {
      obj.cardRecords = [];
    }
    if (message.users) {
      obj.users = message.users.map((e) => (e ? User.toJSON(e) : undefined));
    } else {
      obj.users = [];
    }
    if (message.addresses) {
      obj.addresses = message.addresses.map((e) => e);
    } else {
      obj.addresses = [];
    }
    message.lastCardSchemeId !== undefined &&
      (obj.lastCardSchemeId = message.lastCardSchemeId);
    if (message.matches) {
      obj.matches = message.matches.map((e) =>
        e ? Match.toJSON(e) : undefined
      );
    } else {
      obj.matches = [];
    }
    if (message.collections) {
      obj.collections = message.collections.map((e) =>
        e ? Collection.toJSON(e) : undefined
      );
    } else {
      obj.collections = [];
    }
    if (message.sellOffers) {
      obj.sellOffers = message.sellOffers.map((e) =>
        e ? SellOffer.toJSON(e) : undefined
      );
    } else {
      obj.sellOffers = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.cardRecords = [];
    message.users = [];
    message.addresses = [];
    message.matches = [];
    message.collections = [];
    message.sellOffers = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
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
    if (
      object.lastCardSchemeId !== undefined &&
      object.lastCardSchemeId !== null
    ) {
      message.lastCardSchemeId = object.lastCardSchemeId;
    } else {
      message.lastCardSchemeId = 0;
    }
    if (object.matches !== undefined && object.matches !== null) {
      for (const e of object.matches) {
        message.matches.push(Match.fromPartial(e));
      }
    }
    if (object.collections !== undefined && object.collections !== null) {
      for (const e of object.collections) {
        message.collections.push(Collection.fromPartial(e));
      }
    }
    if (object.sellOffers !== undefined && object.sellOffers !== null) {
      for (const e of object.sellOffers) {
        message.sellOffers.push(SellOffer.fromPartial(e));
      }
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
