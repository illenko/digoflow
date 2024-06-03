package task

import (
	"github.com/Jeffail/gabs/v2"
	"github.com/illenko/digoflow/container"
)

func ToJSON(_ *container.Container, input Input) (Output, error) {

	jsonObj := gabs.New()

	for k, v := range input {
		_, _ = jsonObj.SetP(v, k)
	}

	return Output{
		"result": jsonObj.String(),
	}, nil
}
