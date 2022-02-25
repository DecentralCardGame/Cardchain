package cli

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCommitCouncilResponse() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "commit-council-response [councilId] [response] [secret]",
		Short: "Broadcast message CommitCouncilResponse",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCouncilId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			response := types.Response_name[types.Response_value[args[1]]]
			clearResponse := response + args[2]

			hashResponse := sha256.Sum256([]byte(clearResponse))
			hashStringResponse := hex.EncodeToString(hashResponse[:])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCommitCouncilResponse(
				clientCtx.GetFromAddress().String(),
				hashStringResponse,
				argCouncilId,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
