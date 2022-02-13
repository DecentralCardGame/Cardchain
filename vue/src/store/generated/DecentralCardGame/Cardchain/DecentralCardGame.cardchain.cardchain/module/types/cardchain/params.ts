/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "DecentralCardGame.cardchain.cardchain";

/** Params defines the parameters for the module. */
export interface Params {
  votingRightsExpirationTime: number;
  collectionSize: number;
  collectionPrice: number;
}

const baseParams: object = {
  votingRightsExpirationTime: 0,
  collectionSize: 0,
  collectionPrice: 0,
};

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    if (message.votingRightsExpirationTime !== 0) {
      writer.uint32(8).int64(message.votingRightsExpirationTime);
    }
    if (message.collectionSize !== 0) {
      writer.uint32(16).uint64(message.collectionSize);
    }
    if (message.collectionPrice !== 0) {
      writer.uint32(24).int64(message.collectionPrice);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseParams } as Params;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.votingRightsExpirationTime = longToNumber(
            reader.int64() as Long
          );
          break;
        case 2:
          message.collectionSize = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.collectionPrice = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Params {
    const message = { ...baseParams } as Params;
    if (
      object.votingRightsExpirationTime !== undefined &&
      object.votingRightsExpirationTime !== null
    ) {
      message.votingRightsExpirationTime = Number(
        object.votingRightsExpirationTime
      );
    } else {
      message.votingRightsExpirationTime = 0;
    }
    if (object.collectionSize !== undefined && object.collectionSize !== null) {
      message.collectionSize = Number(object.collectionSize);
    } else {
      message.collectionSize = 0;
    }
    if (
      object.collectionPrice !== undefined &&
      object.collectionPrice !== null
    ) {
      message.collectionPrice = Number(object.collectionPrice);
    } else {
      message.collectionPrice = 0;
    }
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.votingRightsExpirationTime !== undefined &&
      (obj.votingRightsExpirationTime = message.votingRightsExpirationTime);
    message.collectionSize !== undefined &&
      (obj.collectionSize = message.collectionSize);
    message.collectionPrice !== undefined &&
      (obj.collectionPrice = message.collectionPrice);
    return obj;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    if (
      object.votingRightsExpirationTime !== undefined &&
      object.votingRightsExpirationTime !== null
    ) {
      message.votingRightsExpirationTime = object.votingRightsExpirationTime;
    } else {
      message.votingRightsExpirationTime = 0;
    }
    if (object.collectionSize !== undefined && object.collectionSize !== null) {
      message.collectionSize = object.collectionSize;
    } else {
      message.collectionSize = 0;
    }
    if (
      object.collectionPrice !== undefined &&
      object.collectionPrice !== null
    ) {
      message.collectionPrice = object.collectionPrice;
    } else {
      message.collectionPrice = 0;
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
