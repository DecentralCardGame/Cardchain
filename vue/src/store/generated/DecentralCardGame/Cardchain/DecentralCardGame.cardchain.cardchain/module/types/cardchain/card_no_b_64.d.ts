import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "DecentralCardGame.cardchain.cardchain";
export interface CardNoB64 {
    owner: string;
    content: string;
    image: string;
    notes: string;
    fullArt: boolean;
    status: string;
    votePool: string;
    fairEnoughVotes: number;
    overpoweredVotes: number;
    underpoweredVotes: number;
    inappropriateVotes: number;
    nerflevel: number;
}
export declare const CardNoB64: {
    encode(message: CardNoB64, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): CardNoB64;
    fromJSON(object: any): CardNoB64;
    toJSON(message: CardNoB64): unknown;
    fromPartial(object: DeepPartial<CardNoB64>): CardNoB64;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
