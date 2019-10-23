package actions

import (
	"testing"

	"gopkg.in/yaml.v2"
)

func TestStringArrayWithSingle(t *testing.T) {
	inputYaml := "sample string"
	inputYamlBytes := []byte(inputYaml)

	sa := StringArray{}
	err := yaml.Unmarshal(inputYamlBytes, &sa)
	if err != nil {
		t.Error(err)
	}

	if len(sa) != 1 {
		t.Errorf("Expected len(sa) to equal 1, instead got %d", len(sa))
	}

	if sa[0] != inputYaml {
		t.Errorf("Expected sa[0] to equal %s, instead got %s", inputYaml, sa[0])
	}
}

func TestStringArrayWithMultiple(t *testing.T) {
	inputYaml := `["string", "two", "three"]`
	inputYamlBytes := []byte(inputYaml)

	sa := StringArray{}
	err := yaml.Unmarshal(inputYamlBytes, &sa)
	if err != nil {
		t.Error(err)
	}

	if len(sa) != 3 {
		t.Errorf("Expected len(sa) to equal 3, instead got %d", len(sa))
	}

	if sa[0] != "string" {
		t.Errorf("Expected sa[0] to equal 'string', instead got %s", sa[0])
	}

	if sa[2] != "three" {
		t.Errorf("Expected sa[2] to equal 'three', instead got %s", sa[2])
	}
}

func TestStringArrayErrorsOnMap(t *testing.T) {
	inputYaml := `{ "foo": "bar" }`
	inputYamlBytes := []byte(inputYaml)

	sa := StringArray{}
	err := yaml.Unmarshal(inputYamlBytes, &sa)
	if err == nil {
		t.Errorf("Expected to error, but did not. Instead, sa[0] = %s", sa[0])
	}
}
