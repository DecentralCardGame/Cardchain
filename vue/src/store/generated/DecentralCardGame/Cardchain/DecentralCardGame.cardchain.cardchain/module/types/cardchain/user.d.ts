import { Writer, Reader } from "protobufjs/minimal";
import { VoteRight } from "../cardchain/vote_right";
export declare const protobufPackage = "DecentralCardGame.cardchain.cardchain";
export interface User {
    alias: string;
    ownedCardSchemes: number[];
    ownedCards: number[];
    voteRights: VoteRight[];
}
export declare const User: {
    encode(message: User, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): User;
    fromJSON(object: any): User;
    toJSON(message: User): unknown;
    fromPartial(object: DeepPartial<User>): User;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
