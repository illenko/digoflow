package task

import (
	"fmt"
)

func Log(values Input) (Output, error) {

	for _, v := range values {
		fmt.Println(v)
	}

	return make(Output), nil
}
