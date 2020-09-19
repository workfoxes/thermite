package cmd

import (
	"flag"
	"fmt"
	"sort"
)

// Command : is Structure for handling all cmd command
type Command struct {
	fs *flag.FlagSet
	fn func(args []string) error
}

const examples = `
examples:
  echo "GET http://localhost/" | thermite attack -duration=5s | tee results.bin
`

// Commands is list all the acceptable
var (
	Commands map[string]Command
)

func init() {
	Commands = map[string]Command{
		"server": serverCmd(),
	}
}

// ThermiteUsageDocs : Will Display all the Usage Command for the Thermite Libiray
func ThermiteUsageDocs(fs *flag.FlagSet) {
	fs.Usage = func() {
		fmt.Fprintln(fs.Output(), "Usage: thermite [global flags] <command> [command flags]")
		fmt.Fprintf(fs.Output(), "\nglobal flags:\n")
		fs.PrintDefaults()

		names := make([]string, 0, len(Commands))
		for name := range Commands {
			names = append(names, name)
		}

		sort.Strings(names)
		for _, name := range names {
			if cmd := Commands[name]; cmd.fs != nil {
				fmt.Fprintf(fs.Output(), "\n%s command:\n", name)
				cmd.fs.SetOutput(fs.Output())
				cmd.fs.PrintDefaults()
			}
		}

		fmt.Fprintln(fs.Output(), examples)
	}
}
