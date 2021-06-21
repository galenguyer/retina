package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/galenguyer/retina/storage"
)

func StartServer() {
	fileServer := http.FileServer(http.Dir("./web/app/build/"))
	http.HandleFunc("/api/v1/hour", GetLastHour)
	http.Handle("/", http.StripPrefix(strings.TrimRight("/", "/"), fileServer))
	log.Println("starting webserver on port 8000")
	http.ListenAndServe(":8000", nil)
}

func GetLastHour(w http.ResponseWriter, r *http.Request) {
	res := storage.GetLastHour()
	json.NewEncoder(w).Encode(res)
}
