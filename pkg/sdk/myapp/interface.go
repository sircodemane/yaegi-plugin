package myapp

import (
	"github.com/sircodemane/yaegi-plugin/pkg/hello"
)

type MyAppInterface struct{}

func (s MyAppInterface) SayHello(name string) string {
	return hello.SayHello(name) + " (from plugin interface!)"
}
