package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Microservices/eventsservice/rest"
	"github.com/Microservices/lib/configuration"
	"github.com/Microservices/lib/persistence/dblayer"
)

func main() {
	confPath := flag.String("conf", `.\configuration\config.json`, "flag to set the path to the configuration json file")
	flag.Parse()

	//extract configuration
	config, _ := configuration.ExtractConfiguration(*confPath)

	fmt.Println("Connecting to database ")
	dbhandler, _ := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)

	//start Restful API
	log.Fatal(rest.ServeAPI(config.RestfulEndpoint, dbhandler))
}
