package digoflow

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/illenko/digoflow/container"
	"github.com/illenko/digoflow/task"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
)

type App struct {
	Container *container.Container
	Flows     map[string]Flow
	Tasks     map[string]task.ExecutionTask
}

func NewApp(flowsDir string, migrationsDir string) (*App, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	db, err := connectToDB(migrationsDir)
	if err != nil {
		return nil, err
	}

	files, err := filepath.Glob(filepath.Join(flowsDir, "*.yaml"))
	if err != nil {
		return nil, fmt.Errorf("error reading directory: %w", err)
	}

	app := &App{
		Container: container.NewContainer(db),
		Flows:     make(map[string]Flow),
		Tasks:     builtInTasks(),
	}

	for _, file := range files {
		flow, err := readFlow(file)
		if err != nil {
			return nil, err
		}
		app.RegisterFlow(flow)
	}

	return app, nil
}

func connectToDB(migrationsDir string) (*sql.DB, error) {
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		return nil, nil
	}

	db, err := sql.Open("postgres", fmt.Sprintf("%s?user=%s&password=%s&sslmode=disable", dbUrl, os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD")))
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	err = dbMigration(migrationsDir, db)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("error executing migrations: %w", err)
	}

	return db, nil
}

func readFlow(filePath string) (Flow, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return Flow{}, err
	}

	var flow Flow
	err = yaml.Unmarshal(data, &flow)
	if err != nil {
		return Flow{}, err
	}

	return flow, nil
}

func builtInTasks() map[string]task.ExecutionTask {
	return map[string]task.ExecutionTask{
		"digoflow.log":         task.Log,
		"digoflow.httpRequest": task.HTTPRequest,
		"digoflow.toJson":      task.ToJSON,
		"digoflow.sql":         task.SQL,
	}
}

func (a *App) RegisterFlow(flow Flow) {
	flow.Container = a.Container
	a.Flows[flow.ID] = flow
}

func (a *App) RegisterTask(taskType string, task task.ExecutionTask) {
	a.Tasks[taskType] = task
}

func (a *App) Start() error {
	g := gin.Default()

	ep, err := a.registerFlows(g)
	if err != nil {
		fmt.Printf("Error registering flows: %v", err)
		return err
	}

	if slices.Contains(ep, "http-handler") {
		err = g.Run(":8080")

		if err != nil {
			fmt.Printf("Error running server: %v", err)
			return err
		}
	}

	return nil
}

func (a *App) registerFlows(g *gin.Engine) ([]string, error) {
	entrypointTypes := make([]string, 0)
	seen := make(map[string]bool)

	for _, f := range a.Flows {
		err := a.registerTasks(&f)
		if err != nil {
			return nil, err
		}

		if f.Entrypoint.Type == "http-handler" {
			NewHttpHandler(f, g)
		}

		if _, ok := seen[f.Entrypoint.Type]; !ok {
			entrypointTypes = append(entrypointTypes, f.Entrypoint.Type)
			seen[f.Entrypoint.Type] = true
		}
	}

	return entrypointTypes, nil
}

func (a *App) registerTasks(f *Flow) error {
	for _, tc := range f.TaskConfigs {
		t, ok := a.Tasks[tc.Type]

		if !ok {
			return fmt.Errorf("task not found: %s", tc.ID)
		}

		f.ExecutionTasks = append(f.ExecutionTasks, t)
	}
	return nil
}
