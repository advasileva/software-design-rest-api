package http

import (
	"fmt"
	"net/http"
	"server/internal/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

type getGradesRequest struct{}

type getGradesResponse struct {
	Grades []models.Grade `json:"grades"`
	Error  string         `json:"error,omitempty"`
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

	id, err := strconv.ParseInt(ctx.Param("studentId"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, getGradesResponse{Error: "cannot parse studentId"})
	}
	grades, err := h.studentsRepository.GetGrades(id)
	if err != nil {
		return fmt.Errorf("cannot get grades: %v", err)
	}

	return ctx.JSON(http.StatusOK, getGradesResponse{Grades: grades})
}
