package main

import (
	"flag"
	"os"

	"pg_flame/pkg/html"
	"pg_flame/pkg/plan"
)

var (
	// goreleaser automatically overrides this based on the tag
	version  = "dev"
	hFlag    = flag.Bool("h", false, "print help info")
	helpFlag = flag.Bool("help", false, "print help info")
)

func main() {
	flag.Parse()

	if *hFlag || *helpFlag {
		printHelp()
	}

	p, err := plan.New(os.Stdin)
	if err != nil {
		handleErr(err)
	}

	err = html.Generate(os.Stdout, p)
	if err != nil {
		handleErr(err)
	}
}

func handleErr(err error) { _ = "STUB: not implemented"; return }

func printHelp() { _ = "STUB: not implemented"; return }
