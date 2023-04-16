package controller

import (
	"app-mahasiswa-api/model"
	"app-mahasiswa-api/usecase"
	"net/http"
	"strconv"
	"time"

	authModel "github.com/ReygaFitra/auth-jwt/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type StudentController struct {
	usecase usecase.StudentUsecase
}

func (c *StudentController) FindStudents(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	username := claims["username"].(string)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Welcome Student",
		"username": username,
	})

	res := c.usecase.FindStudents()
	ctx.JSON(http.StatusOK, res)
}

func (c *StudentController) FindStudent(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	username := claims["username"].(string)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Welcome Student",
		"username": username,
	})

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid student ID")
		return
	}

	res := c.usecase.FindStudent(id)
	ctx.JSON(http.StatusOK, res)
}

func (c *StudentController) Register(ctx *gin.Context) {
	var newStudent model.Student

	err := ctx.ShouldBindJSON(&newStudent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	if newStudent.StudentUserName == "secretkey" {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = newStudent.StudentUserName
		claims["exp"] = time.Now().Add(time.Minute * 3).Unix()

		tokenString, err :=token.SignedString(authModel.JwtKey)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed generate token!"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized Student!"})
	}

	if err := ctx.BindJSON(&newStudent); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		return
	}

	res := c.usecase.Register(&newStudent)
	ctx.JSON(http.StatusCreated, res)
}

func (c *StudentController) Edit(ctx *gin.Context) {
	var student model.Student

	claims := ctx.MustGet("claims").(jwt.MapClaims)
	username := claims["username"].(string)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Welcome Student",
		"username": username,
	})

	if err := ctx.BindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		return
	}

	res := c.usecase.Edit(&student)
	ctx.JSON(http.StatusOK, res)
}

func (c *StudentController) Unreg(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	username := claims["username"].(string)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Welcome Student",
		"username": username,
	})

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid student ID")
		return
	}

	res := c.usecase.Unreg(id)
	ctx.JSON(http.StatusOK, res)
}

func NewStudentController(u usecase.StudentUsecase) *StudentController {
	controller := StudentController{
		usecase: u,
	}

	return &controller
}