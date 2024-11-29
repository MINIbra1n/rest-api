package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type NotAllowedHandler struct {
}

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"user"`
	Password  string `json:"password"`
	LastLogin int64  `json:"lastlogin"`
	Admin     int    `json:"admin"`
	Active    int    `json:"active"`
}

func SliceToJSON(slice interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(slice)
}

func (h NotAllowedHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	MethodNotAllowedHandler(rw, r)
}

// MethodNotAllowedHandler is executed when the HTTP method is incorrect
func MethodNotAllowedHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host, "with method", r.Method)
	rw.WriteHeader(http.StatusNotFound)
	Body := "Method not allowed!\n"
	fmt.Fprintf(rw, "%s", Body)
}

func DefaultHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("DefaultHandler Serving:", r.URL.Path, "from", r.Host, "with method", r.Method)
	rw.WriteHeader(http.StatusNotFound)
	Body := r.URL.Path + " is not supported. Thanks for visiting!\n"
	fmt.Fprintf(rw, "%s", Body)
}

// TimeHandler is for handling /time – it works with plain text
func TimeHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("TimeHandler Serving:", r.URL.Path, "from", r.Host)
	rw.WriteHeader(http.StatusOK)
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is: " + t + "\n"
	fmt.Fprintf(rw, "%s", Body)
}
