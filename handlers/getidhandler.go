package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/MINIbra1n/restdb"
	"github.com/gorilla/mux"
)

func GetIDHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("GetIDHandler Serving:", r.URL.Path, "from", r.Host)

	username, ok := mux.Vars(r)["username"]
	if !ok {
		log.Println("ID value not set!")
		rw.WriteHeader(http.StatusNotFound)
		return
	}

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

	log.Println("Input user:", user)
	if !restdb.IsUserAdmin(user) {
		log.Println("User", user.Username, "not an admin!")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	t := restdb.FindUserUsername(username)
	if t.ID != 0 {
		err := t.ToJSON(rw)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			log.Println(err)
		}
	} else {
		rw.WriteHeader(http.StatusNotFound)
		log.Println("User " + user.Username + "not found")
	}
}
