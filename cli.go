package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

func readLine(line int, filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var result string
	scanner := bufio.NewScanner(file)
	for i := 1; scanner.Scan(); i++ {
		result += fmt.Sprintf("%s\n", scanner.Text())
		if i >= line {
			break
		}
	}
	return result, nil
}

// Run invokes the CLI with the given arguments.
func (c *CLI) Run(args []string) int {
	var (
		line    int
		version bool
	)

	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(c.outStream)

	flags.IntVar(&line, "lines", 10, "number of print line from head")
	flags.IntVar(&line, "n", 10, "number of print line from head(Short)")

	flag.BoolVar(&version, "version", false, "print version information")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	if version {
		fmt.Fprintf(c.errStream, "%s version %s", Name, Version)
		return ExitCodeOK
	}

	filepath := flags.Args()

	l, err := readLine(line, filepath[0])
	if err != nil {
		fmt.Fprint(c.errStream, err)
		return ExitCodeError
	}
	fmt.Fprint(c.outStream, l)
	return ExitCodeOK
}
