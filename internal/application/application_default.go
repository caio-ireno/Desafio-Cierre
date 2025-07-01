package application

import (
	"app/internal/handler"
	"app/internal/loader"
	"app/internal/repository"
	"app/internal/service"
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	// ...outros imports...
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
)

type ConfigAppDefault struct {
	ServerAddr string
	DbFilePath string
}

func NewServerChi(cfg *ConfigAppDefault) *ApplicationDefault {
	defaultConfig := &ConfigAppDefault{
		ServerAddr: ":8080",
	}

	if cfg != nil {
		if cfg.ServerAddr != "" {
			defaultConfig.ServerAddr = cfg.ServerAddr
		}

		if cfg.DbFilePath != "" {
			defaultConfig.DbFilePath = cfg.DbFilePath
		}
	}

	return &ApplicationDefault{
		serverAddr: defaultConfig.ServerAddr,
		dbFilePath: defaultConfig.DbFilePath,
	}
}

type ApplicationDefault struct {
	serverAddr string
	dbFilePath string
}

func (a *ApplicationDefault) SetUp() (err error) {

	// --- CONEXÃO COM O BANCO ---
	dsn := "root:root@tcp(db:3306)/tickets" // ajuste conforme suas variáveis de ambiente
	dbsql, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Erro ao conectar no banco:", err)
		os.Exit(1)
	}
	defer dbsql.Close()

	// --- IMPORTAÇÃO DO CSV PARA O BANCO ---
	ctx := context.Background()

	loaderDB := loader.NewLoaderTicketCSV(a.dbFilePath)
	err = loaderDB.LoadTicketsToDB(ctx, dbsql, a.dbFilePath)
	if err != nil {
		fmt.Println("Erro ao importar CSV:", err)
		os.Exit(1)
	}

	ld := loader.NewLoaderTicketCSV(a.dbFilePath)
	db, err := ld.Load()
	lastId := len(db)
	if err != nil {
		return
	}

	rp := repository.NewRepositoryTicketMap(db, lastId)

	sv := service.NewServiceTicketDefault(rp)
	hd := handler.NewHandlerTicketDefault(sv)

	rt := chi.NewRouter()

	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)

	rt.Route("/tickets", func(rt chi.Router) {

		rt.Get("/", hd.GetAll())
		rt.Get("/total_amount", hd.GetTotalAmountTickets())

		rt.Patch("/update/{id}", hd.Update())
		rt.Post("/", hd.Create())
	})

	err = http.ListenAndServe(a.serverAddr, rt)

	return
}
