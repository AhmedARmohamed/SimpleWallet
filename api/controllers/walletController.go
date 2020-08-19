package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"wallet/api/models"
	"wallet/api/responses"
)
	//
	func (a *App) createWallet(w http.ResponseWriter, r *http.Request) {
		var resp = map[string]interface{}{"status": "success", "message": "Wallet created successfuly"}

		user := r.Context().Value("userID").(float64)
		wallet := &models.Wallet{}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		err = json.Unmarshal(body, &wallet)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		wallet.Prepare()

		if vne, _ := wallet.GetWallet(a.DB); vne != nil {
			resp["status"] = "Failed"
			resp["message"] = "Wallet already registered, please choose another name"
			responses.JSON(w, http.StatusBadRequest, resp)
			return
		}

		wallet.UserID = uint(user)

		walletCreated, err := wallet.Save(a.DB)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		resp["wallet"] = walletCreated
		responses.JSON(w, http.StatusCreated, resp)
		return
	}

	func (a *App) depositMoney(w http.ResponseWriter, r *http.Request) {
		var resp = map[string]interface{}{"status":"success", "message": "Deposit Successful."}
	fmt.Println("fnknjnvsecknekvneanvkenkee")
		vars := mux.Vars(r)
		wallet := &models.Wallet{}
		user := r.Context().Value("userID").(float64)
		userID := uint(user)

		id, _ := strconv.Atoi(vars["id"])

		wallets, err := wallet.GetWalletById(id, a.DB)

		if wallets.UserID != userID {
			resp["status"] = "Failed"
			resp["message"] = "Unathorized deposit "
			responses.JSON(w, http.StatusUnauthorized, resp)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		walletDeposited := models.Wallet{}
		if err = json.Unmarshal(body, &walletDeposited); err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		responses.JSON(w, http.StatusOK, resp)
		return


}