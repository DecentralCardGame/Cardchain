import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../cardchain/params";
export declare const protobufPackage = "DecentralCardGame.cardchain.cardchain";
/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}
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
export declare const QueryParamsRequest: {
    encode(_: QueryParamsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest;
    fromJSON(_: any): QueryParamsRequest;
    toJSON(_: QueryParamsRequest): unknown;
    fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest;
};
export declare const QueryParamsResponse: {
    encode(message: QueryParamsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse;
    fromJSON(object: any): QueryParamsResponse;
    toJSON(message: QueryParamsResponse): unknown;
    fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse;
};
export declare const QueryQCardRequest: {
    encode(message: QueryQCardRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryQCardRequest;
    fromJSON(object: any): QueryQCardRequest;
    toJSON(message: QueryQCardRequest): unknown;
    fromPartial(object: DeepPartial<QueryQCardRequest>): QueryQCardRequest;
};
export declare const QueryQCardResponse: {
    encode(message: QueryQCardResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryQCardResponse;
    fromJSON(object: any): QueryQCardResponse;
    toJSON(message: QueryQCardResponse): unknown;
    fromPartial(object: DeepPartial<QueryQCardResponse>): QueryQCardResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Parameters queries the parameters of the module. */
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    /** Queries a list of QCard items. */
    QCard(request: QueryQCardRequest): Promise<QueryQCardResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    QCard(request: QueryQCardRequest): Promise<QueryQCardResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
