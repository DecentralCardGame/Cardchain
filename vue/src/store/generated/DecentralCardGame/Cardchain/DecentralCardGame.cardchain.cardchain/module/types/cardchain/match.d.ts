import { Outcome } from "../cardchain/tx";
import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "DecentralCardGame.cardchain.cardchain";
export interface Match {
    matchId: number;
    reporter: string;
    playerA: string;
    playerB: string;
    outcome: Outcome;
}
export declare const Match: {
    encode(message: Match, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Match;
    fromJSON(object: any): Match;
    toJSON(message: Match): unknown;
    fromPartial(object: DeepPartial<Match>): Match;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
