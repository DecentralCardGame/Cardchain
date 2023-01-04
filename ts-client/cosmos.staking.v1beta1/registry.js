"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.msgTypes = void 0;
const tx_1 = require("./types/cosmos/staking/v1beta1/tx");
const tx_2 = require("./types/cosmos/staking/v1beta1/tx");
const tx_3 = require("./types/cosmos/staking/v1beta1/tx");
const tx_4 = require("./types/cosmos/staking/v1beta1/tx");
const tx_5 = require("./types/cosmos/staking/v1beta1/tx");
const tx_6 = require("./types/cosmos/staking/v1beta1/tx");
const msgTypes = [
    ["/cosmos.staking.v1beta1.MsgCancelUnbondingDelegation", tx_1.MsgCancelUnbondingDelegation],
    ["/cosmos.staking.v1beta1.MsgCreateValidator", tx_2.MsgCreateValidator],
    ["/cosmos.staking.v1beta1.MsgEditValidator", tx_3.MsgEditValidator],
    ["/cosmos.staking.v1beta1.MsgDelegate", tx_4.MsgDelegate],
    ["/cosmos.staking.v1beta1.MsgBeginRedelegate", tx_5.MsgBeginRedelegate],
    ["/cosmos.staking.v1beta1.MsgUndelegate", tx_6.MsgUndelegate],
];
exports.msgTypes = msgTypes;
