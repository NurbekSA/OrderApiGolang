package main

import (
	"HalykProject/internal/starter"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
)

func main() {
	flag.Parse()
	// Чтение содержимого JSON-файла
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	config := starter.NewConfig()

	// Распаковка JSON в структуру
	err = json.Unmarshal(file, config)
	if err != nil {
		log.Fatal(err)
	}

	apiServer := starter.ApiServer{Сonfig: config}

	apiServer.Run()
}
