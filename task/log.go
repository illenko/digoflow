package task

import (
	"fmt"

	"github.com/illenko/digoflow/container"
)

type Log struct{}

func (t *Log) Execute(_ *container.Container, input Input) (Output, error) {
	for _, v := range input {
		fmt.Println(v)
	}

	return make(Output), nil
}
