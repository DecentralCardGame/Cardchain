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
exports.Api = exports.HttpClient = exports.ContentType = exports.V1Beta1SignMode = exports.V1Beta1OrderBy = exports.V1Beta1BroadcastMode = exports.TypesSignedMsgType = exports.TypesBlockIDFlag = void 0;
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
/**
* BroadcastMode specifies the broadcast mode for the TxService.Broadcast RPC method.

 - BROADCAST_MODE_UNSPECIFIED: zero-value for mode ordering
 - BROADCAST_MODE_BLOCK: BROADCAST_MODE_BLOCK defines a tx broadcasting mode where the client waits for
the tx to be committed in a block.
 - BROADCAST_MODE_SYNC: BROADCAST_MODE_SYNC defines a tx broadcasting mode where the client waits for
a CheckTx execution response only.
 - BROADCAST_MODE_ASYNC: BROADCAST_MODE_ASYNC defines a tx broadcasting mode where the client returns
immediately.
*/
var V1Beta1BroadcastMode;
(function (V1Beta1BroadcastMode) {
    V1Beta1BroadcastMode["BROADCAST_MODE_UNSPECIFIED"] = "BROADCAST_MODE_UNSPECIFIED";
    V1Beta1BroadcastMode["BROADCAST_MODE_BLOCK"] = "BROADCAST_MODE_BLOCK";
    V1Beta1BroadcastMode["BROADCAST_MODE_SYNC"] = "BROADCAST_MODE_SYNC";
    V1Beta1BroadcastMode["BROADCAST_MODE_ASYNC"] = "BROADCAST_MODE_ASYNC";
})(V1Beta1BroadcastMode = exports.V1Beta1BroadcastMode || (exports.V1Beta1BroadcastMode = {}));
/**
* - ORDER_BY_UNSPECIFIED: ORDER_BY_UNSPECIFIED specifies an unknown sorting order. OrderBy defaults to ASC in this case.
 - ORDER_BY_ASC: ORDER_BY_ASC defines ascending order
 - ORDER_BY_DESC: ORDER_BY_DESC defines descending order
*/
var V1Beta1OrderBy;
(function (V1Beta1OrderBy) {
    V1Beta1OrderBy["ORDER_BY_UNSPECIFIED"] = "ORDER_BY_UNSPECIFIED";
    V1Beta1OrderBy["ORDER_BY_ASC"] = "ORDER_BY_ASC";
    V1Beta1OrderBy["ORDER_BY_DESC"] = "ORDER_BY_DESC";
})(V1Beta1OrderBy = exports.V1Beta1OrderBy || (exports.V1Beta1OrderBy = {}));
/**
* SignMode represents a signing mode with its own security guarantees.

This enum should be considered a registry of all known sign modes
in the Cosmos ecosystem. Apps are not expected to support all known
sign modes. Apps that would like to support custom  sign modes are
encouraged to open a small PR against this file to add a new case
to this SignMode enum describing their sign mode so that different
apps have a consistent version of this enum.

 - SIGN_MODE_UNSPECIFIED: SIGN_MODE_UNSPECIFIED specifies an unknown signing mode and will be
rejected.
 - SIGN_MODE_DIRECT: SIGN_MODE_DIRECT specifies a signing mode which uses SignDoc and is
verified with raw bytes from Tx.
 - SIGN_MODE_TEXTUAL: SIGN_MODE_TEXTUAL is a future signing mode that will verify some
human-readable textual representation on top of the binary representation
from SIGN_MODE_DIRECT. It is currently not supported.
 - SIGN_MODE_DIRECT_AUX: SIGN_MODE_DIRECT_AUX specifies a signing mode which uses
SignDocDirectAux. As opposed to SIGN_MODE_DIRECT, this sign mode does not
require signers signing over other signers' `signer_info`. It also allows
for adding Tips in transactions.

Since: cosmos-sdk 0.46
 - SIGN_MODE_LEGACY_AMINO_JSON: SIGN_MODE_LEGACY_AMINO_JSON is a backwards compatibility mode which uses
Amino JSON and will be removed in the future.
 - SIGN_MODE_EIP_191: SIGN_MODE_EIP_191 specifies the sign mode for EIP 191 signing on the Cosmos
SDK. Ref: https://eips.ethereum.org/EIPS/eip-191

Currently, SIGN_MODE_EIP_191 is registered as a SignMode enum variant,
but is not implemented on the SDK by default. To enable EIP-191, you need
to pass a custom `TxConfig` that has an implementation of
`SignModeHandler` for EIP-191. The SDK may decide to fully support
EIP-191 in the future.

Since: cosmos-sdk 0.45.2
*/
var V1Beta1SignMode;
(function (V1Beta1SignMode) {
    V1Beta1SignMode["SIGN_MODE_UNSPECIFIED"] = "SIGN_MODE_UNSPECIFIED";
    V1Beta1SignMode["SIGN_MODE_DIRECT"] = "SIGN_MODE_DIRECT";
    V1Beta1SignMode["SIGN_MODE_TEXTUAL"] = "SIGN_MODE_TEXTUAL";
    V1Beta1SignMode["SIGN_MODE_DIRECT_AUX"] = "SIGN_MODE_DIRECT_AUX";
    V1Beta1SignMode["SIGN_MODE_LEGACY_AMINO_JSON"] = "SIGN_MODE_LEGACY_AMINO_JSON";
    V1Beta1SignMode["SIGNMODEEIP191"] = "SIGN_MODE_EIP_191";
})(V1Beta1SignMode = exports.V1Beta1SignMode || (exports.V1Beta1SignMode = {}));
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
 * @title cosmos/tx/v1beta1/service.proto
 * @version version not set
 */
class Api extends HttpClient {
    constructor() {
        super(...arguments);
        /**
         * No description
         *
         * @tags Service
         * @name ServiceSimulate
         * @summary Simulate simulates executing a transaction for estimating gas usage.
         * @request POST:/cosmos/tx/v1beta1/simulate
         */
        this.serviceSimulate = (body, params = {}) => this.request({
            path: `/cosmos/tx/v1beta1/simulate`,
            method: "POST",
            body: body,
            type: ContentType.Json,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Service
         * @name ServiceGetTxsEvent
         * @summary GetTxsEvent fetches txs by event.
         * @request GET:/cosmos/tx/v1beta1/txs
         */
        this.serviceGetTxsEvent = (query, params = {}) => this.request({
            path: `/cosmos/tx/v1beta1/txs`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Service
         * @name ServiceBroadcastTx
         * @summary BroadcastTx broadcast transaction.
         * @request POST:/cosmos/tx/v1beta1/txs
         */
        this.serviceBroadcastTx = (body, params = {}) => this.request({
            path: `/cosmos/tx/v1beta1/txs`,
            method: "POST",
            body: body,
            type: ContentType.Json,
            format: "json",
            ...params,
        });
        /**
         * @description Since: cosmos-sdk 0.45.2
         *
         * @tags Service
         * @name ServiceGetBlockWithTxs
         * @summary GetBlockWithTxs fetches a block with decoded txs.
         * @request GET:/cosmos/tx/v1beta1/txs/block/{height}
         */
        this.serviceGetBlockWithTxs = (height, query, params = {}) => this.request({
            path: `/cosmos/tx/v1beta1/txs/block/${height}`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Service
         * @name ServiceGetTx
         * @summary GetTx fetches a tx by hash.
         * @request GET:/cosmos/tx/v1beta1/txs/{hash}
         */
        this.serviceGetTx = (hash, params = {}) => this.request({
            path: `/cosmos/tx/v1beta1/txs/${hash}`,
            method: "GET",
            format: "json",
            ...params,
        });
    }
}
exports.Api = Api;
