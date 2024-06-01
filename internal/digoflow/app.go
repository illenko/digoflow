package digoflow

import (
	"fmt"
	"os"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/illenko/digoflow-protorype/internal/component"
	"github.com/illenko/digoflow-protorype/internal/component/task"
	"github.com/illenko/digoflow-protorype/internal/entrypoint/http"
	"gopkg.in/yaml.v2"
)

type App struct {
	Flows        map[string]component.Flow
	BuiltInTasks map[string]task.ExecutionTask // type -> ExecutionTask
	CustomTasks  map[string]task.ExecutionTask // name -> ExecutionTask
}

func NewApp(flowsDir string) (*App, error) {
	files, err := os.ReadDir(flowsDir)
	if err != nil {
		return nil, err
	}

	app := App{
		Flows: make(map[string]component.Flow),
		BuiltInTasks: map[string]task.ExecutionTask{
			"log": task.Log,
		},
		CustomTasks: make(map[string]task.ExecutionTask),
	}

	for _, file := range files {
		if !file.IsDir() {
			data, err := os.ReadFile(flowsDir + "/" + file.Name())
			if err != nil {
				return nil, err
			}

			var flow component.Flow
			err = yaml.Unmarshal(data, &flow)
			if err != nil {
				return nil, err
			}

			app.RegisterFlow(flow)
		}
	}

	return &app, nil
}

func (a *App) RegisterFlow(flow component.Flow) {
	a.Flows[flow.ID] = flow
}

func (a *App) RegisterCustomTask(taskType string, task task.ExecutionTask) {
	a.CustomTasks[taskType] = task
}

func (a *App) Start() error {
	g := gin.Default()

	ep, err := a.registerFlows(g)
	if err != nil {
		fmt.Printf("Error registering flows: %v", err)
		return err
	}

	if slices.Contains(ep, "http-handler") {
		err = g.Run(":8080")

		if err != nil {
			fmt.Printf("Error running server: %v", err)
			return err
		}
	}

	return nil
}

func (a *App) registerFlows(g *gin.Engine) ([]string, error) {
	entrypointTypes := make([]string, 0)
	seen := make(map[string]bool)

	for _, f := range a.Flows {
		err := a.registerTasks(&f)
		if err != nil {
			return nil, err
		}

		if f.Entrypoint.Type == "http-handler" {
			http.NewHandler(f, g)
		}

		if _, ok := seen[f.Entrypoint.Type]; !ok {
			entrypointTypes = append(entrypointTypes, f.Entrypoint.Type)
			seen[f.Entrypoint.Type] = true
		}
	}

	return entrypointTypes, nil
}

func (a *App) registerTasks(f *component.Flow) error {
	for _, tc := range f.TaskConfigs {
		var t task.ExecutionTask
		var ok bool

		if tc.Type == "custom" {
			t, ok = a.CustomTasks[tc.ID]
		} else {
			t, ok = a.BuiltInTasks[tc.Type]
		}

		if !ok {
			return fmt.Errorf("task not found: %s", tc.ID)
		}

		f.ExecutionTasks = append(f.ExecutionTasks, t)
	}
	return nil
}