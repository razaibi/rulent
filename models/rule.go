package models

type Rule struct {
	Rule       string                   `json:"rule" yaml:"rule"`
	Conditions []map[string]interface{} `json:"conditions" yaml:"conditions"`
}
