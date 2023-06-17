package http

import (
	"fmt"
	"net/http"
	"server/internal/models"

	"github.com/labstack/echo/v4"
)

type getGradesRequest struct {
	Name       string `json:"name"`
	Age        int64  `json:"age"`
	Profession string `json:"profession"`
}

type getGradesResponse struct {
	Error string `json:"error,omitempty"`
}

func newGetGradesHandler(studentsRepository studentsRepository) *getGradesHandler {
	return &getGradesHandler{
		studentsRepository: studentsRepository,
	}
}

type getGradesHandler struct {
	studentsRepository studentsRepository
}

func (h *getGradesHandler) Handle(ctx echo.Context) error {
	var request getGradesRequest
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

	return ctx.JSON(http.StatusOK, getGradesResponse{})
}
