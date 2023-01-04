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
exports.Api = exports.HttpClient = exports.ContentType = exports.V1Beta1BondStatus = void 0;
/**
* BondStatus is the status of a validator.

 - BOND_STATUS_UNSPECIFIED: UNSPECIFIED defines an invalid validator status.
 - BOND_STATUS_UNBONDED: UNBONDED defines a validator that is not bonded.
 - BOND_STATUS_UNBONDING: UNBONDING defines a validator that is unbonding.
 - BOND_STATUS_BONDED: BONDED defines a validator that is bonded.
*/
var V1Beta1BondStatus;
(function (V1Beta1BondStatus) {
    V1Beta1BondStatus["BOND_STATUS_UNSPECIFIED"] = "BOND_STATUS_UNSPECIFIED";
    V1Beta1BondStatus["BOND_STATUS_UNBONDED"] = "BOND_STATUS_UNBONDED";
    V1Beta1BondStatus["BOND_STATUS_UNBONDING"] = "BOND_STATUS_UNBONDING";
    V1Beta1BondStatus["BOND_STATUS_BONDED"] = "BOND_STATUS_BONDED";
})(V1Beta1BondStatus = exports.V1Beta1BondStatus || (exports.V1Beta1BondStatus = {}));
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
 * @title cosmos/staking/v1beta1/authz.proto
 * @version version not set
 */
class Api extends HttpClient {
    constructor() {
        super(...arguments);
        /**
         * No description
         *
         * @tags Query
         * @name QueryDelegatorDelegations
         * @summary DelegatorDelegations queries all delegations of a given delegator address.
         * @request GET:/cosmos/staking/v1beta1/delegations/{delegator_addr}
         */
        this.queryDelegatorDelegations = (delegatorAddr, query, params = {}) => this.request({
            path: `/cosmos/staking/v1beta1/delegations/${delegatorAddr}`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryRedelegations
         * @summary Redelegations queries redelegations of given address.
         * @request GET:/cosmos/staking/v1beta1/delegators/{delegator_addr}/redelegations
         */
        this.queryRedelegations = (delegatorAddr, query, params = {}) => this.request({
            path: `/cosmos/staking/v1beta1/delegators/${delegatorAddr}/redelegations`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
       * No description
       *
       * @tags Query
       * @name QueryDelegatorUnbondingDelegations
       * @summary DelegatorUnbondingDelegations queries all unbonding delegations of a given
      delegator address.
       * @request GET:/cosmos/staking/v1beta1/delegators/{delegator_addr}/unbonding_delegations
       */
        this.queryDelegatorUnbondingDelegations = (delegatorAddr, query, params = {}) => this.request({
            path: `/cosmos/staking/v1beta1/delegators/${delegatorAddr}/unbonding_delegations`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
       * No description
       *
       * @tags Query
       * @name QueryDelegatorValidators
       * @summary DelegatorValidators queries all validators info for given delegator
      address.
       * @request GET:/cosmos/staking/v1beta1/delegators/{delegator_addr}/validators
       */
        this.queryDelegatorValidators = (delegatorAddr, query, params = {}) => this.request({
            path: `/cosmos/staking/v1beta1/delegators/${delegatorAddr}/validators`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
       * No description
       *
       * @tags Query
       * @name QueryDelegatorValidator
       * @summary DelegatorValidator queries validator info for given delegator validator
      pair.
       * @request GET:/cosmos/staking/v1beta1/delegators/{delegator_addr}/validators/{validator_addr}
       */
        this.queryDelegatorValidator = (delegatorAddr, validatorAddr, params = {}) => this.request({
            path: `/cosmos/staking/v1beta1/delegators/${delegatorAddr}/validators/${validatorAddr}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryHistoricalInfo
         * @summary HistoricalInfo queries the historical info for given height.
         * @request GET:/cosmos/staking/v1beta1/historical_info/{height}
         */
        this.queryHistoricalInfo = (height, params = {}) => this.request({
            path: `/cosmos/staking/v1beta1/historical_info/${height}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryParams
         * @summary Parameters queries the staking parameters.
         * @request GET:/cosmos/staking/v1beta1/params
         */
        this.queryParams = (params = {}) => this.request({
            path: `/cosmos/staking/v1beta1/params`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryPool
         * @summary Pool queries the pool info.
         * @request GET:/cosmos/staking/v1beta1/pool
         */
        this.queryPool = (params = {}) => this.request({
            path: `/cosmos/staking/v1beta1/pool`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryValidators
         * @summary Validators queries all validators that match the given status.
         * @request GET:/cosmos/staking/v1beta1/validators
         */
        this.queryValidators = (query, params = {}) => this.request({
            path: `/cosmos/staking/v1beta1/validators`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryValidator
         * @summary Validator queries validator info for given validator address.
         * @request GET:/cosmos/staking/v1beta1/validators/{validator_addr}
         */
        this.queryValidator = (validatorAddr, params = {}) => this.request({
            path: `/cosmos/staking/v1beta1/validators/${validatorAddr}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryValidatorDelegations
         * @summary ValidatorDelegations queries delegate info for given validator.
         * @request GET:/cosmos/staking/v1beta1/validators/{validator_addr}/delegations
         */
        this.queryValidatorDelegations = (validatorAddr, query, params = {}) => this.request({
            path: `/cosmos/staking/v1beta1/validators/${validatorAddr}/delegations`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryDelegation
         * @summary Delegation queries delegate info for given validator delegator pair.
         * @request GET:/cosmos/staking/v1beta1/validators/{validator_addr}/delegations/{delegator_addr}
         */
        this.queryDelegation = (validatorAddr, delegatorAddr, params = {}) => this.request({
            path: `/cosmos/staking/v1beta1/validators/${validatorAddr}/delegations/${delegatorAddr}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
       * No description
       *
       * @tags Query
       * @name QueryUnbondingDelegation
       * @summary UnbondingDelegation queries unbonding info for given validator delegator
      pair.
       * @request GET:/cosmos/staking/v1beta1/validators/{validator_addr}/delegations/{delegator_addr}/unbonding_delegation
       */
        this.queryUnbondingDelegation = (validatorAddr, delegatorAddr, params = {}) => this.request({
            path: `/cosmos/staking/v1beta1/validators/${validatorAddr}/delegations/${delegatorAddr}/unbonding_delegation`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryValidatorUnbondingDelegations
         * @summary ValidatorUnbondingDelegations queries unbonding delegations of a validator.
         * @request GET:/cosmos/staking/v1beta1/validators/{validator_addr}/unbonding_delegations
         */
        this.queryValidatorUnbondingDelegations = (validatorAddr, query, params = {}) => this.request({
            path: `/cosmos/staking/v1beta1/validators/${validatorAddr}/unbonding_delegations`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
    }
}
exports.Api = Api;
