package database

import (
	"app-mahasiswa-api/controller"
	"app-mahasiswa-api/repository"
	"app-mahasiswa-api/usecase"
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func ConnectDB() {
	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "postgres"
	dbPassword := "postgre123"
	dbName := "app_mahasiswa"
	sslMode := "disable"
	serverPort := ":8080"

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