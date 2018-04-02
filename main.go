package main

import (
	"flag"
	"fmt"
)

const version string = "0.1.0"

func main() {
	var line int
	flag.IntVar(&line, "n", 10, "")
	flag.Parse()
	args := flag.Args()
	fmt.Println(args)
}
