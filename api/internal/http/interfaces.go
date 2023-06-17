package http

import "server/internal/models"

type studentsRepository interface {
	GetStudents() ([]models.Student, error)
	AddStudent(model models.Student) error
}
