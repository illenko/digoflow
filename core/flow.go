package core

import (
	"github.com/illenko/digoflow/core/entrypoint"
	"github.com/illenko/digoflow/task"
)

type Flow struct {
	ID             string                `yaml:"id"`
	Name           string                `yaml:"name"`
	Entrypoint     entrypoint.Entrypoint `yaml:"entrypoint"`
	Input          HttpInput             `yaml:"input"`
	TaskConfigs    []task.Config         `yaml:"tasks"`
	ExecutionTasks []task.ExecutionTask
}
