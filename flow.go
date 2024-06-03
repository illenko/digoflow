package digoflow

import (
	"github.com/illenko/digoflow/container"
	"github.com/illenko/digoflow/task"
)

type Flow struct {
	ID             string        `yaml:"id"`
	Name           string        `yaml:"name"`
	Entrypoint     Entrypoint    `yaml:"entrypoint"`
	Input          HttpInput     `yaml:"input"`
	TaskConfigs    []task.Config `yaml:"tasks"`
	Container      *container.Container
	ExecutionTasks []task.ExecutionTask
}
