package actions

// ProcessTemplates : Interface to process templates
type ProcessTemplates interface {
	// ProcessTemplates : Processes the template of the given filename
	ProcessTemplates(filename string) (err error)
	// Placeholder : returns a map with the placeholder names and its values for this type
	Placeholder() map[string]string
}
