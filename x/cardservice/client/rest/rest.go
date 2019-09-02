package rest

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/gorilla/mux"
)

const (
	restName = "card"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	r.HandleFunc(fmt.Sprintf("/%s/buy_card_scheme", storeName), buyCardSchemeHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/save_card_content", storeName), saveCardContentHandler(cliCtx)).Methods("PUT")
	r.HandleFunc(fmt.Sprintf("/%s/vote_card", storeName), voteCardHandler(cliCtx)).Methods("PUT")
	r.HandleFunc(fmt.Sprintf("/%s/transfer_card", storeName), transferCardHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/donate_to_card", storeName), donateToCardHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/create_user", storeName), createUserHandler(cliCtx)).Methods("PUT")

	r.HandleFunc(fmt.Sprintf("/%s/card/{%s}", storeName, restName), resolveCardHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/cards", storeName), resolveCardsHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/whois/{%s}", storeName, restName), whoIsHandler(cliCtx, storeName)).Methods("GET")
}
