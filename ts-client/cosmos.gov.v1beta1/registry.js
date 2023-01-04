"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.msgTypes = void 0;
const tx_1 = require("./types/cosmos/gov/v1beta1/tx");
const tx_2 = require("./types/cosmos/gov/v1beta1/tx");
const tx_3 = require("./types/cosmos/gov/v1beta1/tx");
const tx_4 = require("./types/cosmos/gov/v1beta1/tx");
const msgTypes = [
    ["/cosmos.gov.v1beta1.MsgDeposit", tx_1.MsgDeposit],
    ["/cosmos.gov.v1beta1.MsgSubmitProposal", tx_2.MsgSubmitProposal],
    ["/cosmos.gov.v1beta1.MsgVote", tx_3.MsgVote],
    ["/cosmos.gov.v1beta1.MsgVoteWeighted", tx_4.MsgVoteWeighted],
];
exports.msgTypes = msgTypes;
