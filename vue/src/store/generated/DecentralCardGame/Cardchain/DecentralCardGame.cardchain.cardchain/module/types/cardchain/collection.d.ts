import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "DecentralCardGame.cardchain.cardchain";
export declare enum CStatus {
    design = 0,
    finalized = 1,
    active = 2,
    archived = 3,
    UNRECOGNIZED = -1
}
export declare function cStatusFromJSON(object: any): CStatus;
export declare function cStatusToJSON(object: CStatus): string;
export interface Collection {
    name: string;
    cards: number[];
    contributors: string[];
    story: string;
    artwork: Uint8Array;
    status: CStatus;
    expireBlock: number;
}
export declare const Collection: {
    encode(message: Collection, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Collection;
    fromJSON(object: any): Collection;
    toJSON(message: Collection): unknown;
    fromPartial(object: DeepPartial<Collection>): Collection;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
