# yaegi-plugin

This is a simple example repo of how https://github.com/traefik/yaegi
could be used to build a pluggable application.

MyApp is an HTTP server that only has two endpoints by default:
- `GET /`
- `GET /greet`

Through the use of `yaegi` and some plugin code, the server is
extended by two more endpoints:
- `GET /plugin`
- `GET /plugin/greet`

## Packages

### cmd/myapp
`cmd/myapp` is the main app entry point for MyApp. It creates an HTTP 
server that gets extended via a plugin.

### internal
`internal` houses the internal http handler logic and plugin loader.
The plugin loader makes a bunch of assumptions to keep the example
simple (such as the plugin file name and package name), but it
could use an expected and documented folder structure and naming
conventions to find and load multiple plugins.

### pkg
`pkg` contains `hello` and `sdk`:

`hello` just has a simple function
that is used by both internal app code and exposed to plugins.

`sdk` has a few parts. `myapp` contains an application interface
that is passed to plugins. `plugin.go` defines the Plugin related
interfaces. `lib` uses the `yaegi extract` command to export the
symbols we want to expose to the plugin. This is how we control
what is available to the plugin sandbox. The symbol files are
generated using the `//go:generate` commands in `symbols.go`

### plugins
`plugins` can be considered an external plugin source. It has
an example of how a plugin could be written for MyApp. For demo
purposes, it is embedded using `embed.FS` and loaded directly in
the plugin loader. It could also just as easily be a "plugins"
folder next to the compiled application, and the plugin loader could be
improved as described above in `internal` to facilitate dynamically
loading plugins at runtime.

