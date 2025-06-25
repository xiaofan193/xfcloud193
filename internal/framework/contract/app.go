package contract

// define str credential
const AppKey = "xf:app"

// App define interface
type App interface {
	// APPID represent the unique ID of the current app,which can be used for distributed locks ,etc
	AppID() string
	// Version define current version
	Version() string
	// BaseFolder define base address
	BaseFolder() string
	// ConfigFolder defined the path where the config located
	ConfigFolder() string
	// LogFolder defined the path where the log located
	LogFolder() string
	// ProviderFolder define the service provider address for the business itself
	ProviderFolder() string
	// MiddlewareFolder defined the middleware for the business itself
	MiddlewareFolder() string
	// CommandFolder define the command for business itself
	CommandFolder() string
	// RuntimeFolder() define the runtime status of information for business itself
	RuntimeFolder() string
	// TestFolder Store the information required for testing
	TestFolder() string
	// DeployFolder store the created folder when deployed
	DeployFolder() string

	// AppFolder defined the direcotory where the business code is located for monitoring file changes and usage
	AppFolder() string
	// LoadAppConfig load the new AppConfig,where the key is converted to a lowercase underscore for the corresponging function,
	// for example,ConfigFolder => config_folder
	LoadAppConfig(kv map[string]string)
}
