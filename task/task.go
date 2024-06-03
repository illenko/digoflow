package task

import (
	"github.com/illenko/digoflow/container"
)

type Input map[string]any

type Output map[string]any

type Config struct {
	ID    string        `yaml:"id"`
	Name  string        `yaml:"name"`
	Type  string        `yaml:"type"`
	Input []InputConfig `yaml:"input"`
}

type InputConfig struct {
	Name  string `yaml:"name"`
	Type  string `yaml:"type"`
	Value string `yaml:"value"`
}

type ExecutionTask func(*container.Container, Input) (Output, error)
