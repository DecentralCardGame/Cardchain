import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "DecentralCardGame.cardchain.cardchain";
export declare enum Outcome {
    AWon = 0,
    BWon = 1,
    Draw = 2,
    Aborted = 3,
    UNRECOGNIZED = -1
}
export declare function outcomeFromJSON(object: any): Outcome;
export declare function outcomeToJSON(object: Outcome): string;
export interface MsgCreateuser {
    creator: string;
    newUser: string;
    alias: string;
}
export interface MsgCreateuserResponse {
}
export interface MsgBuyCardScheme {
    creator: string;
    bid: string;
    buyer: string;
}
export interface MsgBuyCardSchemeResponse {
}
export interface MsgVoteCard {
    creator: string;
    cardId: number;
    voteType: string;
    voter: string;
}
export interface MsgVoteCardResponse {
}
export interface MsgSaveCardContent {
    creator: string;
    cardId: number;
    content: Uint8Array;
    /**
     * bytes image = 4;
     *  string fullArt = 5;
     */
    notes: string;
    owner: string;
    artist: string;
}
export interface MsgSaveCardContentResponse {
}
export interface MsgTransferCard {
    creator: string;
    cardId: number;
    receiver: string;
}
export interface MsgTransferCardResponse {
}
export interface MsgDonateToCard {
    creator: string;
    cardId: number;
    donator: string;
    amount: string;
}
export interface MsgDonateToCardResponse {
}
export interface MsgAddArtwork {
    creator: string;
    cardId: number;
    image: Uint8Array;
    fullArt: boolean;
}
export interface MsgAddArtworkResponse {
}
export interface MsgSubmitCopyrightProposal {
    creator: string;
    cardId: number;
    description: string;
    link: string;
}
export interface MsgSubmitCopyrightProposalResponse {
}
export interface MsgChangeArtist {
    creator: string;
    cardID: number;
    artist: string;
}
export interface MsgChangeArtistResponse {
}
export interface MsgRegisterForCouncil {
    creator: string;
}
export interface MsgRegisterForCouncilResponse {
}
export interface MsgReportMatch {
    creator: string;
    playerA: string;
    playerB: string;
    cardsA: number[];
    cardsB: number[];
    outcome: Outcome;
}
export interface MsgReportMatchResponse {
    matchId: number;
}
export interface MsgSubmitMatchReporterProposal {
    creator: string;
    reporter: string;
    deposit: string;
    description: string;
}
export interface MsgSubmitMatchReporterProposalResponse {
}
export interface MsgApointMatchReporter {
    creator: string;
    reporter: string;
}
export interface MsgApointMatchReporterResponse {
}
export interface MsgCreateCollection {
    creator: string;
    name: string;
    contributors: string[];
    story: string;
    artwork: Uint8Array;
}
export interface MsgCreateCollectionResponse {
}
export interface MsgAddCardToCollection {
    creator: string;
    collectionId: number;
    cardId: number;
}
export interface MsgAddCardToCollectionResponse {
}
export interface MsgFinalizeCollection {
    creator: string;
    collectionId: number;
}
export interface MsgFinalizeCollectionResponse {
}
export declare const MsgCreateuser: {
    encode(message: MsgCreateuser, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateuser;
    fromJSON(object: any): MsgCreateuser;
    toJSON(message: MsgCreateuser): unknown;
    fromPartial(object: DeepPartial<MsgCreateuser>): MsgCreateuser;
};
export declare const MsgCreateuserResponse: {
    encode(_: MsgCreateuserResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateuserResponse;
    fromJSON(_: any): MsgCreateuserResponse;
    toJSON(_: MsgCreateuserResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateuserResponse>): MsgCreateuserResponse;
};
export declare const MsgBuyCardScheme: {
    encode(message: MsgBuyCardScheme, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgBuyCardScheme;
    fromJSON(object: any): MsgBuyCardScheme;
    toJSON(message: MsgBuyCardScheme): unknown;
    fromPartial(object: DeepPartial<MsgBuyCardScheme>): MsgBuyCardScheme;
};
export declare const MsgBuyCardSchemeResponse: {
    encode(_: MsgBuyCardSchemeResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgBuyCardSchemeResponse;
    fromJSON(_: any): MsgBuyCardSchemeResponse;
    toJSON(_: MsgBuyCardSchemeResponse): unknown;
    fromPartial(_: DeepPartial<MsgBuyCardSchemeResponse>): MsgBuyCardSchemeResponse;
};
export declare const MsgVoteCard: {
    encode(message: MsgVoteCard, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgVoteCard;
    fromJSON(object: any): MsgVoteCard;
    toJSON(message: MsgVoteCard): unknown;
    fromPartial(object: DeepPartial<MsgVoteCard>): MsgVoteCard;
};
export declare const MsgVoteCardResponse: {
    encode(_: MsgVoteCardResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgVoteCardResponse;
    fromJSON(_: any): MsgVoteCardResponse;
    toJSON(_: MsgVoteCardResponse): unknown;
    fromPartial(_: DeepPartial<MsgVoteCardResponse>): MsgVoteCardResponse;
};
export declare const MsgSaveCardContent: {
    encode(message: MsgSaveCardContent, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSaveCardContent;
    fromJSON(object: any): MsgSaveCardContent;
    toJSON(message: MsgSaveCardContent): unknown;
    fromPartial(object: DeepPartial<MsgSaveCardContent>): MsgSaveCardContent;
};
export declare const MsgSaveCardContentResponse: {
    encode(_: MsgSaveCardContentResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSaveCardContentResponse;
    fromJSON(_: any): MsgSaveCardContentResponse;
    toJSON(_: MsgSaveCardContentResponse): unknown;
    fromPartial(_: DeepPartial<MsgSaveCardContentResponse>): MsgSaveCardContentResponse;
};
export declare const MsgTransferCard: {
    encode(message: MsgTransferCard, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgTransferCard;
    fromJSON(object: any): MsgTransferCard;
    toJSON(message: MsgTransferCard): unknown;
    fromPartial(object: DeepPartial<MsgTransferCard>): MsgTransferCard;
};
export declare const MsgTransferCardResponse: {
    encode(_: MsgTransferCardResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgTransferCardResponse;
    fromJSON(_: any): MsgTransferCardResponse;
    toJSON(_: MsgTransferCardResponse): unknown;
    fromPartial(_: DeepPartial<MsgTransferCardResponse>): MsgTransferCardResponse;
};
export declare const MsgDonateToCard: {
    encode(message: MsgDonateToCard, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDonateToCard;
    fromJSON(object: any): MsgDonateToCard;
    toJSON(message: MsgDonateToCard): unknown;
    fromPartial(object: DeepPartial<MsgDonateToCard>): MsgDonateToCard;
};
export declare const MsgDonateToCardResponse: {
    encode(_: MsgDonateToCardResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDonateToCardResponse;
    fromJSON(_: any): MsgDonateToCardResponse;
    toJSON(_: MsgDonateToCardResponse): unknown;
    fromPartial(_: DeepPartial<MsgDonateToCardResponse>): MsgDonateToCardResponse;
};
export declare const MsgAddArtwork: {
    encode(message: MsgAddArtwork, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAddArtwork;
    fromJSON(object: any): MsgAddArtwork;
    toJSON(message: MsgAddArtwork): unknown;
    fromPartial(object: DeepPartial<MsgAddArtwork>): MsgAddArtwork;
};
export declare const MsgAddArtworkResponse: {
    encode(_: MsgAddArtworkResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAddArtworkResponse;
    fromJSON(_: any): MsgAddArtworkResponse;
    toJSON(_: MsgAddArtworkResponse): unknown;
    fromPartial(_: DeepPartial<MsgAddArtworkResponse>): MsgAddArtworkResponse;
};
export declare const MsgSubmitCopyrightProposal: {
    encode(message: MsgSubmitCopyrightProposal, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSubmitCopyrightProposal;
    fromJSON(object: any): MsgSubmitCopyrightProposal;
    toJSON(message: MsgSubmitCopyrightProposal): unknown;
    fromPartial(object: DeepPartial<MsgSubmitCopyrightProposal>): MsgSubmitCopyrightProposal;
};
export declare const MsgSubmitCopyrightProposalResponse: {
    encode(_: MsgSubmitCopyrightProposalResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSubmitCopyrightProposalResponse;
    fromJSON(_: any): MsgSubmitCopyrightProposalResponse;
    toJSON(_: MsgSubmitCopyrightProposalResponse): unknown;
    fromPartial(_: DeepPartial<MsgSubmitCopyrightProposalResponse>): MsgSubmitCopyrightProposalResponse;
};
export declare const MsgChangeArtist: {
    encode(message: MsgChangeArtist, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgChangeArtist;
    fromJSON(object: any): MsgChangeArtist;
    toJSON(message: MsgChangeArtist): unknown;
    fromPartial(object: DeepPartial<MsgChangeArtist>): MsgChangeArtist;
};
export declare const MsgChangeArtistResponse: {
    encode(_: MsgChangeArtistResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgChangeArtistResponse;
    fromJSON(_: any): MsgChangeArtistResponse;
    toJSON(_: MsgChangeArtistResponse): unknown;
    fromPartial(_: DeepPartial<MsgChangeArtistResponse>): MsgChangeArtistResponse;
};
export declare const MsgRegisterForCouncil: {
    encode(message: MsgRegisterForCouncil, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRegisterForCouncil;
    fromJSON(object: any): MsgRegisterForCouncil;
    toJSON(message: MsgRegisterForCouncil): unknown;
    fromPartial(object: DeepPartial<MsgRegisterForCouncil>): MsgRegisterForCouncil;
};
export declare const MsgRegisterForCouncilResponse: {
    encode(_: MsgRegisterForCouncilResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRegisterForCouncilResponse;
    fromJSON(_: any): MsgRegisterForCouncilResponse;
    toJSON(_: MsgRegisterForCouncilResponse): unknown;
    fromPartial(_: DeepPartial<MsgRegisterForCouncilResponse>): MsgRegisterForCouncilResponse;
};
export declare const MsgReportMatch: {
    encode(message: MsgReportMatch, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgReportMatch;
    fromJSON(object: any): MsgReportMatch;
    toJSON(message: MsgReportMatch): unknown;
    fromPartial(object: DeepPartial<MsgReportMatch>): MsgReportMatch;
};
export declare const MsgReportMatchResponse: {
    encode(message: MsgReportMatchResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgReportMatchResponse;
    fromJSON(object: any): MsgReportMatchResponse;
    toJSON(message: MsgReportMatchResponse): unknown;
    fromPartial(object: DeepPartial<MsgReportMatchResponse>): MsgReportMatchResponse;
};
export declare const MsgSubmitMatchReporterProposal: {
    encode(message: MsgSubmitMatchReporterProposal, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSubmitMatchReporterProposal;
    fromJSON(object: any): MsgSubmitMatchReporterProposal;
    toJSON(message: MsgSubmitMatchReporterProposal): unknown;
    fromPartial(object: DeepPartial<MsgSubmitMatchReporterProposal>): MsgSubmitMatchReporterProposal;
};
export declare const MsgSubmitMatchReporterProposalResponse: {
    encode(_: MsgSubmitMatchReporterProposalResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSubmitMatchReporterProposalResponse;
    fromJSON(_: any): MsgSubmitMatchReporterProposalResponse;
    toJSON(_: MsgSubmitMatchReporterProposalResponse): unknown;
    fromPartial(_: DeepPartial<MsgSubmitMatchReporterProposalResponse>): MsgSubmitMatchReporterProposalResponse;
};
export declare const MsgApointMatchReporter: {
    encode(message: MsgApointMatchReporter, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgApointMatchReporter;
    fromJSON(object: any): MsgApointMatchReporter;
    toJSON(message: MsgApointMatchReporter): unknown;
    fromPartial(object: DeepPartial<MsgApointMatchReporter>): MsgApointMatchReporter;
};
export declare const MsgApointMatchReporterResponse: {
    encode(_: MsgApointMatchReporterResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgApointMatchReporterResponse;
    fromJSON(_: any): MsgApointMatchReporterResponse;
    toJSON(_: MsgApointMatchReporterResponse): unknown;
    fromPartial(_: DeepPartial<MsgApointMatchReporterResponse>): MsgApointMatchReporterResponse;
};
export declare const MsgCreateCollection: {
    encode(message: MsgCreateCollection, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateCollection;
    fromJSON(object: any): MsgCreateCollection;
    toJSON(message: MsgCreateCollection): unknown;
    fromPartial(object: DeepPartial<MsgCreateCollection>): MsgCreateCollection;
};
export declare const MsgCreateCollectionResponse: {
    encode(_: MsgCreateCollectionResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateCollectionResponse;
    fromJSON(_: any): MsgCreateCollectionResponse;
    toJSON(_: MsgCreateCollectionResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateCollectionResponse>): MsgCreateCollectionResponse;
};
export declare const MsgAddCardToCollection: {
    encode(message: MsgAddCardToCollection, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAddCardToCollection;
    fromJSON(object: any): MsgAddCardToCollection;
    toJSON(message: MsgAddCardToCollection): unknown;
    fromPartial(object: DeepPartial<MsgAddCardToCollection>): MsgAddCardToCollection;
};
export declare const MsgAddCardToCollectionResponse: {
    encode(_: MsgAddCardToCollectionResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAddCardToCollectionResponse;
    fromJSON(_: any): MsgAddCardToCollectionResponse;
    toJSON(_: MsgAddCardToCollectionResponse): unknown;
    fromPartial(_: DeepPartial<MsgAddCardToCollectionResponse>): MsgAddCardToCollectionResponse;
};
export declare const MsgFinalizeCollection: {
    encode(message: MsgFinalizeCollection, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgFinalizeCollection;
    fromJSON(object: any): MsgFinalizeCollection;
    toJSON(message: MsgFinalizeCollection): unknown;
    fromPartial(object: DeepPartial<MsgFinalizeCollection>): MsgFinalizeCollection;
};
export declare const MsgFinalizeCollectionResponse: {
    encode(_: MsgFinalizeCollectionResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgFinalizeCollectionResponse;
    fromJSON(_: any): MsgFinalizeCollectionResponse;
    toJSON(_: MsgFinalizeCollectionResponse): unknown;
    fromPartial(_: DeepPartial<MsgFinalizeCollectionResponse>): MsgFinalizeCollectionResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    Createuser(request: MsgCreateuser): Promise<MsgCreateuserResponse>;
    BuyCardScheme(request: MsgBuyCardScheme): Promise<MsgBuyCardSchemeResponse>;
    VoteCard(request: MsgVoteCard): Promise<MsgVoteCardResponse>;
    SaveCardContent(request: MsgSaveCardContent): Promise<MsgSaveCardContentResponse>;
    TransferCard(request: MsgTransferCard): Promise<MsgTransferCardResponse>;
    DonateToCard(request: MsgDonateToCard): Promise<MsgDonateToCardResponse>;
    AddArtwork(request: MsgAddArtwork): Promise<MsgAddArtworkResponse>;
    SubmitCopyrightProposal(request: MsgSubmitCopyrightProposal): Promise<MsgSubmitCopyrightProposalResponse>;
    ChangeArtist(request: MsgChangeArtist): Promise<MsgChangeArtistResponse>;
    RegisterForCouncil(request: MsgRegisterForCouncil): Promise<MsgRegisterForCouncilResponse>;
    ReportMatch(request: MsgReportMatch): Promise<MsgReportMatchResponse>;
    SubmitMatchReporterProposal(request: MsgSubmitMatchReporterProposal): Promise<MsgSubmitMatchReporterProposalResponse>;
    ApointMatchReporter(request: MsgApointMatchReporter): Promise<MsgApointMatchReporterResponse>;
    CreateCollection(request: MsgCreateCollection): Promise<MsgCreateCollectionResponse>;
    AddCardToCollection(request: MsgAddCardToCollection): Promise<MsgAddCardToCollectionResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    FinalizeCollection(request: MsgFinalizeCollection): Promise<MsgFinalizeCollectionResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    Createuser(request: MsgCreateuser): Promise<MsgCreateuserResponse>;
    BuyCardScheme(request: MsgBuyCardScheme): Promise<MsgBuyCardSchemeResponse>;
    VoteCard(request: MsgVoteCard): Promise<MsgVoteCardResponse>;
    SaveCardContent(request: MsgSaveCardContent): Promise<MsgSaveCardContentResponse>;
    TransferCard(request: MsgTransferCard): Promise<MsgTransferCardResponse>;
    DonateToCard(request: MsgDonateToCard): Promise<MsgDonateToCardResponse>;
    AddArtwork(request: MsgAddArtwork): Promise<MsgAddArtworkResponse>;
    SubmitCopyrightProposal(request: MsgSubmitCopyrightProposal): Promise<MsgSubmitCopyrightProposalResponse>;
    ChangeArtist(request: MsgChangeArtist): Promise<MsgChangeArtistResponse>;
    RegisterForCouncil(request: MsgRegisterForCouncil): Promise<MsgRegisterForCouncilResponse>;
    ReportMatch(request: MsgReportMatch): Promise<MsgReportMatchResponse>;
    SubmitMatchReporterProposal(request: MsgSubmitMatchReporterProposal): Promise<MsgSubmitMatchReporterProposalResponse>;
    ApointMatchReporter(request: MsgApointMatchReporter): Promise<MsgApointMatchReporterResponse>;
    CreateCollection(request: MsgCreateCollection): Promise<MsgCreateCollectionResponse>;
    AddCardToCollection(request: MsgAddCardToCollection): Promise<MsgAddCardToCollectionResponse>;
    FinalizeCollection(request: MsgFinalizeCollection): Promise<MsgFinalizeCollectionResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
