# First steps learning Go

## Following this awesome book: Everyday Go by Alex Ellis https://openfaas.gumroad.com/l/everyday-golang

### Learnings:

External dependencies took a bit of getting used to compared to other languages I've had exposure to (mainly javascript, python) - but really nice once I got used to it. https://golang.org/ref/mod 

Use `go mod tidy` to add/remove dependencies imported in main.go. 

Flag built in module implements command line flag parsing. https://pkg.go.dev/flag@go1.17.3 
Can bind the flag to a variable using the Var functions (types e.g. stringVar, intVar). 
Call flag.parse() after all flags have been defined - parse the command line into the ^^ defined flags.
So for the code in my example, the command:
    `go run main.go -message="my message" -secret="terces os"` 
Would result in:
    `Computing hash for: "my message"`
    `Secret: "terces os"`

`%q` adds quotes around string.

Errors vs panics (something that unexpectedly goes wrong)! Want to dive into this more:
https://gobyexample.com/errors 
https://gobyexample.com/panic 

### Come back to: 

How to use `hmac.Validate()` (`--generate` or `--validate` flag modes)










