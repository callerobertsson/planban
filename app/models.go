// Package app model definitions
package app

// Board defines a Planban Board
type Board struct {
	Name        string
	Information string
	Stacks      []Stack
	Config      Config
}

// Config defines the configuration setting of a Board
type Config struct {
	NoColors             bool
	HideBoardInformation bool
	HideTaskDescriptions bool
}

// Stack defines a Planban Stack of Tasks
type Stack struct {
	Name  string
	Tasks []Task
}

// Task defines the contents of a Task
type Task struct {
	Name        string
	Description string
}
