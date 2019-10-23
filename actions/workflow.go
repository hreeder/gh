package actions

import (
	"gopkg.in/yaml.v2"
)

// Workflow represents the main GitHub Workflow definition
type Workflow struct {
	Name string `yaml:"name"`
	// On   []string                 `yaml:"on,flow"`
	Env  map[string]string        `yaml:"env"`
	Jobs map[string]JobDefinition `yaml:"jobs"`
}

// Parse parses a given workflow file
func (wf *Workflow) Parse(data []byte) error {
	return yaml.Unmarshal(data, &wf)
}
