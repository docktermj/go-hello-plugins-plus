package greetings

import (
	"github.com/docktermj/go-hello-plugins/common/runner"
	"github.com/docktermj/go-hello-plugins/subcommand/greetings/english"
	"github.com/docktermj/go-hello-plugins/subcommand/greetings/german"
	"github.com/docktermj/go-hello-plugins/subcommand/greetings/italian"
)

func Command(argv []string) {

	usage := `
Usage:
    go-hello-plugins greetings <subcommand> [<args>...]

Subcommands:
    english    Greetings from U.S.A.!
    german     Grüße aus Deutschland!
    italian    saluti DALL'ITALIA
`

	functions := map[string]interface{}{
		"english": english.Command,
		"german":  german.Command,
		"italian": italian.Command,
	}

	runner.Run(argv, functions, usage)
}
