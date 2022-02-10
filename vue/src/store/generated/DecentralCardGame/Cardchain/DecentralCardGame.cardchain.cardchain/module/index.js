// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgDonateToCard } from "./types/cardchain/tx";
import { MsgChangeArtist } from "./types/cardchain/tx";
import { MsgCreateuser } from "./types/cardchain/tx";
import { MsgBuyCollection } from "./types/cardchain/tx";
import { MsgAddCardToCollection } from "./types/cardchain/tx";
import { MsgFinalizeCollection } from "./types/cardchain/tx";
import { MsgRegisterForCouncil } from "./types/cardchain/tx";
import { MsgReportMatch } from "./types/cardchain/tx";
import { MsgVoteCard } from "./types/cardchain/tx";
import { MsgSaveCardContent } from "./types/cardchain/tx";
import { MsgAddArtwork } from "./types/cardchain/tx";
import { MsgCreateCollection } from "./types/cardchain/tx";
import { MsgBuyCardScheme } from "./types/cardchain/tx";
import { MsgApointMatchReporter } from "./types/cardchain/tx";
import { MsgTransferCard } from "./types/cardchain/tx";
import { MsgSubmitCopyrightProposal } from "./types/cardchain/tx";
import { MsgSubmitMatchReporterProposal } from "./types/cardchain/tx";
const types = [
    ["/DecentralCardGame.cardchain.cardchain.MsgDonateToCard", MsgDonateToCard],
    ["/DecentralCardGame.cardchain.cardchain.MsgChangeArtist", MsgChangeArtist],
    ["/DecentralCardGame.cardchain.cardchain.MsgCreateuser", MsgCreateuser],
    ["/DecentralCardGame.cardchain.cardchain.MsgBuyCollection", MsgBuyCollection],
    ["/DecentralCardGame.cardchain.cardchain.MsgAddCardToCollection", MsgAddCardToCollection],
    ["/DecentralCardGame.cardchain.cardchain.MsgFinalizeCollection", MsgFinalizeCollection],
    ["/DecentralCardGame.cardchain.cardchain.MsgRegisterForCouncil", MsgRegisterForCouncil],
    ["/DecentralCardGame.cardchain.cardchain.MsgReportMatch", MsgReportMatch],
    ["/DecentralCardGame.cardchain.cardchain.MsgVoteCard", MsgVoteCard],
    ["/DecentralCardGame.cardchain.cardchain.MsgSaveCardContent", MsgSaveCardContent],
    ["/DecentralCardGame.cardchain.cardchain.MsgAddArtwork", MsgAddArtwork],
    ["/DecentralCardGame.cardchain.cardchain.MsgCreateCollection", MsgCreateCollection],
    ["/DecentralCardGame.cardchain.cardchain.MsgBuyCardScheme", MsgBuyCardScheme],
    ["/DecentralCardGame.cardchain.cardchain.MsgApointMatchReporter", MsgApointMatchReporter],
    ["/DecentralCardGame.cardchain.cardchain.MsgTransferCard", MsgTransferCard],
    ["/DecentralCardGame.cardchain.cardchain.MsgSubmitCopyrightProposal", MsgSubmitCopyrightProposal],
    ["/DecentralCardGame.cardchain.cardchain.MsgSubmitMatchReporterProposal", MsgSubmitMatchReporterProposal],
];
export const MissingWalletError = new Error("wallet is required");
export const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    let client;
    if (addr) {
        client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    }
    else {
        client = await SigningStargateClient.offline(wallet, { registry });
    }
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgDonateToCard: (data) => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgDonateToCard", value: MsgDonateToCard.fromPartial(data) }),
        msgChangeArtist: (data) => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgChangeArtist", value: MsgChangeArtist.fromPartial(data) }),
        msgCreateuser: (data) => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgCreateuser", value: MsgCreateuser.fromPartial(data) }),
        msgBuyCollection: (data) => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgBuyCollection", value: MsgBuyCollection.fromPartial(data) }),
        msgAddCardToCollection: (data) => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgAddCardToCollection", value: MsgAddCardToCollection.fromPartial(data) }),
        msgFinalizeCollection: (data) => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgFinalizeCollection", value: MsgFinalizeCollection.fromPartial(data) }),
        msgRegisterForCouncil: (data) => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgRegisterForCouncil", value: MsgRegisterForCouncil.fromPartial(data) }),
        msgReportMatch: (data) => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgReportMatch", value: MsgReportMatch.fromPartial(data) }),
        msgVoteCard: (data) => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgVoteCard", value: MsgVoteCard.fromPartial(data) }),
        msgSaveCardContent: (data) => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgSaveCardContent", value: MsgSaveCardContent.fromPartial(data) }),
        msgAddArtwork: (data) => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgAddArtwork", value: MsgAddArtwork.fromPartial(data) }),
        msgCreateCollection: (data) => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgCreateCollection", value: MsgCreateCollection.fromPartial(data) }),
        msgBuyCardScheme: (data) => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgBuyCardScheme", value: MsgBuyCardScheme.fromPartial(data) }),
        msgApointMatchReporter: (data) => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgApointMatchReporter", value: MsgApointMatchReporter.fromPartial(data) }),
        msgTransferCard: (data) => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgTransferCard", value: MsgTransferCard.fromPartial(data) }),
        msgSubmitCopyrightProposal: (data) => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgSubmitCopyrightProposal", value: MsgSubmitCopyrightProposal.fromPartial(data) }),
        msgSubmitMatchReporterProposal: (data) => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgSubmitMatchReporterProposal", value: MsgSubmitMatchReporterProposal.fromPartial(data) }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
