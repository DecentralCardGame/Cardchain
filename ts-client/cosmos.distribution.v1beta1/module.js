"use strict";
// Generated by Ignite ignite.com/cli
Object.defineProperty(exports, "__esModule", { value: true });
exports.queryClient = exports.txClient = exports.registry = exports.MsgWithdrawDelegatorReward = exports.MsgFundCommunityPool = exports.MsgWithdrawValidatorCommission = exports.MsgSetWithdrawAddress = void 0;
const stargate_1 = require("@cosmjs/stargate");
const proto_signing_1 = require("@cosmjs/proto-signing");
const registry_1 = require("./registry");
const rest_1 = require("./rest");
const tx_1 = require("./types/cosmos/distribution/v1beta1/tx");
Object.defineProperty(exports, "MsgSetWithdrawAddress", { enumerable: true, get: function () { return tx_1.MsgSetWithdrawAddress; } });
const tx_2 = require("./types/cosmos/distribution/v1beta1/tx");
Object.defineProperty(exports, "MsgWithdrawValidatorCommission", { enumerable: true, get: function () { return tx_2.MsgWithdrawValidatorCommission; } });
const tx_3 = require("./types/cosmos/distribution/v1beta1/tx");
Object.defineProperty(exports, "MsgFundCommunityPool", { enumerable: true, get: function () { return tx_3.MsgFundCommunityPool; } });
const tx_4 = require("./types/cosmos/distribution/v1beta1/tx");
Object.defineProperty(exports, "MsgWithdrawDelegatorReward", { enumerable: true, get: function () { return tx_4.MsgWithdrawDelegatorReward; } });
exports.registry = new proto_signing_1.Registry(registry_1.msgTypes);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = ({ signer, prefix, addr } = { addr: "http://localhost:26657", prefix: "cosmos" }) => {
    return {
        async sendMsgSetWithdrawAddress({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendMsgSetWithdrawAddress: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await stargate_1.SigningStargateClient.connectWithSigner(addr, signer, { registry: exports.registry, prefix });
                let msg = this.msgSetWithdrawAddress({ value: tx_1.MsgSetWithdrawAddress.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendMsgSetWithdrawAddress: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendMsgWithdrawValidatorCommission({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendMsgWithdrawValidatorCommission: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await stargate_1.SigningStargateClient.connectWithSigner(addr, signer, { registry: exports.registry, prefix });
                let msg = this.msgWithdrawValidatorCommission({ value: tx_2.MsgWithdrawValidatorCommission.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendMsgWithdrawValidatorCommission: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendMsgFundCommunityPool({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendMsgFundCommunityPool: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await stargate_1.SigningStargateClient.connectWithSigner(addr, signer, { registry: exports.registry, prefix });
                let msg = this.msgFundCommunityPool({ value: tx_3.MsgFundCommunityPool.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendMsgFundCommunityPool: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendMsgWithdrawDelegatorReward({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendMsgWithdrawDelegatorReward: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await stargate_1.SigningStargateClient.connectWithSigner(addr, signer, { registry: exports.registry, prefix });
                let msg = this.msgWithdrawDelegatorReward({ value: tx_4.MsgWithdrawDelegatorReward.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendMsgWithdrawDelegatorReward: Could not broadcast Tx: ' + e.message);
            }
        },
        msgSetWithdrawAddress({ value }) {
            try {
                return { typeUrl: "/cosmos.distribution.v1beta1.MsgSetWithdrawAddress", value: tx_1.MsgSetWithdrawAddress.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:MsgSetWithdrawAddress: Could not create message: ' + e.message);
            }
        },
        msgWithdrawValidatorCommission({ value }) {
            try {
                return { typeUrl: "/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission", value: tx_2.MsgWithdrawValidatorCommission.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:MsgWithdrawValidatorCommission: Could not create message: ' + e.message);
            }
        },
        msgFundCommunityPool({ value }) {
            try {
                return { typeUrl: "/cosmos.distribution.v1beta1.MsgFundCommunityPool", value: tx_3.MsgFundCommunityPool.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:MsgFundCommunityPool: Could not create message: ' + e.message);
            }
        },
        msgWithdrawDelegatorReward({ value }) {
            try {
                return { typeUrl: "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward", value: tx_4.MsgWithdrawDelegatorReward.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:MsgWithdrawDelegatorReward: Could not create message: ' + e.message);
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
            CosmosDistributionV1Beta1: new SDKModule(test)
        },
        registry: registry_1.msgTypes
    };
};
exports.default = Module;
