package sdk

import (
	"net/http"

	"github.com/sircodemane/yaegi-plugin/pkg/sdk/myapp"
)

type Plugin interface {
	Routes() map[string]http.HandlerFunc
}

type PluginConstructor func(appInterface *myapp.MyAppInterface) Plugin
