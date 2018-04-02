package main

import (
	"flag"
	"fmt"
	"io"
)

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

// Run invokes the CLI with the given arguments.
func (c *CLI) Run(args []string) int {
	var line int
	var version bool
	flag.IntVar(&line, "n", 10, "")
	flag.BoolVar(&version, "v", false, "print the version")
	flag.Parse()

	if version {
		fmt.Fprintf(c.errStream, "%s version %s\n", Name, Version)
		return 0
	}
	return 0
}
