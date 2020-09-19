package main

import "flag"

const examples = `
examples:
  echo "GET http://localhost/" | thermite attack -duration=5s | tee results.bin | vegeta report
  vegeta report -type=json results.bin > metrics.json
  cat results.bin | vegeta plot > plot.html
  cat results.bin | vegeta report -type="hist[0,100ms,200ms,300ms]"
`

func init() {

}

func main() {
	fs := flag.NewFlagSet("thermite", flag.ExitOnError)
}
