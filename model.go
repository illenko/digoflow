package main

type App struct {
}

type Config struct {
	Method string `yaml:"method"`
	Path   string `yaml:"path"`
}

type Entrypoint struct {
	ID     string `yaml:"id"`
	Type   string `yaml:"type"`
	FlowID string `yaml:"flowId"`
	Config Config `yaml:"config"`
}

type BodyField struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

type Input struct {
	Type string      `yaml:"type"`
	Body []BodyField `yaml:"body"`
}

type MapperConfig struct {
	Mappings []Mapping `yaml:"mappings"`
}

type Mapping struct {
	From string `yaml:"from"`
	To   string `yaml:"to"`
}

type Task struct {
	ID     string   `yaml:"id"`
	Type   string   `yaml:"type"`
	Config []Config `yaml:"config"`
}

type OutputField struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type Output struct {
	Type string        `yaml:"type"`
	Body []OutputField `yaml:"body"`
}

type Flow struct {
	ID     string `yaml:"id"`
	Name   string `yaml:"name"`
	Input  Input  `yaml:"input"`
	Tasks  []Task `yaml:"tasks"`
	Output Output `yaml:"output"`
}
