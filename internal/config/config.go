package config

import (
	"os"

	"github.com/illenko/digoflow-protorype/internal/component"
	"gopkg.in/yaml.v2"
)

func LoadApp(flowsDir string) (*component.App, error) {
	files, err := os.ReadDir(flowsDir)
	if err != nil {
		return nil, err
	}

	app := component.NewApp()

	for _, file := range files {
		if !file.IsDir() {
			data, err := os.ReadFile(flowsDir + "/" + file.Name())
			if err != nil {
				return nil, err
			}

			var flow component.Flow
			err = yaml.Unmarshal(data, &flow)
			if err != nil {
				return nil, err
			}

			app.RegisterFlow(flow)
		}
	}

	return app, nil
}
