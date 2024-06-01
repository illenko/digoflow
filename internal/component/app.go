package component

import "github.com/illenko/digoflow-protorype/internal/component/task"

type App struct {
	Flows        map[string]Flow
	BuiltInTasks map[string]task.ExecutionTask // type -> ExecutionTask
	CustomTasks  map[string]task.ExecutionTask // name -> ExecutionTask
}

func NewApp() *App {
	return &App{
		Flows: make(map[string]Flow),
		BuiltInTasks: map[string]task.ExecutionTask{
			"log": task.Log,
		},
		CustomTasks: make(map[string]task.ExecutionTask),
	}
}

func (a *App) RegisterFlow(flow Flow) {
	a.Flows[flow.ID] = flow
}

func (a *App) RegisterCustomTask(taskType string, task task.ExecutionTask) {
	a.CustomTasks[taskType] = task
}
