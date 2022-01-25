import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "DecentralCardGame.cardchain.cardchain";
export interface VotingResult {
    cardId: number;
    fairEnoughVotes: number;
    overpoweredVotes: number;
    underpoweredVotes: number;
    inappropriateVotes: number;
    result: string;
}
export declare const VotingResult: {
    encode(message: VotingResult, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): VotingResult;
    fromJSON(object: any): VotingResult;
    toJSON(message: VotingResult): unknown;
    fromPartial(object: DeepPartial<VotingResult>): VotingResult;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
