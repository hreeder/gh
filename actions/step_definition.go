package actions

import (
	"errors"
	"regexp"
	"strings"
)

type StepDefinition struct {
	ID               string            `yaml:"id"`
	If               string            `yaml:"if"`
	Name             string            `yaml:"name"`
	Uses             string            `yaml:"uses"`
	Run              string            `yaml:"run"`
	With             map[string]string `yaml:"with"`
	Env              map[string]string `yaml:"env"`
	WorkingDirectory string            `yaml:"working-directory"`
	ContinueOnError  bool              `yaml:"continue-on-error"`
	TimeoutMinutes   int               `yaml:"timeout-minutes"`
}

func (sd *StepDefinition) Repository() (ActionRepository, error) {
	// Local paths
	if strings.HasPrefix(sd.Uses, "./") {
		return ActionRepository{
			Path:  sd.Uses[2:], // Strip the ./
			Local: true,
		}, nil
	}

	// Docker images
	if strings.HasPrefix(sd.Uses, "docker://") {
		dockerImage := sd.Uses[len("docker://"):]
		parts := strings.Split(dockerImage, ":")
		image := parts[0]
		tag := "master"
		if len(parts) == 2 {
			tag = parts[1]
		}
		return ActionRepository{
			Docker: true,
			Image:  image,
			Tag:    tag,
		}, nil
	}

	githubRE := regexp.MustCompile(`^(?P<owner>[a-zA-Z-_]*)/(?P<repo>[a-zA-Z-_]*)(/(?P<path>.*))?@(?P<ref>\w*)$`)
	match := githubRE.FindStringSubmatch(sd.Uses)
	if len(match) == 0 {
		return ActionRepository{}, errors.New("Unable to parse")
	}

	result := make(map[string]string)
	for i, name := range githubRE.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	repo := ActionRepository{
		Owner:      result["owner"],
		Repository: result["repo"],
		Ref:        result["ref"],
	}

	if result["path"] != "" {
		repo.Path = result["path"]
	}

	return repo, nil
}
