package http

import (
	"fmt"
	"io"
	"net/http"

	"github.com/illenko/digoflow-protorype/internal/component/task"

	"github.com/Jeffail/gabs/v2"
	"github.com/gin-gonic/gin"
	"github.com/illenko/digoflow-protorype/internal/component"
)

func NewHandler(f component.Flow, g *gin.Engine) {
	config := f.Entrypoint.Config
	fmt.Printf("registering HTTP entrypoint for %s \n", config["path"])

	switch config["method"] {
	case "GET":
		g.GET(config["path"], handleRequest(f, false))
	case "POST":
		g.POST(config["path"], handleRequest(f, true))
	default:
		fmt.Printf("Method %s is not supported", config["method"])
	}
}

func handleRequest(f component.Flow, body bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		e := component.NewExecution(f.ID)

		handlePathVariables(c, f, &e)
		handleQueryParameters(c, f, &e)
		handleHeaders(c, f, &e)

		if body {
			handleBody(c, f, &e)
		}

		output, err := component.ExecuteTasks(f, &e)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error in task execution: " + err.Error(),
			})
			return
		}

		toResponse(c, output)
	}
}

func toResponse(c *gin.Context, output task.Output) {
	jsonObj := gabs.New()

	for k, v := range output {
		_, _ = jsonObj.SetP(v, k)
	}

	c.JSON(http.StatusOK, jsonObj.Data())
}

func handlePathVariables(c *gin.Context, f component.Flow, e *component.Execution) {
	for _, i := range f.Input.PathVariables {
		e.Values["input.path-variables."+i.Name] = c.Param(i.Name)
	}
}

func handleQueryParameters(c *gin.Context, f component.Flow, e *component.Execution) {
	for _, q := range f.Input.QueryParameters {
		e.Values["input.query-parameters."+q.Name] = c.Query(q.Name)
	}
}

func handleHeaders(c *gin.Context, f component.Flow, e *component.Execution) {
	for _, h := range f.Input.Headers {
		e.Values["input.headers."+h.Name] = c.GetHeader(h.Name)
	}
}

func handleBody(c *gin.Context, f component.Flow, e *component.Execution) {
	if f.Input.Body.Type == "json" {
		handleJSONBody(c, f, e)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Body type is not supported",
		})
	}
}

func handleJSONBody(c *gin.Context, f component.Flow, e *component.Execution) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong request body format",
		})
		return
	}

	bodyParsed, err := gabs.ParseJSON(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong request body format",
		})
		return
	}

	for _, i := range f.Input.Body.Fields {
		if i.Type == "float" {
			e.Values["input.body."+i.Name] = bodyParsed.Path(i.Name).Data().(float64)
		} else if i.Type == "string" {
			e.Values["input.body."+i.Name] = bodyParsed.Path(i.Name).Data().(string)
		}
	}
}
