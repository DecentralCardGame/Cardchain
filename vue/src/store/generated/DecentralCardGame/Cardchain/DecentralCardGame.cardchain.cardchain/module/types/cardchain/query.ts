/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../cardchain/params";

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
  card: Uint8Array;
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

const baseQueryQCardResponse: object = {};

export const QueryQCardResponse = {
  encode(
    message: QueryQCardResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.card.length !== 0) {
      writer.uint32(10).bytes(message.card);
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
          message.card = reader.bytes();
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
    if (object.card !== undefined && object.card !== null) {
      message.card = bytesFromBase64(object.card);
    }
    return message;
  },

  toJSON(message: QueryQCardResponse): unknown {
    const obj: any = {};
    message.card !== undefined &&
      (obj.card = base64FromBytes(
        message.card !== undefined ? message.card : new Uint8Array()
      ));
    return obj;
  },

  fromPartial(object: DeepPartial<QueryQCardResponse>): QueryQCardResponse {
    const message = { ...baseQueryQCardResponse } as QueryQCardResponse;
    if (object.card !== undefined && object.card !== null) {
      message.card = object.card;
    } else {
      message.card = new Uint8Array();
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

const atob: (b64: string) => string =
  globalThis.atob ||
  ((b64) => globalThis.Buffer.from(b64, "base64").toString("binary"));
function bytesFromBase64(b64: string): Uint8Array {
  const bin = atob(b64);
  const arr = new Uint8Array(bin.length);
  for (let i = 0; i < bin.length; ++i) {
    arr[i] = bin.charCodeAt(i);
  }
  return arr;
}

const btoa: (bin: string) => string =
  globalThis.btoa ||
  ((bin) => globalThis.Buffer.from(bin, "binary").toString("base64"));
function base64FromBytes(arr: Uint8Array): string {
  const bin: string[] = [];
  for (let i = 0; i < arr.byteLength; ++i) {
    bin.push(String.fromCharCode(arr[i]));
  }
  return btoa(bin.join(""));
}

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
