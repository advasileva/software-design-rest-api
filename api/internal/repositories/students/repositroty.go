package students

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"server/internal/models"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func NewRepository(db *pg.DB) *repository {
	return &repository{
		db: db,
	}
}

type repository struct {
	db *pg.DB
}

func (r *repository) SetupTable() error {
	models := []interface{}{
		(*student)(nil),
		(*grade)(nil),
	}

	for _, model := range models {
		err := r.db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return fmt.Errorf("cannot setup tables: %v", err)
		}
	}

	return nil
}

func (r *repository) GetStudents() ([]models.Student, error) {
	dto := &[]student{}
	r.db.Model(dto).Select()
	//if err != nil {
	//	return models.Order{}, fmt.Errorf("cannot get order row: %v", err)
	//}

	students := make([]models.Student, 0, len(*dto))
	for _, dtoStudent := range *dto {
		students = append(students, models.Student{
			Name:       dtoStudent.Name,
			Age:        dtoStudent.Age,
			Profession: dtoStudent.Profession,
		})
	}

	return students, nil
}

func (r *repository) AddStudent(model models.Student) error {
	dto := &student{
		Name:       model.Name,
		Age:        model.Age,
		Profession: model.Profession,
	}

	_, err := r.db.Model(dto).Insert()
	if err != nil {
		log.Infof("cannot insert to db", err)
		return err
	}

	return nil
}

func (r *repository) GetGrades(id int64) ([]models.Grade, error) {
	dto := &[]grade{}
	err := r.db.Model(dto).Where("student_id = ?", id).Select()
	if err != nil {
		return []models.Grade{}, fmt.Errorf("cannot get rowS: %v", err)
	}

	grades := make([]models.Grade, 0, len(*dto))
	for _, dtoGrade := range *dto {
		grades = append(grades, models.Grade{
			Subject: dtoGrade.Subject,
			Grade:   dtoGrade.Grade,
		})
	}

	return grades, nil
}

func (r *repository) AddGrade(studentId int64, subject string, value int64) error {
	dto := &grade{
		StudentId: studentId,
		Subject:   subject,
		Grade:     value,
	}

	_, err := r.db.Model(dto).Insert()
	if err != nil {
		log.Infof("cannot insert to db", err)
		return err
	}

	return nil
}
