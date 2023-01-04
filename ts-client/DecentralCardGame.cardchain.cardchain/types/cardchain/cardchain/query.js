"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.QueryClientImpl = exports.QueryQCouncilsResponse = exports.IgnoreCouncils = exports.QueryQCouncilsRequest = exports.QueryRarityDistributionResponse = exports.QueryRarityDistributionRequest = exports.QueryQCollectionsResponse = exports.QueryQCollectionsRequest = exports.QueryQServerResponse = exports.QueryQServerRequest = exports.QueryQSellOffersResponse = exports.IgnoreSellOffers = exports.QueryQSellOffersRequest = exports.QueryQMatchesResponse = exports.IgnoreMatches = exports.QueryQMatchesRequest = exports.QueryQCouncilRequest = exports.QueryQSellOfferRequest = exports.QueryQCollectionRequest = exports.QueryQMatchRequest = exports.QueryQCardsResponse = exports.QueryQCardsRequest = exports.QueryQVotableCardsResponse = exports.QueryQVotableCardsRequest = exports.QueryQVotingResultsResponse = exports.QueryQVotingResultsRequest = exports.QueryQCardchainInfoResponse = exports.QueryQCardchainInfoRequest = exports.QueryQUserRequest = exports.QueryQCardContentResponse = exports.QueryQCardContentRequest = exports.QueryQCardRequest = exports.QueryParamsResponse = exports.QueryParamsRequest = exports.queryQCardsRequest_StatusToJSON = exports.queryQCardsRequest_StatusFromJSON = exports.QueryQCardsRequest_Status = exports.protobufPackage = void 0;
/* eslint-disable */
const long_1 = require("long");
const minimal_1 = require("protobufjs/minimal");
const card_1 = require("./card");
const collection_1 = require("./collection");
const council_1 = require("./council");
const match_1 = require("./match");
const params_1 = require("./params");
const sell_offer_1 = require("./sell_offer");
const server_1 = require("./server");
const tx_1 = require("./tx");
const user_1 = require("./user");
const vote_right_1 = require("./vote_right");
const voting_results_1 = require("./voting_results");
exports.protobufPackage = "DecentralCardGame.cardchain.cardchain";
var QueryQCardsRequest_Status;
(function (QueryQCardsRequest_Status) {
    QueryQCardsRequest_Status[QueryQCardsRequest_Status["scheme"] = 0] = "scheme";
    QueryQCardsRequest_Status[QueryQCardsRequest_Status["prototype"] = 1] = "prototype";
    QueryQCardsRequest_Status[QueryQCardsRequest_Status["trial"] = 2] = "trial";
    QueryQCardsRequest_Status[QueryQCardsRequest_Status["permanent"] = 3] = "permanent";
    QueryQCardsRequest_Status[QueryQCardsRequest_Status["suspended"] = 4] = "suspended";
    QueryQCardsRequest_Status[QueryQCardsRequest_Status["banned"] = 5] = "banned";
    QueryQCardsRequest_Status[QueryQCardsRequest_Status["bannedSoon"] = 6] = "bannedSoon";
    QueryQCardsRequest_Status[QueryQCardsRequest_Status["bannedVerySoon"] = 7] = "bannedVerySoon";
    QueryQCardsRequest_Status[QueryQCardsRequest_Status["none"] = 8] = "none";
    QueryQCardsRequest_Status[QueryQCardsRequest_Status["playable"] = 9] = "playable";
    QueryQCardsRequest_Status[QueryQCardsRequest_Status["unplayable"] = 10] = "unplayable";
    QueryQCardsRequest_Status[QueryQCardsRequest_Status["UNRECOGNIZED"] = -1] = "UNRECOGNIZED";
})(QueryQCardsRequest_Status = exports.QueryQCardsRequest_Status || (exports.QueryQCardsRequest_Status = {}));
function queryQCardsRequest_StatusFromJSON(object) {
    switch (object) {
        case 0:
        case "scheme":
            return QueryQCardsRequest_Status.scheme;
        case 1:
        case "prototype":
            return QueryQCardsRequest_Status.prototype;
        case 2:
        case "trial":
            return QueryQCardsRequest_Status.trial;
        case 3:
        case "permanent":
            return QueryQCardsRequest_Status.permanent;
        case 4:
        case "suspended":
            return QueryQCardsRequest_Status.suspended;
        case 5:
        case "banned":
            return QueryQCardsRequest_Status.banned;
        case 6:
        case "bannedSoon":
            return QueryQCardsRequest_Status.bannedSoon;
        case 7:
        case "bannedVerySoon":
            return QueryQCardsRequest_Status.bannedVerySoon;
        case 8:
        case "none":
            return QueryQCardsRequest_Status.none;
        case 9:
        case "playable":
            return QueryQCardsRequest_Status.playable;
        case 10:
        case "unplayable":
            return QueryQCardsRequest_Status.unplayable;
        case -1:
        case "UNRECOGNIZED":
        default:
            return QueryQCardsRequest_Status.UNRECOGNIZED;
    }
}
exports.queryQCardsRequest_StatusFromJSON = queryQCardsRequest_StatusFromJSON;
function queryQCardsRequest_StatusToJSON(object) {
    switch (object) {
        case QueryQCardsRequest_Status.scheme:
            return "scheme";
        case QueryQCardsRequest_Status.prototype:
            return "prototype";
        case QueryQCardsRequest_Status.trial:
            return "trial";
        case QueryQCardsRequest_Status.permanent:
            return "permanent";
        case QueryQCardsRequest_Status.suspended:
            return "suspended";
        case QueryQCardsRequest_Status.banned:
            return "banned";
        case QueryQCardsRequest_Status.bannedSoon:
            return "bannedSoon";
        case QueryQCardsRequest_Status.bannedVerySoon:
            return "bannedVerySoon";
        case QueryQCardsRequest_Status.none:
            return "none";
        case QueryQCardsRequest_Status.playable:
            return "playable";
        case QueryQCardsRequest_Status.unplayable:
            return "unplayable";
        case QueryQCardsRequest_Status.UNRECOGNIZED:
        default:
            return "UNRECOGNIZED";
    }
}
exports.queryQCardsRequest_StatusToJSON = queryQCardsRequest_StatusToJSON;
function createBaseQueryParamsRequest() {
    return {};
}
exports.QueryParamsRequest = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseQueryParamsRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        return {};
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = createBaseQueryParamsRequest();
        return message;
    },
};
function createBaseQueryParamsResponse() {
    return { params: undefined };
}
exports.QueryParamsResponse = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.params !== undefined) {
            params_1.Params.encode(message.params, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseQueryParamsResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.params = params_1.Params.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return { params: isSet(object.params) ? params_1.Params.fromJSON(object.params) : undefined };
    },
    toJSON(message) {
        const obj = {};
        message.params !== undefined && (obj.params = message.params ? params_1.Params.toJSON(message.params) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseQueryParamsResponse();
        message.params = (object.params !== undefined && object.params !== null)
            ? params_1.Params.fromPartial(object.params)
            : undefined;
        return message;
    },
};
function createBaseQueryQCardRequest() {
    return { cardId: "" };
}
exports.QueryQCardRequest = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.cardId !== "") {
            writer.uint32(10).string(message.cardId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseQueryQCardRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.cardId = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return { cardId: isSet(object.cardId) ? String(object.cardId) : "" };
    },
    toJSON(message) {
        const obj = {};
        message.cardId !== undefined && (obj.cardId = message.cardId);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseQueryQCardRequest();
        message.cardId = object.cardId ?? "";
        return message;
    },
};
function createBaseQueryQCardContentRequest() {
    return { cardId: "" };
}
exports.QueryQCardContentRequest = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.cardId !== "") {
            writer.uint32(10).string(message.cardId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseQueryQCardContentRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.cardId = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return { cardId: isSet(object.cardId) ? String(object.cardId) : "" };
    },
    toJSON(message) {
        const obj = {};
        message.cardId !== undefined && (obj.cardId = message.cardId);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseQueryQCardContentRequest();
        message.cardId = object.cardId ?? "";
        return message;
    },
};
function createBaseQueryQCardContentResponse() {
    return { content: "", hash: "" };
}
exports.QueryQCardContentResponse = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.content !== "") {
            writer.uint32(10).string(message.content);
        }
        if (message.hash !== "") {
            writer.uint32(18).string(message.hash);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseQueryQCardContentResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.content = reader.string();
                    break;
                case 2:
                    message.hash = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return {
            content: isSet(object.content) ? String(object.content) : "",
            hash: isSet(object.hash) ? String(object.hash) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        message.content !== undefined && (obj.content = message.content);
        message.hash !== undefined && (obj.hash = message.hash);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseQueryQCardContentResponse();
        message.content = object.content ?? "";
        message.hash = object.hash ?? "";
        return message;
    },
};
function createBaseQueryQUserRequest() {
    return { address: "" };
}
exports.QueryQUserRequest = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.address !== "") {
            writer.uint32(10).string(message.address);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseQueryQUserRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.address = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return { address: isSet(object.address) ? String(object.address) : "" };
    },
    toJSON(message) {
        const obj = {};
        message.address !== undefined && (obj.address = message.address);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseQueryQUserRequest();
        message.address = object.address ?? "";
        return message;
    },
};
function createBaseQueryQCardchainInfoRequest() {
    return {};
}
exports.QueryQCardchainInfoRequest = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseQueryQCardchainInfoRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        return {};
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = createBaseQueryQCardchainInfoRequest();
        return message;
    },
};
function createBaseQueryQCardchainInfoResponse() {
    return {
        cardAuctionPrice: "",
        activeCollections: [],
        cardsNumber: 0,
        matchesNumber: 0,
        sellOffersNumber: 0,
        councilsNumber: 0,
    };
}
exports.QueryQCardchainInfoResponse = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.cardAuctionPrice !== "") {
            writer.uint32(10).string(message.cardAuctionPrice);
        }
        writer.uint32(18).fork();
        for (const v of message.activeCollections) {
            writer.uint64(v);
        }
        writer.ldelim();
        if (message.cardsNumber !== 0) {
            writer.uint32(24).uint64(message.cardsNumber);
        }
        if (message.matchesNumber !== 0) {
            writer.uint32(32).uint64(message.matchesNumber);
        }
        if (message.sellOffersNumber !== 0) {
            writer.uint32(40).uint64(message.sellOffersNumber);
        }
        if (message.councilsNumber !== 0) {
            writer.uint32(48).uint64(message.councilsNumber);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseQueryQCardchainInfoResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.cardAuctionPrice = reader.string();
                    break;
                case 2:
                    if ((tag & 7) === 2) {
                        const end2 = reader.uint32() + reader.pos;
                        while (reader.pos < end2) {
                            message.activeCollections.push(longToNumber(reader.uint64()));
                        }
                    }
                    else {
                        message.activeCollections.push(longToNumber(reader.uint64()));
                    }
                    break;
                case 3:
                    message.cardsNumber = longToNumber(reader.uint64());
                    break;
                case 4:
                    message.matchesNumber = longToNumber(reader.uint64());
                    break;
                case 5:
                    message.sellOffersNumber = longToNumber(reader.uint64());
                    break;
                case 6:
                    message.councilsNumber = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return {
            cardAuctionPrice: isSet(object.cardAuctionPrice) ? String(object.cardAuctionPrice) : "",
            activeCollections: Array.isArray(object?.activeCollections)
                ? object.activeCollections.map((e) => Number(e))
                : [],
            cardsNumber: isSet(object.cardsNumber) ? Number(object.cardsNumber) : 0,
            matchesNumber: isSet(object.matchesNumber) ? Number(object.matchesNumber) : 0,
            sellOffersNumber: isSet(object.sellOffersNumber) ? Number(object.sellOffersNumber) : 0,
            councilsNumber: isSet(object.councilsNumber) ? Number(object.councilsNumber) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        message.cardAuctionPrice !== undefined && (obj.cardAuctionPrice = message.cardAuctionPrice);
        if (message.activeCollections) {
            obj.activeCollections = message.activeCollections.map((e) => Math.round(e));
        }
        else {
            obj.activeCollections = [];
        }
        message.cardsNumber !== undefined && (obj.cardsNumber = Math.round(message.cardsNumber));
        message.matchesNumber !== undefined && (obj.matchesNumber = Math.round(message.matchesNumber));
        message.sellOffersNumber !== undefined && (obj.sellOffersNumber = Math.round(message.sellOffersNumber));
        message.councilsNumber !== undefined && (obj.councilsNumber = Math.round(message.councilsNumber));
        return obj;
    },
    fromPartial(object) {
        const message = createBaseQueryQCardchainInfoResponse();
        message.cardAuctionPrice = object.cardAuctionPrice ?? "";
        message.activeCollections = object.activeCollections?.map((e) => e) || [];
        message.cardsNumber = object.cardsNumber ?? 0;
        message.matchesNumber = object.matchesNumber ?? 0;
        message.sellOffersNumber = object.sellOffersNumber ?? 0;
        message.councilsNumber = object.councilsNumber ?? 0;
        return message;
    },
};
function createBaseQueryQVotingResultsRequest() {
    return {};
}
exports.QueryQVotingResultsRequest = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseQueryQVotingResultsRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        return {};
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = createBaseQueryQVotingResultsRequest();
        return message;
    },
};
function createBaseQueryQVotingResultsResponse() {
    return { lastVotingResults: undefined };
}
exports.QueryQVotingResultsResponse = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.lastVotingResults !== undefined) {
            voting_results_1.VotingResults.encode(message.lastVotingResults, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseQueryQVotingResultsResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.lastVotingResults = voting_results_1.VotingResults.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return {
            lastVotingResults: isSet(object.lastVotingResults) ? voting_results_1.VotingResults.fromJSON(object.lastVotingResults) : undefined,
        };
    },
    toJSON(message) {
        const obj = {};
        message.lastVotingResults !== undefined && (obj.lastVotingResults = message.lastVotingResults
            ? voting_results_1.VotingResults.toJSON(message.lastVotingResults)
            : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseQueryQVotingResultsResponse();
        message.lastVotingResults = (object.lastVotingResults !== undefined && object.lastVotingResults !== null)
            ? voting_results_1.VotingResults.fromPartial(object.lastVotingResults)
            : undefined;
        return message;
    },
};
function createBaseQueryQVotableCardsRequest() {
    return { address: "" };
}
exports.QueryQVotableCardsRequest = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.address !== "") {
            writer.uint32(10).string(message.address);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseQueryQVotableCardsRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.address = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return { address: isSet(object.address) ? String(object.address) : "" };
    },
    toJSON(message) {
        const obj = {};
        message.address !== undefined && (obj.address = message.address);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseQueryQVotableCardsRequest();
        message.address = object.address ?? "";
        return message;
    },
};
function createBaseQueryQVotableCardsResponse() {
    return { unregistered: false, noVoteRights: false, voteRights: [] };
}
exports.QueryQVotableCardsResponse = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.unregistered === true) {
            writer.uint32(8).bool(message.unregistered);
        }
        if (message.noVoteRights === true) {
            writer.uint32(16).bool(message.noVoteRights);
        }
        for (const v of message.voteRights) {
            vote_right_1.VoteRight.encode(v, writer.uint32(26).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseQueryQVotableCardsResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.unregistered = reader.bool();
                    break;
                case 2:
                    message.noVoteRights = reader.bool();
                    break;
                case 3:
                    message.voteRights.push(vote_right_1.VoteRight.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return {
            unregistered: isSet(object.unregistered) ? Boolean(object.unregistered) : false,
            noVoteRights: isSet(object.noVoteRights) ? Boolean(object.noVoteRights) : false,
            voteRights: Array.isArray(object?.voteRights) ? object.voteRights.map((e) => vote_right_1.VoteRight.fromJSON(e)) : [],
        };
    },
    toJSON(message) {
        const obj = {};
        message.unregistered !== undefined && (obj.unregistered = message.unregistered);
        message.noVoteRights !== undefined && (obj.noVoteRights = message.noVoteRights);
        if (message.voteRights) {
            obj.voteRights = message.voteRights.map((e) => e ? vote_right_1.VoteRight.toJSON(e) : undefined);
        }
        else {
            obj.voteRights = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = createBaseQueryQVotableCardsResponse();
        message.unregistered = object.unregistered ?? false;
        message.noVoteRights = object.noVoteRights ?? false;
        message.voteRights = object.voteRights?.map((e) => vote_right_1.VoteRight.fromPartial(e)) || [];
        return message;
    },
};
function createBaseQueryQCardsRequest() {
    return {
        owner: "",
        status: 0,
        cardType: "",
        classes: "",
        sortBy: "",
        nameContains: "",
        keywordsContains: "",
        notesContains: "",
    };
}
exports.QueryQCardsRequest = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.owner !== "") {
            writer.uint32(10).string(message.owner);
        }
        if (message.status !== 0) {
            writer.uint32(16).int32(message.status);
        }
        if (message.cardType !== "") {
            writer.uint32(26).string(message.cardType);
        }
        if (message.classes !== "") {
            writer.uint32(34).string(message.classes);
        }
        if (message.sortBy !== "") {
            writer.uint32(42).string(message.sortBy);
        }
        if (message.nameContains !== "") {
            writer.uint32(50).string(message.nameContains);
        }
        if (message.keywordsContains !== "") {
            writer.uint32(58).string(message.keywordsContains);
        }
        if (message.notesContains !== "") {
            writer.uint32(66).string(message.notesContains);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseQueryQCardsRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.owner = reader.string();
                    break;
                case 2:
                    message.status = reader.int32();
                    break;
                case 3:
                    message.cardType = reader.string();
                    break;
                case 4:
                    message.classes = reader.string();
                    break;
                case 5:
                    message.sortBy = reader.string();
                    break;
                case 6:
                    message.nameContains = reader.string();
                    break;
                case 7:
                    message.keywordsContains = reader.string();
                    break;
                case 8:
                    message.notesContains = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return {
            owner: isSet(object.owner) ? String(object.owner) : "",
            status: isSet(object.status) ? queryQCardsRequest_StatusFromJSON(object.status) : 0,
            cardType: isSet(object.cardType) ? String(object.cardType) : "",
            classes: isSet(object.classes) ? String(object.classes) : "",
            sortBy: isSet(object.sortBy) ? String(object.sortBy) : "",
            nameContains: isSet(object.nameContains) ? String(object.nameContains) : "",
            keywordsContains: isSet(object.keywordsContains) ? String(object.keywordsContains) : "",
            notesContains: isSet(object.notesContains) ? String(object.notesContains) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        message.owner !== undefined && (obj.owner = message.owner);
        message.status !== undefined && (obj.status = queryQCardsRequest_StatusToJSON(message.status));
        message.cardType !== undefined && (obj.cardType = message.cardType);
        message.classes !== undefined && (obj.classes = message.classes);
        message.sortBy !== undefined && (obj.sortBy = message.sortBy);
        message.nameContains !== undefined && (obj.nameContains = message.nameContains);
        message.keywordsContains !== undefined && (obj.keywordsContains = message.keywordsContains);
        message.notesContains !== undefined && (obj.notesContains = message.notesContains);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseQueryQCardsRequest();
        message.owner = object.owner ?? "";
        message.status = object.status ?? 0;
        message.cardType = object.cardType ?? "";
        message.classes = object.classes ?? "";
        message.sortBy = object.sortBy ?? "";
        message.nameContains = object.nameContains ?? "";
        message.keywordsContains = object.keywordsContains ?? "";
        message.notesContains = object.notesContains ?? "";
        return message;
    },
};
function createBaseQueryQCardsResponse() {
    return { cardsList: [] };
}
exports.QueryQCardsResponse = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        writer.uint32(10).fork();
        for (const v of message.cardsList) {
            writer.uint64(v);
        }
        writer.ldelim();
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseQueryQCardsResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    if ((tag & 7) === 2) {
                        const end2 = reader.uint32() + reader.pos;
                        while (reader.pos < end2) {
                            message.cardsList.push(longToNumber(reader.uint64()));
                        }
                    }
                    else {
                        message.cardsList.push(longToNumber(reader.uint64()));
                    }
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return { cardsList: Array.isArray(object?.cardsList) ? object.cardsList.map((e) => Number(e)) : [] };
    },
    toJSON(message) {
        const obj = {};
        if (message.cardsList) {
            obj.cardsList = message.cardsList.map((e) => Math.round(e));
        }
        else {
            obj.cardsList = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = createBaseQueryQCardsResponse();
        message.cardsList = object.cardsList?.map((e) => e) || [];
        return message;
    },
};
function createBaseQueryQMatchRequest() {
    return { matchId: 0 };
}
exports.QueryQMatchRequest = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.matchId !== 0) {
            writer.uint32(8).uint64(message.matchId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseQueryQMatchRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.matchId = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return { matchId: isSet(object.matchId) ? Number(object.matchId) : 0 };
    },
    toJSON(message) {
        const obj = {};
        message.matchId !== undefined && (obj.matchId = Math.round(message.matchId));
        return obj;
    },
    fromPartial(object) {
        const message = createBaseQueryQMatchRequest();
        message.matchId = object.matchId ?? 0;
        return message;
    },
};
function createBaseQueryQCollectionRequest() {
    return { collectionId: 0 };
}
exports.QueryQCollectionRequest = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.collectionId !== 0) {
            writer.uint32(8).uint64(message.collectionId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseQueryQCollectionRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.collectionId = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return { collectionId: isSet(object.collectionId) ? Number(object.collectionId) : 0 };
    },
    toJSON(message) {
        const obj = {};
        message.collectionId !== undefined && (obj.collectionId = Math.round(message.collectionId));
        return obj;
    },
    fromPartial(object) {
        const message = createBaseQueryQCollectionRequest();
        message.collectionId = object.collectionId ?? 0;
        return message;
    },
};
function createBaseQueryQSellOfferRequest() {
    return { sellOfferId: 0 };
}
exports.QueryQSellOfferRequest = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.sellOfferId !== 0) {
            writer.uint32(8).uint64(message.sellOfferId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseQueryQSellOfferRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.sellOfferId = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return { sellOfferId: isSet(object.sellOfferId) ? Number(object.sellOfferId) : 0 };
    },
    toJSON(message) {
        const obj = {};
        message.sellOfferId !== undefined && (obj.sellOfferId = Math.round(message.sellOfferId));
        return obj;
    },
    fromPartial(object) {
        const message = createBaseQueryQSellOfferRequest();
        message.sellOfferId = object.sellOfferId ?? 0;
        return message;
    },
};
function createBaseQueryQCouncilRequest() {
    return { councilId: 0 };
}
exports.QueryQCouncilRequest = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.councilId !== 0) {
            writer.uint32(8).uint64(message.councilId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseQueryQCouncilRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.councilId = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return { councilId: isSet(object.councilId) ? Number(object.councilId) : 0 };
    },
    toJSON(message) {
        const obj = {};
        message.councilId !== undefined && (obj.councilId = Math.round(message.councilId));
        return obj;
    },
    fromPartial(object) {
        const message = createBaseQueryQCouncilRequest();
        message.councilId = object.councilId ?? 0;
        return message;
    },
};
function createBaseQueryQMatchesRequest() {
    return {
        timestampDown: 0,
        timestampUp: 0,
        containsUsers: [],
        reporter: "",
        outcome: 0,
        cardsPlayed: [],
        ignore: undefined,
    };
}
exports.QueryQMatchesRequest = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.timestampDown !== 0) {
            writer.uint32(8).uint64(message.timestampDown);
        }
        if (message.timestampUp !== 0) {
            writer.uint32(16).uint64(message.timestampUp);
        }
        for (const v of message.containsUsers) {
            writer.uint32(26).string(v);
        }
        if (message.reporter !== "") {
            writer.uint32(34).string(message.reporter);
        }
        if (message.outcome !== 0) {
            writer.uint32(40).int32(message.outcome);
        }
        writer.uint32(50).fork();
        for (const v of message.cardsPlayed) {
            writer.uint64(v);
        }
        writer.ldelim();
        if (message.ignore !== undefined) {
            exports.IgnoreMatches.encode(message.ignore, writer.uint32(58).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseQueryQMatchesRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.timestampDown = longToNumber(reader.uint64());
                    break;
                case 2:
                    message.timestampUp = longToNumber(reader.uint64());
                    break;
                case 3:
                    message.containsUsers.push(reader.string());
                    break;
                case 4:
                    message.reporter = reader.string();
                    break;
                case 5:
                    message.outcome = reader.int32();
                    break;
                case 6:
                    if ((tag & 7) === 2) {
                        const end2 = reader.uint32() + reader.pos;
                        while (reader.pos < end2) {
                            message.cardsPlayed.push(longToNumber(reader.uint64()));
                        }
                    }
                    else {
                        message.cardsPlayed.push(longToNumber(reader.uint64()));
                    }
                    break;
                case 7:
                    message.ignore = exports.IgnoreMatches.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return {
            timestampDown: isSet(object.timestampDown) ? Number(object.timestampDown) : 0,
            timestampUp: isSet(object.timestampUp) ? Number(object.timestampUp) : 0,
            containsUsers: Array.isArray(object?.containsUsers) ? object.containsUsers.map((e) => String(e)) : [],
            reporter: isSet(object.reporter) ? String(object.reporter) : "",
            outcome: isSet(object.outcome) ? (0, tx_1.outcomeFromJSON)(object.outcome) : 0,
            cardsPlayed: Array.isArray(object?.cardsPlayed) ? object.cardsPlayed.map((e) => Number(e)) : [],
            ignore: isSet(object.ignore) ? exports.IgnoreMatches.fromJSON(object.ignore) : undefined,
        };
    },
    toJSON(message) {
        const obj = {};
        message.timestampDown !== undefined && (obj.timestampDown = Math.round(message.timestampDown));
        message.timestampUp !== undefined && (obj.timestampUp = Math.round(message.timestampUp));
        if (message.containsUsers) {
            obj.containsUsers = message.containsUsers.map((e) => e);
        }
        else {
            obj.containsUsers = [];
        }
        message.reporter !== undefined && (obj.reporter = message.reporter);
        message.outcome !== undefined && (obj.outcome = (0, tx_1.outcomeToJSON)(message.outcome));
        if (message.cardsPlayed) {
            obj.cardsPlayed = message.cardsPlayed.map((e) => Math.round(e));
        }
        else {
            obj.cardsPlayed = [];
        }
        message.ignore !== undefined && (obj.ignore = message.ignore ? exports.IgnoreMatches.toJSON(message.ignore) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseQueryQMatchesRequest();
        message.timestampDown = object.timestampDown ?? 0;
        message.timestampUp = object.timestampUp ?? 0;
        message.containsUsers = object.containsUsers?.map((e) => e) || [];
        message.reporter = object.reporter ?? "";
        message.outcome = object.outcome ?? 0;
        message.cardsPlayed = object.cardsPlayed?.map((e) => e) || [];
        message.ignore = (object.ignore !== undefined && object.ignore !== null)
            ? exports.IgnoreMatches.fromPartial(object.ignore)
            : undefined;
        return message;
    },
};
function createBaseIgnoreMatches() {
    return { outcome: false };
}
exports.IgnoreMatches = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.outcome === true) {
            writer.uint32(8).bool(message.outcome);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseIgnoreMatches();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.outcome = reader.bool();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return { outcome: isSet(object.outcome) ? Boolean(object.outcome) : false };
    },
    toJSON(message) {
        const obj = {};
        message.outcome !== undefined && (obj.outcome = message.outcome);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseIgnoreMatches();
        message.outcome = object.outcome ?? false;
        return message;
    },
};
function createBaseQueryQMatchesResponse() {
    return { matchesList: [], matches: [] };
}
exports.QueryQMatchesResponse = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        writer.uint32(10).fork();
        for (const v of message.matchesList) {
            writer.uint64(v);
        }
        writer.ldelim();
        for (const v of message.matches) {
            match_1.Match.encode(v, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseQueryQMatchesResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    if ((tag & 7) === 2) {
                        const end2 = reader.uint32() + reader.pos;
                        while (reader.pos < end2) {
                            message.matchesList.push(longToNumber(reader.uint64()));
                        }
                    }
                    else {
                        message.matchesList.push(longToNumber(reader.uint64()));
                    }
                    break;
                case 2:
                    message.matches.push(match_1.Match.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return {
            matchesList: Array.isArray(object?.matchesList) ? object.matchesList.map((e) => Number(e)) : [],
            matches: Array.isArray(object?.matches) ? object.matches.map((e) => match_1.Match.fromJSON(e)) : [],
        };
    },
    toJSON(message) {
        const obj = {};
        if (message.matchesList) {
            obj.matchesList = message.matchesList.map((e) => Math.round(e));
        }
        else {
            obj.matchesList = [];
        }
        if (message.matches) {
            obj.matches = message.matches.map((e) => e ? match_1.Match.toJSON(e) : undefined);
        }
        else {
            obj.matches = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = createBaseQueryQMatchesResponse();
        message.matchesList = object.matchesList?.map((e) => e) || [];
        message.matches = object.matches?.map((e) => match_1.Match.fromPartial(e)) || [];
        return message;
    },
};
function createBaseQueryQSellOffersRequest() {
    return { priceDown: "", priceUp: "", seller: "", buyer: "", card: 0, status: 0, ignore: undefined };
}
exports.QueryQSellOffersRequest = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.priceDown !== "") {
            writer.uint32(10).string(message.priceDown);
        }
        if (message.priceUp !== "") {
            writer.uint32(18).string(message.priceUp);
        }
        if (message.seller !== "") {
            writer.uint32(26).string(message.seller);
        }
        if (message.buyer !== "") {
            writer.uint32(34).string(message.buyer);
        }
        if (message.card !== 0) {
            writer.uint32(40).uint64(message.card);
        }
        if (message.status !== 0) {
            writer.uint32(48).int32(message.status);
        }
        if (message.ignore !== undefined) {
            exports.IgnoreSellOffers.encode(message.ignore, writer.uint32(58).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseQueryQSellOffersRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.priceDown = reader.string();
                    break;
                case 2:
                    message.priceUp = reader.string();
                    break;
                case 3:
                    message.seller = reader.string();
                    break;
                case 4:
                    message.buyer = reader.string();
                    break;
                case 5:
                    message.card = longToNumber(reader.uint64());
                    break;
                case 6:
                    message.status = reader.int32();
                    break;
                case 7:
                    message.ignore = exports.IgnoreSellOffers.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return {
            priceDown: isSet(object.priceDown) ? String(object.priceDown) : "",
            priceUp: isSet(object.priceUp) ? String(object.priceUp) : "",
            seller: isSet(object.seller) ? String(object.seller) : "",
            buyer: isSet(object.buyer) ? String(object.buyer) : "",
            card: isSet(object.card) ? Number(object.card) : 0,
            status: isSet(object.status) ? (0, sell_offer_1.sellOfferStatusFromJSON)(object.status) : 0,
            ignore: isSet(object.ignore) ? exports.IgnoreSellOffers.fromJSON(object.ignore) : undefined,
        };
    },
    toJSON(message) {
        const obj = {};
        message.priceDown !== undefined && (obj.priceDown = message.priceDown);
        message.priceUp !== undefined && (obj.priceUp = message.priceUp);
        message.seller !== undefined && (obj.seller = message.seller);
        message.buyer !== undefined && (obj.buyer = message.buyer);
        message.card !== undefined && (obj.card = Math.round(message.card));
        message.status !== undefined && (obj.status = (0, sell_offer_1.sellOfferStatusToJSON)(message.status));
        message.ignore !== undefined && (obj.ignore = message.ignore ? exports.IgnoreSellOffers.toJSON(message.ignore) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseQueryQSellOffersRequest();
        message.priceDown = object.priceDown ?? "";
        message.priceUp = object.priceUp ?? "";
        message.seller = object.seller ?? "";
        message.buyer = object.buyer ?? "";
        message.card = object.card ?? 0;
        message.status = object.status ?? 0;
        message.ignore = (object.ignore !== undefined && object.ignore !== null)
            ? exports.IgnoreSellOffers.fromPartial(object.ignore)
            : undefined;
        return message;
    },
};
function createBaseIgnoreSellOffers() {
    return { status: false, card: false };
}
exports.IgnoreSellOffers = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.status === true) {
            writer.uint32(8).bool(message.status);
        }
        if (message.card === true) {
            writer.uint32(16).bool(message.card);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseIgnoreSellOffers();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.status = reader.bool();
                    break;
                case 2:
                    message.card = reader.bool();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return {
            status: isSet(object.status) ? Boolean(object.status) : false,
            card: isSet(object.card) ? Boolean(object.card) : false,
        };
    },
    toJSON(message) {
        const obj = {};
        message.status !== undefined && (obj.status = message.status);
        message.card !== undefined && (obj.card = message.card);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseIgnoreSellOffers();
        message.status = object.status ?? false;
        message.card = object.card ?? false;
        return message;
    },
};
function createBaseQueryQSellOffersResponse() {
    return { sellOffersIds: [], sellOffers: [] };
}
exports.QueryQSellOffersResponse = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        writer.uint32(10).fork();
        for (const v of message.sellOffersIds) {
            writer.uint64(v);
        }
        writer.ldelim();
        for (const v of message.sellOffers) {
            sell_offer_1.SellOffer.encode(v, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseQueryQSellOffersResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    if ((tag & 7) === 2) {
                        const end2 = reader.uint32() + reader.pos;
                        while (reader.pos < end2) {
                            message.sellOffersIds.push(longToNumber(reader.uint64()));
                        }
                    }
                    else {
                        message.sellOffersIds.push(longToNumber(reader.uint64()));
                    }
                    break;
                case 2:
                    message.sellOffers.push(sell_offer_1.SellOffer.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return {
            sellOffersIds: Array.isArray(object?.sellOffersIds) ? object.sellOffersIds.map((e) => Number(e)) : [],
            sellOffers: Array.isArray(object?.sellOffers) ? object.sellOffers.map((e) => sell_offer_1.SellOffer.fromJSON(e)) : [],
        };
    },
    toJSON(message) {
        const obj = {};
        if (message.sellOffersIds) {
            obj.sellOffersIds = message.sellOffersIds.map((e) => Math.round(e));
        }
        else {
            obj.sellOffersIds = [];
        }
        if (message.sellOffers) {
            obj.sellOffers = message.sellOffers.map((e) => e ? sell_offer_1.SellOffer.toJSON(e) : undefined);
        }
        else {
            obj.sellOffers = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = createBaseQueryQSellOffersResponse();
        message.sellOffersIds = object.sellOffersIds?.map((e) => e) || [];
        message.sellOffers = object.sellOffers?.map((e) => sell_offer_1.SellOffer.fromPartial(e)) || [];
        return message;
    },
};
function createBaseQueryQServerRequest() {
    return { id: 0 };
}
exports.QueryQServerRequest = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.id !== 0) {
            writer.uint32(8).uint64(message.id);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseQueryQServerRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.id = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return { id: isSet(object.id) ? Number(object.id) : 0 };
    },
    toJSON(message) {
        const obj = {};
        message.id !== undefined && (obj.id = Math.round(message.id));
        return obj;
    },
    fromPartial(object) {
        const message = createBaseQueryQServerRequest();
        message.id = object.id ?? 0;
        return message;
    },
};
function createBaseQueryQServerResponse() {
    return {};
}
exports.QueryQServerResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseQueryQServerResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        return {};
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = createBaseQueryQServerResponse();
        return message;
    },
};
function createBaseQueryQCollectionsRequest() {
    return { status: 0, ignoreStatus: false, contributors: [], containsCards: [], owner: "" };
}
exports.QueryQCollectionsRequest = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.status !== 0) {
            writer.uint32(8).int32(message.status);
        }
        if (message.ignoreStatus === true) {
            writer.uint32(16).bool(message.ignoreStatus);
        }
        for (const v of message.contributors) {
            writer.uint32(26).string(v);
        }
        writer.uint32(34).fork();
        for (const v of message.containsCards) {
            writer.uint64(v);
        }
        writer.ldelim();
        if (message.owner !== "") {
            writer.uint32(42).string(message.owner);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseQueryQCollectionsRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.status = reader.int32();
                    break;
                case 2:
                    message.ignoreStatus = reader.bool();
                    break;
                case 3:
                    message.contributors.push(reader.string());
                    break;
                case 4:
                    if ((tag & 7) === 2) {
                        const end2 = reader.uint32() + reader.pos;
                        while (reader.pos < end2) {
                            message.containsCards.push(longToNumber(reader.uint64()));
                        }
                    }
                    else {
                        message.containsCards.push(longToNumber(reader.uint64()));
                    }
                    break;
                case 5:
                    message.owner = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return {
            status: isSet(object.status) ? (0, collection_1.cStatusFromJSON)(object.status) : 0,
            ignoreStatus: isSet(object.ignoreStatus) ? Boolean(object.ignoreStatus) : false,
            contributors: Array.isArray(object?.contributors) ? object.contributors.map((e) => String(e)) : [],
            containsCards: Array.isArray(object?.containsCards) ? object.containsCards.map((e) => Number(e)) : [],
            owner: isSet(object.owner) ? String(object.owner) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        message.status !== undefined && (obj.status = (0, collection_1.cStatusToJSON)(message.status));
        message.ignoreStatus !== undefined && (obj.ignoreStatus = message.ignoreStatus);
        if (message.contributors) {
            obj.contributors = message.contributors.map((e) => e);
        }
        else {
            obj.contributors = [];
        }
        if (message.containsCards) {
            obj.containsCards = message.containsCards.map((e) => Math.round(e));
        }
        else {
            obj.containsCards = [];
        }
        message.owner !== undefined && (obj.owner = message.owner);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseQueryQCollectionsRequest();
        message.status = object.status ?? 0;
        message.ignoreStatus = object.ignoreStatus ?? false;
        message.contributors = object.contributors?.map((e) => e) || [];
        message.containsCards = object.containsCards?.map((e) => e) || [];
        message.owner = object.owner ?? "";
        return message;
    },
};
function createBaseQueryQCollectionsResponse() {
    return { collectionIds: [] };
}
exports.QueryQCollectionsResponse = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        writer.uint32(10).fork();
        for (const v of message.collectionIds) {
            writer.uint64(v);
        }
        writer.ldelim();
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseQueryQCollectionsResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    if ((tag & 7) === 2) {
                        const end2 = reader.uint32() + reader.pos;
                        while (reader.pos < end2) {
                            message.collectionIds.push(longToNumber(reader.uint64()));
                        }
                    }
                    else {
                        message.collectionIds.push(longToNumber(reader.uint64()));
                    }
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return {
            collectionIds: Array.isArray(object?.collectionIds) ? object.collectionIds.map((e) => Number(e)) : [],
        };
    },
    toJSON(message) {
        const obj = {};
        if (message.collectionIds) {
            obj.collectionIds = message.collectionIds.map((e) => Math.round(e));
        }
        else {
            obj.collectionIds = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = createBaseQueryQCollectionsResponse();
        message.collectionIds = object.collectionIds?.map((e) => e) || [];
        return message;
    },
};
function createBaseQueryRarityDistributionRequest() {
    return { collectionId: 0 };
}
exports.QueryRarityDistributionRequest = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.collectionId !== 0) {
            writer.uint32(8).uint64(message.collectionId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseQueryRarityDistributionRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.collectionId = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return { collectionId: isSet(object.collectionId) ? Number(object.collectionId) : 0 };
    },
    toJSON(message) {
        const obj = {};
        message.collectionId !== undefined && (obj.collectionId = Math.round(message.collectionId));
        return obj;
    },
    fromPartial(object) {
        const message = createBaseQueryRarityDistributionRequest();
        message.collectionId = object.collectionId ?? 0;
        return message;
    },
};
function createBaseQueryRarityDistributionResponse() {
    return { current: [], wanted: [] };
}
exports.QueryRarityDistributionResponse = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        writer.uint32(10).fork();
        for (const v of message.current) {
            writer.uint64(v);
        }
        writer.ldelim();
        writer.uint32(18).fork();
        for (const v of message.wanted) {
            writer.uint64(v);
        }
        writer.ldelim();
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseQueryRarityDistributionResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    if ((tag & 7) === 2) {
                        const end2 = reader.uint32() + reader.pos;
                        while (reader.pos < end2) {
                            message.current.push(longToNumber(reader.uint64()));
                        }
                    }
                    else {
                        message.current.push(longToNumber(reader.uint64()));
                    }
                    break;
                case 2:
                    if ((tag & 7) === 2) {
                        const end2 = reader.uint32() + reader.pos;
                        while (reader.pos < end2) {
                            message.wanted.push(longToNumber(reader.uint64()));
                        }
                    }
                    else {
                        message.wanted.push(longToNumber(reader.uint64()));
                    }
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return {
            current: Array.isArray(object?.current) ? object.current.map((e) => Number(e)) : [],
            wanted: Array.isArray(object?.wanted) ? object.wanted.map((e) => Number(e)) : [],
        };
    },
    toJSON(message) {
        const obj = {};
        if (message.current) {
            obj.current = message.current.map((e) => Math.round(e));
        }
        else {
            obj.current = [];
        }
        if (message.wanted) {
            obj.wanted = message.wanted.map((e) => Math.round(e));
        }
        else {
            obj.wanted = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = createBaseQueryRarityDistributionResponse();
        message.current = object.current?.map((e) => e) || [];
        message.wanted = object.wanted?.map((e) => e) || [];
        return message;
    },
};
function createBaseQueryQCouncilsRequest() {
    return { status: 0, voters: [], card: 0, creator: "", ignore: undefined };
}
exports.QueryQCouncilsRequest = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.status !== 0) {
            writer.uint32(8).int32(message.status);
        }
        for (const v of message.voters) {
            writer.uint32(26).string(v);
        }
        if (message.card !== 0) {
            writer.uint32(32).uint64(message.card);
        }
        if (message.creator !== "") {
            writer.uint32(42).string(message.creator);
        }
        if (message.ignore !== undefined) {
            exports.IgnoreCouncils.encode(message.ignore, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseQueryQCouncilsRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.status = reader.int32();
                    break;
                case 3:
                    message.voters.push(reader.string());
                    break;
                case 4:
                    message.card = longToNumber(reader.uint64());
                    break;
                case 5:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.ignore = exports.IgnoreCouncils.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return {
            status: isSet(object.status) ? (0, council_1.councelingStatusFromJSON)(object.status) : 0,
            voters: Array.isArray(object?.voters) ? object.voters.map((e) => String(e)) : [],
            card: isSet(object.card) ? Number(object.card) : 0,
            creator: isSet(object.creator) ? String(object.creator) : "",
            ignore: isSet(object.ignore) ? exports.IgnoreCouncils.fromJSON(object.ignore) : undefined,
        };
    },
    toJSON(message) {
        const obj = {};
        message.status !== undefined && (obj.status = (0, council_1.councelingStatusToJSON)(message.status));
        if (message.voters) {
            obj.voters = message.voters.map((e) => e);
        }
        else {
            obj.voters = [];
        }
        message.card !== undefined && (obj.card = Math.round(message.card));
        message.creator !== undefined && (obj.creator = message.creator);
        message.ignore !== undefined && (obj.ignore = message.ignore ? exports.IgnoreCouncils.toJSON(message.ignore) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseQueryQCouncilsRequest();
        message.status = object.status ?? 0;
        message.voters = object.voters?.map((e) => e) || [];
        message.card = object.card ?? 0;
        message.creator = object.creator ?? "";
        message.ignore = (object.ignore !== undefined && object.ignore !== null)
            ? exports.IgnoreCouncils.fromPartial(object.ignore)
            : undefined;
        return message;
    },
};
function createBaseIgnoreCouncils() {
    return { status: false, card: false };
}
exports.IgnoreCouncils = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.status === true) {
            writer.uint32(8).bool(message.status);
        }
        if (message.card === true) {
            writer.uint32(16).bool(message.card);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseIgnoreCouncils();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.status = reader.bool();
                    break;
                case 2:
                    message.card = reader.bool();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return {
            status: isSet(object.status) ? Boolean(object.status) : false,
            card: isSet(object.card) ? Boolean(object.card) : false,
        };
    },
    toJSON(message) {
        const obj = {};
        message.status !== undefined && (obj.status = message.status);
        message.card !== undefined && (obj.card = message.card);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseIgnoreCouncils();
        message.status = object.status ?? false;
        message.card = object.card ?? false;
        return message;
    },
};
function createBaseQueryQCouncilsResponse() {
    return { councilssIds: [], councils: [] };
}
exports.QueryQCouncilsResponse = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        writer.uint32(10).fork();
        for (const v of message.councilssIds) {
            writer.uint64(v);
        }
        writer.ldelim();
        for (const v of message.councils) {
            council_1.Council.encode(v, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseQueryQCouncilsResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    if ((tag & 7) === 2) {
                        const end2 = reader.uint32() + reader.pos;
                        while (reader.pos < end2) {
                            message.councilssIds.push(longToNumber(reader.uint64()));
                        }
                    }
                    else {
                        message.councilssIds.push(longToNumber(reader.uint64()));
                    }
                    break;
                case 2:
                    message.councils.push(council_1.Council.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return {
            councilssIds: Array.isArray(object?.councilssIds) ? object.councilssIds.map((e) => Number(e)) : [],
            councils: Array.isArray(object?.councils) ? object.councils.map((e) => council_1.Council.fromJSON(e)) : [],
        };
    },
    toJSON(message) {
        const obj = {};
        if (message.councilssIds) {
            obj.councilssIds = message.councilssIds.map((e) => Math.round(e));
        }
        else {
            obj.councilssIds = [];
        }
        if (message.councils) {
            obj.councils = message.councils.map((e) => e ? council_1.Council.toJSON(e) : undefined);
        }
        else {
            obj.councils = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = createBaseQueryQCouncilsResponse();
        message.councilssIds = object.councilssIds?.map((e) => e) || [];
        message.councils = object.councils?.map((e) => council_1.Council.fromPartial(e)) || [];
        return message;
    },
};
class QueryClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
        this.Params = this.Params.bind(this);
        this.QCard = this.QCard.bind(this);
        this.QCardContent = this.QCardContent.bind(this);
        this.QUser = this.QUser.bind(this);
        this.QCardchainInfo = this.QCardchainInfo.bind(this);
        this.QVotingResults = this.QVotingResults.bind(this);
        this.QVotableCards = this.QVotableCards.bind(this);
        this.QCards = this.QCards.bind(this);
        this.QMatch = this.QMatch.bind(this);
        this.QCollection = this.QCollection.bind(this);
        this.QSellOffer = this.QSellOffer.bind(this);
        this.QCouncil = this.QCouncil.bind(this);
        this.QMatches = this.QMatches.bind(this);
        this.QSellOffers = this.QSellOffers.bind(this);
        this.QServer = this.QServer.bind(this);
        this.QCollections = this.QCollections.bind(this);
        this.RarityDistribution = this.RarityDistribution.bind(this);
        this.QCouncils = this.QCouncils.bind(this);
    }
    Params(request) {
        const data = exports.QueryParamsRequest.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "Params", data);
        return promise.then((data) => exports.QueryParamsResponse.decode(new minimal_1.default.Reader(data)));
    }
    QCard(request) {
        const data = exports.QueryQCardRequest.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QCard", data);
        return promise.then((data) => card_1.OutpCard.decode(new minimal_1.default.Reader(data)));
    }
    QCardContent(request) {
        const data = exports.QueryQCardContentRequest.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QCardContent", data);
        return promise.then((data) => exports.QueryQCardContentResponse.decode(new minimal_1.default.Reader(data)));
    }
    QUser(request) {
        const data = exports.QueryQUserRequest.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QUser", data);
        return promise.then((data) => user_1.User.decode(new minimal_1.default.Reader(data)));
    }
    QCardchainInfo(request) {
        const data = exports.QueryQCardchainInfoRequest.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QCardchainInfo", data);
        return promise.then((data) => exports.QueryQCardchainInfoResponse.decode(new minimal_1.default.Reader(data)));
    }
    QVotingResults(request) {
        const data = exports.QueryQVotingResultsRequest.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QVotingResults", data);
        return promise.then((data) => exports.QueryQVotingResultsResponse.decode(new minimal_1.default.Reader(data)));
    }
    QVotableCards(request) {
        const data = exports.QueryQVotableCardsRequest.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QVotableCards", data);
        return promise.then((data) => exports.QueryQVotableCardsResponse.decode(new minimal_1.default.Reader(data)));
    }
    QCards(request) {
        const data = exports.QueryQCardsRequest.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QCards", data);
        return promise.then((data) => exports.QueryQCardsResponse.decode(new minimal_1.default.Reader(data)));
    }
    QMatch(request) {
        const data = exports.QueryQMatchRequest.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QMatch", data);
        return promise.then((data) => match_1.Match.decode(new minimal_1.default.Reader(data)));
    }
    QCollection(request) {
        const data = exports.QueryQCollectionRequest.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QCollection", data);
        return promise.then((data) => collection_1.OutpCollection.decode(new minimal_1.default.Reader(data)));
    }
    QSellOffer(request) {
        const data = exports.QueryQSellOfferRequest.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QSellOffer", data);
        return promise.then((data) => sell_offer_1.SellOffer.decode(new minimal_1.default.Reader(data)));
    }
    QCouncil(request) {
        const data = exports.QueryQCouncilRequest.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QCouncil", data);
        return promise.then((data) => council_1.Council.decode(new minimal_1.default.Reader(data)));
    }
    QMatches(request) {
        const data = exports.QueryQMatchesRequest.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QMatches", data);
        return promise.then((data) => exports.QueryQMatchesResponse.decode(new minimal_1.default.Reader(data)));
    }
    QSellOffers(request) {
        const data = exports.QueryQSellOffersRequest.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QSellOffers", data);
        return promise.then((data) => exports.QueryQSellOffersResponse.decode(new minimal_1.default.Reader(data)));
    }
    QServer(request) {
        const data = exports.QueryQServerRequest.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QServer", data);
        return promise.then((data) => server_1.Server.decode(new minimal_1.default.Reader(data)));
    }
    QCollections(request) {
        const data = exports.QueryQCollectionsRequest.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QCollections", data);
        return promise.then((data) => exports.QueryQCollectionsResponse.decode(new minimal_1.default.Reader(data)));
    }
    RarityDistribution(request) {
        const data = exports.QueryRarityDistributionRequest.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "RarityDistribution", data);
        return promise.then((data) => exports.QueryRarityDistributionResponse.decode(new minimal_1.default.Reader(data)));
    }
    QCouncils(request) {
        const data = exports.QueryQCouncilsRequest.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QCouncils", data);
        return promise.then((data) => exports.QueryQCouncilsResponse.decode(new minimal_1.default.Reader(data)));
    }
}
exports.QueryClientImpl = QueryClientImpl;
var globalThis = (() => {
    if (typeof globalThis !== "undefined") {
        return globalThis;
    }
    if (typeof self !== "undefined") {
        return self;
    }
    if (typeof window !== "undefined") {
        return window;
    }
    if (typeof global !== "undefined") {
        return global;
    }
    throw "Unable to locate global object";
})();
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
    }
    return long.toNumber();
}
if (minimal_1.default.util.Long !== long_1.default) {
    minimal_1.default.util.Long = long_1.default;
    minimal_1.default.configure();
}
function isSet(value) {
    return value !== null && value !== undefined;
}
