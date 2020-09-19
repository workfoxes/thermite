package main

import (
	"flag"
	"fmt"
	"sort"
)

// ThermiteUsageDocs : Will Display all the Usage Command for the Thermite Libiray
func ThermiteUsageDocs(fs flag.Flag) {
	fmt.Fprintln(fs.Output(), "Usage: vegeta [global flags] <command> [command flags]")
	fmt.Fprintf(fs.Output(), "\nglobal flags:\n")
	fs.PrintDefaults()

	names := make([]string, 0, len(commands))
	for name := range commands {
		names = append(names, name)
	}

	sort.Strings(names)
	for _, name := range names {
		if cmd := commands[name]; cmd.fs != nil {
			fmt.Fprintf(fs.Output(), "\n%s command:\n", name)
			cmd.fs.SetOutput(fs.Output())
			cmd.fs.PrintDefaults()
		}
	}

	fmt.Fprintln(fs.Output(), examples)
}
