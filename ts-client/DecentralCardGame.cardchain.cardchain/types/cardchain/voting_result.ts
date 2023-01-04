/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "DecentralCardGame.cardchain.cardchain";

export interface VotingResult {
  cardId: number;
  fairEnoughVotes: number;
  overpoweredVotes: number;
  underpoweredVotes: number;
  inappropriateVotes: number;
  result: string;
}

function createBaseVotingResult(): VotingResult {
  return {
    cardId: 0,
    fairEnoughVotes: 0,
    overpoweredVotes: 0,
    underpoweredVotes: 0,
    inappropriateVotes: 0,
    result: "",
  };
}

export const VotingResult = {
  encode(message: VotingResult, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.cardId !== 0) {
      writer.uint32(8).uint64(message.cardId);
    }
    if (message.fairEnoughVotes !== 0) {
      writer.uint32(16).uint64(message.fairEnoughVotes);
    }
    if (message.overpoweredVotes !== 0) {
      writer.uint32(24).uint64(message.overpoweredVotes);
    }
    if (message.underpoweredVotes !== 0) {
      writer.uint32(32).uint64(message.underpoweredVotes);
    }
    if (message.inappropriateVotes !== 0) {
      writer.uint32(40).uint64(message.inappropriateVotes);
    }
    if (message.result !== "") {
      writer.uint32(50).string(message.result);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): VotingResult {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseVotingResult();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.cardId = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.fairEnoughVotes = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.overpoweredVotes = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.underpoweredVotes = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.inappropriateVotes = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.result = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VotingResult {
    return {
      cardId: isSet(object.cardId) ? Number(object.cardId) : 0,
      fairEnoughVotes: isSet(object.fairEnoughVotes) ? Number(object.fairEnoughVotes) : 0,
      overpoweredVotes: isSet(object.overpoweredVotes) ? Number(object.overpoweredVotes) : 0,
      underpoweredVotes: isSet(object.underpoweredVotes) ? Number(object.underpoweredVotes) : 0,
      inappropriateVotes: isSet(object.inappropriateVotes) ? Number(object.inappropriateVotes) : 0,
      result: isSet(object.result) ? String(object.result) : "",
    };
  },

  toJSON(message: VotingResult): unknown {
    const obj: any = {};
    message.cardId !== undefined && (obj.cardId = Math.round(message.cardId));
    message.fairEnoughVotes !== undefined && (obj.fairEnoughVotes = Math.round(message.fairEnoughVotes));
    message.overpoweredVotes !== undefined && (obj.overpoweredVotes = Math.round(message.overpoweredVotes));
    message.underpoweredVotes !== undefined && (obj.underpoweredVotes = Math.round(message.underpoweredVotes));
    message.inappropriateVotes !== undefined && (obj.inappropriateVotes = Math.round(message.inappropriateVotes));
    message.result !== undefined && (obj.result = message.result);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<VotingResult>, I>>(object: I): VotingResult {
    const message = createBaseVotingResult();
    message.cardId = object.cardId ?? 0;
    message.fairEnoughVotes = object.fairEnoughVotes ?? 0;
    message.overpoweredVotes = object.overpoweredVotes ?? 0;
    message.underpoweredVotes = object.underpoweredVotes ?? 0;
    message.inappropriateVotes = object.inappropriateVotes ?? 0;
    message.result = object.result ?? "";
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
