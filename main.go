package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/illenko/digoflow-protorype/internal/digoflow"
	"github.com/illenko/digoflow-protorype/internal/task"
)

func main() {

	app, err := digoflow.NewApp("flows")

	if err != nil {
		fmt.Printf("Error loading app: %v", err)
		return
	}

	app.RegisterTask("uuidGenerator", uuidGenerator)

	err = app.Start()

	if err != nil {
		fmt.Printf("Error starting app: %v", err)
		return
	}
}

func uuidGenerator(_ task.Input) (task.Output, error) {
	return task.Output{"id": uuid.New().String()}, nil
}
