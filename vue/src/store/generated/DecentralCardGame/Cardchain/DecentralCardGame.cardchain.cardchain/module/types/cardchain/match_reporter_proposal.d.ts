import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "DecentralCardGame.cardchain.cardchain";
export interface MatchReporterProposal {
    title: string;
    description: string;
    reporter: string;
}
export declare const MatchReporterProposal: {
    encode(message: MatchReporterProposal, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MatchReporterProposal;
    fromJSON(object: any): MatchReporterProposal;
    toJSON(message: MatchReporterProposal): unknown;
    fromPartial(object: DeepPartial<MatchReporterProposal>): MatchReporterProposal;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
