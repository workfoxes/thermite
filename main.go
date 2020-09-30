package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/airstrik/thermite/cmd"
)

// Set Thermite Global Variable
var (
	Commit  string
	Date    string
	Version string
)

func init() {

}

func main() {
	fs := flag.NewFlagSet("thermite", flag.ExitOnError)

	// Print the version Information of Thermite
	version := fs.Bool("version", false, "Print version and exit")
	if *version {
		fmt.Printf("Version: %s\nCommit: %s\nRuntime: %s %s/%s\nDate: %s\n",
			Version,
			Commit,
			runtime.Version(),
			runtime.GOOS,
			runtime.GOARCH,
			Date,
		)
		return
	}

	cpus := fs.Int("cpus", runtime.NumCPU(), "Number of CPUs to use")
	// profile := fs.String("profile", "", "Enable profiling of [cpu, heap]")

	runtime.GOMAXPROCS(*cpus)

	// for _, prof := range strings.Split(*profile, ",") {
	// 	if prof = strings.TrimSpace(prof); prof == "" {
	// 		continue
	// 	}

	// 	f, err := os.Create(prof + ".pprof")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	defer f.Close()

	// 	switch {
	// 	case strings.HasPrefix(prof, "cpu"):
	// 		pprof.StartCPUProfile(f)
	// 		defer pprof.StopCPUProfile()
	// 	case strings.HasPrefix(prof, "heap"):
	// 		defer pprof.Lookup("heap").WriteTo(f, 0)
	// 	}
	// }
	cmd.ThermiteUsageDocs(fs)
	fs.Parse(os.Args[1:])

	args := fs.Args()
	log.Print(args)
	if len(args) == 0 {
		fs.Usage()
		os.Exit(1)
	}
	_cmd, ok := cmd.Commands[args[0]]
	if !ok {
		log.Fatalf("Invalid Command : %s", args)
	} else if err := _cmd.fn(args[1:]); err != nil {
		log.Fatal(err)
	}

}
