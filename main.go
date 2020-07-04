package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/udayinbiswas21/covid_service/external"
)

var tpl = template.Must(template.ParseFiles("templates/index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	res, err := external.GetCrowdSourcedData()
	if err != nil {
		log.Println(err)
	}
	tpl.Execute(w, res)
}

func main() {
	port := "3000"
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)

	http.ListenAndServe(":"+port, mux)
}
