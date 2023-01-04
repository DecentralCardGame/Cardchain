"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.SellOffer = exports.sellOfferStatusToJSON = exports.sellOfferStatusFromJSON = exports.SellOfferStatus = exports.protobufPackage = void 0;
/* eslint-disable */
const long_1 = require("long");
const minimal_1 = require("protobufjs/minimal");
exports.protobufPackage = "DecentralCardGame.cardchain.cardchain";
var SellOfferStatus;
(function (SellOfferStatus) {
    SellOfferStatus[SellOfferStatus["open"] = 0] = "open";
    SellOfferStatus[SellOfferStatus["sold"] = 1] = "sold";
    SellOfferStatus[SellOfferStatus["removed"] = 2] = "removed";
    SellOfferStatus[SellOfferStatus["UNRECOGNIZED"] = -1] = "UNRECOGNIZED";
})(SellOfferStatus = exports.SellOfferStatus || (exports.SellOfferStatus = {}));
function sellOfferStatusFromJSON(object) {
    switch (object) {
        case 0:
        case "open":
            return SellOfferStatus.open;
        case 1:
        case "sold":
            return SellOfferStatus.sold;
        case 2:
        case "removed":
            return SellOfferStatus.removed;
        case -1:
        case "UNRECOGNIZED":
        default:
            return SellOfferStatus.UNRECOGNIZED;
    }
}
exports.sellOfferStatusFromJSON = sellOfferStatusFromJSON;
function sellOfferStatusToJSON(object) {
    switch (object) {
        case SellOfferStatus.open:
            return "open";
        case SellOfferStatus.sold:
            return "sold";
        case SellOfferStatus.removed:
            return "removed";
        case SellOfferStatus.UNRECOGNIZED:
        default:
            return "UNRECOGNIZED";
    }
}
exports.sellOfferStatusToJSON = sellOfferStatusToJSON;
function createBaseSellOffer() {
    return { seller: "", buyer: "", card: 0, price: "", status: 0 };
}
exports.SellOffer = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.seller !== "") {
            writer.uint32(10).string(message.seller);
        }
        if (message.buyer !== "") {
            writer.uint32(18).string(message.buyer);
        }
        if (message.card !== 0) {
            writer.uint32(24).uint64(message.card);
        }
        if (message.price !== "") {
            writer.uint32(34).string(message.price);
        }
        if (message.status !== 0) {
            writer.uint32(40).int32(message.status);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseSellOffer();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.seller = reader.string();
                    break;
                case 2:
                    message.buyer = reader.string();
                    break;
                case 3:
                    message.card = longToNumber(reader.uint64());
                    break;
                case 4:
                    message.price = reader.string();
                    break;
                case 5:
                    message.status = reader.int32();
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
            seller: isSet(object.seller) ? String(object.seller) : "",
            buyer: isSet(object.buyer) ? String(object.buyer) : "",
            card: isSet(object.card) ? Number(object.card) : 0,
            price: isSet(object.price) ? String(object.price) : "",
            status: isSet(object.status) ? sellOfferStatusFromJSON(object.status) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        message.seller !== undefined && (obj.seller = message.seller);
        message.buyer !== undefined && (obj.buyer = message.buyer);
        message.card !== undefined && (obj.card = Math.round(message.card));
        message.price !== undefined && (obj.price = message.price);
        message.status !== undefined && (obj.status = sellOfferStatusToJSON(message.status));
        return obj;
    },
    fromPartial(object) {
        const message = createBaseSellOffer();
        message.seller = object.seller ?? "";
        message.buyer = object.buyer ?? "";
        message.card = object.card ?? 0;
        message.price = object.price ?? "";
        message.status = object.status ?? 0;
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
