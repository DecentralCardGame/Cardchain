/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "DecentralCardGame.cardchain.cardchain";

export interface MatchReporterProposal {
  title: string;
  description: string;
  reporter: string;
}

function createBaseMatchReporterProposal(): MatchReporterProposal {
  return { title: "", description: "", reporter: "" };
}

export const MatchReporterProposal = {
  encode(message: MatchReporterProposal, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.title !== "") {
      writer.uint32(10).string(message.title);
    }
    if (message.description !== "") {
      writer.uint32(18).string(message.description);
    }
    if (message.reporter !== "") {
      writer.uint32(26).string(message.reporter);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MatchReporterProposal {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMatchReporterProposal();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.title = reader.string();
          break;
        case 2:
          message.description = reader.string();
          break;
        case 3:
          message.reporter = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MatchReporterProposal {
    return {
      title: isSet(object.title) ? String(object.title) : "",
      description: isSet(object.description) ? String(object.description) : "",
      reporter: isSet(object.reporter) ? String(object.reporter) : "",
    };
  },

  toJSON(message: MatchReporterProposal): unknown {
    const obj: any = {};
    message.title !== undefined && (obj.title = message.title);
    message.description !== undefined && (obj.description = message.description);
    message.reporter !== undefined && (obj.reporter = message.reporter);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MatchReporterProposal>, I>>(object: I): MatchReporterProposal {
    const message = createBaseMatchReporterProposal();
    message.title = object.title ?? "";
    message.description = object.description ?? "";
    message.reporter = object.reporter ?? "";
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
