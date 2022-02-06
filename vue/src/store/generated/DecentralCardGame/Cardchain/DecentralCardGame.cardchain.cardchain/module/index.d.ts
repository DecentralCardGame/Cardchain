import { StdFee } from "@cosmjs/launchpad";
import { Registry, OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgRegisterForCouncil } from "./types/cardchain/tx";
import { MsgCreateuser } from "./types/cardchain/tx";
import { MsgTransferCard } from "./types/cardchain/tx";
import { MsgChangeArtist } from "./types/cardchain/tx";
import { MsgReportMatch } from "./types/cardchain/tx";
import { MsgSaveCardContent } from "./types/cardchain/tx";
import { MsgVoteCard } from "./types/cardchain/tx";
import { MsgBuyCardScheme } from "./types/cardchain/tx";
import { MsgAddArtwork } from "./types/cardchain/tx";
import { MsgSubmitCopyrightProposal } from "./types/cardchain/tx";
import { MsgDonateToCard } from "./types/cardchain/tx";
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
    msgRegisterForCouncil: (data: MsgRegisterForCouncil) => EncodeObject;
    msgCreateuser: (data: MsgCreateuser) => EncodeObject;
    msgTransferCard: (data: MsgTransferCard) => EncodeObject;
    msgChangeArtist: (data: MsgChangeArtist) => EncodeObject;
    msgReportMatch: (data: MsgReportMatch) => EncodeObject;
    msgSaveCardContent: (data: MsgSaveCardContent) => EncodeObject;
    msgVoteCard: (data: MsgVoteCard) => EncodeObject;
    msgBuyCardScheme: (data: MsgBuyCardScheme) => EncodeObject;
    msgAddArtwork: (data: MsgAddArtwork) => EncodeObject;
    msgSubmitCopyrightProposal: (data: MsgSubmitCopyrightProposal) => EncodeObject;
    msgDonateToCard: (data: MsgDonateToCard) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
