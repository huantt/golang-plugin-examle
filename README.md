### Plugin implementation
- A plugin package must be identified as main.
- Exported package functions and variables become shared library symbols. In the above, variable Speaker will be exported as a symbol in the compiled shared library.

### Build plugin
```shell
 go build -buildmode=plugin -o ./plugins/eng/eng.so ./plugins/eng/speaker.go
```

You can use Makefile to build all plugins quickly:
```shell
make build-plugins
```

### Open plugin
```go
plugin, err := plugin.Open(mod)
if err != nil {
    return err
}
```

### Look up symbol

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
