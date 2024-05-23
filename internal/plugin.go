package internal

import (
	"github.com/sircodemane/yaegi-plugin/pkg/sdk"
	"github.com/sircodemane/yaegi-plugin/pkg/sdk/lib"
	"github.com/sircodemane/yaegi-plugin/pkg/sdk/myapp"
	"github.com/sircodemane/yaegi-plugin/plugins"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

func LoadPlugin() (sdk.Plugin, error) {
	i := interp.New(interp.Options{})

	i.Use(stdlib.Symbols)
	i.Use(lib.Symbols)

	src, err := plugins.MyPluginFS.ReadFile("myplugin/plugin.go")
	if err != nil {
		return nil, err
	}

	_, err = i.Eval(string(src))
	if err != nil {
		return nil, err
	}

	v, err := i.Eval("myplugin.NewPlugin")
	if err != nil {
		return nil, err
	}

	newPlugin := v.Interface().(sdk.PluginConstructor)
	plugin := newPlugin(&myapp.MyAppInterface{})
	return plugin, nil
}
