package actions

// WorkflowEvent represents an individual workflow event
type WorkflowEvent struct {
	Name string
}

// UnmarshalYAML overrides the default unmarshalling to handle the dynamic type abilities of the GH actions
func (wfe *WorkflowEvent) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var single string
	err := unmarshal(&single)
	if err != nil {
		var mapping map[string]interface{}
		err = unmarshal(&mapping)

		if err != nil {
			return err
		}
		for k := range mapping {
			wfe.Name = k
		}
		return nil
	}

	wfe.Name = single
	return nil
}
