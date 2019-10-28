package actions

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// DiscoverJob finds the job, given the Name
func DiscoverJob(rawJobName string) (*JobDefinition, error) {
	// Load all workflows
	workflows := make(map[string]*Workflow)

	wfDir, err := getWorkflowDirectory()
	if err != nil {
		return nil, err
	}

	// Parse all workflows
	err = filepath.Walk(wfDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			file, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			wf := &Workflow{}
			wf.Parse(file)

			wfNameParts := strings.Split(info.Name(), ".")

			workflows[wfNameParts[0]] = wf
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Handle <workflow>.<job> syntax
	if strings.Contains(rawJobName, ".") {
		rawJobNameParts := strings.Split(rawJobName, ".")
		wfName := rawJobNameParts[0]

		if _, ok := workflows[wfName]; ok {
			jobName := rawJobNameParts[1]
			if _, ok = workflows[wfName].Jobs[jobName]; ok {
				job := workflows[wfName].Jobs[jobName]
				return &job, nil
			}

			fmt.Printf("Job '%s' not found in workflow '%s'. Please check your workflow definitions.\n", jobName, wfName)
			return nil, errors.New("job not found in defined workflow")
		}

		fmt.Printf("Workflow '%s' not found. Please check your workflow definitions.\n", wfName)
		return nil, errors.New("workflow not found")
	}

	// Find the given rawJobName
	var jobFoundInWorkflow string
	for workflowName, workflow := range workflows {
		if _, ok := workflow.Jobs[rawJobName]; ok {
			if jobFoundInWorkflow != "" {
				fmt.Println("Job found in multiple workflows - Please use <workflow>.<job> syntax instead.")
				return nil, errors.New("job found in multiple workflows")
			}
			jobFoundInWorkflow = workflowName
		}
	}

	if jobFoundInWorkflow == "" {
		fmt.Println("Cannot find the given job. Please check your syntax.")
		return nil, errors.New("job not found")
	}

	job := workflows[jobFoundInWorkflow].Jobs[rawJobName]
	return &job, nil
}

func getWorkflowDirectory() (string, error) {
	// workflows reside in `<repo root>/.github/workflows`
	// This could be run anywhere inside the repo.
	// Start at the current directory and go up the tree to
	// locate the .github/workflows directory

	starting, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return getWorkflowDirectoryRecurse(starting)
}

func getWorkflowDirectoryRecurse(dir string) (string, error) {
	gitDirectory := path.Join(dir, ".git")
	if _, err := os.Stat(gitDirectory); os.IsNotExist(err) {
		next := path.Join(dir, "..")
		if next == dir {
			return "", errors.New("reached root of fs, not in a git repo")
		}
		return getWorkflowDirectoryRecurse(path.Join(dir, ".."))
	}

	wfDir := path.Join(dir, ".github", "workflows")
	if _, err := os.Stat(wfDir); os.IsNotExist(err) {
		return "", errors.New("no .github/workflows directory")
	}
	return wfDir, nil
}
