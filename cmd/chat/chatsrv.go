package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/maxisantomil/GoLang2020.git/internal/config"
	"github.com/maxisantomil/GoLang2020.git/internal/database"
	"github.com/maxisantomil/GoLang2020.git/internal/service/chat"
)

func main() {

	cfg := readConfig() // lee la configuracion

	db, err := database.NewDatabase(cfg)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	service, _ := chat.New(db, cfg) // inyecto al chatService una configuracion
	HTTPService := chat.NewHTTPTransport(service)
	r := gin.Default()
	HTTPService.Register(r)
	r.Run()

	/*for _, m := range service.FindAll() {
		fmt.Println(m)
	}*/
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

func createSchema(db *sqlx.DB) error {
	schema := `CREATE TABLE IF NOT EXISTS vino (
		id integer primary key autoincrement,
		Name varchar,
		tipo varchar,
		año int,
		precio);`

	// execute a query on the server
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}
	return nil
}
