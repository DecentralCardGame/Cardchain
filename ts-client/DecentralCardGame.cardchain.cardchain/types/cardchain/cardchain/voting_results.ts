/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { VotingResult } from "./voting_result";

export const protobufPackage = "DecentralCardGame.cardchain.cardchain";

export interface VotingResults {
  totalVotes: number;
  totalFairEnoughVotes: number;
  totalOverpoweredVotes: number;
  totalUnderpoweredVotes: number;
  totalInappropriateVotes: number;
  cardResults: VotingResult[];
  notes: string;
}

function createBaseVotingResults(): VotingResults {
  return {
    totalVotes: 0,
    totalFairEnoughVotes: 0,
    totalOverpoweredVotes: 0,
    totalUnderpoweredVotes: 0,
    totalInappropriateVotes: 0,
    cardResults: [],
    notes: "",
  };
}

export const VotingResults = {
  encode(message: VotingResults, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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
      VotingResult.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    if (message.notes !== "") {
      writer.uint32(58).string(message.notes);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): VotingResults {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseVotingResults();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.totalVotes = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.totalFairEnoughVotes = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.totalOverpoweredVotes = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.totalUnderpoweredVotes = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.totalInappropriateVotes = longToNumber(reader.uint64() as Long);
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

  fromJSON(object: any): VotingResults {
    return {
      totalVotes: isSet(object.totalVotes) ? Number(object.totalVotes) : 0,
      totalFairEnoughVotes: isSet(object.totalFairEnoughVotes) ? Number(object.totalFairEnoughVotes) : 0,
      totalOverpoweredVotes: isSet(object.totalOverpoweredVotes) ? Number(object.totalOverpoweredVotes) : 0,
      totalUnderpoweredVotes: isSet(object.totalUnderpoweredVotes) ? Number(object.totalUnderpoweredVotes) : 0,
      totalInappropriateVotes: isSet(object.totalInappropriateVotes) ? Number(object.totalInappropriateVotes) : 0,
      cardResults: Array.isArray(object?.cardResults)
        ? object.cardResults.map((e: any) => VotingResult.fromJSON(e))
        : [],
      notes: isSet(object.notes) ? String(object.notes) : "",
    };
  },

  toJSON(message: VotingResults): unknown {
    const obj: any = {};
    message.totalVotes !== undefined && (obj.totalVotes = Math.round(message.totalVotes));
    message.totalFairEnoughVotes !== undefined && (obj.totalFairEnoughVotes = Math.round(message.totalFairEnoughVotes));
    message.totalOverpoweredVotes !== undefined
      && (obj.totalOverpoweredVotes = Math.round(message.totalOverpoweredVotes));
    message.totalUnderpoweredVotes !== undefined
      && (obj.totalUnderpoweredVotes = Math.round(message.totalUnderpoweredVotes));
    message.totalInappropriateVotes !== undefined
      && (obj.totalInappropriateVotes = Math.round(message.totalInappropriateVotes));
    if (message.cardResults) {
      obj.cardResults = message.cardResults.map((e) => e ? VotingResult.toJSON(e) : undefined);
    } else {
      obj.cardResults = [];
    }
    message.notes !== undefined && (obj.notes = message.notes);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<VotingResults>, I>>(object: I): VotingResults {
    const message = createBaseVotingResults();
    message.totalVotes = object.totalVotes ?? 0;
    message.totalFairEnoughVotes = object.totalFairEnoughVotes ?? 0;
    message.totalOverpoweredVotes = object.totalOverpoweredVotes ?? 0;
    message.totalUnderpoweredVotes = object.totalUnderpoweredVotes ?? 0;
    message.totalInappropriateVotes = object.totalInappropriateVotes ?? 0;
    message.cardResults = object.cardResults?.map((e) => VotingResult.fromPartial(e)) || [];
    message.notes = object.notes ?? "";
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
