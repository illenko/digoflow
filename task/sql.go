package task

import "github.com/illenko/digoflow/container"

type SQL struct{}

func (t *SQL) Execute(c *container.Container, input Input) (Output, error) {
	_, err := c.Database.Exec(input["query"].(string))

	if err != nil {
		return nil, err
	}

	return Output{}, nil
}
