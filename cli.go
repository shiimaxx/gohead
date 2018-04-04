package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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

func getBody(path string) (io.Reader, error) {
	if strings.HasPrefix(path, "http") {
		return getHTTP(path)
	}
	return getFile(path)
}

func getHTTP(url string) (io.Reader, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("%s: status code was not 200", res.Status)
	}
	return res.Body, nil
}

func getFile(filepath string) (io.Reader, error) {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return nil, fmt.Errorf("%s: No such file or directory", filepath)
	}

	finfo, err := os.Stat(filepath)

	if os.IsNotExist(err) {
		return nil, fmt.Errorf("%s: No such file or directory", filepath)
	}
	if finfo.IsDir() {
		return nil, fmt.Errorf("%s: Is a directory", filepath)
	}
	return os.Open(filepath)
}

func readLine(line int, body io.Reader) (string, error) {
	var result string
	scanner := bufio.NewScanner(body)
	for i := 1; scanner.Scan(); i++ {
		result += fmt.Sprintf("%s\n", scanner.Text())
		if i >= line {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error on reading file: %s", err)
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

	flags.BoolVar(&version, "version", false, "print version information")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	if version {
		fmt.Fprintf(c.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	if len(flags.Args()) < 1 {
		fmt.Fprint(c.errStream, "Missing filename")
		return ExitCodeError
	}

	body, err := getBody(flags.Args()[0])
	if err != nil {
		fmt.Fprint(c.errStream, err)
		return ExitCodeError
	}

	l, err := readLine(line, body)
	if err != nil {
		fmt.Fprint(c.errStream, err)
		return ExitCodeError
	}
	fmt.Fprint(c.outStream, l)
	return ExitCodeOK
}
