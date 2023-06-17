package http

import (
	"fmt"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

func NewServer(
	studentsRepository studentsRepository,
) (*server, error) {
	port, err := strconv.ParseInt(os.Getenv("SERVER_PORT"), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("cannot parse port: %v", err)
	}

	instance := echo.New()
	instance.Server.Addr = fmt.Sprintf(":%d", port)

	instance.Add("GET", "students", newWrapper(newGetStudentsHandler(studentsRepository)).Handle)
	instance.Add("POST", "students", newWrapper(newPostStudentsHandler(studentsRepository)).Handle)
	instance.Add("GET", "grades/:studentId", newWrapper(newGetGradesHandler(studentsRepository)).Handle)
	instance.Add("POST", "grades", newWrapper(newPostGradesHandler(studentsRepository)).Handle)

	return &server{
		echo: instance,
	}, nil
}

type server struct {
	echo *echo.Echo
}

func (s *server) Serve() error {
	return s.echo.StartServer(s.echo.Server)
}
