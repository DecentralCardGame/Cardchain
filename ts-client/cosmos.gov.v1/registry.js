"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.msgTypes = void 0;
const tx_1 = require("./types/cosmos/gov/v1/tx");
const tx_2 = require("./types/cosmos/gov/v1/tx");
const tx_3 = require("./types/cosmos/gov/v1/tx");
const tx_4 = require("./types/cosmos/gov/v1/tx");
const msgTypes = [
    ["/cosmos.gov.v1.MsgVote", tx_1.MsgVote],
    ["/cosmos.gov.v1.MsgSubmitProposal", tx_2.MsgSubmitProposal],
    ["/cosmos.gov.v1.MsgVoteWeighted", tx_3.MsgVoteWeighted],
    ["/cosmos.gov.v1.MsgDeposit", tx_4.MsgDeposit],
];
exports.msgTypes = msgTypes;
