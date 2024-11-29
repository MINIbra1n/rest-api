package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/MINIbra1n/rest-api/handlers"
	"github.com/gorilla/mux"
)

// Create a new ServeMux using Gorilla
var rMux = mux.NewRouter()

// PORT is where the web server listens to
var PORT = ":1234"

func main() {
	arguments := os.Args
	if len(arguments) >= 2 {
		PORT = ":" + arguments[1]
	}

	s := http.Server{
		Addr:         PORT,
		Handler:      rMux,
		ErrorLog:     nil,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	rMux.NotFoundHandler = http.HandlerFunc(handlers.DefaultHandler)

	notAllowed := handlers.NotAllowedHandler{}
	rMux.MethodNotAllowedHandler = notAllowed

	rMux.HandleFunc("/time", handlers.TimeHandler)

	// Define Handler Functions
	// Register GET
	getMux := rMux.Methods(http.MethodGet).Subrouter()

	getMux.HandleFunc("/getall", handlers.GetAllHandler)
	getMux.HandleFunc("/getid/{username}", handlers.GetIDHandler)
	getMux.HandleFunc("/logged", handlers.LoggedUsersHandler)
	getMux.HandleFunc("/username/{id:[0-9]+}", handlers.GetUserDataHandler)

	// Register PUT
	// Update User
	putMux := rMux.Methods(http.MethodPut).Subrouter()
	putMux.HandleFunc("/update", handlers.UpdateHandler)

	// Register POST
	// Add User + Login + Logout
	postMux := rMux.Methods(http.MethodPost).Subrouter()
	postMux.HandleFunc("/add", handlers.Addhandler)
	postMux.HandleFunc("/login", handlers.LoginHandler)
	postMux.HandleFunc("/logout", handlers.LogoutHandler)

	// Register DELETE
	// Delete User
	deleteMux := rMux.Methods(http.MethodDelete).Subrouter()
	deleteMux.HandleFunc("/username/{id:[0-9]+}", handlers.DeleteHandler)

	go func() {
		log.Println("Listening to", PORT)
		err := s.ListenAndServe()
		if err != nil {
			log.Printf("Error starting server: %s\n", err)
			return
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	sig := <-sigs
	log.Println("Quitting after signal:", sig)
	time.Sleep(5 * time.Second)
	s.Shutdown(nil)
}
