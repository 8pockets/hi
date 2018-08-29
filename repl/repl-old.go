package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/8pockets/hi/evaluator"
	"github.com/8pockets/hi/lexer"
	"github.com/8pockets/hi/object"
	"github.com/8pockets/hi/parser"
)

func Start2(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated == nil || evaluated.Type() == object.NULL_OBJ {
			continue
		}
		io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")
	}
}
