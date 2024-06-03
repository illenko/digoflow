package task

import "github.com/illenko/digoflow/container"

func SQL(c *container.Container, input Input) (Output, error) {
	db := c.GetDB()

	_, err := db.Exec(input["query"].(string))

	if err != nil {
		return nil, err
	}

	return Output{}, nil
}
