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
	// Group gov queries under a subcommand
	govQueryCmd := &cobra.Command{
		Use:   "cardservice",
		Short: "Querying commands for the cardservice module",
	}

	govQueryCmd.AddCommand(client.GetCommands(
		cardservicecmd.GetCmdResolveName(mc.storeKey, mc.cdc),
		cardservicecmd.GetCmdWhois(mc.storeKey, mc.cdc),
	)...)

	return govQueryCmd
}

// GetTxCmd returns the transaction commands for this module
func (mc ModuleClient) GetTxCmd() *cobra.Command {
	govTxCmd := &cobra.Command{
		Use:   "cardservice",
		Short: "cardservice transactions subcommands",
	}

	govTxCmd.AddCommand(client.PostCommands(
		cardservicecmd.GetCmdBuyCardScheme(mc.cdc),
		//cardservicecmd.GetCmdSetName(mc.cdc),
		cardservicecmd.GetCmdSaveCardContent(mc.cdc),
		cardservicecmd.GetCmdVoteCard(mc.cdc),
		cardservicecmd.GetCmdTransferCard(mc.cdc),
		cardservicecmd.GetCmdDonateToCard(mc.cdc),
		cardservicecmd.GetCmdCreateUser(mc.cdc),
	)...)

	return govTxCmd
}
