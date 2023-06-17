package http

import "server/internal/models"

type studentsRepository interface {
	GetStudents() ([]models.Student, error)
	AddStudent(model models.Student) error
	GetGrades(id int64) ([]models.Grade, error)
	AddGrade(studentId int64, subject string, value int64) error
}
