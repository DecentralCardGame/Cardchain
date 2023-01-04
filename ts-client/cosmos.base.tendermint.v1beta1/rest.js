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
exports.Api = exports.HttpClient = exports.ContentType = exports.TypesSignedMsgType = exports.TypesBlockIDFlag = void 0;
var TypesBlockIDFlag;
(function (TypesBlockIDFlag) {
    TypesBlockIDFlag["BLOCK_ID_FLAG_UNKNOWN"] = "BLOCK_ID_FLAG_UNKNOWN";
    TypesBlockIDFlag["BLOCK_ID_FLAG_ABSENT"] = "BLOCK_ID_FLAG_ABSENT";
    TypesBlockIDFlag["BLOCK_ID_FLAG_COMMIT"] = "BLOCK_ID_FLAG_COMMIT";
    TypesBlockIDFlag["BLOCK_ID_FLAG_NIL"] = "BLOCK_ID_FLAG_NIL";
})(TypesBlockIDFlag = exports.TypesBlockIDFlag || (exports.TypesBlockIDFlag = {}));
/**
* SignedMsgType is a type of signed message in the consensus.

 - SIGNED_MSG_TYPE_PREVOTE: Votes
 - SIGNED_MSG_TYPE_PROPOSAL: Proposals
*/
var TypesSignedMsgType;
(function (TypesSignedMsgType) {
    TypesSignedMsgType["SIGNED_MSG_TYPE_UNKNOWN"] = "SIGNED_MSG_TYPE_UNKNOWN";
    TypesSignedMsgType["SIGNED_MSG_TYPE_PREVOTE"] = "SIGNED_MSG_TYPE_PREVOTE";
    TypesSignedMsgType["SIGNED_MSG_TYPE_PRECOMMIT"] = "SIGNED_MSG_TYPE_PRECOMMIT";
    TypesSignedMsgType["SIGNED_MSG_TYPE_PROPOSAL"] = "SIGNED_MSG_TYPE_PROPOSAL";
})(TypesSignedMsgType = exports.TypesSignedMsgType || (exports.TypesSignedMsgType = {}));
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
 * @title cosmos/base/tendermint/v1beta1/query.proto
 * @version version not set
 */
class Api extends HttpClient {
    constructor() {
        super(...arguments);
        /**
       * @description Since: cosmos-sdk 0.46
       *
       * @tags Service
       * @name ServiceAbciQuery
       * @summary ABCIQuery defines a query handler that supports ABCI queries directly to
      the application, bypassing Tendermint completely. The ABCI query must
      contain a valid and supported path, including app, custom, p2p, and store.
       * @request GET:/cosmos/base/tendermint/v1beta1/abci_query
       */
        this.serviceAbciQuery = (query, params = {}) => this.request({
            path: `/cosmos/base/tendermint/v1beta1/abci_query`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Service
         * @name ServiceGetLatestBlock
         * @summary GetLatestBlock returns the latest block.
         * @request GET:/cosmos/base/tendermint/v1beta1/blocks/latest
         */
        this.serviceGetLatestBlock = (params = {}) => this.request({
            path: `/cosmos/base/tendermint/v1beta1/blocks/latest`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Service
         * @name ServiceGetBlockByHeight
         * @summary GetBlockByHeight queries block for given height.
         * @request GET:/cosmos/base/tendermint/v1beta1/blocks/{height}
         */
        this.serviceGetBlockByHeight = (height, params = {}) => this.request({
            path: `/cosmos/base/tendermint/v1beta1/blocks/${height}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Service
         * @name ServiceGetNodeInfo
         * @summary GetNodeInfo queries the current node info.
         * @request GET:/cosmos/base/tendermint/v1beta1/node_info
         */
        this.serviceGetNodeInfo = (params = {}) => this.request({
            path: `/cosmos/base/tendermint/v1beta1/node_info`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Service
         * @name ServiceGetSyncing
         * @summary GetSyncing queries node syncing.
         * @request GET:/cosmos/base/tendermint/v1beta1/syncing
         */
        this.serviceGetSyncing = (params = {}) => this.request({
            path: `/cosmos/base/tendermint/v1beta1/syncing`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Service
         * @name ServiceGetLatestValidatorSet
         * @summary GetLatestValidatorSet queries latest validator-set.
         * @request GET:/cosmos/base/tendermint/v1beta1/validatorsets/latest
         */
        this.serviceGetLatestValidatorSet = (query, params = {}) => this.request({
            path: `/cosmos/base/tendermint/v1beta1/validatorsets/latest`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Service
         * @name ServiceGetValidatorSetByHeight
         * @summary GetValidatorSetByHeight queries validator-set at a given height.
         * @request GET:/cosmos/base/tendermint/v1beta1/validatorsets/{height}
         */
        this.serviceGetValidatorSetByHeight = (height, query, params = {}) => this.request({
            path: `/cosmos/base/tendermint/v1beta1/validatorsets/${height}`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
    }
}
exports.Api = Api;
