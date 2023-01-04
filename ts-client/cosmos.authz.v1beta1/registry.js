"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.msgTypes = void 0;
const tx_1 = require("./types/cosmos/authz/v1beta1/tx");
const tx_2 = require("./types/cosmos/authz/v1beta1/tx");
const tx_3 = require("./types/cosmos/authz/v1beta1/tx");
const msgTypes = [
    ["/cosmos.authz.v1beta1.MsgRevoke", tx_1.MsgRevoke],
    ["/cosmos.authz.v1beta1.MsgExec", tx_2.MsgExec],
    ["/cosmos.authz.v1beta1.MsgGrant", tx_3.MsgGrant],
];
exports.msgTypes = msgTypes;
