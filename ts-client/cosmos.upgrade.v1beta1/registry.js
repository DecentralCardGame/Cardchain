"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.msgTypes = void 0;
const tx_1 = require("./types/cosmos/upgrade/v1beta1/tx");
const tx_2 = require("./types/cosmos/upgrade/v1beta1/tx");
const msgTypes = [
    ["/cosmos.upgrade.v1beta1.MsgSoftwareUpgrade", tx_1.MsgSoftwareUpgrade],
    ["/cosmos.upgrade.v1beta1.MsgCancelUpgrade", tx_2.MsgCancelUpgrade],
];
exports.msgTypes = msgTypes;
