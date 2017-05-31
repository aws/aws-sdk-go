Retrieve Credentials with Go Plugin
===

This example demonstrates how you can take advantage of Go 1.8's new Plugin
functionality to retrieve AWS credentials dynamically from a plugin compiled
separate from your application.

Usage
---

Example Plugin
---

You can find the plugin at `plugin/plugin.go` nested within this example. The plugin
demonstrates what symbol the SDK will use when lookup up the credential provider
and the type signature that needs to be implemented.

Compile the plugin with:

   go build -tags example -o plugin.so -buildmode=plugin plugin.go

Example Application
---

The `main.go` file in this folder demonstrates how you can configure the SDK to 
use a plugin to retrieve credentials with.

Compile and run application:

  go build -tags example -o usePlugin usePlugin.go

  ./usePlugin ./plugin/plugin
