package actions

import (
	"testing"
)

func TestGetRepositoryMaster(t *testing.T) {
	step := StepDefinition{
		Uses: "monacorp/action-name@master",
	}

	repo, err := step.Repository()
	if err != nil {
		t.Error(err)
	}

	expected := ActionRepository{
		Owner:      "monacorp",
		Repository: "action-name",
		Ref:        "master",
		Local:      false,
	}

	if repo.Owner != expected.Owner {
		t.Errorf("Wrong repo owner obtained. Expected '%s' but got '%s'", expected.Owner, repo.Owner)
	}

	if repo.Repository != expected.Repository {
		t.Errorf("Wrong repo obtained. Expected '%s' but got '%s'", expected.Repository, repo.Repository)
	}

	if repo.Ref != expected.Ref {
		t.Errorf("Wrong repo ref obtained. Expected '%s' but got '%s'", expected.Ref, repo.Ref)
	}

	if repo.Local != expected.Local {
		t.Errorf("Expected obtained repo to be Local=false, but it was not!")
	}
}

func TestGetRepositoryWithPath(t *testing.T) {
	step := StepDefinition{
		Uses: "monacorp/action-name/path@master",
	}

	repo, err := step.Repository()
	if err != nil {
		t.Error(err)
	}

	if repo.Path != "path" {
		t.Errorf("Wrong path obtained. Expected 'path' but got '%s'", repo.Path)
	}
}

func TestGetRepositoryWithMultilevelPath(t *testing.T) {
	step := StepDefinition{
		Uses: "monacorp/action-name/path/to/action@master",
	}

	repo, err := step.Repository()
	if err != nil {
		t.Error(err)
	}

	if repo.Path != "path/to/action" {
		t.Errorf("Wrong path obtained. Expected 'path/to/action' but got '%s'", repo.Path)
	}
}

func TestGetLocalActionPath(t *testing.T) {
	step := StepDefinition{
		Uses: "./path/to/action",
	}

	repo, err := step.Repository()
	if err != nil {
		t.Error(err)
	}

	if repo.Path != "path/to/action" {
		t.Errorf("Wrong path obtained. Expected 'path/to/action' but got '%s'", repo.Path)
	}

	if repo.Local != true {
		t.Errorf("Expected repo.Local to be true, but it was not")
	}
}

func TestGetDockerAction(t *testing.T) {
	step := StepDefinition{
		Uses: "docker://alpine:3.8",
	}

	repo, err := step.Repository()
	if err != nil {
		t.Error(err)
	}

	if repo.Docker != true {
		t.Errorf("Expected repo.Docker to be true, but it was not")
	}

	if repo.Image != "alpine" {
		t.Errorf("Expected repo.Image to be 'alpine', but got '%s'", repo.Image)
	}

	if repo.Tag != "3.8" {
		t.Errorf("Expected repo.Tag to be '3.8' but got '%s'", repo.Tag)
	}
}

func TestInvalidParsing(t *testing.T) {
	step := StepDefinition{
		Uses: "invalid!",
	}

	_, err := step.Repository()
	if err == nil {
		t.Error("Expected to fail but no error thrown")
	}
}
