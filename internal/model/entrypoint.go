package model

type Entrypoint struct {
	Type   string            `yaml:"type"`
	Config map[string]string `yaml:"config"`
}
