package actions

import "testing"

func TestGetDependencyList(t *testing.T) {
	sd1 := StepDefinition{
		Uses: "monacorp/action-name@master",
	}
	sd2 := StepDefinition{
		Uses: "monacorp/another-action@master",
	}
	jobs := make(map[string]JobDefinition)
	jobs["demo"] = JobDefinition{
		Steps: []StepDefinition{sd1},
	}
	jobs["demo2"] = JobDefinition{
		Steps: []StepDefinition{sd2},
	}
	wf := Workflow{
		Jobs: jobs,
	}

	dependencies := wf.ResolveDependencies()

	if len(dependencies) != 2 {
		t.Errorf("Expected 2 dependencies, but got %d", len(dependencies))
	}

	sd1_get := false
	sd2_get := false
	for _, repo := range dependencies {
		if repo.Owner == "monacorp" && repo.Repository == "action-name" && repo.Ref == "master" {
			sd1_get = true
		}

		if repo.Owner == "monacorp" && repo.Repository == "another-action" && repo.Ref == "master" {
			sd2_get = true
		}
	}

	if sd1_get == false {
		t.Error("monacorp/action-name@master was not properly resolved (SD1)")
	}

	if sd2_get == false {
		t.Error("monacorp/another-action@master was not properly resolved (SD2)")
	}
}
