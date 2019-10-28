package actions

// ResolveDependencies returns all ActionRepositories within the given workflow
func (wf *Workflow) ResolveDependencies() []*ActionRepository {
	var resolved []*ActionRepository

	for _, jobDef := range wf.Jobs {
		for _, step := range jobDef.Steps {
			repo, _ := step.Repository()
			resolved = append(resolved, &repo)
		}
	}

	return resolved
}
