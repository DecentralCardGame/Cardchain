/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface CardchainAirDrops {
  vote?: boolean;
  create?: boolean;
  buy?: boolean;
  play?: boolean;
  user?: boolean;
}

export interface CardchainBoosterPack {
  /** @format uint64 */
  collectionId?: string;

  /** @format int64 */
  timeStamp?: string;
  raritiesPerPack?: string[];
}

export enum CardchainCStatus {
  Design = "design",
  Finalized = "finalized",
  Active = "active",
  Archived = "archived",
}

export enum CardchainCouncelingStatus {
  CouncilOpen = "councilOpen",
  CouncilCreated = "councilCreated",
  CouncilClosed = "councilClosed",
  Commited = "commited",
  Revealed = "revealed",
  SuggestionsMade = "suggestionsMade",
}

export interface CardchainCouncil {
  /** @format uint64 */
  cardId?: string;
  voters?: string[];
  hashResponses?: CardchainWrapHashResponse[];
  clearResponses?: CardchainWrapClearResponse[];
  treasury?: string;
  status?: CardchainCouncelingStatus;

  /** @format uint64 */
  trialStart?: string;
}

export interface CardchainCouncilParticipation {
  status?: CardchainCouncilStatus;

  /** @format uint64 */
  council?: string;
}

export enum CardchainCouncilStatus {
  Available = "available",
  Unavailable = "unavailable",
  OpenCouncil = "openCouncil",
  StartedCouncil = "startedCouncil",
}

export interface CardchainIgnoreCouncils {
  status?: boolean;
  card?: boolean;
}

export interface CardchainIgnoreMatches {
  outcome?: boolean;
}

export interface CardchainIgnoreSellOffers {
  status?: boolean;
  card?: boolean;
}

export interface CardchainMatch {
  /** @format uint64 */
  timestamp?: string;
  reporter?: string;
  playerA?: CardchainMatchPlayer;
  playerB?: CardchainMatchPlayer;
  outcome?: CardchainOutcome;
  coinsDistributed?: boolean;
}

export interface CardchainMatchPlayer {
  addr?: string;
  playedCards?: string[];
  confirmed?: boolean;
  outcome?: CardchainOutcome;
}

export type CardchainMsgAddArtworkResponse = object;

export type CardchainMsgAddArtworkToCollectionResponse = object;

export type CardchainMsgAddCardToCollectionResponse = object;

export type CardchainMsgAddContributorToCollectionResponse = object;

export type CardchainMsgAddStoryToCollectionResponse = object;

export type CardchainMsgApointMatchReporterResponse = object;

export type CardchainMsgBuyCardResponse = object;

export type CardchainMsgBuyCardSchemeResponse = object;

export interface CardchainMsgBuyCollectionResponse {
  airdropClaimed?: boolean;
}

export type CardchainMsgChangeArtistResponse = object;

export type CardchainMsgCommitCouncilResponseResponse = object;

export type CardchainMsgConfirmMatchResponse = object;

export type CardchainMsgCreateCollectionResponse = object;

export type CardchainMsgCreateCouncilResponse = object;

export type CardchainMsgCreateSellOfferResponse = object;

export type CardchainMsgCreateuserResponse = object;

export type CardchainMsgDonateToCardResponse = object;

export type CardchainMsgFinalizeCollectionResponse = object;

export type CardchainMsgOpenBoosterPackResponse = object;

export type CardchainMsgRegisterForCouncilResponse = object;

export type CardchainMsgRemoveCardFromCollectionResponse = object;

export type CardchainMsgRemoveContributorFromCollectionResponse = object;

export type CardchainMsgRemoveSellOfferResponse = object;

export interface CardchainMsgReportMatchResponse {
  /** @format uint64 */
  matchId?: string;
}

export type CardchainMsgRestartCouncilResponse = object;

export type CardchainMsgRevealCouncilResponseResponse = object;

export type CardchainMsgRewokeCouncilRegistrationResponse = object;

export interface CardchainMsgSaveCardContentResponse {
  airdropClaimed?: boolean;
}

export type CardchainMsgSetCardRarityResponse = object;

export type CardchainMsgSetCollectionArtistResponse = object;

export type CardchainMsgSetCollectionStoryWriterResponse = object;

export type CardchainMsgSetProfileCardResponse = object;

export type CardchainMsgSetUserBiographyResponse = object;

export type CardchainMsgSetUserWebsiteResponse = object;

export type CardchainMsgSubmitCollectionProposalResponse = object;

export type CardchainMsgSubmitCopyrightProposalResponse = object;

export type CardchainMsgSubmitMatchReporterProposalResponse = object;

export type CardchainMsgTransferBoosterPackResponse = object;

export type CardchainMsgTransferCardResponse = object;

export interface CardchainMsgVoteCardResponse {
  airdropClaimed?: boolean;
}

export enum CardchainOutcome {
  AWon = "AWon",
  BWon = "BWon",
  Draw = "Draw",
  Aborted = "Aborted",
}

export interface CardchainOutpCard {
  owner?: string;
  artist?: string;
  content?: string;
  image?: string;
  fullArt?: boolean;
  notes?: string;
  status?: CardchaincardchainStatus;
  votePool?: string;
  voters?: string[];

  /** @format uint64 */
  fairEnoughVotes?: string;

  /** @format uint64 */
  overpoweredVotes?: string;

  /** @format uint64 */
  underpoweredVotes?: string;

  /** @format uint64 */
  inappropriateVotes?: string;

  /** @format int64 */
  nerflevel?: string;
}

export interface CardchainOutpCollection {
  name?: string;
  cards?: string[];
  artist?: string;
  storyWriter?: string;
  contributors?: string[];
  story?: string;
  artwork?: string;
  status?: CardchainCStatus;

  /** @format int64 */
  timeStamp?: string;
}

/**
 * Params defines the parameters for the module.
 */
export interface CardchainParams {
  /** @format int64 */
  votingRightsExpirationTime?: string;

  /** @format uint64 */
  collectionSize?: string;
  collectionPrice?: string;

  /** @format uint64 */
  activeCollectionsAmount?: string;
  collectionCreationFee?: string;
  collateralDeposit?: string;

  /** @format int64 */
  winnerReward?: string;
  hourlyFaucet?: string;
  inflationRate?: string;

  /** @format uint64 */
  raresPerPack?: string;

  /** @format uint64 */
  commonsPerPack?: string;

  /** @format uint64 */
  unCommonsPerPack?: string;

  /** @format uint64 */
  trialPeriod?: string;

  /** @format int64 */
  gameVoteRatio?: string;

  /** @format int64 */
  cardAuctionPriceReductionPeriod?: string;
  airDropValue?: string;

  /** @format int64 */
  airDropMaxBlockHeight?: string;
  trialVoteReward?: string;

  /** @format int64 */
  votePoolFraction?: string;

  /** @format int64 */
  votingRewardCap?: string;
}

/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface CardchainQueryParamsResponse {
  /** params holds all the parameters of this module. */
  params?: CardchainParams;
}

export interface CardchainQueryQCardContentResponse {
  content?: string;
  hash?: string;
}

export interface CardchainQueryQCardchainInfoResponse {
  cardAuctionPrice?: string;
  activeCollections?: string[];

  /** @format uint64 */
  cardsNumber?: string;

  /** @format uint64 */
  matchesNumber?: string;

  /** @format uint64 */
  sellOffersNumber?: string;

  /** @format uint64 */
  councilsNumber?: string;
}

export enum CardchainQueryQCardsRequestStatus {
  Scheme = "scheme",
  Prototype = "prototype",
  Trial = "trial",
  Permanent = "permanent",
  Suspended = "suspended",
  Banned = "banned",
  BannedSoon = "bannedSoon",
  BannedVerySoon = "bannedVerySoon",
  None = "none",
  Playable = "playable",
  Unplayable = "unplayable",
}

export interface CardchainQueryQCardsResponse {
  cardsList?: string[];
}

export interface CardchainQueryQCollectionsResponse {
  collectionIds?: string[];
}

export interface CardchainQueryQCouncilsResponse {
  councilssIds?: string[];
  councils?: CardchainCouncil[];
}

export interface CardchainQueryQMatchesResponse {
  matchesList?: string[];
  matches?: CardchainMatch[];
}

export interface CardchainQueryQSellOffersResponse {
  sellOffersIds?: string[];
  sellOffers?: CardchainSellOffer[];
}

export interface CardchainQueryQVotableCardsResponse {
  unregistered?: boolean;
  noVoteRights?: boolean;
  voteRights?: CardchainVoteRight[];
}

export interface CardchainQueryQVotingResultsResponse {
  lastVotingResults?: CardchainVotingResults;
}

export interface CardchainQueryRarityDistributionResponse {
  current?: string[];
  wanted?: string[];
}

export enum CardchainResponse {
  Yes = "Yes",
  No = "No",
  Suggestion = "Suggestion",
}

export interface CardchainSellOffer {
  seller?: string;
  buyer?: string;

  /** @format uint64 */
  card?: string;
  price?: string;
  status?: CardchainSellOfferStatus;
}

export enum CardchainSellOfferStatus {
  Open = "open",
  Sold = "sold",
  Removed = "removed",
}

export interface CardchainServer {
  reporter?: string;

  /** @format uint64 */
  invalidReports?: string;

  /** @format uint64 */
  validReports?: string;
}

export interface CardchainUser {
  alias?: string;
  ownedCardSchemes?: string[];
  ownedPrototypes?: string[];
  cards?: string[];
  voteRights?: CardchainVoteRight[];
  councilParticipation?: CardchainCouncilParticipation;
  ReportMatches?: boolean;

  /** @format uint64 */
  profileCard?: string;
  airDrops?: CardchainAirDrops;
  boosterPacks?: CardchainBoosterPack[];
  website?: string;
  biography?: string;
}

export interface CardchainVoteRight {
  /** @format uint64 */
  cardId?: string;

  /** @format int64 */
  expireBlock?: string;
}

export interface CardchainVotingResult {
  /** @format uint64 */
  cardId?: string;

  /** @format uint64 */
  fairEnoughVotes?: string;

  /** @format uint64 */
  overpoweredVotes?: string;

  /** @format uint64 */
  underpoweredVotes?: string;

  /** @format uint64 */
  inappropriateVotes?: string;
  result?: string;
}

export interface CardchainVotingResults {
  /** @format uint64 */
  totalVotes?: string;

  /** @format uint64 */
  totalFairEnoughVotes?: string;

  /** @format uint64 */
  totalOverpoweredVotes?: string;

  /** @format uint64 */
  totalUnderpoweredVotes?: string;

  /** @format uint64 */
  totalInappropriateVotes?: string;
  cardResults?: CardchainVotingResult[];
  notes?: string;
}

export interface CardchainWrapClearResponse {
  user?: string;
  response?: CardchainResponse;
  suggestion?: string;
}

export interface CardchainWrapHashResponse {
  user?: string;
  hash?: string;
}

export enum CardchaincardchainStatus {
  Scheme = "scheme",
  Prototype = "prototype",
  Trial = "trial",
  Permanent = "permanent",
  Suspended = "suspended",
  Banned = "banned",
  BannedSoon = "bannedSoon",
  BannedVerySoon = "bannedVerySoon",
  None = "none",
  InCouncil = "inCouncil",
}

export interface GooglerpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: ProtobufAny[];
}

export interface ProtobufAny {
  "@type"?: string;
}

import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse, ResponseType } from "axios";

export type QueryParamsType = Record<string | number, any>;

export interface FullRequestParams extends Omit<AxiosRequestConfig, "data" | "params" | "url" | "responseType"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: ResponseType;
  /** request body */
  body?: unknown;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> extends Omit<AxiosRequestConfig, "data" | "cancelToken"> {
  securityWorker?: (
    securityData: SecurityDataType | null,
  ) => Promise<AxiosRequestConfig | void> | AxiosRequestConfig | void;
  secure?: boolean;
  format?: ResponseType;
}

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public instance: AxiosInstance;
  private securityData: SecurityDataType | null = null;
  private securityWorker?: ApiConfig<SecurityDataType>["securityWorker"];
  private secure?: boolean;
  private format?: ResponseType;

  constructor({ securityWorker, secure, format, ...axiosConfig }: ApiConfig<SecurityDataType> = {}) {
    this.instance = axios.create({ ...axiosConfig, baseURL: axiosConfig.baseURL || "" });
    this.secure = secure;
    this.format = format;
    this.securityWorker = securityWorker;
  }

  public setSecurityData = (data: SecurityDataType | null) => {
    this.securityData = data;
  };

  private mergeRequestParams(params1: AxiosRequestConfig, params2?: AxiosRequestConfig): AxiosRequestConfig {
    return {
      ...this.instance.defaults,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.instance.defaults.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createFormData(input: Record<string, unknown>): FormData {
    return Object.keys(input || {}).reduce((formData, key) => {
      const property = input[key];
      formData.append(
        key,
        property instanceof Blob
          ? property
          : typeof property === "object" && property !== null
          ? JSON.stringify(property)
          : `${property}`,
      );
      return formData;
    }, new FormData());
  }

  public request = async <T = any, _E = any>({
    secure,
    path,
    type,
    query,
    format,
    body,
    ...params
  }: FullRequestParams): Promise<AxiosResponse<T>> => {
    const secureParams =
      ((typeof secure === "boolean" ? secure : this.secure) &&
        this.securityWorker &&
        (await this.securityWorker(this.securityData))) ||
      {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const responseFormat = (format && this.format) || void 0;

    if (type === ContentType.FormData && body && body !== null && typeof body === "object") {
      requestParams.headers.common = { Accept: "*/*" };
      requestParams.headers.post = {};
      requestParams.headers.put = {};

      body = this.createFormData(body as Record<string, unknown>);
    }

    return this.instance.request({
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      params: query,
      responseType: responseFormat,
      data: body,
      url: path,
    });
  };
}

/**
 * @title cardchain/cardchain/card.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryQCollections
   * @summary Queries a list of QCollections items.
   * @request GET:/DecentralCardGame/Cardchain/cardchain/q_collections/{status}/{ignoreStatus}
   */
  queryQCollections = (
    status: "design" | "finalized" | "active" | "archived",
    ignoreStatus: boolean,
    query?: { contributors?: string[]; containsCards?: string[]; owner?: string },
    params: RequestParams = {},
  ) =>
    this.request<CardchainQueryQCollectionsResponse, GooglerpcStatus>({
      path: `/DecentralCardGame/Cardchain/cardchain/q_collections/${status}/${ignoreStatus}`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryQCouncils
   * @summary Queries a list of QCouncils items.
   * @request GET:/DecentralCardGame/Cardchain/cardchain/q_councils/{status}
   */
  queryQCouncils = (
    status: "councilOpen" | "councilCreated" | "councilClosed" | "commited" | "revealed" | "suggestionsMade",
    query?: { voters?: string[]; card?: string; creator?: string; "ignore.status"?: boolean; "ignore.card"?: boolean },
    params: RequestParams = {},
  ) =>
    this.request<CardchainQueryQCouncilsResponse, GooglerpcStatus>({
      path: `/DecentralCardGame/Cardchain/cardchain/q_councils/${status}`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryRarityDistribution
   * @summary Queries a list of RarityDistribution items.
   * @request GET:/DecentralCardGame/Cardchain/cardchain/rarity_distribution/{collectionId}
   */
  queryRarityDistribution = (collectionId: string, params: RequestParams = {}) =>
    this.request<CardchainQueryRarityDistributionResponse, GooglerpcStatus>({
      path: `/DecentralCardGame/Cardchain/cardchain/rarity_distribution/${collectionId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryParams
   * @summary Parameters queries the parameters of the module.
   * @request GET:/DecentralCardGame/cardchain/cardchain/params
   */
  queryParams = (params: RequestParams = {}) =>
    this.request<CardchainQueryParamsResponse, GooglerpcStatus>({
      path: `/DecentralCardGame/cardchain/cardchain/params`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryQCard
   * @summary Queries a list of QCard items.
   * @request GET:/DecentralCardGame/cardchain/cardchain/q_card/{cardId}
   */
  queryQCard = (cardId: string, params: RequestParams = {}) =>
    this.request<CardchainOutpCard, GooglerpcStatus>({
      path: `/DecentralCardGame/cardchain/cardchain/q_card/${cardId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryQCardContent
   * @summary Queries a list of QCardContent items.
   * @request GET:/DecentralCardGame/cardchain/cardchain/q_card_content/{cardId}
   */
  queryQCardContent = (cardId: string, params: RequestParams = {}) =>
    this.request<CardchainQueryQCardContentResponse, GooglerpcStatus>({
      path: `/DecentralCardGame/cardchain/cardchain/q_card_content/${cardId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryQCardchainInfo
   * @summary Queries a list of QCardchainInfo items.
   * @request GET:/DecentralCardGame/cardchain/cardchain/q_cardchain_info
   */
  queryQCardchainInfo = (params: RequestParams = {}) =>
    this.request<CardchainQueryQCardchainInfoResponse, GooglerpcStatus>({
      path: `/DecentralCardGame/cardchain/cardchain/q_cardchain_info`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryQCards
   * @summary Queries a list of QCards items.
   * @request GET:/DecentralCardGame/cardchain/cardchain/q_cards/{status}
   */
  queryQCards = (
    status:
      | "scheme"
      | "prototype"
      | "trial"
      | "permanent"
      | "suspended"
      | "banned"
      | "bannedSoon"
      | "bannedVerySoon"
      | "none"
      | "playable"
      | "unplayable",
    query?: {
      owner?: string;
      cardType?: string;
      classes?: string;
      sortBy?: string;
      nameContains?: string;
      keywordsContains?: string;
      notesContains?: string;
    },
    params: RequestParams = {},
  ) =>
    this.request<CardchainQueryQCardsResponse, GooglerpcStatus>({
      path: `/DecentralCardGame/cardchain/cardchain/q_cards/${status}`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryQCollection
   * @summary Queries a list of QCollection items.
   * @request GET:/DecentralCardGame/cardchain/cardchain/q_collection/{collectionId}
   */
  queryQCollection = (collectionId: string, params: RequestParams = {}) =>
    this.request<CardchainOutpCollection, GooglerpcStatus>({
      path: `/DecentralCardGame/cardchain/cardchain/q_collection/${collectionId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryQCouncil
   * @summary Queries a list of QCouncil items.
   * @request GET:/DecentralCardGame/cardchain/cardchain/q_council/{councilId}
   */
  queryQCouncil = (councilId: string, params: RequestParams = {}) =>
    this.request<CardchainCouncil, GooglerpcStatus>({
      path: `/DecentralCardGame/cardchain/cardchain/q_council/${councilId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryQMatch
   * @summary Queries a list of QMatch items.
   * @request GET:/DecentralCardGame/cardchain/cardchain/q_match/{matchId}
   */
  queryQMatch = (matchId: string, params: RequestParams = {}) =>
    this.request<CardchainMatch, GooglerpcStatus>({
      path: `/DecentralCardGame/cardchain/cardchain/q_match/${matchId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryQMatches
   * @summary Queries a list of QMatches items.
   * @request GET:/DecentralCardGame/cardchain/cardchain/q_matches
   */
  queryQMatches = (
    query?: {
      timestampDown?: string;
      timestampUp?: string;
      containsUsers?: string[];
      reporter?: string;
      outcome?: "AWon" | "BWon" | "Draw" | "Aborted";
      cardsPlayed?: string[];
      "ignore.outcome"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<CardchainQueryQMatchesResponse, GooglerpcStatus>({
      path: `/DecentralCardGame/cardchain/cardchain/q_matches`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryQSellOffer
   * @summary Queries a list of QSellOffer items.
   * @request GET:/DecentralCardGame/cardchain/cardchain/q_sell_offer/{sellOfferId}
   */
  queryQSellOffer = (sellOfferId: string, params: RequestParams = {}) =>
    this.request<CardchainSellOffer, GooglerpcStatus>({
      path: `/DecentralCardGame/cardchain/cardchain/q_sell_offer/${sellOfferId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryQSellOffers
   * @summary Queries a list of QSellOffers items.
   * @request GET:/DecentralCardGame/cardchain/cardchain/q_sell_offers/{status}
   */
  queryQSellOffers = (
    status: "open" | "sold" | "removed",
    query?: {
      priceDown?: string;
      priceUp?: string;
      seller?: string;
      buyer?: string;
      card?: string;
      "ignore.status"?: boolean;
      "ignore.card"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<CardchainQueryQSellOffersResponse, GooglerpcStatus>({
      path: `/DecentralCardGame/cardchain/cardchain/q_sell_offers/${status}`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryQServer
   * @summary Queries a list of QServer items.
   * @request GET:/DecentralCardGame/cardchain/cardchain/q_server/{id}
   */
  queryQServer = (id: string, params: RequestParams = {}) =>
    this.request<CardchainServer, GooglerpcStatus>({
      path: `/DecentralCardGame/cardchain/cardchain/q_server/${id}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryQUser
   * @summary Queries a list of QUser items.
   * @request GET:/DecentralCardGame/cardchain/cardchain/q_user/{address}
   */
  queryQUser = (address: string, params: RequestParams = {}) =>
    this.request<CardchainUser, GooglerpcStatus>({
      path: `/DecentralCardGame/cardchain/cardchain/q_user/${address}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryQVotableCards
   * @summary Queries a list of QVotableCards items.
   * @request GET:/DecentralCardGame/cardchain/cardchain/q_votable_cards/{address}
   */
  queryQVotableCards = (address: string, params: RequestParams = {}) =>
    this.request<CardchainQueryQVotableCardsResponse, GooglerpcStatus>({
      path: `/DecentralCardGame/cardchain/cardchain/q_votable_cards/${address}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryQVotingResults
   * @summary Queries a list of QVotingResults items.
   * @request GET:/DecentralCardGame/cardchain/cardchain/q_voting_results
   */
  queryQVotingResults = (params: RequestParams = {}) =>
    this.request<CardchainQueryQVotingResultsResponse, GooglerpcStatus>({
      path: `/DecentralCardGame/cardchain/cardchain/q_voting_results`,
      method: "GET",
      format: "json",
      ...params,
    });
}
