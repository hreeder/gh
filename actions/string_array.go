package actions

// Thanks to https://github.com/go-yaml/yaml/issues/100#issuecomment-324964723

// StringArray is a type to hold a YAML value which can be a string or a slice of strings
type StringArray []string

// UnmarshalYAML loads the YAML into the native Go type
func (a *StringArray) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var multi []string
	err := unmarshal(&multi)
	if err != nil {
		var single string
		err := unmarshal(&single)
		if err != nil {
			return err
		}
		*a = []string{single}
	} else {
		*a = multi
	}
	return nil
}
