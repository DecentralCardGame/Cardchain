import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "DecentralCardGame.cardchain.cardchain";
export declare enum SellOfferStatus {
    open = 0,
    sold = 1,
    removed = 2,
    UNRECOGNIZED = -1
}
export declare function sellOfferStatusFromJSON(object: any): SellOfferStatus;
export declare function sellOfferStatusToJSON(object: SellOfferStatus): string;
export interface SellOffer {
    seller: string;
    buyer: string;
    card: string;
    status: SellOfferStatus;
}
export declare const SellOffer: {
    encode(message: SellOffer, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): SellOffer;
    fromJSON(object: any): SellOffer;
    toJSON(message: SellOffer): unknown;
    fromPartial(object: DeepPartial<SellOffer>): SellOffer;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
