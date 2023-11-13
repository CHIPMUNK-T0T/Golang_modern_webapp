package main

import (
	"flag"
	"go-rest-api/controller"
	"go-rest-api/db"
	"go-rest-api/repository"
	"go-rest-api/router"
	"go-rest-api/usecase"
	"os"
)

func main() {
	// コマンドライン引数から環境変数を設定
	var env = flag.String("GO_ENV", "", "environment")
	flag.Parse()
	if *env != "" {
		os.Setenv("GO_ENV", *env)
	}
	
	db := db.NewDB()
	userPepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userPepository)
	userController := controller.NewUserController(userUsecase)

	taskRepository := repository.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskController := controller.NewTaskController(taskUsecase)

	e := router.NewRouter(userController, taskController)

	e.Logger.Fatal(e.Start(":8080"))
}
