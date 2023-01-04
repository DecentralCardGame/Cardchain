"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.MsgClientImpl = exports.MsgSubmitMisbehaviourResponse = exports.MsgSubmitMisbehaviour = exports.MsgUpgradeClientResponse = exports.MsgUpgradeClient = exports.MsgUpdateClientResponse = exports.MsgUpdateClient = exports.MsgCreateClientResponse = exports.MsgCreateClient = exports.protobufPackage = void 0;
/* eslint-disable */
const minimal_1 = require("protobufjs/minimal");
const any_1 = require("../../../../google/protobuf/any");
exports.protobufPackage = "ibc.core.client.v1";
function createBaseMsgCreateClient() {
    return { clientState: undefined, consensusState: undefined, signer: "" };
}
exports.MsgCreateClient = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.clientState !== undefined) {
            any_1.Any.encode(message.clientState, writer.uint32(10).fork()).ldelim();
        }
        if (message.consensusState !== undefined) {
            any_1.Any.encode(message.consensusState, writer.uint32(18).fork()).ldelim();
        }
        if (message.signer !== "") {
            writer.uint32(26).string(message.signer);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgCreateClient();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.clientState = any_1.Any.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.consensusState = any_1.Any.decode(reader, reader.uint32());
                    break;
                case 3:
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
            clientState: isSet(object.clientState) ? any_1.Any.fromJSON(object.clientState) : undefined,
            consensusState: isSet(object.consensusState) ? any_1.Any.fromJSON(object.consensusState) : undefined,
            signer: isSet(object.signer) ? String(object.signer) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        message.clientState !== undefined
            && (obj.clientState = message.clientState ? any_1.Any.toJSON(message.clientState) : undefined);
        message.consensusState !== undefined
            && (obj.consensusState = message.consensusState ? any_1.Any.toJSON(message.consensusState) : undefined);
        message.signer !== undefined && (obj.signer = message.signer);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgCreateClient();
        message.clientState = (object.clientState !== undefined && object.clientState !== null)
            ? any_1.Any.fromPartial(object.clientState)
            : undefined;
        message.consensusState = (object.consensusState !== undefined && object.consensusState !== null)
            ? any_1.Any.fromPartial(object.consensusState)
            : undefined;
        message.signer = object.signer ?? "";
        return message;
    },
};
function createBaseMsgCreateClientResponse() {
    return {};
}
exports.MsgCreateClientResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgCreateClientResponse();
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
        const message = createBaseMsgCreateClientResponse();
        return message;
    },
};
function createBaseMsgUpdateClient() {
    return { clientId: "", header: undefined, signer: "" };
}
exports.MsgUpdateClient = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.clientId !== "") {
            writer.uint32(10).string(message.clientId);
        }
        if (message.header !== undefined) {
            any_1.Any.encode(message.header, writer.uint32(18).fork()).ldelim();
        }
        if (message.signer !== "") {
            writer.uint32(26).string(message.signer);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgUpdateClient();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.clientId = reader.string();
                    break;
                case 2:
                    message.header = any_1.Any.decode(reader, reader.uint32());
                    break;
                case 3:
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
            header: isSet(object.header) ? any_1.Any.fromJSON(object.header) : undefined,
            signer: isSet(object.signer) ? String(object.signer) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        message.clientId !== undefined && (obj.clientId = message.clientId);
        message.header !== undefined && (obj.header = message.header ? any_1.Any.toJSON(message.header) : undefined);
        message.signer !== undefined && (obj.signer = message.signer);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgUpdateClient();
        message.clientId = object.clientId ?? "";
        message.header = (object.header !== undefined && object.header !== null)
            ? any_1.Any.fromPartial(object.header)
            : undefined;
        message.signer = object.signer ?? "";
        return message;
    },
};
function createBaseMsgUpdateClientResponse() {
    return {};
}
exports.MsgUpdateClientResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgUpdateClientResponse();
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
        const message = createBaseMsgUpdateClientResponse();
        return message;
    },
};
function createBaseMsgUpgradeClient() {
    return {
        clientId: "",
        clientState: undefined,
        consensusState: undefined,
        proofUpgradeClient: new Uint8Array(),
        proofUpgradeConsensusState: new Uint8Array(),
        signer: "",
    };
}
exports.MsgUpgradeClient = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.clientId !== "") {
            writer.uint32(10).string(message.clientId);
        }
        if (message.clientState !== undefined) {
            any_1.Any.encode(message.clientState, writer.uint32(18).fork()).ldelim();
        }
        if (message.consensusState !== undefined) {
            any_1.Any.encode(message.consensusState, writer.uint32(26).fork()).ldelim();
        }
        if (message.proofUpgradeClient.length !== 0) {
            writer.uint32(34).bytes(message.proofUpgradeClient);
        }
        if (message.proofUpgradeConsensusState.length !== 0) {
            writer.uint32(42).bytes(message.proofUpgradeConsensusState);
        }
        if (message.signer !== "") {
            writer.uint32(50).string(message.signer);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgUpgradeClient();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.clientId = reader.string();
                    break;
                case 2:
                    message.clientState = any_1.Any.decode(reader, reader.uint32());
                    break;
                case 3:
                    message.consensusState = any_1.Any.decode(reader, reader.uint32());
                    break;
                case 4:
                    message.proofUpgradeClient = reader.bytes();
                    break;
                case 5:
                    message.proofUpgradeConsensusState = reader.bytes();
                    break;
                case 6:
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
            clientState: isSet(object.clientState) ? any_1.Any.fromJSON(object.clientState) : undefined,
            consensusState: isSet(object.consensusState) ? any_1.Any.fromJSON(object.consensusState) : undefined,
            proofUpgradeClient: isSet(object.proofUpgradeClient)
                ? bytesFromBase64(object.proofUpgradeClient)
                : new Uint8Array(),
            proofUpgradeConsensusState: isSet(object.proofUpgradeConsensusState)
                ? bytesFromBase64(object.proofUpgradeConsensusState)
                : new Uint8Array(),
            signer: isSet(object.signer) ? String(object.signer) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        message.clientId !== undefined && (obj.clientId = message.clientId);
        message.clientState !== undefined
            && (obj.clientState = message.clientState ? any_1.Any.toJSON(message.clientState) : undefined);
        message.consensusState !== undefined
            && (obj.consensusState = message.consensusState ? any_1.Any.toJSON(message.consensusState) : undefined);
        message.proofUpgradeClient !== undefined
            && (obj.proofUpgradeClient = base64FromBytes(message.proofUpgradeClient !== undefined ? message.proofUpgradeClient : new Uint8Array()));
        message.proofUpgradeConsensusState !== undefined
            && (obj.proofUpgradeConsensusState = base64FromBytes(message.proofUpgradeConsensusState !== undefined ? message.proofUpgradeConsensusState : new Uint8Array()));
        message.signer !== undefined && (obj.signer = message.signer);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgUpgradeClient();
        message.clientId = object.clientId ?? "";
        message.clientState = (object.clientState !== undefined && object.clientState !== null)
            ? any_1.Any.fromPartial(object.clientState)
            : undefined;
        message.consensusState = (object.consensusState !== undefined && object.consensusState !== null)
            ? any_1.Any.fromPartial(object.consensusState)
            : undefined;
        message.proofUpgradeClient = object.proofUpgradeClient ?? new Uint8Array();
        message.proofUpgradeConsensusState = object.proofUpgradeConsensusState ?? new Uint8Array();
        message.signer = object.signer ?? "";
        return message;
    },
};
function createBaseMsgUpgradeClientResponse() {
    return {};
}
exports.MsgUpgradeClientResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgUpgradeClientResponse();
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
        const message = createBaseMsgUpgradeClientResponse();
        return message;
    },
};
function createBaseMsgSubmitMisbehaviour() {
    return { clientId: "", misbehaviour: undefined, signer: "" };
}
exports.MsgSubmitMisbehaviour = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.clientId !== "") {
            writer.uint32(10).string(message.clientId);
        }
        if (message.misbehaviour !== undefined) {
            any_1.Any.encode(message.misbehaviour, writer.uint32(18).fork()).ldelim();
        }
        if (message.signer !== "") {
            writer.uint32(26).string(message.signer);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgSubmitMisbehaviour();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.clientId = reader.string();
                    break;
                case 2:
                    message.misbehaviour = any_1.Any.decode(reader, reader.uint32());
                    break;
                case 3:
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
            misbehaviour: isSet(object.misbehaviour) ? any_1.Any.fromJSON(object.misbehaviour) : undefined,
            signer: isSet(object.signer) ? String(object.signer) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        message.clientId !== undefined && (obj.clientId = message.clientId);
        message.misbehaviour !== undefined
            && (obj.misbehaviour = message.misbehaviour ? any_1.Any.toJSON(message.misbehaviour) : undefined);
        message.signer !== undefined && (obj.signer = message.signer);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseMsgSubmitMisbehaviour();
        message.clientId = object.clientId ?? "";
        message.misbehaviour = (object.misbehaviour !== undefined && object.misbehaviour !== null)
            ? any_1.Any.fromPartial(object.misbehaviour)
            : undefined;
        message.signer = object.signer ?? "";
        return message;
    },
};
function createBaseMsgSubmitMisbehaviourResponse() {
    return {};
}
exports.MsgSubmitMisbehaviourResponse = {
    encode(_, writer = minimal_1.default.Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseMsgSubmitMisbehaviourResponse();
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
        const message = createBaseMsgSubmitMisbehaviourResponse();
        return message;
    },
};
class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
        this.CreateClient = this.CreateClient.bind(this);
        this.UpdateClient = this.UpdateClient.bind(this);
        this.UpgradeClient = this.UpgradeClient.bind(this);
        this.SubmitMisbehaviour = this.SubmitMisbehaviour.bind(this);
    }
    CreateClient(request) {
        const data = exports.MsgCreateClient.encode(request).finish();
        const promise = this.rpc.request("ibc.core.client.v1.Msg", "CreateClient", data);
        return promise.then((data) => exports.MsgCreateClientResponse.decode(new minimal_1.default.Reader(data)));
    }
    UpdateClient(request) {
        const data = exports.MsgUpdateClient.encode(request).finish();
        const promise = this.rpc.request("ibc.core.client.v1.Msg", "UpdateClient", data);
        return promise.then((data) => exports.MsgUpdateClientResponse.decode(new minimal_1.default.Reader(data)));
    }
    UpgradeClient(request) {
        const data = exports.MsgUpgradeClient.encode(request).finish();
        const promise = this.rpc.request("ibc.core.client.v1.Msg", "UpgradeClient", data);
        return promise.then((data) => exports.MsgUpgradeClientResponse.decode(new minimal_1.default.Reader(data)));
    }
    SubmitMisbehaviour(request) {
        const data = exports.MsgSubmitMisbehaviour.encode(request).finish();
        const promise = this.rpc.request("ibc.core.client.v1.Msg", "SubmitMisbehaviour", data);
        return promise.then((data) => exports.MsgSubmitMisbehaviourResponse.decode(new minimal_1.default.Reader(data)));
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
function isSet(value) {
    return value !== null && value !== undefined;
}
