package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/illenko/digoflow-protorype/internal/component/task"
	"github.com/illenko/digoflow-protorype/internal/config"
	"github.com/illenko/digoflow-protorype/internal/entrypoint/http"
)

func main() {

	app, err := config.LoadApp("flows")

	app.RegisterCustomTask("use-custom-log", logTask)

	if err != nil {
		panic(err)
	}

	g := gin.Default()

	for _, f := range app.Flows {

		for _, tc := range f.TaskConfigs {
			if tc.Type == "custom" {
				t, ok := app.CustomTasks[tc.ID]

				if !ok {
					panic("task not found: " + tc.ID)
				}

				f.ExecutionTasks = append(f.ExecutionTasks, t)
			} else {
				t, ok := app.BuiltInTasks[tc.Type]

				if !ok {
					panic("task not found: " + tc.ID)
				}

				f.ExecutionTasks = append(f.ExecutionTasks, t)
			}
		}

		if f.Entrypoint.Type == "http-handler" {
			http.NewHandler(f, g)
		}
	}

	err = g.Run(":8080")
	if err != nil {
		return
	}
}

func logTask(values task.Input) (task.Output, error) {
	fmt.Printf("Message from custom log task: %s\n", values["message-to-print"])

	output := map[string]any{}
	output["message"] = fmt.Sprintf("Hello from custom log task with message: %s", values["message-to-print"])

	return output, nil
}
