package task

import (
	"fmt"
	"github.com/illenko/digoflow/container"
)

func Log(_ *container.Container, values Input) (Output, error) {

	for _, v := range values {
		fmt.Println(v)
	}

	return make(Output), nil
}
