package models

type Event struct {
	RulesOperator string    `json:"rules-operator" yaml:"rules-operator"`
	Outcomes      []Outcome `json:"outcomes" yaml:"outcomes"`
	Rules         []Rule    `json:"rules" yaml:"rules"`
}
