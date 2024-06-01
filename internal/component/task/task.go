package task

type Input map[string]any

type Output map[string]any

type Config struct {
	ID    string       `yaml:"id"`
	Name  string       `yaml:"name"`
	Type  string       `yaml:"type"`
	Input []InputValue `yaml:"input"`
}

type InputValue struct {
	Name  string `yaml:"name"`
	Type  string `yaml:"type"`
	Value string `yaml:"value"`
}

type ExecutionTask func(Input) (Output, error)
