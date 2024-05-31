package cfg

import (
	"os"

	"github.com/illenko/digoflow-protorype/internal/model"
	"gopkg.in/yaml.v2"
)

func LoadApp(flowsDir string) (model.App, error) {
	files, err := os.ReadDir(flowsDir)
	if err != nil {
		return model.App{}, err
	}

	flows := make(map[string]model.Flow)

	for _, file := range files {
		if !file.IsDir() {
			data, err := os.ReadFile(flowsDir + "/" + file.Name())
			if err != nil {
				return model.App{}, err
			}

			var flow model.Flow
			err = yaml.Unmarshal(data, &flow)
			if err != nil {
				return model.App{}, err
			}

			flows[flow.ID] = flow
		}
	}

	app := model.App{
		Flows: flows,
	}

	return app, nil
}
