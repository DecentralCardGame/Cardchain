import { Writer, Reader } from "protobufjs/minimal";
import { VotingResult } from "../cardchain/voting_result";
export declare const protobufPackage = "DecentralCardGame.cardchain.cardchain";
export interface VotingResults {
    totalVotes: number;
    totalFairEnoughVotes: number;
    totalOverpoweredVotes: number;
    totalUnderpoweredVotes: number;
    totalInappropriateVotes: number;
    cardResults: VotingResult[];
    notes: string;
}
export declare const VotingResults: {
    encode(message: VotingResults, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): VotingResults;
    fromJSON(object: any): VotingResults;
    toJSON(message: VotingResults): unknown;
    fromPartial(object: DeepPartial<VotingResults>): VotingResults;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
