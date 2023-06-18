package main

import (
	"errors"
	"fmt"
	"os"
	"plugin"
	"strings"
)

type Speaker interface {
	Speak() string
}

func main() {
	err := run(os.Args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func run(args []string) error {
	lang := "english"
	if len(args) == 2 {
		lang = args[1]
	}
	var mod string
	switch lang {
	case "english":
		mod = "./plugins/eng/speaker.so"
	case "vietnamese":
		mod = "plugins/vie/speaker.so"
	default:
		return errors.New("this speakerName is not supported")
	}

	plugin, err := plugin.Open(mod)
	if err != nil {
		return err
	}

	speaker, err := lookUpSymbol[Speaker](plugin, "Speaker")
	if err != nil {
		return err
	}

	speakerName, err := lookUpSymbol[string](plugin, "SpeakerName")
	if err != nil {
		return err
	}
	fmt.Printf("%s says \"%s\" in %s", *speakerName, (*speaker).Speak(), strings.Title(lang))
	return nil
}

func lookUpSymbol[M any](plugin *plugin.Plugin, symbolName string) (*M, error) {
	symbol, err := plugin.Lookup(symbolName)
	if err != nil {
		return nil, err
	}
	switch symbol.(type) {
	case *M:
		return symbol.(*M), nil
	case M:
		result := symbol.(M)
		return &result, nil
	default:
		return nil, errors.New(fmt.Sprintf("unexpected type from module symbol: %T", symbol))
	}
}
