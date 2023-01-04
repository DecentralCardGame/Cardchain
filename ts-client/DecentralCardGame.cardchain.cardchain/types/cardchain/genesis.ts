/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Coin } from "../cosmos/base/v1beta1/coin";
import { Card } from "./card";
import { Collection } from "./collection";
import { Council } from "./council";
import { Image } from "./image";
import { Match } from "./match";
import { Params } from "./params";
import { RunningAverage } from "./running_average";
import { SellOffer } from "./sell_offer";
import { Server } from "./server";
import { User } from "./user";

export const protobufPackage = "DecentralCardGame.cardchain.cardchain";

/** GenesisState defines the cardchain module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  cardRecords: Card[];
  users: User[];
  addresses: string[];
  matches: Match[];
  collections: Collection[];
  sellOffers: SellOffer[];
  pools: Coin[];
  cardAuctionPrice: string;
  councils: Council[];
  RunningAverages: RunningAverage[];
  images: Image[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  Servers: Server[];
}

function createBaseGenesisState(): GenesisState {
  return {
    params: undefined,
    cardRecords: [],
    users: [],
    addresses: [],
    matches: [],
    collections: [],
    sellOffers: [],
    pools: [],
    cardAuctionPrice: "",
    councils: [],
    RunningAverages: [],
    images: [],
    Servers: [],
  };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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
    for (const v of message.matches) {
      Match.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    for (const v of message.collections) {
      Collection.encode(v!, writer.uint32(58).fork()).ldelim();
    }
    for (const v of message.sellOffers) {
      SellOffer.encode(v!, writer.uint32(66).fork()).ldelim();
    }
    for (const v of message.pools) {
      Coin.encode(v!, writer.uint32(74).fork()).ldelim();
    }
    if (message.cardAuctionPrice !== "") {
      writer.uint32(90).string(message.cardAuctionPrice);
    }
    for (const v of message.councils) {
      Council.encode(v!, writer.uint32(98).fork()).ldelim();
    }
    for (const v of message.RunningAverages) {
      RunningAverage.encode(v!, writer.uint32(106).fork()).ldelim();
    }
    for (const v of message.images) {
      Image.encode(v!, writer.uint32(114).fork()).ldelim();
    }
    for (const v of message.Servers) {
      Server.encode(v!, writer.uint32(122).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();
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
        case 6:
          message.matches.push(Match.decode(reader, reader.uint32()));
          break;
        case 7:
          message.collections.push(Collection.decode(reader, reader.uint32()));
          break;
        case 8:
          message.sellOffers.push(SellOffer.decode(reader, reader.uint32()));
          break;
        case 9:
          message.pools.push(Coin.decode(reader, reader.uint32()));
          break;
        case 11:
          message.cardAuctionPrice = reader.string();
          break;
        case 12:
          message.councils.push(Council.decode(reader, reader.uint32()));
          break;
        case 13:
          message.RunningAverages.push(RunningAverage.decode(reader, reader.uint32()));
          break;
        case 14:
          message.images.push(Image.decode(reader, reader.uint32()));
          break;
        case 15:
          message.Servers.push(Server.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    return {
      params: isSet(object.params) ? Params.fromJSON(object.params) : undefined,
      cardRecords: Array.isArray(object?.cardRecords) ? object.cardRecords.map((e: any) => Card.fromJSON(e)) : [],
      users: Array.isArray(object?.users) ? object.users.map((e: any) => User.fromJSON(e)) : [],
      addresses: Array.isArray(object?.addresses) ? object.addresses.map((e: any) => String(e)) : [],
      matches: Array.isArray(object?.matches) ? object.matches.map((e: any) => Match.fromJSON(e)) : [],
      collections: Array.isArray(object?.collections) ? object.collections.map((e: any) => Collection.fromJSON(e)) : [],
      sellOffers: Array.isArray(object?.sellOffers) ? object.sellOffers.map((e: any) => SellOffer.fromJSON(e)) : [],
      pools: Array.isArray(object?.pools) ? object.pools.map((e: any) => Coin.fromJSON(e)) : [],
      cardAuctionPrice: isSet(object.cardAuctionPrice) ? String(object.cardAuctionPrice) : "",
      councils: Array.isArray(object?.councils) ? object.councils.map((e: any) => Council.fromJSON(e)) : [],
      RunningAverages: Array.isArray(object?.RunningAverages)
        ? object.RunningAverages.map((e: any) => RunningAverage.fromJSON(e))
        : [],
      images: Array.isArray(object?.images) ? object.images.map((e: any) => Image.fromJSON(e)) : [],
      Servers: Array.isArray(object?.Servers) ? object.Servers.map((e: any) => Server.fromJSON(e)) : [],
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.cardRecords) {
      obj.cardRecords = message.cardRecords.map((e) => e ? Card.toJSON(e) : undefined);
    } else {
      obj.cardRecords = [];
    }
    if (message.users) {
      obj.users = message.users.map((e) => e ? User.toJSON(e) : undefined);
    } else {
      obj.users = [];
    }
    if (message.addresses) {
      obj.addresses = message.addresses.map((e) => e);
    } else {
      obj.addresses = [];
    }
    if (message.matches) {
      obj.matches = message.matches.map((e) => e ? Match.toJSON(e) : undefined);
    } else {
      obj.matches = [];
    }
    if (message.collections) {
      obj.collections = message.collections.map((e) => e ? Collection.toJSON(e) : undefined);
    } else {
      obj.collections = [];
    }
    if (message.sellOffers) {
      obj.sellOffers = message.sellOffers.map((e) => e ? SellOffer.toJSON(e) : undefined);
    } else {
      obj.sellOffers = [];
    }
    if (message.pools) {
      obj.pools = message.pools.map((e) => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.pools = [];
    }
    message.cardAuctionPrice !== undefined && (obj.cardAuctionPrice = message.cardAuctionPrice);
    if (message.councils) {
      obj.councils = message.councils.map((e) => e ? Council.toJSON(e) : undefined);
    } else {
      obj.councils = [];
    }
    if (message.RunningAverages) {
      obj.RunningAverages = message.RunningAverages.map((e) => e ? RunningAverage.toJSON(e) : undefined);
    } else {
      obj.RunningAverages = [];
    }
    if (message.images) {
      obj.images = message.images.map((e) => e ? Image.toJSON(e) : undefined);
    } else {
      obj.images = [];
    }
    if (message.Servers) {
      obj.Servers = message.Servers.map((e) => e ? Server.toJSON(e) : undefined);
    } else {
      obj.Servers = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    message.cardRecords = object.cardRecords?.map((e) => Card.fromPartial(e)) || [];
    message.users = object.users?.map((e) => User.fromPartial(e)) || [];
    message.addresses = object.addresses?.map((e) => e) || [];
    message.matches = object.matches?.map((e) => Match.fromPartial(e)) || [];
    message.collections = object.collections?.map((e) => Collection.fromPartial(e)) || [];
    message.sellOffers = object.sellOffers?.map((e) => SellOffer.fromPartial(e)) || [];
    message.pools = object.pools?.map((e) => Coin.fromPartial(e)) || [];
    message.cardAuctionPrice = object.cardAuctionPrice ?? "";
    message.councils = object.councils?.map((e) => Council.fromPartial(e)) || [];
    message.RunningAverages = object.RunningAverages?.map((e) => RunningAverage.fromPartial(e)) || [];
    message.images = object.images?.map((e) => Image.fromPartial(e)) || [];
    message.Servers = object.Servers?.map((e) => Server.fromPartial(e)) || [];
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
