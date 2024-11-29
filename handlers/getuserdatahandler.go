package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/MINIbra1n/restdb"
	"github.com/gorilla/mux"
)

func GetUserDataHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("GetUserDataHandler Serving:", r.URL.Path, "from", r.Host)
	id, ok := mux.Vars(r)["id"]
	if !ok {
		log.Println("ID value not set!")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		log.Println("id", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	t := restdb.FindUserID(intID)
	if t.ID != 0 {
		err := t.ToJSON(rw)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			log.Println(err)
		}
		return
	}

	log.Println("User not found:", id)
	rw.WriteHeader(http.StatusBadRequest)
}
