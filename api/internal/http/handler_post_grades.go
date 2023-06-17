package http

import (
	"fmt"
	"net/http"
	"server/internal/models"

	"github.com/labstack/echo/v4"
)

type postGradesRequest struct {
	Name       string `json:"name"`
	Age        int64  `json:"age"`
	Profession string `json:"profession"`
}

type postGradesResponse struct {
	Error string `json:"error,omitempty"`
}

func newPostGradesHandler(studentsRepository studentsRepository) *postGradesHandler {
	return &postGradesHandler{
		studentsRepository: studentsRepository,
	}
}

type postGradesHandler struct {
	studentsRepository studentsRepository
}

func (h *postGradesHandler) Handle(ctx echo.Context) error {
	var request postGradesRequest
	err := ctx.Bind(&request)
	if err != nil {
		return fmt.Errorf("cannot bind request: %v", err)
	}

	student := models.Student{
		Name:       request.Name,
		Age:        request.Age,
		Profession: request.Profession,
	}

	err = h.studentsRepository.AddStudent(student)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, postGradesResponse{})
}
