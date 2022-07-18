package main

import (
	"flag"
	"fmt"

	"github.com/menxqk/my-interpreter/repl"
)

func main() {
	fmt.Println("My 'C-like' interpreter")

	debug := flag.Bool("debug", false, "Show debug information")
	flag.Parse()

	repl.Start(*debug)
}
