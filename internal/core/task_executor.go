package core

import (
	"fmt"
	"github.com/illenko/digoflow-protorype/internal/core/expression"
	"github.com/illenko/digoflow-protorype/internal/task"
	"strconv"
)

func ExecuteTasks(f Flow, e *Execution) (task.Output, error) {
	for i, t := range f.TaskConfigs {
		taskInput, err := createTaskInput(t, e)
		if err != nil {
			return nil, err
		}

		output, err := executeTask(&taskInput, f.ExecutionTasks[i])
		if err != nil {
			return nil, err
		}

		for k, v := range output {
			e.Values["output."+t.ID+"."+k] = v
		}
	}

	output := make(task.Output)

	return output, nil
}

func createTaskInput(t task.Config, e *Execution) (task.Input, error) {
	taskInput := make(task.Input)

	for _, inp := range t.Input {
		placeholders := expression.GetPlaceholders(inp.Value)

		replacement := map[string]string{}

		for _, p := range placeholders {
			value, ok := e.Values[p]
			if !ok {
				return nil, fmt.Errorf("placeholder not found")
			}

			if inp.Type == "float" {
				value = strconv.FormatFloat(value.(float64), 'f', -1, 64)
			}

			replacement[p] = value.(string)
		}

		realValue := expression.ReplacePlaceholders(inp.Value, replacement)

		taskInput[inp.Name] = realValue
	}

	return taskInput, nil
}

func executeTask(taskInput *task.Input, task task.ExecutionTask) (task.Output, error) {
	return task(*taskInput)
}