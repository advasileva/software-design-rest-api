package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type postGradesRequest struct {
	StudentId int64  `json:"student_id"`
	Subject   string `json:"subject"`
	Grade     int64  `json:"grade"`
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

	err = h.studentsRepository.AddGrade(request.StudentId, request.Subject, request.Grade)
	if err != nil {
		return fmt.Errorf("cannot add grade: %v", err)
	}

	return ctx.JSON(http.StatusOK, postGradesResponse{})
}
