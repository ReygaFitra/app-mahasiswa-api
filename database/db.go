package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ReygaFitra/app-mahasiswa-api/controller"
	"github.com/ReygaFitra/app-mahasiswa-api/repository"
	"github.com/ReygaFitra/app-mahasiswa-api/usecase"
	"github.com/ReygaFitra/app-mahasiswa-api/utils"
	authController "github.com/ReygaFitra/auth-jwt/controller"
	authRepository "github.com/ReygaFitra/auth-jwt/repository"
	authUsecase "github.com/ReygaFitra/auth-jwt/usecase"

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
	studentRepo := repository.NewStudentRepo(db)
	studentUsecase := usecase.NewUserUsecase(studentRepo)
	studentCtrl := controller.NewStudentController(studentUsecase)
	
	authCtrl := authController.NewAuthController(authUsecase.NewAuthUsecase(authRepository.NewAuthRepo(db)))

	router := gin.Default()
	// login routes
	router.POST("/auth/login", authCtrl.Login)
	// register routes
	router.POST("/api/v1/students", studentCtrl.Register)

	studentRouter := router.Group("/api/v1/students/profile")
	studentRouter.Use(authController.AuthMiddleware())

	studentRouter.GET("", studentCtrl.FindStudents)
	studentRouter.GET("/:id", studentCtrl.FindStudent)
	studentRouter.PUT("", studentCtrl.Edit)
	studentRouter.DELETE("/:id", studentCtrl.Unreg)

	if err := router.Run(serverPort); err != nil {
		log.Fatal(err)
	}
}