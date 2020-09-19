package cmd

import "flag"

const helpDocs = `

`

type serverOption struct {
	start bool
	name  string
	port  int
}

func serverCmd() Command {
	fs := flag.NewFlagSet("thermite server", flag.ExitOnError)
	opts := &serverOption{}
	fs.BoolVar(&opts.start, "start", false, "Start Thermite as Application server")
	fs.IntVar(&opts.port, "port", 8004, "Port on which the thermite web application need to been run")
	return Command{fs, func(args []string) error {
		fs.Parse(args)
		return server(opts)
	}}
}

func server(opts *serverOption) (err error) {
	return nil
}
