// package db

// import (
// 	"app/pkg/apperrors"
// 	"database/sql"
// 	"fmt"
// 	"os"
// 	"time"
// )

// var conn *sql.DB

// func Connect() (err error) {
// 	connectionsParams := fmt.Sprintf(
// 		"%s:%s@tcp(%s:%s)/%s",
// 		os.Getenv("DB_USER"),
// 		os.Getenv("DB_PASSWORD"),
// 		os.Getenv("DB_HOST"),
// 		os.Getenv("DB_PORT"),
// 		os.Getenv("DB_NAME"),
// 	)

// 	var db *sql.DB
// 	for i := 0; i < 10; i++ {
// 		db, err = sql.Open("mysql", connectionsParams)
// 		if err == nil {
// 			err = db.Ping()
// 			if err == nil {
// 				conn = db
// 				return nil
// 			}
// 		}
// 		fmt.Println("Aguardando banco de dados...", err)
// 		time.Sleep(3 * time.Second)
// 	}

// 	fmt.Println("Erro real ao conectar no banco:", err)
// 	err = apperrors.ErrToConnectDB
// 	return
// }

// func GetConnection() *sql.DB {
// 	return conn
// }

package db

import (
	"app/pkg/apperrors"
	"database/sql"
	"fmt"
	"os"
)

var conn *sql.DB

func Connect() (err error) {
	connectionsParams := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("mysql", connectionsParams)
	if err != nil {
		fmt.Println("Erro real ao conectar no banco:", err)
		err = apperrors.ErrToConnectDB
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Erro real ao dar ping no banco:", err)
		return err
	}

	conn = db

	return nil
}

func GetConnection() *sql.DB {
	return conn
}
