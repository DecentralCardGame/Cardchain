/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { VoteRight } from "../cardchain/vote_right";

export const protobufPackage = "DecentralCardGame.cardchain.cardchain";

export interface User {
  alias: string;
  ownedCardSchemes: number[];
  ownedCards: number[];
  voteRights: VoteRight[];
}

const baseUser: object = { alias: "", ownedCardSchemes: 0, ownedCards: 0 };

export const User = {
  encode(message: User, writer: Writer = Writer.create()): Writer {
    if (message.alias !== "") {
      writer.uint32(10).string(message.alias);
    }
    writer.uint32(18).fork();
    for (const v of message.ownedCardSchemes) {
      writer.uint64(v);
    }
    writer.ldelim();
    writer.uint32(26).fork();
    for (const v of message.ownedCards) {
      writer.uint64(v);
    }
    writer.ldelim();
    for (const v of message.voteRights) {
      VoteRight.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): User {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseUser } as User;
    message.ownedCardSchemes = [];
    message.ownedCards = [];
    message.voteRights = [];
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
              message.ownedCardSchemes.push(
                longToNumber(reader.uint64() as Long)
              );
            }
          } else {
            message.ownedCardSchemes.push(
              longToNumber(reader.uint64() as Long)
            );
          }
          break;
        case 3:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.ownedCards.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.ownedCards.push(longToNumber(reader.uint64() as Long));
          }
          break;
        case 4:
          message.voteRights.push(VoteRight.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): User {
    const message = { ...baseUser } as User;
    message.ownedCardSchemes = [];
    message.ownedCards = [];
    message.voteRights = [];
    if (object.alias !== undefined && object.alias !== null) {
      message.alias = String(object.alias);
    } else {
      message.alias = "";
    }
    if (
      object.ownedCardSchemes !== undefined &&
      object.ownedCardSchemes !== null
    ) {
      for (const e of object.ownedCardSchemes) {
        message.ownedCardSchemes.push(Number(e));
      }
    }
    if (object.ownedCards !== undefined && object.ownedCards !== null) {
      for (const e of object.ownedCards) {
        message.ownedCards.push(Number(e));
      }
    }
    if (object.voteRights !== undefined && object.voteRights !== null) {
      for (const e of object.voteRights) {
        message.voteRights.push(VoteRight.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: User): unknown {
    const obj: any = {};
    message.alias !== undefined && (obj.alias = message.alias);
    if (message.ownedCardSchemes) {
      obj.ownedCardSchemes = message.ownedCardSchemes.map((e) => e);
    } else {
      obj.ownedCardSchemes = [];
    }
    if (message.ownedCards) {
      obj.ownedCards = message.ownedCards.map((e) => e);
    } else {
      obj.ownedCards = [];
    }
    if (message.voteRights) {
      obj.voteRights = message.voteRights.map((e) =>
        e ? VoteRight.toJSON(e) : undefined
      );
    } else {
      obj.voteRights = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<User>): User {
    const message = { ...baseUser } as User;
    message.ownedCardSchemes = [];
    message.ownedCards = [];
    message.voteRights = [];
    if (object.alias !== undefined && object.alias !== null) {
      message.alias = object.alias;
    } else {
      message.alias = "";
    }
    if (
      object.ownedCardSchemes !== undefined &&
      object.ownedCardSchemes !== null
    ) {
      for (const e of object.ownedCardSchemes) {
        message.ownedCardSchemes.push(e);
      }
    }
    if (object.ownedCards !== undefined && object.ownedCards !== null) {
      for (const e of object.ownedCards) {
        message.ownedCards.push(e);
      }
    }
    if (object.voteRights !== undefined && object.voteRights !== null) {
      for (const e of object.voteRights) {
        message.voteRights.push(VoteRight.fromPartial(e));
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
