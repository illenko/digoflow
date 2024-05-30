package main

import (
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"github.com/gin-gonic/gin"
	"github.com/illenko/digoflow-protorype/internal/cfg"
	"io"
	"net/http"
)

type Execution struct {
	ID          string
	FlowID      string
	CurrentTask int
	Values      map[string]any
}

func main() {

	config, err := cfg.LoadConfig()

	if err != nil {
		panic(err)
	}

	g := gin.Default()

	for _, e := range config.Entrypoints {

		f, ok := config.Flows[e.FlowID]

		if !ok {
			fmt.Printf("Flow not found for entrypoint %s", e.ID)
			continue
		}

		if e.Config.Method == "GET" {
			fmt.Printf("Get http method is not supported")
		} else if e.Config.Method == "POST" {
			fmt.Printf("registering POST method for %s", e.Config.Path)
			g.POST(e.Config.Path, func(c *gin.Context) {

				jsonData, err := io.ReadAll(c.Request.Body)
				if err != nil {
					panic(err)
				}

				jsonParsed, err := gabs.ParseJSON(jsonData)

				values := map[string]any{}

				for _, i := range f.Input.Fields {
					fmt.Printf("Field: %s", i.Name)

					if i.Type == "FLOAT" {
						values["input."+i.Name] = jsonParsed.Path(i.Name).Data().(float64)
					} else if i.Type == "STRING" {
						values["input."+i.Name] = jsonParsed.Path(i.Name).Data().(string)
					}
				}

				execution := Execution{
					ID:          "1",
					FlowID:      e.FlowID,
					CurrentTask: 0,
					Values:      values,
				}

				for _, t := range f.Tasks {
					fmt.Printf("Task: %s", t.ID)

					if t.Type == "mapper" {
						for _, m := range t.Mappings {
							fmt.Printf("Mapping: %s -> %s", m.From, m.To)
							execution.Values["output."+t.ID+"."+m.To] = values[m.From]
						}
					}
				}

				fmt.Printf("Execution: %v", execution)

				output := gabs.New()

				for _, o := range f.Output.Body {
					output.SetP(execution.Values[o.Value], o.Name)
				}

				c.JSON(http.StatusOK, output.Data())
			})
		}

	}

	err = g.Run(":8080")
	if err != nil {
		return
	}
}
