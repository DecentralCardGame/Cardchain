/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "DecentralCardGame.cardchain.cardchain";

export interface VoteRight {
  cardId: number;
  expireBlock: number;
}

const baseVoteRight: object = { cardId: 0, expireBlock: 0 };

export const VoteRight = {
  encode(message: VoteRight, writer: Writer = Writer.create()): Writer {
    if (message.cardId !== 0) {
      writer.uint32(8).uint64(message.cardId);
    }
    if (message.expireBlock !== 0) {
      writer.uint32(16).int32(message.expireBlock);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): VoteRight {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseVoteRight } as VoteRight;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.cardId = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.expireBlock = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VoteRight {
    const message = { ...baseVoteRight } as VoteRight;
    if (object.cardId !== undefined && object.cardId !== null) {
      message.cardId = Number(object.cardId);
    } else {
      message.cardId = 0;
    }
    if (object.expireBlock !== undefined && object.expireBlock !== null) {
      message.expireBlock = Number(object.expireBlock);
    } else {
      message.expireBlock = 0;
    }
    return message;
  },

  toJSON(message: VoteRight): unknown {
    const obj: any = {};
    message.cardId !== undefined && (obj.cardId = message.cardId);
    message.expireBlock !== undefined &&
      (obj.expireBlock = message.expireBlock);
    return obj;
  },

  fromPartial(object: DeepPartial<VoteRight>): VoteRight {
    const message = { ...baseVoteRight } as VoteRight;
    if (object.cardId !== undefined && object.cardId !== null) {
      message.cardId = object.cardId;
    } else {
      message.cardId = 0;
    }
    if (object.expireBlock !== undefined && object.expireBlock !== null) {
      message.expireBlock = object.expireBlock;
    } else {
      message.expireBlock = 0;
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
