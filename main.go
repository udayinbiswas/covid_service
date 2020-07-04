package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/udayinbiswas21/covid_service/external"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	res, err := external.GetCrowdSourcedData()
	if err != nil {
		log.Println(err)
	}

	w.Write([]byte("Number of results:" + strconv.Itoa(len(res.CasesTimeSeries))))
}

func main() {
	port := "3000"
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)

	http.ListenAndServe(":"+port, mux)
}
