/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Outcome, outcomeFromJSON, outcomeToJSON } from "./tx";

export const protobufPackage = "DecentralCardGame.cardchain.cardchain";

export interface Match {
  timestamp: number;
  reporter: string;
  playerA: MatchPlayer | undefined;
  playerB: MatchPlayer | undefined;
  outcome: Outcome;
  coinsDistributed: boolean;
}

export interface MatchPlayer {
  addr: string;
  playedCards: number[];
  confirmed: boolean;
  outcome: Outcome;
}

function createBaseMatch(): Match {
  return { timestamp: 0, reporter: "", playerA: undefined, playerB: undefined, outcome: 0, coinsDistributed: false };
}

export const Match = {
  encode(message: Match, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.timestamp !== 0) {
      writer.uint32(8).uint64(message.timestamp);
    }
    if (message.reporter !== "") {
      writer.uint32(18).string(message.reporter);
    }
    if (message.playerA !== undefined) {
      MatchPlayer.encode(message.playerA, writer.uint32(26).fork()).ldelim();
    }
    if (message.playerB !== undefined) {
      MatchPlayer.encode(message.playerB, writer.uint32(34).fork()).ldelim();
    }
    if (message.outcome !== 0) {
      writer.uint32(56).int32(message.outcome);
    }
    if (message.coinsDistributed === true) {
      writer.uint32(80).bool(message.coinsDistributed);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Match {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMatch();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.timestamp = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.reporter = reader.string();
          break;
        case 3:
          message.playerA = MatchPlayer.decode(reader, reader.uint32());
          break;
        case 4:
          message.playerB = MatchPlayer.decode(reader, reader.uint32());
          break;
        case 7:
          message.outcome = reader.int32() as any;
          break;
        case 10:
          message.coinsDistributed = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Match {
    return {
      timestamp: isSet(object.timestamp) ? Number(object.timestamp) : 0,
      reporter: isSet(object.reporter) ? String(object.reporter) : "",
      playerA: isSet(object.playerA) ? MatchPlayer.fromJSON(object.playerA) : undefined,
      playerB: isSet(object.playerB) ? MatchPlayer.fromJSON(object.playerB) : undefined,
      outcome: isSet(object.outcome) ? outcomeFromJSON(object.outcome) : 0,
      coinsDistributed: isSet(object.coinsDistributed) ? Boolean(object.coinsDistributed) : false,
    };
  },

  toJSON(message: Match): unknown {
    const obj: any = {};
    message.timestamp !== undefined && (obj.timestamp = Math.round(message.timestamp));
    message.reporter !== undefined && (obj.reporter = message.reporter);
    message.playerA !== undefined && (obj.playerA = message.playerA ? MatchPlayer.toJSON(message.playerA) : undefined);
    message.playerB !== undefined && (obj.playerB = message.playerB ? MatchPlayer.toJSON(message.playerB) : undefined);
    message.outcome !== undefined && (obj.outcome = outcomeToJSON(message.outcome));
    message.coinsDistributed !== undefined && (obj.coinsDistributed = message.coinsDistributed);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Match>, I>>(object: I): Match {
    const message = createBaseMatch();
    message.timestamp = object.timestamp ?? 0;
    message.reporter = object.reporter ?? "";
    message.playerA = (object.playerA !== undefined && object.playerA !== null)
      ? MatchPlayer.fromPartial(object.playerA)
      : undefined;
    message.playerB = (object.playerB !== undefined && object.playerB !== null)
      ? MatchPlayer.fromPartial(object.playerB)
      : undefined;
    message.outcome = object.outcome ?? 0;
    message.coinsDistributed = object.coinsDistributed ?? false;
    return message;
  },
};

function createBaseMatchPlayer(): MatchPlayer {
  return { addr: "", playedCards: [], confirmed: false, outcome: 0 };
}

export const MatchPlayer = {
  encode(message: MatchPlayer, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.addr !== "") {
      writer.uint32(10).string(message.addr);
    }
    writer.uint32(18).fork();
    for (const v of message.playedCards) {
      writer.uint64(v);
    }
    writer.ldelim();
    if (message.confirmed === true) {
      writer.uint32(24).bool(message.confirmed);
    }
    if (message.outcome !== 0) {
      writer.uint32(32).int32(message.outcome);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MatchPlayer {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMatchPlayer();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.addr = reader.string();
          break;
        case 2:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.playedCards.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.playedCards.push(longToNumber(reader.uint64() as Long));
          }
          break;
        case 3:
          message.confirmed = reader.bool();
          break;
        case 4:
          message.outcome = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MatchPlayer {
    return {
      addr: isSet(object.addr) ? String(object.addr) : "",
      playedCards: Array.isArray(object?.playedCards) ? object.playedCards.map((e: any) => Number(e)) : [],
      confirmed: isSet(object.confirmed) ? Boolean(object.confirmed) : false,
      outcome: isSet(object.outcome) ? outcomeFromJSON(object.outcome) : 0,
    };
  },

  toJSON(message: MatchPlayer): unknown {
    const obj: any = {};
    message.addr !== undefined && (obj.addr = message.addr);
    if (message.playedCards) {
      obj.playedCards = message.playedCards.map((e) => Math.round(e));
    } else {
      obj.playedCards = [];
    }
    message.confirmed !== undefined && (obj.confirmed = message.confirmed);
    message.outcome !== undefined && (obj.outcome = outcomeToJSON(message.outcome));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MatchPlayer>, I>>(object: I): MatchPlayer {
    const message = createBaseMatchPlayer();
    message.addr = object.addr ?? "";
    message.playedCards = object.playedCards?.map((e) => e) || [];
    message.confirmed = object.confirmed ?? false;
    message.outcome = object.outcome ?? 0;
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
