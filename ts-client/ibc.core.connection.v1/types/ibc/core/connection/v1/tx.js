"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.MsgClientImpl = exports.MsgConnectionOpenConfirmResponse = exports.MsgConnectionOpenConfirm = exports.MsgConnectionOpenAckResponse = exports.MsgConnectionOpenAck = exports.MsgConnectionOpenTryResponse = exports.MsgConnectionOpenTry = exports.MsgConnectionOpenInitResponse = exports.MsgConnectionOpenInit = exports.protobufPackage = void 0;
/* eslint-disable */
const long_1 = require("long");
const minimal_1 = require("protobufjs/minimal");
const any_1 = require("../../../../google/protobuf/any");
const client_1 = require("../../client/v1/client");
const connection_1 = require("./connection");
exports.protobufPackage = "ibc.core.connection.v1";
function createBaseMsgConnectionOpenInit() {
    return { clientId: "", counterparty: undefined, version: undefined, delayPeriod: 0, signer: "" };
}
exports.MsgConnectionOpenInit = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.clientId !== "") {
            writer.uint32(10).string(message.clientId);
        }
        if (message.counterparty !== undefined) {
            connection_1.Counterparty.encode(message.counterparty, writer.uint32(18).fork()).ldelim();
        }
        if (message.version !== undefined) {
            connection_1.Version.encode(message.version, writer.uint32(26).fork()).ldelim();
        }
        if (message.delayPeriod !== 0) {
            writer.uint32(32).uint64(message.delayPeriod);
        }
        if (message.signer !== "") {
            writer.uint32(42).string(message.signer);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgConnectionOpenInit();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.clientId = reader.string();
                    break;
                case 2:
                    message.counterparty = connection_1.Counterparty.decode(reader, reader.uint32());
                    break;
                case 3:
                    message.version = connection_1.Version.decode(reader, reader.uint32());
                    break;
                case 4:
                    message.delayPeriod = longToNumber(reader.uint64());
                    break;
                case 5:
                    message.signer = reader.string();
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
            clientId: isSet(object.clientId) ? String(object.clientId) : "",
            counterparty: isSet(object.counterparty) ? connection_1.Counterparty.fromJSON(object.counterparty) : undefined,
            version: isSet(object.version) ? connection_1.Version.fromJSON(object.version) : undefined,
            delayPeriod: isSet(object.delayPeriod) ? Number(object.delayPeriod) : 0,
            signer: isSet(object.signer) ? String(object.signer) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        message.clientId !== undefined && (obj.clientId = message.clientId);
        message.counterparty !== undefined
            && (obj.counterparty = message.counterparty ? connection_1.Counterparty.toJSON(message.counterparty) : undefined);
        message.version !== undefined && (obj.version = message.version ? connection_1.Version.toJSON(message.version) : undefined);
        message.delayPeriod !== undefined && (obj.delayPeriod = Math.round(message.delayPeriod));
        message.signer !== undefined && (obj.signer = message.signer);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgConnectionOpenInit();
        message.clientId = object.clientId ?? "";
        message.counterparty = (object.counterparty !== undefined && object.counterparty !== null)
            ? connection_1.Counterparty.fromPartial(object.counterparty)
            : undefined;
        message.version = (object.version !== undefined && object.version !== null)
            ? connection_1.Version.fromPartial(object.version)
            : undefined;
        message.delayPeriod = object.delayPeriod ?? 0;
        message.signer = object.signer ?? "";
        return message;
    },
};
function createBaseMsgConnectionOpenInitResponse() {
    return {};
}
exports.MsgConnectionOpenInitResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgConnectionOpenInitResponse();
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
        const message = createBaseMsgConnectionOpenInitResponse();
        return message;
    },
};
function createBaseMsgConnectionOpenTry() {
    return {
        clientId: "",
        previousConnectionId: "",
        clientState: undefined,
        counterparty: undefined,
        delayPeriod: 0,
        counterpartyVersions: [],
        proofHeight: undefined,
        proofInit: new Uint8Array(),
        proofClient: new Uint8Array(),
        proofConsensus: new Uint8Array(),
        consensusHeight: undefined,
        signer: "",
    };
}
exports.MsgConnectionOpenTry = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.clientId !== "") {
            writer.uint32(10).string(message.clientId);
        }
        if (message.previousConnectionId !== "") {
            writer.uint32(18).string(message.previousConnectionId);
        }
        if (message.clientState !== undefined) {
            any_1.Any.encode(message.clientState, writer.uint32(26).fork()).ldelim();
        }
        if (message.counterparty !== undefined) {
            connection_1.Counterparty.encode(message.counterparty, writer.uint32(34).fork()).ldelim();
        }
        if (message.delayPeriod !== 0) {
            writer.uint32(40).uint64(message.delayPeriod);
        }
        for (const v of message.counterpartyVersions) {
            connection_1.Version.encode(v, writer.uint32(50).fork()).ldelim();
        }
        if (message.proofHeight !== undefined) {
            client_1.Height.encode(message.proofHeight, writer.uint32(58).fork()).ldelim();
        }
        if (message.proofInit.length !== 0) {
            writer.uint32(66).bytes(message.proofInit);
        }
        if (message.proofClient.length !== 0) {
            writer.uint32(74).bytes(message.proofClient);
        }
        if (message.proofConsensus.length !== 0) {
            writer.uint32(82).bytes(message.proofConsensus);
        }
        if (message.consensusHeight !== undefined) {
            client_1.Height.encode(message.consensusHeight, writer.uint32(90).fork()).ldelim();
        }
        if (message.signer !== "") {
            writer.uint32(98).string(message.signer);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgConnectionOpenTry();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.clientId = reader.string();
                    break;
                case 2:
                    message.previousConnectionId = reader.string();
                    break;
                case 3:
                    message.clientState = any_1.Any.decode(reader, reader.uint32());
                    break;
                case 4:
                    message.counterparty = connection_1.Counterparty.decode(reader, reader.uint32());
                    break;
                case 5:
                    message.delayPeriod = longToNumber(reader.uint64());
                    break;
                case 6:
                    message.counterpartyVersions.push(connection_1.Version.decode(reader, reader.uint32()));
                    break;
                case 7:
                    message.proofHeight = client_1.Height.decode(reader, reader.uint32());
                    break;
                case 8:
                    message.proofInit = reader.bytes();
                    break;
                case 9:
                    message.proofClient = reader.bytes();
                    break;
                case 10:
                    message.proofConsensus = reader.bytes();
                    break;
                case 11:
                    message.consensusHeight = client_1.Height.decode(reader, reader.uint32());
                    break;
                case 12:
                    message.signer = reader.string();
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
            clientId: isSet(object.clientId) ? String(object.clientId) : "",
            previousConnectionId: isSet(object.previousConnectionId) ? String(object.previousConnectionId) : "",
            clientState: isSet(object.clientState) ? any_1.Any.fromJSON(object.clientState) : undefined,
            counterparty: isSet(object.counterparty) ? connection_1.Counterparty.fromJSON(object.counterparty) : undefined,
            delayPeriod: isSet(object.delayPeriod) ? Number(object.delayPeriod) : 0,
            counterpartyVersions: Array.isArray(object?.counterpartyVersions)
                ? object.counterpartyVersions.map((e) => connection_1.Version.fromJSON(e))
                : [],
            proofHeight: isSet(object.proofHeight) ? client_1.Height.fromJSON(object.proofHeight) : undefined,
            proofInit: isSet(object.proofInit) ? bytesFromBase64(object.proofInit) : new Uint8Array(),
            proofClient: isSet(object.proofClient) ? bytesFromBase64(object.proofClient) : new Uint8Array(),
            proofConsensus: isSet(object.proofConsensus) ? bytesFromBase64(object.proofConsensus) : new Uint8Array(),
            consensusHeight: isSet(object.consensusHeight) ? client_1.Height.fromJSON(object.consensusHeight) : undefined,
            signer: isSet(object.signer) ? String(object.signer) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        message.clientId !== undefined && (obj.clientId = message.clientId);
        message.previousConnectionId !== undefined && (obj.previousConnectionId = message.previousConnectionId);
        message.clientState !== undefined
            && (obj.clientState = message.clientState ? any_1.Any.toJSON(message.clientState) : undefined);
        message.counterparty !== undefined
            && (obj.counterparty = message.counterparty ? connection_1.Counterparty.toJSON(message.counterparty) : undefined);
        message.delayPeriod !== undefined && (obj.delayPeriod = Math.round(message.delayPeriod));
        if (message.counterpartyVersions) {
            obj.counterpartyVersions = message.counterpartyVersions.map((e) => e ? connection_1.Version.toJSON(e) : undefined);
        }
        else {
            obj.counterpartyVersions = [];
        }
        message.proofHeight !== undefined
            && (obj.proofHeight = message.proofHeight ? client_1.Height.toJSON(message.proofHeight) : undefined);
        message.proofInit !== undefined
            && (obj.proofInit = base64FromBytes(message.proofInit !== undefined ? message.proofInit : new Uint8Array()));
        message.proofClient !== undefined
            && (obj.proofClient = base64FromBytes(message.proofClient !== undefined ? message.proofClient : new Uint8Array()));
        message.proofConsensus !== undefined
            && (obj.proofConsensus = base64FromBytes(message.proofConsensus !== undefined ? message.proofConsensus : new Uint8Array()));
        message.consensusHeight !== undefined
            && (obj.consensusHeight = message.consensusHeight ? client_1.Height.toJSON(message.consensusHeight) : undefined);
        message.signer !== undefined && (obj.signer = message.signer);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgConnectionOpenTry();
        message.clientId = object.clientId ?? "";
        message.previousConnectionId = object.previousConnectionId ?? "";
        message.clientState = (object.clientState !== undefined && object.clientState !== null)
            ? any_1.Any.fromPartial(object.clientState)
            : undefined;
        message.counterparty = (object.counterparty !== undefined && object.counterparty !== null)
            ? connection_1.Counterparty.fromPartial(object.counterparty)
            : undefined;
        message.delayPeriod = object.delayPeriod ?? 0;
        message.counterpartyVersions = object.counterpartyVersions?.map((e) => connection_1.Version.fromPartial(e)) || [];
        message.proofHeight = (object.proofHeight !== undefined && object.proofHeight !== null)
            ? client_1.Height.fromPartial(object.proofHeight)
            : undefined;
        message.proofInit = object.proofInit ?? new Uint8Array();
        message.proofClient = object.proofClient ?? new Uint8Array();
        message.proofConsensus = object.proofConsensus ?? new Uint8Array();
        message.consensusHeight = (object.consensusHeight !== undefined && object.consensusHeight !== null)
            ? client_1.Height.fromPartial(object.consensusHeight)
            : undefined;
        message.signer = object.signer ?? "";
        return message;
    },
};
function createBaseMsgConnectionOpenTryResponse() {
    return {};
}
exports.MsgConnectionOpenTryResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgConnectionOpenTryResponse();
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
        const message = createBaseMsgConnectionOpenTryResponse();
        return message;
    },
};
function createBaseMsgConnectionOpenAck() {
    return {
        connectionId: "",
        counterpartyConnectionId: "",
        version: undefined,
        clientState: undefined,
        proofHeight: undefined,
        proofTry: new Uint8Array(),
        proofClient: new Uint8Array(),
        proofConsensus: new Uint8Array(),
        consensusHeight: undefined,
        signer: "",
    };
}
exports.MsgConnectionOpenAck = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.connectionId !== "") {
            writer.uint32(10).string(message.connectionId);
        }
        if (message.counterpartyConnectionId !== "") {
            writer.uint32(18).string(message.counterpartyConnectionId);
        }
        if (message.version !== undefined) {
            connection_1.Version.encode(message.version, writer.uint32(26).fork()).ldelim();
        }
        if (message.clientState !== undefined) {
            any_1.Any.encode(message.clientState, writer.uint32(34).fork()).ldelim();
        }
        if (message.proofHeight !== undefined) {
            client_1.Height.encode(message.proofHeight, writer.uint32(42).fork()).ldelim();
        }
        if (message.proofTry.length !== 0) {
            writer.uint32(50).bytes(message.proofTry);
        }
        if (message.proofClient.length !== 0) {
            writer.uint32(58).bytes(message.proofClient);
        }
        if (message.proofConsensus.length !== 0) {
            writer.uint32(66).bytes(message.proofConsensus);
        }
        if (message.consensusHeight !== undefined) {
            client_1.Height.encode(message.consensusHeight, writer.uint32(74).fork()).ldelim();
        }
        if (message.signer !== "") {
            writer.uint32(82).string(message.signer);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgConnectionOpenAck();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.connectionId = reader.string();
                    break;
                case 2:
                    message.counterpartyConnectionId = reader.string();
                    break;
                case 3:
                    message.version = connection_1.Version.decode(reader, reader.uint32());
                    break;
                case 4:
                    message.clientState = any_1.Any.decode(reader, reader.uint32());
                    break;
                case 5:
                    message.proofHeight = client_1.Height.decode(reader, reader.uint32());
                    break;
                case 6:
                    message.proofTry = reader.bytes();
                    break;
                case 7:
                    message.proofClient = reader.bytes();
                    break;
                case 8:
                    message.proofConsensus = reader.bytes();
                    break;
                case 9:
                    message.consensusHeight = client_1.Height.decode(reader, reader.uint32());
                    break;
                case 10:
                    message.signer = reader.string();
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
            connectionId: isSet(object.connectionId) ? String(object.connectionId) : "",
            counterpartyConnectionId: isSet(object.counterpartyConnectionId) ? String(object.counterpartyConnectionId) : "",
            version: isSet(object.version) ? connection_1.Version.fromJSON(object.version) : undefined,
            clientState: isSet(object.clientState) ? any_1.Any.fromJSON(object.clientState) : undefined,
            proofHeight: isSet(object.proofHeight) ? client_1.Height.fromJSON(object.proofHeight) : undefined,
            proofTry: isSet(object.proofTry) ? bytesFromBase64(object.proofTry) : new Uint8Array(),
            proofClient: isSet(object.proofClient) ? bytesFromBase64(object.proofClient) : new Uint8Array(),
            proofConsensus: isSet(object.proofConsensus) ? bytesFromBase64(object.proofConsensus) : new Uint8Array(),
            consensusHeight: isSet(object.consensusHeight) ? client_1.Height.fromJSON(object.consensusHeight) : undefined,
            signer: isSet(object.signer) ? String(object.signer) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        message.connectionId !== undefined && (obj.connectionId = message.connectionId);
        message.counterpartyConnectionId !== undefined && (obj.counterpartyConnectionId = message.counterpartyConnectionId);
        message.version !== undefined && (obj.version = message.version ? connection_1.Version.toJSON(message.version) : undefined);
        message.clientState !== undefined
            && (obj.clientState = message.clientState ? any_1.Any.toJSON(message.clientState) : undefined);
        message.proofHeight !== undefined
            && (obj.proofHeight = message.proofHeight ? client_1.Height.toJSON(message.proofHeight) : undefined);
        message.proofTry !== undefined
            && (obj.proofTry = base64FromBytes(message.proofTry !== undefined ? message.proofTry : new Uint8Array()));
        message.proofClient !== undefined
            && (obj.proofClient = base64FromBytes(message.proofClient !== undefined ? message.proofClient : new Uint8Array()));
        message.proofConsensus !== undefined
            && (obj.proofConsensus = base64FromBytes(message.proofConsensus !== undefined ? message.proofConsensus : new Uint8Array()));
        message.consensusHeight !== undefined
            && (obj.consensusHeight = message.consensusHeight ? client_1.Height.toJSON(message.consensusHeight) : undefined);
        message.signer !== undefined && (obj.signer = message.signer);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgConnectionOpenAck();
        message.connectionId = object.connectionId ?? "";
        message.counterpartyConnectionId = object.counterpartyConnectionId ?? "";
        message.version = (object.version !== undefined && object.version !== null)
            ? connection_1.Version.fromPartial(object.version)
            : undefined;
        message.clientState = (object.clientState !== undefined && object.clientState !== null)
            ? any_1.Any.fromPartial(object.clientState)
            : undefined;
        message.proofHeight = (object.proofHeight !== undefined && object.proofHeight !== null)
            ? client_1.Height.fromPartial(object.proofHeight)
            : undefined;
        message.proofTry = object.proofTry ?? new Uint8Array();
        message.proofClient = object.proofClient ?? new Uint8Array();
        message.proofConsensus = object.proofConsensus ?? new Uint8Array();
        message.consensusHeight = (object.consensusHeight !== undefined && object.consensusHeight !== null)
            ? client_1.Height.fromPartial(object.consensusHeight)
            : undefined;
        message.signer = object.signer ?? "";
        return message;
    },
};
function createBaseMsgConnectionOpenAckResponse() {
    return {};
}
exports.MsgConnectionOpenAckResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgConnectionOpenAckResponse();
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
        const message = createBaseMsgConnectionOpenAckResponse();
        return message;
    },
};
function createBaseMsgConnectionOpenConfirm() {
    return { connectionId: "", proofAck: new Uint8Array(), proofHeight: undefined, signer: "" };
}
exports.MsgConnectionOpenConfirm = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.connectionId !== "") {
            writer.uint32(10).string(message.connectionId);
        }
        if (message.proofAck.length !== 0) {
            writer.uint32(18).bytes(message.proofAck);
        }
        if (message.proofHeight !== undefined) {
            client_1.Height.encode(message.proofHeight, writer.uint32(26).fork()).ldelim();
        }
        if (message.signer !== "") {
            writer.uint32(34).string(message.signer);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgConnectionOpenConfirm();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.connectionId = reader.string();
                    break;
                case 2:
                    message.proofAck = reader.bytes();
                    break;
                case 3:
                    message.proofHeight = client_1.Height.decode(reader, reader.uint32());
                    break;
                case 4:
                    message.signer = reader.string();
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
            connectionId: isSet(object.connectionId) ? String(object.connectionId) : "",
            proofAck: isSet(object.proofAck) ? bytesFromBase64(object.proofAck) : new Uint8Array(),
            proofHeight: isSet(object.proofHeight) ? client_1.Height.fromJSON(object.proofHeight) : undefined,
            signer: isSet(object.signer) ? String(object.signer) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        message.connectionId !== undefined && (obj.connectionId = message.connectionId);
        message.proofAck !== undefined
            && (obj.proofAck = base64FromBytes(message.proofAck !== undefined ? message.proofAck : new Uint8Array()));
        message.proofHeight !== undefined
            && (obj.proofHeight = message.proofHeight ? client_1.Height.toJSON(message.proofHeight) : undefined);
        message.signer !== undefined && (obj.signer = message.signer);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgConnectionOpenConfirm();
        message.connectionId = object.connectionId ?? "";
        message.proofAck = object.proofAck ?? new Uint8Array();
        message.proofHeight = (object.proofHeight !== undefined && object.proofHeight !== null)
            ? client_1.Height.fromPartial(object.proofHeight)
            : undefined;
        message.signer = object.signer ?? "";
        return message;
    },
};
function createBaseMsgConnectionOpenConfirmResponse() {
    return {};
}
exports.MsgConnectionOpenConfirmResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgConnectionOpenConfirmResponse();
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
        const message = createBaseMsgConnectionOpenConfirmResponse();
        return message;
    },
};
class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
        this.ConnectionOpenInit = this.ConnectionOpenInit.bind(this);
        this.ConnectionOpenTry = this.ConnectionOpenTry.bind(this);
        this.ConnectionOpenAck = this.ConnectionOpenAck.bind(this);
        this.ConnectionOpenConfirm = this.ConnectionOpenConfirm.bind(this);
    }
    ConnectionOpenInit(request) {
        const data = exports.MsgConnectionOpenInit.encode(request).finish();
        const promise = this.rpc.request("ibc.core.connection.v1.Msg", "ConnectionOpenInit", data);
        return promise.then((data) => exports.MsgConnectionOpenInitResponse.decode(new minimal_1.default.Reader(data)));
    }
    ConnectionOpenTry(request) {
        const data = exports.MsgConnectionOpenTry.encode(request).finish();
        const promise = this.rpc.request("ibc.core.connection.v1.Msg", "ConnectionOpenTry", data);
        return promise.then((data) => exports.MsgConnectionOpenTryResponse.decode(new minimal_1.default.Reader(data)));
    }
    ConnectionOpenAck(request) {
        const data = exports.MsgConnectionOpenAck.encode(request).finish();
        const promise = this.rpc.request("ibc.core.connection.v1.Msg", "ConnectionOpenAck", data);
        return promise.then((data) => exports.MsgConnectionOpenAckResponse.decode(new minimal_1.default.Reader(data)));
    }
    ConnectionOpenConfirm(request) {
        const data = exports.MsgConnectionOpenConfirm.encode(request).finish();
        const promise = this.rpc.request("ibc.core.connection.v1.Msg", "ConnectionOpenConfirm", data);
        return promise.then((data) => exports.MsgConnectionOpenConfirmResponse.decode(new minimal_1.default.Reader(data)));
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
