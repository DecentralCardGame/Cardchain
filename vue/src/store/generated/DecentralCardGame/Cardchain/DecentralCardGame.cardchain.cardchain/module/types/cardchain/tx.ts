/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "DecentralCardGame.cardchain.cardchain";

export interface MsgCreateuser {
  creator: string;
  newUser: string;
  alias: string;
}

export interface MsgCreateuserResponse {}

export interface MsgBuyCardScheme {
  creator: string;
  bid: string;
  buyer: string;
}

export interface MsgBuyCardSchemeResponse {}

export interface MsgVoteCard {
  creator: string;
  cardId: number;
  voteType: string;
  voter: string;
}

export interface MsgVoteCardResponse {}

const baseMsgCreateuser: object = { creator: "", newUser: "", alias: "" };

export const MsgCreateuser = {
  encode(message: MsgCreateuser, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.newUser !== "") {
      writer.uint32(18).string(message.newUser);
    }
    if (message.alias !== "") {
      writer.uint32(26).string(message.alias);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateuser {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateuser } as MsgCreateuser;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.newUser = reader.string();
          break;
        case 3:
          message.alias = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateuser {
    const message = { ...baseMsgCreateuser } as MsgCreateuser;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.newUser !== undefined && object.newUser !== null) {
      message.newUser = String(object.newUser);
    } else {
      message.newUser = "";
    }
    if (object.alias !== undefined && object.alias !== null) {
      message.alias = String(object.alias);
    } else {
      message.alias = "";
    }
    return message;
  },

  toJSON(message: MsgCreateuser): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.newUser !== undefined && (obj.newUser = message.newUser);
    message.alias !== undefined && (obj.alias = message.alias);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateuser>): MsgCreateuser {
    const message = { ...baseMsgCreateuser } as MsgCreateuser;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.newUser !== undefined && object.newUser !== null) {
      message.newUser = object.newUser;
    } else {
      message.newUser = "";
    }
    if (object.alias !== undefined && object.alias !== null) {
      message.alias = object.alias;
    } else {
      message.alias = "";
    }
    return message;
  },
};

const baseMsgCreateuserResponse: object = {};

export const MsgCreateuserResponse = {
  encode(_: MsgCreateuserResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateuserResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateuserResponse } as MsgCreateuserResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateuserResponse {
    const message = { ...baseMsgCreateuserResponse } as MsgCreateuserResponse;
    return message;
  },

  toJSON(_: MsgCreateuserResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgCreateuserResponse>): MsgCreateuserResponse {
    const message = { ...baseMsgCreateuserResponse } as MsgCreateuserResponse;
    return message;
  },
};

const baseMsgBuyCardScheme: object = { creator: "", bid: "", buyer: "" };

export const MsgBuyCardScheme = {
  encode(message: MsgBuyCardScheme, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.bid !== "") {
      writer.uint32(18).string(message.bid);
    }
    if (message.buyer !== "") {
      writer.uint32(26).string(message.buyer);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgBuyCardScheme {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgBuyCardScheme } as MsgBuyCardScheme;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.bid = reader.string();
          break;
        case 3:
          message.buyer = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgBuyCardScheme {
    const message = { ...baseMsgBuyCardScheme } as MsgBuyCardScheme;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.bid !== undefined && object.bid !== null) {
      message.bid = String(object.bid);
    } else {
      message.bid = "";
    }
    if (object.buyer !== undefined && object.buyer !== null) {
      message.buyer = String(object.buyer);
    } else {
      message.buyer = "";
    }
    return message;
  },

  toJSON(message: MsgBuyCardScheme): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.bid !== undefined && (obj.bid = message.bid);
    message.buyer !== undefined && (obj.buyer = message.buyer);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgBuyCardScheme>): MsgBuyCardScheme {
    const message = { ...baseMsgBuyCardScheme } as MsgBuyCardScheme;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.bid !== undefined && object.bid !== null) {
      message.bid = object.bid;
    } else {
      message.bid = "";
    }
    if (object.buyer !== undefined && object.buyer !== null) {
      message.buyer = object.buyer;
    } else {
      message.buyer = "";
    }
    return message;
  },
};

const baseMsgBuyCardSchemeResponse: object = {};

export const MsgBuyCardSchemeResponse = {
  encode(
    _: MsgBuyCardSchemeResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgBuyCardSchemeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgBuyCardSchemeResponse,
    } as MsgBuyCardSchemeResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgBuyCardSchemeResponse {
    const message = {
      ...baseMsgBuyCardSchemeResponse,
    } as MsgBuyCardSchemeResponse;
    return message;
  },

  toJSON(_: MsgBuyCardSchemeResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgBuyCardSchemeResponse>
  ): MsgBuyCardSchemeResponse {
    const message = {
      ...baseMsgBuyCardSchemeResponse,
    } as MsgBuyCardSchemeResponse;
    return message;
  },
};

const baseMsgVoteCard: object = {
  creator: "",
  cardId: 0,
  voteType: "",
  voter: "",
};

export const MsgVoteCard = {
  encode(message: MsgVoteCard, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.cardId !== 0) {
      writer.uint32(16).uint64(message.cardId);
    }
    if (message.voteType !== "") {
      writer.uint32(26).string(message.voteType);
    }
    if (message.voter !== "") {
      writer.uint32(34).string(message.voter);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgVoteCard {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgVoteCard } as MsgVoteCard;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.cardId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.voteType = reader.string();
          break;
        case 4:
          message.voter = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgVoteCard {
    const message = { ...baseMsgVoteCard } as MsgVoteCard;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.cardId !== undefined && object.cardId !== null) {
      message.cardId = Number(object.cardId);
    } else {
      message.cardId = 0;
    }
    if (object.voteType !== undefined && object.voteType !== null) {
      message.voteType = String(object.voteType);
    } else {
      message.voteType = "";
    }
    if (object.voter !== undefined && object.voter !== null) {
      message.voter = String(object.voter);
    } else {
      message.voter = "";
    }
    return message;
  },

  toJSON(message: MsgVoteCard): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.cardId !== undefined && (obj.cardId = message.cardId);
    message.voteType !== undefined && (obj.voteType = message.voteType);
    message.voter !== undefined && (obj.voter = message.voter);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgVoteCard>): MsgVoteCard {
    const message = { ...baseMsgVoteCard } as MsgVoteCard;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.cardId !== undefined && object.cardId !== null) {
      message.cardId = object.cardId;
    } else {
      message.cardId = 0;
    }
    if (object.voteType !== undefined && object.voteType !== null) {
      message.voteType = object.voteType;
    } else {
      message.voteType = "";
    }
    if (object.voter !== undefined && object.voter !== null) {
      message.voter = object.voter;
    } else {
      message.voter = "";
    }
    return message;
  },
};

const baseMsgVoteCardResponse: object = {};

export const MsgVoteCardResponse = {
  encode(_: MsgVoteCardResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgVoteCardResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgVoteCardResponse } as MsgVoteCardResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgVoteCardResponse {
    const message = { ...baseMsgVoteCardResponse } as MsgVoteCardResponse;
    return message;
  },

  toJSON(_: MsgVoteCardResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgVoteCardResponse>): MsgVoteCardResponse {
    const message = { ...baseMsgVoteCardResponse } as MsgVoteCardResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  Createuser(request: MsgCreateuser): Promise<MsgCreateuserResponse>;
  BuyCardScheme(request: MsgBuyCardScheme): Promise<MsgBuyCardSchemeResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  VoteCard(request: MsgVoteCard): Promise<MsgVoteCardResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Createuser(request: MsgCreateuser): Promise<MsgCreateuserResponse> {
    const data = MsgCreateuser.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Msg",
      "Createuser",
      data
    );
    return promise.then((data) =>
      MsgCreateuserResponse.decode(new Reader(data))
    );
  }

  BuyCardScheme(request: MsgBuyCardScheme): Promise<MsgBuyCardSchemeResponse> {
    const data = MsgBuyCardScheme.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Msg",
      "BuyCardScheme",
      data
    );
    return promise.then((data) =>
      MsgBuyCardSchemeResponse.decode(new Reader(data))
    );
  }

  VoteCard(request: MsgVoteCard): Promise<MsgVoteCardResponse> {
    const data = MsgVoteCard.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Msg",
      "VoteCard",
      data
    );
    return promise.then((data) => MsgVoteCardResponse.decode(new Reader(data)));
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
