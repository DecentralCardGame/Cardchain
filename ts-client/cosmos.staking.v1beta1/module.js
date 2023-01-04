"use strict";
// Generated by Ignite ignite.com/cli
Object.defineProperty(exports, "__esModule", { value: true });
exports.queryClient = exports.txClient = exports.registry = exports.MsgUndelegate = exports.MsgBeginRedelegate = exports.MsgDelegate = exports.MsgEditValidator = exports.MsgCreateValidator = exports.MsgCancelUnbondingDelegation = void 0;
const stargate_1 = require("@cosmjs/stargate");
const proto_signing_1 = require("@cosmjs/proto-signing");
const registry_1 = require("./registry");
const rest_1 = require("./rest");
const tx_1 = require("./types/cosmos/staking/v1beta1/tx");
Object.defineProperty(exports, "MsgCancelUnbondingDelegation", { enumerable: true, get: function () { return tx_1.MsgCancelUnbondingDelegation; } });
const tx_2 = require("./types/cosmos/staking/v1beta1/tx");
Object.defineProperty(exports, "MsgCreateValidator", { enumerable: true, get: function () { return tx_2.MsgCreateValidator; } });
const tx_3 = require("./types/cosmos/staking/v1beta1/tx");
Object.defineProperty(exports, "MsgEditValidator", { enumerable: true, get: function () { return tx_3.MsgEditValidator; } });
const tx_4 = require("./types/cosmos/staking/v1beta1/tx");
Object.defineProperty(exports, "MsgDelegate", { enumerable: true, get: function () { return tx_4.MsgDelegate; } });
const tx_5 = require("./types/cosmos/staking/v1beta1/tx");
Object.defineProperty(exports, "MsgBeginRedelegate", { enumerable: true, get: function () { return tx_5.MsgBeginRedelegate; } });
const tx_6 = require("./types/cosmos/staking/v1beta1/tx");
Object.defineProperty(exports, "MsgUndelegate", { enumerable: true, get: function () { return tx_6.MsgUndelegate; } });
exports.registry = new proto_signing_1.Registry(registry_1.msgTypes);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = ({ signer, prefix, addr } = { addr: "http://localhost:26657", prefix: "cosmos" }) => {
    return {
        async sendMsgCancelUnbondingDelegation({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendMsgCancelUnbondingDelegation: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await stargate_1.SigningStargateClient.connectWithSigner(addr, signer, { registry: exports.registry, prefix });
                let msg = this.msgCancelUnbondingDelegation({ value: tx_1.MsgCancelUnbondingDelegation.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendMsgCancelUnbondingDelegation: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendMsgCreateValidator({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendMsgCreateValidator: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await stargate_1.SigningStargateClient.connectWithSigner(addr, signer, { registry: exports.registry, prefix });
                let msg = this.msgCreateValidator({ value: tx_2.MsgCreateValidator.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendMsgCreateValidator: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendMsgEditValidator({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendMsgEditValidator: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await stargate_1.SigningStargateClient.connectWithSigner(addr, signer, { registry: exports.registry, prefix });
                let msg = this.msgEditValidator({ value: tx_3.MsgEditValidator.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendMsgEditValidator: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendMsgDelegate({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendMsgDelegate: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await stargate_1.SigningStargateClient.connectWithSigner(addr, signer, { registry: exports.registry, prefix });
                let msg = this.msgDelegate({ value: tx_4.MsgDelegate.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendMsgDelegate: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendMsgBeginRedelegate({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendMsgBeginRedelegate: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await stargate_1.SigningStargateClient.connectWithSigner(addr, signer, { registry: exports.registry, prefix });
                let msg = this.msgBeginRedelegate({ value: tx_5.MsgBeginRedelegate.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendMsgBeginRedelegate: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendMsgUndelegate({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendMsgUndelegate: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await stargate_1.SigningStargateClient.connectWithSigner(addr, signer, { registry: exports.registry, prefix });
                let msg = this.msgUndelegate({ value: tx_6.MsgUndelegate.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendMsgUndelegate: Could not broadcast Tx: ' + e.message);
            }
        },
        msgCancelUnbondingDelegation({ value }) {
            try {
                return { typeUrl: "/cosmos.staking.v1beta1.MsgCancelUnbondingDelegation", value: tx_1.MsgCancelUnbondingDelegation.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:MsgCancelUnbondingDelegation: Could not create message: ' + e.message);
            }
        },
        msgCreateValidator({ value }) {
            try {
                return { typeUrl: "/cosmos.staking.v1beta1.MsgCreateValidator", value: tx_2.MsgCreateValidator.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:MsgCreateValidator: Could not create message: ' + e.message);
            }
        },
        msgEditValidator({ value }) {
            try {
                return { typeUrl: "/cosmos.staking.v1beta1.MsgEditValidator", value: tx_3.MsgEditValidator.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:MsgEditValidator: Could not create message: ' + e.message);
            }
        },
        msgDelegate({ value }) {
            try {
                return { typeUrl: "/cosmos.staking.v1beta1.MsgDelegate", value: tx_4.MsgDelegate.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:MsgDelegate: Could not create message: ' + e.message);
            }
        },
        msgBeginRedelegate({ value }) {
            try {
                return { typeUrl: "/cosmos.staking.v1beta1.MsgBeginRedelegate", value: tx_5.MsgBeginRedelegate.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:MsgBeginRedelegate: Could not create message: ' + e.message);
            }
        },
        msgUndelegate({ value }) {
            try {
                return { typeUrl: "/cosmos.staking.v1beta1.MsgUndelegate", value: tx_6.MsgUndelegate.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:MsgUndelegate: Could not create message: ' + e.message);
            }
        },
    };
};
exports.txClient = txClient;
const queryClient = ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new rest_1.Api({ baseURL: addr });
};
exports.queryClient = queryClient;
class SDKModule {
    constructor(client) {
        this.registry = [];
        this.query = (0, exports.queryClient)({ addr: client.env.apiURL });
        this.updateTX(client);
        client.on('signer-changed', (signer) => {
            this.updateTX(client);
        });
    }
    updateTX(client) {
        const methods = (0, exports.txClient)({
            signer: client.signer,
            addr: client.env.rpcURL,
            prefix: client.env.prefix ?? "cosmos",
        });
        this.tx = methods;
        for (let m in methods) {
            this.tx[m] = methods[m].bind(this.tx);
        }
    }
}
;
const Module = (test) => {
    return {
        module: {
            CosmosStakingV1Beta1: new SDKModule(test)
        },
        registry: registry_1.msgTypes
    };
};
exports.default = Module;
