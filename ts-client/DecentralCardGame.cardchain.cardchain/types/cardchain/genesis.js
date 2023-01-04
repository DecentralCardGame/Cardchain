"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.GenesisState = exports.protobufPackage = void 0;
/* eslint-disable */
const minimal_1 = require("protobufjs/minimal");
const coin_1 = require("../cosmos/base/v1beta1/coin");
const card_1 = require("./card");
const collection_1 = require("./collection");
const council_1 = require("./council");
const image_1 = require("./image");
const match_1 = require("./match");
const params_1 = require("./params");
const running_average_1 = require("./running_average");
const sell_offer_1 = require("./sell_offer");
const server_1 = require("./server");
const user_1 = require("./user");
exports.protobufPackage = "DecentralCardGame.cardchain.cardchain";
function createBaseGenesisState() {
    return {
        params: undefined,
        cardRecords: [],
        users: [],
        addresses: [],
        matches: [],
        collections: [],
        sellOffers: [],
        pools: [],
        cardAuctionPrice: "",
        councils: [],
        RunningAverages: [],
        images: [],
        Servers: [],
    };
}
exports.GenesisState = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.params !== undefined) {
            params_1.Params.encode(message.params, writer.uint32(10).fork()).ldelim();
        }
        for (const v of message.cardRecords) {
            card_1.Card.encode(v, writer.uint32(18).fork()).ldelim();
        }
        for (const v of message.users) {
            user_1.User.encode(v, writer.uint32(26).fork()).ldelim();
        }
        for (const v of message.addresses) {
            writer.uint32(34).string(v);
        }
        for (const v of message.matches) {
            match_1.Match.encode(v, writer.uint32(50).fork()).ldelim();
        }
        for (const v of message.collections) {
            collection_1.Collection.encode(v, writer.uint32(58).fork()).ldelim();
        }
        for (const v of message.sellOffers) {
            sell_offer_1.SellOffer.encode(v, writer.uint32(66).fork()).ldelim();
        }
        for (const v of message.pools) {
            coin_1.Coin.encode(v, writer.uint32(74).fork()).ldelim();
        }
        if (message.cardAuctionPrice !== "") {
            writer.uint32(90).string(message.cardAuctionPrice);
        }
        for (const v of message.councils) {
            council_1.Council.encode(v, writer.uint32(98).fork()).ldelim();
        }
        for (const v of message.RunningAverages) {
            running_average_1.RunningAverage.encode(v, writer.uint32(106).fork()).ldelim();
        }
        for (const v of message.images) {
            image_1.Image.encode(v, writer.uint32(114).fork()).ldelim();
        }
        for (const v of message.Servers) {
            server_1.Server.encode(v, writer.uint32(122).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseGenesisState();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.params = params_1.Params.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.cardRecords.push(card_1.Card.decode(reader, reader.uint32()));
                    break;
                case 3:
                    message.users.push(user_1.User.decode(reader, reader.uint32()));
                    break;
                case 4:
                    message.addresses.push(reader.string());
                    break;
                case 6:
                    message.matches.push(match_1.Match.decode(reader, reader.uint32()));
                    break;
                case 7:
                    message.collections.push(collection_1.Collection.decode(reader, reader.uint32()));
                    break;
                case 8:
                    message.sellOffers.push(sell_offer_1.SellOffer.decode(reader, reader.uint32()));
                    break;
                case 9:
                    message.pools.push(coin_1.Coin.decode(reader, reader.uint32()));
                    break;
                case 11:
                    message.cardAuctionPrice = reader.string();
                    break;
                case 12:
                    message.councils.push(council_1.Council.decode(reader, reader.uint32()));
                    break;
                case 13:
                    message.RunningAverages.push(running_average_1.RunningAverage.decode(reader, reader.uint32()));
                    break;
                case 14:
                    message.images.push(image_1.Image.decode(reader, reader.uint32()));
                    break;
                case 15:
                    message.Servers.push(server_1.Server.decode(reader, reader.uint32()));
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
            params: isSet(object.params) ? params_1.Params.fromJSON(object.params) : undefined,
            cardRecords: Array.isArray(object?.cardRecords) ? object.cardRecords.map((e) => card_1.Card.fromJSON(e)) : [],
            users: Array.isArray(object?.users) ? object.users.map((e) => user_1.User.fromJSON(e)) : [],
            addresses: Array.isArray(object?.addresses) ? object.addresses.map((e) => String(e)) : [],
            matches: Array.isArray(object?.matches) ? object.matches.map((e) => match_1.Match.fromJSON(e)) : [],
            collections: Array.isArray(object?.collections) ? object.collections.map((e) => collection_1.Collection.fromJSON(e)) : [],
            sellOffers: Array.isArray(object?.sellOffers) ? object.sellOffers.map((e) => sell_offer_1.SellOffer.fromJSON(e)) : [],
            pools: Array.isArray(object?.pools) ? object.pools.map((e) => coin_1.Coin.fromJSON(e)) : [],
            cardAuctionPrice: isSet(object.cardAuctionPrice) ? String(object.cardAuctionPrice) : "",
            councils: Array.isArray(object?.councils) ? object.councils.map((e) => council_1.Council.fromJSON(e)) : [],
            RunningAverages: Array.isArray(object?.RunningAverages)
                ? object.RunningAverages.map((e) => running_average_1.RunningAverage.fromJSON(e))
                : [],
            images: Array.isArray(object?.images) ? object.images.map((e) => image_1.Image.fromJSON(e)) : [],
            Servers: Array.isArray(object?.Servers) ? object.Servers.map((e) => server_1.Server.fromJSON(e)) : [],
        };
    },
    toJSON(message) {
        const obj = {};
        message.params !== undefined && (obj.params = message.params ? params_1.Params.toJSON(message.params) : undefined);
        if (message.cardRecords) {
            obj.cardRecords = message.cardRecords.map((e) => e ? card_1.Card.toJSON(e) : undefined);
        }
        else {
            obj.cardRecords = [];
        }
        if (message.users) {
            obj.users = message.users.map((e) => e ? user_1.User.toJSON(e) : undefined);
        }
        else {
            obj.users = [];
        }
        if (message.addresses) {
            obj.addresses = message.addresses.map((e) => e);
        }
        else {
            obj.addresses = [];
        }
        if (message.matches) {
            obj.matches = message.matches.map((e) => e ? match_1.Match.toJSON(e) : undefined);
        }
        else {
            obj.matches = [];
        }
        if (message.collections) {
            obj.collections = message.collections.map((e) => e ? collection_1.Collection.toJSON(e) : undefined);
        }
        else {
            obj.collections = [];
        }
        if (message.sellOffers) {
            obj.sellOffers = message.sellOffers.map((e) => e ? sell_offer_1.SellOffer.toJSON(e) : undefined);
        }
        else {
            obj.sellOffers = [];
        }
        if (message.pools) {
            obj.pools = message.pools.map((e) => e ? coin_1.Coin.toJSON(e) : undefined);
        }
        else {
            obj.pools = [];
        }
        message.cardAuctionPrice !== undefined && (obj.cardAuctionPrice = message.cardAuctionPrice);
        if (message.councils) {
            obj.councils = message.councils.map((e) => e ? council_1.Council.toJSON(e) : undefined);
        }
        else {
            obj.councils = [];
        }
        if (message.RunningAverages) {
            obj.RunningAverages = message.RunningAverages.map((e) => e ? running_average_1.RunningAverage.toJSON(e) : undefined);
        }
        else {
            obj.RunningAverages = [];
        }
        if (message.images) {
            obj.images = message.images.map((e) => e ? image_1.Image.toJSON(e) : undefined);
        }
        else {
            obj.images = [];
        }
        if (message.Servers) {
            obj.Servers = message.Servers.map((e) => e ? server_1.Server.toJSON(e) : undefined);
        }
        else {
            obj.Servers = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = createBaseGenesisState();
        message.params = (object.params !== undefined && object.params !== null)
            ? params_1.Params.fromPartial(object.params)
            : undefined;
        message.cardRecords = object.cardRecords?.map((e) => card_1.Card.fromPartial(e)) || [];
        message.users = object.users?.map((e) => user_1.User.fromPartial(e)) || [];
        message.addresses = object.addresses?.map((e) => e) || [];
        message.matches = object.matches?.map((e) => match_1.Match.fromPartial(e)) || [];
        message.collections = object.collections?.map((e) => collection_1.Collection.fromPartial(e)) || [];
        message.sellOffers = object.sellOffers?.map((e) => sell_offer_1.SellOffer.fromPartial(e)) || [];
        message.pools = object.pools?.map((e) => coin_1.Coin.fromPartial(e)) || [];
        message.cardAuctionPrice = object.cardAuctionPrice ?? "";
        message.councils = object.councils?.map((e) => council_1.Council.fromPartial(e)) || [];
        message.RunningAverages = object.RunningAverages?.map((e) => running_average_1.RunningAverage.fromPartial(e)) || [];
        message.images = object.images?.map((e) => image_1.Image.fromPartial(e)) || [];
        message.Servers = object.Servers?.map((e) => server_1.Server.fromPartial(e)) || [];
        return message;
    },
};
function isSet(value) {
    return value !== null && value !== undefined;
}
