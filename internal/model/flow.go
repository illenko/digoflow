package model

type Flow struct {
	ID         string     `yaml:"id"`
	Name       string     `yaml:"name"`
	Entrypoint Entrypoint `yaml:"entrypoint"`
	Input      HttpInput  `yaml:"input"`
}

type HttpInput struct {
	PathVariables   []Variable `yaml:"path-variables"`
	QueryParameters []Variable `yaml:"query-parameters"`
	Headers         []Variable `yaml:"headers"`
	Body            Body       `yaml:"body"`
}

type Variable struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

type Body struct {
	Type   string     `yaml:"type"`
	Fields []Variable `yaml:"fields"`
}
