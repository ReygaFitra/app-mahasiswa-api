package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ReygaFitra/app-mahasiswa-api/controller"
	"github.com/ReygaFitra/app-mahasiswa-api/repository"
	"github.com/ReygaFitra/app-mahasiswa-api/usecase"
	"github.com/ReygaFitra/app-mahasiswa-api/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func ConnectDB() {
	dbHost := utils.DotEnv("DB_HOST")
	dbPort := utils.DotEnv("DB_PORT")
	dbUser := utils.DotEnv("DB_USER")
	dbPassword := utils.DotEnv("DB_PASSWORD")
	dbName := utils.DotEnv("DB_NAME")
	sslMode := utils.DotEnv("SSL_MODE")
	serverPort := utils.DotEnv("SERVER_PORT")

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPassword, dbName, sslMode)
	db, err := sql.Open("postgres", dataSourceName)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	router := gin.Default()

	studentRouter := router.Group("/api/v1/students")

	studentRepo := repository.NewStudentRepo(db)
	studentUsecase := usecase.NewUserUsecase(studentRepo)
	studentCtrl := controller.NewStudentController(studentUsecase)

	studentRouter.GET("", studentCtrl.FindStudents)
	studentRouter.GET("/:id", studentCtrl.FindStudent)
	studentRouter.POST("", studentCtrl.Register)
	studentRouter.PUT("", studentCtrl.Edit)
	studentRouter.DELETE("/:id", studentCtrl.Unreg)

	if err := router.Run(serverPort); err != nil {
		log.Fatal(err)
	}
}