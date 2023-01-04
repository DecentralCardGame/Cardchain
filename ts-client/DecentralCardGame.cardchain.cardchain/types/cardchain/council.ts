/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "DecentralCardGame.cardchain.cardchain";

export enum Response {
  Yes = 0,
  No = 1,
  Suggestion = 2,
  UNRECOGNIZED = -1,
}

export function responseFromJSON(object: any): Response {
  switch (object) {
    case 0:
    case "Yes":
      return Response.Yes;
    case 1:
    case "No":
      return Response.No;
    case 2:
    case "Suggestion":
      return Response.Suggestion;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Response.UNRECOGNIZED;
  }
}

export function responseToJSON(object: Response): string {
  switch (object) {
    case Response.Yes:
      return "Yes";
    case Response.No:
      return "No";
    case Response.Suggestion:
      return "Suggestion";
    case Response.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export enum CouncelingStatus {
  councilOpen = 0,
  councilCreated = 1,
  councilClosed = 2,
  commited = 3,
  revealed = 4,
  suggestionsMade = 5,
  UNRECOGNIZED = -1,
}

export function councelingStatusFromJSON(object: any): CouncelingStatus {
  switch (object) {
    case 0:
    case "councilOpen":
      return CouncelingStatus.councilOpen;
    case 1:
    case "councilCreated":
      return CouncelingStatus.councilCreated;
    case 2:
    case "councilClosed":
      return CouncelingStatus.councilClosed;
    case 3:
    case "commited":
      return CouncelingStatus.commited;
    case 4:
    case "revealed":
      return CouncelingStatus.revealed;
    case 5:
    case "suggestionsMade":
      return CouncelingStatus.suggestionsMade;
    case -1:
    case "UNRECOGNIZED":
    default:
      return CouncelingStatus.UNRECOGNIZED;
  }
}

export function councelingStatusToJSON(object: CouncelingStatus): string {
  switch (object) {
    case CouncelingStatus.councilOpen:
      return "councilOpen";
    case CouncelingStatus.councilCreated:
      return "councilCreated";
    case CouncelingStatus.councilClosed:
      return "councilClosed";
    case CouncelingStatus.commited:
      return "commited";
    case CouncelingStatus.revealed:
      return "revealed";
    case CouncelingStatus.suggestionsMade:
      return "suggestionsMade";
    case CouncelingStatus.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface Council {
  cardId: number;
  voters: string[];
  hashResponses: WrapHashResponse[];
  clearResponses: WrapClearResponse[];
  treasury: string;
  status: CouncelingStatus;
  trialStart: number;
}

export interface WrapClearResponse {
  user: string;
  response: Response;
  suggestion: string;
}

export interface WrapHashResponse {
  user: string;
  hash: string;
}

function createBaseCouncil(): Council {
  return { cardId: 0, voters: [], hashResponses: [], clearResponses: [], treasury: "", status: 0, trialStart: 0 };
}

export const Council = {
  encode(message: Council, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.cardId !== 0) {
      writer.uint32(8).uint64(message.cardId);
    }
    for (const v of message.voters) {
      writer.uint32(18).string(v!);
    }
    for (const v of message.hashResponses) {
      WrapHashResponse.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.clearResponses) {
      WrapClearResponse.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    if (message.treasury !== "") {
      writer.uint32(42).string(message.treasury);
    }
    if (message.status !== 0) {
      writer.uint32(48).int32(message.status);
    }
    if (message.trialStart !== 0) {
      writer.uint32(56).uint64(message.trialStart);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Council {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCouncil();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.cardId = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.voters.push(reader.string());
          break;
        case 3:
          message.hashResponses.push(WrapHashResponse.decode(reader, reader.uint32()));
          break;
        case 4:
          message.clearResponses.push(WrapClearResponse.decode(reader, reader.uint32()));
          break;
        case 5:
          message.treasury = reader.string();
          break;
        case 6:
          message.status = reader.int32() as any;
          break;
        case 7:
          message.trialStart = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Council {
    return {
      cardId: isSet(object.cardId) ? Number(object.cardId) : 0,
      voters: Array.isArray(object?.voters) ? object.voters.map((e: any) => String(e)) : [],
      hashResponses: Array.isArray(object?.hashResponses)
        ? object.hashResponses.map((e: any) => WrapHashResponse.fromJSON(e))
        : [],
      clearResponses: Array.isArray(object?.clearResponses)
        ? object.clearResponses.map((e: any) => WrapClearResponse.fromJSON(e))
        : [],
      treasury: isSet(object.treasury) ? String(object.treasury) : "",
      status: isSet(object.status) ? councelingStatusFromJSON(object.status) : 0,
      trialStart: isSet(object.trialStart) ? Number(object.trialStart) : 0,
    };
  },

  toJSON(message: Council): unknown {
    const obj: any = {};
    message.cardId !== undefined && (obj.cardId = Math.round(message.cardId));
    if (message.voters) {
      obj.voters = message.voters.map((e) => e);
    } else {
      obj.voters = [];
    }
    if (message.hashResponses) {
      obj.hashResponses = message.hashResponses.map((e) => e ? WrapHashResponse.toJSON(e) : undefined);
    } else {
      obj.hashResponses = [];
    }
    if (message.clearResponses) {
      obj.clearResponses = message.clearResponses.map((e) => e ? WrapClearResponse.toJSON(e) : undefined);
    } else {
      obj.clearResponses = [];
    }
    message.treasury !== undefined && (obj.treasury = message.treasury);
    message.status !== undefined && (obj.status = councelingStatusToJSON(message.status));
    message.trialStart !== undefined && (obj.trialStart = Math.round(message.trialStart));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Council>, I>>(object: I): Council {
    const message = createBaseCouncil();
    message.cardId = object.cardId ?? 0;
    message.voters = object.voters?.map((e) => e) || [];
    message.hashResponses = object.hashResponses?.map((e) => WrapHashResponse.fromPartial(e)) || [];
    message.clearResponses = object.clearResponses?.map((e) => WrapClearResponse.fromPartial(e)) || [];
    message.treasury = object.treasury ?? "";
    message.status = object.status ?? 0;
    message.trialStart = object.trialStart ?? 0;
    return message;
  },
};

function createBaseWrapClearResponse(): WrapClearResponse {
  return { user: "", response: 0, suggestion: "" };
}

export const WrapClearResponse = {
  encode(message: WrapClearResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.user !== "") {
      writer.uint32(10).string(message.user);
    }
    if (message.response !== 0) {
      writer.uint32(16).int32(message.response);
    }
    if (message.suggestion !== "") {
      writer.uint32(26).string(message.suggestion);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): WrapClearResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWrapClearResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.user = reader.string();
          break;
        case 2:
          message.response = reader.int32() as any;
          break;
        case 3:
          message.suggestion = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): WrapClearResponse {
    return {
      user: isSet(object.user) ? String(object.user) : "",
      response: isSet(object.response) ? responseFromJSON(object.response) : 0,
      suggestion: isSet(object.suggestion) ? String(object.suggestion) : "",
    };
  },

  toJSON(message: WrapClearResponse): unknown {
    const obj: any = {};
    message.user !== undefined && (obj.user = message.user);
    message.response !== undefined && (obj.response = responseToJSON(message.response));
    message.suggestion !== undefined && (obj.suggestion = message.suggestion);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<WrapClearResponse>, I>>(object: I): WrapClearResponse {
    const message = createBaseWrapClearResponse();
    message.user = object.user ?? "";
    message.response = object.response ?? 0;
    message.suggestion = object.suggestion ?? "";
    return message;
  },
};

function createBaseWrapHashResponse(): WrapHashResponse {
  return { user: "", hash: "" };
}

export const WrapHashResponse = {
  encode(message: WrapHashResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.user !== "") {
      writer.uint32(10).string(message.user);
    }
    if (message.hash !== "") {
      writer.uint32(18).string(message.hash);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): WrapHashResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWrapHashResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.user = reader.string();
          break;
        case 2:
          message.hash = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): WrapHashResponse {
    return { user: isSet(object.user) ? String(object.user) : "", hash: isSet(object.hash) ? String(object.hash) : "" };
  },

  toJSON(message: WrapHashResponse): unknown {
    const obj: any = {};
    message.user !== undefined && (obj.user = message.user);
    message.hash !== undefined && (obj.hash = message.hash);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<WrapHashResponse>, I>>(object: I): WrapHashResponse {
    const message = createBaseWrapHashResponse();
    message.user = object.user ?? "";
    message.hash = object.hash ?? "";
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
