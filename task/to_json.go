package task

import (
	"github.com/Jeffail/gabs/v2"
)

func ToJSON(input Input) (Output, error) {

	jsonObj := gabs.New()

	for k, v := range input {
		_, _ = jsonObj.SetP(v, k)
	}

	return Output{
		"result": jsonObj.String(),
	}, nil
}
