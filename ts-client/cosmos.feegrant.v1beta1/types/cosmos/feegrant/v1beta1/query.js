"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.QueryClientImpl = exports.QueryAllowancesByGranterResponse = exports.QueryAllowancesByGranterRequest = exports.QueryAllowancesResponse = exports.QueryAllowancesRequest = exports.QueryAllowanceResponse = exports.QueryAllowanceRequest = exports.protobufPackage = void 0;
/* eslint-disable */
const minimal_1 = require("protobufjs/minimal");
const pagination_1 = require("../../base/query/v1beta1/pagination");
const feegrant_1 = require("./feegrant");
exports.protobufPackage = "cosmos.feegrant.v1beta1";
function createBaseQueryAllowanceRequest() {
    return { granter: "", grantee: "" };
}
exports.QueryAllowanceRequest = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.granter !== "") {
            writer.uint32(10).string(message.granter);
        }
        if (message.grantee !== "") {
            writer.uint32(18).string(message.grantee);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseQueryAllowanceRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.granter = reader.string();
                    break;
                case 2:
                    message.grantee = reader.string();
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
            granter: isSet(object.granter) ? String(object.granter) : "",
            grantee: isSet(object.grantee) ? String(object.grantee) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        message.granter !== undefined && (obj.granter = message.granter);
        message.grantee !== undefined && (obj.grantee = message.grantee);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseQueryAllowanceRequest();
        message.granter = object.granter ?? "";
        message.grantee = object.grantee ?? "";
        return message;
    },
};
function createBaseQueryAllowanceResponse() {
    return { allowance: undefined };
}
exports.QueryAllowanceResponse = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.allowance !== undefined) {
            feegrant_1.Grant.encode(message.allowance, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseQueryAllowanceResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.allowance = feegrant_1.Grant.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        return { allowance: isSet(object.allowance) ? feegrant_1.Grant.fromJSON(object.allowance) : undefined };
    },
    toJSON(message) {
        const obj = {};
        message.allowance !== undefined
            && (obj.allowance = message.allowance ? feegrant_1.Grant.toJSON(message.allowance) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseQueryAllowanceResponse();
        message.allowance = (object.allowance !== undefined && object.allowance !== null)
            ? feegrant_1.Grant.fromPartial(object.allowance)
            : undefined;
        return message;
    },
};
function createBaseQueryAllowancesRequest() {
    return { grantee: "", pagination: undefined };
}
exports.QueryAllowancesRequest = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.grantee !== "") {
            writer.uint32(10).string(message.grantee);
        }
        if (message.pagination !== undefined) {
            pagination_1.PageRequest.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseQueryAllowancesRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.grantee = reader.string();
                    break;
                case 2:
                    message.pagination = pagination_1.PageRequest.decode(reader, reader.uint32());
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
            grantee: isSet(object.grantee) ? String(object.grantee) : "",
            pagination: isSet(object.pagination) ? pagination_1.PageRequest.fromJSON(object.pagination) : undefined,
        };
    },
    toJSON(message) {
        const obj = {};
        message.grantee !== undefined && (obj.grantee = message.grantee);
        message.pagination !== undefined
            && (obj.pagination = message.pagination ? pagination_1.PageRequest.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseQueryAllowancesRequest();
        message.grantee = object.grantee ?? "";
        message.pagination = (object.pagination !== undefined && object.pagination !== null)
            ? pagination_1.PageRequest.fromPartial(object.pagination)
            : undefined;
        return message;
    },
};
function createBaseQueryAllowancesResponse() {
    return { allowances: [], pagination: undefined };
}
exports.QueryAllowancesResponse = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        for (const v of message.allowances) {
            feegrant_1.Grant.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            pagination_1.PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseQueryAllowancesResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.allowances.push(feegrant_1.Grant.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.pagination = pagination_1.PageResponse.decode(reader, reader.uint32());
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
            allowances: Array.isArray(object?.allowances) ? object.allowances.map((e) => feegrant_1.Grant.fromJSON(e)) : [],
            pagination: isSet(object.pagination) ? pagination_1.PageResponse.fromJSON(object.pagination) : undefined,
        };
    },
    toJSON(message) {
        const obj = {};
        if (message.allowances) {
            obj.allowances = message.allowances.map((e) => e ? feegrant_1.Grant.toJSON(e) : undefined);
        }
        else {
            obj.allowances = [];
        }
        message.pagination !== undefined
            && (obj.pagination = message.pagination ? pagination_1.PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseQueryAllowancesResponse();
        message.allowances = object.allowances?.map((e) => feegrant_1.Grant.fromPartial(e)) || [];
        message.pagination = (object.pagination !== undefined && object.pagination !== null)
            ? pagination_1.PageResponse.fromPartial(object.pagination)
            : undefined;
        return message;
    },
};
function createBaseQueryAllowancesByGranterRequest() {
    return { granter: "", pagination: undefined };
}
exports.QueryAllowancesByGranterRequest = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        if (message.granter !== "") {
            writer.uint32(10).string(message.granter);
        }
        if (message.pagination !== undefined) {
            pagination_1.PageRequest.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseQueryAllowancesByGranterRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.granter = reader.string();
                    break;
                case 2:
                    message.pagination = pagination_1.PageRequest.decode(reader, reader.uint32());
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
            granter: isSet(object.granter) ? String(object.granter) : "",
            pagination: isSet(object.pagination) ? pagination_1.PageRequest.fromJSON(object.pagination) : undefined,
        };
    },
    toJSON(message) {
        const obj = {};
        message.granter !== undefined && (obj.granter = message.granter);
        message.pagination !== undefined
            && (obj.pagination = message.pagination ? pagination_1.PageRequest.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseQueryAllowancesByGranterRequest();
        message.granter = object.granter ?? "";
        message.pagination = (object.pagination !== undefined && object.pagination !== null)
            ? pagination_1.PageRequest.fromPartial(object.pagination)
            : undefined;
        return message;
    },
};
function createBaseQueryAllowancesByGranterResponse() {
    return { allowances: [], pagination: undefined };
}
exports.QueryAllowancesByGranterResponse = {
    encode(message, writer = minimal_1.default.Writer.create()) {
        for (const v of message.allowances) {
            feegrant_1.Grant.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            pagination_1.PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof minimal_1.default.Reader ? input : new minimal_1.default.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseQueryAllowancesByGranterResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.allowances.push(feegrant_1.Grant.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.pagination = pagination_1.PageResponse.decode(reader, reader.uint32());
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
            allowances: Array.isArray(object?.allowances) ? object.allowances.map((e) => feegrant_1.Grant.fromJSON(e)) : [],
            pagination: isSet(object.pagination) ? pagination_1.PageResponse.fromJSON(object.pagination) : undefined,
        };
    },
    toJSON(message) {
        const obj = {};
        if (message.allowances) {
            obj.allowances = message.allowances.map((e) => e ? feegrant_1.Grant.toJSON(e) : undefined);
        }
        else {
            obj.allowances = [];
        }
        message.pagination !== undefined
            && (obj.pagination = message.pagination ? pagination_1.PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = createBaseQueryAllowancesByGranterResponse();
        message.allowances = object.allowances?.map((e) => feegrant_1.Grant.fromPartial(e)) || [];
        message.pagination = (object.pagination !== undefined && object.pagination !== null)
            ? pagination_1.PageResponse.fromPartial(object.pagination)
            : undefined;
        return message;
    },
};
class QueryClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
        this.Allowance = this.Allowance.bind(this);
        this.Allowances = this.Allowances.bind(this);
        this.AllowancesByGranter = this.AllowancesByGranter.bind(this);
    }
    Allowance(request) {
        const data = exports.QueryAllowanceRequest.encode(request).finish();
        const promise = this.rpc.request("cosmos.feegrant.v1beta1.Query", "Allowance", data);
        return promise.then((data) => exports.QueryAllowanceResponse.decode(new minimal_1.default.Reader(data)));
    }
    Allowances(request) {
        const data = exports.QueryAllowancesRequest.encode(request).finish();
        const promise = this.rpc.request("cosmos.feegrant.v1beta1.Query", "Allowances", data);
        return promise.then((data) => exports.QueryAllowancesResponse.decode(new minimal_1.default.Reader(data)));
    }
    AllowancesByGranter(request) {
        const data = exports.QueryAllowancesByGranterRequest.encode(request).finish();
        const promise = this.rpc.request("cosmos.feegrant.v1beta1.Query", "AllowancesByGranter", data);
        return promise.then((data) => exports.QueryAllowancesByGranterResponse.decode(new minimal_1.default.Reader(data)));
    }
}
exports.QueryClientImpl = QueryClientImpl;
function isSet(value) {
    return value !== null && value !== undefined;
}
