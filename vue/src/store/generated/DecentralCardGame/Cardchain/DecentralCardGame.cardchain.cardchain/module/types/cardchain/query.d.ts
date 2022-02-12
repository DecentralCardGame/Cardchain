import { Status } from "../cardchain/card";
import { CouncilStatus } from "../cardchain/user";
import { Outcome } from "../cardchain/tx";
import { CStatus } from "../cardchain/collection";
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../cardchain/params";
import { VoteRight } from "../cardchain/vote_right";
import { VotingResults } from "../cardchain/voting_results";
export declare const protobufPackage = "DecentralCardGame.cardchain.cardchain";
/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}
/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
    /** params holds all the parameters of this module. */
    params: Params | undefined;
}
export interface QueryQCardRequest {
    cardId: string;
}
export interface QueryQCardResponse {
    owner: string;
    content: string;
    artist: string;
    image: string;
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
export interface QueryQCardContentRequest {
    cardId: string;
}
export interface QueryQCardContentResponse {
    content: string;
}
export interface QueryQUserRequest {
    address: string;
}
export interface QueryQUserResponse {
    alias: string;
    ownedCardSchemes: number[];
    ownedPrototypes: number[];
    cards: number[];
    voteRights: VoteRight[];
    councilStatus: CouncilStatus;
    reportMatches: boolean;
}
export interface QueryQCardchainInfoRequest {
}
export interface QueryQCardchainInfoResponse {
    cardAuctionPrice: string;
    activeCollections: number[];
}
export interface QueryQVotingResultsRequest {
}
export interface QueryQVotingResultsResponse {
    lastVotingResults: VotingResults | undefined;
}
export interface QueryQVotableCardsRequest {
    address: string;
}
export interface QueryQVotableCardsResponse {
    unregistered: boolean;
    noVoteRights: boolean;
    voteRights: VoteRight[];
}
export interface QueryQCardsRequest {
    owner: string;
    status: Status;
    cardType: string;
    classes: string;
    sortBy: string;
    nameContains: string;
    keywordsContains: string;
    notesContains: string;
}
export interface QueryQCardsResponse {
    cardsList: number[];
}
export interface QueryQMatchRequest {
    matchId: number;
}
export interface QueryQMatchResponse {
    timestamp: number;
    reporter: string;
    playerA: string;
    playerB: string;
    outcome: Outcome;
}
export interface QueryQCollectionRequest {
    collectionId: number;
}
export interface QueryQCollectionResponse {
    name: string;
    cards: number[];
    contributors: string[];
    story: string;
    artwork: Uint8Array;
    status: CStatus;
    expireBlock: number;
}
export declare const QueryParamsRequest: {
    encode(_: QueryParamsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest;
    fromJSON(_: any): QueryParamsRequest;
    toJSON(_: QueryParamsRequest): unknown;
    fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest;
};
export declare const QueryParamsResponse: {
    encode(message: QueryParamsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse;
    fromJSON(object: any): QueryParamsResponse;
    toJSON(message: QueryParamsResponse): unknown;
    fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse;
};
export declare const QueryQCardRequest: {
    encode(message: QueryQCardRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryQCardRequest;
    fromJSON(object: any): QueryQCardRequest;
    toJSON(message: QueryQCardRequest): unknown;
    fromPartial(object: DeepPartial<QueryQCardRequest>): QueryQCardRequest;
};
export declare const QueryQCardResponse: {
    encode(message: QueryQCardResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryQCardResponse;
    fromJSON(object: any): QueryQCardResponse;
    toJSON(message: QueryQCardResponse): unknown;
    fromPartial(object: DeepPartial<QueryQCardResponse>): QueryQCardResponse;
};
export declare const QueryQCardContentRequest: {
    encode(message: QueryQCardContentRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryQCardContentRequest;
    fromJSON(object: any): QueryQCardContentRequest;
    toJSON(message: QueryQCardContentRequest): unknown;
    fromPartial(object: DeepPartial<QueryQCardContentRequest>): QueryQCardContentRequest;
};
export declare const QueryQCardContentResponse: {
    encode(message: QueryQCardContentResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryQCardContentResponse;
    fromJSON(object: any): QueryQCardContentResponse;
    toJSON(message: QueryQCardContentResponse): unknown;
    fromPartial(object: DeepPartial<QueryQCardContentResponse>): QueryQCardContentResponse;
};
export declare const QueryQUserRequest: {
    encode(message: QueryQUserRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryQUserRequest;
    fromJSON(object: any): QueryQUserRequest;
    toJSON(message: QueryQUserRequest): unknown;
    fromPartial(object: DeepPartial<QueryQUserRequest>): QueryQUserRequest;
};
export declare const QueryQUserResponse: {
    encode(message: QueryQUserResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryQUserResponse;
    fromJSON(object: any): QueryQUserResponse;
    toJSON(message: QueryQUserResponse): unknown;
    fromPartial(object: DeepPartial<QueryQUserResponse>): QueryQUserResponse;
};
export declare const QueryQCardchainInfoRequest: {
    encode(_: QueryQCardchainInfoRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryQCardchainInfoRequest;
    fromJSON(_: any): QueryQCardchainInfoRequest;
    toJSON(_: QueryQCardchainInfoRequest): unknown;
    fromPartial(_: DeepPartial<QueryQCardchainInfoRequest>): QueryQCardchainInfoRequest;
};
export declare const QueryQCardchainInfoResponse: {
    encode(message: QueryQCardchainInfoResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryQCardchainInfoResponse;
    fromJSON(object: any): QueryQCardchainInfoResponse;
    toJSON(message: QueryQCardchainInfoResponse): unknown;
    fromPartial(object: DeepPartial<QueryQCardchainInfoResponse>): QueryQCardchainInfoResponse;
};
export declare const QueryQVotingResultsRequest: {
    encode(_: QueryQVotingResultsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryQVotingResultsRequest;
    fromJSON(_: any): QueryQVotingResultsRequest;
    toJSON(_: QueryQVotingResultsRequest): unknown;
    fromPartial(_: DeepPartial<QueryQVotingResultsRequest>): QueryQVotingResultsRequest;
};
export declare const QueryQVotingResultsResponse: {
    encode(message: QueryQVotingResultsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryQVotingResultsResponse;
    fromJSON(object: any): QueryQVotingResultsResponse;
    toJSON(message: QueryQVotingResultsResponse): unknown;
    fromPartial(object: DeepPartial<QueryQVotingResultsResponse>): QueryQVotingResultsResponse;
};
export declare const QueryQVotableCardsRequest: {
    encode(message: QueryQVotableCardsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryQVotableCardsRequest;
    fromJSON(object: any): QueryQVotableCardsRequest;
    toJSON(message: QueryQVotableCardsRequest): unknown;
    fromPartial(object: DeepPartial<QueryQVotableCardsRequest>): QueryQVotableCardsRequest;
};
export declare const QueryQVotableCardsResponse: {
    encode(message: QueryQVotableCardsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryQVotableCardsResponse;
    fromJSON(object: any): QueryQVotableCardsResponse;
    toJSON(message: QueryQVotableCardsResponse): unknown;
    fromPartial(object: DeepPartial<QueryQVotableCardsResponse>): QueryQVotableCardsResponse;
};
export declare const QueryQCardsRequest: {
    encode(message: QueryQCardsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryQCardsRequest;
    fromJSON(object: any): QueryQCardsRequest;
    toJSON(message: QueryQCardsRequest): unknown;
    fromPartial(object: DeepPartial<QueryQCardsRequest>): QueryQCardsRequest;
};
export declare const QueryQCardsResponse: {
    encode(message: QueryQCardsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryQCardsResponse;
    fromJSON(object: any): QueryQCardsResponse;
    toJSON(message: QueryQCardsResponse): unknown;
    fromPartial(object: DeepPartial<QueryQCardsResponse>): QueryQCardsResponse;
};
export declare const QueryQMatchRequest: {
    encode(message: QueryQMatchRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryQMatchRequest;
    fromJSON(object: any): QueryQMatchRequest;
    toJSON(message: QueryQMatchRequest): unknown;
    fromPartial(object: DeepPartial<QueryQMatchRequest>): QueryQMatchRequest;
};
export declare const QueryQMatchResponse: {
    encode(message: QueryQMatchResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryQMatchResponse;
    fromJSON(object: any): QueryQMatchResponse;
    toJSON(message: QueryQMatchResponse): unknown;
    fromPartial(object: DeepPartial<QueryQMatchResponse>): QueryQMatchResponse;
};
export declare const QueryQCollectionRequest: {
    encode(message: QueryQCollectionRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryQCollectionRequest;
    fromJSON(object: any): QueryQCollectionRequest;
    toJSON(message: QueryQCollectionRequest): unknown;
    fromPartial(object: DeepPartial<QueryQCollectionRequest>): QueryQCollectionRequest;
};
export declare const QueryQCollectionResponse: {
    encode(message: QueryQCollectionResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryQCollectionResponse;
    fromJSON(object: any): QueryQCollectionResponse;
    toJSON(message: QueryQCollectionResponse): unknown;
    fromPartial(object: DeepPartial<QueryQCollectionResponse>): QueryQCollectionResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Parameters queries the parameters of the module. */
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    /** Queries a list of QCard items. */
    QCard(request: QueryQCardRequest): Promise<QueryQCardResponse>;
    /** Queries a list of QCardContent items. */
    QCardContent(request: QueryQCardContentRequest): Promise<QueryQCardContentResponse>;
    /** Queries a list of QUser items. */
    QUser(request: QueryQUserRequest): Promise<QueryQUserResponse>;
    /** Queries a list of QCardchainInfo items. */
    QCardchainInfo(request: QueryQCardchainInfoRequest): Promise<QueryQCardchainInfoResponse>;
    /** Queries a list of QVotingResults items. */
    QVotingResults(request: QueryQVotingResultsRequest): Promise<QueryQVotingResultsResponse>;
    /** Queries a list of QVotableCards items. */
    QVotableCards(request: QueryQVotableCardsRequest): Promise<QueryQVotableCardsResponse>;
    /** Queries a list of QCards items. */
    QCards(request: QueryQCardsRequest): Promise<QueryQCardsResponse>;
    /** Queries a list of QMatch items. */
    QMatch(request: QueryQMatchRequest): Promise<QueryQMatchResponse>;
    /** Queries a list of QCollection items. */
    QCollection(request: QueryQCollectionRequest): Promise<QueryQCollectionResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    QCard(request: QueryQCardRequest): Promise<QueryQCardResponse>;
    QCardContent(request: QueryQCardContentRequest): Promise<QueryQCardContentResponse>;
    QUser(request: QueryQUserRequest): Promise<QueryQUserResponse>;
    QCardchainInfo(request: QueryQCardchainInfoRequest): Promise<QueryQCardchainInfoResponse>;
    QVotingResults(request: QueryQVotingResultsRequest): Promise<QueryQVotingResultsResponse>;
    QVotableCards(request: QueryQVotableCardsRequest): Promise<QueryQVotableCardsResponse>;
    QCards(request: QueryQCardsRequest): Promise<QueryQCardsResponse>;
    QMatch(request: QueryQMatchRequest): Promise<QueryQMatchResponse>;
    QCollection(request: QueryQCollectionRequest): Promise<QueryQCollectionResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
