"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Params = exports.protobufPackage = void 0;
/* eslint-disable */
const long_1 = require("long");
const minimal_1 = require("protobufjs/minimal");
exports.protobufPackage = "DecentralCardGame.cardchain.cardchain";
function createBaseParams() {
    return {
        votingRightsExpirationTime: 0,
        collectionSize: 0,
        collectionPrice: "",
        activeCollectionsAmount: 0,
        collectionCreationFee: "",
        collateralDeposit: "",
        winnerReward: 0,
        voterReward: 0,
        hourlyFaucet: "",
        inflationRate: "",
        raresPerPack: 0,
        commonsPerPack: 0,
        unCommonsPerPack: 0,
        trialPeriod: 0,
        gameVoteRatio: 0,
        cardAuctionPriceReductionPeriod: 0,
        airDropValue: "",
        airDropMaxBlockHeight: 0,
    };
}
exports.Params = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.votingRightsExpirationTime !== 0) {
            writer.uint32(8).int64(message.votingRightsExpirationTime);
        }
        if (message.collectionSize !== 0) {
            writer.uint32(16).uint64(message.collectionSize);
        }
        if (message.collectionPrice !== "") {
            writer.uint32(26).string(message.collectionPrice);
        }
        if (message.activeCollectionsAmount !== 0) {
            writer.uint32(32).uint64(message.activeCollectionsAmount);
        }
        if (message.collectionCreationFee !== "") {
            writer.uint32(42).string(message.collectionCreationFee);
        }
        if (message.collateralDeposit !== "") {
            writer.uint32(50).string(message.collateralDeposit);
        }
        if (message.winnerReward !== 0) {
            writer.uint32(56).int64(message.winnerReward);
        }
        if (message.voterReward !== 0) {
            writer.uint32(64).int64(message.voterReward);
        }
        if (message.hourlyFaucet !== "") {
            writer.uint32(74).string(message.hourlyFaucet);
        }
        if (message.inflationRate !== "") {
            writer.uint32(82).string(message.inflationRate);
        }
        if (message.raresPerPack !== 0) {
            writer.uint32(88).uint64(message.raresPerPack);
        }
        if (message.commonsPerPack !== 0) {
            writer.uint32(96).uint64(message.commonsPerPack);
        }
        if (message.unCommonsPerPack !== 0) {
            writer.uint32(104).uint64(message.unCommonsPerPack);
        }
        if (message.trialPeriod !== 0) {
            writer.uint32(112).uint64(message.trialPeriod);
        }
        if (message.gameVoteRatio !== 0) {
            writer.uint32(120).int64(message.gameVoteRatio);
        }
        if (message.cardAuctionPriceReductionPeriod !== 0) {
            writer.uint32(128).int64(message.cardAuctionPriceReductionPeriod);
        }
        if (message.airDropValue !== "") {
            writer.uint32(138).string(message.airDropValue);
        }
        if (message.airDropMaxBlockHeight !== 0) {
            writer.uint32(144).int64(message.airDropMaxBlockHeight);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseParams();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.votingRightsExpirationTime = longToNumber(reader.int64());
                    break;
                case 2:
                    message.collectionSize = longToNumber(reader.uint64());
                    break;
                case 3:
                    message.collectionPrice = reader.string();
                    break;
                case 4:
                    message.activeCollectionsAmount = longToNumber(reader.uint64());
                    break;
                case 5:
                    message.collectionCreationFee = reader.string();
                    break;
                case 6:
                    message.collateralDeposit = reader.string();
                    break;
                case 7:
                    message.winnerReward = longToNumber(reader.int64());
                    break;
                case 8:
                    message.voterReward = longToNumber(reader.int64());
                    break;
                case 9:
                    message.hourlyFaucet = reader.string();
                    break;
                case 10:
                    message.inflationRate = reader.string();
                    break;
                case 11:
                    message.raresPerPack = longToNumber(reader.uint64());
                    break;
                case 12:
                    message.commonsPerPack = longToNumber(reader.uint64());
                    break;
                case 13:
                    message.unCommonsPerPack = longToNumber(reader.uint64());
                    break;
                case 14:
                    message.trialPeriod = longToNumber(reader.uint64());
                    break;
                case 15:
                    message.gameVoteRatio = longToNumber(reader.int64());
                    break;
                case 16:
                    message.cardAuctionPriceReductionPeriod = longToNumber(reader.int64());
                    break;
                case 17:
                    message.airDropValue = reader.string();
                    break;
                case 18:
                    message.airDropMaxBlockHeight = longToNumber(reader.int64());
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
            votingRightsExpirationTime: isSet(object.votingRightsExpirationTime)
                ? Number(object.votingRightsExpirationTime)
                : 0,
            collectionSize: isSet(object.collectionSize) ? Number(object.collectionSize) : 0,
            collectionPrice: isSet(object.collectionPrice) ? String(object.collectionPrice) : "",
            activeCollectionsAmount: isSet(object.activeCollectionsAmount) ? Number(object.activeCollectionsAmount) : 0,
            collectionCreationFee: isSet(object.collectionCreationFee) ? String(object.collectionCreationFee) : "",
            collateralDeposit: isSet(object.collateralDeposit) ? String(object.collateralDeposit) : "",
            winnerReward: isSet(object.winnerReward) ? Number(object.winnerReward) : 0,
            voterReward: isSet(object.voterReward) ? Number(object.voterReward) : 0,
            hourlyFaucet: isSet(object.hourlyFaucet) ? String(object.hourlyFaucet) : "",
            inflationRate: isSet(object.inflationRate) ? String(object.inflationRate) : "",
            raresPerPack: isSet(object.raresPerPack) ? Number(object.raresPerPack) : 0,
            commonsPerPack: isSet(object.commonsPerPack) ? Number(object.commonsPerPack) : 0,
            unCommonsPerPack: isSet(object.unCommonsPerPack) ? Number(object.unCommonsPerPack) : 0,
            trialPeriod: isSet(object.trialPeriod) ? Number(object.trialPeriod) : 0,
            gameVoteRatio: isSet(object.gameVoteRatio) ? Number(object.gameVoteRatio) : 0,
            cardAuctionPriceReductionPeriod: isSet(object.cardAuctionPriceReductionPeriod)
                ? Number(object.cardAuctionPriceReductionPeriod)
                : 0,
            airDropValue: isSet(object.airDropValue) ? String(object.airDropValue) : "",
            airDropMaxBlockHeight: isSet(object.airDropMaxBlockHeight) ? Number(object.airDropMaxBlockHeight) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        message.votingRightsExpirationTime !== undefined
            && (obj.votingRightsExpirationTime = Math.round(message.votingRightsExpirationTime));
        message.collectionSize !== undefined && (obj.collectionSize = Math.round(message.collectionSize));
        message.collectionPrice !== undefined && (obj.collectionPrice = message.collectionPrice);
        message.activeCollectionsAmount !== undefined
            && (obj.activeCollectionsAmount = Math.round(message.activeCollectionsAmount));
        message.collectionCreationFee !== undefined && (obj.collectionCreationFee = message.collectionCreationFee);
        message.collateralDeposit !== undefined && (obj.collateralDeposit = message.collateralDeposit);
        message.winnerReward !== undefined && (obj.winnerReward = Math.round(message.winnerReward));
        message.voterReward !== undefined && (obj.voterReward = Math.round(message.voterReward));
        message.hourlyFaucet !== undefined && (obj.hourlyFaucet = message.hourlyFaucet);
        message.inflationRate !== undefined && (obj.inflationRate = message.inflationRate);
        message.raresPerPack !== undefined && (obj.raresPerPack = Math.round(message.raresPerPack));
        message.commonsPerPack !== undefined && (obj.commonsPerPack = Math.round(message.commonsPerPack));
        message.unCommonsPerPack !== undefined && (obj.unCommonsPerPack = Math.round(message.unCommonsPerPack));
        message.trialPeriod !== undefined && (obj.trialPeriod = Math.round(message.trialPeriod));
        message.gameVoteRatio !== undefined && (obj.gameVoteRatio = Math.round(message.gameVoteRatio));
        message.cardAuctionPriceReductionPeriod !== undefined
            && (obj.cardAuctionPriceReductionPeriod = Math.round(message.cardAuctionPriceReductionPeriod));
        message.airDropValue !== undefined && (obj.airDropValue = message.airDropValue);
        message.airDropMaxBlockHeight !== undefined
            && (obj.airDropMaxBlockHeight = Math.round(message.airDropMaxBlockHeight));
        return obj;
    },
    fromPartial(object) {
        const message = createBaseParams();
        message.votingRightsExpirationTime = object.votingRightsExpirationTime ?? 0;
        message.collectionSize = object.collectionSize ?? 0;
        message.collectionPrice = object.collectionPrice ?? "";
        message.activeCollectionsAmount = object.activeCollectionsAmount ?? 0;
        message.collectionCreationFee = object.collectionCreationFee ?? "";
        message.collateralDeposit = object.collateralDeposit ?? "";
        message.winnerReward = object.winnerReward ?? 0;
        message.voterReward = object.voterReward ?? 0;
        message.hourlyFaucet = object.hourlyFaucet ?? "";
        message.inflationRate = object.inflationRate ?? "";
        message.raresPerPack = object.raresPerPack ?? 0;
        message.commonsPerPack = object.commonsPerPack ?? 0;
        message.unCommonsPerPack = object.unCommonsPerPack ?? 0;
        message.trialPeriod = object.trialPeriod ?? 0;
        message.gameVoteRatio = object.gameVoteRatio ?? 0;
        message.cardAuctionPriceReductionPeriod = object.cardAuctionPriceReductionPeriod ?? 0;
        message.airDropValue = object.airDropValue ?? "";
        message.airDropMaxBlockHeight = object.airDropMaxBlockHeight ?? 0;
        return message;
    },
};
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
