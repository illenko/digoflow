package main

import (
	"github.com/gin-gonic/gin"
	"github.com/illenko/digoflow-protorype/internal/cfg"
	"github.com/illenko/digoflow-protorype/internal/entrypoint/http"
)

func main() {

	app, err := cfg.LoadApp("flows")

	if err != nil {
		panic(err)
	}

	g := gin.Default()

	for _, f := range app.Flows {
		if f.Entrypoint.Type == "http-handler" {
			http.NewHandler(f, g)
		}
	}

	err = g.Run(":8080")
	if err != nil {
		return
	}
}
