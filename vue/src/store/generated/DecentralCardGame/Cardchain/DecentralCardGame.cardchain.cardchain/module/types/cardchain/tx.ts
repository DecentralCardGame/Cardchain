/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "DecentralCardGame.cardchain.cardchain";

export enum Outcome {
  AWon = 0,
  BWon = 1,
  Draw = 2,
  Aborted = 3,
  UNRECOGNIZED = -1,
}

export function outcomeFromJSON(object: any): Outcome {
  switch (object) {
    case 0:
    case "AWon":
      return Outcome.AWon;
    case 1:
    case "BWon":
      return Outcome.BWon;
    case 2:
    case "Draw":
      return Outcome.Draw;
    case 3:
    case "Aborted":
      return Outcome.Aborted;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Outcome.UNRECOGNIZED;
  }
}

export function outcomeToJSON(object: Outcome): string {
  switch (object) {
    case Outcome.AWon:
      return "AWon";
    case Outcome.BWon:
      return "BWon";
    case Outcome.Draw:
      return "Draw";
    case Outcome.Aborted:
      return "Aborted";
    default:
      return "UNKNOWN";
  }
}

export interface MsgCreateuser {
  creator: string;
  newUser: string;
  alias: string;
}

export interface MsgCreateuserResponse {}

export interface MsgBuyCardScheme {
  creator: string;
  bid: string;
  buyer: string;
}

export interface MsgBuyCardSchemeResponse {}

export interface MsgVoteCard {
  creator: string;
  cardId: number;
  voteType: string;
  voter: string;
}

export interface MsgVoteCardResponse {}

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

export interface MsgSaveCardContentResponse {}

export interface MsgTransferCard {
  creator: string;
  cardId: number;
  receiver: string;
}

export interface MsgTransferCardResponse {}

export interface MsgDonateToCard {
  creator: string;
  cardId: number;
  donator: string;
  amount: string;
}

export interface MsgDonateToCardResponse {}

export interface MsgAddArtwork {
  creator: string;
  cardId: number;
  image: Uint8Array;
  fullArt: boolean;
}

export interface MsgAddArtworkResponse {}

export interface MsgSubmitCopyrightProposal {
  creator: string;
  cardId: number;
  description: string;
  link: string;
}

export interface MsgSubmitCopyrightProposalResponse {}

export interface MsgChangeArtist {
  creator: string;
  cardID: number;
  artist: string;
}

export interface MsgChangeArtistResponse {}

export interface MsgRegisterForCouncil {
  creator: string;
}

export interface MsgRegisterForCouncilResponse {}

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

export interface MsgSubmitMatchReporterProposalResponse {}

export interface MsgApointMatchReporter {
  creator: string;
  reporter: string;
}

export interface MsgApointMatchReporterResponse {}

export interface MsgCreateCollection {
  creator: string;
  name: string;
  artist: string;
  storyWriter: string;
  contributors: string[];
}

export interface MsgCreateCollectionResponse {}

export interface MsgAddCardToCollection {
  creator: string;
  collectionId: number;
  cardId: number;
}

export interface MsgAddCardToCollectionResponse {}

export interface MsgFinalizeCollection {
  creator: string;
  collectionId: number;
}

export interface MsgFinalizeCollectionResponse {}

export interface MsgBuyCollection {
  creator: string;
  collectionId: number;
}

export interface MsgBuyCollectionResponse {}

export interface MsgRemoveCardFromCollection {
  creator: string;
  collectionId: number;
  cardId: number;
}

export interface MsgRemoveCardFromCollectionResponse {}

export interface MsgRemoveContributorFromCollection {
  creator: string;
  collectionId: number;
  user: string;
}

export interface MsgRemoveContributorFromCollectionResponse {}

export interface MsgAddContributorToCollection {
  creator: string;
  collectionId: number;
  user: string;
}

export interface MsgAddContributorToCollectionResponse {}

export interface MsgSubmitCollectionProposal {
  creator: string;
  collectionId: number;
}

export interface MsgSubmitCollectionProposalResponse {}

export interface MsgCreateSellOffer {
  creator: string;
  card: number;
  price: number;
}

export interface MsgCreateSellOfferResponse {}

export interface MsgBuyCard {
  creator: string;
  sellOfferId: number;
}

export interface MsgBuyCardResponse {}

export interface MsgRemoveSellOffer {
  creator: string;
  sellOfferId: number;
}

export interface MsgRemoveSellOfferResponse {}

const baseMsgCreateuser: object = { creator: "", newUser: "", alias: "" };

export const MsgCreateuser = {
  encode(message: MsgCreateuser, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.newUser !== "") {
      writer.uint32(18).string(message.newUser);
    }
    if (message.alias !== "") {
      writer.uint32(26).string(message.alias);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateuser {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateuser } as MsgCreateuser;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.newUser = reader.string();
          break;
        case 3:
          message.alias = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateuser {
    const message = { ...baseMsgCreateuser } as MsgCreateuser;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.newUser !== undefined && object.newUser !== null) {
      message.newUser = String(object.newUser);
    } else {
      message.newUser = "";
    }
    if (object.alias !== undefined && object.alias !== null) {
      message.alias = String(object.alias);
    } else {
      message.alias = "";
    }
    return message;
  },

  toJSON(message: MsgCreateuser): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.newUser !== undefined && (obj.newUser = message.newUser);
    message.alias !== undefined && (obj.alias = message.alias);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateuser>): MsgCreateuser {
    const message = { ...baseMsgCreateuser } as MsgCreateuser;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.newUser !== undefined && object.newUser !== null) {
      message.newUser = object.newUser;
    } else {
      message.newUser = "";
    }
    if (object.alias !== undefined && object.alias !== null) {
      message.alias = object.alias;
    } else {
      message.alias = "";
    }
    return message;
  },
};

const baseMsgCreateuserResponse: object = {};

export const MsgCreateuserResponse = {
  encode(_: MsgCreateuserResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateuserResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateuserResponse } as MsgCreateuserResponse;
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

  fromJSON(_: any): MsgCreateuserResponse {
    const message = { ...baseMsgCreateuserResponse } as MsgCreateuserResponse;
    return message;
  },

  toJSON(_: MsgCreateuserResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgCreateuserResponse>): MsgCreateuserResponse {
    const message = { ...baseMsgCreateuserResponse } as MsgCreateuserResponse;
    return message;
  },
};

const baseMsgBuyCardScheme: object = { creator: "", bid: "", buyer: "" };

export const MsgBuyCardScheme = {
  encode(message: MsgBuyCardScheme, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.bid !== "") {
      writer.uint32(18).string(message.bid);
    }
    if (message.buyer !== "") {
      writer.uint32(26).string(message.buyer);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgBuyCardScheme {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgBuyCardScheme } as MsgBuyCardScheme;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.bid = reader.string();
          break;
        case 3:
          message.buyer = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgBuyCardScheme {
    const message = { ...baseMsgBuyCardScheme } as MsgBuyCardScheme;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.bid !== undefined && object.bid !== null) {
      message.bid = String(object.bid);
    } else {
      message.bid = "";
    }
    if (object.buyer !== undefined && object.buyer !== null) {
      message.buyer = String(object.buyer);
    } else {
      message.buyer = "";
    }
    return message;
  },

  toJSON(message: MsgBuyCardScheme): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.bid !== undefined && (obj.bid = message.bid);
    message.buyer !== undefined && (obj.buyer = message.buyer);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgBuyCardScheme>): MsgBuyCardScheme {
    const message = { ...baseMsgBuyCardScheme } as MsgBuyCardScheme;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.bid !== undefined && object.bid !== null) {
      message.bid = object.bid;
    } else {
      message.bid = "";
    }
    if (object.buyer !== undefined && object.buyer !== null) {
      message.buyer = object.buyer;
    } else {
      message.buyer = "";
    }
    return message;
  },
};

const baseMsgBuyCardSchemeResponse: object = {};

export const MsgBuyCardSchemeResponse = {
  encode(
    _: MsgBuyCardSchemeResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgBuyCardSchemeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgBuyCardSchemeResponse,
    } as MsgBuyCardSchemeResponse;
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

  fromJSON(_: any): MsgBuyCardSchemeResponse {
    const message = {
      ...baseMsgBuyCardSchemeResponse,
    } as MsgBuyCardSchemeResponse;
    return message;
  },

  toJSON(_: MsgBuyCardSchemeResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgBuyCardSchemeResponse>
  ): MsgBuyCardSchemeResponse {
    const message = {
      ...baseMsgBuyCardSchemeResponse,
    } as MsgBuyCardSchemeResponse;
    return message;
  },
};

const baseMsgVoteCard: object = {
  creator: "",
  cardId: 0,
  voteType: "",
  voter: "",
};

export const MsgVoteCard = {
  encode(message: MsgVoteCard, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.cardId !== 0) {
      writer.uint32(16).uint64(message.cardId);
    }
    if (message.voteType !== "") {
      writer.uint32(26).string(message.voteType);
    }
    if (message.voter !== "") {
      writer.uint32(34).string(message.voter);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgVoteCard {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgVoteCard } as MsgVoteCard;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.cardId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.voteType = reader.string();
          break;
        case 4:
          message.voter = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgVoteCard {
    const message = { ...baseMsgVoteCard } as MsgVoteCard;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.cardId !== undefined && object.cardId !== null) {
      message.cardId = Number(object.cardId);
    } else {
      message.cardId = 0;
    }
    if (object.voteType !== undefined && object.voteType !== null) {
      message.voteType = String(object.voteType);
    } else {
      message.voteType = "";
    }
    if (object.voter !== undefined && object.voter !== null) {
      message.voter = String(object.voter);
    } else {
      message.voter = "";
    }
    return message;
  },

  toJSON(message: MsgVoteCard): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.cardId !== undefined && (obj.cardId = message.cardId);
    message.voteType !== undefined && (obj.voteType = message.voteType);
    message.voter !== undefined && (obj.voter = message.voter);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgVoteCard>): MsgVoteCard {
    const message = { ...baseMsgVoteCard } as MsgVoteCard;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.cardId !== undefined && object.cardId !== null) {
      message.cardId = object.cardId;
    } else {
      message.cardId = 0;
    }
    if (object.voteType !== undefined && object.voteType !== null) {
      message.voteType = object.voteType;
    } else {
      message.voteType = "";
    }
    if (object.voter !== undefined && object.voter !== null) {
      message.voter = object.voter;
    } else {
      message.voter = "";
    }
    return message;
  },
};

const baseMsgVoteCardResponse: object = {};

export const MsgVoteCardResponse = {
  encode(_: MsgVoteCardResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgVoteCardResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgVoteCardResponse } as MsgVoteCardResponse;
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

  fromJSON(_: any): MsgVoteCardResponse {
    const message = { ...baseMsgVoteCardResponse } as MsgVoteCardResponse;
    return message;
  },

  toJSON(_: MsgVoteCardResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgVoteCardResponse>): MsgVoteCardResponse {
    const message = { ...baseMsgVoteCardResponse } as MsgVoteCardResponse;
    return message;
  },
};

const baseMsgSaveCardContent: object = {
  creator: "",
  cardId: 0,
  notes: "",
  owner: "",
  artist: "",
};

export const MsgSaveCardContent = {
  encode(
    message: MsgSaveCardContent,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.cardId !== 0) {
      writer.uint32(16).uint64(message.cardId);
    }
    if (message.content.length !== 0) {
      writer.uint32(26).bytes(message.content);
    }
    if (message.notes !== "") {
      writer.uint32(34).string(message.notes);
    }
    if (message.owner !== "") {
      writer.uint32(42).string(message.owner);
    }
    if (message.artist !== "") {
      writer.uint32(50).string(message.artist);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSaveCardContent {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSaveCardContent } as MsgSaveCardContent;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.cardId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.content = reader.bytes();
          break;
        case 4:
          message.notes = reader.string();
          break;
        case 5:
          message.owner = reader.string();
          break;
        case 6:
          message.artist = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSaveCardContent {
    const message = { ...baseMsgSaveCardContent } as MsgSaveCardContent;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.cardId !== undefined && object.cardId !== null) {
      message.cardId = Number(object.cardId);
    } else {
      message.cardId = 0;
    }
    if (object.content !== undefined && object.content !== null) {
      message.content = bytesFromBase64(object.content);
    }
    if (object.notes !== undefined && object.notes !== null) {
      message.notes = String(object.notes);
    } else {
      message.notes = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.artist !== undefined && object.artist !== null) {
      message.artist = String(object.artist);
    } else {
      message.artist = "";
    }
    return message;
  },

  toJSON(message: MsgSaveCardContent): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.cardId !== undefined && (obj.cardId = message.cardId);
    message.content !== undefined &&
      (obj.content = base64FromBytes(
        message.content !== undefined ? message.content : new Uint8Array()
      ));
    message.notes !== undefined && (obj.notes = message.notes);
    message.owner !== undefined && (obj.owner = message.owner);
    message.artist !== undefined && (obj.artist = message.artist);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSaveCardContent>): MsgSaveCardContent {
    const message = { ...baseMsgSaveCardContent } as MsgSaveCardContent;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.cardId !== undefined && object.cardId !== null) {
      message.cardId = object.cardId;
    } else {
      message.cardId = 0;
    }
    if (object.content !== undefined && object.content !== null) {
      message.content = object.content;
    } else {
      message.content = new Uint8Array();
    }
    if (object.notes !== undefined && object.notes !== null) {
      message.notes = object.notes;
    } else {
      message.notes = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.artist !== undefined && object.artist !== null) {
      message.artist = object.artist;
    } else {
      message.artist = "";
    }
    return message;
  },
};

const baseMsgSaveCardContentResponse: object = {};

export const MsgSaveCardContentResponse = {
  encode(
    _: MsgSaveCardContentResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSaveCardContentResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSaveCardContentResponse,
    } as MsgSaveCardContentResponse;
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

  fromJSON(_: any): MsgSaveCardContentResponse {
    const message = {
      ...baseMsgSaveCardContentResponse,
    } as MsgSaveCardContentResponse;
    return message;
  },

  toJSON(_: MsgSaveCardContentResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgSaveCardContentResponse>
  ): MsgSaveCardContentResponse {
    const message = {
      ...baseMsgSaveCardContentResponse,
    } as MsgSaveCardContentResponse;
    return message;
  },
};

const baseMsgTransferCard: object = { creator: "", cardId: 0, receiver: "" };

export const MsgTransferCard = {
  encode(message: MsgTransferCard, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.cardId !== 0) {
      writer.uint32(16).uint64(message.cardId);
    }
    if (message.receiver !== "") {
      writer.uint32(34).string(message.receiver);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgTransferCard {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgTransferCard } as MsgTransferCard;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.cardId = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.receiver = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgTransferCard {
    const message = { ...baseMsgTransferCard } as MsgTransferCard;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.cardId !== undefined && object.cardId !== null) {
      message.cardId = Number(object.cardId);
    } else {
      message.cardId = 0;
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = String(object.receiver);
    } else {
      message.receiver = "";
    }
    return message;
  },

  toJSON(message: MsgTransferCard): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.cardId !== undefined && (obj.cardId = message.cardId);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgTransferCard>): MsgTransferCard {
    const message = { ...baseMsgTransferCard } as MsgTransferCard;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.cardId !== undefined && object.cardId !== null) {
      message.cardId = object.cardId;
    } else {
      message.cardId = 0;
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = object.receiver;
    } else {
      message.receiver = "";
    }
    return message;
  },
};

const baseMsgTransferCardResponse: object = {};

export const MsgTransferCardResponse = {
  encode(_: MsgTransferCardResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgTransferCardResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgTransferCardResponse,
    } as MsgTransferCardResponse;
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

  fromJSON(_: any): MsgTransferCardResponse {
    const message = {
      ...baseMsgTransferCardResponse,
    } as MsgTransferCardResponse;
    return message;
  },

  toJSON(_: MsgTransferCardResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgTransferCardResponse>
  ): MsgTransferCardResponse {
    const message = {
      ...baseMsgTransferCardResponse,
    } as MsgTransferCardResponse;
    return message;
  },
};

const baseMsgDonateToCard: object = {
  creator: "",
  cardId: 0,
  donator: "",
  amount: "",
};

export const MsgDonateToCard = {
  encode(message: MsgDonateToCard, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.cardId !== 0) {
      writer.uint32(16).uint64(message.cardId);
    }
    if (message.donator !== "") {
      writer.uint32(26).string(message.donator);
    }
    if (message.amount !== "") {
      writer.uint32(34).string(message.amount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDonateToCard {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDonateToCard } as MsgDonateToCard;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.cardId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.donator = reader.string();
          break;
        case 4:
          message.amount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDonateToCard {
    const message = { ...baseMsgDonateToCard } as MsgDonateToCard;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.cardId !== undefined && object.cardId !== null) {
      message.cardId = Number(object.cardId);
    } else {
      message.cardId = 0;
    }
    if (object.donator !== undefined && object.donator !== null) {
      message.donator = String(object.donator);
    } else {
      message.donator = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    return message;
  },

  toJSON(message: MsgDonateToCard): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.cardId !== undefined && (obj.cardId = message.cardId);
    message.donator !== undefined && (obj.donator = message.donator);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDonateToCard>): MsgDonateToCard {
    const message = { ...baseMsgDonateToCard } as MsgDonateToCard;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.cardId !== undefined && object.cardId !== null) {
      message.cardId = object.cardId;
    } else {
      message.cardId = 0;
    }
    if (object.donator !== undefined && object.donator !== null) {
      message.donator = object.donator;
    } else {
      message.donator = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    return message;
  },
};

const baseMsgDonateToCardResponse: object = {};

export const MsgDonateToCardResponse = {
  encode(_: MsgDonateToCardResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDonateToCardResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDonateToCardResponse,
    } as MsgDonateToCardResponse;
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

  fromJSON(_: any): MsgDonateToCardResponse {
    const message = {
      ...baseMsgDonateToCardResponse,
    } as MsgDonateToCardResponse;
    return message;
  },

  toJSON(_: MsgDonateToCardResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDonateToCardResponse>
  ): MsgDonateToCardResponse {
    const message = {
      ...baseMsgDonateToCardResponse,
    } as MsgDonateToCardResponse;
    return message;
  },
};

const baseMsgAddArtwork: object = { creator: "", cardId: 0, fullArt: false };

export const MsgAddArtwork = {
  encode(message: MsgAddArtwork, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.cardId !== 0) {
      writer.uint32(16).uint64(message.cardId);
    }
    if (message.image.length !== 0) {
      writer.uint32(26).bytes(message.image);
    }
    if (message.fullArt === true) {
      writer.uint32(32).bool(message.fullArt);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddArtwork {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddArtwork } as MsgAddArtwork;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.cardId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.image = reader.bytes();
          break;
        case 4:
          message.fullArt = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddArtwork {
    const message = { ...baseMsgAddArtwork } as MsgAddArtwork;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.cardId !== undefined && object.cardId !== null) {
      message.cardId = Number(object.cardId);
    } else {
      message.cardId = 0;
    }
    if (object.image !== undefined && object.image !== null) {
      message.image = bytesFromBase64(object.image);
    }
    if (object.fullArt !== undefined && object.fullArt !== null) {
      message.fullArt = Boolean(object.fullArt);
    } else {
      message.fullArt = false;
    }
    return message;
  },

  toJSON(message: MsgAddArtwork): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.cardId !== undefined && (obj.cardId = message.cardId);
    message.image !== undefined &&
      (obj.image = base64FromBytes(
        message.image !== undefined ? message.image : new Uint8Array()
      ));
    message.fullArt !== undefined && (obj.fullArt = message.fullArt);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgAddArtwork>): MsgAddArtwork {
    const message = { ...baseMsgAddArtwork } as MsgAddArtwork;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.cardId !== undefined && object.cardId !== null) {
      message.cardId = object.cardId;
    } else {
      message.cardId = 0;
    }
    if (object.image !== undefined && object.image !== null) {
      message.image = object.image;
    } else {
      message.image = new Uint8Array();
    }
    if (object.fullArt !== undefined && object.fullArt !== null) {
      message.fullArt = object.fullArt;
    } else {
      message.fullArt = false;
    }
    return message;
  },
};

const baseMsgAddArtworkResponse: object = {};

export const MsgAddArtworkResponse = {
  encode(_: MsgAddArtworkResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddArtworkResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddArtworkResponse } as MsgAddArtworkResponse;
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

  fromJSON(_: any): MsgAddArtworkResponse {
    const message = { ...baseMsgAddArtworkResponse } as MsgAddArtworkResponse;
    return message;
  },

  toJSON(_: MsgAddArtworkResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgAddArtworkResponse>): MsgAddArtworkResponse {
    const message = { ...baseMsgAddArtworkResponse } as MsgAddArtworkResponse;
    return message;
  },
};

const baseMsgSubmitCopyrightProposal: object = {
  creator: "",
  cardId: 0,
  description: "",
  link: "",
};

export const MsgSubmitCopyrightProposal = {
  encode(
    message: MsgSubmitCopyrightProposal,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.cardId !== 0) {
      writer.uint32(16).uint64(message.cardId);
    }
    if (message.description !== "") {
      writer.uint32(26).string(message.description);
    }
    if (message.link !== "") {
      writer.uint32(34).string(message.link);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSubmitCopyrightProposal {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSubmitCopyrightProposal,
    } as MsgSubmitCopyrightProposal;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.cardId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.description = reader.string();
          break;
        case 4:
          message.link = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSubmitCopyrightProposal {
    const message = {
      ...baseMsgSubmitCopyrightProposal,
    } as MsgSubmitCopyrightProposal;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.cardId !== undefined && object.cardId !== null) {
      message.cardId = Number(object.cardId);
    } else {
      message.cardId = 0;
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    if (object.link !== undefined && object.link !== null) {
      message.link = String(object.link);
    } else {
      message.link = "";
    }
    return message;
  },

  toJSON(message: MsgSubmitCopyrightProposal): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.cardId !== undefined && (obj.cardId = message.cardId);
    message.description !== undefined &&
      (obj.description = message.description);
    message.link !== undefined && (obj.link = message.link);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSubmitCopyrightProposal>
  ): MsgSubmitCopyrightProposal {
    const message = {
      ...baseMsgSubmitCopyrightProposal,
    } as MsgSubmitCopyrightProposal;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.cardId !== undefined && object.cardId !== null) {
      message.cardId = object.cardId;
    } else {
      message.cardId = 0;
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    if (object.link !== undefined && object.link !== null) {
      message.link = object.link;
    } else {
      message.link = "";
    }
    return message;
  },
};

const baseMsgSubmitCopyrightProposalResponse: object = {};

export const MsgSubmitCopyrightProposalResponse = {
  encode(
    _: MsgSubmitCopyrightProposalResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSubmitCopyrightProposalResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSubmitCopyrightProposalResponse,
    } as MsgSubmitCopyrightProposalResponse;
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

  fromJSON(_: any): MsgSubmitCopyrightProposalResponse {
    const message = {
      ...baseMsgSubmitCopyrightProposalResponse,
    } as MsgSubmitCopyrightProposalResponse;
    return message;
  },

  toJSON(_: MsgSubmitCopyrightProposalResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgSubmitCopyrightProposalResponse>
  ): MsgSubmitCopyrightProposalResponse {
    const message = {
      ...baseMsgSubmitCopyrightProposalResponse,
    } as MsgSubmitCopyrightProposalResponse;
    return message;
  },
};

const baseMsgChangeArtist: object = { creator: "", cardID: 0, artist: "" };

export const MsgChangeArtist = {
  encode(message: MsgChangeArtist, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.cardID !== 0) {
      writer.uint32(16).uint64(message.cardID);
    }
    if (message.artist !== "") {
      writer.uint32(26).string(message.artist);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgChangeArtist {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgChangeArtist } as MsgChangeArtist;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.cardID = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.artist = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgChangeArtist {
    const message = { ...baseMsgChangeArtist } as MsgChangeArtist;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.cardID !== undefined && object.cardID !== null) {
      message.cardID = Number(object.cardID);
    } else {
      message.cardID = 0;
    }
    if (object.artist !== undefined && object.artist !== null) {
      message.artist = String(object.artist);
    } else {
      message.artist = "";
    }
    return message;
  },

  toJSON(message: MsgChangeArtist): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.cardID !== undefined && (obj.cardID = message.cardID);
    message.artist !== undefined && (obj.artist = message.artist);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgChangeArtist>): MsgChangeArtist {
    const message = { ...baseMsgChangeArtist } as MsgChangeArtist;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.cardID !== undefined && object.cardID !== null) {
      message.cardID = object.cardID;
    } else {
      message.cardID = 0;
    }
    if (object.artist !== undefined && object.artist !== null) {
      message.artist = object.artist;
    } else {
      message.artist = "";
    }
    return message;
  },
};

const baseMsgChangeArtistResponse: object = {};

export const MsgChangeArtistResponse = {
  encode(_: MsgChangeArtistResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgChangeArtistResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgChangeArtistResponse,
    } as MsgChangeArtistResponse;
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

  fromJSON(_: any): MsgChangeArtistResponse {
    const message = {
      ...baseMsgChangeArtistResponse,
    } as MsgChangeArtistResponse;
    return message;
  },

  toJSON(_: MsgChangeArtistResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgChangeArtistResponse>
  ): MsgChangeArtistResponse {
    const message = {
      ...baseMsgChangeArtistResponse,
    } as MsgChangeArtistResponse;
    return message;
  },
};

const baseMsgRegisterForCouncil: object = { creator: "" };

export const MsgRegisterForCouncil = {
  encode(
    message: MsgRegisterForCouncil,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRegisterForCouncil {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRegisterForCouncil } as MsgRegisterForCouncil;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRegisterForCouncil {
    const message = { ...baseMsgRegisterForCouncil } as MsgRegisterForCouncil;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: MsgRegisterForCouncil): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgRegisterForCouncil>
  ): MsgRegisterForCouncil {
    const message = { ...baseMsgRegisterForCouncil } as MsgRegisterForCouncil;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseMsgRegisterForCouncilResponse: object = {};

export const MsgRegisterForCouncilResponse = {
  encode(
    _: MsgRegisterForCouncilResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgRegisterForCouncilResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgRegisterForCouncilResponse,
    } as MsgRegisterForCouncilResponse;
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

  fromJSON(_: any): MsgRegisterForCouncilResponse {
    const message = {
      ...baseMsgRegisterForCouncilResponse,
    } as MsgRegisterForCouncilResponse;
    return message;
  },

  toJSON(_: MsgRegisterForCouncilResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgRegisterForCouncilResponse>
  ): MsgRegisterForCouncilResponse {
    const message = {
      ...baseMsgRegisterForCouncilResponse,
    } as MsgRegisterForCouncilResponse;
    return message;
  },
};

const baseMsgReportMatch: object = {
  creator: "",
  playerA: "",
  playerB: "",
  cardsA: 0,
  cardsB: 0,
  outcome: 0,
};

export const MsgReportMatch = {
  encode(message: MsgReportMatch, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.playerA !== "") {
      writer.uint32(18).string(message.playerA);
    }
    if (message.playerB !== "") {
      writer.uint32(26).string(message.playerB);
    }
    writer.uint32(42).fork();
    for (const v of message.cardsA) {
      writer.uint64(v);
    }
    writer.ldelim();
    writer.uint32(50).fork();
    for (const v of message.cardsB) {
      writer.uint64(v);
    }
    writer.ldelim();
    if (message.outcome !== 0) {
      writer.uint32(56).int32(message.outcome);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgReportMatch {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgReportMatch } as MsgReportMatch;
    message.cardsA = [];
    message.cardsB = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.playerA = reader.string();
          break;
        case 3:
          message.playerB = reader.string();
          break;
        case 5:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.cardsA.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.cardsA.push(longToNumber(reader.uint64() as Long));
          }
          break;
        case 6:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.cardsB.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.cardsB.push(longToNumber(reader.uint64() as Long));
          }
          break;
        case 7:
          message.outcome = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgReportMatch {
    const message = { ...baseMsgReportMatch } as MsgReportMatch;
    message.cardsA = [];
    message.cardsB = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
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
    if (object.cardsA !== undefined && object.cardsA !== null) {
      for (const e of object.cardsA) {
        message.cardsA.push(Number(e));
      }
    }
    if (object.cardsB !== undefined && object.cardsB !== null) {
      for (const e of object.cardsB) {
        message.cardsB.push(Number(e));
      }
    }
    if (object.outcome !== undefined && object.outcome !== null) {
      message.outcome = outcomeFromJSON(object.outcome);
    } else {
      message.outcome = 0;
    }
    return message;
  },

  toJSON(message: MsgReportMatch): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.playerA !== undefined && (obj.playerA = message.playerA);
    message.playerB !== undefined && (obj.playerB = message.playerB);
    if (message.cardsA) {
      obj.cardsA = message.cardsA.map((e) => e);
    } else {
      obj.cardsA = [];
    }
    if (message.cardsB) {
      obj.cardsB = message.cardsB.map((e) => e);
    } else {
      obj.cardsB = [];
    }
    message.outcome !== undefined &&
      (obj.outcome = outcomeToJSON(message.outcome));
    return obj;
  },

  fromPartial(object: DeepPartial<MsgReportMatch>): MsgReportMatch {
    const message = { ...baseMsgReportMatch } as MsgReportMatch;
    message.cardsA = [];
    message.cardsB = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
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
    if (object.cardsA !== undefined && object.cardsA !== null) {
      for (const e of object.cardsA) {
        message.cardsA.push(e);
      }
    }
    if (object.cardsB !== undefined && object.cardsB !== null) {
      for (const e of object.cardsB) {
        message.cardsB.push(e);
      }
    }
    if (object.outcome !== undefined && object.outcome !== null) {
      message.outcome = object.outcome;
    } else {
      message.outcome = 0;
    }
    return message;
  },
};

const baseMsgReportMatchResponse: object = { matchId: 0 };

export const MsgReportMatchResponse = {
  encode(
    message: MsgReportMatchResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.matchId !== 0) {
      writer.uint32(8).uint64(message.matchId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgReportMatchResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgReportMatchResponse } as MsgReportMatchResponse;
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

  fromJSON(object: any): MsgReportMatchResponse {
    const message = { ...baseMsgReportMatchResponse } as MsgReportMatchResponse;
    if (object.matchId !== undefined && object.matchId !== null) {
      message.matchId = Number(object.matchId);
    } else {
      message.matchId = 0;
    }
    return message;
  },

  toJSON(message: MsgReportMatchResponse): unknown {
    const obj: any = {};
    message.matchId !== undefined && (obj.matchId = message.matchId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgReportMatchResponse>
  ): MsgReportMatchResponse {
    const message = { ...baseMsgReportMatchResponse } as MsgReportMatchResponse;
    if (object.matchId !== undefined && object.matchId !== null) {
      message.matchId = object.matchId;
    } else {
      message.matchId = 0;
    }
    return message;
  },
};

const baseMsgSubmitMatchReporterProposal: object = {
  creator: "",
  reporter: "",
  deposit: "",
  description: "",
};

export const MsgSubmitMatchReporterProposal = {
  encode(
    message: MsgSubmitMatchReporterProposal,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.reporter !== "") {
      writer.uint32(18).string(message.reporter);
    }
    if (message.deposit !== "") {
      writer.uint32(26).string(message.deposit);
    }
    if (message.description !== "") {
      writer.uint32(34).string(message.description);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSubmitMatchReporterProposal {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSubmitMatchReporterProposal,
    } as MsgSubmitMatchReporterProposal;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.reporter = reader.string();
          break;
        case 3:
          message.deposit = reader.string();
          break;
        case 4:
          message.description = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSubmitMatchReporterProposal {
    const message = {
      ...baseMsgSubmitMatchReporterProposal,
    } as MsgSubmitMatchReporterProposal;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.reporter !== undefined && object.reporter !== null) {
      message.reporter = String(object.reporter);
    } else {
      message.reporter = "";
    }
    if (object.deposit !== undefined && object.deposit !== null) {
      message.deposit = String(object.deposit);
    } else {
      message.deposit = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    return message;
  },

  toJSON(message: MsgSubmitMatchReporterProposal): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.reporter !== undefined && (obj.reporter = message.reporter);
    message.deposit !== undefined && (obj.deposit = message.deposit);
    message.description !== undefined &&
      (obj.description = message.description);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSubmitMatchReporterProposal>
  ): MsgSubmitMatchReporterProposal {
    const message = {
      ...baseMsgSubmitMatchReporterProposal,
    } as MsgSubmitMatchReporterProposal;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.reporter !== undefined && object.reporter !== null) {
      message.reporter = object.reporter;
    } else {
      message.reporter = "";
    }
    if (object.deposit !== undefined && object.deposit !== null) {
      message.deposit = object.deposit;
    } else {
      message.deposit = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    return message;
  },
};

const baseMsgSubmitMatchReporterProposalResponse: object = {};

export const MsgSubmitMatchReporterProposalResponse = {
  encode(
    _: MsgSubmitMatchReporterProposalResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSubmitMatchReporterProposalResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSubmitMatchReporterProposalResponse,
    } as MsgSubmitMatchReporterProposalResponse;
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

  fromJSON(_: any): MsgSubmitMatchReporterProposalResponse {
    const message = {
      ...baseMsgSubmitMatchReporterProposalResponse,
    } as MsgSubmitMatchReporterProposalResponse;
    return message;
  },

  toJSON(_: MsgSubmitMatchReporterProposalResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgSubmitMatchReporterProposalResponse>
  ): MsgSubmitMatchReporterProposalResponse {
    const message = {
      ...baseMsgSubmitMatchReporterProposalResponse,
    } as MsgSubmitMatchReporterProposalResponse;
    return message;
  },
};

const baseMsgApointMatchReporter: object = { creator: "", reporter: "" };

export const MsgApointMatchReporter = {
  encode(
    message: MsgApointMatchReporter,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.reporter !== "") {
      writer.uint32(18).string(message.reporter);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgApointMatchReporter {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgApointMatchReporter } as MsgApointMatchReporter;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.reporter = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgApointMatchReporter {
    const message = { ...baseMsgApointMatchReporter } as MsgApointMatchReporter;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.reporter !== undefined && object.reporter !== null) {
      message.reporter = String(object.reporter);
    } else {
      message.reporter = "";
    }
    return message;
  },

  toJSON(message: MsgApointMatchReporter): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.reporter !== undefined && (obj.reporter = message.reporter);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgApointMatchReporter>
  ): MsgApointMatchReporter {
    const message = { ...baseMsgApointMatchReporter } as MsgApointMatchReporter;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.reporter !== undefined && object.reporter !== null) {
      message.reporter = object.reporter;
    } else {
      message.reporter = "";
    }
    return message;
  },
};

const baseMsgApointMatchReporterResponse: object = {};

export const MsgApointMatchReporterResponse = {
  encode(
    _: MsgApointMatchReporterResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgApointMatchReporterResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgApointMatchReporterResponse,
    } as MsgApointMatchReporterResponse;
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

  fromJSON(_: any): MsgApointMatchReporterResponse {
    const message = {
      ...baseMsgApointMatchReporterResponse,
    } as MsgApointMatchReporterResponse;
    return message;
  },

  toJSON(_: MsgApointMatchReporterResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgApointMatchReporterResponse>
  ): MsgApointMatchReporterResponse {
    const message = {
      ...baseMsgApointMatchReporterResponse,
    } as MsgApointMatchReporterResponse;
    return message;
  },
};

const baseMsgCreateCollection: object = {
  creator: "",
  name: "",
  artist: "",
  storyWriter: "",
  contributors: "",
};

export const MsgCreateCollection = {
  encode(
    message: MsgCreateCollection,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.artist !== "") {
      writer.uint32(26).string(message.artist);
    }
    if (message.storyWriter !== "") {
      writer.uint32(34).string(message.storyWriter);
    }
    for (const v of message.contributors) {
      writer.uint32(42).string(v!);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateCollection {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateCollection } as MsgCreateCollection;
    message.contributors = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.artist = reader.string();
          break;
        case 4:
          message.storyWriter = reader.string();
          break;
        case 5:
          message.contributors.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateCollection {
    const message = { ...baseMsgCreateCollection } as MsgCreateCollection;
    message.contributors = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.artist !== undefined && object.artist !== null) {
      message.artist = String(object.artist);
    } else {
      message.artist = "";
    }
    if (object.storyWriter !== undefined && object.storyWriter !== null) {
      message.storyWriter = String(object.storyWriter);
    } else {
      message.storyWriter = "";
    }
    if (object.contributors !== undefined && object.contributors !== null) {
      for (const e of object.contributors) {
        message.contributors.push(String(e));
      }
    }
    return message;
  },

  toJSON(message: MsgCreateCollection): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    message.artist !== undefined && (obj.artist = message.artist);
    message.storyWriter !== undefined &&
      (obj.storyWriter = message.storyWriter);
    if (message.contributors) {
      obj.contributors = message.contributors.map((e) => e);
    } else {
      obj.contributors = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateCollection>): MsgCreateCollection {
    const message = { ...baseMsgCreateCollection } as MsgCreateCollection;
    message.contributors = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.artist !== undefined && object.artist !== null) {
      message.artist = object.artist;
    } else {
      message.artist = "";
    }
    if (object.storyWriter !== undefined && object.storyWriter !== null) {
      message.storyWriter = object.storyWriter;
    } else {
      message.storyWriter = "";
    }
    if (object.contributors !== undefined && object.contributors !== null) {
      for (const e of object.contributors) {
        message.contributors.push(e);
      }
    }
    return message;
  },
};

const baseMsgCreateCollectionResponse: object = {};

export const MsgCreateCollectionResponse = {
  encode(
    _: MsgCreateCollectionResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateCollectionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateCollectionResponse,
    } as MsgCreateCollectionResponse;
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

  fromJSON(_: any): MsgCreateCollectionResponse {
    const message = {
      ...baseMsgCreateCollectionResponse,
    } as MsgCreateCollectionResponse;
    return message;
  },

  toJSON(_: MsgCreateCollectionResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateCollectionResponse>
  ): MsgCreateCollectionResponse {
    const message = {
      ...baseMsgCreateCollectionResponse,
    } as MsgCreateCollectionResponse;
    return message;
  },
};

const baseMsgAddCardToCollection: object = {
  creator: "",
  collectionId: 0,
  cardId: 0,
};

export const MsgAddCardToCollection = {
  encode(
    message: MsgAddCardToCollection,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.collectionId !== 0) {
      writer.uint32(16).uint64(message.collectionId);
    }
    if (message.cardId !== 0) {
      writer.uint32(24).uint64(message.cardId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddCardToCollection {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddCardToCollection } as MsgAddCardToCollection;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.collectionId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.cardId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddCardToCollection {
    const message = { ...baseMsgAddCardToCollection } as MsgAddCardToCollection;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.collectionId !== undefined && object.collectionId !== null) {
      message.collectionId = Number(object.collectionId);
    } else {
      message.collectionId = 0;
    }
    if (object.cardId !== undefined && object.cardId !== null) {
      message.cardId = Number(object.cardId);
    } else {
      message.cardId = 0;
    }
    return message;
  },

  toJSON(message: MsgAddCardToCollection): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.collectionId !== undefined &&
      (obj.collectionId = message.collectionId);
    message.cardId !== undefined && (obj.cardId = message.cardId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgAddCardToCollection>
  ): MsgAddCardToCollection {
    const message = { ...baseMsgAddCardToCollection } as MsgAddCardToCollection;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.collectionId !== undefined && object.collectionId !== null) {
      message.collectionId = object.collectionId;
    } else {
      message.collectionId = 0;
    }
    if (object.cardId !== undefined && object.cardId !== null) {
      message.cardId = object.cardId;
    } else {
      message.cardId = 0;
    }
    return message;
  },
};

const baseMsgAddCardToCollectionResponse: object = {};

export const MsgAddCardToCollectionResponse = {
  encode(
    _: MsgAddCardToCollectionResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgAddCardToCollectionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgAddCardToCollectionResponse,
    } as MsgAddCardToCollectionResponse;
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

  fromJSON(_: any): MsgAddCardToCollectionResponse {
    const message = {
      ...baseMsgAddCardToCollectionResponse,
    } as MsgAddCardToCollectionResponse;
    return message;
  },

  toJSON(_: MsgAddCardToCollectionResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgAddCardToCollectionResponse>
  ): MsgAddCardToCollectionResponse {
    const message = {
      ...baseMsgAddCardToCollectionResponse,
    } as MsgAddCardToCollectionResponse;
    return message;
  },
};

const baseMsgFinalizeCollection: object = { creator: "", collectionId: 0 };

export const MsgFinalizeCollection = {
  encode(
    message: MsgFinalizeCollection,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.collectionId !== 0) {
      writer.uint32(16).uint64(message.collectionId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgFinalizeCollection {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgFinalizeCollection } as MsgFinalizeCollection;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.collectionId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgFinalizeCollection {
    const message = { ...baseMsgFinalizeCollection } as MsgFinalizeCollection;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.collectionId !== undefined && object.collectionId !== null) {
      message.collectionId = Number(object.collectionId);
    } else {
      message.collectionId = 0;
    }
    return message;
  },

  toJSON(message: MsgFinalizeCollection): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.collectionId !== undefined &&
      (obj.collectionId = message.collectionId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgFinalizeCollection>
  ): MsgFinalizeCollection {
    const message = { ...baseMsgFinalizeCollection } as MsgFinalizeCollection;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.collectionId !== undefined && object.collectionId !== null) {
      message.collectionId = object.collectionId;
    } else {
      message.collectionId = 0;
    }
    return message;
  },
};

const baseMsgFinalizeCollectionResponse: object = {};

export const MsgFinalizeCollectionResponse = {
  encode(
    _: MsgFinalizeCollectionResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgFinalizeCollectionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgFinalizeCollectionResponse,
    } as MsgFinalizeCollectionResponse;
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

  fromJSON(_: any): MsgFinalizeCollectionResponse {
    const message = {
      ...baseMsgFinalizeCollectionResponse,
    } as MsgFinalizeCollectionResponse;
    return message;
  },

  toJSON(_: MsgFinalizeCollectionResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgFinalizeCollectionResponse>
  ): MsgFinalizeCollectionResponse {
    const message = {
      ...baseMsgFinalizeCollectionResponse,
    } as MsgFinalizeCollectionResponse;
    return message;
  },
};

const baseMsgBuyCollection: object = { creator: "", collectionId: 0 };

export const MsgBuyCollection = {
  encode(message: MsgBuyCollection, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.collectionId !== 0) {
      writer.uint32(16).uint64(message.collectionId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgBuyCollection {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgBuyCollection } as MsgBuyCollection;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.collectionId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgBuyCollection {
    const message = { ...baseMsgBuyCollection } as MsgBuyCollection;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.collectionId !== undefined && object.collectionId !== null) {
      message.collectionId = Number(object.collectionId);
    } else {
      message.collectionId = 0;
    }
    return message;
  },

  toJSON(message: MsgBuyCollection): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.collectionId !== undefined &&
      (obj.collectionId = message.collectionId);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgBuyCollection>): MsgBuyCollection {
    const message = { ...baseMsgBuyCollection } as MsgBuyCollection;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.collectionId !== undefined && object.collectionId !== null) {
      message.collectionId = object.collectionId;
    } else {
      message.collectionId = 0;
    }
    return message;
  },
};

const baseMsgBuyCollectionResponse: object = {};

export const MsgBuyCollectionResponse = {
  encode(
    _: MsgBuyCollectionResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgBuyCollectionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgBuyCollectionResponse,
    } as MsgBuyCollectionResponse;
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

  fromJSON(_: any): MsgBuyCollectionResponse {
    const message = {
      ...baseMsgBuyCollectionResponse,
    } as MsgBuyCollectionResponse;
    return message;
  },

  toJSON(_: MsgBuyCollectionResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgBuyCollectionResponse>
  ): MsgBuyCollectionResponse {
    const message = {
      ...baseMsgBuyCollectionResponse,
    } as MsgBuyCollectionResponse;
    return message;
  },
};

const baseMsgRemoveCardFromCollection: object = {
  creator: "",
  collectionId: 0,
  cardId: 0,
};

export const MsgRemoveCardFromCollection = {
  encode(
    message: MsgRemoveCardFromCollection,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.collectionId !== 0) {
      writer.uint32(16).uint64(message.collectionId);
    }
    if (message.cardId !== 0) {
      writer.uint32(24).uint64(message.cardId);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgRemoveCardFromCollection {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgRemoveCardFromCollection,
    } as MsgRemoveCardFromCollection;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.collectionId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.cardId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRemoveCardFromCollection {
    const message = {
      ...baseMsgRemoveCardFromCollection,
    } as MsgRemoveCardFromCollection;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.collectionId !== undefined && object.collectionId !== null) {
      message.collectionId = Number(object.collectionId);
    } else {
      message.collectionId = 0;
    }
    if (object.cardId !== undefined && object.cardId !== null) {
      message.cardId = Number(object.cardId);
    } else {
      message.cardId = 0;
    }
    return message;
  },

  toJSON(message: MsgRemoveCardFromCollection): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.collectionId !== undefined &&
      (obj.collectionId = message.collectionId);
    message.cardId !== undefined && (obj.cardId = message.cardId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgRemoveCardFromCollection>
  ): MsgRemoveCardFromCollection {
    const message = {
      ...baseMsgRemoveCardFromCollection,
    } as MsgRemoveCardFromCollection;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.collectionId !== undefined && object.collectionId !== null) {
      message.collectionId = object.collectionId;
    } else {
      message.collectionId = 0;
    }
    if (object.cardId !== undefined && object.cardId !== null) {
      message.cardId = object.cardId;
    } else {
      message.cardId = 0;
    }
    return message;
  },
};

const baseMsgRemoveCardFromCollectionResponse: object = {};

export const MsgRemoveCardFromCollectionResponse = {
  encode(
    _: MsgRemoveCardFromCollectionResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgRemoveCardFromCollectionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgRemoveCardFromCollectionResponse,
    } as MsgRemoveCardFromCollectionResponse;
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

  fromJSON(_: any): MsgRemoveCardFromCollectionResponse {
    const message = {
      ...baseMsgRemoveCardFromCollectionResponse,
    } as MsgRemoveCardFromCollectionResponse;
    return message;
  },

  toJSON(_: MsgRemoveCardFromCollectionResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgRemoveCardFromCollectionResponse>
  ): MsgRemoveCardFromCollectionResponse {
    const message = {
      ...baseMsgRemoveCardFromCollectionResponse,
    } as MsgRemoveCardFromCollectionResponse;
    return message;
  },
};

const baseMsgRemoveContributorFromCollection: object = {
  creator: "",
  collectionId: 0,
  user: "",
};

export const MsgRemoveContributorFromCollection = {
  encode(
    message: MsgRemoveContributorFromCollection,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.collectionId !== 0) {
      writer.uint32(16).uint64(message.collectionId);
    }
    if (message.user !== "") {
      writer.uint32(26).string(message.user);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgRemoveContributorFromCollection {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgRemoveContributorFromCollection,
    } as MsgRemoveContributorFromCollection;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.collectionId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.user = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRemoveContributorFromCollection {
    const message = {
      ...baseMsgRemoveContributorFromCollection,
    } as MsgRemoveContributorFromCollection;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.collectionId !== undefined && object.collectionId !== null) {
      message.collectionId = Number(object.collectionId);
    } else {
      message.collectionId = 0;
    }
    if (object.user !== undefined && object.user !== null) {
      message.user = String(object.user);
    } else {
      message.user = "";
    }
    return message;
  },

  toJSON(message: MsgRemoveContributorFromCollection): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.collectionId !== undefined &&
      (obj.collectionId = message.collectionId);
    message.user !== undefined && (obj.user = message.user);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgRemoveContributorFromCollection>
  ): MsgRemoveContributorFromCollection {
    const message = {
      ...baseMsgRemoveContributorFromCollection,
    } as MsgRemoveContributorFromCollection;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.collectionId !== undefined && object.collectionId !== null) {
      message.collectionId = object.collectionId;
    } else {
      message.collectionId = 0;
    }
    if (object.user !== undefined && object.user !== null) {
      message.user = object.user;
    } else {
      message.user = "";
    }
    return message;
  },
};

const baseMsgRemoveContributorFromCollectionResponse: object = {};

export const MsgRemoveContributorFromCollectionResponse = {
  encode(
    _: MsgRemoveContributorFromCollectionResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgRemoveContributorFromCollectionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgRemoveContributorFromCollectionResponse,
    } as MsgRemoveContributorFromCollectionResponse;
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

  fromJSON(_: any): MsgRemoveContributorFromCollectionResponse {
    const message = {
      ...baseMsgRemoveContributorFromCollectionResponse,
    } as MsgRemoveContributorFromCollectionResponse;
    return message;
  },

  toJSON(_: MsgRemoveContributorFromCollectionResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgRemoveContributorFromCollectionResponse>
  ): MsgRemoveContributorFromCollectionResponse {
    const message = {
      ...baseMsgRemoveContributorFromCollectionResponse,
    } as MsgRemoveContributorFromCollectionResponse;
    return message;
  },
};

const baseMsgAddContributorToCollection: object = {
  creator: "",
  collectionId: 0,
  user: "",
};

export const MsgAddContributorToCollection = {
  encode(
    message: MsgAddContributorToCollection,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.collectionId !== 0) {
      writer.uint32(16).uint64(message.collectionId);
    }
    if (message.user !== "") {
      writer.uint32(26).string(message.user);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgAddContributorToCollection {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgAddContributorToCollection,
    } as MsgAddContributorToCollection;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.collectionId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.user = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddContributorToCollection {
    const message = {
      ...baseMsgAddContributorToCollection,
    } as MsgAddContributorToCollection;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.collectionId !== undefined && object.collectionId !== null) {
      message.collectionId = Number(object.collectionId);
    } else {
      message.collectionId = 0;
    }
    if (object.user !== undefined && object.user !== null) {
      message.user = String(object.user);
    } else {
      message.user = "";
    }
    return message;
  },

  toJSON(message: MsgAddContributorToCollection): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.collectionId !== undefined &&
      (obj.collectionId = message.collectionId);
    message.user !== undefined && (obj.user = message.user);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgAddContributorToCollection>
  ): MsgAddContributorToCollection {
    const message = {
      ...baseMsgAddContributorToCollection,
    } as MsgAddContributorToCollection;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.collectionId !== undefined && object.collectionId !== null) {
      message.collectionId = object.collectionId;
    } else {
      message.collectionId = 0;
    }
    if (object.user !== undefined && object.user !== null) {
      message.user = object.user;
    } else {
      message.user = "";
    }
    return message;
  },
};

const baseMsgAddContributorToCollectionResponse: object = {};

export const MsgAddContributorToCollectionResponse = {
  encode(
    _: MsgAddContributorToCollectionResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgAddContributorToCollectionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgAddContributorToCollectionResponse,
    } as MsgAddContributorToCollectionResponse;
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

  fromJSON(_: any): MsgAddContributorToCollectionResponse {
    const message = {
      ...baseMsgAddContributorToCollectionResponse,
    } as MsgAddContributorToCollectionResponse;
    return message;
  },

  toJSON(_: MsgAddContributorToCollectionResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgAddContributorToCollectionResponse>
  ): MsgAddContributorToCollectionResponse {
    const message = {
      ...baseMsgAddContributorToCollectionResponse,
    } as MsgAddContributorToCollectionResponse;
    return message;
  },
};

const baseMsgSubmitCollectionProposal: object = {
  creator: "",
  collectionId: 0,
};

export const MsgSubmitCollectionProposal = {
  encode(
    message: MsgSubmitCollectionProposal,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.collectionId !== 0) {
      writer.uint32(16).uint64(message.collectionId);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSubmitCollectionProposal {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSubmitCollectionProposal,
    } as MsgSubmitCollectionProposal;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.collectionId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSubmitCollectionProposal {
    const message = {
      ...baseMsgSubmitCollectionProposal,
    } as MsgSubmitCollectionProposal;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.collectionId !== undefined && object.collectionId !== null) {
      message.collectionId = Number(object.collectionId);
    } else {
      message.collectionId = 0;
    }
    return message;
  },

  toJSON(message: MsgSubmitCollectionProposal): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.collectionId !== undefined &&
      (obj.collectionId = message.collectionId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSubmitCollectionProposal>
  ): MsgSubmitCollectionProposal {
    const message = {
      ...baseMsgSubmitCollectionProposal,
    } as MsgSubmitCollectionProposal;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.collectionId !== undefined && object.collectionId !== null) {
      message.collectionId = object.collectionId;
    } else {
      message.collectionId = 0;
    }
    return message;
  },
};

const baseMsgSubmitCollectionProposalResponse: object = {};

export const MsgSubmitCollectionProposalResponse = {
  encode(
    _: MsgSubmitCollectionProposalResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSubmitCollectionProposalResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSubmitCollectionProposalResponse,
    } as MsgSubmitCollectionProposalResponse;
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

  fromJSON(_: any): MsgSubmitCollectionProposalResponse {
    const message = {
      ...baseMsgSubmitCollectionProposalResponse,
    } as MsgSubmitCollectionProposalResponse;
    return message;
  },

  toJSON(_: MsgSubmitCollectionProposalResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgSubmitCollectionProposalResponse>
  ): MsgSubmitCollectionProposalResponse {
    const message = {
      ...baseMsgSubmitCollectionProposalResponse,
    } as MsgSubmitCollectionProposalResponse;
    return message;
  },
};

const baseMsgCreateSellOffer: object = { creator: "", card: 0, price: 0 };

export const MsgCreateSellOffer = {
  encode(
    message: MsgCreateSellOffer,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.card !== 0) {
      writer.uint32(16).uint64(message.card);
    }
    if (message.price !== 0) {
      writer.uint32(24).uint64(message.price);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateSellOffer {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateSellOffer } as MsgCreateSellOffer;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.card = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.price = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateSellOffer {
    const message = { ...baseMsgCreateSellOffer } as MsgCreateSellOffer;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.card !== undefined && object.card !== null) {
      message.card = Number(object.card);
    } else {
      message.card = 0;
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = Number(object.price);
    } else {
      message.price = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateSellOffer): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.card !== undefined && (obj.card = message.card);
    message.price !== undefined && (obj.price = message.price);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateSellOffer>): MsgCreateSellOffer {
    const message = { ...baseMsgCreateSellOffer } as MsgCreateSellOffer;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.card !== undefined && object.card !== null) {
      message.card = object.card;
    } else {
      message.card = 0;
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = object.price;
    } else {
      message.price = 0;
    }
    return message;
  },
};

const baseMsgCreateSellOfferResponse: object = {};

export const MsgCreateSellOfferResponse = {
  encode(
    _: MsgCreateSellOfferResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateSellOfferResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateSellOfferResponse,
    } as MsgCreateSellOfferResponse;
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

  fromJSON(_: any): MsgCreateSellOfferResponse {
    const message = {
      ...baseMsgCreateSellOfferResponse,
    } as MsgCreateSellOfferResponse;
    return message;
  },

  toJSON(_: MsgCreateSellOfferResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateSellOfferResponse>
  ): MsgCreateSellOfferResponse {
    const message = {
      ...baseMsgCreateSellOfferResponse,
    } as MsgCreateSellOfferResponse;
    return message;
  },
};

const baseMsgBuyCard: object = { creator: "", sellOfferId: 0 };

export const MsgBuyCard = {
  encode(message: MsgBuyCard, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.sellOfferId !== 0) {
      writer.uint32(16).uint64(message.sellOfferId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgBuyCard {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgBuyCard } as MsgBuyCard;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.sellOfferId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgBuyCard {
    const message = { ...baseMsgBuyCard } as MsgBuyCard;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.sellOfferId !== undefined && object.sellOfferId !== null) {
      message.sellOfferId = Number(object.sellOfferId);
    } else {
      message.sellOfferId = 0;
    }
    return message;
  },

  toJSON(message: MsgBuyCard): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.sellOfferId !== undefined &&
      (obj.sellOfferId = message.sellOfferId);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgBuyCard>): MsgBuyCard {
    const message = { ...baseMsgBuyCard } as MsgBuyCard;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.sellOfferId !== undefined && object.sellOfferId !== null) {
      message.sellOfferId = object.sellOfferId;
    } else {
      message.sellOfferId = 0;
    }
    return message;
  },
};

const baseMsgBuyCardResponse: object = {};

export const MsgBuyCardResponse = {
  encode(_: MsgBuyCardResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgBuyCardResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgBuyCardResponse } as MsgBuyCardResponse;
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

  fromJSON(_: any): MsgBuyCardResponse {
    const message = { ...baseMsgBuyCardResponse } as MsgBuyCardResponse;
    return message;
  },

  toJSON(_: MsgBuyCardResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgBuyCardResponse>): MsgBuyCardResponse {
    const message = { ...baseMsgBuyCardResponse } as MsgBuyCardResponse;
    return message;
  },
};

const baseMsgRemoveSellOffer: object = { creator: "", sellOfferId: 0 };

export const MsgRemoveSellOffer = {
  encode(
    message: MsgRemoveSellOffer,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.sellOfferId !== 0) {
      writer.uint32(16).uint64(message.sellOfferId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRemoveSellOffer {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRemoveSellOffer } as MsgRemoveSellOffer;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.sellOfferId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRemoveSellOffer {
    const message = { ...baseMsgRemoveSellOffer } as MsgRemoveSellOffer;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.sellOfferId !== undefined && object.sellOfferId !== null) {
      message.sellOfferId = Number(object.sellOfferId);
    } else {
      message.sellOfferId = 0;
    }
    return message;
  },

  toJSON(message: MsgRemoveSellOffer): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.sellOfferId !== undefined &&
      (obj.sellOfferId = message.sellOfferId);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgRemoveSellOffer>): MsgRemoveSellOffer {
    const message = { ...baseMsgRemoveSellOffer } as MsgRemoveSellOffer;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.sellOfferId !== undefined && object.sellOfferId !== null) {
      message.sellOfferId = object.sellOfferId;
    } else {
      message.sellOfferId = 0;
    }
    return message;
  },
};

const baseMsgRemoveSellOfferResponse: object = {};

export const MsgRemoveSellOfferResponse = {
  encode(
    _: MsgRemoveSellOfferResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgRemoveSellOfferResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgRemoveSellOfferResponse,
    } as MsgRemoveSellOfferResponse;
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

  fromJSON(_: any): MsgRemoveSellOfferResponse {
    const message = {
      ...baseMsgRemoveSellOfferResponse,
    } as MsgRemoveSellOfferResponse;
    return message;
  },

  toJSON(_: MsgRemoveSellOfferResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgRemoveSellOfferResponse>
  ): MsgRemoveSellOfferResponse {
    const message = {
      ...baseMsgRemoveSellOfferResponse,
    } as MsgRemoveSellOfferResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  Createuser(request: MsgCreateuser): Promise<MsgCreateuserResponse>;
  BuyCardScheme(request: MsgBuyCardScheme): Promise<MsgBuyCardSchemeResponse>;
  VoteCard(request: MsgVoteCard): Promise<MsgVoteCardResponse>;
  SaveCardContent(
    request: MsgSaveCardContent
  ): Promise<MsgSaveCardContentResponse>;
  TransferCard(request: MsgTransferCard): Promise<MsgTransferCardResponse>;
  DonateToCard(request: MsgDonateToCard): Promise<MsgDonateToCardResponse>;
  AddArtwork(request: MsgAddArtwork): Promise<MsgAddArtworkResponse>;
  SubmitCopyrightProposal(
    request: MsgSubmitCopyrightProposal
  ): Promise<MsgSubmitCopyrightProposalResponse>;
  ChangeArtist(request: MsgChangeArtist): Promise<MsgChangeArtistResponse>;
  RegisterForCouncil(
    request: MsgRegisterForCouncil
  ): Promise<MsgRegisterForCouncilResponse>;
  ReportMatch(request: MsgReportMatch): Promise<MsgReportMatchResponse>;
  SubmitMatchReporterProposal(
    request: MsgSubmitMatchReporterProposal
  ): Promise<MsgSubmitMatchReporterProposalResponse>;
  ApointMatchReporter(
    request: MsgApointMatchReporter
  ): Promise<MsgApointMatchReporterResponse>;
  CreateCollection(
    request: MsgCreateCollection
  ): Promise<MsgCreateCollectionResponse>;
  AddCardToCollection(
    request: MsgAddCardToCollection
  ): Promise<MsgAddCardToCollectionResponse>;
  FinalizeCollection(
    request: MsgFinalizeCollection
  ): Promise<MsgFinalizeCollectionResponse>;
  BuyCollection(request: MsgBuyCollection): Promise<MsgBuyCollectionResponse>;
  RemoveCardFromCollection(
    request: MsgRemoveCardFromCollection
  ): Promise<MsgRemoveCardFromCollectionResponse>;
  RemoveContributorFromCollection(
    request: MsgRemoveContributorFromCollection
  ): Promise<MsgRemoveContributorFromCollectionResponse>;
  AddContributorToCollection(
    request: MsgAddContributorToCollection
  ): Promise<MsgAddContributorToCollectionResponse>;
  SubmitCollectionProposal(
    request: MsgSubmitCollectionProposal
  ): Promise<MsgSubmitCollectionProposalResponse>;
  CreateSellOffer(
    request: MsgCreateSellOffer
  ): Promise<MsgCreateSellOfferResponse>;
  BuyCard(request: MsgBuyCard): Promise<MsgBuyCardResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  RemoveSellOffer(
    request: MsgRemoveSellOffer
  ): Promise<MsgRemoveSellOfferResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Createuser(request: MsgCreateuser): Promise<MsgCreateuserResponse> {
    const data = MsgCreateuser.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Msg",
      "Createuser",
      data
    );
    return promise.then((data) =>
      MsgCreateuserResponse.decode(new Reader(data))
    );
  }

  BuyCardScheme(request: MsgBuyCardScheme): Promise<MsgBuyCardSchemeResponse> {
    const data = MsgBuyCardScheme.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Msg",
      "BuyCardScheme",
      data
    );
    return promise.then((data) =>
      MsgBuyCardSchemeResponse.decode(new Reader(data))
    );
  }

  VoteCard(request: MsgVoteCard): Promise<MsgVoteCardResponse> {
    const data = MsgVoteCard.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Msg",
      "VoteCard",
      data
    );
    return promise.then((data) => MsgVoteCardResponse.decode(new Reader(data)));
  }

  SaveCardContent(
    request: MsgSaveCardContent
  ): Promise<MsgSaveCardContentResponse> {
    const data = MsgSaveCardContent.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Msg",
      "SaveCardContent",
      data
    );
    return promise.then((data) =>
      MsgSaveCardContentResponse.decode(new Reader(data))
    );
  }

  TransferCard(request: MsgTransferCard): Promise<MsgTransferCardResponse> {
    const data = MsgTransferCard.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Msg",
      "TransferCard",
      data
    );
    return promise.then((data) =>
      MsgTransferCardResponse.decode(new Reader(data))
    );
  }

  DonateToCard(request: MsgDonateToCard): Promise<MsgDonateToCardResponse> {
    const data = MsgDonateToCard.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Msg",
      "DonateToCard",
      data
    );
    return promise.then((data) =>
      MsgDonateToCardResponse.decode(new Reader(data))
    );
  }

  AddArtwork(request: MsgAddArtwork): Promise<MsgAddArtworkResponse> {
    const data = MsgAddArtwork.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Msg",
      "AddArtwork",
      data
    );
    return promise.then((data) =>
      MsgAddArtworkResponse.decode(new Reader(data))
    );
  }

  SubmitCopyrightProposal(
    request: MsgSubmitCopyrightProposal
  ): Promise<MsgSubmitCopyrightProposalResponse> {
    const data = MsgSubmitCopyrightProposal.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Msg",
      "SubmitCopyrightProposal",
      data
    );
    return promise.then((data) =>
      MsgSubmitCopyrightProposalResponse.decode(new Reader(data))
    );
  }

  ChangeArtist(request: MsgChangeArtist): Promise<MsgChangeArtistResponse> {
    const data = MsgChangeArtist.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Msg",
      "ChangeArtist",
      data
    );
    return promise.then((data) =>
      MsgChangeArtistResponse.decode(new Reader(data))
    );
  }

  RegisterForCouncil(
    request: MsgRegisterForCouncil
  ): Promise<MsgRegisterForCouncilResponse> {
    const data = MsgRegisterForCouncil.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Msg",
      "RegisterForCouncil",
      data
    );
    return promise.then((data) =>
      MsgRegisterForCouncilResponse.decode(new Reader(data))
    );
  }

  ReportMatch(request: MsgReportMatch): Promise<MsgReportMatchResponse> {
    const data = MsgReportMatch.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Msg",
      "ReportMatch",
      data
    );
    return promise.then((data) =>
      MsgReportMatchResponse.decode(new Reader(data))
    );
  }

  SubmitMatchReporterProposal(
    request: MsgSubmitMatchReporterProposal
  ): Promise<MsgSubmitMatchReporterProposalResponse> {
    const data = MsgSubmitMatchReporterProposal.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Msg",
      "SubmitMatchReporterProposal",
      data
    );
    return promise.then((data) =>
      MsgSubmitMatchReporterProposalResponse.decode(new Reader(data))
    );
  }

  ApointMatchReporter(
    request: MsgApointMatchReporter
  ): Promise<MsgApointMatchReporterResponse> {
    const data = MsgApointMatchReporter.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Msg",
      "ApointMatchReporter",
      data
    );
    return promise.then((data) =>
      MsgApointMatchReporterResponse.decode(new Reader(data))
    );
  }

  CreateCollection(
    request: MsgCreateCollection
  ): Promise<MsgCreateCollectionResponse> {
    const data = MsgCreateCollection.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Msg",
      "CreateCollection",
      data
    );
    return promise.then((data) =>
      MsgCreateCollectionResponse.decode(new Reader(data))
    );
  }

  AddCardToCollection(
    request: MsgAddCardToCollection
  ): Promise<MsgAddCardToCollectionResponse> {
    const data = MsgAddCardToCollection.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Msg",
      "AddCardToCollection",
      data
    );
    return promise.then((data) =>
      MsgAddCardToCollectionResponse.decode(new Reader(data))
    );
  }

  FinalizeCollection(
    request: MsgFinalizeCollection
  ): Promise<MsgFinalizeCollectionResponse> {
    const data = MsgFinalizeCollection.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Msg",
      "FinalizeCollection",
      data
    );
    return promise.then((data) =>
      MsgFinalizeCollectionResponse.decode(new Reader(data))
    );
  }

  BuyCollection(request: MsgBuyCollection): Promise<MsgBuyCollectionResponse> {
    const data = MsgBuyCollection.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Msg",
      "BuyCollection",
      data
    );
    return promise.then((data) =>
      MsgBuyCollectionResponse.decode(new Reader(data))
    );
  }

  RemoveCardFromCollection(
    request: MsgRemoveCardFromCollection
  ): Promise<MsgRemoveCardFromCollectionResponse> {
    const data = MsgRemoveCardFromCollection.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Msg",
      "RemoveCardFromCollection",
      data
    );
    return promise.then((data) =>
      MsgRemoveCardFromCollectionResponse.decode(new Reader(data))
    );
  }

  RemoveContributorFromCollection(
    request: MsgRemoveContributorFromCollection
  ): Promise<MsgRemoveContributorFromCollectionResponse> {
    const data = MsgRemoveContributorFromCollection.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Msg",
      "RemoveContributorFromCollection",
      data
    );
    return promise.then((data) =>
      MsgRemoveContributorFromCollectionResponse.decode(new Reader(data))
    );
  }

  AddContributorToCollection(
    request: MsgAddContributorToCollection
  ): Promise<MsgAddContributorToCollectionResponse> {
    const data = MsgAddContributorToCollection.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Msg",
      "AddContributorToCollection",
      data
    );
    return promise.then((data) =>
      MsgAddContributorToCollectionResponse.decode(new Reader(data))
    );
  }

  SubmitCollectionProposal(
    request: MsgSubmitCollectionProposal
  ): Promise<MsgSubmitCollectionProposalResponse> {
    const data = MsgSubmitCollectionProposal.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Msg",
      "SubmitCollectionProposal",
      data
    );
    return promise.then((data) =>
      MsgSubmitCollectionProposalResponse.decode(new Reader(data))
    );
  }

  CreateSellOffer(
    request: MsgCreateSellOffer
  ): Promise<MsgCreateSellOfferResponse> {
    const data = MsgCreateSellOffer.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Msg",
      "CreateSellOffer",
      data
    );
    return promise.then((data) =>
      MsgCreateSellOfferResponse.decode(new Reader(data))
    );
  }

  BuyCard(request: MsgBuyCard): Promise<MsgBuyCardResponse> {
    const data = MsgBuyCard.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Msg",
      "BuyCard",
      data
    );
    return promise.then((data) => MsgBuyCardResponse.decode(new Reader(data)));
  }

  RemoveSellOffer(
    request: MsgRemoveSellOffer
  ): Promise<MsgRemoveSellOfferResponse> {
    const data = MsgRemoveSellOffer.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Msg",
      "RemoveSellOffer",
      data
    );
    return promise.then((data) =>
      MsgRemoveSellOfferResponse.decode(new Reader(data))
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
