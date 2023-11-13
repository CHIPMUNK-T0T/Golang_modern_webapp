package main

import (
	"flag"
	"fmt"
	"go-rest-api/db"
	"go-rest-api/model"
	"os"
)

func main() {
	// コマンドライン引数から環境変数を設定
	var env = flag.String("GO_ENV", "", "environment")
	flag.Parse()
	if *env != "" {
		os.Setenv("GO_ENV", *env)
	}

	dbConn := db.NewDB()
	defer fmt.Println("Successfuly Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}