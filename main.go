package main

import (
	"fmt"
	errorHandler "github.com/fossyy/WebAppTemplate/handler/error"
	indexHandler "github.com/fossyy/WebAppTemplate/handler/index"
	logoutHandler "github.com/fossyy/WebAppTemplate/handler/logout"
	miscHandler "github.com/fossyy/WebAppTemplate/handler/misc"
	signinHandler "github.com/fossyy/WebAppTemplate/handler/signin"
	signupHandler "github.com/fossyy/WebAppTemplate/handler/signup"
	userHandler "github.com/fossyy/WebAppTemplate/handler/user"
	"github.com/fossyy/WebAppTemplate/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	serverAddr := "localhost:8000"
	handler := mux.NewRouter()
	server := http.Server{
		Addr:    serverAddr,
		Handler: middleware.Handler(handler),
	}

	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			middleware.Auth(indexHandler.GET, w, r)
		case http.MethodPost:
			middleware.Auth(indexHandler.POST, w, r)
		}
	})

	handler.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			middleware.Guest(signinHandler.GET, w, r)
		case http.MethodPost:
			middleware.Guest(signinHandler.POST, w, r)
		}
	})

	handler.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			middleware.Guest(signupHandler.GET, w, r)
		case http.MethodPost:
			middleware.Guest(signupHandler.POST, w, r)
		}
	})

	handler.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			middleware.Auth(userHandler.GET, w, r)
		}
	})

	handler.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		middleware.Auth(logoutHandler.GET, w, r)
	})

	handler.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		miscHandler.Robot(w, r)
	})

	handler.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		miscHandler.Favicon(w, r)
	})

	handler.NotFoundHandler = http.HandlerFunc(errorHandler.ALL)

	fileServer := http.FileServer(http.Dir("./public"))
	handler.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fileServer))

	fmt.Printf("Listening on http://%s\n", serverAddr)
	err := server.ListenAndServe()
	if err != nil {
		return
	}

}
