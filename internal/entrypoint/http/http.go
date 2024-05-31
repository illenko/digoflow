package http

import (
	"fmt"
	"io"
	"net/http"

	"github.com/illenko/digoflow-protorype/internal/expression"

	"github.com/Jeffail/gabs/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/illenko/digoflow-protorype/internal/model"
)

func NewHandler(f model.Flow, g *gin.Engine) {
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

func handleRequest(f model.Flow, body bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		e := createNewExecution(f)

		handlePathVariables(c, f, &e)
		handleQueryParameters(c, f, &e)
		handleHeaders(c, f, &e)

		if body {
			handleBody(c, f, &e)
		}

		err := handleTasks(f, &e)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error in task execution: " + err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, e)
	}
}

func createNewExecution(f model.Flow) model.Execution {
	return model.Execution{
		ID:     uuid.New().String(),
		FlowID: f.ID,
		Values: map[string]any{},
	}
}

func handleTasks(f model.Flow, e *model.Execution) error {
	for i, t := range f.Tasks {
		taskInput, err := createTaskInput(t, e)
		if err != nil {
			return err
		}

		err = executeTask(f, i, t, taskInput, e)
		if err != nil {
			return err
		}
	}
	return nil
}

func createTaskInput(t model.TaskConfig, e *model.Execution) (model.TaskInput, error) {
	taskInput := make(model.TaskInput)

	for _, inp := range t.Input {
		placeholders := expression.GetPlaceholders(inp.Value)

		replacement := map[string]string{}

		for _, p := range placeholders {
			value, ok := e.Values[p]
			if !ok {
				return nil, fmt.Errorf("placeholder not found")
			}

			if inp.Type == "float" {
				value = fmt.Sprintf("%f", value)
			}

			replacement[p] = value.(string)
		}

		realValue := expression.ReplacePlaceholders(inp.Value, replacement)

		taskInput[inp.Name] = realValue
	}

	return taskInput, nil
}

func executeTask(f model.Flow, i int, t model.TaskConfig, taskInput model.TaskInput, e *model.Execution) error {
	executionTask := f.ExecutionTasks[i]

	taskOutput := executionTask(taskInput)

	for k, v := range taskOutput {
		e.Values["outputs."+t.ID+"."+k] = v
	}

	return nil
}

func handlePathVariables(c *gin.Context, f model.Flow, e *model.Execution) {
	for _, i := range f.Input.PathVariables {
		e.Values["input.path-variables."+i.Name] = c.Param(i.Name)
	}
}

func handleQueryParameters(c *gin.Context, f model.Flow, e *model.Execution) {
	for _, q := range f.Input.QueryParameters {
		e.Values["input.query-parameters."+q.Name] = c.Query(q.Name)
	}
}

func handleHeaders(c *gin.Context, f model.Flow, e *model.Execution) {
	for _, h := range f.Input.Headers {
		e.Values["input.headers."+h.Name] = c.GetHeader(h.Name)
	}
}

func handleBody(c *gin.Context, f model.Flow, e *model.Execution) {
	if f.Input.Body.Type == "json" {
		handleJSONBody(c, f, e)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Body type is not supported",
		})
	}
}

func handleJSONBody(c *gin.Context, f model.Flow, e *model.Execution) {
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
