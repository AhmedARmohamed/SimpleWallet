package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"wallet/api/models"
	"wallet/api/responses"
	"wallet/utils"
)

// UserSignUp controller for creating new users
func (a *App) userSignUp(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": "success", "message": "Registered successfully"}

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}
	user.Prepare() // here strip the text of white spaces


	err = user.Validate("") // default were all fields(email, lastname, firstname, password, profileimage) are validated
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	//log.Printf("%#v",user)

	usr, err := user.GetUser(a.DB)
	if err != nil {
		log.Printf("user %v", usr)
	}
	if usr != nil {
		resp["status"] = "failed"
		resp["message"] = "User already registered, please login"
		responses.JSON(w, http.StatusBadRequest, resp)
		return
	}

	userCreated, err := user.SaveUser(a.DB)
	log.Printf("created user %v", userCreated)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	resp["user"] = userCreated
	responses.JSON(w, http.StatusCreated, resp)
	return
}

//Login: signs in users
func (a *App) login(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": "success", "message": "logged in"}

	user := &models.User{}
	body, err := ioutil.ReadAll(r.Body) //read user input from request
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user.Prepare() //to strip the text of white spaces

	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}


	usr, err := user.GetUser(a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if usr == nil { // user is not registered
		resp["status"] = "failed"
		resp["message"] = "Login failed, please signup"
		responses.JSON(w, http.StatusBadRequest, resp)
		return
	}

	err = models.CheckPasswordHash(user.Password, usr.Password)
	if err != nil {
		resp["status"] = "failed"
		resp["message"] = "Login failed, please try again"
		responses.JSON(w, http.StatusBadRequest, resp)
		return
	}

	token, err := utils.EncodeAuthToken(uint(usr.ID))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	resp["token"] = token
	responses.JSON(w, http.StatusOK, resp)
	return

}

func(a *App) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 {
		count = 10
	}

	if start < 0 {
		start = 0
	}
	users, err :=models.GetAllUsers(a.DB, start, count)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)

}
