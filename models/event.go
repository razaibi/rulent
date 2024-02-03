package models

type Event struct {
	RulesOperator string    `yaml:"rules-operator"`
	Outcomes      []Outcome `yaml:"outcomes"`
	Rules         []Rule    `yaml:"rules"`
}
