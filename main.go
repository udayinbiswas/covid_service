package main

import (
	"log"

	"github.com/udayinbiswas21/covid_service/external"
)

func main() {
	res, err := external.GetCrowdSourcedData()
	if err != nil {
		log.Println(err)
	}
	log.Println("Result is", res)
	log.Println("Result length is", len(res.CasesTimeSeries))
}
