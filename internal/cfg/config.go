package cfg

import (
	"github.com/illenko/digoflow-protorype/internal/model"
	"gopkg.in/yaml.v2"
	"os"
)

func LoadConfig() (model.App, error) {
	data, err := os.ReadFile("flows/purchase.yaml")

	if err != nil {
		return model.App{}, err
	}

	var purchaseFlow model.Flow

	err = yaml.Unmarshal(data, &purchaseFlow)

	if err != nil {
		return model.App{}, err
	}

	appData, err := os.ReadFile("entrypoints/entrypoints.yaml")

	if err != nil {
		return model.App{}, err
	}

	var entrypoints model.Entrypoints

	err = yaml.Unmarshal(appData, &entrypoints)

	if err != nil {
		return model.App{}, err
	}

	flows := make(map[string]model.Flow)
	flows[purchaseFlow.ID] = purchaseFlow

	app := model.App{
		Entrypoints: entrypoints.Values,
		Flows:       flows,
	}

	return app, nil

}
