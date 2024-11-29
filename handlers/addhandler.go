package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/MINIbra1n/restdb"
)

func Addhandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("AddHandler Serving:", r.URL.Path, "from", r.Host)
	d, err := io.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	if len(d) == 0 {
		rw.WriteHeader(http.StatusBadRequest)
		log.Println("No input")
		return
	}

	var users = []restdb.User{}
	err = json.Unmarshal(d, &users)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println(users)

	if !restdb.IsUserAdmin(users[0]) {
		log.Println("Command issued by non-admin user:", users[0].Username)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	result := restdb.InsertUser(users[1])
	if !result {
		rw.WriteHeader(http.StatusBadRequest)
	}

}
