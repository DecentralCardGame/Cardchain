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
exports.Api = exports.HttpClient = exports.ContentType = exports.V1State = void 0;
/**
* State defines if a connection is in one of the following states:
INIT, TRYOPEN, OPEN or UNINITIALIZED.

 - STATE_UNINITIALIZED_UNSPECIFIED: Default State
 - STATE_INIT: A connection end has just started the opening handshake.
 - STATE_TRYOPEN: A connection end has acknowledged the handshake step on the counterparty
chain.
 - STATE_OPEN: A connection end has completed the handshake.
*/
var V1State;
(function (V1State) {
    V1State["STATE_UNINITIALIZED_UNSPECIFIED"] = "STATE_UNINITIALIZED_UNSPECIFIED";
    V1State["STATE_INIT"] = "STATE_INIT";
    V1State["STATE_TRYOPEN"] = "STATE_TRYOPEN";
    V1State["STATE_OPEN"] = "STATE_OPEN";
})(V1State = exports.V1State || (exports.V1State = {}));
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
 * @title ibc/core/connection/v1/connection.proto
 * @version version not set
 */
class Api extends HttpClient {
    constructor() {
        super(...arguments);
        /**
       * No description
       *
       * @tags Query
       * @name QueryClientConnections
       * @summary ClientConnections queries the connection paths associated with a client
      state.
       * @request GET:/ibc/core/connection/v1/client_connections/{client_id}
       */
        this.queryClientConnections = (clientId, params = {}) => this.request({
            path: `/ibc/core/connection/v1/client_connections/${clientId}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryConnections
         * @summary Connections queries all the IBC connections of a chain.
         * @request GET:/ibc/core/connection/v1/connections
         */
        this.queryConnections = (query, params = {}) => this.request({
            path: `/ibc/core/connection/v1/connections`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryConnection
         * @summary Connection queries an IBC connection end.
         * @request GET:/ibc/core/connection/v1/connections/{connection_id}
         */
        this.queryConnection = (connectionId, params = {}) => this.request({
            path: `/ibc/core/connection/v1/connections/${connectionId}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
       * No description
       *
       * @tags Query
       * @name QueryConnectionClientState
       * @summary ConnectionClientState queries the client state associated with the
      connection.
       * @request GET:/ibc/core/connection/v1/connections/{connection_id}/client_state
       */
        this.queryConnectionClientState = (connectionId, params = {}) => this.request({
            path: `/ibc/core/connection/v1/connections/${connectionId}/client_state`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
       * No description
       *
       * @tags Query
       * @name QueryConnectionConsensusState
       * @summary ConnectionConsensusState queries the consensus state associated with the
      connection.
       * @request GET:/ibc/core/connection/v1/connections/{connection_id}/consensus_state/revision/{revision_number}/height/{revision_height}
       */
        this.queryConnectionConsensusState = (connectionId, revisionNumber, revisionHeight, params = {}) => this.request({
            path: `/ibc/core/connection/v1/connections/${connectionId}/consensus_state/revision/${revisionNumber}/height/${revisionHeight}`,
            method: "GET",
            format: "json",
            ...params,
        });
    }
}
exports.Api = Api;
