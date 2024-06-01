package component

import "github.com/illenko/digoflow-protorype/internal/component/task"

type Flow struct {
	ID             string        `yaml:"id"`
	Name           string        `yaml:"name"`
	Entrypoint     Entrypoint    `yaml:"entrypoint"`
	Input          HttpInput     `yaml:"input"`
	TaskConfigs    []task.Config `yaml:"tasks"`
	ExecutionTasks []task.ExecutionTask
}
