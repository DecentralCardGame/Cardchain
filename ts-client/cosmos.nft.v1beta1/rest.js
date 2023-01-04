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
 * @title cosmos/nft/v1beta1/event.proto
 * @version version not set
 */
class Api extends HttpClient {
    constructor() {
        super(...arguments);
        /**
         * No description
         *
         * @tags Query
         * @name QueryBalance
         * @summary Balance queries the number of NFTs of a given class owned by the owner, same as balanceOf in ERC721
         * @request GET:/cosmos/nft/v1beta1/balance/{owner}/{class_id}
         */
        this.queryBalance = (owner, classId, params = {}) => this.request({
            path: `/cosmos/nft/v1beta1/balance/${owner}/${classId}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryClasses
         * @summary Classes queries all NFT classes
         * @request GET:/cosmos/nft/v1beta1/classes
         */
        this.queryClasses = (query, params = {}) => this.request({
            path: `/cosmos/nft/v1beta1/classes`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryClass
         * @summary Class queries an NFT class based on its id
         * @request GET:/cosmos/nft/v1beta1/classes/{class_id}
         */
        this.queryClass = (classId, params = {}) => this.request({
            path: `/cosmos/nft/v1beta1/classes/${classId}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
       * No description
       *
       * @tags Query
       * @name QueryNfTs
       * @summary NFTs queries all NFTs of a given class or owner,choose at least one of the two, similar to tokenByIndex in
      ERC721Enumerable
       * @request GET:/cosmos/nft/v1beta1/nfts
       */
        this.queryNfTs = (query, params = {}) => this.request({
            path: `/cosmos/nft/v1beta1/nfts`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryNft
         * @summary NFT queries an NFT based on its class and id.
         * @request GET:/cosmos/nft/v1beta1/nfts/{class_id}/{id}
         */
        this.queryNft = (classId, id, params = {}) => this.request({
            path: `/cosmos/nft/v1beta1/nfts/${classId}/${id}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryOwner
         * @summary Owner queries the owner of the NFT based on its class and id, same as ownerOf in ERC721
         * @request GET:/cosmos/nft/v1beta1/owner/{class_id}/{id}
         */
        this.queryOwner = (classId, id, params = {}) => this.request({
            path: `/cosmos/nft/v1beta1/owner/${classId}/${id}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QuerySupply
         * @summary Supply queries the number of NFTs from the given class, same as totalSupply of ERC721.
         * @request GET:/cosmos/nft/v1beta1/supply/{class_id}
         */
        this.querySupply = (classId, params = {}) => this.request({
            path: `/cosmos/nft/v1beta1/supply/${classId}`,
            method: "GET",
            format: "json",
            ...params,
        });
    }
}
exports.Api = Api;
