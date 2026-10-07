package handler

import (
	"encoding/json"
	"net/http"

	"github.com/lucassresende/wallet-service/internal/dto"
	"github.com/lucassresende/wallet-service/internal/service"
)

type WalletHandler struct {
	service *service.WalletService
}

func NewWalletHandler(
	service *service.WalletService,
) *WalletHandler {
	return &WalletHandler{service}
}

func (h *WalletHandler) CreateWallet(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req dto.CreateWalletRequest

	json.NewDecoder(r.Body).Decode(&req)

	wallet, err :=
		h.service.CreateWallet(
			req.PlayerID,
			req.Currency,
		)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(wallet)
}
func (h *WalletHandler) GetWallet(
	w http.ResponseWriter,
	r *http.Request,
) {

	id := r.URL.Query().Get("id")

	wallet, err := h.service.GetWallet(id)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusNotFound,
		)

		return
	}

	json.NewEncoder(w).Encode(wallet)
}
func (h *WalletHandler) Bet(
	w http.ResponseWriter,
	r *http.Request,
) {

	type Request struct {
		WalletID string `json:"walletId"`
		Amount   int64  `json:"amount"`
	}

	var req Request

	json.NewDecoder(r.Body).Decode(&req)

	err :=
		h.service.Bet(
			req.WalletID,
			req.Amount,
		)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *WalletHandler) Win(
	w http.ResponseWriter,
	r *http.Request,
) {

	type Request struct {
		WalletID string `json:"walletId"`
		Amount   int64  `json:"amount"`
	}

	var req Request

	json.NewDecoder(r.Body).Decode(&req)

	err :=
		h.service.Win(
			req.WalletID,
			req.Amount,
		)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	w.WriteHeader(http.StatusOK)
}
