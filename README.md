# Write in Go

[![Go](https://github.com/tylfin/writeingo/actions/workflows/go.yml/badge.svg)](https://github.com/tylfin/writeingo/actions/workflows/go.yml)

Repository to cover contents of "Writing An Interpreter in Go" and "Writing a Compiler in Go"

## REPL

The monkey language supports a REPL that can be used like so:

```bash
go run main.go
Hello tylfin! This is the Monkey programming language!
Feel free to type in commands
>> let x = 5;
let x = 5;
>> 5 + 4 * (2 + 4) / 5
(5 + ((4 * (2 + 4)) / 5))
>> fn(x, y) { x + y; }
fn(x, y)(x + y)
>> let x 5 2
Woops! We ran into some monkey business here!
 parser errors:
        expected next token to be =, got INT instead
```

## Project Structure

The repository is arranged like:

```bash
├── main.go
├── ast
├── lexer
├── parser
├── repl
└── token
```

Where:

- main.go: Entrypoint to the REPL
- AST: Abstract Syntax Tree interface/structs used to store results from parser
- Lexer: Takes raw input from the user and converts to tokens
- Parser: Takes tokens from lexer and creates a program as an AST using Top Down Operator Precedence (or: Pratt Parsing)
- REPL: Read-Eval-Print Loop
- Token: Raw tokens that are valid in the monkey language

## More on Monkey

Here is how we bind values to names in Monkey:

```bash
let age = 1;
let name = "Monkey";
let result = 10 * (20 / 2);
```

Besides integers, booleans and strings, the Monkey interpreter supports arrays and hashes.
Here’s what binding an array of integers to a name looks like:

```bash
let myArray = [1, 2, 3, 4, 5];
And here is a hash, where values are associated with keys:
let thorsten = {"name": "Thorsten", "age": 28};
Accessing the elements in arrays and hashes is done with index expressions:
myArray[0] // => 1 thorsten["name"] // => "Thorsten"
```

The let statements can also be used to bind functions to names. Here’s a small function that adds two numbers:

```bash
let add = fn(a, b) { return a + b; };
```

Implicit return values are also possible,

```bash
let add = fn(a, b) { a + b; };
```

And calling a function is as easy as you’d expect:

```bash
add(1, 2);
```

A more complex function, such as a fibonacci function that returns the Nth Fibonacci number,
might look like this:

```bash
let fibonacci = fn(x) {
    if (x == 0) {
        0
    } else {
        if (x == 1) {
            1
        } else {
            fibonacci(x - 1) + fibonacci(x - 2);
        }
    }
};
```

Note the recursive calls to fibonacci itself!

Monkey also supports a special type of functions, called higher order functions. These are functions that take other
functions as arguments. Here is an example:

```bash
let twice = fn(f, x) { return f(f(x)); };
let addTwo = fn(x) { return x + 2; };
twice(addTwo, 2); // => 6
```

## Benchmark

From the final chapter of Writing a Compiler in Go, my compiled version was 3.05x as fast:

```bash
$ go build -o fibonacci ./benchmark
$ ./fibonacci -engine=eval
engine=eval, result=9227465, duration=12.267151708s
$ ./fibonacci -engine=vm
engine=vm, result=9227465, duration=4.013193583s
```

## Book References

- [Writing An Interpreter in Go](https://interpreterbook.com/)
- [Writing a Compiler in Go](https://compilerbook.com/)
