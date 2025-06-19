package app

import (
	"errors"
	"github.com/google/uuid"
	"github.com/xiaofan193/xifancloud193/internal/framework"
	"github.com/xiaofan193/xifancloud193/internal/pkg/systemutil"
	"os"
	"path/filepath"
	"strings"
)

// App Impementiontation Represnting the Framework
type XfApp struct {
	container  framework.Container // service container
	baseFolder string              //  base route path
	appId      string              //  Indicates the unique ID for the current app

	configMap map[string]string
	envMap    map[string]string
	argsMap   map[string]string
}

// AppID indicates the unique ID
func (app *XfApp) AppID() string {
	return app.appId
}

// Version
func (app *XfApp) Version() string {
	return XfVersion
}

// BaseFolder Represents the basic directory,which can represent the directory of the development
// scenario or the directory at runtime
func (app *XfApp) BaseFolder() string {
	if app.baseFolder != "" {
		return app.baseFolder
	}
	baseFolder := app.getConfigBySequence("base_folder", "BASE_FOLDER", "app.path.base_folder")
	if baseFolder != "" {
		return baseFolder
	}
	// if no parameters ,use default path route
	return systemutil.GetExecDirectory()
}

// Indicate the configuration file address
func (app *XfApp) ConfigFolder() string {
	val := app.getConfigBySequence("config_folder", "CONFIG_FOLDER", "app.path.config_folder")
	if val != "" {
		return val
	}
	return filepath.Join(app.BaseFolder(), "config")
}

// LogFolder Indicate log storage address
func (app *XfApp) LogFolder() string {
	val := app.getConfigBySequence("log_folder", "LOG_FOLDER", "app.path.log_folder")
	if val != "" {
		return val
	}
	return filepath.Join(app.StorageFolder(), "log")
}

func (app *XfApp) StorageFolder() string {
	val := app.getConfigBySequence("storage_folder", "STORAGE_FOLDER", "app.path.storage_folder")
	if val != "" {
		return val
	}
	return filepath.Join(app.BaseFolder(), "storage")
}

func (app *XfApp) HttpFolder() string {
	val := app.getConfigBySequence("http_folder", "HTTP_FOLDER", "app.path.http_folder")
	if val != "" {
		return val
	}
	return filepath.Join(app.BaseFolder(), "app", "console")
}

func (app *XfApp) ConsoleFolder() string {
	val := app.getConfigBySequence("console_folder", "CONSOLE_FOLDER", "app.path.console_folder")
	if val != "" {
		return val
	}
	return filepath.Join(app.BaseFolder(), "app", "console")
}

// define business provider address
func (app *XfApp) ProviderFolder() string {
	val := app.getConfigBySequence("provider_folder", "PROVIDER_FOLDER", "app.path.provider_folder")
	if val != "" {
		return val
	}
	return filepath.Join(app.BaseFolder(), "app", "provider")
}

// MiddlewareFolder define self middleware
func (app *XfApp) MiddlewareFolder() string {
	val := app.getConfigBySequence("middleware_folder", "MIDDLEWARE_FOLDER", "app.path.middleware_folder")
	if val != "" {
		return val
	}
	return filepath.Join(app.ConsoleFolder(), "middleware")
}

// define business command
func (app *XfApp) CommandFolder() string {
	val := app.getConfigBySequence("command_folder", "COMMAND_FOLDER", "app.path.command_folder")
	if val != "" {
		return val
	}
	return filepath.Join(app.ConsoleFolder(), "command")
}

// RuntimeFolder define business Runtime status info
func (app *XfApp) RuntimeFolder() string {
	val := app.getConfigBySequence("runtime_folder", "RUNTIME_FOLDER", "app.path.runtime_folder")
	if val != "" {
		return val
	}
	return filepath.Join(app.StorageFolder(), "runtime")
}

// define test require infomation
func (app *XfApp) TestFolder() string {
	val := app.getConfigBySequence("test_folder", "TEST_FOLDER", "app.path.test_folder")
	if val != "" {
		return val
	}
	return filepath.Join(app.BaseFolder(), "test")
}

// initialization XfApp
func NewXfApp(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}
	// two param one is container another baseFolder
	container := params[0].(framework.Container)
	baseFolder := params[1].(string)

	appId := uuid.New().String()
	configMap := map[string]string{}
	xfApp := &XfApp{baseFolder: baseFolder, container: container, appId: appId, configMap: configMap}
	_ = xfApp.loadEnvMaps()
	_ = xfApp.loadArgsMaps()
	return xfApp, nil
}

// define deploy required infomation
func (app *XfApp) DeployFolder() string {
	val := app.getConfigBySequence("deploy_folder", "DEPLOY_FOLDER", "app.path.deploy_folder")
	if val != "" {
		return val
	}
	return filepath.Join(app.BaseFolder(), "deploy")
}

// AppFolder indicates app folder
func (app *XfApp) AppFolder() string {
	val := app.getConfigBySequence("app_folder", "APP_FOLDER", "app.path.app_folder")
	if val != "" {
		return val
	}
	return filepath.Join(app.BaseFolder(), "app")
}

// get default config method
//
//	parameters > Enviroment Variables > Configuration Files
func (app *XfApp) getConfigBySequence(argsKey string, envKey string, configKey string) string {
	if app.argsMap != nil {
		if val, ok := app.argsMap[argsKey]; ok {
			return val
		}
	}

	if app.envMap != nil {
		if val, ok := app.envMap[envKey]; ok {
			return val
		}
	}

	if app.configMap != nil {
		if val, ok := app.configMap[configKey]; ok {
			return val
		}
	}
	return ""
}

func (app *XfApp) loadEnvMaps() error {
	if app.envMap == nil {
		app.envMap = map[string]string{}
	}
	// Read environment variables
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		app.envMap[pair[0]] = pair[1]
	}
	return nil
}

func (app *XfApp) loadArgsMaps() error {
	if app.argsMap == nil {
		app.argsMap = map[string]string{}
	}

	// load args,must be formate : --key=value
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "--") {
			pair := strings.SplitN(strings.TrimPrefix(arg, "--"), "=", 2)
			if len(pair) == 2 {
				app.argsMap[pair[0]] = pair[1]
			}
		}
	}
	return nil
}

// load config map
func (app *XfApp) LoadAppConfig(kv map[string]string) {
	for key, val := range kv {
		app.configMap[key] = val
	}
}
