// help functionality.

package help

import (
	"github.com/docopt/docopt-go"
)

// Show help and potentially exit.
func ShowHelp(usage string) {
	help := true
	version := ""
	optionsFirst := false
	docopt.Parse(usage, nil, help, version, optionsFirst)
}
