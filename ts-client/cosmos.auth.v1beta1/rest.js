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
 * @title cosmos/auth/v1beta1/auth.proto
 * @version version not set
 */
class Api extends HttpClient {
    constructor() {
        super(...arguments);
        /**
         * @description Since: cosmos-sdk 0.43
         *
         * @tags Query
         * @name QueryAccounts
         * @summary Accounts returns all the existing accounts
         * @request GET:/cosmos/auth/v1beta1/accounts
         */
        this.queryAccounts = (query, params = {}) => this.request({
            path: `/cosmos/auth/v1beta1/accounts`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryAccount
         * @summary Account returns account details based on address.
         * @request GET:/cosmos/auth/v1beta1/accounts/{address}
         */
        this.queryAccount = (address, params = {}) => this.request({
            path: `/cosmos/auth/v1beta1/accounts/${address}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * @description Since: cosmos-sdk 0.46.2
         *
         * @tags Query
         * @name QueryAccountAddressById
         * @summary AccountAddressByID returns account address based on account number.
         * @request GET:/cosmos/auth/v1beta1/address_by_id/{id}
         */
        this.queryAccountAddressById = (id, params = {}) => this.request({
            path: `/cosmos/auth/v1beta1/address_by_id/${id}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * @description Since: cosmos-sdk 0.46
         *
         * @tags Query
         * @name QueryBech32Prefix
         * @summary Bech32Prefix queries bech32Prefix
         * @request GET:/cosmos/auth/v1beta1/bech32
         */
        this.queryBech32Prefix = (params = {}) => this.request({
            path: `/cosmos/auth/v1beta1/bech32`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * @description Since: cosmos-sdk 0.46
         *
         * @tags Query
         * @name QueryAddressBytesToString
         * @summary AddressBytesToString converts Account Address bytes to string
         * @request GET:/cosmos/auth/v1beta1/bech32/{address_bytes}
         */
        this.queryAddressBytesToString = (addressBytes, params = {}) => this.request({
            path: `/cosmos/auth/v1beta1/bech32/${addressBytes}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * @description Since: cosmos-sdk 0.46
         *
         * @tags Query
         * @name QueryAddressStringToBytes
         * @summary AddressStringToBytes converts Address string to bytes
         * @request GET:/cosmos/auth/v1beta1/bech32/{address_string}
         */
        this.queryAddressStringToBytes = (addressString, params = {}) => this.request({
            path: `/cosmos/auth/v1beta1/bech32/${addressString}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * @description Since: cosmos-sdk 0.46
         *
         * @tags Query
         * @name QueryModuleAccounts
         * @summary ModuleAccounts returns all the existing module accounts.
         * @request GET:/cosmos/auth/v1beta1/module_accounts
         */
        this.queryModuleAccounts = (params = {}) => this.request({
            path: `/cosmos/auth/v1beta1/module_accounts`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryParams
         * @summary Params queries all parameters.
         * @request GET:/cosmos/auth/v1beta1/params
         */
        this.queryParams = (params = {}) => this.request({
            path: `/cosmos/auth/v1beta1/params`,
            method: "GET",
            format: "json",
            ...params,
        });
    }
}
exports.Api = Api;
