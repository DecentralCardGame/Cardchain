/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Params } from "../cardchain/params";
import { VoteRight } from "../cardchain/vote_right";

export const protobufPackage = "DecentralCardGame.cardchain.cardchain";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryQCardRequest {
  cardId: string;
}

export interface QueryQCardResponse {
  owner: string;
  content: string;
  image: string;
  fullArt: string;
  notes: string;
  status: string;
  votePool: string;
  fairEnoughVotes: number;
  overpoweredVotes: number;
  underpoweredVotes: number;
  inappropriateVotes: number;
  nerflevel: number;
}

export interface QueryQCardContentRequest {
  cardId: string;
}

export interface QueryQCardContentResponse {
  content: string;
}

export interface QueryQUserRequest {
  address: string;
}

export interface QueryQUserResponse {
  alias: string;
  ownedCardSchemes: number[];
  ownedCards: number[];
  voteRights: VoteRight[];
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
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

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryQCardRequest: object = { cardId: "" };

export const QueryQCardRequest = {
  encode(message: QueryQCardRequest, writer: Writer = Writer.create()): Writer {
    if (message.cardId !== "") {
      writer.uint32(10).string(message.cardId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryQCardRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryQCardRequest } as QueryQCardRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.cardId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryQCardRequest {
    const message = { ...baseQueryQCardRequest } as QueryQCardRequest;
    if (object.cardId !== undefined && object.cardId !== null) {
      message.cardId = String(object.cardId);
    } else {
      message.cardId = "";
    }
    return message;
  },

  toJSON(message: QueryQCardRequest): unknown {
    const obj: any = {};
    message.cardId !== undefined && (obj.cardId = message.cardId);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryQCardRequest>): QueryQCardRequest {
    const message = { ...baseQueryQCardRequest } as QueryQCardRequest;
    if (object.cardId !== undefined && object.cardId !== null) {
      message.cardId = object.cardId;
    } else {
      message.cardId = "";
    }
    return message;
  },
};

const baseQueryQCardResponse: object = {
  owner: "",
  content: "",
  image: "",
  fullArt: "",
  notes: "",
  status: "",
  votePool: "",
  fairEnoughVotes: 0,
  overpoweredVotes: 0,
  underpoweredVotes: 0,
  inappropriateVotes: 0,
  nerflevel: 0,
};

export const QueryQCardResponse = {
  encode(
    message: QueryQCardResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.content !== "") {
      writer.uint32(18).string(message.content);
    }
    if (message.image !== "") {
      writer.uint32(26).string(message.image);
    }
    if (message.fullArt !== "") {
      writer.uint32(34).string(message.fullArt);
    }
    if (message.notes !== "") {
      writer.uint32(42).string(message.notes);
    }
    if (message.status !== "") {
      writer.uint32(50).string(message.status);
    }
    if (message.votePool !== "") {
      writer.uint32(58).string(message.votePool);
    }
    if (message.fairEnoughVotes !== 0) {
      writer.uint32(64).uint64(message.fairEnoughVotes);
    }
    if (message.overpoweredVotes !== 0) {
      writer.uint32(72).uint64(message.overpoweredVotes);
    }
    if (message.underpoweredVotes !== 0) {
      writer.uint32(80).uint64(message.underpoweredVotes);
    }
    if (message.inappropriateVotes !== 0) {
      writer.uint32(88).uint64(message.inappropriateVotes);
    }
    if (message.nerflevel !== 0) {
      writer.uint32(96).int64(message.nerflevel);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryQCardResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryQCardResponse } as QueryQCardResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.content = reader.string();
          break;
        case 3:
          message.image = reader.string();
          break;
        case 4:
          message.fullArt = reader.string();
          break;
        case 5:
          message.notes = reader.string();
          break;
        case 6:
          message.status = reader.string();
          break;
        case 7:
          message.votePool = reader.string();
          break;
        case 8:
          message.fairEnoughVotes = longToNumber(reader.uint64() as Long);
          break;
        case 9:
          message.overpoweredVotes = longToNumber(reader.uint64() as Long);
          break;
        case 10:
          message.underpoweredVotes = longToNumber(reader.uint64() as Long);
          break;
        case 11:
          message.inappropriateVotes = longToNumber(reader.uint64() as Long);
          break;
        case 12:
          message.nerflevel = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryQCardResponse {
    const message = { ...baseQueryQCardResponse } as QueryQCardResponse;
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.content !== undefined && object.content !== null) {
      message.content = String(object.content);
    } else {
      message.content = "";
    }
    if (object.image !== undefined && object.image !== null) {
      message.image = String(object.image);
    } else {
      message.image = "";
    }
    if (object.fullArt !== undefined && object.fullArt !== null) {
      message.fullArt = String(object.fullArt);
    } else {
      message.fullArt = "";
    }
    if (object.notes !== undefined && object.notes !== null) {
      message.notes = String(object.notes);
    } else {
      message.notes = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = String(object.status);
    } else {
      message.status = "";
    }
    if (object.votePool !== undefined && object.votePool !== null) {
      message.votePool = String(object.votePool);
    } else {
      message.votePool = "";
    }
    if (
      object.fairEnoughVotes !== undefined &&
      object.fairEnoughVotes !== null
    ) {
      message.fairEnoughVotes = Number(object.fairEnoughVotes);
    } else {
      message.fairEnoughVotes = 0;
    }
    if (
      object.overpoweredVotes !== undefined &&
      object.overpoweredVotes !== null
    ) {
      message.overpoweredVotes = Number(object.overpoweredVotes);
    } else {
      message.overpoweredVotes = 0;
    }
    if (
      object.underpoweredVotes !== undefined &&
      object.underpoweredVotes !== null
    ) {
      message.underpoweredVotes = Number(object.underpoweredVotes);
    } else {
      message.underpoweredVotes = 0;
    }
    if (
      object.inappropriateVotes !== undefined &&
      object.inappropriateVotes !== null
    ) {
      message.inappropriateVotes = Number(object.inappropriateVotes);
    } else {
      message.inappropriateVotes = 0;
    }
    if (object.nerflevel !== undefined && object.nerflevel !== null) {
      message.nerflevel = Number(object.nerflevel);
    } else {
      message.nerflevel = 0;
    }
    return message;
  },

  toJSON(message: QueryQCardResponse): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.content !== undefined && (obj.content = message.content);
    message.image !== undefined && (obj.image = message.image);
    message.fullArt !== undefined && (obj.fullArt = message.fullArt);
    message.notes !== undefined && (obj.notes = message.notes);
    message.status !== undefined && (obj.status = message.status);
    message.votePool !== undefined && (obj.votePool = message.votePool);
    message.fairEnoughVotes !== undefined &&
      (obj.fairEnoughVotes = message.fairEnoughVotes);
    message.overpoweredVotes !== undefined &&
      (obj.overpoweredVotes = message.overpoweredVotes);
    message.underpoweredVotes !== undefined &&
      (obj.underpoweredVotes = message.underpoweredVotes);
    message.inappropriateVotes !== undefined &&
      (obj.inappropriateVotes = message.inappropriateVotes);
    message.nerflevel !== undefined && (obj.nerflevel = message.nerflevel);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryQCardResponse>): QueryQCardResponse {
    const message = { ...baseQueryQCardResponse } as QueryQCardResponse;
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.content !== undefined && object.content !== null) {
      message.content = object.content;
    } else {
      message.content = "";
    }
    if (object.image !== undefined && object.image !== null) {
      message.image = object.image;
    } else {
      message.image = "";
    }
    if (object.fullArt !== undefined && object.fullArt !== null) {
      message.fullArt = object.fullArt;
    } else {
      message.fullArt = "";
    }
    if (object.notes !== undefined && object.notes !== null) {
      message.notes = object.notes;
    } else {
      message.notes = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = "";
    }
    if (object.votePool !== undefined && object.votePool !== null) {
      message.votePool = object.votePool;
    } else {
      message.votePool = "";
    }
    if (
      object.fairEnoughVotes !== undefined &&
      object.fairEnoughVotes !== null
    ) {
      message.fairEnoughVotes = object.fairEnoughVotes;
    } else {
      message.fairEnoughVotes = 0;
    }
    if (
      object.overpoweredVotes !== undefined &&
      object.overpoweredVotes !== null
    ) {
      message.overpoweredVotes = object.overpoweredVotes;
    } else {
      message.overpoweredVotes = 0;
    }
    if (
      object.underpoweredVotes !== undefined &&
      object.underpoweredVotes !== null
    ) {
      message.underpoweredVotes = object.underpoweredVotes;
    } else {
      message.underpoweredVotes = 0;
    }
    if (
      object.inappropriateVotes !== undefined &&
      object.inappropriateVotes !== null
    ) {
      message.inappropriateVotes = object.inappropriateVotes;
    } else {
      message.inappropriateVotes = 0;
    }
    if (object.nerflevel !== undefined && object.nerflevel !== null) {
      message.nerflevel = object.nerflevel;
    } else {
      message.nerflevel = 0;
    }
    return message;
  },
};

const baseQueryQCardContentRequest: object = { cardId: "" };

export const QueryQCardContentRequest = {
  encode(
    message: QueryQCardContentRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.cardId !== "") {
      writer.uint32(10).string(message.cardId);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryQCardContentRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryQCardContentRequest,
    } as QueryQCardContentRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.cardId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryQCardContentRequest {
    const message = {
      ...baseQueryQCardContentRequest,
    } as QueryQCardContentRequest;
    if (object.cardId !== undefined && object.cardId !== null) {
      message.cardId = String(object.cardId);
    } else {
      message.cardId = "";
    }
    return message;
  },

  toJSON(message: QueryQCardContentRequest): unknown {
    const obj: any = {};
    message.cardId !== undefined && (obj.cardId = message.cardId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryQCardContentRequest>
  ): QueryQCardContentRequest {
    const message = {
      ...baseQueryQCardContentRequest,
    } as QueryQCardContentRequest;
    if (object.cardId !== undefined && object.cardId !== null) {
      message.cardId = object.cardId;
    } else {
      message.cardId = "";
    }
    return message;
  },
};

const baseQueryQCardContentResponse: object = { content: "" };

export const QueryQCardContentResponse = {
  encode(
    message: QueryQCardContentResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.content !== "") {
      writer.uint32(10).string(message.content);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryQCardContentResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryQCardContentResponse,
    } as QueryQCardContentResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.content = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryQCardContentResponse {
    const message = {
      ...baseQueryQCardContentResponse,
    } as QueryQCardContentResponse;
    if (object.content !== undefined && object.content !== null) {
      message.content = String(object.content);
    } else {
      message.content = "";
    }
    return message;
  },

  toJSON(message: QueryQCardContentResponse): unknown {
    const obj: any = {};
    message.content !== undefined && (obj.content = message.content);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryQCardContentResponse>
  ): QueryQCardContentResponse {
    const message = {
      ...baseQueryQCardContentResponse,
    } as QueryQCardContentResponse;
    if (object.content !== undefined && object.content !== null) {
      message.content = object.content;
    } else {
      message.content = "";
    }
    return message;
  },
};

const baseQueryQUserRequest: object = { address: "" };

export const QueryQUserRequest = {
  encode(message: QueryQUserRequest, writer: Writer = Writer.create()): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryQUserRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryQUserRequest } as QueryQUserRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryQUserRequest {
    const message = { ...baseQueryQUserRequest } as QueryQUserRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: QueryQUserRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryQUserRequest>): QueryQUserRequest {
    const message = { ...baseQueryQUserRequest } as QueryQUserRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseQueryQUserResponse: object = {
  alias: "",
  ownedCardSchemes: 0,
  ownedCards: 0,
};

export const QueryQUserResponse = {
  encode(
    message: QueryQUserResponse,
    writer: Writer = Writer.create()
  ): Writer {
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

  decode(input: Reader | Uint8Array, length?: number): QueryQUserResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryQUserResponse } as QueryQUserResponse;
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

  fromJSON(object: any): QueryQUserResponse {
    const message = { ...baseQueryQUserResponse } as QueryQUserResponse;
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

  toJSON(message: QueryQUserResponse): unknown {
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

  fromPartial(object: DeepPartial<QueryQUserResponse>): QueryQUserResponse {
    const message = { ...baseQueryQUserResponse } as QueryQUserResponse;
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

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of QCard items. */
  QCard(request: QueryQCardRequest): Promise<QueryQCardResponse>;
  /** Queries a list of QCardContent items. */
  QCardContent(
    request: QueryQCardContentRequest
  ): Promise<QueryQCardContentResponse>;
  /** Queries a list of QUser items. */
  QUser(request: QueryQUserRequest): Promise<QueryQUserResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  QCard(request: QueryQCardRequest): Promise<QueryQCardResponse> {
    const data = QueryQCardRequest.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Query",
      "QCard",
      data
    );
    return promise.then((data) => QueryQCardResponse.decode(new Reader(data)));
  }

  QCardContent(
    request: QueryQCardContentRequest
  ): Promise<QueryQCardContentResponse> {
    const data = QueryQCardContentRequest.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Query",
      "QCardContent",
      data
    );
    return promise.then((data) =>
      QueryQCardContentResponse.decode(new Reader(data))
    );
  }

  QUser(request: QueryQUserRequest): Promise<QueryQUserResponse> {
    const data = QueryQUserRequest.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Query",
      "QUser",
      data
    );
    return promise.then((data) => QueryQUserResponse.decode(new Reader(data)));
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
