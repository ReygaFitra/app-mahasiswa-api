package controller

import (
	"app-mahasiswa-api/model"
	"app-mahasiswa-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StudentController struct {
	usecase usecase.StudentUsecase
}

func (c *StudentController) FindStudents(ctx *gin.Context) {
	res := c.usecase.FindStudents()

	ctx.JSON(http.StatusOK, res)
}

func (c *StudentController) FindStudent(ctx *gin.Context) {
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

	if err := ctx.BindJSON(&newStudent); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		return
	}

	res := c.usecase.Register(&newStudent)
	ctx.JSON(http.StatusCreated, res)
}

func (c *StudentController) Edit(ctx *gin.Context) {
	var student model.Student

	if err := ctx.BindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		return
	}

	res := c.usecase.Edit(&student)
	ctx.JSON(http.StatusOK, res)
}

func (c *StudentController) Unreg(ctx *gin.Context) {
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