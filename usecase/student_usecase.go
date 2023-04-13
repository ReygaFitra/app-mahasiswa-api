package usecase

import (
	"app-mahasiswa-api/model"
	"app-mahasiswa-api/repository"
)

type StudentUsecase interface {
	FindStudents() any
	FindStudent(id int) any
	Register(newStudent *model.Student) string
	Edit(student *model.Student) string
	Unreg(id int) string
}

type studentUsecase struct {
	studentRepo repository.StudentRepo
}

func (u *studentUsecase) FindStudents() any {
	return u.studentRepo.GetAll()
}

func (u *studentUsecase) FindStudent(id int) any {
	return u.studentRepo.GetById(id)
}

func (u *studentUsecase) Register(newStudent *model.Student) string {

	var studentMaximum [] model.Student
	if len(studentMaximum) > 10 {
		return "Maksimal 10 data Mahasiswa!"
	}

		if len(newStudent.Name) < 3 {
			return("nama minimal 3 Karakter dan maksimal 20 karakter")
		} else if len(newStudent.Name) >= 20 {
			return("nama minimal 3 Karakter dan maksimal 20 karakter")
		} else if newStudent.Age < 17 {
			return("Umur Minimal 17 Tahun")
		} else if len(newStudent.Major) > 10 {
			return("Jurusan Maksimal 10 Karakter")
		}
	return u.studentRepo.Create(newStudent)
}

func (u *studentUsecase) Edit(student *model.Student) string {
	if len(student.Name) < 3 {
		return("nama minimal 3 Karakter dan maksimal 20 karakter")
	} else if len(student.Name) >= 20 {
		return("nama minimal 3 Karakter dan maksimal 20 karakter")
	} else if student.Age < 17 {
		return("Umur Minimal 17 Tahun")
	} else if len(student.Major) > 10 {
		return("Jurusan Maksimal 10 Karakter")
	}

	return u.studentRepo.Update(student)
}

func (u *studentUsecase) Unreg(id int) string {
	return u.studentRepo.Delete(id)
}

func NewUserUsecase(studentRepo repository.StudentRepo) StudentUsecase {
	return &studentUsecase{
		studentRepo: studentRepo,
	}
}