package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/gorilla/mux"
)

func getCardHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		paramType := vars[restName]

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/card/%s", storeName, paramType), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

type getCardsReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Amount  string       `json:"amount"`
	Buyer   string       `json:"buyer"`
}

func getCardsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		query := r.URL.Query()

		status := ""
		if len(query["status"]) > 0 {
			status = query["status"][0]
		}
		owner := ""
		if len(query["owner"]) > 0 {
			owner = query["owner"][0]
		}
		nameContains := ""
		if len(query["nameContains"]) > 0 {
			nameContains = query["nameContains"][0]
		}

		fmt.Println(status, owner, nameContains)


		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/cards/%s/%s/%s", storeName, owner, status, nameContains), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func getVotableCardsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(storeName)
		vars := mux.Vars(r)
		name := vars[restName]
		fmt.Println(name)

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/votable-cards/%s", storeName, name), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func getUserHandler(cliCtx context.CLIContext, userAddress string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		paramType := vars[restName]

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/user/%s", userAddress, paramType), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func getCardchainInfoHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/cardchain-info", storeName), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}
