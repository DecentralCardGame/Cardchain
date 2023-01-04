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
exports.Api = exports.HttpClient = exports.ContentType = exports.V1State = exports.V1ResponseResultType = exports.V1Order = void 0;
/**
* - ORDER_NONE_UNSPECIFIED: zero-value for channel ordering
 - ORDER_UNORDERED: packets can be delivered in any order, which may differ from the order in
which they were sent.
 - ORDER_ORDERED: packets are delivered exactly in the order which they were sent
*/
var V1Order;
(function (V1Order) {
    V1Order["ORDER_NONE_UNSPECIFIED"] = "ORDER_NONE_UNSPECIFIED";
    V1Order["ORDER_UNORDERED"] = "ORDER_UNORDERED";
    V1Order["ORDER_ORDERED"] = "ORDER_ORDERED";
})(V1Order = exports.V1Order || (exports.V1Order = {}));
/**
* - RESPONSE_RESULT_TYPE_UNSPECIFIED: Default zero value enumeration
 - RESPONSE_RESULT_TYPE_NOOP: The message did not call the IBC application callbacks (because, for example, the packet had already been relayed)
 - RESPONSE_RESULT_TYPE_SUCCESS: The message was executed successfully
*/
var V1ResponseResultType;
(function (V1ResponseResultType) {
    V1ResponseResultType["RESPONSE_RESULT_TYPE_UNSPECIFIED"] = "RESPONSE_RESULT_TYPE_UNSPECIFIED";
    V1ResponseResultType["RESPONSE_RESULT_TYPE_NOOP"] = "RESPONSE_RESULT_TYPE_NOOP";
    V1ResponseResultType["RESPONSE_RESULT_TYPE_SUCCESS"] = "RESPONSE_RESULT_TYPE_SUCCESS";
})(V1ResponseResultType = exports.V1ResponseResultType || (exports.V1ResponseResultType = {}));
/**
* State defines if a channel is in one of the following states:
CLOSED, INIT, TRYOPEN, OPEN or UNINITIALIZED.

 - STATE_UNINITIALIZED_UNSPECIFIED: Default State
 - STATE_INIT: A channel has just started the opening handshake.
 - STATE_TRYOPEN: A channel has acknowledged the handshake step on the counterparty chain.
 - STATE_OPEN: A channel has completed the handshake. Open channels are
ready to send and receive packets.
 - STATE_CLOSED: A channel has been closed and can no longer be used to send or receive
packets.
*/
var V1State;
(function (V1State) {
    V1State["STATE_UNINITIALIZED_UNSPECIFIED"] = "STATE_UNINITIALIZED_UNSPECIFIED";
    V1State["STATE_INIT"] = "STATE_INIT";
    V1State["STATE_TRYOPEN"] = "STATE_TRYOPEN";
    V1State["STATE_OPEN"] = "STATE_OPEN";
    V1State["STATE_CLOSED"] = "STATE_CLOSED";
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
 * @title ibc/core/channel/v1/channel.proto
 * @version version not set
 */
class Api extends HttpClient {
    constructor() {
        super(...arguments);
        /**
         * No description
         *
         * @tags Query
         * @name QueryChannels
         * @summary Channels queries all the IBC channels of a chain.
         * @request GET:/ibc/core/channel/v1/channels
         */
        this.queryChannels = (query, params = {}) => this.request({
            path: `/ibc/core/channel/v1/channels`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryChannel
         * @summary Channel queries an IBC Channel.
         * @request GET:/ibc/core/channel/v1/channels/{channel_id}/ports/{port_id}
         */
        this.queryChannel = (channelId, portId, params = {}) => this.request({
            path: `/ibc/core/channel/v1/channels/${channelId}/ports/${portId}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
       * No description
       *
       * @tags Query
       * @name QueryChannelClientState
       * @summary ChannelClientState queries for the client state for the channel associated
      with the provided channel identifiers.
       * @request GET:/ibc/core/channel/v1/channels/{channel_id}/ports/{port_id}/client_state
       */
        this.queryChannelClientState = (channelId, portId, params = {}) => this.request({
            path: `/ibc/core/channel/v1/channels/${channelId}/ports/${portId}/client_state`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
       * No description
       *
       * @tags Query
       * @name QueryChannelConsensusState
       * @summary ChannelConsensusState queries for the consensus state for the channel
      associated with the provided channel identifiers.
       * @request GET:/ibc/core/channel/v1/channels/{channel_id}/ports/{port_id}/consensus_state/revision/{revision_number}/height/{revision_height}
       */
        this.queryChannelConsensusState = (channelId, portId, revisionNumber, revisionHeight, params = {}) => this.request({
            path: `/ibc/core/channel/v1/channels/${channelId}/ports/${portId}/consensus_state/revision/${revisionNumber}/height/${revisionHeight}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryNextSequenceReceive
         * @summary NextSequenceReceive returns the next receive sequence for a given channel.
         * @request GET:/ibc/core/channel/v1/channels/{channel_id}/ports/{port_id}/next_sequence
         */
        this.queryNextSequenceReceive = (channelId, portId, params = {}) => this.request({
            path: `/ibc/core/channel/v1/channels/${channelId}/ports/${portId}/next_sequence`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
       * No description
       *
       * @tags Query
       * @name QueryPacketAcknowledgements
       * @summary PacketAcknowledgements returns all the packet acknowledgements associated
      with a channel.
       * @request GET:/ibc/core/channel/v1/channels/{channel_id}/ports/{port_id}/packet_acknowledgements
       */
        this.queryPacketAcknowledgements = (channelId, portId, query, params = {}) => this.request({
            path: `/ibc/core/channel/v1/channels/${channelId}/ports/${portId}/packet_acknowledgements`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryPacketAcknowledgement
         * @summary PacketAcknowledgement queries a stored packet acknowledgement hash.
         * @request GET:/ibc/core/channel/v1/channels/{channel_id}/ports/{port_id}/packet_acks/{sequence}
         */
        this.queryPacketAcknowledgement = (channelId, portId, sequence, params = {}) => this.request({
            path: `/ibc/core/channel/v1/channels/${channelId}/ports/${portId}/packet_acks/${sequence}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
       * No description
       *
       * @tags Query
       * @name QueryPacketCommitments
       * @summary PacketCommitments returns all the packet commitments hashes associated
      with a channel.
       * @request GET:/ibc/core/channel/v1/channels/{channel_id}/ports/{port_id}/packet_commitments
       */
        this.queryPacketCommitments = (channelId, portId, query, params = {}) => this.request({
            path: `/ibc/core/channel/v1/channels/${channelId}/ports/${portId}/packet_commitments`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
       * No description
       *
       * @tags Query
       * @name QueryUnreceivedAcks
       * @summary UnreceivedAcks returns all the unreceived IBC acknowledgements associated
      with a channel and sequences.
       * @request GET:/ibc/core/channel/v1/channels/{channel_id}/ports/{port_id}/packet_commitments/{packet_ack_sequences}/unreceived_acks
       */
        this.queryUnreceivedAcks = (channelId, portId, packetAckSequences, params = {}) => this.request({
            path: `/ibc/core/channel/v1/channels/${channelId}/ports/${portId}/packet_commitments/${packetAckSequences}/unreceived_acks`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
       * No description
       *
       * @tags Query
       * @name QueryUnreceivedPackets
       * @summary UnreceivedPackets returns all the unreceived IBC packets associated with a
      channel and sequences.
       * @request GET:/ibc/core/channel/v1/channels/{channel_id}/ports/{port_id}/packet_commitments/{packet_commitment_sequences}/unreceived_packets
       */
        this.queryUnreceivedPackets = (channelId, portId, packetCommitmentSequences, params = {}) => this.request({
            path: `/ibc/core/channel/v1/channels/${channelId}/ports/${portId}/packet_commitments/${packetCommitmentSequences}/unreceived_packets`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryPacketCommitment
         * @summary PacketCommitment queries a stored packet commitment hash.
         * @request GET:/ibc/core/channel/v1/channels/{channel_id}/ports/{port_id}/packet_commitments/{sequence}
         */
        this.queryPacketCommitment = (channelId, portId, sequence, params = {}) => this.request({
            path: `/ibc/core/channel/v1/channels/${channelId}/ports/${portId}/packet_commitments/${sequence}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
       * No description
       *
       * @tags Query
       * @name QueryPacketReceipt
       * @summary PacketReceipt queries if a given packet sequence has been received on the
      queried chain
       * @request GET:/ibc/core/channel/v1/channels/{channel_id}/ports/{port_id}/packet_receipts/{sequence}
       */
        this.queryPacketReceipt = (channelId, portId, sequence, params = {}) => this.request({
            path: `/ibc/core/channel/v1/channels/${channelId}/ports/${portId}/packet_receipts/${sequence}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
       * No description
       *
       * @tags Query
       * @name QueryConnectionChannels
       * @summary ConnectionChannels queries all the channels associated with a connection
      end.
       * @request GET:/ibc/core/channel/v1/connections/{connection}/channels
       */
        this.queryConnectionChannels = (connection, query, params = {}) => this.request({
            path: `/ibc/core/channel/v1/connections/${connection}/channels`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
    }
}
exports.Api = Api;
