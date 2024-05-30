package model

type Entrypoint struct {
	ID     string `yaml:"id"`
	Type   string `yaml:"type"`
	FlowID string `yaml:"flowId"`
	Config Config `yaml:"config"`
}

type Config struct {
	Method string `yaml:"method"`
	Path   string `yaml:"path"`
}
