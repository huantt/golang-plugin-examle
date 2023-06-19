# Golang plugin implementation
This repository explores the plugin feature introduced in Go 1.8, providing a example to help you understand and utilize this powerful functionality.
## About plugin feature in Golang
Starting from Go 1.8, developers have the ability to write plugins and load them dynamically at runtime. This feature adds a new level of flexibility to Go programs, allowing for modular and extensible designs. With plugins, you can extend your application's functionality without the need to recompile or redeploy the entire codebase.

## Key benefit
- Dynamic Loading: Plugins can be loaded and unloaded dynamically at runtime, enabling the addition or removal of functionality without disrupting the main program.
- Flexibility: Unlike traditional library imports, which are resolved at build time, plugins offer a more flexible approach. They can be swapped or updated independently, giving you the freedom to iterate and experiment with different implementations.
- Modularity: The plugin feature promotes a modular design pattern by encapsulating specific functionalities in separate components. This enhances code organization, reusability, and maintainability.
- Encapsulation: Plugins run in their own isolated namespaces, preventing conflicts between different plugins or the main program. This allows for secure and reliable extensibility.

## How to run this example:
1. Build plugins
```shell
 go build -buildmode=plugin -o ./plugins/eng/eng.so ./plugins/eng/speaker.go
```

You can use the Makefile to quickly build all plugins:
```shell
make build-plugins
```

2. Run
```shell
go run main.go english

# Alice says "hello" in English

go run main.go vietnamese

# Anh Thư says "xin chào" in Vietnamese
```

## How to implement a plugin
- A plugin package must be identified as main.
- Exported package functions and variables become shared library symbols. In the above, variable Speaker will be exported as a symbol in the compiled shared library.

Example implementation:
```go
package main

type speaker struct {
}

func (s *speaker) Speak() string {
	return "hello"
}

// Exported
var Speaker speaker
var SpeakerName = "Alice"
```

In this example, we are exporting both the `Speaker` type and the `SpeakerName` variable, which are referred to as `symbols` in the context of plugins. By exporting these symbols, we enable their accessibility and visibility to other packages and modules that import them.

## How to open plugin
```go
plugin, err := plugin.Open("path/to/plugin.so")
if err != nil {
    return err
}
```

## How to look up symbol

You can look up the symbol that has been exported in plugin implementation by the following syntax:
```go
symSpeaker, err := plugin.Lookup("Speaker")
if err != nil {
  return err
}

var speaker Speaker
speaker, ok := symSpeaker.(Speaker)
if !ok {
    return errors.New("unexpected type from module symbol")
}
speaker.Speak()
```
