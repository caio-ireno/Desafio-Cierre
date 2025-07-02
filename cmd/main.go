package main

import (
	"app/db"
	"app/internal/application"
	"app/internal/loader"
	"context"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	if err := db.Connect(); err != nil {
		panic(err)
	}

	defer db.GetConnection().Close()

	r := application.NewRouter()

	ld := loader.NewLoaderTicketCSV("docs/db/tickets.csv")
	tickets, err := ld.Load()
	if err != nil {
		return
	}

	err = loader.LoadTicketsToDB(context.Background(), db.GetConnection(), tickets)
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":8080", r.TicketsRoutes())
}
