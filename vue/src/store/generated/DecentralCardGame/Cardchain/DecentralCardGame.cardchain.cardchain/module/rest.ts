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

export enum CardchainCouncilStatus {
  Available = "available",
  Unavailable = "unavailable",
  OpenCouncil = "openCouncil",
  StartedCouncil = "startedCouncil",
}

export type CardchainMsgAddArtworkResponse = object;

export type CardchainMsgApointMatchReporterResponse = object;

export type CardchainMsgBuyCardSchemeResponse = object;

export type CardchainMsgChangeArtistResponse = object;

export type CardchainMsgCreateCollectionResponse = object;

export type CardchainMsgCreateuserResponse = object;

export type CardchainMsgDonateToCardResponse = object;

export type CardchainMsgRegisterForCouncilResponse = object;

export interface CardchainMsgReportMatchResponse {
  /** @format uint64 */
  matchId?: string;
}

export type CardchainMsgSaveCardContentResponse = object;

export type CardchainMsgSubmitCopyrightProposalResponse = object;

export type CardchainMsgSubmitMatchReporterProposalResponse = object;

export type CardchainMsgTransferCardResponse = object;

export type CardchainMsgVoteCardResponse = object;

export enum CardchainOutcome {
  AWon = "AWon",
  BWon = "BWon",
  Draw = "Draw",
  Aborted = "Aborted",
}

/**
 * Params defines the parameters for the module.
 */
export type CardchainParams = object;

/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface CardchainQueryParamsResponse {
  /** params holds all the parameters of this module. */
  params?: CardchainParams;
}

export interface CardchainQueryQCardContentResponse {
  content?: string;
}

export interface CardchainQueryQCardResponse {
  owner?: string;
  content?: string;
  artist?: string;
  image?: string;
  fullArt?: boolean;
  notes?: string;
  status?: CardchaincardchainStatus;
  votePool?: string;

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

export interface CardchainQueryQCardchainInfoResponse {
  cardAuctionPrice?: string;
}

export interface CardchainQueryQCardsResponse {
  cardsList?: string[];
}

export interface CardchainQueryQMatchResponse {
  /** @format uint64 */
  timestamp?: string;
  reporter?: string;
  playerA?: string;
  playerB?: string;
  outcome?: CardchainOutcome;
}

export interface CardchainQueryQUserResponse {
  alias?: string;
  ownedCardSchemes?: string[];
  ownedCards?: string[];
  voteRights?: CardchainVoteRight[];
  councilStatus?: CardchainCouncilStatus;
  reportMatches?: boolean;
}

export interface CardchainQueryQVotableCardsResponse {
  unregistered?: boolean;
  noVoteRights?: boolean;
  voteRights?: CardchainVoteRight[];
}

export interface CardchainQueryQVotingResultsResponse {
  lastVotingResults?: CardchainVotingResults;
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

export type QueryParamsType = Record<string | number, any>;
export type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;

export interface FullRequestParams extends Omit<RequestInit, "body"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: keyof Omit<Body, "body" | "bodyUsed">;
  /** request body */
  body?: unknown;
  /** base url */
  baseUrl?: string;
  /** request cancellation token */
  cancelToken?: CancelToken;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> {
  baseUrl?: string;
  baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
  securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}

export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
  data: D;
  error: E;
}

type CancelToken = Symbol | string | number;

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public baseUrl: string = "";
  private securityData: SecurityDataType = null as any;
  private securityWorker: null | ApiConfig<SecurityDataType>["securityWorker"] = null;
  private abortControllers = new Map<CancelToken, AbortController>();

  private baseApiParams: RequestParams = {
    credentials: "same-origin",
    headers: {},
    redirect: "follow",
    referrerPolicy: "no-referrer",
  };

  constructor(apiConfig: ApiConfig<SecurityDataType> = {}) {
    Object.assign(this, apiConfig);
  }

  public setSecurityData = (data: SecurityDataType) => {
    this.securityData = data;
  };

  private addQueryParam(query: QueryParamsType, key: string) {
    const value = query[key];

    return (
      encodeURIComponent(key) +
      "=" +
      encodeURIComponent(Array.isArray(value) ? value.join(",") : typeof value === "number" ? value : `${value}`)
    );
  }

  protected toQueryString(rawQuery?: QueryParamsType): string {
    const query = rawQuery || {};
    const keys = Object.keys(query).filter((key) => "undefined" !== typeof query[key]);
    return keys
      .map((key) =>
        typeof query[key] === "object" && !Array.isArray(query[key])
          ? this.toQueryString(query[key] as QueryParamsType)
          : this.addQueryParam(query, key),
      )
      .join("&");
  }

  protected addQueryParams(rawQuery?: QueryParamsType): string {
    const queryString = this.toQueryString(rawQuery);
    return queryString ? `?${queryString}` : "";
  }

  private contentFormatters: Record<ContentType, (input: any) => any> = {
    [ContentType.Json]: (input: any) =>
      input !== null && (typeof input === "object" || typeof input === "string") ? JSON.stringify(input) : input,
    [ContentType.FormData]: (input: any) =>
      Object.keys(input || {}).reduce((data, key) => {
        data.append(key, input[key]);
        return data;
      }, new FormData()),
    [ContentType.UrlEncoded]: (input: any) => this.toQueryString(input),
  };

  private mergeRequestParams(params1: RequestParams, params2?: RequestParams): RequestParams {
    return {
      ...this.baseApiParams,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.baseApiParams.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createAbortSignal = (cancelToken: CancelToken): AbortSignal | undefined => {
    if (this.abortControllers.has(cancelToken)) {
      const abortController = this.abortControllers.get(cancelToken);
      if (abortController) {
        return abortController.signal;
      }
      return void 0;
    }

    const abortController = new AbortController();
    this.abortControllers.set(cancelToken, abortController);
    return abortController.signal;
  };

  public abortRequest = (cancelToken: CancelToken) => {
    const abortController = this.abortControllers.get(cancelToken);

    if (abortController) {
      abortController.abort();
      this.abortControllers.delete(cancelToken);
    }
  };

  public request = <T = any, E = any>({
    body,
    secure,
    path,
    type,
    query,
    format = "json",
    baseUrl,
    cancelToken,
    ...params
  }: FullRequestParams): Promise<HttpResponse<T, E>> => {
    const secureParams = (secure && this.securityWorker && this.securityWorker(this.securityData)) || {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const queryString = query && this.toQueryString(query);
    const payloadFormatter = this.contentFormatters[type || ContentType.Json];

    return fetch(`${baseUrl || this.baseUrl || ""}${path}${queryString ? `?${queryString}` : ""}`, {
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      signal: cancelToken ? this.createAbortSignal(cancelToken) : void 0,
      body: typeof body === "undefined" || body === null ? null : payloadFormatter(body),
    }).then(async (response) => {
      const r = response as HttpResponse<T, E>;
      r.data = (null as unknown) as T;
      r.error = (null as unknown) as E;

      const data = await response[format]()
        .then((data) => {
          if (r.ok) {
            r.data = data;
          } else {
            r.error = data;
          }
          return r;
        })
        .catch((e) => {
          r.error = e;
          return r;
        });

      if (cancelToken) {
        this.abortControllers.delete(cancelToken);
      }

      if (!response.ok) throw data;
      return data;
    });
  };
}

/**
 * @title cardchain/card.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
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
    this.request<CardchainQueryQCardResponse, GooglerpcStatus>({
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
   * @request GET:/DecentralCardGame/cardchain/cardchain/q_cards/{owner}/{status}/{cardType}/{classes}/{sortBy}/{nameContains}/{keywordsContains}/{notesContains}
   */
  queryQCards = (
    owner: string,
    status:
      | "scheme"
      | "prototype"
      | "trial"
      | "permanent"
      | "suspended"
      | "banned"
      | "bannedSoon"
      | "bannedVerySoon"
      | "none",
    cardType: string,
    classes: string,
    sortBy: string,
    nameContains: string,
    keywordsContains: string,
    notesContains: string,
    params: RequestParams = {},
  ) =>
    this.request<CardchainQueryQCardsResponse, GooglerpcStatus>({
      path: `/DecentralCardGame/cardchain/cardchain/q_cards/${owner}/${status}/${cardType}/${classes}/${sortBy}/${nameContains}/${keywordsContains}/${notesContains}`,
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
    this.request<CardchainQueryQMatchResponse, GooglerpcStatus>({
      path: `/DecentralCardGame/cardchain/cardchain/q_match/${matchId}`,
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
    this.request<CardchainQueryQUserResponse, GooglerpcStatus>({
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
