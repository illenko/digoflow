package model

type Flow struct {
	ID          string `yaml:"id"`
	Name        string `yaml:"name"`
	Input       Input  `yaml:"input"`
	CurrentTask int
	Tasks       []MappingTask `yaml:"tasks"`
	Output      Output        `yaml:"output"`
}

type Input struct {
	Type   string           `yaml:"type"`
	Fields []InputBodyField `yaml:"fields"`
}

type InputBodyField struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

type Mapping struct {
	From string `yaml:"from"`
	To   string `yaml:"to"`
}

type MappingTask struct {
	ID       string    `yaml:"id"`
	Type     string    `yaml:"type"`
	Mappings []Mapping `yaml:"config"`
}

type Output struct {
	Type string        `yaml:"type"`
	Body []OutputField `yaml:"body"`
}

type OutputField struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}
