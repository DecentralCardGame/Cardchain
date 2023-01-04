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
exports.Api = exports.HttpClient = exports.ContentType = exports.V1VoteOption = exports.V1ProposalStatus = exports.V1ProposalExecutorResult = exports.V1Exec = void 0;
/**
* Exec defines modes of execution of a proposal on creation or on new vote.

 - EXEC_UNSPECIFIED: An empty value means that there should be a separate
MsgExec request for the proposal to execute.
 - EXEC_TRY: Try to execute the proposal immediately.
If the proposal is not allowed per the DecisionPolicy,
the proposal will still be open and could
be executed at a later point.
*/
var V1Exec;
(function (V1Exec) {
    V1Exec["EXEC_UNSPECIFIED"] = "EXEC_UNSPECIFIED";
    V1Exec["EXEC_TRY"] = "EXEC_TRY";
})(V1Exec = exports.V1Exec || (exports.V1Exec = {}));
/**
* ProposalExecutorResult defines types of proposal executor results.

 - PROPOSAL_EXECUTOR_RESULT_UNSPECIFIED: An empty value is not allowed.
 - PROPOSAL_EXECUTOR_RESULT_NOT_RUN: We have not yet run the executor.
 - PROPOSAL_EXECUTOR_RESULT_SUCCESS: The executor was successful and proposed action updated state.
 - PROPOSAL_EXECUTOR_RESULT_FAILURE: The executor returned an error and proposed action didn't update state.
*/
var V1ProposalExecutorResult;
(function (V1ProposalExecutorResult) {
    V1ProposalExecutorResult["PROPOSAL_EXECUTOR_RESULT_UNSPECIFIED"] = "PROPOSAL_EXECUTOR_RESULT_UNSPECIFIED";
    V1ProposalExecutorResult["PROPOSAL_EXECUTOR_RESULT_NOT_RUN"] = "PROPOSAL_EXECUTOR_RESULT_NOT_RUN";
    V1ProposalExecutorResult["PROPOSAL_EXECUTOR_RESULT_SUCCESS"] = "PROPOSAL_EXECUTOR_RESULT_SUCCESS";
    V1ProposalExecutorResult["PROPOSAL_EXECUTOR_RESULT_FAILURE"] = "PROPOSAL_EXECUTOR_RESULT_FAILURE";
})(V1ProposalExecutorResult = exports.V1ProposalExecutorResult || (exports.V1ProposalExecutorResult = {}));
/**
* ProposalStatus defines proposal statuses.

 - PROPOSAL_STATUS_UNSPECIFIED: An empty value is invalid and not allowed.
 - PROPOSAL_STATUS_SUBMITTED: Initial status of a proposal when submitted.
 - PROPOSAL_STATUS_ACCEPTED: Final status of a proposal when the final tally is done and the outcome
passes the group policy's decision policy.
 - PROPOSAL_STATUS_REJECTED: Final status of a proposal when the final tally is done and the outcome
is rejected by the group policy's decision policy.
 - PROPOSAL_STATUS_ABORTED: Final status of a proposal when the group policy is modified before the
final tally.
 - PROPOSAL_STATUS_WITHDRAWN: A proposal can be withdrawn before the voting start time by the owner.
When this happens the final status is Withdrawn.
*/
var V1ProposalStatus;
(function (V1ProposalStatus) {
    V1ProposalStatus["PROPOSAL_STATUS_UNSPECIFIED"] = "PROPOSAL_STATUS_UNSPECIFIED";
    V1ProposalStatus["PROPOSAL_STATUS_SUBMITTED"] = "PROPOSAL_STATUS_SUBMITTED";
    V1ProposalStatus["PROPOSAL_STATUS_ACCEPTED"] = "PROPOSAL_STATUS_ACCEPTED";
    V1ProposalStatus["PROPOSAL_STATUS_REJECTED"] = "PROPOSAL_STATUS_REJECTED";
    V1ProposalStatus["PROPOSAL_STATUS_ABORTED"] = "PROPOSAL_STATUS_ABORTED";
    V1ProposalStatus["PROPOSAL_STATUS_WITHDRAWN"] = "PROPOSAL_STATUS_WITHDRAWN";
})(V1ProposalStatus = exports.V1ProposalStatus || (exports.V1ProposalStatus = {}));
/**
* VoteOption enumerates the valid vote options for a given proposal.

 - VOTE_OPTION_UNSPECIFIED: VOTE_OPTION_UNSPECIFIED defines an unspecified vote option which will
return an error.
 - VOTE_OPTION_YES: VOTE_OPTION_YES defines a yes vote option.
 - VOTE_OPTION_ABSTAIN: VOTE_OPTION_ABSTAIN defines an abstain vote option.
 - VOTE_OPTION_NO: VOTE_OPTION_NO defines a no vote option.
 - VOTE_OPTION_NO_WITH_VETO: VOTE_OPTION_NO_WITH_VETO defines a no with veto vote option.
*/
var V1VoteOption;
(function (V1VoteOption) {
    V1VoteOption["VOTE_OPTION_UNSPECIFIED"] = "VOTE_OPTION_UNSPECIFIED";
    V1VoteOption["VOTE_OPTION_YES"] = "VOTE_OPTION_YES";
    V1VoteOption["VOTE_OPTION_ABSTAIN"] = "VOTE_OPTION_ABSTAIN";
    V1VoteOption["VOTE_OPTION_NO"] = "VOTE_OPTION_NO";
    V1VoteOption["VOTE_OPTION_NO_WITH_VETO"] = "VOTE_OPTION_NO_WITH_VETO";
})(V1VoteOption = exports.V1VoteOption || (exports.V1VoteOption = {}));
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
 * @title cosmos/group/v1/events.proto
 * @version version not set
 */
class Api extends HttpClient {
    constructor() {
        super(...arguments);
        /**
         * No description
         *
         * @tags Query
         * @name QueryGroupInfo
         * @summary GroupInfo queries group info based on group id.
         * @request GET:/cosmos/group/v1/group_info/{group_id}
         */
        this.queryGroupInfo = (groupId, params = {}) => this.request({
            path: `/cosmos/group/v1/group_info/${groupId}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryGroupMembers
         * @summary GroupMembers queries members of a group
         * @request GET:/cosmos/group/v1/group_members/{group_id}
         */
        this.queryGroupMembers = (groupId, query, params = {}) => this.request({
            path: `/cosmos/group/v1/group_members/${groupId}`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryGroupPoliciesByAdmin
         * @summary GroupsByAdmin queries group policies by admin address.
         * @request GET:/cosmos/group/v1/group_policies_by_admin/{admin}
         */
        this.queryGroupPoliciesByAdmin = (admin, query, params = {}) => this.request({
            path: `/cosmos/group/v1/group_policies_by_admin/${admin}`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryGroupPoliciesByGroup
         * @summary GroupPoliciesByGroup queries group policies by group id.
         * @request GET:/cosmos/group/v1/group_policies_by_group/{group_id}
         */
        this.queryGroupPoliciesByGroup = (groupId, query, params = {}) => this.request({
            path: `/cosmos/group/v1/group_policies_by_group/${groupId}`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryGroupPolicyInfo
         * @summary GroupPolicyInfo queries group policy info based on account address of group policy.
         * @request GET:/cosmos/group/v1/group_policy_info/{address}
         */
        this.queryGroupPolicyInfo = (address, params = {}) => this.request({
            path: `/cosmos/group/v1/group_policy_info/${address}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryGroupsByAdmin
         * @summary GroupsByAdmin queries groups by admin address.
         * @request GET:/cosmos/group/v1/groups_by_admin/{admin}
         */
        this.queryGroupsByAdmin = (admin, query, params = {}) => this.request({
            path: `/cosmos/group/v1/groups_by_admin/${admin}`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryGroupsByMember
         * @summary GroupsByMember queries groups by member address.
         * @request GET:/cosmos/group/v1/groups_by_member/{address}
         */
        this.queryGroupsByMember = (address, query, params = {}) => this.request({
            path: `/cosmos/group/v1/groups_by_member/${address}`,
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
         * @summary Proposal queries a proposal based on proposal id.
         * @request GET:/cosmos/group/v1/proposal/{proposal_id}
         */
        this.queryProposal = (proposalId, params = {}) => this.request({
            path: `/cosmos/group/v1/proposal/${proposalId}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
       * No description
       *
       * @tags Query
       * @name QueryTallyResult
       * @summary TallyResult returns the tally result of a proposal. If the proposal is
      still in voting period, then this query computes the current tally state,
      which might not be final. On the other hand, if the proposal is final,
      then it simply returns the `final_tally_result` state stored in the
      proposal itself.
       * @request GET:/cosmos/group/v1/proposals/{proposal_id}/tally
       */
        this.queryTallyResult = (proposalId, params = {}) => this.request({
            path: `/cosmos/group/v1/proposals/${proposalId}/tally`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryProposalsByGroupPolicy
         * @summary ProposalsByGroupPolicy queries proposals based on account address of group policy.
         * @request GET:/cosmos/group/v1/proposals_by_group_policy/{address}
         */
        this.queryProposalsByGroupPolicy = (address, query, params = {}) => this.request({
            path: `/cosmos/group/v1/proposals_by_group_policy/${address}`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryVoteByProposalVoter
         * @summary VoteByProposalVoter queries a vote by proposal id and voter.
         * @request GET:/cosmos/group/v1/vote_by_proposal_voter/{proposal_id}/{voter}
         */
        this.queryVoteByProposalVoter = (proposalId, voter, params = {}) => this.request({
            path: `/cosmos/group/v1/vote_by_proposal_voter/${proposalId}/${voter}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryVotesByProposal
         * @summary VotesByProposal queries a vote by proposal.
         * @request GET:/cosmos/group/v1/votes_by_proposal/{proposal_id}
         */
        this.queryVotesByProposal = (proposalId, query, params = {}) => this.request({
            path: `/cosmos/group/v1/votes_by_proposal/${proposalId}`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryVotesByVoter
         * @summary VotesByVoter queries a vote by voter.
         * @request GET:/cosmos/group/v1/votes_by_voter/{voter}
         */
        this.queryVotesByVoter = (voter, query, params = {}) => this.request({
            path: `/cosmos/group/v1/votes_by_voter/${voter}`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
    }
}
exports.Api = Api;
