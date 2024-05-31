package http

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Jeffail/gabs/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/illenko/digoflow-protorype/internal/model"
)

func NewHandler(f model.Flow, g *gin.Engine) {
	config := f.Entrypoint.Config
	fmt.Printf("registering HTTP entrypoint for %s \n", config["path"])

	if config["method"] == "GET" {
		fmt.Printf("Get http method is not supported")
	} else if config["method"] == "POST" {
		g.POST(config["path"], func(c *gin.Context) {
			e := model.Execution{
				ID:          uuid.New().String(),
				FlowID:      f.ID,
				CurrentTask: 0,
				Values:      map[string]any{},
			}

			for _, i := range f.Input.PathVariables {
				fmt.Printf("Path variable: %s \n", i.Name)
				e.Values["input.path-variables."+i.Name] = c.Param(i.Name)
			}

			for _, q := range f.Input.QueryParameters {
				fmt.Printf("Query parameter: %s \n", q.Name)
				e.Values["input.query-parameters."+q.Name] = c.Query(q.Name)
			}

			for _, h := range f.Input.Headers {
				fmt.Printf("Header: %s \n", h.Name)
				e.Values["input.headers."+h.Name] = c.GetHeader(h.Name)
			}

			if f.Input.Body.Type == "json" {
				fmt.Printf("Parsing JSON body \n")

				body, err := io.ReadAll(c.Request.Body)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"message": "Wrong request body format",
					})
				}

				bodyParsed, err := gabs.ParseJSON(body)

				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"message": "Wrong request body format",
					})
				}

				for _, i := range f.Input.Body.Fields {
					fmt.Printf("Body Field: %s \n", i.Name)

					if i.Type == "float" {
						e.Values["input.body."+i.Name] = bodyParsed.Path(i.Name).Data().(float64)
					} else if i.Type == "string" {
						e.Values["input.body."+i.Name] = bodyParsed.Path(i.Name).Data().(string)
					}
				}
			} else {
				fmt.Printf("Body type %s is not supported \n", f.Input.Body.Type)
			}

			fmt.Printf("Started execution of flow %s \n", f.ID)
			c.JSON(http.StatusOK, e)
		})
	} else {
		fmt.Printf("Method %s is not supported", config["method"])
	}
}
