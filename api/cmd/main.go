package main

import (
	"fmt"
	"server/internal/database"
	"server/internal/http"
	"server/internal/repositories/students"

	"github.com/go-pg/pg/v10"
	"github.com/labstack/gommon/log"
)

func main() {
	db := database.Connect()
	defer func(db *pg.DB) {
		err := database.Disconnect(db)
		if err != nil {
			log.Error(fmt.Errorf("cannot disconnect with db: %v", err))
		}
	}(db)

	studentsRepository := students.NewRepository(db)
	err := studentsRepository.SetupTable()
	if err != nil {
		log.Error(fmt.Errorf("cannot setup table in students repository: %v", err))
		return
	}

	server, err := http.NewServer(studentsRepository)
	if err != nil {
		log.Error(fmt.Errorf("cannot create server: %v", err))
		return
	}

	err = server.Serve()
	if err != nil {
		log.Error(fmt.Errorf("error during servant: %v", err))
	}
}
