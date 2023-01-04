"use strict";
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
Object.defineProperty(exports, "__esModule", { value: true });
exports.Api = exports.HttpClient = exports.ContentType = exports.CardchaincardchainStatus = exports.CardchainSellOfferStatus = exports.CardchainResponse = exports.CardchainQueryQCardsRequestStatus = exports.CardchainOutcome = exports.CardchainCouncilStatus = exports.CardchainCouncelingStatus = exports.CardchainCStatus = void 0;
var CardchainCStatus;
(function (CardchainCStatus) {
    CardchainCStatus["Design"] = "design";
    CardchainCStatus["Finalized"] = "finalized";
    CardchainCStatus["Active"] = "active";
    CardchainCStatus["Archived"] = "archived";
})(CardchainCStatus = exports.CardchainCStatus || (exports.CardchainCStatus = {}));
var CardchainCouncelingStatus;
(function (CardchainCouncelingStatus) {
    CardchainCouncelingStatus["CouncilOpen"] = "councilOpen";
    CardchainCouncelingStatus["CouncilCreated"] = "councilCreated";
    CardchainCouncelingStatus["CouncilClosed"] = "councilClosed";
    CardchainCouncelingStatus["Commited"] = "commited";
    CardchainCouncelingStatus["Revealed"] = "revealed";
    CardchainCouncelingStatus["SuggestionsMade"] = "suggestionsMade";
})(CardchainCouncelingStatus = exports.CardchainCouncelingStatus || (exports.CardchainCouncelingStatus = {}));
var CardchainCouncilStatus;
(function (CardchainCouncilStatus) {
    CardchainCouncilStatus["Available"] = "available";
    CardchainCouncilStatus["Unavailable"] = "unavailable";
    CardchainCouncilStatus["OpenCouncil"] = "openCouncil";
    CardchainCouncilStatus["StartedCouncil"] = "startedCouncil";
})(CardchainCouncilStatus = exports.CardchainCouncilStatus || (exports.CardchainCouncilStatus = {}));
var CardchainOutcome;
(function (CardchainOutcome) {
    CardchainOutcome["AWon"] = "AWon";
    CardchainOutcome["BWon"] = "BWon";
    CardchainOutcome["Draw"] = "Draw";
    CardchainOutcome["Aborted"] = "Aborted";
})(CardchainOutcome = exports.CardchainOutcome || (exports.CardchainOutcome = {}));
var CardchainQueryQCardsRequestStatus;
(function (CardchainQueryQCardsRequestStatus) {
    CardchainQueryQCardsRequestStatus["Scheme"] = "scheme";
    CardchainQueryQCardsRequestStatus["Prototype"] = "prototype";
    CardchainQueryQCardsRequestStatus["Trial"] = "trial";
    CardchainQueryQCardsRequestStatus["Permanent"] = "permanent";
    CardchainQueryQCardsRequestStatus["Suspended"] = "suspended";
    CardchainQueryQCardsRequestStatus["Banned"] = "banned";
    CardchainQueryQCardsRequestStatus["BannedSoon"] = "bannedSoon";
    CardchainQueryQCardsRequestStatus["BannedVerySoon"] = "bannedVerySoon";
    CardchainQueryQCardsRequestStatus["None"] = "none";
    CardchainQueryQCardsRequestStatus["Playable"] = "playable";
    CardchainQueryQCardsRequestStatus["Unplayable"] = "unplayable";
})(CardchainQueryQCardsRequestStatus = exports.CardchainQueryQCardsRequestStatus || (exports.CardchainQueryQCardsRequestStatus = {}));
var CardchainResponse;
(function (CardchainResponse) {
    CardchainResponse["Yes"] = "Yes";
    CardchainResponse["No"] = "No";
    CardchainResponse["Suggestion"] = "Suggestion";
})(CardchainResponse = exports.CardchainResponse || (exports.CardchainResponse = {}));
var CardchainSellOfferStatus;
(function (CardchainSellOfferStatus) {
    CardchainSellOfferStatus["Open"] = "open";
    CardchainSellOfferStatus["Sold"] = "sold";
    CardchainSellOfferStatus["Removed"] = "removed";
})(CardchainSellOfferStatus = exports.CardchainSellOfferStatus || (exports.CardchainSellOfferStatus = {}));
var CardchaincardchainStatus;
(function (CardchaincardchainStatus) {
    CardchaincardchainStatus["Scheme"] = "scheme";
    CardchaincardchainStatus["Prototype"] = "prototype";
    CardchaincardchainStatus["Trial"] = "trial";
    CardchaincardchainStatus["Permanent"] = "permanent";
    CardchaincardchainStatus["Suspended"] = "suspended";
    CardchaincardchainStatus["Banned"] = "banned";
    CardchaincardchainStatus["BannedSoon"] = "bannedSoon";
    CardchaincardchainStatus["BannedVerySoon"] = "bannedVerySoon";
    CardchaincardchainStatus["None"] = "none";
})(CardchaincardchainStatus = exports.CardchaincardchainStatus || (exports.CardchaincardchainStatus = {}));
const axios_1 = require("axios");
var ContentType;
(function (ContentType) {
    ContentType["Json"] = "application/json";
    ContentType["FormData"] = "multipart/form-data";
    ContentType["UrlEncoded"] = "application/x-www-form-urlencoded";
})(ContentType = exports.ContentType || (exports.ContentType = {}));
class HttpClient {
    constructor({ securityWorker, secure, format, ...axiosConfig } = {}) {
        this.securityData = null;
        this.setSecurityData = (data) => {
            this.securityData = data;
        };
        this.request = async ({ secure, path, type, query, format, body, ...params }) => {
            const secureParams = ((typeof secure === "boolean" ? secure : this.secure) &&
                this.securityWorker &&
                (await this.securityWorker(this.securityData))) ||
                {};
            const requestParams = this.mergeRequestParams(params, secureParams);
            const responseFormat = (format && this.format) || void 0;
            if (type === ContentType.FormData && body && body !== null && typeof body === "object") {
                requestParams.headers.common = { Accept: "*/*" };
                requestParams.headers.post = {};
                requestParams.headers.put = {};
                body = this.createFormData(body);
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
        this.instance = axios_1.default.create({ ...axiosConfig, baseURL: axiosConfig.baseURL || "" });
        this.secure = secure;
        this.format = format;
        this.securityWorker = securityWorker;
    }
    mergeRequestParams(params1, params2) {
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
    createFormData(input) {
        return Object.keys(input || {}).reduce((formData, key) => {
            const property = input[key];
            formData.append(key, property instanceof Blob
                ? property
                : typeof property === "object" && property !== null
                    ? JSON.stringify(property)
                    : `${property}`);
            return formData;
        }, new FormData());
    }
}
exports.HttpClient = HttpClient;
/**
 * @title cardchain/card.proto
 * @version version not set
 */
class Api extends HttpClient {
    constructor() {
        super(...arguments);
        /**
         * No description
         *
         * @tags Query
         * @name QueryQCollections
         * @summary Queries a list of QCollections items.
         * @request GET:/DecentralCardGame/Cardchain/cardchain/q_collections/{status}/{ignoreStatus}
         */
        this.queryQCollections = (status, ignoreStatus, query, params = {}) => this.request({
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
         * @name QueryRarityDistribution
         * @summary Queries a list of RarityDistribution items.
         * @request GET:/DecentralCardGame/Cardchain/cardchain/rarity_distribution/{collectionId}
         */
        this.queryRarityDistribution = (collectionId, params = {}) => this.request({
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
        this.queryParams = (params = {}) => this.request({
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
        this.queryQCard = (cardId, params = {}) => this.request({
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
        this.queryQCardContent = (cardId, params = {}) => this.request({
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
        this.queryQCardchainInfo = (params = {}) => this.request({
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
        this.queryQCards = (status, query, params = {}) => this.request({
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
        this.queryQCollection = (collectionId, params = {}) => this.request({
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
        this.queryQCouncil = (councilId, params = {}) => this.request({
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
        this.queryQMatch = (matchId, params = {}) => this.request({
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
        this.queryQMatches = (query, params = {}) => this.request({
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
        this.queryQSellOffer = (sellOfferId, params = {}) => this.request({
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
        this.queryQSellOffers = (status, query, params = {}) => this.request({
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
        this.queryQServer = (id, params = {}) => this.request({
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
        this.queryQUser = (address, params = {}) => this.request({
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
        this.queryQVotableCards = (address, params = {}) => this.request({
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
        this.queryQVotingResults = (params = {}) => this.request({
            path: `/DecentralCardGame/cardchain/cardchain/q_voting_results`,
            method: "GET",
            format: "json",
            ...params,
        });
    }
}
exports.Api = Api;
