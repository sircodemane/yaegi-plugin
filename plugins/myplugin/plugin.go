package myplugin

import (
	"net/http"

	"github.com/sircodemane/yaegi-plugin/pkg/sdk"
	"github.com/sircodemane/yaegi-plugin/pkg/sdk/myapp"
)

type MyPlugin struct {
	app *myapp.MyAppInterface
}

func NewMyPlugin(app *myapp.MyAppInterface) sdk.Plugin {
	return &MyPlugin{app: app}
}

func (s *MyPlugin) Routes() map[string]http.HandlerFunc {
	return map[string]http.HandlerFunc{
		"/":      s.HelloWorld,
		"/greet": s.Greet,
	}
}

func (s *MyPlugin) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World, from a plugin!"))
}

func (s *MyPlugin) Greet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(s.app.SayHello("plugin pal")))
}

var NewPlugin sdk.PluginConstructor = NewMyPlugin