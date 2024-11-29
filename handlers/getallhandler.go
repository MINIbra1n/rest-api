package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/MINIbra1n/restdb"
)

func GetAllHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("GetAllHandler Serving:", r.URL.Path, "from", r.Host)
	d, err := io.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	if len(d) == 0 {
		rw.WriteHeader(http.StatusBadRequest)
		log.Println("No input!")
		return
	}

	var user = restdb.User{}
	err = json.Unmarshal(d, &user)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	if !restdb.IsUserAdmin(user) {
		log.Println("User", user, "is not an admin!")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err = SliceToJSON(restdb.ListAllUsers(), rw)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
}
