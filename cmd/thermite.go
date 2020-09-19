package cmd

import "flag"

type command struct {
	fs *flag.FlagSet
	fn func(args []string) error
}

func init() {
	commands := map[string]command{
		"server": serverCmd(),
		// "attack": attackCmd(),
		// "report": reportCmd(),
		// "plot":   plotCmd(),
		// "encode": encodeCmd(),
		// "dump":   dumpCmd(),
	}
}
