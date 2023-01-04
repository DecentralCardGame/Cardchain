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
 * @title cosmos/upgrade/v1beta1/query.proto
 * @version version not set
 */
class Api extends HttpClient {
    constructor() {
        super(...arguments);
        /**
         * No description
         *
         * @tags Query
         * @name QueryAppliedPlan
         * @summary AppliedPlan queries a previously applied upgrade plan by its name.
         * @request GET:/cosmos/upgrade/v1beta1/applied_plan/{name}
         */
        this.queryAppliedPlan = (name, params = {}) => this.request({
            path: `/cosmos/upgrade/v1beta1/applied_plan/${name}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * @description Since: cosmos-sdk 0.46
         *
         * @tags Query
         * @name QueryAuthority
         * @summary Returns the account with authority to conduct upgrades
         * @request GET:/cosmos/upgrade/v1beta1/authority
         */
        this.queryAuthority = (params = {}) => this.request({
            path: `/cosmos/upgrade/v1beta1/authority`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryCurrentPlan
         * @summary CurrentPlan queries the current upgrade plan.
         * @request GET:/cosmos/upgrade/v1beta1/current_plan
         */
        this.queryCurrentPlan = (params = {}) => this.request({
            path: `/cosmos/upgrade/v1beta1/current_plan`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * @description Since: cosmos-sdk 0.43
         *
         * @tags Query
         * @name QueryModuleVersions
         * @summary ModuleVersions queries the list of module versions from state.
         * @request GET:/cosmos/upgrade/v1beta1/module_versions
         */
        this.queryModuleVersions = (query, params = {}) => this.request({
            path: `/cosmos/upgrade/v1beta1/module_versions`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
       * No description
       *
       * @tags Query
       * @name QueryUpgradedConsensusState
       * @summary UpgradedConsensusState queries the consensus state that will serve
      as a trusted kernel for the next version of this chain. It will only be
      stored at the last height of this chain.
      UpgradedConsensusState RPC not supported with legacy querier
      This rpc is deprecated now that IBC has its own replacement
      (https://github.com/cosmos/ibc-go/blob/2c880a22e9f9cc75f62b527ca94aa75ce1106001/proto/ibc/core/client/v1/query.proto#L54)
       * @request GET:/cosmos/upgrade/v1beta1/upgraded_consensus_state/{last_height}
       */
        this.queryUpgradedConsensusState = (lastHeight, params = {}) => this.request({
            path: `/cosmos/upgrade/v1beta1/upgraded_consensus_state/${lastHeight}`,
            method: "GET",
            format: "json",
            ...params,
        });
    }
}
exports.Api = Api;
