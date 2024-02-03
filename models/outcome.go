package models

type Outcome struct {
	Name       string                 `yaml:"name"`
	Action     string                 `yaml:"action"`
	Mode       string                 `yaml:"mode"` // "sync" or "async"
	Parameters map[string]interface{} `yaml:"parameters"`
}
