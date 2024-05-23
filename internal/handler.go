package internal

import (
	"net/http"

	"github.com/sircodemane/yaegi-plugin/pkg/hello"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	var greeting = hello.SayHello("friend")
	w.Write([]byte(greeting))
}