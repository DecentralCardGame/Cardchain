import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "DecentralCardGame.cardchain.cardchain";
export interface CopyrightProposal {
    title: string;
    description: string;
    link: string;
    cardId: number;
}
export declare const CopyrightProposal: {
    encode(message: CopyrightProposal, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): CopyrightProposal;
    fromJSON(object: any): CopyrightProposal;
    toJSON(message: CopyrightProposal): unknown;
    fromPartial(object: DeepPartial<CopyrightProposal>): CopyrightProposal;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
