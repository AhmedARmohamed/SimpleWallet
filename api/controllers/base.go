package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"wallet/api/middlewares"
	"wallet/api/responses"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

//Initialize connect to the database and wire up routes
func (a *App) Initialize(DBHost, DBPort, DBUser, DBName string)  {
	var err error
	DBUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable ", DBHost, DBPort, DBUser, DBName)

	a.DB, err = sql.Open("postgres", DBUri)
	if err != nil {
		fmt.Printf("\n Cannot connect to the databse %s", DBName)
		 log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the database %s", DBName)
	}

	//a.DB.Debug().AutoMigrate(&models.User{}) //database migration

	a.Router = mux.NewRouter().StrictSlash(true)
	a.intializeRoutes()
}

func (a *App) intializeRoutes() {
	a.Router.Use(middlewares.SetContentTypeMiddleware) // setting content-type to json

	a.Router.HandleFunc("/", home).Methods("GET")
	a.Router.HandleFunc("/register", a.userSignUp).Methods("POST")
	a.Router.HandleFunc("/login", a.login).Methods("POST")

	s := a.Router.PathPrefix("/api").Subrouter() // subrouter to add auth middleware
	s.Use(middlewares.AuthJwtVerify)

	//s.HandleFunc("/users", a.GetAllUsers).Methods("GET")
	s.HandleFunc("/balance/{id:[0-9]}", a.balanceInquiry).Methods("GET")
	s.HandleFunc("/wallet", a.createWallet).Methods("POST")
	s.HandleFunc("/deposit/{id:[0-9]}", a.depositMoney).Methods("POST")
	s.HandleFunc("/withdraw/{id:[0-9]}", a.withdrawMoney).Methods("POST")
}

func (a  *App) RunServer() {
	log.Printf("\nServer starting on port 8020")
	log.Fatal(http.ListenAndServe(":8020", a.Router))
}

func home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to my wallet service")
}