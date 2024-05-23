package lib

import "reflect"

//go:generate go run github.com/traefik/yaegi/cmd/yaegi@v0.16.1 extract github.com/go-chi/chi/v5/middleware
//go:generate go run github.com/traefik/yaegi/cmd/yaegi@v0.16.1 extract github.com/sircodemane/yaegi-plugin/pkg/sdk
//go:generate go run github.com/traefik/yaegi/cmd/yaegi@v0.16.1 extract github.com/sircodemane/yaegi-plugin/pkg/sdk/myapp

var Symbols = map[string]map[string]reflect.Value{}
