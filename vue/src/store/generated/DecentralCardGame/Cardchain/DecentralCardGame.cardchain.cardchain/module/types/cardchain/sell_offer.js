/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "DecentralCardGame.cardchain.cardchain";
export var SellOfferStatus;
(function (SellOfferStatus) {
    SellOfferStatus[SellOfferStatus["open"] = 0] = "open";
    SellOfferStatus[SellOfferStatus["sold"] = 1] = "sold";
    SellOfferStatus[SellOfferStatus["removed"] = 2] = "removed";
    SellOfferStatus[SellOfferStatus["UNRECOGNIZED"] = -1] = "UNRECOGNIZED";
})(SellOfferStatus || (SellOfferStatus = {}));
export function sellOfferStatusFromJSON(object) {
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
export function sellOfferStatusToJSON(object) {
    switch (object) {
        case SellOfferStatus.open:
            return "open";
        case SellOfferStatus.sold:
            return "sold";
        case SellOfferStatus.removed:
            return "removed";
        default:
            return "UNKNOWN";
    }
}
const baseSellOffer = {
    seller: "",
    buyer: "",
    card: 0,
    price: 0,
    status: 0,
};
export const SellOffer = {
    encode(message, writer = Writer.create()) {
        if (message.seller !== "") {
            writer.uint32(10).string(message.seller);
        }
        if (message.buyer !== "") {
            writer.uint32(18).string(message.buyer);
        }
        if (message.card !== 0) {
            writer.uint32(24).uint64(message.card);
        }
        if (message.price !== 0) {
            writer.uint32(32).uint64(message.price);
        }
        if (message.status !== 0) {
            writer.uint32(40).int32(message.status);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseSellOffer };
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
                    message.price = longToNumber(reader.uint64());
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
        const message = { ...baseSellOffer };
        if (object.seller !== undefined && object.seller !== null) {
            message.seller = String(object.seller);
        }
        else {
            message.seller = "";
        }
        if (object.buyer !== undefined && object.buyer !== null) {
            message.buyer = String(object.buyer);
        }
        else {
            message.buyer = "";
        }
        if (object.card !== undefined && object.card !== null) {
            message.card = Number(object.card);
        }
        else {
            message.card = 0;
        }
        if (object.price !== undefined && object.price !== null) {
            message.price = Number(object.price);
        }
        else {
            message.price = 0;
        }
        if (object.status !== undefined && object.status !== null) {
            message.status = sellOfferStatusFromJSON(object.status);
        }
        else {
            message.status = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.seller !== undefined && (obj.seller = message.seller);
        message.buyer !== undefined && (obj.buyer = message.buyer);
        message.card !== undefined && (obj.card = message.card);
        message.price !== undefined && (obj.price = message.price);
        message.status !== undefined &&
            (obj.status = sellOfferStatusToJSON(message.status));
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseSellOffer };
        if (object.seller !== undefined && object.seller !== null) {
            message.seller = object.seller;
        }
        else {
            message.seller = "";
        }
        if (object.buyer !== undefined && object.buyer !== null) {
            message.buyer = object.buyer;
        }
        else {
            message.buyer = "";
        }
        if (object.card !== undefined && object.card !== null) {
            message.card = object.card;
        }
        else {
            message.card = 0;
        }
        if (object.price !== undefined && object.price !== null) {
            message.price = object.price;
        }
        else {
            message.price = 0;
        }
        if (object.status !== undefined && object.status !== null) {
            message.status = object.status;
        }
        else {
            message.status = 0;
        }
        return message;
    },
};
var globalThis = (() => {
    if (typeof globalThis !== "undefined")
        return globalThis;
    if (typeof self !== "undefined")
        return self;
    if (typeof window !== "undefined")
        return window;
    if (typeof global !== "undefined")
        return global;
    throw "Unable to locate global object";
})();
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}
