package models

type Rule struct {
	Rule       string                   `yaml:"rule"`
	Conditions []map[string]interface{} `yaml:"conditions"`
}
