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
exports.Api = exports.HttpClient = exports.ContentType = void 0;
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
 * @title ibc/core/client/v1/client.proto
 * @version version not set
 */
class Api extends HttpClient {
    constructor() {
        super(...arguments);
        /**
         * No description
         *
         * @tags Query
         * @name QueryClientParams
         * @summary ClientParams queries all parameters of the ibc client.
         * @request GET:/ibc/client/v1/params
         */
        this.queryClientParams = (params = {}) => this.request({
            path: `/ibc/client/v1/params`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryClientStates
         * @summary ClientStates queries all the IBC light clients of a chain.
         * @request GET:/ibc/core/client/v1/client_states
         */
        this.queryClientStates = (query, params = {}) => this.request({
            path: `/ibc/core/client/v1/client_states`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryClientState
         * @summary ClientState queries an IBC light client.
         * @request GET:/ibc/core/client/v1/client_states/{client_id}
         */
        this.queryClientState = (clientId, params = {}) => this.request({
            path: `/ibc/core/client/v1/client_states/${clientId}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryClientStatus
         * @summary Status queries the status of an IBC client.
         * @request GET:/ibc/core/client/v1/client_status/{client_id}
         */
        this.queryClientStatus = (clientId, params = {}) => this.request({
            path: `/ibc/core/client/v1/client_status/${clientId}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
       * No description
       *
       * @tags Query
       * @name QueryConsensusStates
       * @summary ConsensusStates queries all the consensus state associated with a given
      client.
       * @request GET:/ibc/core/client/v1/consensus_states/{client_id}
       */
        this.queryConsensusStates = (clientId, query, params = {}) => this.request({
            path: `/ibc/core/client/v1/consensus_states/${clientId}`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryConsensusStateHeights
         * @summary ConsensusStateHeights queries the height of every consensus states associated with a given client.
         * @request GET:/ibc/core/client/v1/consensus_states/{client_id}/heights
         */
        this.queryConsensusStateHeights = (clientId, query, params = {}) => this.request({
            path: `/ibc/core/client/v1/consensus_states/${clientId}/heights`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
       * No description
       *
       * @tags Query
       * @name QueryConsensusState
       * @summary ConsensusState queries a consensus state associated with a client state at
      a given height.
       * @request GET:/ibc/core/client/v1/consensus_states/{client_id}/revision/{revision_number}/height/{revision_height}
       */
        this.queryConsensusState = (clientId, revisionNumber, revisionHeight, query, params = {}) => this.request({
            path: `/ibc/core/client/v1/consensus_states/${clientId}/revision/${revisionNumber}/height/${revisionHeight}`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryUpgradedClientState
         * @summary UpgradedClientState queries an Upgraded IBC light client.
         * @request GET:/ibc/core/client/v1/upgraded_client_states
         */
        this.queryUpgradedClientState = (params = {}) => this.request({
            path: `/ibc/core/client/v1/upgraded_client_states`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryUpgradedConsensusState
         * @summary UpgradedConsensusState queries an Upgraded IBC consensus state.
         * @request GET:/ibc/core/client/v1/upgraded_consensus_states
         */
        this.queryUpgradedConsensusState = (params = {}) => this.request({
            path: `/ibc/core/client/v1/upgraded_consensus_states`,
            method: "GET",
            format: "json",
            ...params,
        });
    }
}
exports.Api = Api;
