import { Writer, Reader } from "protobufjs/minimal";
import { VoteRight } from "../cardchain/vote_right";
export declare const protobufPackage = "DecentralCardGame.cardchain.cardchain";
export declare enum CouncilStatus {
    available = 0,
    unavailable = 1,
    openCouncil = 2,
    startedCouncil = 3,
    UNRECOGNIZED = -1
}
export declare function councilStatusFromJSON(object: any): CouncilStatus;
export declare function councilStatusToJSON(object: CouncilStatus): string;
export interface User {
    alias: string;
    ownedCardSchemes: number[];
    ownedPrototypes: number[];
    cards: number[];
    voteRights: VoteRight[];
    CouncilStatus: CouncilStatus;
    ReportMatches: boolean;
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
