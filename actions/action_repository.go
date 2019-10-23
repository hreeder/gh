package actions

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