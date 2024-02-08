package models

type Outcome struct {
	Name       string                 `json:"name" yaml:"name"`
	Action     string                 `json:"action" yaml:"action"`
	Mode       string                 `json:"mode" yaml:"mode"` // "sync" or "async"
	Parameters map[string]interface{} `json:"parameters" yaml:"parameters"`
}
