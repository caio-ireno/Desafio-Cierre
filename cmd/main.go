package main

import (
	"app/db"
	"app/internal/application"
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

	http.ListenAndServe(":8080", r.TicketsRoutes())
}
