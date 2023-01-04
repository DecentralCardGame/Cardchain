"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.IgniteClient = void 0;
const proto_signing_1 = require("@cosmjs/proto-signing");
const stargate_1 = require("@cosmjs/stargate");
const events_1 = require("events");
const defaultFee = {
    amount: [],
    gas: "200000",
};
class IgniteClient extends events_1.EventEmitter {
    constructor(env, signer) {
        super();
        this.registry = [];
        this.env = env;
        this.setMaxListeners(0);
        this.signer = signer;
        const classConstructor = this.constructor;
        classConstructor.plugins.forEach(plugin => {
            const pluginInstance = plugin(this);
            Object.assign(this, pluginInstance.module);
            if (this.registry) {
                this.registry = this.registry.concat(pluginInstance.registry);
            }
        });
    }
    static plugin(plugin) {
        const currentPlugins = this.plugins;
        class AugmentedClient extends this {
        }
        AugmentedClient.plugins = currentPlugins.concat(plugin);
        if (Array.isArray(plugin)) {
            return AugmentedClient;
        }
        return AugmentedClient;
    }
    async signAndBroadcast(msgs, fee, memo) {
        if (this.signer) {
            const { address } = (await this.signer.getAccounts())[0];
            const signingClient = await stargate_1.SigningStargateClient.connectWithSigner(this.env.rpcURL, this.signer, { registry: new proto_signing_1.Registry(this.registry), prefix: this.env.prefix });
            return await signingClient.signAndBroadcast(address, msgs, fee ? fee : defaultFee, memo);
        }
        else {
            throw new Error(" Signer is not present.");
        }
    }
    async useSigner(signer) {
        this.signer = signer;
        this.emit("signer-changed", this.signer);
    }
    async useKeplr(keplrChainInfo = {}) {
        // Using queryClients directly because BaseClient has no knowledge of the modules at this stage
        try {
            const queryClient = (await Promise.resolve().then(() => require("./cosmos.base.tendermint.v1beta1/module"))).queryClient;
            const stakingQueryClient = (await Promise.resolve().then(() => require("./cosmos.staking.v1beta1/module"))).queryClient;
            const bankQueryClient = (await Promise.resolve().then(() => require("./cosmos.bank.v1beta1/module")))
                .queryClient;
            const stakingqc = stakingQueryClient({ addr: this.env.apiURL });
            const qc = queryClient({ addr: this.env.apiURL });
            const node_info = await (await qc.serviceGetNodeInfo()).data;
            const chainId = node_info.default_node_info?.network ?? "";
            const chainName = chainId?.toUpperCase() + " Network";
            const staking = await (await stakingqc.queryParams()).data;
            const bankqc = bankQueryClient({ addr: this.env.apiURL });
            const tokens = await (await bankqc.queryTotalSupply()).data;
            const addrPrefix = this.env.prefix ?? "cosmos";
            const rpc = this.env.rpcURL;
            const rest = this.env.apiURL;
            let stakeCurrency = {
                coinDenom: staking.params?.bond_denom?.toUpperCase() ?? "",
                coinMinimalDenom: staking.params?.bond_denom ?? "",
                coinDecimals: 0,
            };
            let bip44 = {
                coinType: 118,
            };
            let bech32Config = {
                bech32PrefixAccAddr: addrPrefix,
                bech32PrefixAccPub: addrPrefix + "pub",
                bech32PrefixValAddr: addrPrefix + "valoper",
                bech32PrefixValPub: addrPrefix + "valoperpub",
                bech32PrefixConsAddr: addrPrefix + "valcons",
                bech32PrefixConsPub: addrPrefix + "valconspub",
            };
            let currencies = tokens.supply?.map((x) => {
                const y = {
                    coinDenom: x.denom?.toUpperCase() ?? "",
                    coinMinimalDenom: x.denom ?? "",
                    coinDecimals: 0,
                };
                return y;
            }) ?? [];
            let feeCurrencies = tokens.supply?.map((x) => {
                const y = {
                    coinDenom: x.denom?.toUpperCase() ?? "",
                    coinMinimalDenom: x.denom ?? "",
                    coinDecimals: 0,
                };
                return y;
            }) ?? [];
            let coinType = 118;
            if (chainId) {
                const suggestOptions = {
                    chainId,
                    chainName,
                    rpc,
                    rest,
                    stakeCurrency,
                    bip44,
                    bech32Config,
                    currencies,
                    feeCurrencies,
                    coinType,
                    ...keplrChainInfo,
                };
                await window.keplr.experimentalSuggestChain(suggestOptions);
                window.keplr.defaultOptions = {
                    sign: {
                        preferNoSetFee: true,
                        preferNoSetMemo: true,
                    },
                };
            }
            await window.keplr.enable(chainId);
            this.signer = window.keplr.getOfflineSigner(chainId);
            this.emit("signer-changed", this.signer);
        }
        catch (e) {
            throw new Error("Could not load tendermint, staking and bank modules. Please ensure your client loads them to use useKeplr()");
        }
    }
}
exports.IgniteClient = IgniteClient;
IgniteClient.plugins = [];
