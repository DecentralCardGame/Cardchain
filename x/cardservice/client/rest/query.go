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

// TODO change this to give a list of cards
func getCardsHandler(cliCtx context.CLIContext, storeName string, status bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqStatus string

		if status {
			vars := mux.Vars(r)
			reqStatus = vars[restName]
		} else {
			reqStatus = ""
		}
		fmt.Println(restName)

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/cards/%s", storeName, reqStatus), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func getVotableCardsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("storeName")
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
