package cfg

import (
	"os"

	"github.com/illenko/digoflow-protorype/internal/model"
	"gopkg.in/yaml.v2"
)

func LoadApp(flowsDir string) (model.App, error) {
	data, err := os.ReadFile(flowsDir + "/purchase.yaml")

	if err != nil {
		return model.App{}, err
	}

	var purchaseFlow model.Flow

	err = yaml.Unmarshal(data, &purchaseFlow)

	if err != nil {
		return model.App{}, err
	}

	flows := make(map[string]model.Flow)
	flows[purchaseFlow.ID] = purchaseFlow

	app := model.App{
		Flows: flows,
	}

	return app, nil

}
