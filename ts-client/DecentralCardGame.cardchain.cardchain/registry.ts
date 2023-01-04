import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCommitCouncilResponse } from "./types/cardchain/cardchain/tx";
import { MsgAddStoryToCollection } from "./types/cardchain/cardchain/tx";
import { MsgCreateuser } from "./types/cardchain/cardchain/tx";
import { MsgRemoveCardFromCollection } from "./types/cardchain/cardchain/tx";
import { MsgDonateToCard } from "./types/cardchain/cardchain/tx";
import { MsgConfirmMatch } from "./types/cardchain/cardchain/tx";
import { MsgCreateSellOffer } from "./types/cardchain/cardchain/tx";
import { MsgAddArtwork } from "./types/cardchain/cardchain/tx";
import { MsgCreateCollection } from "./types/cardchain/cardchain/tx";
import { MsgAddArtworkToCollection } from "./types/cardchain/cardchain/tx";
import { MsgAddContributorToCollection } from "./types/cardchain/cardchain/tx";
import { MsgBuyCollection } from "./types/cardchain/cardchain/tx";
import { MsgRemoveContributorFromCollection } from "./types/cardchain/cardchain/tx";
import { MsgCreateCouncil } from "./types/cardchain/cardchain/tx";
import { MsgSubmitCollectionProposal } from "./types/cardchain/cardchain/tx";
import { MsgRestartCouncil } from "./types/cardchain/cardchain/tx";
import { MsgTransferBoosterPack } from "./types/cardchain/cardchain/tx";
import { MsgSetUserBiography } from "./types/cardchain/cardchain/tx";
import { MsgSetCardRarity } from "./types/cardchain/cardchain/tx";
import { MsgBuyCardScheme } from "./types/cardchain/cardchain/tx";
import { MsgRemoveSellOffer } from "./types/cardchain/cardchain/tx";
import { MsgChangeArtist } from "./types/cardchain/cardchain/tx";
import { MsgOpenBoosterPack } from "./types/cardchain/cardchain/tx";
import { MsgReportMatch } from "./types/cardchain/cardchain/tx";
import { MsgSetUserWebsite } from "./types/cardchain/cardchain/tx";
import { MsgApointMatchReporter } from "./types/cardchain/cardchain/tx";
import { MsgSetProfileCard } from "./types/cardchain/cardchain/tx";
import { MsgAddCardToCollection } from "./types/cardchain/cardchain/tx";
import { MsgVoteCard } from "./types/cardchain/cardchain/tx";
import { MsgBuyCard } from "./types/cardchain/cardchain/tx";
import { MsgSaveCardContent } from "./types/cardchain/cardchain/tx";
import { MsgRewokeCouncilRegistration } from "./types/cardchain/cardchain/tx";
import { MsgRevealCouncilResponse } from "./types/cardchain/cardchain/tx";
import { MsgSubmitCopyrightProposal } from "./types/cardchain/cardchain/tx";
import { MsgTransferCard } from "./types/cardchain/cardchain/tx";
import { MsgFinalizeCollection } from "./types/cardchain/cardchain/tx";
import { MsgRegisterForCouncil } from "./types/cardchain/cardchain/tx";
import { MsgSubmitMatchReporterProposal } from "./types/cardchain/cardchain/tx";
import { MsgSetCollectionStoryWriter } from "./types/cardchain/cardchain/tx";
import { MsgSetCollectionArtist } from "./types/cardchain/cardchain/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/DecentralCardGame.cardchain.cardchain.MsgCommitCouncilResponse", MsgCommitCouncilResponse],
    ["/DecentralCardGame.cardchain.cardchain.MsgAddStoryToCollection", MsgAddStoryToCollection],
    ["/DecentralCardGame.cardchain.cardchain.MsgCreateuser", MsgCreateuser],
    ["/DecentralCardGame.cardchain.cardchain.MsgRemoveCardFromCollection", MsgRemoveCardFromCollection],
    ["/DecentralCardGame.cardchain.cardchain.MsgDonateToCard", MsgDonateToCard],
    ["/DecentralCardGame.cardchain.cardchain.MsgConfirmMatch", MsgConfirmMatch],
    ["/DecentralCardGame.cardchain.cardchain.MsgCreateSellOffer", MsgCreateSellOffer],
    ["/DecentralCardGame.cardchain.cardchain.MsgAddArtwork", MsgAddArtwork],
    ["/DecentralCardGame.cardchain.cardchain.MsgCreateCollection", MsgCreateCollection],
    ["/DecentralCardGame.cardchain.cardchain.MsgAddArtworkToCollection", MsgAddArtworkToCollection],
    ["/DecentralCardGame.cardchain.cardchain.MsgAddContributorToCollection", MsgAddContributorToCollection],
    ["/DecentralCardGame.cardchain.cardchain.MsgBuyCollection", MsgBuyCollection],
    ["/DecentralCardGame.cardchain.cardchain.MsgRemoveContributorFromCollection", MsgRemoveContributorFromCollection],
    ["/DecentralCardGame.cardchain.cardchain.MsgCreateCouncil", MsgCreateCouncil],
    ["/DecentralCardGame.cardchain.cardchain.MsgSubmitCollectionProposal", MsgSubmitCollectionProposal],
    ["/DecentralCardGame.cardchain.cardchain.MsgRestartCouncil", MsgRestartCouncil],
    ["/DecentralCardGame.cardchain.cardchain.MsgTransferBoosterPack", MsgTransferBoosterPack],
    ["/DecentralCardGame.cardchain.cardchain.MsgSetUserBiography", MsgSetUserBiography],
    ["/DecentralCardGame.cardchain.cardchain.MsgSetCardRarity", MsgSetCardRarity],
    ["/DecentralCardGame.cardchain.cardchain.MsgBuyCardScheme", MsgBuyCardScheme],
    ["/DecentralCardGame.cardchain.cardchain.MsgRemoveSellOffer", MsgRemoveSellOffer],
    ["/DecentralCardGame.cardchain.cardchain.MsgChangeArtist", MsgChangeArtist],
    ["/DecentralCardGame.cardchain.cardchain.MsgOpenBoosterPack", MsgOpenBoosterPack],
    ["/DecentralCardGame.cardchain.cardchain.MsgReportMatch", MsgReportMatch],
    ["/DecentralCardGame.cardchain.cardchain.MsgSetUserWebsite", MsgSetUserWebsite],
    ["/DecentralCardGame.cardchain.cardchain.MsgApointMatchReporter", MsgApointMatchReporter],
    ["/DecentralCardGame.cardchain.cardchain.MsgSetProfileCard", MsgSetProfileCard],
    ["/DecentralCardGame.cardchain.cardchain.MsgAddCardToCollection", MsgAddCardToCollection],
    ["/DecentralCardGame.cardchain.cardchain.MsgVoteCard", MsgVoteCard],
    ["/DecentralCardGame.cardchain.cardchain.MsgBuyCard", MsgBuyCard],
    ["/DecentralCardGame.cardchain.cardchain.MsgSaveCardContent", MsgSaveCardContent],
    ["/DecentralCardGame.cardchain.cardchain.MsgRewokeCouncilRegistration", MsgRewokeCouncilRegistration],
    ["/DecentralCardGame.cardchain.cardchain.MsgRevealCouncilResponse", MsgRevealCouncilResponse],
    ["/DecentralCardGame.cardchain.cardchain.MsgSubmitCopyrightProposal", MsgSubmitCopyrightProposal],
    ["/DecentralCardGame.cardchain.cardchain.MsgTransferCard", MsgTransferCard],
    ["/DecentralCardGame.cardchain.cardchain.MsgFinalizeCollection", MsgFinalizeCollection],
    ["/DecentralCardGame.cardchain.cardchain.MsgRegisterForCouncil", MsgRegisterForCouncil],
    ["/DecentralCardGame.cardchain.cardchain.MsgSubmitMatchReporterProposal", MsgSubmitMatchReporterProposal],
    ["/DecentralCardGame.cardchain.cardchain.MsgSetCollectionStoryWriter", MsgSetCollectionStoryWriter],
    ["/DecentralCardGame.cardchain.cardchain.MsgSetCollectionArtist", MsgSetCollectionArtist],
    
];

export { msgTypes }