package http

import (
	"fmt"
	"net/http"
	"server/internal/models"

	"github.com/labstack/echo/v4"
)

type postStudentsRequest struct {
	Name       string `json:"name"`
	Age        int64  `json:"age"`
	Profession string `json:"profession"`
}

type postStudentsResponse struct {
	Error string `json:"error,omitempty"`
}

func newPostStudentsHandler(studentsRepository studentsRepository) *postStudentsHandler {
	return &postStudentsHandler{
		studentsRepository: studentsRepository,
	}
}

type postStudentsHandler struct {
	studentsRepository studentsRepository
}

func (h *postStudentsHandler) Handle(ctx echo.Context) error {
	var request postStudentsRequest
	err := ctx.Bind(&request)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, getGradesResponse{})
	}

	student := models.Student{
		Name:       request.Name,
		Age:        request.Age,
		Profession: request.Profession,
	}

	err = h.studentsRepository.AddStudent(student)
	if err != nil {
		return fmt.Errorf("cannot add student: %v", err)
	}

	return ctx.JSON(http.StatusOK, postStudentsResponse{})
}
