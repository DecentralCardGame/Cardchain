/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { VoteRight } from "./vote_right";

export const protobufPackage = "DecentralCardGame.cardchain.cardchain";

export enum CouncilStatus {
  available = 0,
  unavailable = 1,
  openCouncil = 2,
  startedCouncil = 3,
  UNRECOGNIZED = -1,
}

export function councilStatusFromJSON(object: any): CouncilStatus {
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

export function councilStatusToJSON(object: CouncilStatus): string {
  switch (object) {
    case CouncilStatus.available:
      return "available";
    case CouncilStatus.unavailable:
      return "unavailable";
    case CouncilStatus.openCouncil:
      return "openCouncil";
    case CouncilStatus.startedCouncil:
      return "startedCouncil";
    case CouncilStatus.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export enum AirDrop {
  play = 0,
  vote = 1,
  create = 2,
  buy = 3,
  user = 4,
  UNRECOGNIZED = -1,
}

export function airDropFromJSON(object: any): AirDrop {
  switch (object) {
    case 0:
    case "play":
      return AirDrop.play;
    case 1:
    case "vote":
      return AirDrop.vote;
    case 2:
    case "create":
      return AirDrop.create;
    case 3:
    case "buy":
      return AirDrop.buy;
    case 4:
    case "user":
      return AirDrop.user;
    case -1:
    case "UNRECOGNIZED":
    default:
      return AirDrop.UNRECOGNIZED;
  }
}

export function airDropToJSON(object: AirDrop): string {
  switch (object) {
    case AirDrop.play:
      return "play";
    case AirDrop.vote:
      return "vote";
    case AirDrop.create:
      return "create";
    case AirDrop.buy:
      return "buy";
    case AirDrop.user:
      return "user";
    case AirDrop.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface User {
  alias: string;
  ownedCardSchemes: number[];
  ownedPrototypes: number[];
  cards: number[];
  voteRights: VoteRight[];
  CouncilStatus: CouncilStatus;
  ReportMatches: boolean;
  profileCard: number;
  airDrops: AirDrops | undefined;
  boosterPacks: BoosterPack[];
  website: string;
  biography: string;
}

export interface BoosterPack {
  collectionId: number;
  timeStamp: number;
  raritiesPerPack: number[];
}

export interface AirDrops {
  vote: boolean;
  create: boolean;
  buy: boolean;
  play: boolean;
  user: boolean;
}

function createBaseUser(): User {
  return {
    alias: "",
    ownedCardSchemes: [],
    ownedPrototypes: [],
    cards: [],
    voteRights: [],
    CouncilStatus: 0,
    ReportMatches: false,
    profileCard: 0,
    airDrops: undefined,
    boosterPacks: [],
    website: "",
    biography: "",
  };
}

export const User = {
  encode(message: User, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.alias !== "") {
      writer.uint32(10).string(message.alias);
    }
    writer.uint32(18).fork();
    for (const v of message.ownedCardSchemes) {
      writer.uint64(v);
    }
    writer.ldelim();
    writer.uint32(26).fork();
    for (const v of message.ownedPrototypes) {
      writer.uint64(v);
    }
    writer.ldelim();
    writer.uint32(34).fork();
    for (const v of message.cards) {
      writer.uint64(v);
    }
    writer.ldelim();
    for (const v of message.voteRights) {
      VoteRight.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    if (message.CouncilStatus !== 0) {
      writer.uint32(48).int32(message.CouncilStatus);
    }
    if (message.ReportMatches === true) {
      writer.uint32(56).bool(message.ReportMatches);
    }
    if (message.profileCard !== 0) {
      writer.uint32(64).uint64(message.profileCard);
    }
    if (message.airDrops !== undefined) {
      AirDrops.encode(message.airDrops, writer.uint32(74).fork()).ldelim();
    }
    for (const v of message.boosterPacks) {
      BoosterPack.encode(v!, writer.uint32(82).fork()).ldelim();
    }
    if (message.website !== "") {
      writer.uint32(90).string(message.website);
    }
    if (message.biography !== "") {
      writer.uint32(98).string(message.biography);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): User {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUser();
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
              message.ownedCardSchemes.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.ownedCardSchemes.push(longToNumber(reader.uint64() as Long));
          }
          break;
        case 3:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.ownedPrototypes.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.ownedPrototypes.push(longToNumber(reader.uint64() as Long));
          }
          break;
        case 4:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.cards.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.cards.push(longToNumber(reader.uint64() as Long));
          }
          break;
        case 5:
          message.voteRights.push(VoteRight.decode(reader, reader.uint32()));
          break;
        case 6:
          message.CouncilStatus = reader.int32() as any;
          break;
        case 7:
          message.ReportMatches = reader.bool();
          break;
        case 8:
          message.profileCard = longToNumber(reader.uint64() as Long);
          break;
        case 9:
          message.airDrops = AirDrops.decode(reader, reader.uint32());
          break;
        case 10:
          message.boosterPacks.push(BoosterPack.decode(reader, reader.uint32()));
          break;
        case 11:
          message.website = reader.string();
          break;
        case 12:
          message.biography = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): User {
    return {
      alias: isSet(object.alias) ? String(object.alias) : "",
      ownedCardSchemes: Array.isArray(object?.ownedCardSchemes)
        ? object.ownedCardSchemes.map((e: any) => Number(e))
        : [],
      ownedPrototypes: Array.isArray(object?.ownedPrototypes) ? object.ownedPrototypes.map((e: any) => Number(e)) : [],
      cards: Array.isArray(object?.cards) ? object.cards.map((e: any) => Number(e)) : [],
      voteRights: Array.isArray(object?.voteRights) ? object.voteRights.map((e: any) => VoteRight.fromJSON(e)) : [],
      CouncilStatus: isSet(object.CouncilStatus) ? councilStatusFromJSON(object.CouncilStatus) : 0,
      ReportMatches: isSet(object.ReportMatches) ? Boolean(object.ReportMatches) : false,
      profileCard: isSet(object.profileCard) ? Number(object.profileCard) : 0,
      airDrops: isSet(object.airDrops) ? AirDrops.fromJSON(object.airDrops) : undefined,
      boosterPacks: Array.isArray(object?.boosterPacks)
        ? object.boosterPacks.map((e: any) => BoosterPack.fromJSON(e))
        : [],
      website: isSet(object.website) ? String(object.website) : "",
      biography: isSet(object.biography) ? String(object.biography) : "",
    };
  },

  toJSON(message: User): unknown {
    const obj: any = {};
    message.alias !== undefined && (obj.alias = message.alias);
    if (message.ownedCardSchemes) {
      obj.ownedCardSchemes = message.ownedCardSchemes.map((e) => Math.round(e));
    } else {
      obj.ownedCardSchemes = [];
    }
    if (message.ownedPrototypes) {
      obj.ownedPrototypes = message.ownedPrototypes.map((e) => Math.round(e));
    } else {
      obj.ownedPrototypes = [];
    }
    if (message.cards) {
      obj.cards = message.cards.map((e) => Math.round(e));
    } else {
      obj.cards = [];
    }
    if (message.voteRights) {
      obj.voteRights = message.voteRights.map((e) => e ? VoteRight.toJSON(e) : undefined);
    } else {
      obj.voteRights = [];
    }
    message.CouncilStatus !== undefined && (obj.CouncilStatus = councilStatusToJSON(message.CouncilStatus));
    message.ReportMatches !== undefined && (obj.ReportMatches = message.ReportMatches);
    message.profileCard !== undefined && (obj.profileCard = Math.round(message.profileCard));
    message.airDrops !== undefined && (obj.airDrops = message.airDrops ? AirDrops.toJSON(message.airDrops) : undefined);
    if (message.boosterPacks) {
      obj.boosterPacks = message.boosterPacks.map((e) => e ? BoosterPack.toJSON(e) : undefined);
    } else {
      obj.boosterPacks = [];
    }
    message.website !== undefined && (obj.website = message.website);
    message.biography !== undefined && (obj.biography = message.biography);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<User>, I>>(object: I): User {
    const message = createBaseUser();
    message.alias = object.alias ?? "";
    message.ownedCardSchemes = object.ownedCardSchemes?.map((e) => e) || [];
    message.ownedPrototypes = object.ownedPrototypes?.map((e) => e) || [];
    message.cards = object.cards?.map((e) => e) || [];
    message.voteRights = object.voteRights?.map((e) => VoteRight.fromPartial(e)) || [];
    message.CouncilStatus = object.CouncilStatus ?? 0;
    message.ReportMatches = object.ReportMatches ?? false;
    message.profileCard = object.profileCard ?? 0;
    message.airDrops = (object.airDrops !== undefined && object.airDrops !== null)
      ? AirDrops.fromPartial(object.airDrops)
      : undefined;
    message.boosterPacks = object.boosterPacks?.map((e) => BoosterPack.fromPartial(e)) || [];
    message.website = object.website ?? "";
    message.biography = object.biography ?? "";
    return message;
  },
};

function createBaseBoosterPack(): BoosterPack {
  return { collectionId: 0, timeStamp: 0, raritiesPerPack: [] };
}

export const BoosterPack = {
  encode(message: BoosterPack, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.collectionId !== 0) {
      writer.uint32(8).uint64(message.collectionId);
    }
    if (message.timeStamp !== 0) {
      writer.uint32(16).int64(message.timeStamp);
    }
    writer.uint32(26).fork();
    for (const v of message.raritiesPerPack) {
      writer.uint64(v);
    }
    writer.ldelim();
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): BoosterPack {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBoosterPack();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.collectionId = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.timeStamp = longToNumber(reader.int64() as Long);
          break;
        case 3:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.raritiesPerPack.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.raritiesPerPack.push(longToNumber(reader.uint64() as Long));
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): BoosterPack {
    return {
      collectionId: isSet(object.collectionId) ? Number(object.collectionId) : 0,
      timeStamp: isSet(object.timeStamp) ? Number(object.timeStamp) : 0,
      raritiesPerPack: Array.isArray(object?.raritiesPerPack) ? object.raritiesPerPack.map((e: any) => Number(e)) : [],
    };
  },

  toJSON(message: BoosterPack): unknown {
    const obj: any = {};
    message.collectionId !== undefined && (obj.collectionId = Math.round(message.collectionId));
    message.timeStamp !== undefined && (obj.timeStamp = Math.round(message.timeStamp));
    if (message.raritiesPerPack) {
      obj.raritiesPerPack = message.raritiesPerPack.map((e) => Math.round(e));
    } else {
      obj.raritiesPerPack = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<BoosterPack>, I>>(object: I): BoosterPack {
    const message = createBaseBoosterPack();
    message.collectionId = object.collectionId ?? 0;
    message.timeStamp = object.timeStamp ?? 0;
    message.raritiesPerPack = object.raritiesPerPack?.map((e) => e) || [];
    return message;
  },
};

function createBaseAirDrops(): AirDrops {
  return { vote: false, create: false, buy: false, play: false, user: false };
}

export const AirDrops = {
  encode(message: AirDrops, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.vote === true) {
      writer.uint32(8).bool(message.vote);
    }
    if (message.create === true) {
      writer.uint32(16).bool(message.create);
    }
    if (message.buy === true) {
      writer.uint32(24).bool(message.buy);
    }
    if (message.play === true) {
      writer.uint32(32).bool(message.play);
    }
    if (message.user === true) {
      writer.uint32(40).bool(message.user);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AirDrops {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAirDrops();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.vote = reader.bool();
          break;
        case 2:
          message.create = reader.bool();
          break;
        case 3:
          message.buy = reader.bool();
          break;
        case 4:
          message.play = reader.bool();
          break;
        case 5:
          message.user = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AirDrops {
    return {
      vote: isSet(object.vote) ? Boolean(object.vote) : false,
      create: isSet(object.create) ? Boolean(object.create) : false,
      buy: isSet(object.buy) ? Boolean(object.buy) : false,
      play: isSet(object.play) ? Boolean(object.play) : false,
      user: isSet(object.user) ? Boolean(object.user) : false,
    };
  },

  toJSON(message: AirDrops): unknown {
    const obj: any = {};
    message.vote !== undefined && (obj.vote = message.vote);
    message.create !== undefined && (obj.create = message.create);
    message.buy !== undefined && (obj.buy = message.buy);
    message.play !== undefined && (obj.play = message.play);
    message.user !== undefined && (obj.user = message.user);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<AirDrops>, I>>(object: I): AirDrops {
    const message = createBaseAirDrops();
    message.vote = object.vote ?? false;
    message.create = object.create ?? false;
    message.buy = object.buy ?? false;
    message.play = object.play ?? false;
    message.user = object.user ?? false;
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
