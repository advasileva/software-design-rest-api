package http

import (
	"fmt"
	"net/http"
	"server/internal/models"

	"github.com/labstack/echo/v4"
)

type getStudentsRequest struct{}

type getStudentsResponse struct {
	Students []models.Student `json:"students"`
	Error    string           `json:"error,omitempty"`
}

func newGetStudentsHandler(studentsRepository studentsRepository) *getStudentsHandler {
	return &getStudentsHandler{
		studentsRepository: studentsRepository,
	}
}

type getStudentsHandler struct {
	studentsRepository studentsRepository
}

func (h *getStudentsHandler) Handle(ctx echo.Context) error {
	var request getStudentsRequest
	err := ctx.Bind(&request)
	if err != nil {
		return fmt.Errorf("cannot bind request: %v", err)
	}

	students, err := h.studentsRepository.GetStudents()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, getStudentsResponse{
		Students: students,
	})
}
