package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sichangSun/course-enrollment-manager/router"
)

func main() {
	// Todo: log setting

	// load config from 環境変数
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	dsn := os.Getenv("DSN")
	fmt.Println(dsn)

	// dbコネクション作成
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := router.New(&router.RouterConfig{
		DB: db,
	})
	log.Fatal(r.Run(fmt.Sprintf(":%s", port)))
}
