package apiserver

import "github.com/spf13/cobra"

// App is the main structure of a cli application
// It is recommeded that an app be created with the app.NewApp() function()
type App struct {
	baseName    string
	name        string
	description string
	options     CliOptions
	funFunc     RunFunc
	silence     bool
	noVersion   bool
	//commands    []*commands
	args cobra.PositionalArgs
	cmd  *cobra.Command
}

// Option defines optional parameters for initializing the application
// structure.
type Option func(*App)

// RunFunc defines the application`s startup callback function.
type RunFunc func(basename string) error
