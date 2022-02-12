import { StdFee } from "@cosmjs/launchpad";
import { Registry, OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgSubmitMatchReporterProposal } from "./types/cardchain/tx";
import { MsgChangeArtist } from "./types/cardchain/tx";
import { MsgVoteCard } from "./types/cardchain/tx";
import { MsgDonateToCard } from "./types/cardchain/tx";
import { MsgApointMatchReporter } from "./types/cardchain/tx";
import { MsgRemoveCardFromCollection } from "./types/cardchain/tx";
import { MsgSaveCardContent } from "./types/cardchain/tx";
import { MsgRemoveContributorFromCollection } from "./types/cardchain/tx";
import { MsgAddArtwork } from "./types/cardchain/tx";
import { MsgBuyCardScheme } from "./types/cardchain/tx";
import { MsgBuyCollection } from "./types/cardchain/tx";
import { MsgSubmitCopyrightProposal } from "./types/cardchain/tx";
import { MsgReportMatch } from "./types/cardchain/tx";
import { MsgFinalizeCollection } from "./types/cardchain/tx";
import { MsgAddContributorToCollection } from "./types/cardchain/tx";
import { MsgRegisterForCouncil } from "./types/cardchain/tx";
import { MsgCreateuser } from "./types/cardchain/tx";
import { MsgSubmitCollectionProposal } from "./types/cardchain/tx";
import { MsgTransferCard } from "./types/cardchain/tx";
import { MsgAddCardToCollection } from "./types/cardchain/tx";
import { MsgCreateCollection } from "./types/cardchain/tx";
export declare const MissingWalletError: Error;
export declare const registry: Registry;
interface TxClientOptions {
    addr: string;
}
interface SignAndBroadcastOptions {
    fee: StdFee;
    memo?: string;
}
declare const txClient: (wallet: OfflineSigner, { addr: addr }?: TxClientOptions) => Promise<{
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }?: SignAndBroadcastOptions) => any;
    msgSubmitMatchReporterProposal: (data: MsgSubmitMatchReporterProposal) => EncodeObject;
    msgChangeArtist: (data: MsgChangeArtist) => EncodeObject;
    msgVoteCard: (data: MsgVoteCard) => EncodeObject;
    msgDonateToCard: (data: MsgDonateToCard) => EncodeObject;
    msgApointMatchReporter: (data: MsgApointMatchReporter) => EncodeObject;
    msgRemoveCardFromCollection: (data: MsgRemoveCardFromCollection) => EncodeObject;
    msgSaveCardContent: (data: MsgSaveCardContent) => EncodeObject;
    msgRemoveContributorFromCollection: (data: MsgRemoveContributorFromCollection) => EncodeObject;
    msgAddArtwork: (data: MsgAddArtwork) => EncodeObject;
    msgBuyCardScheme: (data: MsgBuyCardScheme) => EncodeObject;
    msgBuyCollection: (data: MsgBuyCollection) => EncodeObject;
    msgSubmitCopyrightProposal: (data: MsgSubmitCopyrightProposal) => EncodeObject;
    msgReportMatch: (data: MsgReportMatch) => EncodeObject;
    msgFinalizeCollection: (data: MsgFinalizeCollection) => EncodeObject;
    msgAddContributorToCollection: (data: MsgAddContributorToCollection) => EncodeObject;
    msgRegisterForCouncil: (data: MsgRegisterForCouncil) => EncodeObject;
    msgCreateuser: (data: MsgCreateuser) => EncodeObject;
    msgSubmitCollectionProposal: (data: MsgSubmitCollectionProposal) => EncodeObject;
    msgTransferCard: (data: MsgTransferCard) => EncodeObject;
    msgAddCardToCollection: (data: MsgAddCardToCollection) => EncodeObject;
    msgCreateCollection: (data: MsgCreateCollection) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
