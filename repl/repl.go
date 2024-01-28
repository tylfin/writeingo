package repl

import (
	"bufio"
	"fmt"
	"io"
	"log"

	"monkey/compiler"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"monkey/vm"

	"net/http"
	_ "net/http/pprof"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	scanner := bufio.NewScanner(in)
	// env := object.NewEnvironment()
	constants := []object.Object{}
	globals := make([]object.Object, vm.GlobalsSize)
	symbolTable := compiler.NewSymbolTable()
	for i, v := range object.Builtins {
		symbolTable.DefineBuiltin(i, v.Name)
	}

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

		comp := compiler.NewWithState(symbolTable, constants)
		err := comp.Compile(program)
		if err != nil {
			fmt.Fprintf(out, "Woops! Compilation failed:\n %s\n", err)
			continue
		}

		machine := vm.NewWithGlobalsStore(comp.Bytecode(), globals)
		err = machine.Run()
		if err != nil {
			fmt.Fprintf(out, "Woops! Executing bytecode failed:\n %s\n", err)
			continue
		}

		stackTop := machine.LastPoppedStackElem()
		_, _ = io.WriteString(out, stackTop.Inspect())
		_, _ = io.WriteString(out, "\n")

		// Eval is the interper-based solution from book 1.
		// evaluated := evaluator.Eval(program, env)
		// if evaluated != nil {
		// 	_, _ = io.WriteString(out, evaluated.Inspect())
		// 	_, _ = io.WriteString(out, "\n")
		// }
	}
}

func printParserErrors(out io.Writer, errors []string) {
	_, _ = io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	_, _ = io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		_, _ = io.WriteString(out, "\t"+msg+"\n")
	}
}
