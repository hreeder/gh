package actions

// JobDefinition represents a single job within a workflow
type JobDefinition struct {
	Name           string            `yaml:"name"`
	Needs          StringArray       `yaml:"needs"`
	RunsOn         string            `yaml:"runs-on"`
	Env            map[string]string `yaml:"env"`
	If             string            `yaml:"if"`
	Steps          []StepDefinition  `yaml:"steps"`
	TimeoutMinutes int               `yaml:"timeout-minutes"`
}
