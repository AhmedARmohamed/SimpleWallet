package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"wallet/api/models"
	"wallet/api/responses"
)

func (a *App) createWallet(w http.ResponseWriter, r *http.Request) {
	//var resp = map[string]interface{}{"status": "success", "message": "Wallet successfully created"}

	//user := r.Context().Value("userID").(float64)
	wallet := &models.Wallet{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&wallet); err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	if err := wallet.CreateWallet(a.DB); err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusCreated, wallet)
}

func (a *App) depositMoney(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	var wa models.Wallet
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&wa); err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()
	wa.ID = id

	balance, err := wa.DepositMoney(a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	newBalance :=map[string]int{"balance": balance.Amount}
	responses.JSON(w, http.StatusCreated, newBalance)
}

func (a *App) withdrawMoney(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	var wa models.Wallet
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&wa); err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()
	wa.ID = id
	newBalance, err := wa.WithdrawMoney(a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	balance := map[string]int{"balance": newBalance.Amount}
	responses.JSON(w, http.StatusCreated, balance)

}

func (a *App) balanceInquiry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	wa := models.Wallet{ID: id}
	if err := wa.CheckBalance(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			responses.ERROR(w, http.StatusNotFound, err)
		default:
			responses.ERROR(w, http.StatusInternalServerError, err)
		}
		return
	}
	responses.JSON(w, http.StatusOK, wa)
}