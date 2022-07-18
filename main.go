package main

import (
	"fmt"

	"github.com/menxqk/my-interpreter/repl"
)

func main() {
	fmt.Println("My 'C-like' interpreter")
	repl.Start()
}
