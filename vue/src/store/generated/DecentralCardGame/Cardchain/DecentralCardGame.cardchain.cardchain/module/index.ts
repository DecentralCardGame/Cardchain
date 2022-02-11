// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgAddCardToCollection } from "./types/cardchain/tx";
import { MsgRemoveContributorFromCollection } from "./types/cardchain/tx";
import { MsgRemoveCardFromCollection } from "./types/cardchain/tx";
import { MsgSubmitCopyrightProposal } from "./types/cardchain/tx";
import { MsgSubmitMatchReporterProposal } from "./types/cardchain/tx";
import { MsgSaveCardContent } from "./types/cardchain/tx";
import { MsgCreateCollection } from "./types/cardchain/tx";
import { MsgDonateToCard } from "./types/cardchain/tx";
import { MsgReportMatch } from "./types/cardchain/tx";
import { MsgAddArtwork } from "./types/cardchain/tx";
import { MsgTransferCard } from "./types/cardchain/tx";
import { MsgAddContributorToCollection } from "./types/cardchain/tx";
import { MsgApointMatchReporter } from "./types/cardchain/tx";
import { MsgBuyCollection } from "./types/cardchain/tx";
import { MsgVoteCard } from "./types/cardchain/tx";
import { MsgChangeArtist } from "./types/cardchain/tx";
import { MsgRegisterForCouncil } from "./types/cardchain/tx";
import { MsgFinalizeCollection } from "./types/cardchain/tx";
import { MsgCreateuser } from "./types/cardchain/tx";
import { MsgBuyCardScheme } from "./types/cardchain/tx";


const types = [
  ["/DecentralCardGame.cardchain.cardchain.MsgAddCardToCollection", MsgAddCardToCollection],
  ["/DecentralCardGame.cardchain.cardchain.MsgRemoveContributorFromCollection", MsgRemoveContributorFromCollection],
  ["/DecentralCardGame.cardchain.cardchain.MsgRemoveCardFromCollection", MsgRemoveCardFromCollection],
  ["/DecentralCardGame.cardchain.cardchain.MsgSubmitCopyrightProposal", MsgSubmitCopyrightProposal],
  ["/DecentralCardGame.cardchain.cardchain.MsgSubmitMatchReporterProposal", MsgSubmitMatchReporterProposal],
  ["/DecentralCardGame.cardchain.cardchain.MsgSaveCardContent", MsgSaveCardContent],
  ["/DecentralCardGame.cardchain.cardchain.MsgCreateCollection", MsgCreateCollection],
  ["/DecentralCardGame.cardchain.cardchain.MsgDonateToCard", MsgDonateToCard],
  ["/DecentralCardGame.cardchain.cardchain.MsgReportMatch", MsgReportMatch],
  ["/DecentralCardGame.cardchain.cardchain.MsgAddArtwork", MsgAddArtwork],
  ["/DecentralCardGame.cardchain.cardchain.MsgTransferCard", MsgTransferCard],
  ["/DecentralCardGame.cardchain.cardchain.MsgAddContributorToCollection", MsgAddContributorToCollection],
  ["/DecentralCardGame.cardchain.cardchain.MsgApointMatchReporter", MsgApointMatchReporter],
  ["/DecentralCardGame.cardchain.cardchain.MsgBuyCollection", MsgBuyCollection],
  ["/DecentralCardGame.cardchain.cardchain.MsgVoteCard", MsgVoteCard],
  ["/DecentralCardGame.cardchain.cardchain.MsgChangeArtist", MsgChangeArtist],
  ["/DecentralCardGame.cardchain.cardchain.MsgRegisterForCouncil", MsgRegisterForCouncil],
  ["/DecentralCardGame.cardchain.cardchain.MsgFinalizeCollection", MsgFinalizeCollection],
  ["/DecentralCardGame.cardchain.cardchain.MsgCreateuser", MsgCreateuser],
  ["/DecentralCardGame.cardchain.cardchain.MsgBuyCardScheme", MsgBuyCardScheme],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgAddCardToCollection: (data: MsgAddCardToCollection): EncodeObject => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgAddCardToCollection", value: MsgAddCardToCollection.fromPartial( data ) }),
    msgRemoveContributorFromCollection: (data: MsgRemoveContributorFromCollection): EncodeObject => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgRemoveContributorFromCollection", value: MsgRemoveContributorFromCollection.fromPartial( data ) }),
    msgRemoveCardFromCollection: (data: MsgRemoveCardFromCollection): EncodeObject => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgRemoveCardFromCollection", value: MsgRemoveCardFromCollection.fromPartial( data ) }),
    msgSubmitCopyrightProposal: (data: MsgSubmitCopyrightProposal): EncodeObject => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgSubmitCopyrightProposal", value: MsgSubmitCopyrightProposal.fromPartial( data ) }),
    msgSubmitMatchReporterProposal: (data: MsgSubmitMatchReporterProposal): EncodeObject => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgSubmitMatchReporterProposal", value: MsgSubmitMatchReporterProposal.fromPartial( data ) }),
    msgSaveCardContent: (data: MsgSaveCardContent): EncodeObject => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgSaveCardContent", value: MsgSaveCardContent.fromPartial( data ) }),
    msgCreateCollection: (data: MsgCreateCollection): EncodeObject => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgCreateCollection", value: MsgCreateCollection.fromPartial( data ) }),
    msgDonateToCard: (data: MsgDonateToCard): EncodeObject => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgDonateToCard", value: MsgDonateToCard.fromPartial( data ) }),
    msgReportMatch: (data: MsgReportMatch): EncodeObject => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgReportMatch", value: MsgReportMatch.fromPartial( data ) }),
    msgAddArtwork: (data: MsgAddArtwork): EncodeObject => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgAddArtwork", value: MsgAddArtwork.fromPartial( data ) }),
    msgTransferCard: (data: MsgTransferCard): EncodeObject => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgTransferCard", value: MsgTransferCard.fromPartial( data ) }),
    msgAddContributorToCollection: (data: MsgAddContributorToCollection): EncodeObject => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgAddContributorToCollection", value: MsgAddContributorToCollection.fromPartial( data ) }),
    msgApointMatchReporter: (data: MsgApointMatchReporter): EncodeObject => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgApointMatchReporter", value: MsgApointMatchReporter.fromPartial( data ) }),
    msgBuyCollection: (data: MsgBuyCollection): EncodeObject => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgBuyCollection", value: MsgBuyCollection.fromPartial( data ) }),
    msgVoteCard: (data: MsgVoteCard): EncodeObject => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgVoteCard", value: MsgVoteCard.fromPartial( data ) }),
    msgChangeArtist: (data: MsgChangeArtist): EncodeObject => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgChangeArtist", value: MsgChangeArtist.fromPartial( data ) }),
    msgRegisterForCouncil: (data: MsgRegisterForCouncil): EncodeObject => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgRegisterForCouncil", value: MsgRegisterForCouncil.fromPartial( data ) }),
    msgFinalizeCollection: (data: MsgFinalizeCollection): EncodeObject => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgFinalizeCollection", value: MsgFinalizeCollection.fromPartial( data ) }),
    msgCreateuser: (data: MsgCreateuser): EncodeObject => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgCreateuser", value: MsgCreateuser.fromPartial( data ) }),
    msgBuyCardScheme: (data: MsgBuyCardScheme): EncodeObject => ({ typeUrl: "/DecentralCardGame.cardchain.cardchain.MsgBuyCardScheme", value: MsgBuyCardScheme.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
