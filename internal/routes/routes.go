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
		func(w http.ResponseWriter, r *http.Request) {

			switch r.Method {

			case http.MethodPost:
				handler.CreateWallet(w, r)

			case http.MethodGet:
				handler.GetWallet(w, r)

			default:
				http.Error(
					w,
					"method not allowed",
					http.StatusMethodNotAllowed,
				)
			}
		},
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