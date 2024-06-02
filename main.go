package main

import (
	"fmt"

	"github.com/illenko/digoflow-protorype/internal/component/task"
	"github.com/illenko/digoflow-protorype/internal/digoflow"
)

func main() {

	app, err := digoflow.NewApp("flows")

	if err != nil {
		fmt.Printf("Error loading app: %v", err)
		return
	}

	app.RegisterTask("custom-log", logTask)

	err = app.Start()

	if err != nil {
		fmt.Printf("Error starting app: %v", err)
		return
	}
}

func logTask(values task.Input) (task.Output, error) {
	fmt.Printf("Message from custom log task: %s\n", values["message-to-print"])

	output := map[string]any{}
	output["message"] = fmt.Sprintf("Hello from custom log task with message: %s", values["message-to-print"])

	return output, nil
}
