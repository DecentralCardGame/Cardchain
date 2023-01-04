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
 * @title cosmos/bank/v1beta1/authz.proto
 * @version version not set
 */
class Api extends HttpClient {
    constructor() {
        super(...arguments);
        /**
         * No description
         *
         * @tags Query
         * @name QueryAllBalances
         * @summary AllBalances queries the balance of all coins for a single account.
         * @request GET:/cosmos/bank/v1beta1/balances/{address}
         */
        this.queryAllBalances = (address, query, params = {}) => this.request({
            path: `/cosmos/bank/v1beta1/balances/${address}`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryBalance
         * @summary Balance queries the balance of a single coin for a single account.
         * @request GET:/cosmos/bank/v1beta1/balances/{address}/by_denom
         */
        this.queryBalance = (address, query, params = {}) => this.request({
            path: `/cosmos/bank/v1beta1/balances/${address}/by_denom`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
       * @description Since: cosmos-sdk 0.46
       *
       * @tags Query
       * @name QueryDenomOwners
       * @summary DenomOwners queries for all account addresses that own a particular token
      denomination.
       * @request GET:/cosmos/bank/v1beta1/denom_owners/{denom}
       */
        this.queryDenomOwners = (denom, query, params = {}) => this.request({
            path: `/cosmos/bank/v1beta1/denom_owners/${denom}`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
       * No description
       *
       * @tags Query
       * @name QueryDenomsMetadata
       * @summary DenomsMetadata queries the client metadata for all registered coin
      denominations.
       * @request GET:/cosmos/bank/v1beta1/denoms_metadata
       */
        this.queryDenomsMetadata = (query, params = {}) => this.request({
            path: `/cosmos/bank/v1beta1/denoms_metadata`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryDenomMetadata
         * @summary DenomsMetadata queries the client metadata of a given coin denomination.
         * @request GET:/cosmos/bank/v1beta1/denoms_metadata/{denom}
         */
        this.queryDenomMetadata = (denom, params = {}) => this.request({
            path: `/cosmos/bank/v1beta1/denoms_metadata/${denom}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryParams
         * @summary Params queries the parameters of x/bank module.
         * @request GET:/cosmos/bank/v1beta1/params
         */
        this.queryParams = (params = {}) => this.request({
            path: `/cosmos/bank/v1beta1/params`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
       * @description Since: cosmos-sdk 0.46
       *
       * @tags Query
       * @name QuerySpendableBalances
       * @summary SpendableBalances queries the spenable balance of all coins for a single
      account.
       * @request GET:/cosmos/bank/v1beta1/spendable_balances/{address}
       */
        this.querySpendableBalances = (address, query, params = {}) => this.request({
            path: `/cosmos/bank/v1beta1/spendable_balances/${address}`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryTotalSupply
         * @summary TotalSupply queries the total supply of all coins.
         * @request GET:/cosmos/bank/v1beta1/supply
         */
        this.queryTotalSupply = (query, params = {}) => this.request({
            path: `/cosmos/bank/v1beta1/supply`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QuerySupplyOf
         * @summary SupplyOf queries the supply of a single coin.
         * @request GET:/cosmos/bank/v1beta1/supply/by_denom
         */
        this.querySupplyOf = (query, params = {}) => this.request({
            path: `/cosmos/bank/v1beta1/supply/by_denom`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
    }
}
exports.Api = Api;
