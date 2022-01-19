import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "DecentralCardGame.cardchain.cardchain";
export interface Card {
    owner: string;
    content: Uint8Array;
    image: Uint8Array;
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
export declare const Card: {
    encode(message: Card, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Card;
    fromJSON(object: any): Card;
    toJSON(message: Card): unknown;
    fromPartial(object: DeepPartial<Card>): Card;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
