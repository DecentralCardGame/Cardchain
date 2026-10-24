package keeper_test

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/stretchr/testify/require"
)

func TestRunDeckRESTGateway(t *testing.T) {
	f := setupRunDeck(t)
	_, err := f.msg.SubmitRunDeck(f.ctx, f.submit("gateway-run"))
	require.NoError(t, err)
	mux := runtime.NewServeMux()
	require.NoError(t, types.RegisterRunDeckQueryHandlerServer(f.ctx, mux, f.k))
	for _, path := range []string{
		"/DecentralCardGame/cardchain/cardchain/run_decks/0",
		"/DecentralCardGame/cardchain/cardchain/owners/" + f.owner + "/run_decks?limit=1",
		"/DecentralCardGame/cardchain/cardchain/owners/" + f.owner + "/all_stars",
	} {
		req := httptest.NewRequest("GET", path, nil).WithContext(f.ctx)
		res := httptest.NewRecorder()
		mux.ServeHTTP(res, req)
		require.Equal(t, 200, res.Code, res.Body.String())
		require.True(t, json.Valid(res.Body.Bytes()))
	}
	req := httptest.NewRequest("GET", "/DecentralCardGame/cardchain/cardchain/run_decks/99", nil).WithContext(f.ctx)
	res := httptest.NewRecorder()
	mux.ServeHTTP(res, req)
	require.Equal(t, 404, res.Code, res.Body.String())
}
