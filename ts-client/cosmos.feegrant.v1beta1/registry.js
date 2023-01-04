"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.msgTypes = void 0;
const tx_1 = require("./types/cosmos/feegrant/v1beta1/tx");
const tx_2 = require("./types/cosmos/feegrant/v1beta1/tx");
const msgTypes = [
    ["/cosmos.feegrant.v1beta1.MsgRevokeAllowance", tx_1.MsgRevokeAllowance],
    ["/cosmos.feegrant.v1beta1.MsgGrantAllowance", tx_2.MsgGrantAllowance],
];
exports.msgTypes = msgTypes;
