package api

import (
	"log"
	"net/http"
)

func (api *API) ping(w http.ResponseWriter, r *http.Request) {
	log.Println("OK")
}
