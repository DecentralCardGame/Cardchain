"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.MsgBuyCardResponse = exports.MsgBuyCard = exports.MsgCreateSellOfferResponse = exports.MsgCreateSellOffer = exports.MsgSubmitCollectionProposalResponse = exports.MsgSubmitCollectionProposal = exports.MsgAddContributorToCollectionResponse = exports.MsgAddContributorToCollection = exports.MsgRemoveContributorFromCollectionResponse = exports.MsgRemoveContributorFromCollection = exports.MsgRemoveCardFromCollectionResponse = exports.MsgRemoveCardFromCollection = exports.MsgBuyCollectionResponse = exports.MsgBuyCollection = exports.MsgFinalizeCollectionResponse = exports.MsgFinalizeCollection = exports.MsgAddCardToCollectionResponse = exports.MsgAddCardToCollection = exports.MsgCreateCollectionResponse = exports.MsgCreateCollection = exports.MsgApointMatchReporterResponse = exports.MsgApointMatchReporter = exports.MsgSubmitMatchReporterProposalResponse = exports.MsgSubmitMatchReporterProposal = exports.MsgReportMatchResponse = exports.MsgReportMatch = exports.MsgRegisterForCouncilResponse = exports.MsgRegisterForCouncil = exports.MsgChangeArtistResponse = exports.MsgChangeArtist = exports.MsgSubmitCopyrightProposalResponse = exports.MsgSubmitCopyrightProposal = exports.MsgAddArtworkResponse = exports.MsgAddArtwork = exports.MsgDonateToCardResponse = exports.MsgDonateToCard = exports.MsgTransferCardResponse = exports.MsgTransferCard = exports.MsgSaveCardContentResponse = exports.MsgSaveCardContent = exports.MsgVoteCardResponse = exports.MsgVoteCard = exports.MsgBuyCardSchemeResponse = exports.MsgBuyCardScheme = exports.MsgCreateuserResponse = exports.MsgCreateuser = exports.outcomeToJSON = exports.outcomeFromJSON = exports.Outcome = exports.protobufPackage = void 0;
exports.MsgClientImpl = exports.MsgSetUserBiographyResponse = exports.MsgSetUserBiography = exports.MsgSetUserWebsiteResponse = exports.MsgSetUserWebsite = exports.MsgSetCollectionArtistResponse = exports.MsgSetCollectionArtist = exports.MsgSetCollectionStoryWriterResponse = exports.MsgSetCollectionStoryWriter = exports.MsgTransferBoosterPackResponse = exports.MsgTransferBoosterPack = exports.MsgOpenBoosterPackResponse = exports.MsgOpenBoosterPack = exports.MsgSetProfileCardResponse = exports.MsgSetProfileCard = exports.MsgConfirmMatchResponse = exports.MsgConfirmMatch = exports.MsgRewokeCouncilRegistrationResponse = exports.MsgRewokeCouncilRegistration = exports.MsgRestartCouncilResponse = exports.MsgRestartCouncil = exports.MsgRevealCouncilResponseResponse = exports.MsgRevealCouncilResponse = exports.MsgCommitCouncilResponseResponse = exports.MsgCommitCouncilResponse = exports.MsgCreateCouncilResponse = exports.MsgCreateCouncil = exports.MsgSetCardRarityResponse = exports.MsgSetCardRarity = exports.MsgAddStoryToCollectionResponse = exports.MsgAddStoryToCollection = exports.MsgAddArtworkToCollectionResponse = exports.MsgAddArtworkToCollection = exports.MsgRemoveSellOfferResponse = exports.MsgRemoveSellOffer = void 0;
/* eslint-disable */
const long_1 = require("long");
const minimal_1 = require("protobufjs/minimal");
const council_1 = require("./council");
exports.protobufPackage = "DecentralCardGame.cardchain.cardchain";
var Outcome;
(function (Outcome) {
    Outcome[Outcome["AWon"] = 0] = "AWon";
    Outcome[Outcome["BWon"] = 1] = "BWon";
    Outcome[Outcome["Draw"] = 2] = "Draw";
    Outcome[Outcome["Aborted"] = 3] = "Aborted";
    Outcome[Outcome["UNRECOGNIZED"] = -1] = "UNRECOGNIZED";
})(Outcome = exports.Outcome || (exports.Outcome = {}));
function outcomeFromJSON(object) {
    switch (object) {
        case 0:
        case "AWon":
            return Outcome.AWon;
        case 1:
        case "BWon":
            return Outcome.BWon;
        case 2:
        case "Draw":
            return Outcome.Draw;
        case 3:
        case "Aborted":
            return Outcome.Aborted;
        case -1:
        case "UNRECOGNIZED":
        default:
            return Outcome.UNRECOGNIZED;
    }
}
exports.outcomeFromJSON = outcomeFromJSON;
function outcomeToJSON(object) {
    switch (object) {
        case Outcome.AWon:
            return "AWon";
        case Outcome.BWon:
            return "BWon";
        case Outcome.Draw:
            return "Draw";
        case Outcome.Aborted:
            return "Aborted";
        case Outcome.UNRECOGNIZED:
        default:
            return "UNRECOGNIZED";
    }
}
exports.outcomeToJSON = outcomeToJSON;
function createBaseMsgCreateuser() {
    return { creator: "", newUser: "", alias: "" };
}
exports.MsgCreateuser = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.newUser !== "") {
            writer.uint32(18).string(message.newUser);
        }
        if (message.alias !== "") {
            writer.uint32(26).string(message.alias);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgCreateuser();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.newUser = reader.string();
                    break;
                case 3:
                    message.alias = reader.string();
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
            creator: isSet(object.creator) ? String(object.creator) : "",
            newUser: isSet(object.newUser) ? String(object.newUser) : "",
            alias: isSet(object.alias) ? String(object.alias) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.newUser !== undefined && (obj.newUser = message.newUser);
        message.alias !== undefined && (obj.alias = message.alias);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgCreateuser();
        message.creator = object.creator ?? "";
        message.newUser = object.newUser ?? "";
        message.alias = object.alias ?? "";
        return message;
    },
};
function createBaseMsgCreateuserResponse() {
    return {};
}
exports.MsgCreateuserResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgCreateuserResponse();
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
        const message = createBaseMsgCreateuserResponse();
        return message;
    },
};
function createBaseMsgBuyCardScheme() {
    return { creator: "", bid: "" };
}
exports.MsgBuyCardScheme = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.bid !== "") {
            writer.uint32(18).string(message.bid);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgBuyCardScheme();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.bid = reader.string();
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
            creator: isSet(object.creator) ? String(object.creator) : "",
            bid: isSet(object.bid) ? String(object.bid) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.bid !== undefined && (obj.bid = message.bid);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgBuyCardScheme();
        message.creator = object.creator ?? "";
        message.bid = object.bid ?? "";
        return message;
    },
};
function createBaseMsgBuyCardSchemeResponse() {
    return {};
}
exports.MsgBuyCardSchemeResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgBuyCardSchemeResponse();
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
        const message = createBaseMsgBuyCardSchemeResponse();
        return message;
    },
};
function createBaseMsgVoteCard() {
    return { creator: "", cardId: 0, voteType: "" };
}
exports.MsgVoteCard = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.cardId !== 0) {
            writer.uint32(16).uint64(message.cardId);
        }
        if (message.voteType !== "") {
            writer.uint32(26).string(message.voteType);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgVoteCard();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.cardId = longToNumber(reader.uint64());
                    break;
                case 3:
                    message.voteType = reader.string();
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
            creator: isSet(object.creator) ? String(object.creator) : "",
            cardId: isSet(object.cardId) ? Number(object.cardId) : 0,
            voteType: isSet(object.voteType) ? String(object.voteType) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.cardId !== undefined && (obj.cardId = Math.round(message.cardId));
        message.voteType !== undefined && (obj.voteType = message.voteType);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgVoteCard();
        message.creator = object.creator ?? "";
        message.cardId = object.cardId ?? 0;
        message.voteType = object.voteType ?? "";
        return message;
    },
};
function createBaseMsgVoteCardResponse() {
    return { airdropClaimed: false };
}
exports.MsgVoteCardResponse = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.airdropClaimed === true) {
            writer.uint32(8).bool(message.airdropClaimed);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgVoteCardResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.airdropClaimed = reader.bool();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return { airdropClaimed: isSet(object.airdropClaimed) ? Boolean(object.airdropClaimed) : false };
    },
    toJSON(message) {
        const obj = {};
        message.airdropClaimed !== undefined && (obj.airdropClaimed = message.airdropClaimed);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgVoteCardResponse();
        message.airdropClaimed = object.airdropClaimed ?? false;
        return message;
    },
};
function createBaseMsgSaveCardContent() {
    return { creator: "", cardId: 0, content: new Uint8Array(), notes: "", artist: "" };
}
exports.MsgSaveCardContent = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.cardId !== 0) {
            writer.uint32(16).uint64(message.cardId);
        }
        if (message.content.length !== 0) {
            writer.uint32(26).bytes(message.content);
        }
        if (message.notes !== "") {
            writer.uint32(34).string(message.notes);
        }
        if (message.artist !== "") {
            writer.uint32(42).string(message.artist);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgSaveCardContent();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.cardId = longToNumber(reader.uint64());
                    break;
                case 3:
                    message.content = reader.bytes();
                    break;
                case 4:
                    message.notes = reader.string();
                    break;
                case 5:
                    message.artist = reader.string();
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
            creator: isSet(object.creator) ? String(object.creator) : "",
            cardId: isSet(object.cardId) ? Number(object.cardId) : 0,
            content: isSet(object.content) ? bytesFromBase64(object.content) : new Uint8Array(),
            notes: isSet(object.notes) ? String(object.notes) : "",
            artist: isSet(object.artist) ? String(object.artist) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.cardId !== undefined && (obj.cardId = Math.round(message.cardId));
        message.content !== undefined
            && (obj.content = base64FromBytes(message.content !== undefined ? message.content : new Uint8Array()));
        message.notes !== undefined && (obj.notes = message.notes);
        message.artist !== undefined && (obj.artist = message.artist);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgSaveCardContent();
        message.creator = object.creator ?? "";
        message.cardId = object.cardId ?? 0;
        message.content = object.content ?? new Uint8Array();
        message.notes = object.notes ?? "";
        message.artist = object.artist ?? "";
        return message;
    },
};
function createBaseMsgSaveCardContentResponse() {
    return { airdropClaimed: false };
}
exports.MsgSaveCardContentResponse = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.airdropClaimed === true) {
            writer.uint32(8).bool(message.airdropClaimed);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgSaveCardContentResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.airdropClaimed = reader.bool();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return { airdropClaimed: isSet(object.airdropClaimed) ? Boolean(object.airdropClaimed) : false };
    },
    toJSON(message) {
        const obj = {};
        message.airdropClaimed !== undefined && (obj.airdropClaimed = message.airdropClaimed);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgSaveCardContentResponse();
        message.airdropClaimed = object.airdropClaimed ?? false;
        return message;
    },
};
function createBaseMsgTransferCard() {
    return { creator: "", cardId: 0, receiver: "" };
}
exports.MsgTransferCard = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.cardId !== 0) {
            writer.uint32(16).uint64(message.cardId);
        }
        if (message.receiver !== "") {
            writer.uint32(34).string(message.receiver);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgTransferCard();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.cardId = longToNumber(reader.uint64());
                    break;
                case 4:
                    message.receiver = reader.string();
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
            creator: isSet(object.creator) ? String(object.creator) : "",
            cardId: isSet(object.cardId) ? Number(object.cardId) : 0,
            receiver: isSet(object.receiver) ? String(object.receiver) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.cardId !== undefined && (obj.cardId = Math.round(message.cardId));
        message.receiver !== undefined && (obj.receiver = message.receiver);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgTransferCard();
        message.creator = object.creator ?? "";
        message.cardId = object.cardId ?? 0;
        message.receiver = object.receiver ?? "";
        return message;
    },
};
function createBaseMsgTransferCardResponse() {
    return {};
}
exports.MsgTransferCardResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgTransferCardResponse();
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
        const message = createBaseMsgTransferCardResponse();
        return message;
    },
};
function createBaseMsgDonateToCard() {
    return { creator: "", cardId: 0, amount: "" };
}
exports.MsgDonateToCard = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.cardId !== 0) {
            writer.uint32(16).uint64(message.cardId);
        }
        if (message.amount !== "") {
            writer.uint32(26).string(message.amount);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgDonateToCard();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.cardId = longToNumber(reader.uint64());
                    break;
                case 3:
                    message.amount = reader.string();
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
            creator: isSet(object.creator) ? String(object.creator) : "",
            cardId: isSet(object.cardId) ? Number(object.cardId) : 0,
            amount: isSet(object.amount) ? String(object.amount) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.cardId !== undefined && (obj.cardId = Math.round(message.cardId));
        message.amount !== undefined && (obj.amount = message.amount);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgDonateToCard();
        message.creator = object.creator ?? "";
        message.cardId = object.cardId ?? 0;
        message.amount = object.amount ?? "";
        return message;
    },
};
function createBaseMsgDonateToCardResponse() {
    return {};
}
exports.MsgDonateToCardResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgDonateToCardResponse();
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
        const message = createBaseMsgDonateToCardResponse();
        return message;
    },
};
function createBaseMsgAddArtwork() {
    return { creator: "", cardId: 0, image: new Uint8Array(), fullArt: false };
}
exports.MsgAddArtwork = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.cardId !== 0) {
            writer.uint32(16).uint64(message.cardId);
        }
        if (message.image.length !== 0) {
            writer.uint32(26).bytes(message.image);
        }
        if (message.fullArt === true) {
            writer.uint32(32).bool(message.fullArt);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgAddArtwork();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.cardId = longToNumber(reader.uint64());
                    break;
                case 3:
                    message.image = reader.bytes();
                    break;
                case 4:
                    message.fullArt = reader.bool();
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
            creator: isSet(object.creator) ? String(object.creator) : "",
            cardId: isSet(object.cardId) ? Number(object.cardId) : 0,
            image: isSet(object.image) ? bytesFromBase64(object.image) : new Uint8Array(),
            fullArt: isSet(object.fullArt) ? Boolean(object.fullArt) : false,
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.cardId !== undefined && (obj.cardId = Math.round(message.cardId));
        message.image !== undefined
            && (obj.image = base64FromBytes(message.image !== undefined ? message.image : new Uint8Array()));
        message.fullArt !== undefined && (obj.fullArt = message.fullArt);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgAddArtwork();
        message.creator = object.creator ?? "";
        message.cardId = object.cardId ?? 0;
        message.image = object.image ?? new Uint8Array();
        message.fullArt = object.fullArt ?? false;
        return message;
    },
};
function createBaseMsgAddArtworkResponse() {
    return {};
}
exports.MsgAddArtworkResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgAddArtworkResponse();
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
        const message = createBaseMsgAddArtworkResponse();
        return message;
    },
};
function createBaseMsgSubmitCopyrightProposal() {
    return { creator: "", cardId: 0, description: "", link: "" };
}
exports.MsgSubmitCopyrightProposal = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.cardId !== 0) {
            writer.uint32(16).uint64(message.cardId);
        }
        if (message.description !== "") {
            writer.uint32(26).string(message.description);
        }
        if (message.link !== "") {
            writer.uint32(34).string(message.link);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgSubmitCopyrightProposal();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.cardId = longToNumber(reader.uint64());
                    break;
                case 3:
                    message.description = reader.string();
                    break;
                case 4:
                    message.link = reader.string();
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
            creator: isSet(object.creator) ? String(object.creator) : "",
            cardId: isSet(object.cardId) ? Number(object.cardId) : 0,
            description: isSet(object.description) ? String(object.description) : "",
            link: isSet(object.link) ? String(object.link) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.cardId !== undefined && (obj.cardId = Math.round(message.cardId));
        message.description !== undefined && (obj.description = message.description);
        message.link !== undefined && (obj.link = message.link);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgSubmitCopyrightProposal();
        message.creator = object.creator ?? "";
        message.cardId = object.cardId ?? 0;
        message.description = object.description ?? "";
        message.link = object.link ?? "";
        return message;
    },
};
function createBaseMsgSubmitCopyrightProposalResponse() {
    return {};
}
exports.MsgSubmitCopyrightProposalResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgSubmitCopyrightProposalResponse();
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
        const message = createBaseMsgSubmitCopyrightProposalResponse();
        return message;
    },
};
function createBaseMsgChangeArtist() {
    return { creator: "", cardID: 0, artist: "" };
}
exports.MsgChangeArtist = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.cardID !== 0) {
            writer.uint32(16).uint64(message.cardID);
        }
        if (message.artist !== "") {
            writer.uint32(26).string(message.artist);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgChangeArtist();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.cardID = longToNumber(reader.uint64());
                    break;
                case 3:
                    message.artist = reader.string();
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
            creator: isSet(object.creator) ? String(object.creator) : "",
            cardID: isSet(object.cardID) ? Number(object.cardID) : 0,
            artist: isSet(object.artist) ? String(object.artist) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.cardID !== undefined && (obj.cardID = Math.round(message.cardID));
        message.artist !== undefined && (obj.artist = message.artist);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgChangeArtist();
        message.creator = object.creator ?? "";
        message.cardID = object.cardID ?? 0;
        message.artist = object.artist ?? "";
        return message;
    },
};
function createBaseMsgChangeArtistResponse() {
    return {};
}
exports.MsgChangeArtistResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgChangeArtistResponse();
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
        const message = createBaseMsgChangeArtistResponse();
        return message;
    },
};
function createBaseMsgRegisterForCouncil() {
    return { creator: "" };
}
exports.MsgRegisterForCouncil = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgRegisterForCouncil();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return { creator: isSet(object.creator) ? String(object.creator) : "" };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgRegisterForCouncil();
        message.creator = object.creator ?? "";
        return message;
    },
};
function createBaseMsgRegisterForCouncilResponse() {
    return {};
}
exports.MsgRegisterForCouncilResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgRegisterForCouncilResponse();
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
        const message = createBaseMsgRegisterForCouncilResponse();
        return message;
    },
};
function createBaseMsgReportMatch() {
    return { creator: "", playerA: "", playerB: "", cardsA: [], cardsB: [], outcome: 0 };
}
exports.MsgReportMatch = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.playerA !== "") {
            writer.uint32(18).string(message.playerA);
        }
        if (message.playerB !== "") {
            writer.uint32(26).string(message.playerB);
        }
        writer.uint32(42).fork();
        for (const v of message.cardsA) {
            writer.uint64(v);
        }
        writer.ldelim();
        writer.uint32(50).fork();
        for (const v of message.cardsB) {
            writer.uint64(v);
        }
        writer.ldelim();
        if (message.outcome !== 0) {
            writer.uint32(56).int32(message.outcome);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgReportMatch();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.playerA = reader.string();
                    break;
                case 3:
                    message.playerB = reader.string();
                    break;
                case 5:
                    if ((tag & 7) === 2) {
                        const end2 = reader.uint32() + reader.pos;
                        while (reader.pos < end2) {
                            message.cardsA.push(longToNumber(reader.uint64()));
                        }
                    }
                    else {
                        message.cardsA.push(longToNumber(reader.uint64()));
                    }
                    break;
                case 6:
                    if ((tag & 7) === 2) {
                        const end2 = reader.uint32() + reader.pos;
                        while (reader.pos < end2) {
                            message.cardsB.push(longToNumber(reader.uint64()));
                        }
                    }
                    else {
                        message.cardsB.push(longToNumber(reader.uint64()));
                    }
                    break;
                case 7:
                    message.outcome = reader.int32();
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
            creator: isSet(object.creator) ? String(object.creator) : "",
            playerA: isSet(object.playerA) ? String(object.playerA) : "",
            playerB: isSet(object.playerB) ? String(object.playerB) : "",
            cardsA: Array.isArray(object?.cardsA) ? object.cardsA.map((e) => Number(e)) : [],
            cardsB: Array.isArray(object?.cardsB) ? object.cardsB.map((e) => Number(e)) : [],
            outcome: isSet(object.outcome) ? outcomeFromJSON(object.outcome) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.playerA !== undefined && (obj.playerA = message.playerA);
        message.playerB !== undefined && (obj.playerB = message.playerB);
        if (message.cardsA) {
            obj.cardsA = message.cardsA.map((e) => Math.round(e));
        }
        else {
            obj.cardsA = [];
        }
        if (message.cardsB) {
            obj.cardsB = message.cardsB.map((e) => Math.round(e));
        }
        else {
            obj.cardsB = [];
        }
        message.outcome !== undefined && (obj.outcome = outcomeToJSON(message.outcome));
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgReportMatch();
        message.creator = object.creator ?? "";
        message.playerA = object.playerA ?? "";
        message.playerB = object.playerB ?? "";
        message.cardsA = object.cardsA?.map((e) => e) || [];
        message.cardsB = object.cardsB?.map((e) => e) || [];
        message.outcome = object.outcome ?? 0;
        return message;
    },
};
function createBaseMsgReportMatchResponse() {
    return { matchId: 0 };
}
exports.MsgReportMatchResponse = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.matchId !== 0) {
            writer.uint32(8).uint64(message.matchId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgReportMatchResponse();
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
        const message = createBaseMsgReportMatchResponse();
        message.matchId = object.matchId ?? 0;
        return message;
    },
};
function createBaseMsgSubmitMatchReporterProposal() {
    return { creator: "", reporter: "", deposit: "", description: "" };
}
exports.MsgSubmitMatchReporterProposal = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.reporter !== "") {
            writer.uint32(18).string(message.reporter);
        }
        if (message.deposit !== "") {
            writer.uint32(26).string(message.deposit);
        }
        if (message.description !== "") {
            writer.uint32(34).string(message.description);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgSubmitMatchReporterProposal();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.reporter = reader.string();
                    break;
                case 3:
                    message.deposit = reader.string();
                    break;
                case 4:
                    message.description = reader.string();
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
            creator: isSet(object.creator) ? String(object.creator) : "",
            reporter: isSet(object.reporter) ? String(object.reporter) : "",
            deposit: isSet(object.deposit) ? String(object.deposit) : "",
            description: isSet(object.description) ? String(object.description) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.reporter !== undefined && (obj.reporter = message.reporter);
        message.deposit !== undefined && (obj.deposit = message.deposit);
        message.description !== undefined && (obj.description = message.description);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgSubmitMatchReporterProposal();
        message.creator = object.creator ?? "";
        message.reporter = object.reporter ?? "";
        message.deposit = object.deposit ?? "";
        message.description = object.description ?? "";
        return message;
    },
};
function createBaseMsgSubmitMatchReporterProposalResponse() {
    return {};
}
exports.MsgSubmitMatchReporterProposalResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgSubmitMatchReporterProposalResponse();
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
        const message = createBaseMsgSubmitMatchReporterProposalResponse();
        return message;
    },
};
function createBaseMsgApointMatchReporter() {
    return { creator: "", reporter: "" };
}
exports.MsgApointMatchReporter = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.reporter !== "") {
            writer.uint32(18).string(message.reporter);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgApointMatchReporter();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.reporter = reader.string();
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
            creator: isSet(object.creator) ? String(object.creator) : "",
            reporter: isSet(object.reporter) ? String(object.reporter) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.reporter !== undefined && (obj.reporter = message.reporter);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgApointMatchReporter();
        message.creator = object.creator ?? "";
        message.reporter = object.reporter ?? "";
        return message;
    },
};
function createBaseMsgApointMatchReporterResponse() {
    return {};
}
exports.MsgApointMatchReporterResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgApointMatchReporterResponse();
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
        const message = createBaseMsgApointMatchReporterResponse();
        return message;
    },
};
function createBaseMsgCreateCollection() {
    return { creator: "", name: "", artist: "", storyWriter: "", contributors: [] };
}
exports.MsgCreateCollection = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.name !== "") {
            writer.uint32(18).string(message.name);
        }
        if (message.artist !== "") {
            writer.uint32(26).string(message.artist);
        }
        if (message.storyWriter !== "") {
            writer.uint32(34).string(message.storyWriter);
        }
        for (const v of message.contributors) {
            writer.uint32(42).string(v);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgCreateCollection();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.name = reader.string();
                    break;
                case 3:
                    message.artist = reader.string();
                    break;
                case 4:
                    message.storyWriter = reader.string();
                    break;
                case 5:
                    message.contributors.push(reader.string());
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
            creator: isSet(object.creator) ? String(object.creator) : "",
            name: isSet(object.name) ? String(object.name) : "",
            artist: isSet(object.artist) ? String(object.artist) : "",
            storyWriter: isSet(object.storyWriter) ? String(object.storyWriter) : "",
            contributors: Array.isArray(object?.contributors) ? object.contributors.map((e) => String(e)) : [],
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.name !== undefined && (obj.name = message.name);
        message.artist !== undefined && (obj.artist = message.artist);
        message.storyWriter !== undefined && (obj.storyWriter = message.storyWriter);
        if (message.contributors) {
            obj.contributors = message.contributors.map((e) => e);
        }
        else {
            obj.contributors = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgCreateCollection();
        message.creator = object.creator ?? "";
        message.name = object.name ?? "";
        message.artist = object.artist ?? "";
        message.storyWriter = object.storyWriter ?? "";
        message.contributors = object.contributors?.map((e) => e) || [];
        return message;
    },
};
function createBaseMsgCreateCollectionResponse() {
    return {};
}
exports.MsgCreateCollectionResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgCreateCollectionResponse();
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
        const message = createBaseMsgCreateCollectionResponse();
        return message;
    },
};
function createBaseMsgAddCardToCollection() {
    return { creator: "", collectionId: 0, cardId: 0 };
}
exports.MsgAddCardToCollection = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.collectionId !== 0) {
            writer.uint32(16).uint64(message.collectionId);
        }
        if (message.cardId !== 0) {
            writer.uint32(24).uint64(message.cardId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgAddCardToCollection();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.collectionId = longToNumber(reader.uint64());
                    break;
                case 3:
                    message.cardId = longToNumber(reader.uint64());
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
            creator: isSet(object.creator) ? String(object.creator) : "",
            collectionId: isSet(object.collectionId) ? Number(object.collectionId) : 0,
            cardId: isSet(object.cardId) ? Number(object.cardId) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.collectionId !== undefined && (obj.collectionId = Math.round(message.collectionId));
        message.cardId !== undefined && (obj.cardId = Math.round(message.cardId));
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgAddCardToCollection();
        message.creator = object.creator ?? "";
        message.collectionId = object.collectionId ?? 0;
        message.cardId = object.cardId ?? 0;
        return message;
    },
};
function createBaseMsgAddCardToCollectionResponse() {
    return {};
}
exports.MsgAddCardToCollectionResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgAddCardToCollectionResponse();
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
        const message = createBaseMsgAddCardToCollectionResponse();
        return message;
    },
};
function createBaseMsgFinalizeCollection() {
    return { creator: "", collectionId: 0 };
}
exports.MsgFinalizeCollection = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.collectionId !== 0) {
            writer.uint32(16).uint64(message.collectionId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgFinalizeCollection();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
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
        return {
            creator: isSet(object.creator) ? String(object.creator) : "",
            collectionId: isSet(object.collectionId) ? Number(object.collectionId) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.collectionId !== undefined && (obj.collectionId = Math.round(message.collectionId));
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgFinalizeCollection();
        message.creator = object.creator ?? "";
        message.collectionId = object.collectionId ?? 0;
        return message;
    },
};
function createBaseMsgFinalizeCollectionResponse() {
    return {};
}
exports.MsgFinalizeCollectionResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgFinalizeCollectionResponse();
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
        const message = createBaseMsgFinalizeCollectionResponse();
        return message;
    },
};
function createBaseMsgBuyCollection() {
    return { creator: "", collectionId: 0 };
}
exports.MsgBuyCollection = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.collectionId !== 0) {
            writer.uint32(16).uint64(message.collectionId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgBuyCollection();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
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
        return {
            creator: isSet(object.creator) ? String(object.creator) : "",
            collectionId: isSet(object.collectionId) ? Number(object.collectionId) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.collectionId !== undefined && (obj.collectionId = Math.round(message.collectionId));
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgBuyCollection();
        message.creator = object.creator ?? "";
        message.collectionId = object.collectionId ?? 0;
        return message;
    },
};
function createBaseMsgBuyCollectionResponse() {
    return { airdropClaimed: false };
}
exports.MsgBuyCollectionResponse = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.airdropClaimed === true) {
            writer.uint32(8).bool(message.airdropClaimed);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgBuyCollectionResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.airdropClaimed = reader.bool();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return { airdropClaimed: isSet(object.airdropClaimed) ? Boolean(object.airdropClaimed) : false };
    },
    toJSON(message) {
        const obj = {};
        message.airdropClaimed !== undefined && (obj.airdropClaimed = message.airdropClaimed);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgBuyCollectionResponse();
        message.airdropClaimed = object.airdropClaimed ?? false;
        return message;
    },
};
function createBaseMsgRemoveCardFromCollection() {
    return { creator: "", collectionId: 0, cardId: 0 };
}
exports.MsgRemoveCardFromCollection = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.collectionId !== 0) {
            writer.uint32(16).uint64(message.collectionId);
        }
        if (message.cardId !== 0) {
            writer.uint32(24).uint64(message.cardId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgRemoveCardFromCollection();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.collectionId = longToNumber(reader.uint64());
                    break;
                case 3:
                    message.cardId = longToNumber(reader.uint64());
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
            creator: isSet(object.creator) ? String(object.creator) : "",
            collectionId: isSet(object.collectionId) ? Number(object.collectionId) : 0,
            cardId: isSet(object.cardId) ? Number(object.cardId) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.collectionId !== undefined && (obj.collectionId = Math.round(message.collectionId));
        message.cardId !== undefined && (obj.cardId = Math.round(message.cardId));
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgRemoveCardFromCollection();
        message.creator = object.creator ?? "";
        message.collectionId = object.collectionId ?? 0;
        message.cardId = object.cardId ?? 0;
        return message;
    },
};
function createBaseMsgRemoveCardFromCollectionResponse() {
    return {};
}
exports.MsgRemoveCardFromCollectionResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgRemoveCardFromCollectionResponse();
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
        const message = createBaseMsgRemoveCardFromCollectionResponse();
        return message;
    },
};
function createBaseMsgRemoveContributorFromCollection() {
    return { creator: "", collectionId: 0, user: "" };
}
exports.MsgRemoveContributorFromCollection = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.collectionId !== 0) {
            writer.uint32(16).uint64(message.collectionId);
        }
        if (message.user !== "") {
            writer.uint32(26).string(message.user);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgRemoveContributorFromCollection();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.collectionId = longToNumber(reader.uint64());
                    break;
                case 3:
                    message.user = reader.string();
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
            creator: isSet(object.creator) ? String(object.creator) : "",
            collectionId: isSet(object.collectionId) ? Number(object.collectionId) : 0,
            user: isSet(object.user) ? String(object.user) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.collectionId !== undefined && (obj.collectionId = Math.round(message.collectionId));
        message.user !== undefined && (obj.user = message.user);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgRemoveContributorFromCollection();
        message.creator = object.creator ?? "";
        message.collectionId = object.collectionId ?? 0;
        message.user = object.user ?? "";
        return message;
    },
};
function createBaseMsgRemoveContributorFromCollectionResponse() {
    return {};
}
exports.MsgRemoveContributorFromCollectionResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgRemoveContributorFromCollectionResponse();
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
        const message = createBaseMsgRemoveContributorFromCollectionResponse();
        return message;
    },
};
function createBaseMsgAddContributorToCollection() {
    return { creator: "", collectionId: 0, user: "" };
}
exports.MsgAddContributorToCollection = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.collectionId !== 0) {
            writer.uint32(16).uint64(message.collectionId);
        }
        if (message.user !== "") {
            writer.uint32(26).string(message.user);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgAddContributorToCollection();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.collectionId = longToNumber(reader.uint64());
                    break;
                case 3:
                    message.user = reader.string();
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
            creator: isSet(object.creator) ? String(object.creator) : "",
            collectionId: isSet(object.collectionId) ? Number(object.collectionId) : 0,
            user: isSet(object.user) ? String(object.user) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.collectionId !== undefined && (obj.collectionId = Math.round(message.collectionId));
        message.user !== undefined && (obj.user = message.user);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgAddContributorToCollection();
        message.creator = object.creator ?? "";
        message.collectionId = object.collectionId ?? 0;
        message.user = object.user ?? "";
        return message;
    },
};
function createBaseMsgAddContributorToCollectionResponse() {
    return {};
}
exports.MsgAddContributorToCollectionResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgAddContributorToCollectionResponse();
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
        const message = createBaseMsgAddContributorToCollectionResponse();
        return message;
    },
};
function createBaseMsgSubmitCollectionProposal() {
    return { creator: "", collectionId: 0 };
}
exports.MsgSubmitCollectionProposal = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.collectionId !== 0) {
            writer.uint32(16).uint64(message.collectionId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgSubmitCollectionProposal();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
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
        return {
            creator: isSet(object.creator) ? String(object.creator) : "",
            collectionId: isSet(object.collectionId) ? Number(object.collectionId) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.collectionId !== undefined && (obj.collectionId = Math.round(message.collectionId));
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgSubmitCollectionProposal();
        message.creator = object.creator ?? "";
        message.collectionId = object.collectionId ?? 0;
        return message;
    },
};
function createBaseMsgSubmitCollectionProposalResponse() {
    return {};
}
exports.MsgSubmitCollectionProposalResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgSubmitCollectionProposalResponse();
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
        const message = createBaseMsgSubmitCollectionProposalResponse();
        return message;
    },
};
function createBaseMsgCreateSellOffer() {
    return { creator: "", card: 0, price: "" };
}
exports.MsgCreateSellOffer = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.card !== 0) {
            writer.uint32(16).uint64(message.card);
        }
        if (message.price !== "") {
            writer.uint32(26).string(message.price);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgCreateSellOffer();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.card = longToNumber(reader.uint64());
                    break;
                case 3:
                    message.price = reader.string();
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
            creator: isSet(object.creator) ? String(object.creator) : "",
            card: isSet(object.card) ? Number(object.card) : 0,
            price: isSet(object.price) ? String(object.price) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.card !== undefined && (obj.card = Math.round(message.card));
        message.price !== undefined && (obj.price = message.price);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgCreateSellOffer();
        message.creator = object.creator ?? "";
        message.card = object.card ?? 0;
        message.price = object.price ?? "";
        return message;
    },
};
function createBaseMsgCreateSellOfferResponse() {
    return {};
}
exports.MsgCreateSellOfferResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgCreateSellOfferResponse();
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
        const message = createBaseMsgCreateSellOfferResponse();
        return message;
    },
};
function createBaseMsgBuyCard() {
    return { creator: "", sellOfferId: 0 };
}
exports.MsgBuyCard = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.sellOfferId !== 0) {
            writer.uint32(16).uint64(message.sellOfferId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgBuyCard();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
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
        return {
            creator: isSet(object.creator) ? String(object.creator) : "",
            sellOfferId: isSet(object.sellOfferId) ? Number(object.sellOfferId) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.sellOfferId !== undefined && (obj.sellOfferId = Math.round(message.sellOfferId));
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgBuyCard();
        message.creator = object.creator ?? "";
        message.sellOfferId = object.sellOfferId ?? 0;
        return message;
    },
};
function createBaseMsgBuyCardResponse() {
    return {};
}
exports.MsgBuyCardResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgBuyCardResponse();
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
        const message = createBaseMsgBuyCardResponse();
        return message;
    },
};
function createBaseMsgRemoveSellOffer() {
    return { creator: "", sellOfferId: 0 };
}
exports.MsgRemoveSellOffer = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.sellOfferId !== 0) {
            writer.uint32(16).uint64(message.sellOfferId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgRemoveSellOffer();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
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
        return {
            creator: isSet(object.creator) ? String(object.creator) : "",
            sellOfferId: isSet(object.sellOfferId) ? Number(object.sellOfferId) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.sellOfferId !== undefined && (obj.sellOfferId = Math.round(message.sellOfferId));
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgRemoveSellOffer();
        message.creator = object.creator ?? "";
        message.sellOfferId = object.sellOfferId ?? 0;
        return message;
    },
};
function createBaseMsgRemoveSellOfferResponse() {
    return {};
}
exports.MsgRemoveSellOfferResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgRemoveSellOfferResponse();
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
        const message = createBaseMsgRemoveSellOfferResponse();
        return message;
    },
};
function createBaseMsgAddArtworkToCollection() {
    return { creator: "", collectionId: 0, image: new Uint8Array() };
}
exports.MsgAddArtworkToCollection = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.collectionId !== 0) {
            writer.uint32(16).uint64(message.collectionId);
        }
        if (message.image.length !== 0) {
            writer.uint32(26).bytes(message.image);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgAddArtworkToCollection();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.collectionId = longToNumber(reader.uint64());
                    break;
                case 3:
                    message.image = reader.bytes();
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
            creator: isSet(object.creator) ? String(object.creator) : "",
            collectionId: isSet(object.collectionId) ? Number(object.collectionId) : 0,
            image: isSet(object.image) ? bytesFromBase64(object.image) : new Uint8Array(),
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.collectionId !== undefined && (obj.collectionId = Math.round(message.collectionId));
        message.image !== undefined
            && (obj.image = base64FromBytes(message.image !== undefined ? message.image : new Uint8Array()));
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgAddArtworkToCollection();
        message.creator = object.creator ?? "";
        message.collectionId = object.collectionId ?? 0;
        message.image = object.image ?? new Uint8Array();
        return message;
    },
};
function createBaseMsgAddArtworkToCollectionResponse() {
    return {};
}
exports.MsgAddArtworkToCollectionResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgAddArtworkToCollectionResponse();
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
        const message = createBaseMsgAddArtworkToCollectionResponse();
        return message;
    },
};
function createBaseMsgAddStoryToCollection() {
    return { creator: "", collectionId: 0, story: "" };
}
exports.MsgAddStoryToCollection = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.collectionId !== 0) {
            writer.uint32(16).uint64(message.collectionId);
        }
        if (message.story !== "") {
            writer.uint32(26).string(message.story);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgAddStoryToCollection();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.collectionId = longToNumber(reader.uint64());
                    break;
                case 3:
                    message.story = reader.string();
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
            creator: isSet(object.creator) ? String(object.creator) : "",
            collectionId: isSet(object.collectionId) ? Number(object.collectionId) : 0,
            story: isSet(object.story) ? String(object.story) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.collectionId !== undefined && (obj.collectionId = Math.round(message.collectionId));
        message.story !== undefined && (obj.story = message.story);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgAddStoryToCollection();
        message.creator = object.creator ?? "";
        message.collectionId = object.collectionId ?? 0;
        message.story = object.story ?? "";
        return message;
    },
};
function createBaseMsgAddStoryToCollectionResponse() {
    return {};
}
exports.MsgAddStoryToCollectionResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgAddStoryToCollectionResponse();
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
        const message = createBaseMsgAddStoryToCollectionResponse();
        return message;
    },
};
function createBaseMsgSetCardRarity() {
    return { creator: "", cardId: 0, collectionId: 0, rarity: "" };
}
exports.MsgSetCardRarity = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.cardId !== 0) {
            writer.uint32(16).uint64(message.cardId);
        }
        if (message.collectionId !== 0) {
            writer.uint32(24).uint64(message.collectionId);
        }
        if (message.rarity !== "") {
            writer.uint32(34).string(message.rarity);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgSetCardRarity();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.cardId = longToNumber(reader.uint64());
                    break;
                case 3:
                    message.collectionId = longToNumber(reader.uint64());
                    break;
                case 4:
                    message.rarity = reader.string();
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
            creator: isSet(object.creator) ? String(object.creator) : "",
            cardId: isSet(object.cardId) ? Number(object.cardId) : 0,
            collectionId: isSet(object.collectionId) ? Number(object.collectionId) : 0,
            rarity: isSet(object.rarity) ? String(object.rarity) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.cardId !== undefined && (obj.cardId = Math.round(message.cardId));
        message.collectionId !== undefined && (obj.collectionId = Math.round(message.collectionId));
        message.rarity !== undefined && (obj.rarity = message.rarity);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgSetCardRarity();
        message.creator = object.creator ?? "";
        message.cardId = object.cardId ?? 0;
        message.collectionId = object.collectionId ?? 0;
        message.rarity = object.rarity ?? "";
        return message;
    },
};
function createBaseMsgSetCardRarityResponse() {
    return {};
}
exports.MsgSetCardRarityResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgSetCardRarityResponse();
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
        const message = createBaseMsgSetCardRarityResponse();
        return message;
    },
};
function createBaseMsgCreateCouncil() {
    return { creator: "", cardId: 0 };
}
exports.MsgCreateCouncil = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.cardId !== 0) {
            writer.uint32(16).uint64(message.cardId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgCreateCouncil();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.cardId = longToNumber(reader.uint64());
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
            creator: isSet(object.creator) ? String(object.creator) : "",
            cardId: isSet(object.cardId) ? Number(object.cardId) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.cardId !== undefined && (obj.cardId = Math.round(message.cardId));
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgCreateCouncil();
        message.creator = object.creator ?? "";
        message.cardId = object.cardId ?? 0;
        return message;
    },
};
function createBaseMsgCreateCouncilResponse() {
    return {};
}
exports.MsgCreateCouncilResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgCreateCouncilResponse();
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
        const message = createBaseMsgCreateCouncilResponse();
        return message;
    },
};
function createBaseMsgCommitCouncilResponse() {
    return { creator: "", response: "", councilId: 0, suggestion: "" };
}
exports.MsgCommitCouncilResponse = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.response !== "") {
            writer.uint32(18).string(message.response);
        }
        if (message.councilId !== 0) {
            writer.uint32(24).uint64(message.councilId);
        }
        if (message.suggestion !== "") {
            writer.uint32(34).string(message.suggestion);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgCommitCouncilResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.response = reader.string();
                    break;
                case 3:
                    message.councilId = longToNumber(reader.uint64());
                    break;
                case 4:
                    message.suggestion = reader.string();
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
            creator: isSet(object.creator) ? String(object.creator) : "",
            response: isSet(object.response) ? String(object.response) : "",
            councilId: isSet(object.councilId) ? Number(object.councilId) : 0,
            suggestion: isSet(object.suggestion) ? String(object.suggestion) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.response !== undefined && (obj.response = message.response);
        message.councilId !== undefined && (obj.councilId = Math.round(message.councilId));
        message.suggestion !== undefined && (obj.suggestion = message.suggestion);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgCommitCouncilResponse();
        message.creator = object.creator ?? "";
        message.response = object.response ?? "";
        message.councilId = object.councilId ?? 0;
        message.suggestion = object.suggestion ?? "";
        return message;
    },
};
function createBaseMsgCommitCouncilResponseResponse() {
    return {};
}
exports.MsgCommitCouncilResponseResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgCommitCouncilResponseResponse();
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
        const message = createBaseMsgCommitCouncilResponseResponse();
        return message;
    },
};
function createBaseMsgRevealCouncilResponse() {
    return { creator: "", response: 0, secret: "", councilId: 0 };
}
exports.MsgRevealCouncilResponse = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.response !== 0) {
            writer.uint32(16).int32(message.response);
        }
        if (message.secret !== "") {
            writer.uint32(26).string(message.secret);
        }
        if (message.councilId !== 0) {
            writer.uint32(32).uint64(message.councilId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgRevealCouncilResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.response = reader.int32();
                    break;
                case 3:
                    message.secret = reader.string();
                    break;
                case 4:
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
        return {
            creator: isSet(object.creator) ? String(object.creator) : "",
            response: isSet(object.response) ? (0, council_1.responseFromJSON)(object.response) : 0,
            secret: isSet(object.secret) ? String(object.secret) : "",
            councilId: isSet(object.councilId) ? Number(object.councilId) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.response !== undefined && (obj.response = (0, council_1.responseToJSON)(message.response));
        message.secret !== undefined && (obj.secret = message.secret);
        message.councilId !== undefined && (obj.councilId = Math.round(message.councilId));
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgRevealCouncilResponse();
        message.creator = object.creator ?? "";
        message.response = object.response ?? 0;
        message.secret = object.secret ?? "";
        message.councilId = object.councilId ?? 0;
        return message;
    },
};
function createBaseMsgRevealCouncilResponseResponse() {
    return {};
}
exports.MsgRevealCouncilResponseResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgRevealCouncilResponseResponse();
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
        const message = createBaseMsgRevealCouncilResponseResponse();
        return message;
    },
};
function createBaseMsgRestartCouncil() {
    return { creator: "", councilId: 0 };
}
exports.MsgRestartCouncil = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.councilId !== 0) {
            writer.uint32(16).uint64(message.councilId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgRestartCouncil();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
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
        return {
            creator: isSet(object.creator) ? String(object.creator) : "",
            councilId: isSet(object.councilId) ? Number(object.councilId) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.councilId !== undefined && (obj.councilId = Math.round(message.councilId));
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgRestartCouncil();
        message.creator = object.creator ?? "";
        message.councilId = object.councilId ?? 0;
        return message;
    },
};
function createBaseMsgRestartCouncilResponse() {
    return {};
}
exports.MsgRestartCouncilResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgRestartCouncilResponse();
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
        const message = createBaseMsgRestartCouncilResponse();
        return message;
    },
};
function createBaseMsgRewokeCouncilRegistration() {
    return { creator: "" };
}
exports.MsgRewokeCouncilRegistration = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgRewokeCouncilRegistration();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return { creator: isSet(object.creator) ? String(object.creator) : "" };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgRewokeCouncilRegistration();
        message.creator = object.creator ?? "";
        return message;
    },
};
function createBaseMsgRewokeCouncilRegistrationResponse() {
    return {};
}
exports.MsgRewokeCouncilRegistrationResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgRewokeCouncilRegistrationResponse();
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
        const message = createBaseMsgRewokeCouncilRegistrationResponse();
        return message;
    },
};
function createBaseMsgConfirmMatch() {
    return { creator: "", matchId: 0, outcome: 0 };
}
exports.MsgConfirmMatch = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.matchId !== 0) {
            writer.uint32(16).uint64(message.matchId);
        }
        if (message.outcome !== 0) {
            writer.uint32(24).int32(message.outcome);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgConfirmMatch();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.matchId = longToNumber(reader.uint64());
                    break;
                case 3:
                    message.outcome = reader.int32();
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
            creator: isSet(object.creator) ? String(object.creator) : "",
            matchId: isSet(object.matchId) ? Number(object.matchId) : 0,
            outcome: isSet(object.outcome) ? outcomeFromJSON(object.outcome) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.matchId !== undefined && (obj.matchId = Math.round(message.matchId));
        message.outcome !== undefined && (obj.outcome = outcomeToJSON(message.outcome));
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgConfirmMatch();
        message.creator = object.creator ?? "";
        message.matchId = object.matchId ?? 0;
        message.outcome = object.outcome ?? 0;
        return message;
    },
};
function createBaseMsgConfirmMatchResponse() {
    return {};
}
exports.MsgConfirmMatchResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgConfirmMatchResponse();
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
        const message = createBaseMsgConfirmMatchResponse();
        return message;
    },
};
function createBaseMsgSetProfileCard() {
    return { creator: "", cardId: 0 };
}
exports.MsgSetProfileCard = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.cardId !== 0) {
            writer.uint32(16).uint64(message.cardId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgSetProfileCard();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.cardId = longToNumber(reader.uint64());
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
            creator: isSet(object.creator) ? String(object.creator) : "",
            cardId: isSet(object.cardId) ? Number(object.cardId) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.cardId !== undefined && (obj.cardId = Math.round(message.cardId));
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgSetProfileCard();
        message.creator = object.creator ?? "";
        message.cardId = object.cardId ?? 0;
        return message;
    },
};
function createBaseMsgSetProfileCardResponse() {
    return {};
}
exports.MsgSetProfileCardResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgSetProfileCardResponse();
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
        const message = createBaseMsgSetProfileCardResponse();
        return message;
    },
};
function createBaseMsgOpenBoosterPack() {
    return { creator: "", boosterPackId: 0 };
}
exports.MsgOpenBoosterPack = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.boosterPackId !== 0) {
            writer.uint32(16).uint64(message.boosterPackId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgOpenBoosterPack();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.boosterPackId = longToNumber(reader.uint64());
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
            creator: isSet(object.creator) ? String(object.creator) : "",
            boosterPackId: isSet(object.boosterPackId) ? Number(object.boosterPackId) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.boosterPackId !== undefined && (obj.boosterPackId = Math.round(message.boosterPackId));
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgOpenBoosterPack();
        message.creator = object.creator ?? "";
        message.boosterPackId = object.boosterPackId ?? 0;
        return message;
    },
};
function createBaseMsgOpenBoosterPackResponse() {
    return {};
}
exports.MsgOpenBoosterPackResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgOpenBoosterPackResponse();
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
        const message = createBaseMsgOpenBoosterPackResponse();
        return message;
    },
};
function createBaseMsgTransferBoosterPack() {
    return { creator: "", boosterPackId: 0, receiver: "" };
}
exports.MsgTransferBoosterPack = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.boosterPackId !== 0) {
            writer.uint32(16).uint64(message.boosterPackId);
        }
        if (message.receiver !== "") {
            writer.uint32(26).string(message.receiver);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgTransferBoosterPack();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.boosterPackId = longToNumber(reader.uint64());
                    break;
                case 3:
                    message.receiver = reader.string();
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
            creator: isSet(object.creator) ? String(object.creator) : "",
            boosterPackId: isSet(object.boosterPackId) ? Number(object.boosterPackId) : 0,
            receiver: isSet(object.receiver) ? String(object.receiver) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.boosterPackId !== undefined && (obj.boosterPackId = Math.round(message.boosterPackId));
        message.receiver !== undefined && (obj.receiver = message.receiver);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgTransferBoosterPack();
        message.creator = object.creator ?? "";
        message.boosterPackId = object.boosterPackId ?? 0;
        message.receiver = object.receiver ?? "";
        return message;
    },
};
function createBaseMsgTransferBoosterPackResponse() {
    return {};
}
exports.MsgTransferBoosterPackResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgTransferBoosterPackResponse();
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
        const message = createBaseMsgTransferBoosterPackResponse();
        return message;
    },
};
function createBaseMsgSetCollectionStoryWriter() {
    return { creator: "", collectionId: 0, storyWriter: "" };
}
exports.MsgSetCollectionStoryWriter = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.collectionId !== 0) {
            writer.uint32(16).uint64(message.collectionId);
        }
        if (message.storyWriter !== "") {
            writer.uint32(26).string(message.storyWriter);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgSetCollectionStoryWriter();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.collectionId = longToNumber(reader.uint64());
                    break;
                case 3:
                    message.storyWriter = reader.string();
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
            creator: isSet(object.creator) ? String(object.creator) : "",
            collectionId: isSet(object.collectionId) ? Number(object.collectionId) : 0,
            storyWriter: isSet(object.storyWriter) ? String(object.storyWriter) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.collectionId !== undefined && (obj.collectionId = Math.round(message.collectionId));
        message.storyWriter !== undefined && (obj.storyWriter = message.storyWriter);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgSetCollectionStoryWriter();
        message.creator = object.creator ?? "";
        message.collectionId = object.collectionId ?? 0;
        message.storyWriter = object.storyWriter ?? "";
        return message;
    },
};
function createBaseMsgSetCollectionStoryWriterResponse() {
    return {};
}
exports.MsgSetCollectionStoryWriterResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgSetCollectionStoryWriterResponse();
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
        const message = createBaseMsgSetCollectionStoryWriterResponse();
        return message;
    },
};
function createBaseMsgSetCollectionArtist() {
    return { creator: "", collectionId: 0, artist: "" };
}
exports.MsgSetCollectionArtist = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.collectionId !== 0) {
            writer.uint32(16).uint64(message.collectionId);
        }
        if (message.artist !== "") {
            writer.uint32(26).string(message.artist);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgSetCollectionArtist();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.collectionId = longToNumber(reader.uint64());
                    break;
                case 3:
                    message.artist = reader.string();
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
            creator: isSet(object.creator) ? String(object.creator) : "",
            collectionId: isSet(object.collectionId) ? Number(object.collectionId) : 0,
            artist: isSet(object.artist) ? String(object.artist) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.collectionId !== undefined && (obj.collectionId = Math.round(message.collectionId));
        message.artist !== undefined && (obj.artist = message.artist);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgSetCollectionArtist();
        message.creator = object.creator ?? "";
        message.collectionId = object.collectionId ?? 0;
        message.artist = object.artist ?? "";
        return message;
    },
};
function createBaseMsgSetCollectionArtistResponse() {
    return {};
}
exports.MsgSetCollectionArtistResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgSetCollectionArtistResponse();
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
        const message = createBaseMsgSetCollectionArtistResponse();
        return message;
    },
};
function createBaseMsgSetUserWebsite() {
    return { creator: "", website: "" };
}
exports.MsgSetUserWebsite = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.website !== "") {
            writer.uint32(18).string(message.website);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgSetUserWebsite();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.website = reader.string();
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
            creator: isSet(object.creator) ? String(object.creator) : "",
            website: isSet(object.website) ? String(object.website) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.website !== undefined && (obj.website = message.website);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgSetUserWebsite();
        message.creator = object.creator ?? "";
        message.website = object.website ?? "";
        return message;
    },
};
function createBaseMsgSetUserWebsiteResponse() {
    return {};
}
exports.MsgSetUserWebsiteResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgSetUserWebsiteResponse();
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
        const message = createBaseMsgSetUserWebsiteResponse();
        return message;
    },
};
function createBaseMsgSetUserBiography() {
    return { creator: "", biography: "" };
}
exports.MsgSetUserBiography = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.biography !== "") {
            writer.uint32(18).string(message.biography);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgSetUserBiography();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.biography = reader.string();
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
            creator: isSet(object.creator) ? String(object.creator) : "",
            biography: isSet(object.biography) ? String(object.biography) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.biography !== undefined && (obj.biography = message.biography);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgSetUserBiography();
        message.creator = object.creator ?? "";
        message.biography = object.biography ?? "";
        return message;
    },
};
function createBaseMsgSetUserBiographyResponse() {
    return {};
}
exports.MsgSetUserBiographyResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgSetUserBiographyResponse();
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
        const message = createBaseMsgSetUserBiographyResponse();
        return message;
    },
};
class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
        this.Createuser = this.Createuser.bind(this);
        this.BuyCardScheme = this.BuyCardScheme.bind(this);
        this.VoteCard = this.VoteCard.bind(this);
        this.SaveCardContent = this.SaveCardContent.bind(this);
        this.TransferCard = this.TransferCard.bind(this);
        this.DonateToCard = this.DonateToCard.bind(this);
        this.AddArtwork = this.AddArtwork.bind(this);
        this.SubmitCopyrightProposal = this.SubmitCopyrightProposal.bind(this);
        this.ChangeArtist = this.ChangeArtist.bind(this);
        this.RegisterForCouncil = this.RegisterForCouncil.bind(this);
        this.ReportMatch = this.ReportMatch.bind(this);
        this.SubmitMatchReporterProposal = this.SubmitMatchReporterProposal.bind(this);
        this.ApointMatchReporter = this.ApointMatchReporter.bind(this);
        this.CreateCollection = this.CreateCollection.bind(this);
        this.AddCardToCollection = this.AddCardToCollection.bind(this);
        this.FinalizeCollection = this.FinalizeCollection.bind(this);
        this.BuyCollection = this.BuyCollection.bind(this);
        this.RemoveCardFromCollection = this.RemoveCardFromCollection.bind(this);
        this.RemoveContributorFromCollection = this.RemoveContributorFromCollection.bind(this);
        this.AddContributorToCollection = this.AddContributorToCollection.bind(this);
        this.SubmitCollectionProposal = this.SubmitCollectionProposal.bind(this);
        this.CreateSellOffer = this.CreateSellOffer.bind(this);
        this.BuyCard = this.BuyCard.bind(this);
        this.RemoveSellOffer = this.RemoveSellOffer.bind(this);
        this.AddArtworkToCollection = this.AddArtworkToCollection.bind(this);
        this.AddStoryToCollection = this.AddStoryToCollection.bind(this);
        this.SetCardRarity = this.SetCardRarity.bind(this);
        this.CreateCouncil = this.CreateCouncil.bind(this);
        this.CommitCouncilResponse = this.CommitCouncilResponse.bind(this);
        this.RevealCouncilResponse = this.RevealCouncilResponse.bind(this);
        this.RestartCouncil = this.RestartCouncil.bind(this);
        this.RewokeCouncilRegistration = this.RewokeCouncilRegistration.bind(this);
        this.ConfirmMatch = this.ConfirmMatch.bind(this);
        this.SetProfileCard = this.SetProfileCard.bind(this);
        this.OpenBoosterPack = this.OpenBoosterPack.bind(this);
        this.TransferBoosterPack = this.TransferBoosterPack.bind(this);
        this.SetCollectionStoryWriter = this.SetCollectionStoryWriter.bind(this);
        this.SetCollectionArtist = this.SetCollectionArtist.bind(this);
        this.SetUserWebsite = this.SetUserWebsite.bind(this);
        this.SetUserBiography = this.SetUserBiography.bind(this);
    }
    Createuser(request) {
        const data = exports.MsgCreateuser.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "Createuser", data);
        return promise.then((data) => exports.MsgCreateuserResponse.decode(new minimal_1.default.Reader(data)));
    }
    BuyCardScheme(request) {
        const data = exports.MsgBuyCardScheme.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "BuyCardScheme", data);
        return promise.then((data) => exports.MsgBuyCardSchemeResponse.decode(new minimal_1.default.Reader(data)));
    }
    VoteCard(request) {
        const data = exports.MsgVoteCard.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "VoteCard", data);
        return promise.then((data) => exports.MsgVoteCardResponse.decode(new minimal_1.default.Reader(data)));
    }
    SaveCardContent(request) {
        const data = exports.MsgSaveCardContent.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "SaveCardContent", data);
        return promise.then((data) => exports.MsgSaveCardContentResponse.decode(new minimal_1.default.Reader(data)));
    }
    TransferCard(request) {
        const data = exports.MsgTransferCard.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "TransferCard", data);
        return promise.then((data) => exports.MsgTransferCardResponse.decode(new minimal_1.default.Reader(data)));
    }
    DonateToCard(request) {
        const data = exports.MsgDonateToCard.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "DonateToCard", data);
        return promise.then((data) => exports.MsgDonateToCardResponse.decode(new minimal_1.default.Reader(data)));
    }
    AddArtwork(request) {
        const data = exports.MsgAddArtwork.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "AddArtwork", data);
        return promise.then((data) => exports.MsgAddArtworkResponse.decode(new minimal_1.default.Reader(data)));
    }
    SubmitCopyrightProposal(request) {
        const data = exports.MsgSubmitCopyrightProposal.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "SubmitCopyrightProposal", data);
        return promise.then((data) => exports.MsgSubmitCopyrightProposalResponse.decode(new minimal_1.default.Reader(data)));
    }
    ChangeArtist(request) {
        const data = exports.MsgChangeArtist.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "ChangeArtist", data);
        return promise.then((data) => exports.MsgChangeArtistResponse.decode(new minimal_1.default.Reader(data)));
    }
    RegisterForCouncil(request) {
        const data = exports.MsgRegisterForCouncil.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "RegisterForCouncil", data);
        return promise.then((data) => exports.MsgRegisterForCouncilResponse.decode(new minimal_1.default.Reader(data)));
    }
    ReportMatch(request) {
        const data = exports.MsgReportMatch.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "ReportMatch", data);
        return promise.then((data) => exports.MsgReportMatchResponse.decode(new minimal_1.default.Reader(data)));
    }
    SubmitMatchReporterProposal(request) {
        const data = exports.MsgSubmitMatchReporterProposal.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "SubmitMatchReporterProposal", data);
        return promise.then((data) => exports.MsgSubmitMatchReporterProposalResponse.decode(new minimal_1.default.Reader(data)));
    }
    ApointMatchReporter(request) {
        const data = exports.MsgApointMatchReporter.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "ApointMatchReporter", data);
        return promise.then((data) => exports.MsgApointMatchReporterResponse.decode(new minimal_1.default.Reader(data)));
    }
    CreateCollection(request) {
        const data = exports.MsgCreateCollection.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "CreateCollection", data);
        return promise.then((data) => exports.MsgCreateCollectionResponse.decode(new minimal_1.default.Reader(data)));
    }
    AddCardToCollection(request) {
        const data = exports.MsgAddCardToCollection.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "AddCardToCollection", data);
        return promise.then((data) => exports.MsgAddCardToCollectionResponse.decode(new minimal_1.default.Reader(data)));
    }
    FinalizeCollection(request) {
        const data = exports.MsgFinalizeCollection.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "FinalizeCollection", data);
        return promise.then((data) => exports.MsgFinalizeCollectionResponse.decode(new minimal_1.default.Reader(data)));
    }
    BuyCollection(request) {
        const data = exports.MsgBuyCollection.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "BuyCollection", data);
        return promise.then((data) => exports.MsgBuyCollectionResponse.decode(new minimal_1.default.Reader(data)));
    }
    RemoveCardFromCollection(request) {
        const data = exports.MsgRemoveCardFromCollection.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "RemoveCardFromCollection", data);
        return promise.then((data) => exports.MsgRemoveCardFromCollectionResponse.decode(new minimal_1.default.Reader(data)));
    }
    RemoveContributorFromCollection(request) {
        const data = exports.MsgRemoveContributorFromCollection.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "RemoveContributorFromCollection", data);
        return promise.then((data) => exports.MsgRemoveContributorFromCollectionResponse.decode(new minimal_1.default.Reader(data)));
    }
    AddContributorToCollection(request) {
        const data = exports.MsgAddContributorToCollection.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "AddContributorToCollection", data);
        return promise.then((data) => exports.MsgAddContributorToCollectionResponse.decode(new minimal_1.default.Reader(data)));
    }
    SubmitCollectionProposal(request) {
        const data = exports.MsgSubmitCollectionProposal.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "SubmitCollectionProposal", data);
        return promise.then((data) => exports.MsgSubmitCollectionProposalResponse.decode(new minimal_1.default.Reader(data)));
    }
    CreateSellOffer(request) {
        const data = exports.MsgCreateSellOffer.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "CreateSellOffer", data);
        return promise.then((data) => exports.MsgCreateSellOfferResponse.decode(new minimal_1.default.Reader(data)));
    }
    BuyCard(request) {
        const data = exports.MsgBuyCard.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "BuyCard", data);
        return promise.then((data) => exports.MsgBuyCardResponse.decode(new minimal_1.default.Reader(data)));
    }
    RemoveSellOffer(request) {
        const data = exports.MsgRemoveSellOffer.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "RemoveSellOffer", data);
        return promise.then((data) => exports.MsgRemoveSellOfferResponse.decode(new minimal_1.default.Reader(data)));
    }
    AddArtworkToCollection(request) {
        const data = exports.MsgAddArtworkToCollection.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "AddArtworkToCollection", data);
        return promise.then((data) => exports.MsgAddArtworkToCollectionResponse.decode(new minimal_1.default.Reader(data)));
    }
    AddStoryToCollection(request) {
        const data = exports.MsgAddStoryToCollection.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "AddStoryToCollection", data);
        return promise.then((data) => exports.MsgAddStoryToCollectionResponse.decode(new minimal_1.default.Reader(data)));
    }
    SetCardRarity(request) {
        const data = exports.MsgSetCardRarity.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "SetCardRarity", data);
        return promise.then((data) => exports.MsgSetCardRarityResponse.decode(new minimal_1.default.Reader(data)));
    }
    CreateCouncil(request) {
        const data = exports.MsgCreateCouncil.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "CreateCouncil", data);
        return promise.then((data) => exports.MsgCreateCouncilResponse.decode(new minimal_1.default.Reader(data)));
    }
    CommitCouncilResponse(request) {
        const data = exports.MsgCommitCouncilResponse.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "CommitCouncilResponse", data);
        return promise.then((data) => exports.MsgCommitCouncilResponseResponse.decode(new minimal_1.default.Reader(data)));
    }
    RevealCouncilResponse(request) {
        const data = exports.MsgRevealCouncilResponse.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "RevealCouncilResponse", data);
        return promise.then((data) => exports.MsgRevealCouncilResponseResponse.decode(new minimal_1.default.Reader(data)));
    }
    RestartCouncil(request) {
        const data = exports.MsgRestartCouncil.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "RestartCouncil", data);
        return promise.then((data) => exports.MsgRestartCouncilResponse.decode(new minimal_1.default.Reader(data)));
    }
    RewokeCouncilRegistration(request) {
        const data = exports.MsgRewokeCouncilRegistration.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "RewokeCouncilRegistration", data);
        return promise.then((data) => exports.MsgRewokeCouncilRegistrationResponse.decode(new minimal_1.default.Reader(data)));
    }
    ConfirmMatch(request) {
        const data = exports.MsgConfirmMatch.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "ConfirmMatch", data);
        return promise.then((data) => exports.MsgConfirmMatchResponse.decode(new minimal_1.default.Reader(data)));
    }
    SetProfileCard(request) {
        const data = exports.MsgSetProfileCard.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "SetProfileCard", data);
        return promise.then((data) => exports.MsgSetProfileCardResponse.decode(new minimal_1.default.Reader(data)));
    }
    OpenBoosterPack(request) {
        const data = exports.MsgOpenBoosterPack.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "OpenBoosterPack", data);
        return promise.then((data) => exports.MsgOpenBoosterPackResponse.decode(new minimal_1.default.Reader(data)));
    }
    TransferBoosterPack(request) {
        const data = exports.MsgTransferBoosterPack.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "TransferBoosterPack", data);
        return promise.then((data) => exports.MsgTransferBoosterPackResponse.decode(new minimal_1.default.Reader(data)));
    }
    SetCollectionStoryWriter(request) {
        const data = exports.MsgSetCollectionStoryWriter.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "SetCollectionStoryWriter", data);
        return promise.then((data) => exports.MsgSetCollectionStoryWriterResponse.decode(new minimal_1.default.Reader(data)));
    }
    SetCollectionArtist(request) {
        const data = exports.MsgSetCollectionArtist.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "SetCollectionArtist", data);
        return promise.then((data) => exports.MsgSetCollectionArtistResponse.decode(new minimal_1.default.Reader(data)));
    }
    SetUserWebsite(request) {
        const data = exports.MsgSetUserWebsite.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "SetUserWebsite", data);
        return promise.then((data) => exports.MsgSetUserWebsiteResponse.decode(new minimal_1.default.Reader(data)));
    }
    SetUserBiography(request) {
        const data = exports.MsgSetUserBiography.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Msg", "SetUserBiography", data);
        return promise.then((data) => exports.MsgSetUserBiographyResponse.decode(new minimal_1.default.Reader(data)));
    }
}
exports.MsgClientImpl = MsgClientImpl;
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
function bytesFromBase64(b64) {
    if (globalThis.Buffer) {
        return Uint8Array.from(globalThis.Buffer.from(b64, "base64"));
    }
    else {
        const bin = globalThis.atob(b64);
        const arr = new Uint8Array(bin.length);
        for (let i = 0; i < bin.length; ++i) {
            arr[i] = bin.charCodeAt(i);
        }
        return arr;
    }
}
function base64FromBytes(arr) {
    if (globalThis.Buffer) {
        return globalThis.Buffer.from(arr).toString("base64");
    }
    else {
        const bin = [];
        arr.forEach((byte) => {
            bin.push(String.fromCharCode(byte));
        });
        return globalThis.btoa(bin.join(""));
    }
}
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
