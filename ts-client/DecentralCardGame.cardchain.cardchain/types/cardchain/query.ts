/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { OutpCard } from "./card";
import { CStatus, cStatusFromJSON, cStatusToJSON, OutpCollection } from "./collection";
import { Council } from "./council";
import { Match } from "./match";
import { Params } from "./params";
import { SellOffer, SellOfferStatus, sellOfferStatusFromJSON, sellOfferStatusToJSON } from "./sell_offer";
import { Server } from "./server";
import { Outcome, outcomeFromJSON, outcomeToJSON } from "./tx";
import { User } from "./user";
import { VoteRight } from "./vote_right";
import { VotingResults } from "./voting_results";

export const protobufPackage = "DecentralCardGame.cardchain.cardchain";

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

export interface QueryQCardContentRequest {
  cardId: string;
}

export interface QueryQCardContentResponse {
  content: string;
  hash: string;
}

export interface QueryQUserRequest {
  address: string;
}

export interface QueryQCardchainInfoRequest {
}

export interface QueryQCardchainInfoResponse {
  cardAuctionPrice: string;
  activeCollections: number[];
  cardsNumber: number;
  matchesNumber: number;
  sellOffersNumber: number;
  councilsNumber: number;
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
  status: QueryQCardsRequest_Status;
  cardType: string;
  classes: string;
  sortBy: string;
  nameContains: string;
  keywordsContains: string;
  notesContains: string;
}

export enum QueryQCardsRequest_Status {
  scheme = 0,
  prototype = 1,
  trial = 2,
  permanent = 3,
  suspended = 4,
  banned = 5,
  bannedSoon = 6,
  bannedVerySoon = 7,
  none = 8,
  playable = 9,
  unplayable = 10,
  UNRECOGNIZED = -1,
}

export function queryQCardsRequest_StatusFromJSON(object: any): QueryQCardsRequest_Status {
  switch (object) {
    case 0:
    case "scheme":
      return QueryQCardsRequest_Status.scheme;
    case 1:
    case "prototype":
      return QueryQCardsRequest_Status.prototype;
    case 2:
    case "trial":
      return QueryQCardsRequest_Status.trial;
    case 3:
    case "permanent":
      return QueryQCardsRequest_Status.permanent;
    case 4:
    case "suspended":
      return QueryQCardsRequest_Status.suspended;
    case 5:
    case "banned":
      return QueryQCardsRequest_Status.banned;
    case 6:
    case "bannedSoon":
      return QueryQCardsRequest_Status.bannedSoon;
    case 7:
    case "bannedVerySoon":
      return QueryQCardsRequest_Status.bannedVerySoon;
    case 8:
    case "none":
      return QueryQCardsRequest_Status.none;
    case 9:
    case "playable":
      return QueryQCardsRequest_Status.playable;
    case 10:
    case "unplayable":
      return QueryQCardsRequest_Status.unplayable;
    case -1:
    case "UNRECOGNIZED":
    default:
      return QueryQCardsRequest_Status.UNRECOGNIZED;
  }
}

export function queryQCardsRequest_StatusToJSON(object: QueryQCardsRequest_Status): string {
  switch (object) {
    case QueryQCardsRequest_Status.scheme:
      return "scheme";
    case QueryQCardsRequest_Status.prototype:
      return "prototype";
    case QueryQCardsRequest_Status.trial:
      return "trial";
    case QueryQCardsRequest_Status.permanent:
      return "permanent";
    case QueryQCardsRequest_Status.suspended:
      return "suspended";
    case QueryQCardsRequest_Status.banned:
      return "banned";
    case QueryQCardsRequest_Status.bannedSoon:
      return "bannedSoon";
    case QueryQCardsRequest_Status.bannedVerySoon:
      return "bannedVerySoon";
    case QueryQCardsRequest_Status.none:
      return "none";
    case QueryQCardsRequest_Status.playable:
      return "playable";
    case QueryQCardsRequest_Status.unplayable:
      return "unplayable";
    case QueryQCardsRequest_Status.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface QueryQCardsResponse {
  cardsList: number[];
}

export interface QueryQMatchRequest {
  matchId: number;
}

export interface QueryQCollectionRequest {
  collectionId: number;
}

export interface QueryQSellOfferRequest {
  sellOfferId: number;
}

export interface QueryQCouncilRequest {
  councilId: number;
}

export interface QueryQMatchesRequest {
  timestampDown: number;
  timestampUp: number;
  containsUsers: string[];
  reporter: string;
  outcome: Outcome;
  cardsPlayed: number[];
  ignore: IgnoreMatches | undefined;
}

export interface IgnoreMatches {
  outcome: boolean;
}

export interface QueryQMatchesResponse {
  matchesList: number[];
  matches: Match[];
}

export interface QueryQSellOffersRequest {
  priceDown: string;
  priceUp: string;
  seller: string;
  buyer: string;
  card: number;
  status: SellOfferStatus;
  ignore: IgnoreSellOffers | undefined;
}

export interface IgnoreSellOffers {
  status: boolean;
  card: boolean;
}

export interface QueryQSellOffersResponse {
  sellOffersIds: number[];
  sellOffers: SellOffer[];
}

export interface QueryQServerRequest {
  id: number;
}

export interface QueryQServerResponse {
}

export interface QueryQCollectionsRequest {
  status: CStatus;
  ignoreStatus: boolean;
  contributors: string[];
  containsCards: number[];
  owner: string;
}

export interface QueryQCollectionsResponse {
  collectionIds: number[];
}

export interface QueryRarityDistributionRequest {
  collectionId: number;
}

export interface QueryRarityDistributionResponse {
  current: number[];
  wanted: number[];
}

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    return {};
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(_: I): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  },
};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return { params: undefined };
}

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    return { params: isSet(object.params) ? Params.fromJSON(object.params) : undefined };
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(object: I): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseQueryQCardRequest(): QueryQCardRequest {
  return { cardId: "" };
}

export const QueryQCardRequest = {
  encode(message: QueryQCardRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.cardId !== "") {
      writer.uint32(10).string(message.cardId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryQCardRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryQCardRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.cardId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryQCardRequest {
    return { cardId: isSet(object.cardId) ? String(object.cardId) : "" };
  },

  toJSON(message: QueryQCardRequest): unknown {
    const obj: any = {};
    message.cardId !== undefined && (obj.cardId = message.cardId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryQCardRequest>, I>>(object: I): QueryQCardRequest {
    const message = createBaseQueryQCardRequest();
    message.cardId = object.cardId ?? "";
    return message;
  },
};

function createBaseQueryQCardContentRequest(): QueryQCardContentRequest {
  return { cardId: "" };
}

export const QueryQCardContentRequest = {
  encode(message: QueryQCardContentRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.cardId !== "") {
      writer.uint32(10).string(message.cardId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryQCardContentRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryQCardContentRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.cardId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryQCardContentRequest {
    return { cardId: isSet(object.cardId) ? String(object.cardId) : "" };
  },

  toJSON(message: QueryQCardContentRequest): unknown {
    const obj: any = {};
    message.cardId !== undefined && (obj.cardId = message.cardId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryQCardContentRequest>, I>>(object: I): QueryQCardContentRequest {
    const message = createBaseQueryQCardContentRequest();
    message.cardId = object.cardId ?? "";
    return message;
  },
};

function createBaseQueryQCardContentResponse(): QueryQCardContentResponse {
  return { content: "", hash: "" };
}

export const QueryQCardContentResponse = {
  encode(message: QueryQCardContentResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.content !== "") {
      writer.uint32(10).string(message.content);
    }
    if (message.hash !== "") {
      writer.uint32(18).string(message.hash);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryQCardContentResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryQCardContentResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.content = reader.string();
          break;
        case 2:
          message.hash = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryQCardContentResponse {
    return {
      content: isSet(object.content) ? String(object.content) : "",
      hash: isSet(object.hash) ? String(object.hash) : "",
    };
  },

  toJSON(message: QueryQCardContentResponse): unknown {
    const obj: any = {};
    message.content !== undefined && (obj.content = message.content);
    message.hash !== undefined && (obj.hash = message.hash);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryQCardContentResponse>, I>>(object: I): QueryQCardContentResponse {
    const message = createBaseQueryQCardContentResponse();
    message.content = object.content ?? "";
    message.hash = object.hash ?? "";
    return message;
  },
};

function createBaseQueryQUserRequest(): QueryQUserRequest {
  return { address: "" };
}

export const QueryQUserRequest = {
  encode(message: QueryQUserRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryQUserRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryQUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryQUserRequest {
    return { address: isSet(object.address) ? String(object.address) : "" };
  },

  toJSON(message: QueryQUserRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryQUserRequest>, I>>(object: I): QueryQUserRequest {
    const message = createBaseQueryQUserRequest();
    message.address = object.address ?? "";
    return message;
  },
};

function createBaseQueryQCardchainInfoRequest(): QueryQCardchainInfoRequest {
  return {};
}

export const QueryQCardchainInfoRequest = {
  encode(_: QueryQCardchainInfoRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryQCardchainInfoRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryQCardchainInfoRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryQCardchainInfoRequest {
    return {};
  },

  toJSON(_: QueryQCardchainInfoRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryQCardchainInfoRequest>, I>>(_: I): QueryQCardchainInfoRequest {
    const message = createBaseQueryQCardchainInfoRequest();
    return message;
  },
};

function createBaseQueryQCardchainInfoResponse(): QueryQCardchainInfoResponse {
  return {
    cardAuctionPrice: "",
    activeCollections: [],
    cardsNumber: 0,
    matchesNumber: 0,
    sellOffersNumber: 0,
    councilsNumber: 0,
  };
}

export const QueryQCardchainInfoResponse = {
  encode(message: QueryQCardchainInfoResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.cardAuctionPrice !== "") {
      writer.uint32(10).string(message.cardAuctionPrice);
    }
    writer.uint32(18).fork();
    for (const v of message.activeCollections) {
      writer.uint64(v);
    }
    writer.ldelim();
    if (message.cardsNumber !== 0) {
      writer.uint32(24).uint64(message.cardsNumber);
    }
    if (message.matchesNumber !== 0) {
      writer.uint32(32).uint64(message.matchesNumber);
    }
    if (message.sellOffersNumber !== 0) {
      writer.uint32(40).uint64(message.sellOffersNumber);
    }
    if (message.councilsNumber !== 0) {
      writer.uint32(48).uint64(message.councilsNumber);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryQCardchainInfoResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryQCardchainInfoResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.cardAuctionPrice = reader.string();
          break;
        case 2:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.activeCollections.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.activeCollections.push(longToNumber(reader.uint64() as Long));
          }
          break;
        case 3:
          message.cardsNumber = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.matchesNumber = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.sellOffersNumber = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.councilsNumber = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryQCardchainInfoResponse {
    return {
      cardAuctionPrice: isSet(object.cardAuctionPrice) ? String(object.cardAuctionPrice) : "",
      activeCollections: Array.isArray(object?.activeCollections)
        ? object.activeCollections.map((e: any) => Number(e))
        : [],
      cardsNumber: isSet(object.cardsNumber) ? Number(object.cardsNumber) : 0,
      matchesNumber: isSet(object.matchesNumber) ? Number(object.matchesNumber) : 0,
      sellOffersNumber: isSet(object.sellOffersNumber) ? Number(object.sellOffersNumber) : 0,
      councilsNumber: isSet(object.councilsNumber) ? Number(object.councilsNumber) : 0,
    };
  },

  toJSON(message: QueryQCardchainInfoResponse): unknown {
    const obj: any = {};
    message.cardAuctionPrice !== undefined && (obj.cardAuctionPrice = message.cardAuctionPrice);
    if (message.activeCollections) {
      obj.activeCollections = message.activeCollections.map((e) => Math.round(e));
    } else {
      obj.activeCollections = [];
    }
    message.cardsNumber !== undefined && (obj.cardsNumber = Math.round(message.cardsNumber));
    message.matchesNumber !== undefined && (obj.matchesNumber = Math.round(message.matchesNumber));
    message.sellOffersNumber !== undefined && (obj.sellOffersNumber = Math.round(message.sellOffersNumber));
    message.councilsNumber !== undefined && (obj.councilsNumber = Math.round(message.councilsNumber));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryQCardchainInfoResponse>, I>>(object: I): QueryQCardchainInfoResponse {
    const message = createBaseQueryQCardchainInfoResponse();
    message.cardAuctionPrice = object.cardAuctionPrice ?? "";
    message.activeCollections = object.activeCollections?.map((e) => e) || [];
    message.cardsNumber = object.cardsNumber ?? 0;
    message.matchesNumber = object.matchesNumber ?? 0;
    message.sellOffersNumber = object.sellOffersNumber ?? 0;
    message.councilsNumber = object.councilsNumber ?? 0;
    return message;
  },
};

function createBaseQueryQVotingResultsRequest(): QueryQVotingResultsRequest {
  return {};
}

export const QueryQVotingResultsRequest = {
  encode(_: QueryQVotingResultsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryQVotingResultsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryQVotingResultsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryQVotingResultsRequest {
    return {};
  },

  toJSON(_: QueryQVotingResultsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryQVotingResultsRequest>, I>>(_: I): QueryQVotingResultsRequest {
    const message = createBaseQueryQVotingResultsRequest();
    return message;
  },
};

function createBaseQueryQVotingResultsResponse(): QueryQVotingResultsResponse {
  return { lastVotingResults: undefined };
}

export const QueryQVotingResultsResponse = {
  encode(message: QueryQVotingResultsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.lastVotingResults !== undefined) {
      VotingResults.encode(message.lastVotingResults, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryQVotingResultsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryQVotingResultsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.lastVotingResults = VotingResults.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryQVotingResultsResponse {
    return {
      lastVotingResults: isSet(object.lastVotingResults) ? VotingResults.fromJSON(object.lastVotingResults) : undefined,
    };
  },

  toJSON(message: QueryQVotingResultsResponse): unknown {
    const obj: any = {};
    message.lastVotingResults !== undefined && (obj.lastVotingResults = message.lastVotingResults
      ? VotingResults.toJSON(message.lastVotingResults)
      : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryQVotingResultsResponse>, I>>(object: I): QueryQVotingResultsResponse {
    const message = createBaseQueryQVotingResultsResponse();
    message.lastVotingResults = (object.lastVotingResults !== undefined && object.lastVotingResults !== null)
      ? VotingResults.fromPartial(object.lastVotingResults)
      : undefined;
    return message;
  },
};

function createBaseQueryQVotableCardsRequest(): QueryQVotableCardsRequest {
  return { address: "" };
}

export const QueryQVotableCardsRequest = {
  encode(message: QueryQVotableCardsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryQVotableCardsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryQVotableCardsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryQVotableCardsRequest {
    return { address: isSet(object.address) ? String(object.address) : "" };
  },

  toJSON(message: QueryQVotableCardsRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryQVotableCardsRequest>, I>>(object: I): QueryQVotableCardsRequest {
    const message = createBaseQueryQVotableCardsRequest();
    message.address = object.address ?? "";
    return message;
  },
};

function createBaseQueryQVotableCardsResponse(): QueryQVotableCardsResponse {
  return { unregistered: false, noVoteRights: false, voteRights: [] };
}

export const QueryQVotableCardsResponse = {
  encode(message: QueryQVotableCardsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.unregistered === true) {
      writer.uint32(8).bool(message.unregistered);
    }
    if (message.noVoteRights === true) {
      writer.uint32(16).bool(message.noVoteRights);
    }
    for (const v of message.voteRights) {
      VoteRight.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryQVotableCardsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryQVotableCardsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.unregistered = reader.bool();
          break;
        case 2:
          message.noVoteRights = reader.bool();
          break;
        case 3:
          message.voteRights.push(VoteRight.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryQVotableCardsResponse {
    return {
      unregistered: isSet(object.unregistered) ? Boolean(object.unregistered) : false,
      noVoteRights: isSet(object.noVoteRights) ? Boolean(object.noVoteRights) : false,
      voteRights: Array.isArray(object?.voteRights) ? object.voteRights.map((e: any) => VoteRight.fromJSON(e)) : [],
    };
  },

  toJSON(message: QueryQVotableCardsResponse): unknown {
    const obj: any = {};
    message.unregistered !== undefined && (obj.unregistered = message.unregistered);
    message.noVoteRights !== undefined && (obj.noVoteRights = message.noVoteRights);
    if (message.voteRights) {
      obj.voteRights = message.voteRights.map((e) => e ? VoteRight.toJSON(e) : undefined);
    } else {
      obj.voteRights = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryQVotableCardsResponse>, I>>(object: I): QueryQVotableCardsResponse {
    const message = createBaseQueryQVotableCardsResponse();
    message.unregistered = object.unregistered ?? false;
    message.noVoteRights = object.noVoteRights ?? false;
    message.voteRights = object.voteRights?.map((e) => VoteRight.fromPartial(e)) || [];
    return message;
  },
};

function createBaseQueryQCardsRequest(): QueryQCardsRequest {
  return {
    owner: "",
    status: 0,
    cardType: "",
    classes: "",
    sortBy: "",
    nameContains: "",
    keywordsContains: "",
    notesContains: "",
  };
}

export const QueryQCardsRequest = {
  encode(message: QueryQCardsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.status !== 0) {
      writer.uint32(16).int32(message.status);
    }
    if (message.cardType !== "") {
      writer.uint32(26).string(message.cardType);
    }
    if (message.classes !== "") {
      writer.uint32(34).string(message.classes);
    }
    if (message.sortBy !== "") {
      writer.uint32(42).string(message.sortBy);
    }
    if (message.nameContains !== "") {
      writer.uint32(50).string(message.nameContains);
    }
    if (message.keywordsContains !== "") {
      writer.uint32(58).string(message.keywordsContains);
    }
    if (message.notesContains !== "") {
      writer.uint32(66).string(message.notesContains);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryQCardsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryQCardsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.status = reader.int32() as any;
          break;
        case 3:
          message.cardType = reader.string();
          break;
        case 4:
          message.classes = reader.string();
          break;
        case 5:
          message.sortBy = reader.string();
          break;
        case 6:
          message.nameContains = reader.string();
          break;
        case 7:
          message.keywordsContains = reader.string();
          break;
        case 8:
          message.notesContains = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryQCardsRequest {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      status: isSet(object.status) ? queryQCardsRequest_StatusFromJSON(object.status) : 0,
      cardType: isSet(object.cardType) ? String(object.cardType) : "",
      classes: isSet(object.classes) ? String(object.classes) : "",
      sortBy: isSet(object.sortBy) ? String(object.sortBy) : "",
      nameContains: isSet(object.nameContains) ? String(object.nameContains) : "",
      keywordsContains: isSet(object.keywordsContains) ? String(object.keywordsContains) : "",
      notesContains: isSet(object.notesContains) ? String(object.notesContains) : "",
    };
  },

  toJSON(message: QueryQCardsRequest): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.status !== undefined && (obj.status = queryQCardsRequest_StatusToJSON(message.status));
    message.cardType !== undefined && (obj.cardType = message.cardType);
    message.classes !== undefined && (obj.classes = message.classes);
    message.sortBy !== undefined && (obj.sortBy = message.sortBy);
    message.nameContains !== undefined && (obj.nameContains = message.nameContains);
    message.keywordsContains !== undefined && (obj.keywordsContains = message.keywordsContains);
    message.notesContains !== undefined && (obj.notesContains = message.notesContains);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryQCardsRequest>, I>>(object: I): QueryQCardsRequest {
    const message = createBaseQueryQCardsRequest();
    message.owner = object.owner ?? "";
    message.status = object.status ?? 0;
    message.cardType = object.cardType ?? "";
    message.classes = object.classes ?? "";
    message.sortBy = object.sortBy ?? "";
    message.nameContains = object.nameContains ?? "";
    message.keywordsContains = object.keywordsContains ?? "";
    message.notesContains = object.notesContains ?? "";
    return message;
  },
};

function createBaseQueryQCardsResponse(): QueryQCardsResponse {
  return { cardsList: [] };
}

export const QueryQCardsResponse = {
  encode(message: QueryQCardsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    writer.uint32(10).fork();
    for (const v of message.cardsList) {
      writer.uint64(v);
    }
    writer.ldelim();
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryQCardsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryQCardsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.cardsList.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.cardsList.push(longToNumber(reader.uint64() as Long));
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryQCardsResponse {
    return { cardsList: Array.isArray(object?.cardsList) ? object.cardsList.map((e: any) => Number(e)) : [] };
  },

  toJSON(message: QueryQCardsResponse): unknown {
    const obj: any = {};
    if (message.cardsList) {
      obj.cardsList = message.cardsList.map((e) => Math.round(e));
    } else {
      obj.cardsList = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryQCardsResponse>, I>>(object: I): QueryQCardsResponse {
    const message = createBaseQueryQCardsResponse();
    message.cardsList = object.cardsList?.map((e) => e) || [];
    return message;
  },
};

function createBaseQueryQMatchRequest(): QueryQMatchRequest {
  return { matchId: 0 };
}

export const QueryQMatchRequest = {
  encode(message: QueryQMatchRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.matchId !== 0) {
      writer.uint32(8).uint64(message.matchId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryQMatchRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryQMatchRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.matchId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryQMatchRequest {
    return { matchId: isSet(object.matchId) ? Number(object.matchId) : 0 };
  },

  toJSON(message: QueryQMatchRequest): unknown {
    const obj: any = {};
    message.matchId !== undefined && (obj.matchId = Math.round(message.matchId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryQMatchRequest>, I>>(object: I): QueryQMatchRequest {
    const message = createBaseQueryQMatchRequest();
    message.matchId = object.matchId ?? 0;
    return message;
  },
};

function createBaseQueryQCollectionRequest(): QueryQCollectionRequest {
  return { collectionId: 0 };
}

export const QueryQCollectionRequest = {
  encode(message: QueryQCollectionRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.collectionId !== 0) {
      writer.uint32(8).uint64(message.collectionId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryQCollectionRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryQCollectionRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.collectionId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryQCollectionRequest {
    return { collectionId: isSet(object.collectionId) ? Number(object.collectionId) : 0 };
  },

  toJSON(message: QueryQCollectionRequest): unknown {
    const obj: any = {};
    message.collectionId !== undefined && (obj.collectionId = Math.round(message.collectionId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryQCollectionRequest>, I>>(object: I): QueryQCollectionRequest {
    const message = createBaseQueryQCollectionRequest();
    message.collectionId = object.collectionId ?? 0;
    return message;
  },
};

function createBaseQueryQSellOfferRequest(): QueryQSellOfferRequest {
  return { sellOfferId: 0 };
}

export const QueryQSellOfferRequest = {
  encode(message: QueryQSellOfferRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.sellOfferId !== 0) {
      writer.uint32(8).uint64(message.sellOfferId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryQSellOfferRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryQSellOfferRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.sellOfferId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryQSellOfferRequest {
    return { sellOfferId: isSet(object.sellOfferId) ? Number(object.sellOfferId) : 0 };
  },

  toJSON(message: QueryQSellOfferRequest): unknown {
    const obj: any = {};
    message.sellOfferId !== undefined && (obj.sellOfferId = Math.round(message.sellOfferId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryQSellOfferRequest>, I>>(object: I): QueryQSellOfferRequest {
    const message = createBaseQueryQSellOfferRequest();
    message.sellOfferId = object.sellOfferId ?? 0;
    return message;
  },
};

function createBaseQueryQCouncilRequest(): QueryQCouncilRequest {
  return { councilId: 0 };
}

export const QueryQCouncilRequest = {
  encode(message: QueryQCouncilRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.councilId !== 0) {
      writer.uint32(8).uint64(message.councilId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryQCouncilRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryQCouncilRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.councilId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryQCouncilRequest {
    return { councilId: isSet(object.councilId) ? Number(object.councilId) : 0 };
  },

  toJSON(message: QueryQCouncilRequest): unknown {
    const obj: any = {};
    message.councilId !== undefined && (obj.councilId = Math.round(message.councilId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryQCouncilRequest>, I>>(object: I): QueryQCouncilRequest {
    const message = createBaseQueryQCouncilRequest();
    message.councilId = object.councilId ?? 0;
    return message;
  },
};

function createBaseQueryQMatchesRequest(): QueryQMatchesRequest {
  return {
    timestampDown: 0,
    timestampUp: 0,
    containsUsers: [],
    reporter: "",
    outcome: 0,
    cardsPlayed: [],
    ignore: undefined,
  };
}

export const QueryQMatchesRequest = {
  encode(message: QueryQMatchesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.timestampDown !== 0) {
      writer.uint32(8).uint64(message.timestampDown);
    }
    if (message.timestampUp !== 0) {
      writer.uint32(16).uint64(message.timestampUp);
    }
    for (const v of message.containsUsers) {
      writer.uint32(26).string(v!);
    }
    if (message.reporter !== "") {
      writer.uint32(34).string(message.reporter);
    }
    if (message.outcome !== 0) {
      writer.uint32(40).int32(message.outcome);
    }
    writer.uint32(50).fork();
    for (const v of message.cardsPlayed) {
      writer.uint64(v);
    }
    writer.ldelim();
    if (message.ignore !== undefined) {
      IgnoreMatches.encode(message.ignore, writer.uint32(58).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryQMatchesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryQMatchesRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.timestampDown = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.timestampUp = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.containsUsers.push(reader.string());
          break;
        case 4:
          message.reporter = reader.string();
          break;
        case 5:
          message.outcome = reader.int32() as any;
          break;
        case 6:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.cardsPlayed.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.cardsPlayed.push(longToNumber(reader.uint64() as Long));
          }
          break;
        case 7:
          message.ignore = IgnoreMatches.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryQMatchesRequest {
    return {
      timestampDown: isSet(object.timestampDown) ? Number(object.timestampDown) : 0,
      timestampUp: isSet(object.timestampUp) ? Number(object.timestampUp) : 0,
      containsUsers: Array.isArray(object?.containsUsers) ? object.containsUsers.map((e: any) => String(e)) : [],
      reporter: isSet(object.reporter) ? String(object.reporter) : "",
      outcome: isSet(object.outcome) ? outcomeFromJSON(object.outcome) : 0,
      cardsPlayed: Array.isArray(object?.cardsPlayed) ? object.cardsPlayed.map((e: any) => Number(e)) : [],
      ignore: isSet(object.ignore) ? IgnoreMatches.fromJSON(object.ignore) : undefined,
    };
  },

  toJSON(message: QueryQMatchesRequest): unknown {
    const obj: any = {};
    message.timestampDown !== undefined && (obj.timestampDown = Math.round(message.timestampDown));
    message.timestampUp !== undefined && (obj.timestampUp = Math.round(message.timestampUp));
    if (message.containsUsers) {
      obj.containsUsers = message.containsUsers.map((e) => e);
    } else {
      obj.containsUsers = [];
    }
    message.reporter !== undefined && (obj.reporter = message.reporter);
    message.outcome !== undefined && (obj.outcome = outcomeToJSON(message.outcome));
    if (message.cardsPlayed) {
      obj.cardsPlayed = message.cardsPlayed.map((e) => Math.round(e));
    } else {
      obj.cardsPlayed = [];
    }
    message.ignore !== undefined && (obj.ignore = message.ignore ? IgnoreMatches.toJSON(message.ignore) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryQMatchesRequest>, I>>(object: I): QueryQMatchesRequest {
    const message = createBaseQueryQMatchesRequest();
    message.timestampDown = object.timestampDown ?? 0;
    message.timestampUp = object.timestampUp ?? 0;
    message.containsUsers = object.containsUsers?.map((e) => e) || [];
    message.reporter = object.reporter ?? "";
    message.outcome = object.outcome ?? 0;
    message.cardsPlayed = object.cardsPlayed?.map((e) => e) || [];
    message.ignore = (object.ignore !== undefined && object.ignore !== null)
      ? IgnoreMatches.fromPartial(object.ignore)
      : undefined;
    return message;
  },
};

function createBaseIgnoreMatches(): IgnoreMatches {
  return { outcome: false };
}

export const IgnoreMatches = {
  encode(message: IgnoreMatches, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.outcome === true) {
      writer.uint32(8).bool(message.outcome);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): IgnoreMatches {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIgnoreMatches();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.outcome = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): IgnoreMatches {
    return { outcome: isSet(object.outcome) ? Boolean(object.outcome) : false };
  },

  toJSON(message: IgnoreMatches): unknown {
    const obj: any = {};
    message.outcome !== undefined && (obj.outcome = message.outcome);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<IgnoreMatches>, I>>(object: I): IgnoreMatches {
    const message = createBaseIgnoreMatches();
    message.outcome = object.outcome ?? false;
    return message;
  },
};

function createBaseQueryQMatchesResponse(): QueryQMatchesResponse {
  return { matchesList: [], matches: [] };
}

export const QueryQMatchesResponse = {
  encode(message: QueryQMatchesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    writer.uint32(10).fork();
    for (const v of message.matchesList) {
      writer.uint64(v);
    }
    writer.ldelim();
    for (const v of message.matches) {
      Match.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryQMatchesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryQMatchesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.matchesList.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.matchesList.push(longToNumber(reader.uint64() as Long));
          }
          break;
        case 2:
          message.matches.push(Match.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryQMatchesResponse {
    return {
      matchesList: Array.isArray(object?.matchesList) ? object.matchesList.map((e: any) => Number(e)) : [],
      matches: Array.isArray(object?.matches) ? object.matches.map((e: any) => Match.fromJSON(e)) : [],
    };
  },

  toJSON(message: QueryQMatchesResponse): unknown {
    const obj: any = {};
    if (message.matchesList) {
      obj.matchesList = message.matchesList.map((e) => Math.round(e));
    } else {
      obj.matchesList = [];
    }
    if (message.matches) {
      obj.matches = message.matches.map((e) => e ? Match.toJSON(e) : undefined);
    } else {
      obj.matches = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryQMatchesResponse>, I>>(object: I): QueryQMatchesResponse {
    const message = createBaseQueryQMatchesResponse();
    message.matchesList = object.matchesList?.map((e) => e) || [];
    message.matches = object.matches?.map((e) => Match.fromPartial(e)) || [];
    return message;
  },
};

function createBaseQueryQSellOffersRequest(): QueryQSellOffersRequest {
  return { priceDown: "", priceUp: "", seller: "", buyer: "", card: 0, status: 0, ignore: undefined };
}

export const QueryQSellOffersRequest = {
  encode(message: QueryQSellOffersRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.priceDown !== "") {
      writer.uint32(10).string(message.priceDown);
    }
    if (message.priceUp !== "") {
      writer.uint32(18).string(message.priceUp);
    }
    if (message.seller !== "") {
      writer.uint32(26).string(message.seller);
    }
    if (message.buyer !== "") {
      writer.uint32(34).string(message.buyer);
    }
    if (message.card !== 0) {
      writer.uint32(40).uint64(message.card);
    }
    if (message.status !== 0) {
      writer.uint32(48).int32(message.status);
    }
    if (message.ignore !== undefined) {
      IgnoreSellOffers.encode(message.ignore, writer.uint32(58).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryQSellOffersRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryQSellOffersRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.priceDown = reader.string();
          break;
        case 2:
          message.priceUp = reader.string();
          break;
        case 3:
          message.seller = reader.string();
          break;
        case 4:
          message.buyer = reader.string();
          break;
        case 5:
          message.card = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.status = reader.int32() as any;
          break;
        case 7:
          message.ignore = IgnoreSellOffers.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryQSellOffersRequest {
    return {
      priceDown: isSet(object.priceDown) ? String(object.priceDown) : "",
      priceUp: isSet(object.priceUp) ? String(object.priceUp) : "",
      seller: isSet(object.seller) ? String(object.seller) : "",
      buyer: isSet(object.buyer) ? String(object.buyer) : "",
      card: isSet(object.card) ? Number(object.card) : 0,
      status: isSet(object.status) ? sellOfferStatusFromJSON(object.status) : 0,
      ignore: isSet(object.ignore) ? IgnoreSellOffers.fromJSON(object.ignore) : undefined,
    };
  },

  toJSON(message: QueryQSellOffersRequest): unknown {
    const obj: any = {};
    message.priceDown !== undefined && (obj.priceDown = message.priceDown);
    message.priceUp !== undefined && (obj.priceUp = message.priceUp);
    message.seller !== undefined && (obj.seller = message.seller);
    message.buyer !== undefined && (obj.buyer = message.buyer);
    message.card !== undefined && (obj.card = Math.round(message.card));
    message.status !== undefined && (obj.status = sellOfferStatusToJSON(message.status));
    message.ignore !== undefined && (obj.ignore = message.ignore ? IgnoreSellOffers.toJSON(message.ignore) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryQSellOffersRequest>, I>>(object: I): QueryQSellOffersRequest {
    const message = createBaseQueryQSellOffersRequest();
    message.priceDown = object.priceDown ?? "";
    message.priceUp = object.priceUp ?? "";
    message.seller = object.seller ?? "";
    message.buyer = object.buyer ?? "";
    message.card = object.card ?? 0;
    message.status = object.status ?? 0;
    message.ignore = (object.ignore !== undefined && object.ignore !== null)
      ? IgnoreSellOffers.fromPartial(object.ignore)
      : undefined;
    return message;
  },
};

function createBaseIgnoreSellOffers(): IgnoreSellOffers {
  return { status: false, card: false };
}

export const IgnoreSellOffers = {
  encode(message: IgnoreSellOffers, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.status === true) {
      writer.uint32(8).bool(message.status);
    }
    if (message.card === true) {
      writer.uint32(16).bool(message.card);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): IgnoreSellOffers {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIgnoreSellOffers();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.status = reader.bool();
          break;
        case 2:
          message.card = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): IgnoreSellOffers {
    return {
      status: isSet(object.status) ? Boolean(object.status) : false,
      card: isSet(object.card) ? Boolean(object.card) : false,
    };
  },

  toJSON(message: IgnoreSellOffers): unknown {
    const obj: any = {};
    message.status !== undefined && (obj.status = message.status);
    message.card !== undefined && (obj.card = message.card);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<IgnoreSellOffers>, I>>(object: I): IgnoreSellOffers {
    const message = createBaseIgnoreSellOffers();
    message.status = object.status ?? false;
    message.card = object.card ?? false;
    return message;
  },
};

function createBaseQueryQSellOffersResponse(): QueryQSellOffersResponse {
  return { sellOffersIds: [], sellOffers: [] };
}

export const QueryQSellOffersResponse = {
  encode(message: QueryQSellOffersResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    writer.uint32(10).fork();
    for (const v of message.sellOffersIds) {
      writer.uint64(v);
    }
    writer.ldelim();
    for (const v of message.sellOffers) {
      SellOffer.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryQSellOffersResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryQSellOffersResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.sellOffersIds.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.sellOffersIds.push(longToNumber(reader.uint64() as Long));
          }
          break;
        case 2:
          message.sellOffers.push(SellOffer.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryQSellOffersResponse {
    return {
      sellOffersIds: Array.isArray(object?.sellOffersIds) ? object.sellOffersIds.map((e: any) => Number(e)) : [],
      sellOffers: Array.isArray(object?.sellOffers) ? object.sellOffers.map((e: any) => SellOffer.fromJSON(e)) : [],
    };
  },

  toJSON(message: QueryQSellOffersResponse): unknown {
    const obj: any = {};
    if (message.sellOffersIds) {
      obj.sellOffersIds = message.sellOffersIds.map((e) => Math.round(e));
    } else {
      obj.sellOffersIds = [];
    }
    if (message.sellOffers) {
      obj.sellOffers = message.sellOffers.map((e) => e ? SellOffer.toJSON(e) : undefined);
    } else {
      obj.sellOffers = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryQSellOffersResponse>, I>>(object: I): QueryQSellOffersResponse {
    const message = createBaseQueryQSellOffersResponse();
    message.sellOffersIds = object.sellOffersIds?.map((e) => e) || [];
    message.sellOffers = object.sellOffers?.map((e) => SellOffer.fromPartial(e)) || [];
    return message;
  },
};

function createBaseQueryQServerRequest(): QueryQServerRequest {
  return { id: 0 };
}

export const QueryQServerRequest = {
  encode(message: QueryQServerRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryQServerRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryQServerRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryQServerRequest {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: QueryQServerRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryQServerRequest>, I>>(object: I): QueryQServerRequest {
    const message = createBaseQueryQServerRequest();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseQueryQServerResponse(): QueryQServerResponse {
  return {};
}

export const QueryQServerResponse = {
  encode(_: QueryQServerResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryQServerResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryQServerResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryQServerResponse {
    return {};
  },

  toJSON(_: QueryQServerResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryQServerResponse>, I>>(_: I): QueryQServerResponse {
    const message = createBaseQueryQServerResponse();
    return message;
  },
};

function createBaseQueryQCollectionsRequest(): QueryQCollectionsRequest {
  return { status: 0, ignoreStatus: false, contributors: [], containsCards: [], owner: "" };
}

export const QueryQCollectionsRequest = {
  encode(message: QueryQCollectionsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.status !== 0) {
      writer.uint32(8).int32(message.status);
    }
    if (message.ignoreStatus === true) {
      writer.uint32(16).bool(message.ignoreStatus);
    }
    for (const v of message.contributors) {
      writer.uint32(26).string(v!);
    }
    writer.uint32(34).fork();
    for (const v of message.containsCards) {
      writer.uint64(v);
    }
    writer.ldelim();
    if (message.owner !== "") {
      writer.uint32(42).string(message.owner);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryQCollectionsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryQCollectionsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.status = reader.int32() as any;
          break;
        case 2:
          message.ignoreStatus = reader.bool();
          break;
        case 3:
          message.contributors.push(reader.string());
          break;
        case 4:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.containsCards.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.containsCards.push(longToNumber(reader.uint64() as Long));
          }
          break;
        case 5:
          message.owner = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryQCollectionsRequest {
    return {
      status: isSet(object.status) ? cStatusFromJSON(object.status) : 0,
      ignoreStatus: isSet(object.ignoreStatus) ? Boolean(object.ignoreStatus) : false,
      contributors: Array.isArray(object?.contributors) ? object.contributors.map((e: any) => String(e)) : [],
      containsCards: Array.isArray(object?.containsCards) ? object.containsCards.map((e: any) => Number(e)) : [],
      owner: isSet(object.owner) ? String(object.owner) : "",
    };
  },

  toJSON(message: QueryQCollectionsRequest): unknown {
    const obj: any = {};
    message.status !== undefined && (obj.status = cStatusToJSON(message.status));
    message.ignoreStatus !== undefined && (obj.ignoreStatus = message.ignoreStatus);
    if (message.contributors) {
      obj.contributors = message.contributors.map((e) => e);
    } else {
      obj.contributors = [];
    }
    if (message.containsCards) {
      obj.containsCards = message.containsCards.map((e) => Math.round(e));
    } else {
      obj.containsCards = [];
    }
    message.owner !== undefined && (obj.owner = message.owner);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryQCollectionsRequest>, I>>(object: I): QueryQCollectionsRequest {
    const message = createBaseQueryQCollectionsRequest();
    message.status = object.status ?? 0;
    message.ignoreStatus = object.ignoreStatus ?? false;
    message.contributors = object.contributors?.map((e) => e) || [];
    message.containsCards = object.containsCards?.map((e) => e) || [];
    message.owner = object.owner ?? "";
    return message;
  },
};

function createBaseQueryQCollectionsResponse(): QueryQCollectionsResponse {
  return { collectionIds: [] };
}

export const QueryQCollectionsResponse = {
  encode(message: QueryQCollectionsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    writer.uint32(10).fork();
    for (const v of message.collectionIds) {
      writer.uint64(v);
    }
    writer.ldelim();
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryQCollectionsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryQCollectionsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.collectionIds.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.collectionIds.push(longToNumber(reader.uint64() as Long));
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryQCollectionsResponse {
    return {
      collectionIds: Array.isArray(object?.collectionIds) ? object.collectionIds.map((e: any) => Number(e)) : [],
    };
  },

  toJSON(message: QueryQCollectionsResponse): unknown {
    const obj: any = {};
    if (message.collectionIds) {
      obj.collectionIds = message.collectionIds.map((e) => Math.round(e));
    } else {
      obj.collectionIds = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryQCollectionsResponse>, I>>(object: I): QueryQCollectionsResponse {
    const message = createBaseQueryQCollectionsResponse();
    message.collectionIds = object.collectionIds?.map((e) => e) || [];
    return message;
  },
};

function createBaseQueryRarityDistributionRequest(): QueryRarityDistributionRequest {
  return { collectionId: 0 };
}

export const QueryRarityDistributionRequest = {
  encode(message: QueryRarityDistributionRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.collectionId !== 0) {
      writer.uint32(8).uint64(message.collectionId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryRarityDistributionRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryRarityDistributionRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.collectionId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryRarityDistributionRequest {
    return { collectionId: isSet(object.collectionId) ? Number(object.collectionId) : 0 };
  },

  toJSON(message: QueryRarityDistributionRequest): unknown {
    const obj: any = {};
    message.collectionId !== undefined && (obj.collectionId = Math.round(message.collectionId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryRarityDistributionRequest>, I>>(
    object: I,
  ): QueryRarityDistributionRequest {
    const message = createBaseQueryRarityDistributionRequest();
    message.collectionId = object.collectionId ?? 0;
    return message;
  },
};

function createBaseQueryRarityDistributionResponse(): QueryRarityDistributionResponse {
  return { current: [], wanted: [] };
}

export const QueryRarityDistributionResponse = {
  encode(message: QueryRarityDistributionResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    writer.uint32(10).fork();
    for (const v of message.current) {
      writer.uint64(v);
    }
    writer.ldelim();
    writer.uint32(18).fork();
    for (const v of message.wanted) {
      writer.uint64(v);
    }
    writer.ldelim();
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryRarityDistributionResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryRarityDistributionResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.current.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.current.push(longToNumber(reader.uint64() as Long));
          }
          break;
        case 2:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.wanted.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.wanted.push(longToNumber(reader.uint64() as Long));
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryRarityDistributionResponse {
    return {
      current: Array.isArray(object?.current) ? object.current.map((e: any) => Number(e)) : [],
      wanted: Array.isArray(object?.wanted) ? object.wanted.map((e: any) => Number(e)) : [],
    };
  },

  toJSON(message: QueryRarityDistributionResponse): unknown {
    const obj: any = {};
    if (message.current) {
      obj.current = message.current.map((e) => Math.round(e));
    } else {
      obj.current = [];
    }
    if (message.wanted) {
      obj.wanted = message.wanted.map((e) => Math.round(e));
    } else {
      obj.wanted = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryRarityDistributionResponse>, I>>(
    object: I,
  ): QueryRarityDistributionResponse {
    const message = createBaseQueryRarityDistributionResponse();
    message.current = object.current?.map((e) => e) || [];
    message.wanted = object.wanted?.map((e) => e) || [];
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of QCard items. */
  QCard(request: QueryQCardRequest): Promise<OutpCard>;
  /** Queries a list of QCardContent items. */
  QCardContent(request: QueryQCardContentRequest): Promise<QueryQCardContentResponse>;
  /** Queries a list of QUser items. */
  QUser(request: QueryQUserRequest): Promise<User>;
  /** Queries a list of QCardchainInfo items. */
  QCardchainInfo(request: QueryQCardchainInfoRequest): Promise<QueryQCardchainInfoResponse>;
  /** Queries a list of QVotingResults items. */
  QVotingResults(request: QueryQVotingResultsRequest): Promise<QueryQVotingResultsResponse>;
  /** Queries a list of QVotableCards items. */
  QVotableCards(request: QueryQVotableCardsRequest): Promise<QueryQVotableCardsResponse>;
  /** Queries a list of QCards items. */
  QCards(request: QueryQCardsRequest): Promise<QueryQCardsResponse>;
  /** Queries a list of QMatch items. */
  QMatch(request: QueryQMatchRequest): Promise<Match>;
  /** Queries a list of QCollection items. */
  QCollection(request: QueryQCollectionRequest): Promise<OutpCollection>;
  /** Queries a list of QSellOffer items. */
  QSellOffer(request: QueryQSellOfferRequest): Promise<SellOffer>;
  /** Queries a list of QCouncil items. */
  QCouncil(request: QueryQCouncilRequest): Promise<Council>;
  /** Queries a list of QMatches items. */
  QMatches(request: QueryQMatchesRequest): Promise<QueryQMatchesResponse>;
  /** Queries a list of QSellOffers items. */
  QSellOffers(request: QueryQSellOffersRequest): Promise<QueryQSellOffersResponse>;
  /** Queries a list of QServer items. */
  QServer(request: QueryQServerRequest): Promise<Server>;
  /** Queries a list of QCollections items. */
  QCollections(request: QueryQCollectionsRequest): Promise<QueryQCollectionsResponse>;
  /** Queries a list of RarityDistribution items. */
  RarityDistribution(request: QueryRarityDistributionRequest): Promise<QueryRarityDistributionResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.QCard = this.QCard.bind(this);
    this.QCardContent = this.QCardContent.bind(this);
    this.QUser = this.QUser.bind(this);
    this.QCardchainInfo = this.QCardchainInfo.bind(this);
    this.QVotingResults = this.QVotingResults.bind(this);
    this.QVotableCards = this.QVotableCards.bind(this);
    this.QCards = this.QCards.bind(this);
    this.QMatch = this.QMatch.bind(this);
    this.QCollection = this.QCollection.bind(this);
    this.QSellOffer = this.QSellOffer.bind(this);
    this.QCouncil = this.QCouncil.bind(this);
    this.QMatches = this.QMatches.bind(this);
    this.QSellOffers = this.QSellOffers.bind(this);
    this.QServer = this.QServer.bind(this);
    this.QCollections = this.QCollections.bind(this);
    this.RarityDistribution = this.RarityDistribution.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  QCard(request: QueryQCardRequest): Promise<OutpCard> {
    const data = QueryQCardRequest.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QCard", data);
    return promise.then((data) => OutpCard.decode(new _m0.Reader(data)));
  }

  QCardContent(request: QueryQCardContentRequest): Promise<QueryQCardContentResponse> {
    const data = QueryQCardContentRequest.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QCardContent", data);
    return promise.then((data) => QueryQCardContentResponse.decode(new _m0.Reader(data)));
  }

  QUser(request: QueryQUserRequest): Promise<User> {
    const data = QueryQUserRequest.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QUser", data);
    return promise.then((data) => User.decode(new _m0.Reader(data)));
  }

  QCardchainInfo(request: QueryQCardchainInfoRequest): Promise<QueryQCardchainInfoResponse> {
    const data = QueryQCardchainInfoRequest.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QCardchainInfo", data);
    return promise.then((data) => QueryQCardchainInfoResponse.decode(new _m0.Reader(data)));
  }

  QVotingResults(request: QueryQVotingResultsRequest): Promise<QueryQVotingResultsResponse> {
    const data = QueryQVotingResultsRequest.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QVotingResults", data);
    return promise.then((data) => QueryQVotingResultsResponse.decode(new _m0.Reader(data)));
  }

  QVotableCards(request: QueryQVotableCardsRequest): Promise<QueryQVotableCardsResponse> {
    const data = QueryQVotableCardsRequest.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QVotableCards", data);
    return promise.then((data) => QueryQVotableCardsResponse.decode(new _m0.Reader(data)));
  }

  QCards(request: QueryQCardsRequest): Promise<QueryQCardsResponse> {
    const data = QueryQCardsRequest.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QCards", data);
    return promise.then((data) => QueryQCardsResponse.decode(new _m0.Reader(data)));
  }

  QMatch(request: QueryQMatchRequest): Promise<Match> {
    const data = QueryQMatchRequest.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QMatch", data);
    return promise.then((data) => Match.decode(new _m0.Reader(data)));
  }

  QCollection(request: QueryQCollectionRequest): Promise<OutpCollection> {
    const data = QueryQCollectionRequest.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QCollection", data);
    return promise.then((data) => OutpCollection.decode(new _m0.Reader(data)));
  }

  QSellOffer(request: QueryQSellOfferRequest): Promise<SellOffer> {
    const data = QueryQSellOfferRequest.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QSellOffer", data);
    return promise.then((data) => SellOffer.decode(new _m0.Reader(data)));
  }

  QCouncil(request: QueryQCouncilRequest): Promise<Council> {
    const data = QueryQCouncilRequest.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QCouncil", data);
    return promise.then((data) => Council.decode(new _m0.Reader(data)));
  }

  QMatches(request: QueryQMatchesRequest): Promise<QueryQMatchesResponse> {
    const data = QueryQMatchesRequest.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QMatches", data);
    return promise.then((data) => QueryQMatchesResponse.decode(new _m0.Reader(data)));
  }

  QSellOffers(request: QueryQSellOffersRequest): Promise<QueryQSellOffersResponse> {
    const data = QueryQSellOffersRequest.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QSellOffers", data);
    return promise.then((data) => QueryQSellOffersResponse.decode(new _m0.Reader(data)));
  }

  QServer(request: QueryQServerRequest): Promise<Server> {
    const data = QueryQServerRequest.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QServer", data);
    return promise.then((data) => Server.decode(new _m0.Reader(data)));
  }

  QCollections(request: QueryQCollectionsRequest): Promise<QueryQCollectionsResponse> {
    const data = QueryQCollectionsRequest.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QCollections", data);
    return promise.then((data) => QueryQCollectionsResponse.decode(new _m0.Reader(data)));
  }

  RarityDistribution(request: QueryRarityDistributionRequest): Promise<QueryRarityDistributionResponse> {
    const data = QueryRarityDistributionRequest.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "RarityDistribution", data);
    return promise.then((data) => QueryRarityDistributionResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
