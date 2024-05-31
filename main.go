package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/illenko/digoflow-protorype/internal/cfg"
	"github.com/illenko/digoflow-protorype/internal/entrypoint/http"
	"github.com/illenko/digoflow-protorype/internal/model"
)

func main() {

	app, err := cfg.LoadApp("flows")
	app.CustomTasks = map[string]model.ExecutionTask{
		"log-task": logTask,
	}

	if err != nil {
		panic(err)
	}

	g := gin.Default()

	for _, f := range app.Flows {
		f.ExecutionTasks = append(f.ExecutionTasks, app.CustomTasks["log-task"])
		if f.Entrypoint.Type == "http-handler" {
			http.NewHandler(f, g)
		}
	}

	err = g.Run(":8080")
	if err != nil {
		return
	}
}

func logTask(values model.TaskInput) model.TaskOutput {
	fmt.Printf("Message from custom log task: %s\n", values["message-to-print"])

	output := map[string]any{}
	output["message"] = fmt.Sprintf("Hello from custom log task with message: %s", values["message-to-print"])

	return output
}
