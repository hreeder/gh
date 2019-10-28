package actions

// ActionRepository represents a target of a step within a job
// It can be either a git repository (the current one, or separate), or it can be a Docker image
type ActionRepository struct {
	Owner      string
	Repository string
	Path       string
	Ref        string

	Local bool

	Docker bool
	Image  string
	Tag    string
}
