package actions

import (
	"testing"

	"gopkg.in/yaml.v2"
)

func TestWorkflowEventParsesSingleString(t *testing.T) {
	inputYaml := "push"
	inputYamlBytes := []byte(inputYaml)

	workflowEvent := WorkflowEvent{}
	err := yaml.Unmarshal(inputYamlBytes, &workflowEvent)
	if err != nil {
		t.Error(err)
	}

	if workflowEvent.Name != "push" {
		t.Errorf("Expected 'push' but got '%s'", workflowEvent.Name)
	}
}

func TestWorkflowEventParsesMapping(t *testing.T) {
	inputYaml := `push:
  branches:
  - master
  - 'mona/octocat'`
	inputYamlBytes := []byte(inputYaml)

	workflowEvent := WorkflowEvent{}
	err := yaml.Unmarshal(inputYamlBytes, &workflowEvent)
	if err != nil {
		t.Error(err)
	}

	if workflowEvent.Name != "push" {
		t.Errorf("Expected 'push' but got '%s'", workflowEvent.Name)
	}
}
