import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "DecentralCardGame.cardchain.cardchain";
export declare enum Status {
    scheme = 0,
    prototype = 1,
    trial = 2,
    permanent = 3,
    suspended = 4,
    banned = 5,
    bannedSoon = 6,
    bannedVerySoon = 7,
    none = 8,
    UNRECOGNIZED = -1
}
export declare function statusFromJSON(object: any): Status;
export declare function statusToJSON(object: Status): string;
export interface Card {
    owner: string;
    artist: string;
    content: Uint8Array;
    image: Uint8Array;
    fullArt: boolean;
    notes: string;
    status: Status;
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
