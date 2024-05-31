package model

type TaskInput map[string]any

type TaskOutput map[string]any

type ExecutionTask func(TaskInput) TaskOutput

type App struct {
	Flows       map[string]Flow
	CustomTasks map[string]ExecutionTask
}
