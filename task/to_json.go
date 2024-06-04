package task

import (
	"github.com/Jeffail/gabs/v2"
	"github.com/illenko/digoflow/container"
)

type ToJson struct{}

func (t *ToJson) Execute(_ *container.Container, input Input) (Output, error) {
	jsonObj := gabs.New()

	for k, v := range input {
		_, _ = jsonObj.SetP(v, k)
	}

	return Output{
		"result": jsonObj.String(),
	}, nil
}
