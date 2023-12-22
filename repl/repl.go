package repl

import (
	"bufio"
	"fmt"
	"io"
	"log"

	"monkey/evaluator"
	"monkey/lexer"
	"monkey/parser"

	"net/http"
	_ "net/http/pprof"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	scanner := bufio.NewScanner(in)
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

		evaluated := evaluator.Eval(program)
		_, _ = io.WriteString(out, evaluated.Inspect())
		_, _ = io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	_, _ = io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	_, _ = io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		_, _ = io.WriteString(out, "\t"+msg+"\n")
	}
}
