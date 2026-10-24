package types

import (
	"bytes"
	"testing"

	"cosmossdk.io/x/tx/signing"
	"github.com/cosmos/cosmos-sdk/codec"
	addresscodec "github.com/cosmos/cosmos-sdk/codec/address"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/require"
)

func TestRunDeckSDKSigners(t *testing.T) {
	prefix := sdk.GetConfig().GetBech32AccountAddrPrefix()
	reg, err := codectypes.NewInterfaceRegistryWithOptions(codectypes.InterfaceRegistryOptions{
		ProtoFiles:     proto.HybridResolver,
		SigningOptions: signing.Options{AddressCodec: addresscodec.NewBech32Codec(prefix), ValidatorAddressCodec: addresscodec.NewBech32Codec(prefix + "valoper")},
	})
	require.NoError(t, err)
	RegisterInterfaces(reg)
	cdc := codec.NewProtoCodec(reg)
	owner := sdk.AccAddress(bytes.Repeat([]byte{1}, 20))
	reporter := sdk.AccAddress(bytes.Repeat([]byte{2}, 20))
	signers, _, err := cdc.GetMsgV1Signers(&MsgSubmitRunDeck{Owner: owner.String(), Reporter: reporter.String()})
	require.NoError(t, err)
	require.Equal(t, [][]byte{reporter.Bytes()}, signers)
	signers, _, err = cdc.GetMsgV1Signers(&MsgDistillRunDeck{Owner: owner.String()})
	require.NoError(t, err)
	require.Equal(t, [][]byte{owner.Bytes()}, signers)
}
