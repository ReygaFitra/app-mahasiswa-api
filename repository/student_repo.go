package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ReygaFitra/app-mahasiswa-api/model"
)

type StudentRepo interface {
	GetAll() any
	GetById(id int) any
	Create(newStudent *model.Student) string
	Update(student *model.Student) string
	Delete(id int) string
}

type studentRepo struct {
	db *sql.DB
}

func (r *studentRepo) GetAll() any {
	var students []model.Student

	query := "SELECT id, name, age, major, student_user_name FROM student"
	rows, err := r.db.Query(query)

	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var student model.Student

		if err := rows.Scan(&student.Id, &student.Name, &student.Age, &student.Major, &student.StudentUserName); err != nil {
			log.Println(err)
		}

		students = append(students, student)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	if len(students) == 0 {
		return "no data"
	}

	return students
}

func (r *studentRepo) GetById(id int) any {
	var studentDB model.Student

	query := "SELECT id, name, age, major, student_user_name FROM student WHERE id = $1"
	row := r.db.QueryRow(query, id)

	err := row.Scan(&studentDB.Id, &studentDB.Name, &studentDB.Age, &studentDB.Major, &studentDB.StudentUserName)

	if err != nil {
		log.Println(err)
	}

	if studentDB.Id == 0 {
		return "student not found!"
	}

	return studentDB
}

func (r *studentRepo) Create(newStudent *model.Student) string {
	query := "INSERT INTO student (name, age, major, student_user_name) VALUES ($1, $2, $3, $4)"
	_, err := r.db.Exec(query, newStudent.Name, newStudent.Age, newStudent.Major, newStudent.StudentUserName)

	if err != nil {
		log.Println(err)
		return "failed to create student!"
	}

	return "created successfully"
}

func (r *studentRepo) Update(student *model.Student) string {
	res := r.GetById(student.Id)

	if res == "student not found!" {
		return res.(string)
	}

	query := "UPDATE student SET name = $1, age = $2, major = $3, student_user_name = $4 WHERE id = $5"
	_, err := r.db.Exec(query, student.Name, student.Age, student.Major, student.StudentUserName, student.Id)

	if err != nil {
		log.Println(err)
	}

	return fmt.Sprintf("student with id %d updated successfully", student.Id)
}

func (r *studentRepo) Delete(id int) string {
	res := r.GetById(id)

	if res == "student not found!" {
		return res.(string)
	}

	query := "DELETE FROM student WHERE id = $1"
	_, err := r.db.Exec(query, id)

	if err != nil {
		log.Println(err)
		return "failed to delete student"
	}

	return fmt.Sprintf("student with id %d deleted successfully", id)
}

func NewStudentRepo(db *sql.DB) StudentRepo {
	repo := new(studentRepo)
	repo.db = db

	return repo
}