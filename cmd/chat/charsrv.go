package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/maxisantomil/GoLang2020.git/internal/chat"
	"github.com/maxisantomil/GoLang2020.git/internal/config"
)

func main() {

	cfg := readConfig()         // lee la configuracion
	service, _ := chat.New(cfg) // inyecto al chatService una configuracion

	for _, m := range service.FindAll() {
		fmt.Println(m)
	}
}
func readConfig() *config.Config {
	//puntero a un string
	configFile := flag.String("config", "./config.yaml", "this is the service config")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return cfg
}
