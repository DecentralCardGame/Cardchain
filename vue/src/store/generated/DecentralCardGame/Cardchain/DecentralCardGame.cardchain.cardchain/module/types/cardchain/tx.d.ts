import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "DecentralCardGame.cardchain.cardchain";
export interface MsgCreateuser {
    creator: string;
    newUser: string;
    alias: string;
}
export interface MsgCreateuserResponse {
}
export interface MsgBuyCardScheme {
    creator: string;
    bid: string;
    buyer: string;
}
export interface MsgBuyCardSchemeResponse {
}
export declare const MsgCreateuser: {
    encode(message: MsgCreateuser, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateuser;
    fromJSON(object: any): MsgCreateuser;
    toJSON(message: MsgCreateuser): unknown;
    fromPartial(object: DeepPartial<MsgCreateuser>): MsgCreateuser;
};
export declare const MsgCreateuserResponse: {
    encode(_: MsgCreateuserResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateuserResponse;
    fromJSON(_: any): MsgCreateuserResponse;
    toJSON(_: MsgCreateuserResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateuserResponse>): MsgCreateuserResponse;
};
export declare const MsgBuyCardScheme: {
    encode(message: MsgBuyCardScheme, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgBuyCardScheme;
    fromJSON(object: any): MsgBuyCardScheme;
    toJSON(message: MsgBuyCardScheme): unknown;
    fromPartial(object: DeepPartial<MsgBuyCardScheme>): MsgBuyCardScheme;
};
export declare const MsgBuyCardSchemeResponse: {
    encode(_: MsgBuyCardSchemeResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgBuyCardSchemeResponse;
    fromJSON(_: any): MsgBuyCardSchemeResponse;
    toJSON(_: MsgBuyCardSchemeResponse): unknown;
    fromPartial(_: DeepPartial<MsgBuyCardSchemeResponse>): MsgBuyCardSchemeResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    Createuser(request: MsgCreateuser): Promise<MsgCreateuserResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    BuyCardScheme(request: MsgBuyCardScheme): Promise<MsgBuyCardSchemeResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    Createuser(request: MsgCreateuser): Promise<MsgCreateuserResponse>;
    BuyCardScheme(request: MsgBuyCardScheme): Promise<MsgBuyCardSchemeResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
