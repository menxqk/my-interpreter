package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/menxqk/my-interpreter/evaluator"
	"github.com/menxqk/my-interpreter/lexer"
	"github.com/menxqk/my-interpreter/parser"
)

const (
	PROMPT = ">>> "
	DOTS   = "... "
)

var (
	in  = os.Stdin
	out = os.Stdout
)

func Start(debug bool) {
	eval := evaluator.New()

	scanner := bufio.NewScanner(in)
	for {
		out.WriteString(PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			out.WriteString("\n")
			return
		}

		line := scanner.Text()
		l := lexer.New(line, debug)
		p := parser.New(l, debug)
		program := p.ParseProgram()

		if p.HasErrors() {
			printErrors(p.Errors())
		} else {
			res := eval.Eval(program)
			out.WriteString(res.Inspect() + "\n")
		}
	}
}

func printErrors(errors []string) {
	for _, e := range errors {
		fmt.Printf("%s\n", e)
	}
}
