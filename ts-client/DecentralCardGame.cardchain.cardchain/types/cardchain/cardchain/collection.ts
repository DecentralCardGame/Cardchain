/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "DecentralCardGame.cardchain.cardchain";

export enum CStatus {
  design = 0,
  finalized = 1,
  active = 2,
  archived = 3,
  UNRECOGNIZED = -1,
}

export function cStatusFromJSON(object: any): CStatus {
  switch (object) {
    case 0:
    case "design":
      return CStatus.design;
    case 1:
    case "finalized":
      return CStatus.finalized;
    case 2:
    case "active":
      return CStatus.active;
    case 3:
    case "archived":
      return CStatus.archived;
    case -1:
    case "UNRECOGNIZED":
    default:
      return CStatus.UNRECOGNIZED;
  }
}

export function cStatusToJSON(object: CStatus): string {
  switch (object) {
    case CStatus.design:
      return "design";
    case CStatus.finalized:
      return "finalized";
    case CStatus.active:
      return "active";
    case CStatus.archived:
      return "archived";
    case CStatus.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface Collection {
  name: string;
  cards: number[];
  artist: string;
  storyWriter: string;
  contributors: string[];
  story: string;
  artworkId: number;
  status: CStatus;
  timeStamp: number;
}

export interface OutpCollection {
  name: string;
  cards: number[];
  artist: string;
  storyWriter: string;
  contributors: string[];
  story: string;
  artwork: string;
  status: CStatus;
  timeStamp: number;
}

function createBaseCollection(): Collection {
  return {
    name: "",
    cards: [],
    artist: "",
    storyWriter: "",
    contributors: [],
    story: "",
    artworkId: 0,
    status: 0,
    timeStamp: 0,
  };
}

export const Collection = {
  encode(message: Collection, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    writer.uint32(18).fork();
    for (const v of message.cards) {
      writer.uint64(v);
    }
    writer.ldelim();
    if (message.artist !== "") {
      writer.uint32(26).string(message.artist);
    }
    if (message.storyWriter !== "") {
      writer.uint32(34).string(message.storyWriter);
    }
    for (const v of message.contributors) {
      writer.uint32(42).string(v!);
    }
    if (message.story !== "") {
      writer.uint32(50).string(message.story);
    }
    if (message.artworkId !== 0) {
      writer.uint32(56).uint64(message.artworkId);
    }
    if (message.status !== 0) {
      writer.uint32(64).int32(message.status);
    }
    if (message.timeStamp !== 0) {
      writer.uint32(72).int64(message.timeStamp);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Collection {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCollection();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.cards.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.cards.push(longToNumber(reader.uint64() as Long));
          }
          break;
        case 3:
          message.artist = reader.string();
          break;
        case 4:
          message.storyWriter = reader.string();
          break;
        case 5:
          message.contributors.push(reader.string());
          break;
        case 6:
          message.story = reader.string();
          break;
        case 7:
          message.artworkId = longToNumber(reader.uint64() as Long);
          break;
        case 8:
          message.status = reader.int32() as any;
          break;
        case 9:
          message.timeStamp = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Collection {
    return {
      name: isSet(object.name) ? String(object.name) : "",
      cards: Array.isArray(object?.cards) ? object.cards.map((e: any) => Number(e)) : [],
      artist: isSet(object.artist) ? String(object.artist) : "",
      storyWriter: isSet(object.storyWriter) ? String(object.storyWriter) : "",
      contributors: Array.isArray(object?.contributors) ? object.contributors.map((e: any) => String(e)) : [],
      story: isSet(object.story) ? String(object.story) : "",
      artworkId: isSet(object.artworkId) ? Number(object.artworkId) : 0,
      status: isSet(object.status) ? cStatusFromJSON(object.status) : 0,
      timeStamp: isSet(object.timeStamp) ? Number(object.timeStamp) : 0,
    };
  },

  toJSON(message: Collection): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    if (message.cards) {
      obj.cards = message.cards.map((e) => Math.round(e));
    } else {
      obj.cards = [];
    }
    message.artist !== undefined && (obj.artist = message.artist);
    message.storyWriter !== undefined && (obj.storyWriter = message.storyWriter);
    if (message.contributors) {
      obj.contributors = message.contributors.map((e) => e);
    } else {
      obj.contributors = [];
    }
    message.story !== undefined && (obj.story = message.story);
    message.artworkId !== undefined && (obj.artworkId = Math.round(message.artworkId));
    message.status !== undefined && (obj.status = cStatusToJSON(message.status));
    message.timeStamp !== undefined && (obj.timeStamp = Math.round(message.timeStamp));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Collection>, I>>(object: I): Collection {
    const message = createBaseCollection();
    message.name = object.name ?? "";
    message.cards = object.cards?.map((e) => e) || [];
    message.artist = object.artist ?? "";
    message.storyWriter = object.storyWriter ?? "";
    message.contributors = object.contributors?.map((e) => e) || [];
    message.story = object.story ?? "";
    message.artworkId = object.artworkId ?? 0;
    message.status = object.status ?? 0;
    message.timeStamp = object.timeStamp ?? 0;
    return message;
  },
};

function createBaseOutpCollection(): OutpCollection {
  return {
    name: "",
    cards: [],
    artist: "",
    storyWriter: "",
    contributors: [],
    story: "",
    artwork: "",
    status: 0,
    timeStamp: 0,
  };
}

export const OutpCollection = {
  encode(message: OutpCollection, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    writer.uint32(18).fork();
    for (const v of message.cards) {
      writer.uint64(v);
    }
    writer.ldelim();
    if (message.artist !== "") {
      writer.uint32(26).string(message.artist);
    }
    if (message.storyWriter !== "") {
      writer.uint32(34).string(message.storyWriter);
    }
    for (const v of message.contributors) {
      writer.uint32(42).string(v!);
    }
    if (message.story !== "") {
      writer.uint32(50).string(message.story);
    }
    if (message.artwork !== "") {
      writer.uint32(58).string(message.artwork);
    }
    if (message.status !== 0) {
      writer.uint32(64).int32(message.status);
    }
    if (message.timeStamp !== 0) {
      writer.uint32(72).int64(message.timeStamp);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): OutpCollection {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOutpCollection();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.cards.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.cards.push(longToNumber(reader.uint64() as Long));
          }
          break;
        case 3:
          message.artist = reader.string();
          break;
        case 4:
          message.storyWriter = reader.string();
          break;
        case 5:
          message.contributors.push(reader.string());
          break;
        case 6:
          message.story = reader.string();
          break;
        case 7:
          message.artwork = reader.string();
          break;
        case 8:
          message.status = reader.int32() as any;
          break;
        case 9:
          message.timeStamp = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OutpCollection {
    return {
      name: isSet(object.name) ? String(object.name) : "",
      cards: Array.isArray(object?.cards) ? object.cards.map((e: any) => Number(e)) : [],
      artist: isSet(object.artist) ? String(object.artist) : "",
      storyWriter: isSet(object.storyWriter) ? String(object.storyWriter) : "",
      contributors: Array.isArray(object?.contributors) ? object.contributors.map((e: any) => String(e)) : [],
      story: isSet(object.story) ? String(object.story) : "",
      artwork: isSet(object.artwork) ? String(object.artwork) : "",
      status: isSet(object.status) ? cStatusFromJSON(object.status) : 0,
      timeStamp: isSet(object.timeStamp) ? Number(object.timeStamp) : 0,
    };
  },

  toJSON(message: OutpCollection): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    if (message.cards) {
      obj.cards = message.cards.map((e) => Math.round(e));
    } else {
      obj.cards = [];
    }
    message.artist !== undefined && (obj.artist = message.artist);
    message.storyWriter !== undefined && (obj.storyWriter = message.storyWriter);
    if (message.contributors) {
      obj.contributors = message.contributors.map((e) => e);
    } else {
      obj.contributors = [];
    }
    message.story !== undefined && (obj.story = message.story);
    message.artwork !== undefined && (obj.artwork = message.artwork);
    message.status !== undefined && (obj.status = cStatusToJSON(message.status));
    message.timeStamp !== undefined && (obj.timeStamp = Math.round(message.timeStamp));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<OutpCollection>, I>>(object: I): OutpCollection {
    const message = createBaseOutpCollection();
    message.name = object.name ?? "";
    message.cards = object.cards?.map((e) => e) || [];
    message.artist = object.artist ?? "";
    message.storyWriter = object.storyWriter ?? "";
    message.contributors = object.contributors?.map((e) => e) || [];
    message.story = object.story ?? "";
    message.artwork = object.artwork ?? "";
    message.status = object.status ?? 0;
    message.timeStamp = object.timeStamp ?? 0;
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
