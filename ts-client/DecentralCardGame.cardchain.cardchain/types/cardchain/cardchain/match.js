"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.MatchPlayer = exports.Match = exports.protobufPackage = void 0;
/* eslint-disable */
const long_1 = require("long");
const minimal_1 = require("protobufjs/minimal");
const tx_1 = require("./tx");
exports.protobufPackage = "DecentralCardGame.cardchain.cardchain";
function createBaseMatch() {
    return { timestamp: 0, reporter: "", playerA: undefined, playerB: undefined, outcome: 0, coinsDistributed: false };
}
exports.Match = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.timestamp !== 0) {
            writer.uint32(8).uint64(message.timestamp);
        }
        if (message.reporter !== "") {
            writer.uint32(18).string(message.reporter);
        }
        if (message.playerA !== undefined) {
            exports.MatchPlayer.encode(message.playerA, writer.uint32(26).fork()).ldelim();
        }
        if (message.playerB !== undefined) {
            exports.MatchPlayer.encode(message.playerB, writer.uint32(34).fork()).ldelim();
        }
        if (message.outcome !== 0) {
            writer.uint32(56).int32(message.outcome);
        }
        if (message.coinsDistributed === true) {
            writer.uint32(80).bool(message.coinsDistributed);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMatch();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.timestamp = longToNumber(reader.uint64());
                    break;
                case 2:
                    message.reporter = reader.string();
                    break;
                case 3:
                    message.playerA = exports.MatchPlayer.decode(reader, reader.uint32());
                    break;
                case 4:
                    message.playerB = exports.MatchPlayer.decode(reader, reader.uint32());
                    break;
                case 7:
                    message.outcome = reader.int32();
                    break;
                case 10:
                    message.coinsDistributed = reader.bool();
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
            timestamp: isSet(object.timestamp) ? Number(object.timestamp) : 0,
            reporter: isSet(object.reporter) ? String(object.reporter) : "",
            playerA: isSet(object.playerA) ? exports.MatchPlayer.fromJSON(object.playerA) : undefined,
            playerB: isSet(object.playerB) ? exports.MatchPlayer.fromJSON(object.playerB) : undefined,
            outcome: isSet(object.outcome) ? (0, tx_1.outcomeFromJSON)(object.outcome) : 0,
            coinsDistributed: isSet(object.coinsDistributed) ? Boolean(object.coinsDistributed) : false,
        };
    },
    toJSON(message) {
        const obj = {};
        message.timestamp !== undefined && (obj.timestamp = Math.round(message.timestamp));
        message.reporter !== undefined && (obj.reporter = message.reporter);
        message.playerA !== undefined && (obj.playerA = message.playerA ? exports.MatchPlayer.toJSON(message.playerA) : undefined);
        message.playerB !== undefined && (obj.playerB = message.playerB ? exports.MatchPlayer.toJSON(message.playerB) : undefined);
        message.outcome !== undefined && (obj.outcome = (0, tx_1.outcomeToJSON)(message.outcome));
        message.coinsDistributed !== undefined && (obj.coinsDistributed = message.coinsDistributed);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMatch();
        message.timestamp = object.timestamp ?? 0;
        message.reporter = object.reporter ?? "";
        message.playerA = (object.playerA !== undefined && object.playerA !== null)
            ? exports.MatchPlayer.fromPartial(object.playerA)
            : undefined;
        message.playerB = (object.playerB !== undefined && object.playerB !== null)
            ? exports.MatchPlayer.fromPartial(object.playerB)
            : undefined;
        message.outcome = object.outcome ?? 0;
        message.coinsDistributed = object.coinsDistributed ?? false;
        return message;
    },
};
function createBaseMatchPlayer() {
    return { addr: "", playedCards: [], confirmed: false, outcome: 0 };
}
exports.MatchPlayer = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.addr !== "") {
            writer.uint32(10).string(message.addr);
        }
        writer.uint32(18).fork();
        for (const v of message.playedCards) {
            writer.uint64(v);
        }
        writer.ldelim();
        if (message.confirmed === true) {
            writer.uint32(24).bool(message.confirmed);
        }
        if (message.outcome !== 0) {
            writer.uint32(32).int32(message.outcome);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMatchPlayer();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.addr = reader.string();
                    break;
                case 2:
                    if ((tag & 7) === 2) {
                        const end2 = reader.uint32() + reader.pos;
                        while (reader.pos < end2) {
                            message.playedCards.push(longToNumber(reader.uint64()));
                        }
                    }
                    else {
                        message.playedCards.push(longToNumber(reader.uint64()));
                    }
                    break;
                case 3:
                    message.confirmed = reader.bool();
                    break;
                case 4:
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
            addr: isSet(object.addr) ? String(object.addr) : "",
            playedCards: Array.isArray(object?.playedCards) ? object.playedCards.map((e) => Number(e)) : [],
            confirmed: isSet(object.confirmed) ? Boolean(object.confirmed) : false,
            outcome: isSet(object.outcome) ? (0, tx_1.outcomeFromJSON)(object.outcome) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        message.addr !== undefined && (obj.addr = message.addr);
        if (message.playedCards) {
            obj.playedCards = message.playedCards.map((e) => Math.round(e));
        }
        else {
            obj.playedCards = [];
        }
        message.confirmed !== undefined && (obj.confirmed = message.confirmed);
        message.outcome !== undefined && (obj.outcome = (0, tx_1.outcomeToJSON)(message.outcome));
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMatchPlayer();
        message.addr = object.addr ?? "";
        message.playedCards = object.playedCards?.map((e) => e) || [];
        message.confirmed = object.confirmed ?? false;
        message.outcome = object.outcome ?? 0;
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
