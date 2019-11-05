package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	cardservicecmd "github.com/DecentralCardGame/Cardchain/x/cardservice/client/cli"
	"github.com/spf13/cobra"
	amino "github.com/tendermint/go-amino"
)

// ModuleClient exports all client functionality from this module
type ModuleClient struct {
	storeKey string
	cdc      *amino.Codec
}

func NewModuleClient(storeKey string, cdc *amino.Codec) ModuleClient {
	return ModuleClient{storeKey, cdc}
}

// GetQueryCmd returns the cli query commands for this module
func (mc ModuleClient) GetQueryCmd() *cobra.Command {
	// Group cardservice queries under a subcommand
	cardserviceQueryCmd := &cobra.Command{
		Use:   "cardservice",
		Short: "Querying commands for the cardservice module",
	}

	cardserviceQueryCmd.AddCommand(client.GetCommands(
		cardservicecmd.GetCmdCard(mc.storeKey, mc.cdc),
		cardservicecmd.GetCmdUser(mc.storeKey, mc.cdc),
		cardservicecmd.GetCmdCardList(mc.storeKey, mc.cdc),
		cardservicecmd.GetCmdVotableCardList(mc.storeKey, mc.cdc),
		cardservicecmd.GetCmdCardchainInfo(mc.storeKey, mc.cdc),
	)...)

	return cardserviceQueryCmd
}

// GetTxCmd returns the transaction commands for this module
func (mc ModuleClient) GetTxCmd() *cobra.Command {
	cardserviceTxCmd := &cobra.Command{
		Use:   "cardservice",
		Short: "cardservice transactions subcommands",
	}

	cardserviceTxCmd.AddCommand(client.PostCommands(
		cardservicecmd.GetCmdBuyCardScheme(mc.cdc),
		cardservicecmd.GetCmdSaveCardContent(mc.cdc),
		cardservicecmd.GetCmdVoteCard(mc.cdc),
		cardservicecmd.GetCmdTransferCard(mc.cdc),
		cardservicecmd.GetCmdDonateToCard(mc.cdc),
		cardservicecmd.GetCmdCreateUser(mc.cdc),
	)...)

	return cardserviceTxCmd
}
