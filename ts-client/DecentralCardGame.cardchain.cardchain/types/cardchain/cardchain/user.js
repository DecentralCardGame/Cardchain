"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.AirDrops = exports.BoosterPack = exports.CouncilParticipation = exports.User = exports.airDropToJSON = exports.airDropFromJSON = exports.AirDrop = exports.councilStatusToJSON = exports.councilStatusFromJSON = exports.CouncilStatus = exports.protobufPackage = void 0;
/* eslint-disable */
const long_1 = require("long");
const minimal_1 = require("protobufjs/minimal");
const vote_right_1 = require("./vote_right");
exports.protobufPackage = "DecentralCardGame.cardchain.cardchain";
var CouncilStatus;
(function (CouncilStatus) {
    CouncilStatus[CouncilStatus["available"] = 0] = "available";
    CouncilStatus[CouncilStatus["unavailable"] = 1] = "unavailable";
    CouncilStatus[CouncilStatus["openCouncil"] = 2] = "openCouncil";
    CouncilStatus[CouncilStatus["startedCouncil"] = 3] = "startedCouncil";
    CouncilStatus[CouncilStatus["UNRECOGNIZED"] = -1] = "UNRECOGNIZED";
})(CouncilStatus = exports.CouncilStatus || (exports.CouncilStatus = {}));
function councilStatusFromJSON(object) {
    switch (object) {
        case 0:
        case "available":
            return CouncilStatus.available;
        case 1:
        case "unavailable":
            return CouncilStatus.unavailable;
        case 2:
        case "openCouncil":
            return CouncilStatus.openCouncil;
        case 3:
        case "startedCouncil":
            return CouncilStatus.startedCouncil;
        case -1:
        case "UNRECOGNIZED":
        default:
            return CouncilStatus.UNRECOGNIZED;
    }
}
exports.councilStatusFromJSON = councilStatusFromJSON;
function councilStatusToJSON(object) {
    switch (object) {
        case CouncilStatus.available:
            return "available";
        case CouncilStatus.unavailable:
            return "unavailable";
        case CouncilStatus.openCouncil:
            return "openCouncil";
        case CouncilStatus.startedCouncil:
            return "startedCouncil";
        case CouncilStatus.UNRECOGNIZED:
        default:
            return "UNRECOGNIZED";
    }
}
exports.councilStatusToJSON = councilStatusToJSON;
var AirDrop;
(function (AirDrop) {
    AirDrop[AirDrop["play"] = 0] = "play";
    AirDrop[AirDrop["vote"] = 1] = "vote";
    AirDrop[AirDrop["create"] = 2] = "create";
    AirDrop[AirDrop["buy"] = 3] = "buy";
    AirDrop[AirDrop["user"] = 4] = "user";
    AirDrop[AirDrop["UNRECOGNIZED"] = -1] = "UNRECOGNIZED";
})(AirDrop = exports.AirDrop || (exports.AirDrop = {}));
function airDropFromJSON(object) {
    switch (object) {
        case 0:
        case "play":
            return AirDrop.play;
        case 1:
        case "vote":
            return AirDrop.vote;
        case 2:
        case "create":
            return AirDrop.create;
        case 3:
        case "buy":
            return AirDrop.buy;
        case 4:
        case "user":
            return AirDrop.user;
        case -1:
        case "UNRECOGNIZED":
        default:
            return AirDrop.UNRECOGNIZED;
    }
}
exports.airDropFromJSON = airDropFromJSON;
function airDropToJSON(object) {
    switch (object) {
        case AirDrop.play:
            return "play";
        case AirDrop.vote:
            return "vote";
        case AirDrop.create:
            return "create";
        case AirDrop.buy:
            return "buy";
        case AirDrop.user:
            return "user";
        case AirDrop.UNRECOGNIZED:
        default:
            return "UNRECOGNIZED";
    }
}
exports.airDropToJSON = airDropToJSON;
function createBaseUser() {
    return {
        alias: "",
        ownedCardSchemes: [],
        ownedPrototypes: [],
        cards: [],
        voteRights: [],
        councilParticipation: undefined,
        ReportMatches: false,
        profileCard: 0,
        airDrops: undefined,
        boosterPacks: [],
        website: "",
        biography: "",
    };
}
exports.User = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.alias !== "") {
            writer.uint32(10).string(message.alias);
        }
        writer.uint32(18).fork();
        for (const v of message.ownedCardSchemes) {
            writer.uint64(v);
        }
        writer.ldelim();
        writer.uint32(26).fork();
        for (const v of message.ownedPrototypes) {
            writer.uint64(v);
        }
        writer.ldelim();
        writer.uint32(34).fork();
        for (const v of message.cards) {
            writer.uint64(v);
        }
        writer.ldelim();
        for (const v of message.voteRights) {
            vote_right_1.VoteRight.encode(v, writer.uint32(42).fork()).ldelim();
        }
        if (message.councilParticipation !== undefined) {
            exports.CouncilParticipation.encode(message.councilParticipation, writer.uint32(50).fork()).ldelim();
        }
        if (message.ReportMatches === true) {
            writer.uint32(56).bool(message.ReportMatches);
        }
        if (message.profileCard !== 0) {
            writer.uint32(64).uint64(message.profileCard);
        }
        if (message.airDrops !== undefined) {
            exports.AirDrops.encode(message.airDrops, writer.uint32(74).fork()).ldelim();
        }
        for (const v of message.boosterPacks) {
            exports.BoosterPack.encode(v, writer.uint32(82).fork()).ldelim();
        }
        if (message.website !== "") {
            writer.uint32(90).string(message.website);
        }
        if (message.biography !== "") {
            writer.uint32(98).string(message.biography);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseUser();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.alias = reader.string();
                    break;
                case 2:
                    if ((tag & 7) === 2) {
                        const end2 = reader.uint32() + reader.pos;
                        while (reader.pos < end2) {
                            message.ownedCardSchemes.push(longToNumber(reader.uint64()));
                        }
                    }
                    else {
                        message.ownedCardSchemes.push(longToNumber(reader.uint64()));
                    }
                    break;
                case 3:
                    if ((tag & 7) === 2) {
                        const end2 = reader.uint32() + reader.pos;
                        while (reader.pos < end2) {
                            message.ownedPrototypes.push(longToNumber(reader.uint64()));
                        }
                    }
                    else {
                        message.ownedPrototypes.push(longToNumber(reader.uint64()));
                    }
                    break;
                case 4:
                    if ((tag & 7) === 2) {
                        const end2 = reader.uint32() + reader.pos;
                        while (reader.pos < end2) {
                            message.cards.push(longToNumber(reader.uint64()));
                        }
                    }
                    else {
                        message.cards.push(longToNumber(reader.uint64()));
                    }
                    break;
                case 5:
                    message.voteRights.push(vote_right_1.VoteRight.decode(reader, reader.uint32()));
                    break;
                case 6:
                    message.councilParticipation = exports.CouncilParticipation.decode(reader, reader.uint32());
                    break;
                case 7:
                    message.ReportMatches = reader.bool();
                    break;
                case 8:
                    message.profileCard = longToNumber(reader.uint64());
                    break;
                case 9:
                    message.airDrops = exports.AirDrops.decode(reader, reader.uint32());
                    break;
                case 10:
                    message.boosterPacks.push(exports.BoosterPack.decode(reader, reader.uint32()));
                    break;
                case 11:
                    message.website = reader.string();
                    break;
                case 12:
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
            alias: isSet(object.alias) ? String(object.alias) : "",
            ownedCardSchemes: Array.isArray(object?.ownedCardSchemes)
                ? object.ownedCardSchemes.map((e) => Number(e))
                : [],
            ownedPrototypes: Array.isArray(object?.ownedPrototypes) ? object.ownedPrototypes.map((e) => Number(e)) : [],
            cards: Array.isArray(object?.cards) ? object.cards.map((e) => Number(e)) : [],
            voteRights: Array.isArray(object?.voteRights) ? object.voteRights.map((e) => vote_right_1.VoteRight.fromJSON(e)) : [],
            councilParticipation: isSet(object.councilParticipation)
                ? exports.CouncilParticipation.fromJSON(object.councilParticipation)
                : undefined,
            ReportMatches: isSet(object.ReportMatches) ? Boolean(object.ReportMatches) : false,
            profileCard: isSet(object.profileCard) ? Number(object.profileCard) : 0,
            airDrops: isSet(object.airDrops) ? exports.AirDrops.fromJSON(object.airDrops) : undefined,
            boosterPacks: Array.isArray(object?.boosterPacks)
                ? object.boosterPacks.map((e) => exports.BoosterPack.fromJSON(e))
                : [],
            website: isSet(object.website) ? String(object.website) : "",
            biography: isSet(object.biography) ? String(object.biography) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        message.alias !== undefined && (obj.alias = message.alias);
        if (message.ownedCardSchemes) {
            obj.ownedCardSchemes = message.ownedCardSchemes.map((e) => Math.round(e));
        }
        else {
            obj.ownedCardSchemes = [];
        }
        if (message.ownedPrototypes) {
            obj.ownedPrototypes = message.ownedPrototypes.map((e) => Math.round(e));
        }
        else {
            obj.ownedPrototypes = [];
        }
        if (message.cards) {
            obj.cards = message.cards.map((e) => Math.round(e));
        }
        else {
            obj.cards = [];
        }
        if (message.voteRights) {
            obj.voteRights = message.voteRights.map((e) => e ? vote_right_1.VoteRight.toJSON(e) : undefined);
        }
        else {
            obj.voteRights = [];
        }
        message.councilParticipation !== undefined && (obj.councilParticipation = message.councilParticipation
            ? exports.CouncilParticipation.toJSON(message.councilParticipation)
            : undefined);
        message.ReportMatches !== undefined && (obj.ReportMatches = message.ReportMatches);
        message.profileCard !== undefined && (obj.profileCard = Math.round(message.profileCard));
        message.airDrops !== undefined && (obj.airDrops = message.airDrops ? exports.AirDrops.toJSON(message.airDrops) : undefined);
        if (message.boosterPacks) {
            obj.boosterPacks = message.boosterPacks.map((e) => e ? exports.BoosterPack.toJSON(e) : undefined);
        }
        else {
            obj.boosterPacks = [];
        }
        message.website !== undefined && (obj.website = message.website);
        message.biography !== undefined && (obj.biography = message.biography);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseUser();
        message.alias = object.alias ?? "";
        message.ownedCardSchemes = object.ownedCardSchemes?.map((e) => e) || [];
        message.ownedPrototypes = object.ownedPrototypes?.map((e) => e) || [];
        message.cards = object.cards?.map((e) => e) || [];
        message.voteRights = object.voteRights?.map((e) => vote_right_1.VoteRight.fromPartial(e)) || [];
        message.councilParticipation = (object.councilParticipation !== undefined && object.councilParticipation !== null)
            ? exports.CouncilParticipation.fromPartial(object.councilParticipation)
            : undefined;
        message.ReportMatches = object.ReportMatches ?? false;
        message.profileCard = object.profileCard ?? 0;
        message.airDrops = (object.airDrops !== undefined && object.airDrops !== null)
            ? exports.AirDrops.fromPartial(object.airDrops)
            : undefined;
        message.boosterPacks = object.boosterPacks?.map((e) => exports.BoosterPack.fromPartial(e)) || [];
        message.website = object.website ?? "";
        message.biography = object.biography ?? "";
        return message;
    },
};
function createBaseCouncilParticipation() {
    return { status: 0, council: 0 };
}
exports.CouncilParticipation = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.status !== 0) {
            writer.uint32(8).int32(message.status);
        }
        if (message.council !== 0) {
            writer.uint32(16).uint64(message.council);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseCouncilParticipation();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.status = reader.int32();
                    break;
                case 2:
                    message.council = longToNumber(reader.uint64());
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
            status: isSet(object.status) ? councilStatusFromJSON(object.status) : 0,
            council: isSet(object.council) ? Number(object.council) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        message.status !== undefined && (obj.status = councilStatusToJSON(message.status));
        message.council !== undefined && (obj.council = Math.round(message.council));
        return obj;
    },
    fromPartial(object) {
        const message = createBaseCouncilParticipation();
        message.status = object.status ?? 0;
        message.council = object.council ?? 0;
        return message;
    },
};
function createBaseBoosterPack() {
    return { collectionId: 0, timeStamp: 0, raritiesPerPack: [] };
}
exports.BoosterPack = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.collectionId !== 0) {
            writer.uint32(8).uint64(message.collectionId);
        }
        if (message.timeStamp !== 0) {
            writer.uint32(16).int64(message.timeStamp);
        }
        writer.uint32(26).fork();
        for (const v of message.raritiesPerPack) {
            writer.uint64(v);
        }
        writer.ldelim();
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseBoosterPack();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.collectionId = longToNumber(reader.uint64());
                    break;
                case 2:
                    message.timeStamp = longToNumber(reader.int64());
                    break;
                case 3:
                    if ((tag & 7) === 2) {
                        const end2 = reader.uint32() + reader.pos;
                        while (reader.pos < end2) {
                            message.raritiesPerPack.push(longToNumber(reader.uint64()));
                        }
                    }
                    else {
                        message.raritiesPerPack.push(longToNumber(reader.uint64()));
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
            collectionId: isSet(object.collectionId) ? Number(object.collectionId) : 0,
            timeStamp: isSet(object.timeStamp) ? Number(object.timeStamp) : 0,
            raritiesPerPack: Array.isArray(object?.raritiesPerPack) ? object.raritiesPerPack.map((e) => Number(e)) : [],
        };
    },
    toJSON(message) {
        const obj = {};
        message.collectionId !== undefined && (obj.collectionId = Math.round(message.collectionId));
        message.timeStamp !== undefined && (obj.timeStamp = Math.round(message.timeStamp));
        if (message.raritiesPerPack) {
            obj.raritiesPerPack = message.raritiesPerPack.map((e) => Math.round(e));
        }
        else {
            obj.raritiesPerPack = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = createBaseBoosterPack();
        message.collectionId = object.collectionId ?? 0;
        message.timeStamp = object.timeStamp ?? 0;
        message.raritiesPerPack = object.raritiesPerPack?.map((e) => e) || [];
        return message;
    },
};
function createBaseAirDrops() {
    return { vote: false, create: false, buy: false, play: false, user: false };
}
exports.AirDrops = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.vote === true) {
            writer.uint32(8).bool(message.vote);
        }
        if (message.create === true) {
            writer.uint32(16).bool(message.create);
        }
        if (message.buy === true) {
            writer.uint32(24).bool(message.buy);
        }
        if (message.play === true) {
            writer.uint32(32).bool(message.play);
        }
        if (message.user === true) {
            writer.uint32(40).bool(message.user);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseAirDrops();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.vote = reader.bool();
                    break;
                case 2:
                    message.create = reader.bool();
                    break;
                case 3:
                    message.buy = reader.bool();
                    break;
                case 4:
                    message.play = reader.bool();
                    break;
                case 5:
                    message.user = reader.bool();
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
            vote: isSet(object.vote) ? Boolean(object.vote) : false,
            create: isSet(object.create) ? Boolean(object.create) : false,
            buy: isSet(object.buy) ? Boolean(object.buy) : false,
            play: isSet(object.play) ? Boolean(object.play) : false,
            user: isSet(object.user) ? Boolean(object.user) : false,
        };
    },
    toJSON(message) {
        const obj = {};
        message.vote !== undefined && (obj.vote = message.vote);
        message.create !== undefined && (obj.create = message.create);
        message.buy !== undefined && (obj.buy = message.buy);
        message.play !== undefined && (obj.play = message.play);
        message.user !== undefined && (obj.user = message.user);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseAirDrops();
        message.vote = object.vote ?? false;
        message.create = object.create ?? false;
        message.buy = object.buy ?? false;
        message.play = object.play ?? false;
        message.user = object.user ?? false;
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
