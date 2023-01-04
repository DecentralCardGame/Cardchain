/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Response, responseFromJSON, responseToJSON } from "./council";

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
    case Outcome.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface MsgCreateuser {
  creator: string;
  newUser: string;
  alias: string;
}

export interface MsgCreateuserResponse {
}

export interface MsgBuyCardScheme {
  creator: string;
  /** cosmos.base.v1beta1.Coin bid = 2; */
  bid: string;
}

export interface MsgBuyCardSchemeResponse {
}

export interface MsgVoteCard {
  creator: string;
  cardId: number;
  voteType: string;
}

export interface MsgVoteCardResponse {
  airdropClaimed: boolean;
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
  artist: string;
}

export interface MsgSaveCardContentResponse {
  airdropClaimed: boolean;
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
  artist: string;
  storyWriter: string;
  contributors: string[];
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

export interface MsgBuyCollection {
  creator: string;
  collectionId: number;
}

export interface MsgBuyCollectionResponse {
  airdropClaimed: boolean;
}

export interface MsgRemoveCardFromCollection {
  creator: string;
  collectionId: number;
  cardId: number;
}

export interface MsgRemoveCardFromCollectionResponse {
}

export interface MsgRemoveContributorFromCollection {
  creator: string;
  collectionId: number;
  user: string;
}

export interface MsgRemoveContributorFromCollectionResponse {
}

export interface MsgAddContributorToCollection {
  creator: string;
  collectionId: number;
  user: string;
}

export interface MsgAddContributorToCollectionResponse {
}

export interface MsgSubmitCollectionProposal {
  creator: string;
  collectionId: number;
}

export interface MsgSubmitCollectionProposalResponse {
}

export interface MsgCreateSellOffer {
  creator: string;
  card: number;
  price: string;
}

export interface MsgCreateSellOfferResponse {
}

export interface MsgBuyCard {
  creator: string;
  sellOfferId: number;
}

export interface MsgBuyCardResponse {
}

export interface MsgRemoveSellOffer {
  creator: string;
  sellOfferId: number;
}

export interface MsgRemoveSellOfferResponse {
}

export interface MsgAddArtworkToCollection {
  creator: string;
  collectionId: number;
  image: Uint8Array;
}

export interface MsgAddArtworkToCollectionResponse {
}

export interface MsgAddStoryToCollection {
  creator: string;
  collectionId: number;
  story: string;
}

export interface MsgAddStoryToCollectionResponse {
}

export interface MsgSetCardRarity {
  creator: string;
  cardId: number;
  collectionId: number;
  rarity: string;
}

export interface MsgSetCardRarityResponse {
}

export interface MsgCreateCouncil {
  creator: string;
  cardId: number;
}

export interface MsgCreateCouncilResponse {
}

export interface MsgCommitCouncilResponse {
  creator: string;
  response: string;
  councilId: number;
  suggestion: string;
}

export interface MsgCommitCouncilResponseResponse {
}

export interface MsgRevealCouncilResponse {
  creator: string;
  response: Response;
  secret: string;
  councilId: number;
}

export interface MsgRevealCouncilResponseResponse {
}

export interface MsgRestartCouncil {
  creator: string;
  councilId: number;
}

export interface MsgRestartCouncilResponse {
}

export interface MsgRewokeCouncilRegistration {
  creator: string;
}

export interface MsgRewokeCouncilRegistrationResponse {
}

export interface MsgConfirmMatch {
  creator: string;
  matchId: number;
  outcome: Outcome;
}

export interface MsgConfirmMatchResponse {
}

export interface MsgSetProfileCard {
  creator: string;
  cardId: number;
}

export interface MsgSetProfileCardResponse {
}

export interface MsgOpenBoosterPack {
  creator: string;
  boosterPackId: number;
}

export interface MsgOpenBoosterPackResponse {
}

export interface MsgTransferBoosterPack {
  creator: string;
  boosterPackId: number;
  receiver: string;
}

export interface MsgTransferBoosterPackResponse {
}

export interface MsgSetCollectionStoryWriter {
  creator: string;
  collectionId: number;
  storyWriter: string;
}

export interface MsgSetCollectionStoryWriterResponse {
}

export interface MsgSetCollectionArtist {
  creator: string;
  collectionId: number;
  artist: string;
}

export interface MsgSetCollectionArtistResponse {
}

export interface MsgSetUserWebsite {
  creator: string;
  website: string;
}

export interface MsgSetUserWebsiteResponse {
}

export interface MsgSetUserBiography {
  creator: string;
  biography: string;
}

export interface MsgSetUserBiographyResponse {
}

function createBaseMsgCreateuser(): MsgCreateuser {
  return { creator: "", newUser: "", alias: "" };
}

export const MsgCreateuser = {
  encode(message: MsgCreateuser, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateuser {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateuser();
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
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      newUser: isSet(object.newUser) ? String(object.newUser) : "",
      alias: isSet(object.alias) ? String(object.alias) : "",
    };
  },

  toJSON(message: MsgCreateuser): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.newUser !== undefined && (obj.newUser = message.newUser);
    message.alias !== undefined && (obj.alias = message.alias);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateuser>, I>>(object: I): MsgCreateuser {
    const message = createBaseMsgCreateuser();
    message.creator = object.creator ?? "";
    message.newUser = object.newUser ?? "";
    message.alias = object.alias ?? "";
    return message;
  },
};

function createBaseMsgCreateuserResponse(): MsgCreateuserResponse {
  return {};
}

export const MsgCreateuserResponse = {
  encode(_: MsgCreateuserResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateuserResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateuserResponse();
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
    return {};
  },

  toJSON(_: MsgCreateuserResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateuserResponse>, I>>(_: I): MsgCreateuserResponse {
    const message = createBaseMsgCreateuserResponse();
    return message;
  },
};

function createBaseMsgBuyCardScheme(): MsgBuyCardScheme {
  return { creator: "", bid: "" };
}

export const MsgBuyCardScheme = {
  encode(message: MsgBuyCardScheme, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.bid !== "") {
      writer.uint32(18).string(message.bid);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgBuyCardScheme {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgBuyCardScheme();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.bid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgBuyCardScheme {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      bid: isSet(object.bid) ? String(object.bid) : "",
    };
  },

  toJSON(message: MsgBuyCardScheme): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.bid !== undefined && (obj.bid = message.bid);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgBuyCardScheme>, I>>(object: I): MsgBuyCardScheme {
    const message = createBaseMsgBuyCardScheme();
    message.creator = object.creator ?? "";
    message.bid = object.bid ?? "";
    return message;
  },
};

function createBaseMsgBuyCardSchemeResponse(): MsgBuyCardSchemeResponse {
  return {};
}

export const MsgBuyCardSchemeResponse = {
  encode(_: MsgBuyCardSchemeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgBuyCardSchemeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgBuyCardSchemeResponse();
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
    return {};
  },

  toJSON(_: MsgBuyCardSchemeResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgBuyCardSchemeResponse>, I>>(_: I): MsgBuyCardSchemeResponse {
    const message = createBaseMsgBuyCardSchemeResponse();
    return message;
  },
};

function createBaseMsgVoteCard(): MsgVoteCard {
  return { creator: "", cardId: 0, voteType: "" };
}

export const MsgVoteCard = {
  encode(message: MsgVoteCard, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.cardId !== 0) {
      writer.uint32(16).uint64(message.cardId);
    }
    if (message.voteType !== "") {
      writer.uint32(26).string(message.voteType);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgVoteCard {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgVoteCard();
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
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgVoteCard {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      cardId: isSet(object.cardId) ? Number(object.cardId) : 0,
      voteType: isSet(object.voteType) ? String(object.voteType) : "",
    };
  },

  toJSON(message: MsgVoteCard): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.cardId !== undefined && (obj.cardId = Math.round(message.cardId));
    message.voteType !== undefined && (obj.voteType = message.voteType);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgVoteCard>, I>>(object: I): MsgVoteCard {
    const message = createBaseMsgVoteCard();
    message.creator = object.creator ?? "";
    message.cardId = object.cardId ?? 0;
    message.voteType = object.voteType ?? "";
    return message;
  },
};

function createBaseMsgVoteCardResponse(): MsgVoteCardResponse {
  return { airdropClaimed: false };
}

export const MsgVoteCardResponse = {
  encode(message: MsgVoteCardResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.airdropClaimed === true) {
      writer.uint32(8).bool(message.airdropClaimed);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgVoteCardResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgVoteCardResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.airdropClaimed = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgVoteCardResponse {
    return { airdropClaimed: isSet(object.airdropClaimed) ? Boolean(object.airdropClaimed) : false };
  },

  toJSON(message: MsgVoteCardResponse): unknown {
    const obj: any = {};
    message.airdropClaimed !== undefined && (obj.airdropClaimed = message.airdropClaimed);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgVoteCardResponse>, I>>(object: I): MsgVoteCardResponse {
    const message = createBaseMsgVoteCardResponse();
    message.airdropClaimed = object.airdropClaimed ?? false;
    return message;
  },
};

function createBaseMsgSaveCardContent(): MsgSaveCardContent {
  return { creator: "", cardId: 0, content: new Uint8Array(), notes: "", artist: "" };
}

export const MsgSaveCardContent = {
  encode(message: MsgSaveCardContent, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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
    if (message.artist !== "") {
      writer.uint32(42).string(message.artist);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSaveCardContent {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSaveCardContent();
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
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      cardId: isSet(object.cardId) ? Number(object.cardId) : 0,
      content: isSet(object.content) ? bytesFromBase64(object.content) : new Uint8Array(),
      notes: isSet(object.notes) ? String(object.notes) : "",
      artist: isSet(object.artist) ? String(object.artist) : "",
    };
  },

  toJSON(message: MsgSaveCardContent): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.cardId !== undefined && (obj.cardId = Math.round(message.cardId));
    message.content !== undefined
      && (obj.content = base64FromBytes(message.content !== undefined ? message.content : new Uint8Array()));
    message.notes !== undefined && (obj.notes = message.notes);
    message.artist !== undefined && (obj.artist = message.artist);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSaveCardContent>, I>>(object: I): MsgSaveCardContent {
    const message = createBaseMsgSaveCardContent();
    message.creator = object.creator ?? "";
    message.cardId = object.cardId ?? 0;
    message.content = object.content ?? new Uint8Array();
    message.notes = object.notes ?? "";
    message.artist = object.artist ?? "";
    return message;
  },
};

function createBaseMsgSaveCardContentResponse(): MsgSaveCardContentResponse {
  return { airdropClaimed: false };
}

export const MsgSaveCardContentResponse = {
  encode(message: MsgSaveCardContentResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.airdropClaimed === true) {
      writer.uint32(8).bool(message.airdropClaimed);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSaveCardContentResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSaveCardContentResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.airdropClaimed = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSaveCardContentResponse {
    return { airdropClaimed: isSet(object.airdropClaimed) ? Boolean(object.airdropClaimed) : false };
  },

  toJSON(message: MsgSaveCardContentResponse): unknown {
    const obj: any = {};
    message.airdropClaimed !== undefined && (obj.airdropClaimed = message.airdropClaimed);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSaveCardContentResponse>, I>>(object: I): MsgSaveCardContentResponse {
    const message = createBaseMsgSaveCardContentResponse();
    message.airdropClaimed = object.airdropClaimed ?? false;
    return message;
  },
};

function createBaseMsgTransferCard(): MsgTransferCard {
  return { creator: "", cardId: 0, receiver: "" };
}

export const MsgTransferCard = {
  encode(message: MsgTransferCard, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgTransferCard {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgTransferCard();
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
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      cardId: isSet(object.cardId) ? Number(object.cardId) : 0,
      receiver: isSet(object.receiver) ? String(object.receiver) : "",
    };
  },

  toJSON(message: MsgTransferCard): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.cardId !== undefined && (obj.cardId = Math.round(message.cardId));
    message.receiver !== undefined && (obj.receiver = message.receiver);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgTransferCard>, I>>(object: I): MsgTransferCard {
    const message = createBaseMsgTransferCard();
    message.creator = object.creator ?? "";
    message.cardId = object.cardId ?? 0;
    message.receiver = object.receiver ?? "";
    return message;
  },
};

function createBaseMsgTransferCardResponse(): MsgTransferCardResponse {
  return {};
}

export const MsgTransferCardResponse = {
  encode(_: MsgTransferCardResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgTransferCardResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgTransferCardResponse();
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
    return {};
  },

  toJSON(_: MsgTransferCardResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgTransferCardResponse>, I>>(_: I): MsgTransferCardResponse {
    const message = createBaseMsgTransferCardResponse();
    return message;
  },
};

function createBaseMsgDonateToCard(): MsgDonateToCard {
  return { creator: "", cardId: 0, amount: "" };
}

export const MsgDonateToCard = {
  encode(message: MsgDonateToCard, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.cardId !== 0) {
      writer.uint32(16).uint64(message.cardId);
    }
    if (message.amount !== "") {
      writer.uint32(26).string(message.amount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDonateToCard {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDonateToCard();
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
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      cardId: isSet(object.cardId) ? Number(object.cardId) : 0,
      amount: isSet(object.amount) ? String(object.amount) : "",
    };
  },

  toJSON(message: MsgDonateToCard): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.cardId !== undefined && (obj.cardId = Math.round(message.cardId));
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDonateToCard>, I>>(object: I): MsgDonateToCard {
    const message = createBaseMsgDonateToCard();
    message.creator = object.creator ?? "";
    message.cardId = object.cardId ?? 0;
    message.amount = object.amount ?? "";
    return message;
  },
};

function createBaseMsgDonateToCardResponse(): MsgDonateToCardResponse {
  return {};
}

export const MsgDonateToCardResponse = {
  encode(_: MsgDonateToCardResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDonateToCardResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDonateToCardResponse();
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
    return {};
  },

  toJSON(_: MsgDonateToCardResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDonateToCardResponse>, I>>(_: I): MsgDonateToCardResponse {
    const message = createBaseMsgDonateToCardResponse();
    return message;
  },
};

function createBaseMsgAddArtwork(): MsgAddArtwork {
  return { creator: "", cardId: 0, image: new Uint8Array(), fullArt: false };
}

export const MsgAddArtwork = {
  encode(message: MsgAddArtwork, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddArtwork {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddArtwork();
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
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      cardId: isSet(object.cardId) ? Number(object.cardId) : 0,
      image: isSet(object.image) ? bytesFromBase64(object.image) : new Uint8Array(),
      fullArt: isSet(object.fullArt) ? Boolean(object.fullArt) : false,
    };
  },

  toJSON(message: MsgAddArtwork): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.cardId !== undefined && (obj.cardId = Math.round(message.cardId));
    message.image !== undefined
      && (obj.image = base64FromBytes(message.image !== undefined ? message.image : new Uint8Array()));
    message.fullArt !== undefined && (obj.fullArt = message.fullArt);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddArtwork>, I>>(object: I): MsgAddArtwork {
    const message = createBaseMsgAddArtwork();
    message.creator = object.creator ?? "";
    message.cardId = object.cardId ?? 0;
    message.image = object.image ?? new Uint8Array();
    message.fullArt = object.fullArt ?? false;
    return message;
  },
};

function createBaseMsgAddArtworkResponse(): MsgAddArtworkResponse {
  return {};
}

export const MsgAddArtworkResponse = {
  encode(_: MsgAddArtworkResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddArtworkResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddArtworkResponse();
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
    return {};
  },

  toJSON(_: MsgAddArtworkResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddArtworkResponse>, I>>(_: I): MsgAddArtworkResponse {
    const message = createBaseMsgAddArtworkResponse();
    return message;
  },
};

function createBaseMsgSubmitCopyrightProposal(): MsgSubmitCopyrightProposal {
  return { creator: "", cardId: 0, description: "", link: "" };
}

export const MsgSubmitCopyrightProposal = {
  encode(message: MsgSubmitCopyrightProposal, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSubmitCopyrightProposal {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSubmitCopyrightProposal();
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
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      cardId: isSet(object.cardId) ? Number(object.cardId) : 0,
      description: isSet(object.description) ? String(object.description) : "",
      link: isSet(object.link) ? String(object.link) : "",
    };
  },

  toJSON(message: MsgSubmitCopyrightProposal): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.cardId !== undefined && (obj.cardId = Math.round(message.cardId));
    message.description !== undefined && (obj.description = message.description);
    message.link !== undefined && (obj.link = message.link);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSubmitCopyrightProposal>, I>>(object: I): MsgSubmitCopyrightProposal {
    const message = createBaseMsgSubmitCopyrightProposal();
    message.creator = object.creator ?? "";
    message.cardId = object.cardId ?? 0;
    message.description = object.description ?? "";
    message.link = object.link ?? "";
    return message;
  },
};

function createBaseMsgSubmitCopyrightProposalResponse(): MsgSubmitCopyrightProposalResponse {
  return {};
}

export const MsgSubmitCopyrightProposalResponse = {
  encode(_: MsgSubmitCopyrightProposalResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSubmitCopyrightProposalResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSubmitCopyrightProposalResponse();
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
    return {};
  },

  toJSON(_: MsgSubmitCopyrightProposalResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSubmitCopyrightProposalResponse>, I>>(
    _: I,
  ): MsgSubmitCopyrightProposalResponse {
    const message = createBaseMsgSubmitCopyrightProposalResponse();
    return message;
  },
};

function createBaseMsgChangeArtist(): MsgChangeArtist {
  return { creator: "", cardID: 0, artist: "" };
}

export const MsgChangeArtist = {
  encode(message: MsgChangeArtist, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgChangeArtist {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgChangeArtist();
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
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      cardID: isSet(object.cardID) ? Number(object.cardID) : 0,
      artist: isSet(object.artist) ? String(object.artist) : "",
    };
  },

  toJSON(message: MsgChangeArtist): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.cardID !== undefined && (obj.cardID = Math.round(message.cardID));
    message.artist !== undefined && (obj.artist = message.artist);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgChangeArtist>, I>>(object: I): MsgChangeArtist {
    const message = createBaseMsgChangeArtist();
    message.creator = object.creator ?? "";
    message.cardID = object.cardID ?? 0;
    message.artist = object.artist ?? "";
    return message;
  },
};

function createBaseMsgChangeArtistResponse(): MsgChangeArtistResponse {
  return {};
}

export const MsgChangeArtistResponse = {
  encode(_: MsgChangeArtistResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgChangeArtistResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgChangeArtistResponse();
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
    return {};
  },

  toJSON(_: MsgChangeArtistResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgChangeArtistResponse>, I>>(_: I): MsgChangeArtistResponse {
    const message = createBaseMsgChangeArtistResponse();
    return message;
  },
};

function createBaseMsgRegisterForCouncil(): MsgRegisterForCouncil {
  return { creator: "" };
}

export const MsgRegisterForCouncil = {
  encode(message: MsgRegisterForCouncil, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRegisterForCouncil {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRegisterForCouncil();
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
    return { creator: isSet(object.creator) ? String(object.creator) : "" };
  },

  toJSON(message: MsgRegisterForCouncil): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRegisterForCouncil>, I>>(object: I): MsgRegisterForCouncil {
    const message = createBaseMsgRegisterForCouncil();
    message.creator = object.creator ?? "";
    return message;
  },
};

function createBaseMsgRegisterForCouncilResponse(): MsgRegisterForCouncilResponse {
  return {};
}

export const MsgRegisterForCouncilResponse = {
  encode(_: MsgRegisterForCouncilResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRegisterForCouncilResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRegisterForCouncilResponse();
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
    return {};
  },

  toJSON(_: MsgRegisterForCouncilResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRegisterForCouncilResponse>, I>>(_: I): MsgRegisterForCouncilResponse {
    const message = createBaseMsgRegisterForCouncilResponse();
    return message;
  },
};

function createBaseMsgReportMatch(): MsgReportMatch {
  return { creator: "", playerA: "", playerB: "", cardsA: [], cardsB: [], outcome: 0 };
}

export const MsgReportMatch = {
  encode(message: MsgReportMatch, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgReportMatch {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgReportMatch();
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
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      playerA: isSet(object.playerA) ? String(object.playerA) : "",
      playerB: isSet(object.playerB) ? String(object.playerB) : "",
      cardsA: Array.isArray(object?.cardsA) ? object.cardsA.map((e: any) => Number(e)) : [],
      cardsB: Array.isArray(object?.cardsB) ? object.cardsB.map((e: any) => Number(e)) : [],
      outcome: isSet(object.outcome) ? outcomeFromJSON(object.outcome) : 0,
    };
  },

  toJSON(message: MsgReportMatch): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.playerA !== undefined && (obj.playerA = message.playerA);
    message.playerB !== undefined && (obj.playerB = message.playerB);
    if (message.cardsA) {
      obj.cardsA = message.cardsA.map((e) => Math.round(e));
    } else {
      obj.cardsA = [];
    }
    if (message.cardsB) {
      obj.cardsB = message.cardsB.map((e) => Math.round(e));
    } else {
      obj.cardsB = [];
    }
    message.outcome !== undefined && (obj.outcome = outcomeToJSON(message.outcome));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgReportMatch>, I>>(object: I): MsgReportMatch {
    const message = createBaseMsgReportMatch();
    message.creator = object.creator ?? "";
    message.playerA = object.playerA ?? "";
    message.playerB = object.playerB ?? "";
    message.cardsA = object.cardsA?.map((e) => e) || [];
    message.cardsB = object.cardsB?.map((e) => e) || [];
    message.outcome = object.outcome ?? 0;
    return message;
  },
};

function createBaseMsgReportMatchResponse(): MsgReportMatchResponse {
  return { matchId: 0 };
}

export const MsgReportMatchResponse = {
  encode(message: MsgReportMatchResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.matchId !== 0) {
      writer.uint32(8).uint64(message.matchId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgReportMatchResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgReportMatchResponse();
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
    return { matchId: isSet(object.matchId) ? Number(object.matchId) : 0 };
  },

  toJSON(message: MsgReportMatchResponse): unknown {
    const obj: any = {};
    message.matchId !== undefined && (obj.matchId = Math.round(message.matchId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgReportMatchResponse>, I>>(object: I): MsgReportMatchResponse {
    const message = createBaseMsgReportMatchResponse();
    message.matchId = object.matchId ?? 0;
    return message;
  },
};

function createBaseMsgSubmitMatchReporterProposal(): MsgSubmitMatchReporterProposal {
  return { creator: "", reporter: "", deposit: "", description: "" };
}

export const MsgSubmitMatchReporterProposal = {
  encode(message: MsgSubmitMatchReporterProposal, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSubmitMatchReporterProposal {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSubmitMatchReporterProposal();
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
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      reporter: isSet(object.reporter) ? String(object.reporter) : "",
      deposit: isSet(object.deposit) ? String(object.deposit) : "",
      description: isSet(object.description) ? String(object.description) : "",
    };
  },

  toJSON(message: MsgSubmitMatchReporterProposal): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.reporter !== undefined && (obj.reporter = message.reporter);
    message.deposit !== undefined && (obj.deposit = message.deposit);
    message.description !== undefined && (obj.description = message.description);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSubmitMatchReporterProposal>, I>>(
    object: I,
  ): MsgSubmitMatchReporterProposal {
    const message = createBaseMsgSubmitMatchReporterProposal();
    message.creator = object.creator ?? "";
    message.reporter = object.reporter ?? "";
    message.deposit = object.deposit ?? "";
    message.description = object.description ?? "";
    return message;
  },
};

function createBaseMsgSubmitMatchReporterProposalResponse(): MsgSubmitMatchReporterProposalResponse {
  return {};
}

export const MsgSubmitMatchReporterProposalResponse = {
  encode(_: MsgSubmitMatchReporterProposalResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSubmitMatchReporterProposalResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSubmitMatchReporterProposalResponse();
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
    return {};
  },

  toJSON(_: MsgSubmitMatchReporterProposalResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSubmitMatchReporterProposalResponse>, I>>(
    _: I,
  ): MsgSubmitMatchReporterProposalResponse {
    const message = createBaseMsgSubmitMatchReporterProposalResponse();
    return message;
  },
};

function createBaseMsgApointMatchReporter(): MsgApointMatchReporter {
  return { creator: "", reporter: "" };
}

export const MsgApointMatchReporter = {
  encode(message: MsgApointMatchReporter, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.reporter !== "") {
      writer.uint32(18).string(message.reporter);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgApointMatchReporter {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgApointMatchReporter();
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
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      reporter: isSet(object.reporter) ? String(object.reporter) : "",
    };
  },

  toJSON(message: MsgApointMatchReporter): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.reporter !== undefined && (obj.reporter = message.reporter);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgApointMatchReporter>, I>>(object: I): MsgApointMatchReporter {
    const message = createBaseMsgApointMatchReporter();
    message.creator = object.creator ?? "";
    message.reporter = object.reporter ?? "";
    return message;
  },
};

function createBaseMsgApointMatchReporterResponse(): MsgApointMatchReporterResponse {
  return {};
}

export const MsgApointMatchReporterResponse = {
  encode(_: MsgApointMatchReporterResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgApointMatchReporterResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgApointMatchReporterResponse();
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
    return {};
  },

  toJSON(_: MsgApointMatchReporterResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgApointMatchReporterResponse>, I>>(_: I): MsgApointMatchReporterResponse {
    const message = createBaseMsgApointMatchReporterResponse();
    return message;
  },
};

function createBaseMsgCreateCollection(): MsgCreateCollection {
  return { creator: "", name: "", artist: "", storyWriter: "", contributors: [] };
}

export const MsgCreateCollection = {
  encode(message: MsgCreateCollection, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateCollection {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateCollection();
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
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      name: isSet(object.name) ? String(object.name) : "",
      artist: isSet(object.artist) ? String(object.artist) : "",
      storyWriter: isSet(object.storyWriter) ? String(object.storyWriter) : "",
      contributors: Array.isArray(object?.contributors) ? object.contributors.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: MsgCreateCollection): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    message.artist !== undefined && (obj.artist = message.artist);
    message.storyWriter !== undefined && (obj.storyWriter = message.storyWriter);
    if (message.contributors) {
      obj.contributors = message.contributors.map((e) => e);
    } else {
      obj.contributors = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateCollection>, I>>(object: I): MsgCreateCollection {
    const message = createBaseMsgCreateCollection();
    message.creator = object.creator ?? "";
    message.name = object.name ?? "";
    message.artist = object.artist ?? "";
    message.storyWriter = object.storyWriter ?? "";
    message.contributors = object.contributors?.map((e) => e) || [];
    return message;
  },
};

function createBaseMsgCreateCollectionResponse(): MsgCreateCollectionResponse {
  return {};
}

export const MsgCreateCollectionResponse = {
  encode(_: MsgCreateCollectionResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateCollectionResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateCollectionResponse();
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
    return {};
  },

  toJSON(_: MsgCreateCollectionResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateCollectionResponse>, I>>(_: I): MsgCreateCollectionResponse {
    const message = createBaseMsgCreateCollectionResponse();
    return message;
  },
};

function createBaseMsgAddCardToCollection(): MsgAddCardToCollection {
  return { creator: "", collectionId: 0, cardId: 0 };
}

export const MsgAddCardToCollection = {
  encode(message: MsgAddCardToCollection, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddCardToCollection {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddCardToCollection();
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
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      collectionId: isSet(object.collectionId) ? Number(object.collectionId) : 0,
      cardId: isSet(object.cardId) ? Number(object.cardId) : 0,
    };
  },

  toJSON(message: MsgAddCardToCollection): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.collectionId !== undefined && (obj.collectionId = Math.round(message.collectionId));
    message.cardId !== undefined && (obj.cardId = Math.round(message.cardId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddCardToCollection>, I>>(object: I): MsgAddCardToCollection {
    const message = createBaseMsgAddCardToCollection();
    message.creator = object.creator ?? "";
    message.collectionId = object.collectionId ?? 0;
    message.cardId = object.cardId ?? 0;
    return message;
  },
};

function createBaseMsgAddCardToCollectionResponse(): MsgAddCardToCollectionResponse {
  return {};
}

export const MsgAddCardToCollectionResponse = {
  encode(_: MsgAddCardToCollectionResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddCardToCollectionResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddCardToCollectionResponse();
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
    return {};
  },

  toJSON(_: MsgAddCardToCollectionResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddCardToCollectionResponse>, I>>(_: I): MsgAddCardToCollectionResponse {
    const message = createBaseMsgAddCardToCollectionResponse();
    return message;
  },
};

function createBaseMsgFinalizeCollection(): MsgFinalizeCollection {
  return { creator: "", collectionId: 0 };
}

export const MsgFinalizeCollection = {
  encode(message: MsgFinalizeCollection, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.collectionId !== 0) {
      writer.uint32(16).uint64(message.collectionId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgFinalizeCollection {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgFinalizeCollection();
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
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      collectionId: isSet(object.collectionId) ? Number(object.collectionId) : 0,
    };
  },

  toJSON(message: MsgFinalizeCollection): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.collectionId !== undefined && (obj.collectionId = Math.round(message.collectionId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgFinalizeCollection>, I>>(object: I): MsgFinalizeCollection {
    const message = createBaseMsgFinalizeCollection();
    message.creator = object.creator ?? "";
    message.collectionId = object.collectionId ?? 0;
    return message;
  },
};

function createBaseMsgFinalizeCollectionResponse(): MsgFinalizeCollectionResponse {
  return {};
}

export const MsgFinalizeCollectionResponse = {
  encode(_: MsgFinalizeCollectionResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgFinalizeCollectionResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgFinalizeCollectionResponse();
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
    return {};
  },

  toJSON(_: MsgFinalizeCollectionResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgFinalizeCollectionResponse>, I>>(_: I): MsgFinalizeCollectionResponse {
    const message = createBaseMsgFinalizeCollectionResponse();
    return message;
  },
};

function createBaseMsgBuyCollection(): MsgBuyCollection {
  return { creator: "", collectionId: 0 };
}

export const MsgBuyCollection = {
  encode(message: MsgBuyCollection, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.collectionId !== 0) {
      writer.uint32(16).uint64(message.collectionId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgBuyCollection {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgBuyCollection();
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
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      collectionId: isSet(object.collectionId) ? Number(object.collectionId) : 0,
    };
  },

  toJSON(message: MsgBuyCollection): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.collectionId !== undefined && (obj.collectionId = Math.round(message.collectionId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgBuyCollection>, I>>(object: I): MsgBuyCollection {
    const message = createBaseMsgBuyCollection();
    message.creator = object.creator ?? "";
    message.collectionId = object.collectionId ?? 0;
    return message;
  },
};

function createBaseMsgBuyCollectionResponse(): MsgBuyCollectionResponse {
  return { airdropClaimed: false };
}

export const MsgBuyCollectionResponse = {
  encode(message: MsgBuyCollectionResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.airdropClaimed === true) {
      writer.uint32(8).bool(message.airdropClaimed);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgBuyCollectionResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgBuyCollectionResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.airdropClaimed = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgBuyCollectionResponse {
    return { airdropClaimed: isSet(object.airdropClaimed) ? Boolean(object.airdropClaimed) : false };
  },

  toJSON(message: MsgBuyCollectionResponse): unknown {
    const obj: any = {};
    message.airdropClaimed !== undefined && (obj.airdropClaimed = message.airdropClaimed);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgBuyCollectionResponse>, I>>(object: I): MsgBuyCollectionResponse {
    const message = createBaseMsgBuyCollectionResponse();
    message.airdropClaimed = object.airdropClaimed ?? false;
    return message;
  },
};

function createBaseMsgRemoveCardFromCollection(): MsgRemoveCardFromCollection {
  return { creator: "", collectionId: 0, cardId: 0 };
}

export const MsgRemoveCardFromCollection = {
  encode(message: MsgRemoveCardFromCollection, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRemoveCardFromCollection {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRemoveCardFromCollection();
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
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      collectionId: isSet(object.collectionId) ? Number(object.collectionId) : 0,
      cardId: isSet(object.cardId) ? Number(object.cardId) : 0,
    };
  },

  toJSON(message: MsgRemoveCardFromCollection): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.collectionId !== undefined && (obj.collectionId = Math.round(message.collectionId));
    message.cardId !== undefined && (obj.cardId = Math.round(message.cardId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRemoveCardFromCollection>, I>>(object: I): MsgRemoveCardFromCollection {
    const message = createBaseMsgRemoveCardFromCollection();
    message.creator = object.creator ?? "";
    message.collectionId = object.collectionId ?? 0;
    message.cardId = object.cardId ?? 0;
    return message;
  },
};

function createBaseMsgRemoveCardFromCollectionResponse(): MsgRemoveCardFromCollectionResponse {
  return {};
}

export const MsgRemoveCardFromCollectionResponse = {
  encode(_: MsgRemoveCardFromCollectionResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRemoveCardFromCollectionResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRemoveCardFromCollectionResponse();
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
    return {};
  },

  toJSON(_: MsgRemoveCardFromCollectionResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRemoveCardFromCollectionResponse>, I>>(
    _: I,
  ): MsgRemoveCardFromCollectionResponse {
    const message = createBaseMsgRemoveCardFromCollectionResponse();
    return message;
  },
};

function createBaseMsgRemoveContributorFromCollection(): MsgRemoveContributorFromCollection {
  return { creator: "", collectionId: 0, user: "" };
}

export const MsgRemoveContributorFromCollection = {
  encode(message: MsgRemoveContributorFromCollection, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRemoveContributorFromCollection {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRemoveContributorFromCollection();
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
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      collectionId: isSet(object.collectionId) ? Number(object.collectionId) : 0,
      user: isSet(object.user) ? String(object.user) : "",
    };
  },

  toJSON(message: MsgRemoveContributorFromCollection): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.collectionId !== undefined && (obj.collectionId = Math.round(message.collectionId));
    message.user !== undefined && (obj.user = message.user);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRemoveContributorFromCollection>, I>>(
    object: I,
  ): MsgRemoveContributorFromCollection {
    const message = createBaseMsgRemoveContributorFromCollection();
    message.creator = object.creator ?? "";
    message.collectionId = object.collectionId ?? 0;
    message.user = object.user ?? "";
    return message;
  },
};

function createBaseMsgRemoveContributorFromCollectionResponse(): MsgRemoveContributorFromCollectionResponse {
  return {};
}

export const MsgRemoveContributorFromCollectionResponse = {
  encode(_: MsgRemoveContributorFromCollectionResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRemoveContributorFromCollectionResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRemoveContributorFromCollectionResponse();
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
    return {};
  },

  toJSON(_: MsgRemoveContributorFromCollectionResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRemoveContributorFromCollectionResponse>, I>>(
    _: I,
  ): MsgRemoveContributorFromCollectionResponse {
    const message = createBaseMsgRemoveContributorFromCollectionResponse();
    return message;
  },
};

function createBaseMsgAddContributorToCollection(): MsgAddContributorToCollection {
  return { creator: "", collectionId: 0, user: "" };
}

export const MsgAddContributorToCollection = {
  encode(message: MsgAddContributorToCollection, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddContributorToCollection {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddContributorToCollection();
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
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      collectionId: isSet(object.collectionId) ? Number(object.collectionId) : 0,
      user: isSet(object.user) ? String(object.user) : "",
    };
  },

  toJSON(message: MsgAddContributorToCollection): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.collectionId !== undefined && (obj.collectionId = Math.round(message.collectionId));
    message.user !== undefined && (obj.user = message.user);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddContributorToCollection>, I>>(
    object: I,
  ): MsgAddContributorToCollection {
    const message = createBaseMsgAddContributorToCollection();
    message.creator = object.creator ?? "";
    message.collectionId = object.collectionId ?? 0;
    message.user = object.user ?? "";
    return message;
  },
};

function createBaseMsgAddContributorToCollectionResponse(): MsgAddContributorToCollectionResponse {
  return {};
}

export const MsgAddContributorToCollectionResponse = {
  encode(_: MsgAddContributorToCollectionResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddContributorToCollectionResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddContributorToCollectionResponse();
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
    return {};
  },

  toJSON(_: MsgAddContributorToCollectionResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddContributorToCollectionResponse>, I>>(
    _: I,
  ): MsgAddContributorToCollectionResponse {
    const message = createBaseMsgAddContributorToCollectionResponse();
    return message;
  },
};

function createBaseMsgSubmitCollectionProposal(): MsgSubmitCollectionProposal {
  return { creator: "", collectionId: 0 };
}

export const MsgSubmitCollectionProposal = {
  encode(message: MsgSubmitCollectionProposal, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.collectionId !== 0) {
      writer.uint32(16).uint64(message.collectionId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSubmitCollectionProposal {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSubmitCollectionProposal();
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
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      collectionId: isSet(object.collectionId) ? Number(object.collectionId) : 0,
    };
  },

  toJSON(message: MsgSubmitCollectionProposal): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.collectionId !== undefined && (obj.collectionId = Math.round(message.collectionId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSubmitCollectionProposal>, I>>(object: I): MsgSubmitCollectionProposal {
    const message = createBaseMsgSubmitCollectionProposal();
    message.creator = object.creator ?? "";
    message.collectionId = object.collectionId ?? 0;
    return message;
  },
};

function createBaseMsgSubmitCollectionProposalResponse(): MsgSubmitCollectionProposalResponse {
  return {};
}

export const MsgSubmitCollectionProposalResponse = {
  encode(_: MsgSubmitCollectionProposalResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSubmitCollectionProposalResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSubmitCollectionProposalResponse();
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
    return {};
  },

  toJSON(_: MsgSubmitCollectionProposalResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSubmitCollectionProposalResponse>, I>>(
    _: I,
  ): MsgSubmitCollectionProposalResponse {
    const message = createBaseMsgSubmitCollectionProposalResponse();
    return message;
  },
};

function createBaseMsgCreateSellOffer(): MsgCreateSellOffer {
  return { creator: "", card: 0, price: "" };
}

export const MsgCreateSellOffer = {
  encode(message: MsgCreateSellOffer, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.card !== 0) {
      writer.uint32(16).uint64(message.card);
    }
    if (message.price !== "") {
      writer.uint32(26).string(message.price);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateSellOffer {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateSellOffer();
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
          message.price = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateSellOffer {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      card: isSet(object.card) ? Number(object.card) : 0,
      price: isSet(object.price) ? String(object.price) : "",
    };
  },

  toJSON(message: MsgCreateSellOffer): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.card !== undefined && (obj.card = Math.round(message.card));
    message.price !== undefined && (obj.price = message.price);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateSellOffer>, I>>(object: I): MsgCreateSellOffer {
    const message = createBaseMsgCreateSellOffer();
    message.creator = object.creator ?? "";
    message.card = object.card ?? 0;
    message.price = object.price ?? "";
    return message;
  },
};

function createBaseMsgCreateSellOfferResponse(): MsgCreateSellOfferResponse {
  return {};
}

export const MsgCreateSellOfferResponse = {
  encode(_: MsgCreateSellOfferResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateSellOfferResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateSellOfferResponse();
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
    return {};
  },

  toJSON(_: MsgCreateSellOfferResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateSellOfferResponse>, I>>(_: I): MsgCreateSellOfferResponse {
    const message = createBaseMsgCreateSellOfferResponse();
    return message;
  },
};

function createBaseMsgBuyCard(): MsgBuyCard {
  return { creator: "", sellOfferId: 0 };
}

export const MsgBuyCard = {
  encode(message: MsgBuyCard, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.sellOfferId !== 0) {
      writer.uint32(16).uint64(message.sellOfferId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgBuyCard {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgBuyCard();
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
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      sellOfferId: isSet(object.sellOfferId) ? Number(object.sellOfferId) : 0,
    };
  },

  toJSON(message: MsgBuyCard): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.sellOfferId !== undefined && (obj.sellOfferId = Math.round(message.sellOfferId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgBuyCard>, I>>(object: I): MsgBuyCard {
    const message = createBaseMsgBuyCard();
    message.creator = object.creator ?? "";
    message.sellOfferId = object.sellOfferId ?? 0;
    return message;
  },
};

function createBaseMsgBuyCardResponse(): MsgBuyCardResponse {
  return {};
}

export const MsgBuyCardResponse = {
  encode(_: MsgBuyCardResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgBuyCardResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgBuyCardResponse();
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
    return {};
  },

  toJSON(_: MsgBuyCardResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgBuyCardResponse>, I>>(_: I): MsgBuyCardResponse {
    const message = createBaseMsgBuyCardResponse();
    return message;
  },
};

function createBaseMsgRemoveSellOffer(): MsgRemoveSellOffer {
  return { creator: "", sellOfferId: 0 };
}

export const MsgRemoveSellOffer = {
  encode(message: MsgRemoveSellOffer, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.sellOfferId !== 0) {
      writer.uint32(16).uint64(message.sellOfferId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRemoveSellOffer {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRemoveSellOffer();
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
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      sellOfferId: isSet(object.sellOfferId) ? Number(object.sellOfferId) : 0,
    };
  },

  toJSON(message: MsgRemoveSellOffer): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.sellOfferId !== undefined && (obj.sellOfferId = Math.round(message.sellOfferId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRemoveSellOffer>, I>>(object: I): MsgRemoveSellOffer {
    const message = createBaseMsgRemoveSellOffer();
    message.creator = object.creator ?? "";
    message.sellOfferId = object.sellOfferId ?? 0;
    return message;
  },
};

function createBaseMsgRemoveSellOfferResponse(): MsgRemoveSellOfferResponse {
  return {};
}

export const MsgRemoveSellOfferResponse = {
  encode(_: MsgRemoveSellOfferResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRemoveSellOfferResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRemoveSellOfferResponse();
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
    return {};
  },

  toJSON(_: MsgRemoveSellOfferResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRemoveSellOfferResponse>, I>>(_: I): MsgRemoveSellOfferResponse {
    const message = createBaseMsgRemoveSellOfferResponse();
    return message;
  },
};

function createBaseMsgAddArtworkToCollection(): MsgAddArtworkToCollection {
  return { creator: "", collectionId: 0, image: new Uint8Array() };
}

export const MsgAddArtworkToCollection = {
  encode(message: MsgAddArtworkToCollection, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.collectionId !== 0) {
      writer.uint32(16).uint64(message.collectionId);
    }
    if (message.image.length !== 0) {
      writer.uint32(26).bytes(message.image);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddArtworkToCollection {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddArtworkToCollection();
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
          message.image = reader.bytes();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddArtworkToCollection {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      collectionId: isSet(object.collectionId) ? Number(object.collectionId) : 0,
      image: isSet(object.image) ? bytesFromBase64(object.image) : new Uint8Array(),
    };
  },

  toJSON(message: MsgAddArtworkToCollection): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.collectionId !== undefined && (obj.collectionId = Math.round(message.collectionId));
    message.image !== undefined
      && (obj.image = base64FromBytes(message.image !== undefined ? message.image : new Uint8Array()));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddArtworkToCollection>, I>>(object: I): MsgAddArtworkToCollection {
    const message = createBaseMsgAddArtworkToCollection();
    message.creator = object.creator ?? "";
    message.collectionId = object.collectionId ?? 0;
    message.image = object.image ?? new Uint8Array();
    return message;
  },
};

function createBaseMsgAddArtworkToCollectionResponse(): MsgAddArtworkToCollectionResponse {
  return {};
}

export const MsgAddArtworkToCollectionResponse = {
  encode(_: MsgAddArtworkToCollectionResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddArtworkToCollectionResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddArtworkToCollectionResponse();
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

  fromJSON(_: any): MsgAddArtworkToCollectionResponse {
    return {};
  },

  toJSON(_: MsgAddArtworkToCollectionResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddArtworkToCollectionResponse>, I>>(
    _: I,
  ): MsgAddArtworkToCollectionResponse {
    const message = createBaseMsgAddArtworkToCollectionResponse();
    return message;
  },
};

function createBaseMsgAddStoryToCollection(): MsgAddStoryToCollection {
  return { creator: "", collectionId: 0, story: "" };
}

export const MsgAddStoryToCollection = {
  encode(message: MsgAddStoryToCollection, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.collectionId !== 0) {
      writer.uint32(16).uint64(message.collectionId);
    }
    if (message.story !== "") {
      writer.uint32(26).string(message.story);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddStoryToCollection {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddStoryToCollection();
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
          message.story = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddStoryToCollection {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      collectionId: isSet(object.collectionId) ? Number(object.collectionId) : 0,
      story: isSet(object.story) ? String(object.story) : "",
    };
  },

  toJSON(message: MsgAddStoryToCollection): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.collectionId !== undefined && (obj.collectionId = Math.round(message.collectionId));
    message.story !== undefined && (obj.story = message.story);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddStoryToCollection>, I>>(object: I): MsgAddStoryToCollection {
    const message = createBaseMsgAddStoryToCollection();
    message.creator = object.creator ?? "";
    message.collectionId = object.collectionId ?? 0;
    message.story = object.story ?? "";
    return message;
  },
};

function createBaseMsgAddStoryToCollectionResponse(): MsgAddStoryToCollectionResponse {
  return {};
}

export const MsgAddStoryToCollectionResponse = {
  encode(_: MsgAddStoryToCollectionResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddStoryToCollectionResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddStoryToCollectionResponse();
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

  fromJSON(_: any): MsgAddStoryToCollectionResponse {
    return {};
  },

  toJSON(_: MsgAddStoryToCollectionResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddStoryToCollectionResponse>, I>>(_: I): MsgAddStoryToCollectionResponse {
    const message = createBaseMsgAddStoryToCollectionResponse();
    return message;
  },
};

function createBaseMsgSetCardRarity(): MsgSetCardRarity {
  return { creator: "", cardId: 0, collectionId: 0, rarity: "" };
}

export const MsgSetCardRarity = {
  encode(message: MsgSetCardRarity, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.cardId !== 0) {
      writer.uint32(16).uint64(message.cardId);
    }
    if (message.collectionId !== 0) {
      writer.uint32(24).uint64(message.collectionId);
    }
    if (message.rarity !== "") {
      writer.uint32(34).string(message.rarity);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSetCardRarity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSetCardRarity();
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
          message.collectionId = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.rarity = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSetCardRarity {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      cardId: isSet(object.cardId) ? Number(object.cardId) : 0,
      collectionId: isSet(object.collectionId) ? Number(object.collectionId) : 0,
      rarity: isSet(object.rarity) ? String(object.rarity) : "",
    };
  },

  toJSON(message: MsgSetCardRarity): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.cardId !== undefined && (obj.cardId = Math.round(message.cardId));
    message.collectionId !== undefined && (obj.collectionId = Math.round(message.collectionId));
    message.rarity !== undefined && (obj.rarity = message.rarity);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSetCardRarity>, I>>(object: I): MsgSetCardRarity {
    const message = createBaseMsgSetCardRarity();
    message.creator = object.creator ?? "";
    message.cardId = object.cardId ?? 0;
    message.collectionId = object.collectionId ?? 0;
    message.rarity = object.rarity ?? "";
    return message;
  },
};

function createBaseMsgSetCardRarityResponse(): MsgSetCardRarityResponse {
  return {};
}

export const MsgSetCardRarityResponse = {
  encode(_: MsgSetCardRarityResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSetCardRarityResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSetCardRarityResponse();
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

  fromJSON(_: any): MsgSetCardRarityResponse {
    return {};
  },

  toJSON(_: MsgSetCardRarityResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSetCardRarityResponse>, I>>(_: I): MsgSetCardRarityResponse {
    const message = createBaseMsgSetCardRarityResponse();
    return message;
  },
};

function createBaseMsgCreateCouncil(): MsgCreateCouncil {
  return { creator: "", cardId: 0 };
}

export const MsgCreateCouncil = {
  encode(message: MsgCreateCouncil, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.cardId !== 0) {
      writer.uint32(16).uint64(message.cardId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateCouncil {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateCouncil();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.cardId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateCouncil {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      cardId: isSet(object.cardId) ? Number(object.cardId) : 0,
    };
  },

  toJSON(message: MsgCreateCouncil): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.cardId !== undefined && (obj.cardId = Math.round(message.cardId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateCouncil>, I>>(object: I): MsgCreateCouncil {
    const message = createBaseMsgCreateCouncil();
    message.creator = object.creator ?? "";
    message.cardId = object.cardId ?? 0;
    return message;
  },
};

function createBaseMsgCreateCouncilResponse(): MsgCreateCouncilResponse {
  return {};
}

export const MsgCreateCouncilResponse = {
  encode(_: MsgCreateCouncilResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateCouncilResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateCouncilResponse();
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

  fromJSON(_: any): MsgCreateCouncilResponse {
    return {};
  },

  toJSON(_: MsgCreateCouncilResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateCouncilResponse>, I>>(_: I): MsgCreateCouncilResponse {
    const message = createBaseMsgCreateCouncilResponse();
    return message;
  },
};

function createBaseMsgCommitCouncilResponse(): MsgCommitCouncilResponse {
  return { creator: "", response: "", councilId: 0, suggestion: "" };
}

export const MsgCommitCouncilResponse = {
  encode(message: MsgCommitCouncilResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.response !== "") {
      writer.uint32(18).string(message.response);
    }
    if (message.councilId !== 0) {
      writer.uint32(24).uint64(message.councilId);
    }
    if (message.suggestion !== "") {
      writer.uint32(34).string(message.suggestion);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCommitCouncilResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCommitCouncilResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.response = reader.string();
          break;
        case 3:
          message.councilId = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.suggestion = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCommitCouncilResponse {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      response: isSet(object.response) ? String(object.response) : "",
      councilId: isSet(object.councilId) ? Number(object.councilId) : 0,
      suggestion: isSet(object.suggestion) ? String(object.suggestion) : "",
    };
  },

  toJSON(message: MsgCommitCouncilResponse): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.response !== undefined && (obj.response = message.response);
    message.councilId !== undefined && (obj.councilId = Math.round(message.councilId));
    message.suggestion !== undefined && (obj.suggestion = message.suggestion);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCommitCouncilResponse>, I>>(object: I): MsgCommitCouncilResponse {
    const message = createBaseMsgCommitCouncilResponse();
    message.creator = object.creator ?? "";
    message.response = object.response ?? "";
    message.councilId = object.councilId ?? 0;
    message.suggestion = object.suggestion ?? "";
    return message;
  },
};

function createBaseMsgCommitCouncilResponseResponse(): MsgCommitCouncilResponseResponse {
  return {};
}

export const MsgCommitCouncilResponseResponse = {
  encode(_: MsgCommitCouncilResponseResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCommitCouncilResponseResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCommitCouncilResponseResponse();
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

  fromJSON(_: any): MsgCommitCouncilResponseResponse {
    return {};
  },

  toJSON(_: MsgCommitCouncilResponseResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCommitCouncilResponseResponse>, I>>(
    _: I,
  ): MsgCommitCouncilResponseResponse {
    const message = createBaseMsgCommitCouncilResponseResponse();
    return message;
  },
};

function createBaseMsgRevealCouncilResponse(): MsgRevealCouncilResponse {
  return { creator: "", response: 0, secret: "", councilId: 0 };
}

export const MsgRevealCouncilResponse = {
  encode(message: MsgRevealCouncilResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.response !== 0) {
      writer.uint32(16).int32(message.response);
    }
    if (message.secret !== "") {
      writer.uint32(26).string(message.secret);
    }
    if (message.councilId !== 0) {
      writer.uint32(32).uint64(message.councilId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRevealCouncilResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRevealCouncilResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.response = reader.int32() as any;
          break;
        case 3:
          message.secret = reader.string();
          break;
        case 4:
          message.councilId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRevealCouncilResponse {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      response: isSet(object.response) ? responseFromJSON(object.response) : 0,
      secret: isSet(object.secret) ? String(object.secret) : "",
      councilId: isSet(object.councilId) ? Number(object.councilId) : 0,
    };
  },

  toJSON(message: MsgRevealCouncilResponse): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.response !== undefined && (obj.response = responseToJSON(message.response));
    message.secret !== undefined && (obj.secret = message.secret);
    message.councilId !== undefined && (obj.councilId = Math.round(message.councilId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRevealCouncilResponse>, I>>(object: I): MsgRevealCouncilResponse {
    const message = createBaseMsgRevealCouncilResponse();
    message.creator = object.creator ?? "";
    message.response = object.response ?? 0;
    message.secret = object.secret ?? "";
    message.councilId = object.councilId ?? 0;
    return message;
  },
};

function createBaseMsgRevealCouncilResponseResponse(): MsgRevealCouncilResponseResponse {
  return {};
}

export const MsgRevealCouncilResponseResponse = {
  encode(_: MsgRevealCouncilResponseResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRevealCouncilResponseResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRevealCouncilResponseResponse();
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

  fromJSON(_: any): MsgRevealCouncilResponseResponse {
    return {};
  },

  toJSON(_: MsgRevealCouncilResponseResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRevealCouncilResponseResponse>, I>>(
    _: I,
  ): MsgRevealCouncilResponseResponse {
    const message = createBaseMsgRevealCouncilResponseResponse();
    return message;
  },
};

function createBaseMsgRestartCouncil(): MsgRestartCouncil {
  return { creator: "", councilId: 0 };
}

export const MsgRestartCouncil = {
  encode(message: MsgRestartCouncil, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.councilId !== 0) {
      writer.uint32(16).uint64(message.councilId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRestartCouncil {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRestartCouncil();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.councilId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRestartCouncil {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      councilId: isSet(object.councilId) ? Number(object.councilId) : 0,
    };
  },

  toJSON(message: MsgRestartCouncil): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.councilId !== undefined && (obj.councilId = Math.round(message.councilId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRestartCouncil>, I>>(object: I): MsgRestartCouncil {
    const message = createBaseMsgRestartCouncil();
    message.creator = object.creator ?? "";
    message.councilId = object.councilId ?? 0;
    return message;
  },
};

function createBaseMsgRestartCouncilResponse(): MsgRestartCouncilResponse {
  return {};
}

export const MsgRestartCouncilResponse = {
  encode(_: MsgRestartCouncilResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRestartCouncilResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRestartCouncilResponse();
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

  fromJSON(_: any): MsgRestartCouncilResponse {
    return {};
  },

  toJSON(_: MsgRestartCouncilResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRestartCouncilResponse>, I>>(_: I): MsgRestartCouncilResponse {
    const message = createBaseMsgRestartCouncilResponse();
    return message;
  },
};

function createBaseMsgRewokeCouncilRegistration(): MsgRewokeCouncilRegistration {
  return { creator: "" };
}

export const MsgRewokeCouncilRegistration = {
  encode(message: MsgRewokeCouncilRegistration, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRewokeCouncilRegistration {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRewokeCouncilRegistration();
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

  fromJSON(object: any): MsgRewokeCouncilRegistration {
    return { creator: isSet(object.creator) ? String(object.creator) : "" };
  },

  toJSON(message: MsgRewokeCouncilRegistration): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRewokeCouncilRegistration>, I>>(object: I): MsgRewokeCouncilRegistration {
    const message = createBaseMsgRewokeCouncilRegistration();
    message.creator = object.creator ?? "";
    return message;
  },
};

function createBaseMsgRewokeCouncilRegistrationResponse(): MsgRewokeCouncilRegistrationResponse {
  return {};
}

export const MsgRewokeCouncilRegistrationResponse = {
  encode(_: MsgRewokeCouncilRegistrationResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRewokeCouncilRegistrationResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRewokeCouncilRegistrationResponse();
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

  fromJSON(_: any): MsgRewokeCouncilRegistrationResponse {
    return {};
  },

  toJSON(_: MsgRewokeCouncilRegistrationResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRewokeCouncilRegistrationResponse>, I>>(
    _: I,
  ): MsgRewokeCouncilRegistrationResponse {
    const message = createBaseMsgRewokeCouncilRegistrationResponse();
    return message;
  },
};

function createBaseMsgConfirmMatch(): MsgConfirmMatch {
  return { creator: "", matchId: 0, outcome: 0 };
}

export const MsgConfirmMatch = {
  encode(message: MsgConfirmMatch, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.matchId !== 0) {
      writer.uint32(16).uint64(message.matchId);
    }
    if (message.outcome !== 0) {
      writer.uint32(24).int32(message.outcome);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgConfirmMatch {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgConfirmMatch();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.matchId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.outcome = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgConfirmMatch {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      matchId: isSet(object.matchId) ? Number(object.matchId) : 0,
      outcome: isSet(object.outcome) ? outcomeFromJSON(object.outcome) : 0,
    };
  },

  toJSON(message: MsgConfirmMatch): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.matchId !== undefined && (obj.matchId = Math.round(message.matchId));
    message.outcome !== undefined && (obj.outcome = outcomeToJSON(message.outcome));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgConfirmMatch>, I>>(object: I): MsgConfirmMatch {
    const message = createBaseMsgConfirmMatch();
    message.creator = object.creator ?? "";
    message.matchId = object.matchId ?? 0;
    message.outcome = object.outcome ?? 0;
    return message;
  },
};

function createBaseMsgConfirmMatchResponse(): MsgConfirmMatchResponse {
  return {};
}

export const MsgConfirmMatchResponse = {
  encode(_: MsgConfirmMatchResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgConfirmMatchResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgConfirmMatchResponse();
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

  fromJSON(_: any): MsgConfirmMatchResponse {
    return {};
  },

  toJSON(_: MsgConfirmMatchResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgConfirmMatchResponse>, I>>(_: I): MsgConfirmMatchResponse {
    const message = createBaseMsgConfirmMatchResponse();
    return message;
  },
};

function createBaseMsgSetProfileCard(): MsgSetProfileCard {
  return { creator: "", cardId: 0 };
}

export const MsgSetProfileCard = {
  encode(message: MsgSetProfileCard, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.cardId !== 0) {
      writer.uint32(16).uint64(message.cardId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSetProfileCard {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSetProfileCard();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.cardId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSetProfileCard {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      cardId: isSet(object.cardId) ? Number(object.cardId) : 0,
    };
  },

  toJSON(message: MsgSetProfileCard): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.cardId !== undefined && (obj.cardId = Math.round(message.cardId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSetProfileCard>, I>>(object: I): MsgSetProfileCard {
    const message = createBaseMsgSetProfileCard();
    message.creator = object.creator ?? "";
    message.cardId = object.cardId ?? 0;
    return message;
  },
};

function createBaseMsgSetProfileCardResponse(): MsgSetProfileCardResponse {
  return {};
}

export const MsgSetProfileCardResponse = {
  encode(_: MsgSetProfileCardResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSetProfileCardResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSetProfileCardResponse();
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

  fromJSON(_: any): MsgSetProfileCardResponse {
    return {};
  },

  toJSON(_: MsgSetProfileCardResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSetProfileCardResponse>, I>>(_: I): MsgSetProfileCardResponse {
    const message = createBaseMsgSetProfileCardResponse();
    return message;
  },
};

function createBaseMsgOpenBoosterPack(): MsgOpenBoosterPack {
  return { creator: "", boosterPackId: 0 };
}

export const MsgOpenBoosterPack = {
  encode(message: MsgOpenBoosterPack, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.boosterPackId !== 0) {
      writer.uint32(16).uint64(message.boosterPackId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgOpenBoosterPack {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgOpenBoosterPack();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.boosterPackId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgOpenBoosterPack {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      boosterPackId: isSet(object.boosterPackId) ? Number(object.boosterPackId) : 0,
    };
  },

  toJSON(message: MsgOpenBoosterPack): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.boosterPackId !== undefined && (obj.boosterPackId = Math.round(message.boosterPackId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgOpenBoosterPack>, I>>(object: I): MsgOpenBoosterPack {
    const message = createBaseMsgOpenBoosterPack();
    message.creator = object.creator ?? "";
    message.boosterPackId = object.boosterPackId ?? 0;
    return message;
  },
};

function createBaseMsgOpenBoosterPackResponse(): MsgOpenBoosterPackResponse {
  return {};
}

export const MsgOpenBoosterPackResponse = {
  encode(_: MsgOpenBoosterPackResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgOpenBoosterPackResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgOpenBoosterPackResponse();
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

  fromJSON(_: any): MsgOpenBoosterPackResponse {
    return {};
  },

  toJSON(_: MsgOpenBoosterPackResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgOpenBoosterPackResponse>, I>>(_: I): MsgOpenBoosterPackResponse {
    const message = createBaseMsgOpenBoosterPackResponse();
    return message;
  },
};

function createBaseMsgTransferBoosterPack(): MsgTransferBoosterPack {
  return { creator: "", boosterPackId: 0, receiver: "" };
}

export const MsgTransferBoosterPack = {
  encode(message: MsgTransferBoosterPack, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.boosterPackId !== 0) {
      writer.uint32(16).uint64(message.boosterPackId);
    }
    if (message.receiver !== "") {
      writer.uint32(26).string(message.receiver);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgTransferBoosterPack {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgTransferBoosterPack();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.boosterPackId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.receiver = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgTransferBoosterPack {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      boosterPackId: isSet(object.boosterPackId) ? Number(object.boosterPackId) : 0,
      receiver: isSet(object.receiver) ? String(object.receiver) : "",
    };
  },

  toJSON(message: MsgTransferBoosterPack): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.boosterPackId !== undefined && (obj.boosterPackId = Math.round(message.boosterPackId));
    message.receiver !== undefined && (obj.receiver = message.receiver);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgTransferBoosterPack>, I>>(object: I): MsgTransferBoosterPack {
    const message = createBaseMsgTransferBoosterPack();
    message.creator = object.creator ?? "";
    message.boosterPackId = object.boosterPackId ?? 0;
    message.receiver = object.receiver ?? "";
    return message;
  },
};

function createBaseMsgTransferBoosterPackResponse(): MsgTransferBoosterPackResponse {
  return {};
}

export const MsgTransferBoosterPackResponse = {
  encode(_: MsgTransferBoosterPackResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgTransferBoosterPackResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgTransferBoosterPackResponse();
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

  fromJSON(_: any): MsgTransferBoosterPackResponse {
    return {};
  },

  toJSON(_: MsgTransferBoosterPackResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgTransferBoosterPackResponse>, I>>(_: I): MsgTransferBoosterPackResponse {
    const message = createBaseMsgTransferBoosterPackResponse();
    return message;
  },
};

function createBaseMsgSetCollectionStoryWriter(): MsgSetCollectionStoryWriter {
  return { creator: "", collectionId: 0, storyWriter: "" };
}

export const MsgSetCollectionStoryWriter = {
  encode(message: MsgSetCollectionStoryWriter, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.collectionId !== 0) {
      writer.uint32(16).uint64(message.collectionId);
    }
    if (message.storyWriter !== "") {
      writer.uint32(26).string(message.storyWriter);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSetCollectionStoryWriter {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSetCollectionStoryWriter();
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
          message.storyWriter = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSetCollectionStoryWriter {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      collectionId: isSet(object.collectionId) ? Number(object.collectionId) : 0,
      storyWriter: isSet(object.storyWriter) ? String(object.storyWriter) : "",
    };
  },

  toJSON(message: MsgSetCollectionStoryWriter): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.collectionId !== undefined && (obj.collectionId = Math.round(message.collectionId));
    message.storyWriter !== undefined && (obj.storyWriter = message.storyWriter);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSetCollectionStoryWriter>, I>>(object: I): MsgSetCollectionStoryWriter {
    const message = createBaseMsgSetCollectionStoryWriter();
    message.creator = object.creator ?? "";
    message.collectionId = object.collectionId ?? 0;
    message.storyWriter = object.storyWriter ?? "";
    return message;
  },
};

function createBaseMsgSetCollectionStoryWriterResponse(): MsgSetCollectionStoryWriterResponse {
  return {};
}

export const MsgSetCollectionStoryWriterResponse = {
  encode(_: MsgSetCollectionStoryWriterResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSetCollectionStoryWriterResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSetCollectionStoryWriterResponse();
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

  fromJSON(_: any): MsgSetCollectionStoryWriterResponse {
    return {};
  },

  toJSON(_: MsgSetCollectionStoryWriterResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSetCollectionStoryWriterResponse>, I>>(
    _: I,
  ): MsgSetCollectionStoryWriterResponse {
    const message = createBaseMsgSetCollectionStoryWriterResponse();
    return message;
  },
};

function createBaseMsgSetCollectionArtist(): MsgSetCollectionArtist {
  return { creator: "", collectionId: 0, artist: "" };
}

export const MsgSetCollectionArtist = {
  encode(message: MsgSetCollectionArtist, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.collectionId !== 0) {
      writer.uint32(16).uint64(message.collectionId);
    }
    if (message.artist !== "") {
      writer.uint32(26).string(message.artist);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSetCollectionArtist {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSetCollectionArtist();
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
          message.artist = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSetCollectionArtist {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      collectionId: isSet(object.collectionId) ? Number(object.collectionId) : 0,
      artist: isSet(object.artist) ? String(object.artist) : "",
    };
  },

  toJSON(message: MsgSetCollectionArtist): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.collectionId !== undefined && (obj.collectionId = Math.round(message.collectionId));
    message.artist !== undefined && (obj.artist = message.artist);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSetCollectionArtist>, I>>(object: I): MsgSetCollectionArtist {
    const message = createBaseMsgSetCollectionArtist();
    message.creator = object.creator ?? "";
    message.collectionId = object.collectionId ?? 0;
    message.artist = object.artist ?? "";
    return message;
  },
};

function createBaseMsgSetCollectionArtistResponse(): MsgSetCollectionArtistResponse {
  return {};
}

export const MsgSetCollectionArtistResponse = {
  encode(_: MsgSetCollectionArtistResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSetCollectionArtistResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSetCollectionArtistResponse();
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

  fromJSON(_: any): MsgSetCollectionArtistResponse {
    return {};
  },

  toJSON(_: MsgSetCollectionArtistResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSetCollectionArtistResponse>, I>>(_: I): MsgSetCollectionArtistResponse {
    const message = createBaseMsgSetCollectionArtistResponse();
    return message;
  },
};

function createBaseMsgSetUserWebsite(): MsgSetUserWebsite {
  return { creator: "", website: "" };
}

export const MsgSetUserWebsite = {
  encode(message: MsgSetUserWebsite, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.website !== "") {
      writer.uint32(18).string(message.website);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSetUserWebsite {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSetUserWebsite();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.website = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSetUserWebsite {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      website: isSet(object.website) ? String(object.website) : "",
    };
  },

  toJSON(message: MsgSetUserWebsite): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.website !== undefined && (obj.website = message.website);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSetUserWebsite>, I>>(object: I): MsgSetUserWebsite {
    const message = createBaseMsgSetUserWebsite();
    message.creator = object.creator ?? "";
    message.website = object.website ?? "";
    return message;
  },
};

function createBaseMsgSetUserWebsiteResponse(): MsgSetUserWebsiteResponse {
  return {};
}

export const MsgSetUserWebsiteResponse = {
  encode(_: MsgSetUserWebsiteResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSetUserWebsiteResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSetUserWebsiteResponse();
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

  fromJSON(_: any): MsgSetUserWebsiteResponse {
    return {};
  },

  toJSON(_: MsgSetUserWebsiteResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSetUserWebsiteResponse>, I>>(_: I): MsgSetUserWebsiteResponse {
    const message = createBaseMsgSetUserWebsiteResponse();
    return message;
  },
};

function createBaseMsgSetUserBiography(): MsgSetUserBiography {
  return { creator: "", biography: "" };
}

export const MsgSetUserBiography = {
  encode(message: MsgSetUserBiography, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.biography !== "") {
      writer.uint32(18).string(message.biography);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSetUserBiography {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSetUserBiography();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.biography = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSetUserBiography {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      biography: isSet(object.biography) ? String(object.biography) : "",
    };
  },

  toJSON(message: MsgSetUserBiography): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.biography !== undefined && (obj.biography = message.biography);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSetUserBiography>, I>>(object: I): MsgSetUserBiography {
    const message = createBaseMsgSetUserBiography();
    message.creator = object.creator ?? "";
    message.biography = object.biography ?? "";
    return message;
  },
};

function createBaseMsgSetUserBiographyResponse(): MsgSetUserBiographyResponse {
  return {};
}

export const MsgSetUserBiographyResponse = {
  encode(_: MsgSetUserBiographyResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSetUserBiographyResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSetUserBiographyResponse();
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

  fromJSON(_: any): MsgSetUserBiographyResponse {
    return {};
  },

  toJSON(_: MsgSetUserBiographyResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSetUserBiographyResponse>, I>>(_: I): MsgSetUserBiographyResponse {
    const message = createBaseMsgSetUserBiographyResponse();
    return message;
  },
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
  FinalizeCollection(request: MsgFinalizeCollection): Promise<MsgFinalizeCollectionResponse>;
  BuyCollection(request: MsgBuyCollection): Promise<MsgBuyCollectionResponse>;
  RemoveCardFromCollection(request: MsgRemoveCardFromCollection): Promise<MsgRemoveCardFromCollectionResponse>;
  RemoveContributorFromCollection(
    request: MsgRemoveContributorFromCollection,
  ): Promise<MsgRemoveContributorFromCollectionResponse>;
  AddContributorToCollection(request: MsgAddContributorToCollection): Promise<MsgAddContributorToCollectionResponse>;
  SubmitCollectionProposal(request: MsgSubmitCollectionProposal): Promise<MsgSubmitCollectionProposalResponse>;
  CreateSellOffer(request: MsgCreateSellOffer): Promise<MsgCreateSellOfferResponse>;
  BuyCard(request: MsgBuyCard): Promise<MsgBuyCardResponse>;
  RemoveSellOffer(request: MsgRemoveSellOffer): Promise<MsgRemoveSellOfferResponse>;
  AddArtworkToCollection(request: MsgAddArtworkToCollection): Promise<MsgAddArtworkToCollectionResponse>;
  AddStoryToCollection(request: MsgAddStoryToCollection): Promise<MsgAddStoryToCollectionResponse>;
  SetCardRarity(request: MsgSetCardRarity): Promise<MsgSetCardRarityResponse>;
  CreateCouncil(request: MsgCreateCouncil): Promise<MsgCreateCouncilResponse>;
  CommitCouncilResponse(request: MsgCommitCouncilResponse): Promise<MsgCommitCouncilResponseResponse>;
  RevealCouncilResponse(request: MsgRevealCouncilResponse): Promise<MsgRevealCouncilResponseResponse>;
  RestartCouncil(request: MsgRestartCouncil): Promise<MsgRestartCouncilResponse>;
  RewokeCouncilRegistration(request: MsgRewokeCouncilRegistration): Promise<MsgRewokeCouncilRegistrationResponse>;
  ConfirmMatch(request: MsgConfirmMatch): Promise<MsgConfirmMatchResponse>;
  SetProfileCard(request: MsgSetProfileCard): Promise<MsgSetProfileCardResponse>;
  OpenBoosterPack(request: MsgOpenBoosterPack): Promise<MsgOpenBoosterPackResponse>;
  TransferBoosterPack(request: MsgTransferBoosterPack): Promise<MsgTransferBoosterPackResponse>;
  SetCollectionStoryWriter(request: MsgSetCollectionStoryWriter): Promise<MsgSetCollectionStoryWriterResponse>;
  SetCollectionArtist(request: MsgSetCollectionArtist): Promise<MsgSetCollectionArtistResponse>;
  SetUserWebsite(request: MsgSetUserWebsite): Promise<MsgSetUserWebsiteResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  SetUserBiography(request: MsgSetUserBiography): Promise<MsgSetUserBiographyResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Createuser = this.Createuser.bind(this);
    this.BuyCardScheme = this.BuyCardScheme.bind(this);
    this.VoteCard = this.VoteCard.bind(this);
    this.SaveCardContent = this.SaveCardContent.bind(this);
    this.TransferCard = this.TransferCard.bind(this);
    this.DonateToCard = this.DonateToCard.bind(this);
    this.AddArtwork = this.AddArtwork.bind(this);
    this.SubmitCopyrightProposal = this.SubmitCopyrightProposal.bind(this);
    this.ChangeArtist = this.ChangeArtist.bind(this);
    this.RegisterForCouncil = this.RegisterForCouncil.bind(this);
    this.ReportMatch = this.ReportMatch.bind(this);
    this.SubmitMatchReporterProposal = this.SubmitMatchReporterProposal.bind(this);
    this.ApointMatchReporter = this.ApointMatchReporter.bind(this);
    this.CreateCollection = this.CreateCollection.bind(this);
    this.AddCardToCollection = this.AddCardToCollection.bind(this);
    this.FinalizeCollection = this.FinalizeCollection.bind(this);
    this.BuyCollection = this.BuyCollection.bind(this);
    this.RemoveCardFromCollection = this.RemoveCardFromCollection.bind(this);
    this.RemoveContributorFromCollection = this.RemoveContributorFromCollection.bind(this);
    this.AddContributorToCollection = this.AddContributorToCollection.bind(this);
    this.SubmitCollectionProposal = this.SubmitCollectionProposal.bind(this);
    this.CreateSellOffer = this.CreateSellOffer.bind(this);
    this.BuyCard = this.BuyCard.bind(this);
    this.RemoveSellOffer = this.RemoveSellOffer.bind(this);
    this.AddArtworkToCollection = this.AddArtworkToCollection.bind(this);
    this.AddStoryToCollection = this.AddStoryToCollection.bind(this);
    this.SetCardRarity = this.SetCardRarity.bind(this);
    this.CreateCouncil = this.CreateCouncil.bind(this);
    this.CommitCouncilResponse = this.CommitCouncilResponse.bind(this);
    this.RevealCouncilResponse = this.RevealCouncilResponse.bind(this);
    this.RestartCouncil = this.RestartCouncil.bind(this);
    this.RewokeCouncilRegistration = this.RewokeCouncilRegistration.bind(this);
    this.ConfirmMatch = this.ConfirmMatch.bind(this);
    this.SetProfileCard = this.SetProfileCard.bind(this);
    this.OpenBoosterPack = this.OpenBoosterPack.bind(this);
    this.TransferBoosterPack = this.TransferBoosterPack.bind(this);
    this.SetCollectionStoryWriter = this.SetCollectionStoryWriter.bind(this);
    this.SetCollectionArtist = this.SetCollectionArtist.bind(this);
    this.SetUserWebsite = this.SetUserWebsite.bind(this);
    this.SetUserBiography = this.SetUserBiography.bind(this);
  }
  Createuser(request: MsgCreateuser): Promise<MsgCreateuserResponse> {
    const data = MsgCreateuser.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "Createuser", data);
    return promise.then((data) => MsgCreateuserResponse.decode(new _m0.Reader(data)));
  }

  BuyCardScheme(request: MsgBuyCardScheme): Promise<MsgBuyCardSchemeResponse> {
    const data = MsgBuyCardScheme.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "BuyCardScheme", data);
    return promise.then((data) => MsgBuyCardSchemeResponse.decode(new _m0.Reader(data)));
  }

  VoteCard(request: MsgVoteCard): Promise<MsgVoteCardResponse> {
    const data = MsgVoteCard.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "VoteCard", data);
    return promise.then((data) => MsgVoteCardResponse.decode(new _m0.Reader(data)));
  }

  SaveCardContent(request: MsgSaveCardContent): Promise<MsgSaveCardContentResponse> {
    const data = MsgSaveCardContent.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "SaveCardContent", data);
    return promise.then((data) => MsgSaveCardContentResponse.decode(new _m0.Reader(data)));
  }

  TransferCard(request: MsgTransferCard): Promise<MsgTransferCardResponse> {
    const data = MsgTransferCard.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "TransferCard", data);
    return promise.then((data) => MsgTransferCardResponse.decode(new _m0.Reader(data)));
  }

  DonateToCard(request: MsgDonateToCard): Promise<MsgDonateToCardResponse> {
    const data = MsgDonateToCard.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "DonateToCard", data);
    return promise.then((data) => MsgDonateToCardResponse.decode(new _m0.Reader(data)));
  }

  AddArtwork(request: MsgAddArtwork): Promise<MsgAddArtworkResponse> {
    const data = MsgAddArtwork.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "AddArtwork", data);
    return promise.then((data) => MsgAddArtworkResponse.decode(new _m0.Reader(data)));
  }

  SubmitCopyrightProposal(request: MsgSubmitCopyrightProposal): Promise<MsgSubmitCopyrightProposalResponse> {
    const data = MsgSubmitCopyrightProposal.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "SubmitCopyrightProposal", data);
    return promise.then((data) => MsgSubmitCopyrightProposalResponse.decode(new _m0.Reader(data)));
  }

  ChangeArtist(request: MsgChangeArtist): Promise<MsgChangeArtistResponse> {
    const data = MsgChangeArtist.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "ChangeArtist", data);
    return promise.then((data) => MsgChangeArtistResponse.decode(new _m0.Reader(data)));
  }

  RegisterForCouncil(request: MsgRegisterForCouncil): Promise<MsgRegisterForCouncilResponse> {
    const data = MsgRegisterForCouncil.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "RegisterForCouncil", data);
    return promise.then((data) => MsgRegisterForCouncilResponse.decode(new _m0.Reader(data)));
  }

  ReportMatch(request: MsgReportMatch): Promise<MsgReportMatchResponse> {
    const data = MsgReportMatch.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "ReportMatch", data);
    return promise.then((data) => MsgReportMatchResponse.decode(new _m0.Reader(data)));
  }

  SubmitMatchReporterProposal(
    request: MsgSubmitMatchReporterProposal,
  ): Promise<MsgSubmitMatchReporterProposalResponse> {
    const data = MsgSubmitMatchReporterProposal.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "SubmitMatchReporterProposal", data);
    return promise.then((data) => MsgSubmitMatchReporterProposalResponse.decode(new _m0.Reader(data)));
  }

  ApointMatchReporter(request: MsgApointMatchReporter): Promise<MsgApointMatchReporterResponse> {
    const data = MsgApointMatchReporter.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "ApointMatchReporter", data);
    return promise.then((data) => MsgApointMatchReporterResponse.decode(new _m0.Reader(data)));
  }

  CreateCollection(request: MsgCreateCollection): Promise<MsgCreateCollectionResponse> {
    const data = MsgCreateCollection.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "CreateCollection", data);
    return promise.then((data) => MsgCreateCollectionResponse.decode(new _m0.Reader(data)));
  }

  AddCardToCollection(request: MsgAddCardToCollection): Promise<MsgAddCardToCollectionResponse> {
    const data = MsgAddCardToCollection.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "AddCardToCollection", data);
    return promise.then((data) => MsgAddCardToCollectionResponse.decode(new _m0.Reader(data)));
  }

  FinalizeCollection(request: MsgFinalizeCollection): Promise<MsgFinalizeCollectionResponse> {
    const data = MsgFinalizeCollection.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "FinalizeCollection", data);
    return promise.then((data) => MsgFinalizeCollectionResponse.decode(new _m0.Reader(data)));
  }

  BuyCollection(request: MsgBuyCollection): Promise<MsgBuyCollectionResponse> {
    const data = MsgBuyCollection.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "BuyCollection", data);
    return promise.then((data) => MsgBuyCollectionResponse.decode(new _m0.Reader(data)));
  }

  RemoveCardFromCollection(request: MsgRemoveCardFromCollection): Promise<MsgRemoveCardFromCollectionResponse> {
    const data = MsgRemoveCardFromCollection.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "RemoveCardFromCollection", data);
    return promise.then((data) => MsgRemoveCardFromCollectionResponse.decode(new _m0.Reader(data)));
  }

  RemoveContributorFromCollection(
    request: MsgRemoveContributorFromCollection,
  ): Promise<MsgRemoveContributorFromCollectionResponse> {
    const data = MsgRemoveContributorFromCollection.encode(request).finish();
    const promise = this.rpc.request(
      "DecentralCardGame.cardchain.cardchain.Msg",
      "RemoveContributorFromCollection",
      data,
    );
    return promise.then((data) => MsgRemoveContributorFromCollectionResponse.decode(new _m0.Reader(data)));
  }

  AddContributorToCollection(request: MsgAddContributorToCollection): Promise<MsgAddContributorToCollectionResponse> {
    const data = MsgAddContributorToCollection.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "AddContributorToCollection", data);
    return promise.then((data) => MsgAddContributorToCollectionResponse.decode(new _m0.Reader(data)));
  }

  SubmitCollectionProposal(request: MsgSubmitCollectionProposal): Promise<MsgSubmitCollectionProposalResponse> {
    const data = MsgSubmitCollectionProposal.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "SubmitCollectionProposal", data);
    return promise.then((data) => MsgSubmitCollectionProposalResponse.decode(new _m0.Reader(data)));
  }

  CreateSellOffer(request: MsgCreateSellOffer): Promise<MsgCreateSellOfferResponse> {
    const data = MsgCreateSellOffer.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "CreateSellOffer", data);
    return promise.then((data) => MsgCreateSellOfferResponse.decode(new _m0.Reader(data)));
  }

  BuyCard(request: MsgBuyCard): Promise<MsgBuyCardResponse> {
    const data = MsgBuyCard.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "BuyCard", data);
    return promise.then((data) => MsgBuyCardResponse.decode(new _m0.Reader(data)));
  }

  RemoveSellOffer(request: MsgRemoveSellOffer): Promise<MsgRemoveSellOfferResponse> {
    const data = MsgRemoveSellOffer.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "RemoveSellOffer", data);
    return promise.then((data) => MsgRemoveSellOfferResponse.decode(new _m0.Reader(data)));
  }

  AddArtworkToCollection(request: MsgAddArtworkToCollection): Promise<MsgAddArtworkToCollectionResponse> {
    const data = MsgAddArtworkToCollection.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "AddArtworkToCollection", data);
    return promise.then((data) => MsgAddArtworkToCollectionResponse.decode(new _m0.Reader(data)));
  }

  AddStoryToCollection(request: MsgAddStoryToCollection): Promise<MsgAddStoryToCollectionResponse> {
    const data = MsgAddStoryToCollection.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "AddStoryToCollection", data);
    return promise.then((data) => MsgAddStoryToCollectionResponse.decode(new _m0.Reader(data)));
  }

  SetCardRarity(request: MsgSetCardRarity): Promise<MsgSetCardRarityResponse> {
    const data = MsgSetCardRarity.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "SetCardRarity", data);
    return promise.then((data) => MsgSetCardRarityResponse.decode(new _m0.Reader(data)));
  }

  CreateCouncil(request: MsgCreateCouncil): Promise<MsgCreateCouncilResponse> {
    const data = MsgCreateCouncil.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "CreateCouncil", data);
    return promise.then((data) => MsgCreateCouncilResponse.decode(new _m0.Reader(data)));
  }

  CommitCouncilResponse(request: MsgCommitCouncilResponse): Promise<MsgCommitCouncilResponseResponse> {
    const data = MsgCommitCouncilResponse.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "CommitCouncilResponse", data);
    return promise.then((data) => MsgCommitCouncilResponseResponse.decode(new _m0.Reader(data)));
  }

  RevealCouncilResponse(request: MsgRevealCouncilResponse): Promise<MsgRevealCouncilResponseResponse> {
    const data = MsgRevealCouncilResponse.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "RevealCouncilResponse", data);
    return promise.then((data) => MsgRevealCouncilResponseResponse.decode(new _m0.Reader(data)));
  }

  RestartCouncil(request: MsgRestartCouncil): Promise<MsgRestartCouncilResponse> {
    const data = MsgRestartCouncil.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "RestartCouncil", data);
    return promise.then((data) => MsgRestartCouncilResponse.decode(new _m0.Reader(data)));
  }

  RewokeCouncilRegistration(request: MsgRewokeCouncilRegistration): Promise<MsgRewokeCouncilRegistrationResponse> {
    const data = MsgRewokeCouncilRegistration.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "RewokeCouncilRegistration", data);
    return promise.then((data) => MsgRewokeCouncilRegistrationResponse.decode(new _m0.Reader(data)));
  }

  ConfirmMatch(request: MsgConfirmMatch): Promise<MsgConfirmMatchResponse> {
    const data = MsgConfirmMatch.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "ConfirmMatch", data);
    return promise.then((data) => MsgConfirmMatchResponse.decode(new _m0.Reader(data)));
  }

  SetProfileCard(request: MsgSetProfileCard): Promise<MsgSetProfileCardResponse> {
    const data = MsgSetProfileCard.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "SetProfileCard", data);
    return promise.then((data) => MsgSetProfileCardResponse.decode(new _m0.Reader(data)));
  }

  OpenBoosterPack(request: MsgOpenBoosterPack): Promise<MsgOpenBoosterPackResponse> {
    const data = MsgOpenBoosterPack.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "OpenBoosterPack", data);
    return promise.then((data) => MsgOpenBoosterPackResponse.decode(new _m0.Reader(data)));
  }

  TransferBoosterPack(request: MsgTransferBoosterPack): Promise<MsgTransferBoosterPackResponse> {
    const data = MsgTransferBoosterPack.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "TransferBoosterPack", data);
    return promise.then((data) => MsgTransferBoosterPackResponse.decode(new _m0.Reader(data)));
  }

  SetCollectionStoryWriter(request: MsgSetCollectionStoryWriter): Promise<MsgSetCollectionStoryWriterResponse> {
    const data = MsgSetCollectionStoryWriter.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "SetCollectionStoryWriter", data);
    return promise.then((data) => MsgSetCollectionStoryWriterResponse.decode(new _m0.Reader(data)));
  }

  SetCollectionArtist(request: MsgSetCollectionArtist): Promise<MsgSetCollectionArtistResponse> {
    const data = MsgSetCollectionArtist.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "SetCollectionArtist", data);
    return promise.then((data) => MsgSetCollectionArtistResponse.decode(new _m0.Reader(data)));
  }

  SetUserWebsite(request: MsgSetUserWebsite): Promise<MsgSetUserWebsiteResponse> {
    const data = MsgSetUserWebsite.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "SetUserWebsite", data);
    return promise.then((data) => MsgSetUserWebsiteResponse.decode(new _m0.Reader(data)));
  }

  SetUserBiography(request: MsgSetUserBiography): Promise<MsgSetUserBiographyResponse> {
    const data = MsgSetUserBiography.encode(request).finish();
    const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "SetUserBiography", data);
    return promise.then((data) => MsgSetUserBiographyResponse.decode(new _m0.Reader(data)));
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

function bytesFromBase64(b64: string): Uint8Array {
  if (globalThis.Buffer) {
    return Uint8Array.from(globalThis.Buffer.from(b64, "base64"));
  } else {
    const bin = globalThis.atob(b64);
    const arr = new Uint8Array(bin.length);
    for (let i = 0; i < bin.length; ++i) {
      arr[i] = bin.charCodeAt(i);
    }
    return arr;
  }
}

function base64FromBytes(arr: Uint8Array): string {
  if (globalThis.Buffer) {
    return globalThis.Buffer.from(arr).toString("base64");
  } else {
    const bin: string[] = [];
    arr.forEach((byte) => {
      bin.push(String.fromCharCode(byte));
    });
    return globalThis.btoa(bin.join(""));
  }
}

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
