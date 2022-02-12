/* eslint-disable */
import { Status, statusFromJSON, statusToJSON } from "../cardchain/card";
import {
  CouncilStatus,
  councilStatusFromJSON,
  councilStatusToJSON,
} from "../cardchain/user";
import { Outcome, outcomeFromJSON, outcomeToJSON } from "../cardchain/tx";
import {
  CStatus,
  cStatusFromJSON,
  cStatusToJSON,
} from "../cardchain/collection";
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Params } from "../cardchain/params";
import { VoteRight } from "../cardchain/vote_right";
import { VotingResults } from "../cardchain/voting_results";

export const protobufPackage = "DecentralCardGame.cardchain.cardchain";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

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

export interface QueryQCardchainInfoRequest {}

export interface QueryQCardchainInfoResponse {
  cardAuctionPrice: string;
  activeCollections: number[];
}

export interface QueryQVotingResultsRequest {}

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

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
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
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
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
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryQCardRequest: object = { cardId: "" };

export const QueryQCardRequest = {
  encode(message: QueryQCardRequest, writer: Writer = Writer.create()): Writer {
    if (message.cardId !== "") {
      writer.uint32(10).string(message.cardId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryQCardRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryQCardRequest } as QueryQCardRequest;
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
    const message = { ...baseQueryQCardRequest } as QueryQCardRequest;
    if (object.cardId !== undefined && object.cardId !== null) {
      message.cardId = String(object.cardId);
    } else {
      message.cardId = "";
    }
    return message;
  },

  toJSON(message: QueryQCardRequest): unknown {
    const obj: any = {};
    message.cardId !== undefined && (obj.cardId = message.cardId);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryQCardRequest>): QueryQCardRequest {
    const message = { ...baseQueryQCardRequest } as QueryQCardRequest;
    if (object.cardId !== undefined && object.cardId !== null) {
      message.cardId = object.cardId;
    } else {
      message.cardId = "";
    }
    return message;
  },
};

const baseQueryQCardResponse: object = {
  owner: "",
  content: "",
  artist: "",
  image: "",
  fullArt: false,
  notes: "",
  status: 0,
  votePool: "",
  fairEnoughVotes: 0,
  overpoweredVotes: 0,
  underpoweredVotes: 0,
  inappropriateVotes: 0,
  nerflevel: 0,
};

export const QueryQCardResponse = {
  encode(
    message: QueryQCardResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.content !== "") {
      writer.uint32(18).string(message.content);
    }
    if (message.artist !== "") {
      writer.uint32(26).string(message.artist);
    }
    if (message.image !== "") {
      writer.uint32(34).string(message.image);
    }
    if (message.fullArt === true) {
      writer.uint32(40).bool(message.fullArt);
    }
    if (message.notes !== "") {
      writer.uint32(50).string(message.notes);
    }
    if (message.status !== 0) {
      writer.uint32(56).int32(message.status);
    }
    if (message.votePool !== "") {
      writer.uint32(66).string(message.votePool);
    }
    if (message.fairEnoughVotes !== 0) {
      writer.uint32(72).uint64(message.fairEnoughVotes);
    }
    if (message.overpoweredVotes !== 0) {
      writer.uint32(80).uint64(message.overpoweredVotes);
    }
    if (message.underpoweredVotes !== 0) {
      writer.uint32(88).uint64(message.underpoweredVotes);
    }
    if (message.inappropriateVotes !== 0) {
      writer.uint32(96).uint64(message.inappropriateVotes);
    }
    if (message.nerflevel !== 0) {
      writer.uint32(104).int64(message.nerflevel);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryQCardResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryQCardResponse } as QueryQCardResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.content = reader.string();
          break;
        case 3:
          message.artist = reader.string();
          break;
        case 4:
          message.image = reader.string();
          break;
        case 5:
          message.fullArt = reader.bool();
          break;
        case 6:
          message.notes = reader.string();
          break;
        case 7:
          message.status = reader.int32() as any;
          break;
        case 8:
          message.votePool = reader.string();
          break;
        case 9:
          message.fairEnoughVotes = longToNumber(reader.uint64() as Long);
          break;
        case 10:
          message.overpoweredVotes = longToNumber(reader.uint64() as Long);
          break;
        case 11:
          message.underpoweredVotes = longToNumber(reader.uint64() as Long);
          break;
        case 12:
          message.inappropriateVotes = longToNumber(reader.uint64() as Long);
          break;
        case 13:
          message.nerflevel = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryQCardResponse {
    const message = { ...baseQueryQCardResponse } as QueryQCardResponse;
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.content !== undefined && object.content !== null) {
      message.content = String(object.content);
    } else {
      message.content = "";
    }
    if (object.artist !== undefined && object.artist !== null) {
      message.artist = String(object.artist);
    } else {
      message.artist = "";
    }
    if (object.image !== undefined && object.image !== null) {
      message.image = String(object.image);
    } else {
      message.image = "";
    }
    if (object.fullArt !== undefined && object.fullArt !== null) {
      message.fullArt = Boolean(object.fullArt);
    } else {
      message.fullArt = false;
    }
    if (object.notes !== undefined && object.notes !== null) {
      message.notes = String(object.notes);
    } else {
      message.notes = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = statusFromJSON(object.status);
    } else {
      message.status = 0;
    }
    if (object.votePool !== undefined && object.votePool !== null) {
      message.votePool = String(object.votePool);
    } else {
      message.votePool = "";
    }
    if (
      object.fairEnoughVotes !== undefined &&
      object.fairEnoughVotes !== null
    ) {
      message.fairEnoughVotes = Number(object.fairEnoughVotes);
    } else {
      message.fairEnoughVotes = 0;
    }
    if (
      object.overpoweredVotes !== undefined &&
      object.overpoweredVotes !== null
    ) {
      message.overpoweredVotes = Number(object.overpoweredVotes);
    } else {
      message.overpoweredVotes = 0;
    }
    if (
      object.underpoweredVotes !== undefined &&
      object.underpoweredVotes !== null
    ) {
      message.underpoweredVotes = Number(object.underpoweredVotes);
    } else {
      message.underpoweredVotes = 0;
    }
    if (
      object.inappropriateVotes !== undefined &&
      object.inappropriateVotes !== null
    ) {
      message.inappropriateVotes = Number(object.inappropriateVotes);
    } else {
      message.inappropriateVotes = 0;
    }
    if (object.nerflevel !== undefined && object.nerflevel !== null) {
      message.nerflevel = Number(object.nerflevel);
    } else {
      message.nerflevel = 0;
    }
    return message;
  },

  toJSON(message: QueryQCardResponse): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.content !== undefined && (obj.content = message.content);
    message.artist !== undefined && (obj.artist = message.artist);
    message.image !== undefined && (obj.image = message.image);
    message.fullArt !== undefined && (obj.fullArt = message.fullArt);
    message.notes !== undefined && (obj.notes = message.notes);
    message.status !== undefined && (obj.status = statusToJSON(message.status));
    message.votePool !== undefined && (obj.votePool = message.votePool);
    message.fairEnoughVotes !== undefined &&
      (obj.fairEnoughVotes = message.fairEnoughVotes);
    message.overpoweredVotes !== undefined &&
      (obj.overpoweredVotes = message.overpoweredVotes);
    message.underpoweredVotes !== undefined &&
      (obj.underpoweredVotes = message.underpoweredVotes);
    message.inappropriateVotes !== undefined &&
      (obj.inappropriateVotes = message.inappropriateVotes);
    message.nerflevel !== undefined && (obj.nerflevel = message.nerflevel);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryQCardResponse>): QueryQCardResponse {
    const message = { ...baseQueryQCardResponse } as QueryQCardResponse;
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.content !== undefined && object.content !== null) {
      message.content = object.content;
    } else {
      message.content = "";
    }
    if (object.artist !== undefined && object.artist !== null) {
      message.artist = object.artist;
    } else {
      message.artist = "";
    }
    if (object.image !== undefined && object.image !== null) {
      message.image = object.image;
    } else {
      message.image = "";
    }
    if (object.fullArt !== undefined && object.fullArt !== null) {
      message.fullArt = object.fullArt;
    } else {
      message.fullArt = false;
    }
    if (object.notes !== undefined && object.notes !== null) {
      message.notes = object.notes;
    } else {
      message.notes = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = 0;
    }
    if (object.votePool !== undefined && object.votePool !== null) {
      message.votePool = object.votePool;
    } else {
      message.votePool = "";
    }
    if (
      object.fairEnoughVotes !== undefined &&
      object.fairEnoughVotes !== null
    ) {
      message.fairEnoughVotes = object.fairEnoughVotes;
    } else {
      message.fairEnoughVotes = 0;
    }
    if (
      object.overpoweredVotes !== undefined &&
      object.overpoweredVotes !== null
    ) {
      message.overpoweredVotes = object.overpoweredVotes;
    } else {
      message.overpoweredVotes = 0;
    }
    if (
      object.underpoweredVotes !== undefined &&
      object.underpoweredVotes !== null
    ) {
      message.underpoweredVotes = object.underpoweredVotes;
    } else {
      message.underpoweredVotes = 0;
    }
    if (
      object.inappropriateVotes !== undefined &&
      object.inappropriateVotes !== null
    ) {
      message.inappropriateVotes = object.inappropriateVotes;
    } else {
      message.inappropriateVotes = 0;
    }
    if (object.nerflevel !== undefined && object.nerflevel !== null) {
      message.nerflevel = object.nerflevel;
    } else {
      message.nerflevel = 0;
    }
    return message;
  },
};

const baseQueryQCardContentRequest: object = { cardId: "" };

export const QueryQCardContentRequest = {
  encode(
    message: QueryQCardContentRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.cardId !== "") {
      writer.uint32(10).string(message.cardId);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryQCardContentRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryQCardContentRequest,
    } as QueryQCardContentRequest;
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
    const message = {
      ...baseQueryQCardContentRequest,
    } as QueryQCardContentRequest;
    if (object.cardId !== undefined && object.cardId !== null) {
      message.cardId = String(object.cardId);
    } else {
      message.cardId = "";
    }
    return message;
  },

  toJSON(message: QueryQCardContentRequest): unknown {
    const obj: any = {};
    message.cardId !== undefined && (obj.cardId = message.cardId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryQCardContentRequest>
  ): QueryQCardContentRequest {
    const message = {
      ...baseQueryQCardContentRequest,
    } as QueryQCardContentRequest;
    if (object.cardId !== undefined && object.cardId !== null) {
      message.cardId = object.cardId;
    } else {
      message.cardId = "";
    }
    return message;
  },
};

const baseQueryQCardContentResponse: object = { content: "" };

export const QueryQCardContentResponse = {
  encode(
    message: QueryQCardContentResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.content !== "") {
      writer.uint32(10).string(message.content);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryQCardContentResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryQCardContentResponse,
    } as QueryQCardContentResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.content = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryQCardContentResponse {
    const message = {
      ...baseQueryQCardContentResponse,
    } as QueryQCardContentResponse;
    if (object.content !== undefined && object.content !== null) {
      message.content = String(object.content);
    } else {
      message.content = "";
    }
    return message;
  },

  toJSON(message: QueryQCardContentResponse): unknown {
    const obj: any = {};
    message.content !== undefined && (obj.content = message.content);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryQCardContentResponse>
  ): QueryQCardContentResponse {
    const message = {
      ...baseQueryQCardContentResponse,
    } as QueryQCardContentResponse;
    if (object.content !== undefined && object.content !== null) {
      message.content = object.content;
    } else {
      message.content = "";
    }
    return message;
  },
};

const baseQueryQUserRequest: object = { address: "" };

export const QueryQUserRequest = {
  encode(message: QueryQUserRequest, writer: Writer = Writer.create()): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryQUserRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryQUserRequest } as QueryQUserRequest;
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
    const message = { ...baseQueryQUserRequest } as QueryQUserRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: QueryQUserRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryQUserRequest>): QueryQUserRequest {
    const message = { ...baseQueryQUserRequest } as QueryQUserRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseQueryQUserResponse: object = {
  alias: "",
  ownedCardSchemes: 0,
  ownedPrototypes: 0,
  cards: 0,
  councilStatus: 0,
  reportMatches: false,
};

export const QueryQUserResponse = {
  encode(
    message: QueryQUserResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.alias !== "") {
      writer.uint32(10).string(message.alias);
    }
    writer.uint32(18).fork();
    for (const v of message.ownedCardSchemes) {
      writer.uint64(v);
    }
    writer.ldelim();
    writer.uint32(26).fork();
    for (const v of message.ownedPrototypes) {
      writer.uint64(v);
    }
    writer.ldelim();
    writer.uint32(34).fork();
    for (const v of message.cards) {
      writer.uint64(v);
    }
    writer.ldelim();
    for (const v of message.voteRights) {
      VoteRight.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    if (message.councilStatus !== 0) {
      writer.uint32(48).int32(message.councilStatus);
    }
    if (message.reportMatches === true) {
      writer.uint32(56).bool(message.reportMatches);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryQUserResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryQUserResponse } as QueryQUserResponse;
    message.ownedCardSchemes = [];
    message.ownedPrototypes = [];
    message.cards = [];
    message.voteRights = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.alias = reader.string();
          break;
        case 2:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.ownedCardSchemes.push(
                longToNumber(reader.uint64() as Long)
              );
            }
          } else {
            message.ownedCardSchemes.push(
              longToNumber(reader.uint64() as Long)
            );
          }
          break;
        case 3:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.ownedPrototypes.push(
                longToNumber(reader.uint64() as Long)
              );
            }
          } else {
            message.ownedPrototypes.push(longToNumber(reader.uint64() as Long));
          }
          break;
        case 4:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.cards.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.cards.push(longToNumber(reader.uint64() as Long));
          }
          break;
        case 5:
          message.voteRights.push(VoteRight.decode(reader, reader.uint32()));
          break;
        case 6:
          message.councilStatus = reader.int32() as any;
          break;
        case 7:
          message.reportMatches = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryQUserResponse {
    const message = { ...baseQueryQUserResponse } as QueryQUserResponse;
    message.ownedCardSchemes = [];
    message.ownedPrototypes = [];
    message.cards = [];
    message.voteRights = [];
    if (object.alias !== undefined && object.alias !== null) {
      message.alias = String(object.alias);
    } else {
      message.alias = "";
    }
    if (
      object.ownedCardSchemes !== undefined &&
      object.ownedCardSchemes !== null
    ) {
      for (const e of object.ownedCardSchemes) {
        message.ownedCardSchemes.push(Number(e));
      }
    }
    if (
      object.ownedPrototypes !== undefined &&
      object.ownedPrototypes !== null
    ) {
      for (const e of object.ownedPrototypes) {
        message.ownedPrototypes.push(Number(e));
      }
    }
    if (object.cards !== undefined && object.cards !== null) {
      for (const e of object.cards) {
        message.cards.push(Number(e));
      }
    }
    if (object.voteRights !== undefined && object.voteRights !== null) {
      for (const e of object.voteRights) {
        message.voteRights.push(VoteRight.fromJSON(e));
      }
    }
    if (object.councilStatus !== undefined && object.councilStatus !== null) {
      message.councilStatus = councilStatusFromJSON(object.councilStatus);
    } else {
      message.councilStatus = 0;
    }
    if (object.reportMatches !== undefined && object.reportMatches !== null) {
      message.reportMatches = Boolean(object.reportMatches);
    } else {
      message.reportMatches = false;
    }
    return message;
  },

  toJSON(message: QueryQUserResponse): unknown {
    const obj: any = {};
    message.alias !== undefined && (obj.alias = message.alias);
    if (message.ownedCardSchemes) {
      obj.ownedCardSchemes = message.ownedCardSchemes.map((e) => e);
    } else {
      obj.ownedCardSchemes = [];
    }
    if (message.ownedPrototypes) {
      obj.ownedPrototypes = message.ownedPrototypes.map((e) => e);
    } else {
      obj.ownedPrototypes = [];
    }
    if (message.cards) {
      obj.cards = message.cards.map((e) => e);
    } else {
      obj.cards = [];
    }
    if (message.voteRights) {
      obj.voteRights = message.voteRights.map((e) =>
        e ? VoteRight.toJSON(e) : undefined
      );
    } else {
      obj.voteRights = [];
    }
    message.councilStatus !== undefined &&
      (obj.councilStatus = councilStatusToJSON(message.councilStatus));
    message.reportMatches !== undefined &&
      (obj.reportMatches = message.reportMatches);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryQUserResponse>): QueryQUserResponse {
    const message = { ...baseQueryQUserResponse } as QueryQUserResponse;
    message.ownedCardSchemes = [];
    message.ownedPrototypes = [];
    message.cards = [];
    message.voteRights = [];
    if (object.alias !== undefined && object.alias !== null) {
      message.alias = object.alias;
    } else {
      message.alias = "";
    }
    if (
      object.ownedCardSchemes !== undefined &&
      object.ownedCardSchemes !== null
    ) {
      for (const e of object.ownedCardSchemes) {
        message.ownedCardSchemes.push(e);
      }
    }
    if (
      object.ownedPrototypes !== undefined &&
      object.ownedPrototypes !== null
    ) {
      for (const e of object.ownedPrototypes) {
        message.ownedPrototypes.push(e);
      }
    }
    if (object.cards !== undefined && object.cards !== null) {
      for (const e of object.cards) {
        message.cards.push(e);
      }
    }
    if (object.voteRights !== undefined && object.voteRights !== null) {
      for (const e of object.voteRights) {
        message.voteRights.push(VoteRight.fromPartial(e));
      }
    }
    if (object.councilStatus !== undefined && object.councilStatus !== null) {
      message.councilStatus = object.councilStatus;
    } else {
      message.councilStatus = 0;
    }
    if (object.reportMatches !== undefined && object.reportMatches !== null) {
      message.reportMatches = object.reportMatches;
    } else {
      message.reportMatches = false;
    }
    return message;
  },
};

const baseQueryQCardchainInfoRequest: object = {};

export const QueryQCardchainInfoRequest = {
  encode(
    _: QueryQCardchainInfoRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryQCardchainInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryQCardchainInfoRequest,
    } as QueryQCardchainInfoRequest;
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
    const message = {
      ...baseQueryQCardchainInfoRequest,
    } as QueryQCardchainInfoRequest;
    return message;
  },

  toJSON(_: QueryQCardchainInfoRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryQCardchainInfoRequest>
  ): QueryQCardchainInfoRequest {
    const message = {
      ...baseQueryQCardchainInfoRequest,
    } as QueryQCardchainInfoRequest;
    return message;
  },
};

const baseQueryQCardchainInfoResponse: object = {
  cardAuctionPrice: "",
  activeCollections: 0,
};

export const QueryQCardchainInfoResponse = {
  encode(
    message: QueryQCardchainInfoResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.cardAuctionPrice !== "") {
      writer.uint32(10).string(message.cardAuctionPrice);
    }
    writer.uint32(18).fork();
    for (const v of message.activeCollections) {
      writer.uint64(v);
    }
    writer.ldelim();
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryQCardchainInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryQCardchainInfoResponse,
    } as QueryQCardchainInfoResponse;
    message.activeCollections = [];
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
              message.activeCollections.push(
                longToNumber(reader.uint64() as Long)
              );
            }
          } else {
            message.activeCollections.push(
              longToNumber(reader.uint64() as Long)
            );
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryQCardchainInfoResponse {
    const message = {
      ...baseQueryQCardchainInfoResponse,
    } as QueryQCardchainInfoResponse;
    message.activeCollections = [];
    if (
      object.cardAuctionPrice !== undefined &&
      object.cardAuctionPrice !== null
    ) {
      message.cardAuctionPrice = String(object.cardAuctionPrice);
    } else {
      message.cardAuctionPrice = "";
    }
    if (
      object.activeCollections !== undefined &&
      object.activeCollections !== null
    ) {
      for (const e of object.activeCollections) {
        message.activeCollections.push(Number(e));
      }
    }
    return message;
  },

  toJSON(message: QueryQCardchainInfoResponse): unknown {
    const obj: any = {};
    message.cardAuctionPrice !== undefined &&
      (obj.cardAuctionPrice = message.cardAuctionPrice);
    if (message.activeCollections) {
      obj.activeCollections = message.activeCollections.map((e) => e);
    } else {
      obj.activeCollections = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryQCardchainInfoResponse>
  ): QueryQCardchainInfoResponse {
    const message = {
      ...baseQueryQCardchainInfoResponse,
    } as QueryQCardchainInfoResponse;
    message.activeCollections = [];
    if (
      object.cardAuctionPrice !== undefined &&
      object.cardAuctionPrice !== null
    ) {
      message.cardAuctionPrice = object.cardAuctionPrice;
    } else {
      message.cardAuctionPrice = "";
    }
    if (
      object.activeCollections !== undefined &&
      object.activeCollections !== null
    ) {
      for (const e of object.activeCollections) {
        message.activeCollections.push(e);
      }
    }
    return message;
  },
};

const baseQueryQVotingResultsRequest: object = {};

export const QueryQVotingResultsRequest = {
  encode(
    _: QueryQVotingResultsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryQVotingResultsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryQVotingResultsRequest,
    } as QueryQVotingResultsRequest;
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
    const message = {
      ...baseQueryQVotingResultsRequest,
    } as QueryQVotingResultsRequest;
    return message;
  },

  toJSON(_: QueryQVotingResultsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryQVotingResultsRequest>
  ): QueryQVotingResultsRequest {
    const message = {
      ...baseQueryQVotingResultsRequest,
    } as QueryQVotingResultsRequest;
    return message;
  },
};

const baseQueryQVotingResultsResponse: object = {};

export const QueryQVotingResultsResponse = {
  encode(
    message: QueryQVotingResultsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.lastVotingResults !== undefined) {
      VotingResults.encode(
        message.lastVotingResults,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryQVotingResultsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryQVotingResultsResponse,
    } as QueryQVotingResultsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.lastVotingResults = VotingResults.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryQVotingResultsResponse {
    const message = {
      ...baseQueryQVotingResultsResponse,
    } as QueryQVotingResultsResponse;
    if (
      object.lastVotingResults !== undefined &&
      object.lastVotingResults !== null
    ) {
      message.lastVotingResults = VotingResults.fromJSON(
        object.lastVotingResults
      );
    } else {
      message.lastVotingResults = undefined;
    }
    return message;
  },

  toJSON(message: QueryQVotingResultsResponse): unknown {
    const obj: any = {};
    message.lastVotingResults !== undefined &&
      (obj.lastVotingResults = message.lastVotingResults
        ? VotingResults.toJSON(message.lastVotingResults)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryQVotingResultsResponse>
  ): QueryQVotingResultsResponse {
    const message = {
      ...baseQueryQVotingResultsResponse,
    } as QueryQVotingResultsResponse;
    if (
      object.lastVotingResults !== undefined &&
      object.lastVotingResults !== null
    ) {
      message.lastVotingResults = VotingResults.fromPartial(
        object.lastVotingResults
      );
    } else {
      message.lastVotingResults = undefined;
    }
    return message;
  },
};

const baseQueryQVotableCardsRequest: object = { address: "" };

export const QueryQVotableCardsRequest = {
  encode(
    message: QueryQVotableCardsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryQVotableCardsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryQVotableCardsRequest,
    } as QueryQVotableCardsRequest;
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
    const message = {
      ...baseQueryQVotableCardsRequest,
    } as QueryQVotableCardsRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: QueryQVotableCardsRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryQVotableCardsRequest>
  ): QueryQVotableCardsRequest {
    const message = {
      ...baseQueryQVotableCardsRequest,
    } as QueryQVotableCardsRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseQueryQVotableCardsResponse: object = {
  unregistered: false,
  noVoteRights: false,
};

export const QueryQVotableCardsResponse = {
  encode(
    message: QueryQVotableCardsResponse,
    writer: Writer = Writer.create()
  ): Writer {
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

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryQVotableCardsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryQVotableCardsResponse,
    } as QueryQVotableCardsResponse;
    message.voteRights = [];
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
    const message = {
      ...baseQueryQVotableCardsResponse,
    } as QueryQVotableCardsResponse;
    message.voteRights = [];
    if (object.unregistered !== undefined && object.unregistered !== null) {
      message.unregistered = Boolean(object.unregistered);
    } else {
      message.unregistered = false;
    }
    if (object.noVoteRights !== undefined && object.noVoteRights !== null) {
      message.noVoteRights = Boolean(object.noVoteRights);
    } else {
      message.noVoteRights = false;
    }
    if (object.voteRights !== undefined && object.voteRights !== null) {
      for (const e of object.voteRights) {
        message.voteRights.push(VoteRight.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: QueryQVotableCardsResponse): unknown {
    const obj: any = {};
    message.unregistered !== undefined &&
      (obj.unregistered = message.unregistered);
    message.noVoteRights !== undefined &&
      (obj.noVoteRights = message.noVoteRights);
    if (message.voteRights) {
      obj.voteRights = message.voteRights.map((e) =>
        e ? VoteRight.toJSON(e) : undefined
      );
    } else {
      obj.voteRights = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryQVotableCardsResponse>
  ): QueryQVotableCardsResponse {
    const message = {
      ...baseQueryQVotableCardsResponse,
    } as QueryQVotableCardsResponse;
    message.voteRights = [];
    if (object.unregistered !== undefined && object.unregistered !== null) {
      message.unregistered = object.unregistered;
    } else {
      message.unregistered = false;
    }
    if (object.noVoteRights !== undefined && object.noVoteRights !== null) {
      message.noVoteRights = object.noVoteRights;
    } else {
      message.noVoteRights = false;
    }
    if (object.voteRights !== undefined && object.voteRights !== null) {
      for (const e of object.voteRights) {
        message.voteRights.push(VoteRight.fromPartial(e));
      }
    }
    return message;
  },
};

const baseQueryQCardsRequest: object = {
  owner: "",
  status: 0,
  cardType: "",
  classes: "",
  sortBy: "",
  nameContains: "",
  keywordsContains: "",
  notesContains: "",
};

export const QueryQCardsRequest = {
  encode(
    message: QueryQCardsRequest,
    writer: Writer = Writer.create()
  ): Writer {
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

  decode(input: Reader | Uint8Array, length?: number): QueryQCardsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryQCardsRequest } as QueryQCardsRequest;
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
    const message = { ...baseQueryQCardsRequest } as QueryQCardsRequest;
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = statusFromJSON(object.status);
    } else {
      message.status = 0;
    }
    if (object.cardType !== undefined && object.cardType !== null) {
      message.cardType = String(object.cardType);
    } else {
      message.cardType = "";
    }
    if (object.classes !== undefined && object.classes !== null) {
      message.classes = String(object.classes);
    } else {
      message.classes = "";
    }
    if (object.sortBy !== undefined && object.sortBy !== null) {
      message.sortBy = String(object.sortBy);
    } else {
      message.sortBy = "";
    }
    if (object.nameContains !== undefined && object.nameContains !== null) {
      message.nameContains = String(object.nameContains);
    } else {
      message.nameContains = "";
    }
    if (
      object.keywordsContains !== undefined &&
      object.keywordsContains !== null
    ) {
      message.keywordsContains = String(object.keywordsContains);
    } else {
      message.keywordsContains = "";
    }
    if (object.notesContains !== undefined && object.notesContains !== null) {
      message.notesContains = String(object.notesContains);
    } else {
      message.notesContains = "";
    }
    return message;
  },

  toJSON(message: QueryQCardsRequest): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.status !== undefined && (obj.status = statusToJSON(message.status));
    message.cardType !== undefined && (obj.cardType = message.cardType);
    message.classes !== undefined && (obj.classes = message.classes);
    message.sortBy !== undefined && (obj.sortBy = message.sortBy);
    message.nameContains !== undefined &&
      (obj.nameContains = message.nameContains);
    message.keywordsContains !== undefined &&
      (obj.keywordsContains = message.keywordsContains);
    message.notesContains !== undefined &&
      (obj.notesContains = message.notesContains);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryQCardsRequest>): QueryQCardsRequest {
    const message = { ...baseQueryQCardsRequest } as QueryQCardsRequest;
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = 0;
    }
    if (object.cardType !== undefined && object.cardType !== null) {
      message.cardType = object.cardType;
    } else {
      message.cardType = "";
    }
    if (object.classes !== undefined && object.classes !== null) {
      message.classes = object.classes;
    } else {
      message.classes = "";
    }
    if (object.sortBy !== undefined && object.sortBy !== null) {
      message.sortBy = object.sortBy;
    } else {
      message.sortBy = "";
    }
    if (object.nameContains !== undefined && object.nameContains !== null) {
      message.nameContains = object.nameContains;
    } else {
      message.nameContains = "";
    }
    if (
      object.keywordsContains !== undefined &&
      object.keywordsContains !== null
    ) {
      message.keywordsContains = object.keywordsContains;
    } else {
      message.keywordsContains = "";
    }
    if (object.notesContains !== undefined && object.notesContains !== null) {
      message.notesContains = object.notesContains;
    } else {
      message.notesContains = "";
    }
    return message;
  },
};

const baseQueryQCardsResponse: object = { cardsList: 0 };

export const QueryQCardsResponse = {
  encode(
    message: QueryQCardsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    writer.uint32(10).fork();
    for (const v of message.cardsList) {
      writer.uint64(v);
    }
    writer.ldelim();
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryQCardsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryQCardsResponse } as QueryQCardsResponse;
    message.cardsList = [];
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
    const message = { ...baseQueryQCardsResponse } as QueryQCardsResponse;
    message.cardsList = [];
    if (object.cardsList !== undefined && object.cardsList !== null) {
      for (const e of object.cardsList) {
        message.cardsList.push(Number(e));
      }
    }
    return message;
  },

  toJSON(message: QueryQCardsResponse): unknown {
    const obj: any = {};
    if (message.cardsList) {
      obj.cardsList = message.cardsList.map((e) => e);
    } else {
      obj.cardsList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<QueryQCardsResponse>): QueryQCardsResponse {
    const message = { ...baseQueryQCardsResponse } as QueryQCardsResponse;
    message.cardsList = [];
    if (object.cardsList !== undefined && object.cardsList !== null) {
      for (const e of object.cardsList) {
        message.cardsList.push(e);
      }
    }
    return message;
  },
};

const baseQueryQMatchRequest: object = { matchId: 0 };

export const QueryQMatchRequest = {
  encode(
    message: QueryQMatchRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.matchId !== 0) {
      writer.uint32(8).uint64(message.matchId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryQMatchRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryQMatchRequest } as QueryQMatchRequest;
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
    const message = { ...baseQueryQMatchRequest } as QueryQMatchRequest;
    if (object.matchId !== undefined && object.matchId !== null) {
      message.matchId = Number(object.matchId);
    } else {
      message.matchId = 0;
    }
    return message;
  },

  toJSON(message: QueryQMatchRequest): unknown {
    const obj: any = {};
    message.matchId !== undefined && (obj.matchId = message.matchId);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryQMatchRequest>): QueryQMatchRequest {
    const message = { ...baseQueryQMatchRequest } as QueryQMatchRequest;
    if (object.matchId !== undefined && object.matchId !== null) {
      message.matchId = object.matchId;
    } else {
      message.matchId = 0;
    }
    return message;
  },
};

const baseQueryQMatchResponse: object = {
  timestamp: 0,
  reporter: "",
  playerA: "",
  playerB: "",
  outcome: 0,
};

export const QueryQMatchResponse = {
  encode(
    message: QueryQMatchResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.timestamp !== 0) {
      writer.uint32(8).uint64(message.timestamp);
    }
    if (message.reporter !== "") {
      writer.uint32(18).string(message.reporter);
    }
    if (message.playerA !== "") {
      writer.uint32(26).string(message.playerA);
    }
    if (message.playerB !== "") {
      writer.uint32(34).string(message.playerB);
    }
    if (message.outcome !== 0) {
      writer.uint32(40).int32(message.outcome);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryQMatchResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryQMatchResponse } as QueryQMatchResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.timestamp = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.reporter = reader.string();
          break;
        case 3:
          message.playerA = reader.string();
          break;
        case 4:
          message.playerB = reader.string();
          break;
        case 5:
          message.outcome = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryQMatchResponse {
    const message = { ...baseQueryQMatchResponse } as QueryQMatchResponse;
    if (object.timestamp !== undefined && object.timestamp !== null) {
      message.timestamp = Number(object.timestamp);
    } else {
      message.timestamp = 0;
    }
    if (object.reporter !== undefined && object.reporter !== null) {
      message.reporter = String(object.reporter);
    } else {
      message.reporter = "";
    }
    if (object.playerA !== undefined && object.playerA !== null) {
      message.playerA = String(object.playerA);
    } else {
      message.playerA = "";
    }
    if (object.playerB !== undefined && object.playerB !== null) {
      message.playerB = String(object.playerB);
    } else {
      message.playerB = "";
    }
    if (object.outcome !== undefined && object.outcome !== null) {
      message.outcome = outcomeFromJSON(object.outcome);
    } else {
      message.outcome = 0;
    }
    return message;
  },

  toJSON(message: QueryQMatchResponse): unknown {
    const obj: any = {};
    message.timestamp !== undefined && (obj.timestamp = message.timestamp);
    message.reporter !== undefined && (obj.reporter = message.reporter);
    message.playerA !== undefined && (obj.playerA = message.playerA);
    message.playerB !== undefined && (obj.playerB = message.playerB);
    message.outcome !== undefined &&
      (obj.outcome = outcomeToJSON(message.outcome));
    return obj;
  },

  fromPartial(object: DeepPartial<QueryQMatchResponse>): QueryQMatchResponse {
    const message = { ...baseQueryQMatchResponse } as QueryQMatchResponse;
    if (object.timestamp !== undefined && object.timestamp !== null) {
      message.timestamp = object.timestamp;
    } else {
      message.timestamp = 0;
    }
    if (object.reporter !== undefined && object.reporter !== null) {
      message.reporter = object.reporter;
    } else {
      message.reporter = "";
    }
    if (object.playerA !== undefined && object.playerA !== null) {
      message.playerA = object.playerA;
    } else {
      message.playerA = "";
    }
    if (object.playerB !== undefined && object.playerB !== null) {
      message.playerB = object.playerB;
    } else {
      message.playerB = "";
    }
    if (object.outcome !== undefined && object.outcome !== null) {
      message.outcome = object.outcome;
    } else {
      message.outcome = 0;
    }
    return message;
  },
};

const baseQueryQCollectionRequest: object = { collectionId: 0 };

export const QueryQCollectionRequest = {
  encode(
    message: QueryQCollectionRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.collectionId !== 0) {
      writer.uint32(8).uint64(message.collectionId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryQCollectionRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryQCollectionRequest,
    } as QueryQCollectionRequest;
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
    const message = {
      ...baseQueryQCollectionRequest,
    } as QueryQCollectionRequest;
    if (object.collectionId !== undefined && object.collectionId !== null) {
      message.collectionId = Number(object.collectionId);
    } else {
      message.collectionId = 0;
    }
    return message;
  },

  toJSON(message: QueryQCollectionRequest): unknown {
    const obj: any = {};
    message.collectionId !== undefined &&
      (obj.collectionId = message.collectionId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryQCollectionRequest>
  ): QueryQCollectionRequest {
    const message = {
      ...baseQueryQCollectionRequest,
    } as QueryQCollectionRequest;
    if (object.collectionId !== undefined && object.collectionId !== null) {
      message.collectionId = object.collectionId;
    } else {
      message.collectionId = 0;
    }
    return message;
  },
};

const baseQueryQCollectionResponse: object = {
  name: "",
  cards: 0,
  contributors: "",
  story: "",
  status: 0,
  expireBlock: 0,
};

export const QueryQCollectionResponse = {
  encode(
    message: QueryQCollectionResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    writer.uint32(18).fork();
    for (const v of message.cards) {
      writer.uint64(v);
    }
    writer.ldelim();
    for (const v of message.contributors) {
      writer.uint32(26).string(v!);
    }
    if (message.story !== "") {
      writer.uint32(34).string(message.story);
    }
    if (message.artwork.length !== 0) {
      writer.uint32(42).bytes(message.artwork);
    }
    if (message.status !== 0) {
      writer.uint32(48).int32(message.status);
    }
    if (message.expireBlock !== 0) {
      writer.uint32(56).int64(message.expireBlock);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryQCollectionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryQCollectionResponse,
    } as QueryQCollectionResponse;
    message.cards = [];
    message.contributors = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.cards.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.cards.push(longToNumber(reader.uint64() as Long));
          }
          break;
        case 3:
          message.contributors.push(reader.string());
          break;
        case 4:
          message.story = reader.string();
          break;
        case 5:
          message.artwork = reader.bytes();
          break;
        case 6:
          message.status = reader.int32() as any;
          break;
        case 7:
          message.expireBlock = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryQCollectionResponse {
    const message = {
      ...baseQueryQCollectionResponse,
    } as QueryQCollectionResponse;
    message.cards = [];
    message.contributors = [];
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.cards !== undefined && object.cards !== null) {
      for (const e of object.cards) {
        message.cards.push(Number(e));
      }
    }
    if (object.contributors !== undefined && object.contributors !== null) {
      for (const e of object.contributors) {
        message.contributors.push(String(e));
      }
    }
    if (object.story !== undefined && object.story !== null) {
      message.story = String(object.story);
    } else {
      message.story = "";
    }
    if (object.artwork !== undefined && object.artwork !== null) {
      message.artwork = bytesFromBase64(object.artwork);
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = cStatusFromJSON(object.status);
    } else {
      message.status = 0;
    }
    if (object.expireBlock !== undefined && object.expireBlock !== null) {
      message.expireBlock = Number(object.expireBlock);
    } else {
      message.expireBlock = 0;
    }
    return message;
  },

  toJSON(message: QueryQCollectionResponse): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    if (message.cards) {
      obj.cards = message.cards.map((e) => e);
    } else {
      obj.cards = [];
    }
    if (message.contributors) {
      obj.contributors = message.contributors.map((e) => e);
    } else {
      obj.contributors = [];
    }
    message.story !== undefined && (obj.story = message.story);
    message.artwork !== undefined &&
      (obj.artwork = base64FromBytes(
        message.artwork !== undefined ? message.artwork : new Uint8Array()
      ));
    message.status !== undefined &&
      (obj.status = cStatusToJSON(message.status));
    message.expireBlock !== undefined &&
      (obj.expireBlock = message.expireBlock);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryQCollectionResponse>
  ): QueryQCollectionResponse {
    const message = {
      ...baseQueryQCollectionResponse,
    } as QueryQCollectionResponse;
    message.cards = [];
    message.contributors = [];
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.cards !== undefined && object.cards !== null) {
      for (const e of object.cards) {
        message.cards.push(e);
      }
    }
    if (object.contributors !== undefined && object.contributors !== null) {
      for (const e of object.contributors) {
        message.contributors.push(e);
      }
    }
    if (object.story !== undefined && object.story !== null) {
      message.story = object.story;
    } else {
      message.story = "";
    }
    if (object.artwork !== undefined && object.artwork !== null) {
      message.artwork = object.artwork;
    } else {
      message.artwork = new Uint8Array();
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = 0;
    }
    if (object.expireBlock !== undefined && object.expireBlock !== null) {
      message.expireBlock = object.expireBlock;
    } else {
      message.expireBlock = 0;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of QCard items. */
  QCard(request: QueryQCardRequest): Promise<QueryQCardResponse>;
  /** Queries a list of QCardContent items. */
  QCardContent(
    request: QueryQCardContentRequest
  ): Promise<QueryQCardContentResponse>;
  /** Queries a list of QUser items. */
  QUser(request: QueryQUserRequest): Promise<QueryQUserResponse>;
  /** Queries a list of QCardchainInfo items. */
  QCardchainInfo(
    request: QueryQCardchainInfoRequest
  ): Promise<QueryQCardchainInfoResponse>;
  /** Queries a list of QVotingResults items. */
  QVotingResults(
    request: QueryQVotingResultsRequest
  ): Promise<QueryQVotingResultsResponse>;
  /** Queries a list of QVotableCards items. */
  QVotableCards(
    request: QueryQVotableCardsRequest
  ): Promise<QueryQVotableCardsResponse>;
  /** Queries a list of QCards items. */
  QCards(request: QueryQCardsRequest): Promise<QueryQCardsResponse>;
  /** Queries a list of QMatch items. */
  QMatch(request: QueryQMatchRequest): Promise<QueryQMatchResponse>;
  /** Queries a list of QCollection items. */
  QCollection(
    request: QueryQCollectionRequest
  ): Promise<QueryQCollectionResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  QCard(request: QueryQCardRequest): Promise<QueryQCardResponse> {
    const data = QueryQCardRequest.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Query",
      "QCard",
      data
    );
    return promise.then((data) => QueryQCardResponse.decode(new Reader(data)));
  }

  QCardContent(
    request: QueryQCardContentRequest
  ): Promise<QueryQCardContentResponse> {
    const data = QueryQCardContentRequest.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Query",
      "QCardContent",
      data
    );
    return promise.then((data) =>
      QueryQCardContentResponse.decode(new Reader(data))
    );
  }

  QUser(request: QueryQUserRequest): Promise<QueryQUserResponse> {
    const data = QueryQUserRequest.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Query",
      "QUser",
      data
    );
    return promise.then((data) => QueryQUserResponse.decode(new Reader(data)));
  }

  QCardchainInfo(
    request: QueryQCardchainInfoRequest
  ): Promise<QueryQCardchainInfoResponse> {
    const data = QueryQCardchainInfoRequest.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Query",
      "QCardchainInfo",
      data
    );
    return promise.then((data) =>
      QueryQCardchainInfoResponse.decode(new Reader(data))
    );
  }

  QVotingResults(
    request: QueryQVotingResultsRequest
  ): Promise<QueryQVotingResultsResponse> {
    const data = QueryQVotingResultsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Query",
      "QVotingResults",
      data
    );
    return promise.then((data) =>
      QueryQVotingResultsResponse.decode(new Reader(data))
    );
  }

  QVotableCards(
    request: QueryQVotableCardsRequest
  ): Promise<QueryQVotableCardsResponse> {
    const data = QueryQVotableCardsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Query",
      "QVotableCards",
      data
    );
    return promise.then((data) =>
      QueryQVotableCardsResponse.decode(new Reader(data))
    );
  }

  QCards(request: QueryQCardsRequest): Promise<QueryQCardsResponse> {
    const data = QueryQCardsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Query",
      "QCards",
      data
    );
    return promise.then((data) => QueryQCardsResponse.decode(new Reader(data)));
  }

  QMatch(request: QueryQMatchRequest): Promise<QueryQMatchResponse> {
    const data = QueryQMatchRequest.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Query",
      "QMatch",
      data
    );
    return promise.then((data) => QueryQMatchResponse.decode(new Reader(data)));
  }

  QCollection(
    request: QueryQCollectionRequest
  ): Promise<QueryQCollectionResponse> {
    const data = QueryQCollectionRequest.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Query",
      "QCollection",
      data
    );
    return promise.then((data) =>
      QueryQCollectionResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

const atob: (b64: string) => string =
  globalThis.atob ||
  ((b64) => globalThis.Buffer.from(b64, "base64").toString("binary"));
function bytesFromBase64(b64: string): Uint8Array {
  const bin = atob(b64);
  const arr = new Uint8Array(bin.length);
  for (let i = 0; i < bin.length; ++i) {
    arr[i] = bin.charCodeAt(i);
  }
  return arr;
}

const btoa: (bin: string) => string =
  globalThis.btoa ||
  ((bin) => globalThis.Buffer.from(bin, "binary").toString("base64"));
function base64FromBytes(arr: Uint8Array): string {
  const bin: string[] = [];
  for (let i = 0; i < arr.byteLength; ++i) {
    bin.push(String.fromCharCode(arr[i]));
  }
  return btoa(bin.join(""));
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
