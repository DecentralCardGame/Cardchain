import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "DecentralCardGame.cardchain.cardchain";
export interface CollectionProposal {
    title: string;
    description: string;
    collectionId: number;
}
export declare const CollectionProposal: {
    encode(message: CollectionProposal, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): CollectionProposal;
    fromJSON(object: any): CollectionProposal;
    toJSON(message: CollectionProposal): unknown;
    fromPartial(object: DeepPartial<CollectionProposal>): CollectionProposal;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
