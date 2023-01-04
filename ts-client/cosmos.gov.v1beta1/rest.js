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
exports.Api = exports.HttpClient = exports.ContentType = exports.V1Beta1VoteOption = exports.V1Beta1ProposalStatus = void 0;
/**
* ProposalStatus enumerates the valid statuses of a proposal.

 - PROPOSAL_STATUS_UNSPECIFIED: PROPOSAL_STATUS_UNSPECIFIED defines the default proposal status.
 - PROPOSAL_STATUS_DEPOSIT_PERIOD: PROPOSAL_STATUS_DEPOSIT_PERIOD defines a proposal status during the deposit
period.
 - PROPOSAL_STATUS_VOTING_PERIOD: PROPOSAL_STATUS_VOTING_PERIOD defines a proposal status during the voting
period.
 - PROPOSAL_STATUS_PASSED: PROPOSAL_STATUS_PASSED defines a proposal status of a proposal that has
passed.
 - PROPOSAL_STATUS_REJECTED: PROPOSAL_STATUS_REJECTED defines a proposal status of a proposal that has
been rejected.
 - PROPOSAL_STATUS_FAILED: PROPOSAL_STATUS_FAILED defines a proposal status of a proposal that has
failed.
*/
var V1Beta1ProposalStatus;
(function (V1Beta1ProposalStatus) {
    V1Beta1ProposalStatus["PROPOSAL_STATUS_UNSPECIFIED"] = "PROPOSAL_STATUS_UNSPECIFIED";
    V1Beta1ProposalStatus["PROPOSAL_STATUS_DEPOSIT_PERIOD"] = "PROPOSAL_STATUS_DEPOSIT_PERIOD";
    V1Beta1ProposalStatus["PROPOSAL_STATUS_VOTING_PERIOD"] = "PROPOSAL_STATUS_VOTING_PERIOD";
    V1Beta1ProposalStatus["PROPOSAL_STATUS_PASSED"] = "PROPOSAL_STATUS_PASSED";
    V1Beta1ProposalStatus["PROPOSAL_STATUS_REJECTED"] = "PROPOSAL_STATUS_REJECTED";
    V1Beta1ProposalStatus["PROPOSAL_STATUS_FAILED"] = "PROPOSAL_STATUS_FAILED";
})(V1Beta1ProposalStatus = exports.V1Beta1ProposalStatus || (exports.V1Beta1ProposalStatus = {}));
/**
* VoteOption enumerates the valid vote options for a given governance proposal.

 - VOTE_OPTION_UNSPECIFIED: VOTE_OPTION_UNSPECIFIED defines a no-op vote option.
 - VOTE_OPTION_YES: VOTE_OPTION_YES defines a yes vote option.
 - VOTE_OPTION_ABSTAIN: VOTE_OPTION_ABSTAIN defines an abstain vote option.
 - VOTE_OPTION_NO: VOTE_OPTION_NO defines a no vote option.
 - VOTE_OPTION_NO_WITH_VETO: VOTE_OPTION_NO_WITH_VETO defines a no with veto vote option.
*/
var V1Beta1VoteOption;
(function (V1Beta1VoteOption) {
    V1Beta1VoteOption["VOTE_OPTION_UNSPECIFIED"] = "VOTE_OPTION_UNSPECIFIED";
    V1Beta1VoteOption["VOTE_OPTION_YES"] = "VOTE_OPTION_YES";
    V1Beta1VoteOption["VOTE_OPTION_ABSTAIN"] = "VOTE_OPTION_ABSTAIN";
    V1Beta1VoteOption["VOTE_OPTION_NO"] = "VOTE_OPTION_NO";
    V1Beta1VoteOption["VOTE_OPTION_NO_WITH_VETO"] = "VOTE_OPTION_NO_WITH_VETO";
})(V1Beta1VoteOption = exports.V1Beta1VoteOption || (exports.V1Beta1VoteOption = {}));
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
 * @title cosmos/gov/v1beta1/genesis.proto
 * @version version not set
 */
class Api extends HttpClient {
    constructor() {
        super(...arguments);
        /**
         * No description
         *
         * @tags Query
         * @name QueryParams
         * @summary Params queries all parameters of the gov module.
         * @request GET:/cosmos/gov/v1beta1/params/{params_type}
         */
        this.queryParams = (paramsType, params = {}) => this.request({
            path: `/cosmos/gov/v1beta1/params/${paramsType}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryProposals
         * @summary Proposals queries all proposals based on given status.
         * @request GET:/cosmos/gov/v1beta1/proposals
         */
        this.queryProposals = (query, params = {}) => this.request({
            path: `/cosmos/gov/v1beta1/proposals`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryProposal
         * @summary Proposal queries proposal details based on ProposalID.
         * @request GET:/cosmos/gov/v1beta1/proposals/{proposal_id}
         */
        this.queryProposal = (proposalId, params = {}) => this.request({
            path: `/cosmos/gov/v1beta1/proposals/${proposalId}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryDeposits
         * @summary Deposits queries all deposits of a single proposal.
         * @request GET:/cosmos/gov/v1beta1/proposals/{proposal_id}/deposits
         */
        this.queryDeposits = (proposalId, query, params = {}) => this.request({
            path: `/cosmos/gov/v1beta1/proposals/${proposalId}/deposits`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryDeposit
         * @summary Deposit queries single deposit information based proposalID, depositAddr.
         * @request GET:/cosmos/gov/v1beta1/proposals/{proposal_id}/deposits/{depositor}
         */
        this.queryDeposit = (proposalId, depositor, params = {}) => this.request({
            path: `/cosmos/gov/v1beta1/proposals/${proposalId}/deposits/${depositor}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryTallyResult
         * @summary TallyResult queries the tally of a proposal vote.
         * @request GET:/cosmos/gov/v1beta1/proposals/{proposal_id}/tally
         */
        this.queryTallyResult = (proposalId, params = {}) => this.request({
            path: `/cosmos/gov/v1beta1/proposals/${proposalId}/tally`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryVotes
         * @summary Votes queries votes of a given proposal.
         * @request GET:/cosmos/gov/v1beta1/proposals/{proposal_id}/votes
         */
        this.queryVotes = (proposalId, query, params = {}) => this.request({
            path: `/cosmos/gov/v1beta1/proposals/${proposalId}/votes`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryVote
         * @summary Vote queries voted information based on proposalID, voterAddr.
         * @request GET:/cosmos/gov/v1beta1/proposals/{proposal_id}/votes/{voter}
         */
        this.queryVote = (proposalId, voter, params = {}) => this.request({
            path: `/cosmos/gov/v1beta1/proposals/${proposalId}/votes/${voter}`,
            method: "GET",
            format: "json",
            ...params,
        });
    }
}
exports.Api = Api;
