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

// this was mainly built to get cards on untap.in but this might have no use at all... TODO
func getCardSVGHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		paramType := vars[restName]

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/cardsvg/%s", storeName, paramType), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		stringres := string("<?xml version=\"1.0\" encoding=\"ISO-8859-1\" standalone=\"no\"?> <!DOCTYPE svg PUBLIC \"-//W3C//DTD SVG 1.1//EN\" \"http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd\"> <svg mtc:dpi=\"90\" xmlns=\"http://www.w3.org/2000/svg\" version=\"1.1\" width=\"29.0882in\" height=\"22in\" viewBox=\"0 0 2617.942383 1980\" xmlns:inkscape=\"http://www.inkscape.org/namespaces/inkscape\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xmlns:mtc=\"http://www.make-the-cut.com/namespaces/make-the-cut\"> <circle cx=\"50\" cy=\"50\" r=\"40\" stroke=\"black\" stroke-width=\"3\" fill=\"red\" /> </svg> ")
		res = []byte(stringres)

		w.Header().Set("Content-Type", "image/svg+xml")
		_, _ = w.Write(res)
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
