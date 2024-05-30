package model

type App struct {
	Entrypoints []Entrypoint
	Flows       map[string]Flow
}

type Entrypoints struct {
	Values []Entrypoint `yaml:"entrypoints"`
}
