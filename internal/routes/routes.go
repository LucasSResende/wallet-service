package routes

import (
	"net/http"

	"github.com/lucassresende/wallet-service/internal/handler"
)

func RegisterRoutes(
	handler *handler.WalletHandler,
) {

	http.HandleFunc(
		"/wallets",
		handler.CreateWallet,
	)

	http.HandleFunc(
	"/wallets/bet",
	handler.Bet,
	)

	http.HandleFunc(
	"/wallets/win",
	handler.Win,
	)



}