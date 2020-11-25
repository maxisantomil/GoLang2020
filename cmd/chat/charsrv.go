package main

import (
	"flag"
	"fmt"
	"os"
	"time"

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
	httpService := chat.NewHTTPTransport(service)
	r := gin.Default()
	httpService.Register(r)
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
	schema := `CREATE TABLE IF NOT EXISTS messages (
		id integer primary key autoincrement,
		text varchar);`

	// execute a query on the server
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	// or, you can use MustExec, which panics on error
	insertMessage := `INSERT INTO messages (text) VALUES (?)`
	s := fmt.Sprintf("Message number %v", time.Now().Nanosecond())
	db.MustExec(insertMessage, s)
	return nil
}
