package main

import (
	"fmt"
	"log"
	"mediaboost/config"
)

func main() {
	appConfig, err := config.MustConfigApp()
	if err != nil {
		log.Fatal(err)
	}
	dbConfig, err := config.MustConfigDB()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(appConfig.Host, dbConfig.Host)
}
