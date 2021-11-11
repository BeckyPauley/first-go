# First steps learning Go

## Following this awesome book: Everyday Go by Alex Ellis </br>https://openfaas.gumroad.com/l/everyday-golang

### Learnings:

- External dependencies took a bit of getting used to compared to other languages I've had exposure to (mainly javascript, python) - but really nice once I got used to it. https://golang.org/ref/mod 

- Use `go mod tidy` to add/remove dependencies imported in main.go. 

- Flag built in module implements command line flag parsing. https://pkg.go.dev/flag@go1.17.3 </br>
Can bind the flag to a variable using the Var functions (types e.g. stringVar, intVar). </br>
- Call flag.parse() after all flags have been defined - parse the command line into the ^^ defined flags. </br>
So for the code in my example, the command: </br>
    `go run main.go -message="my message" -secret="terces os"` </br> 
Would result in: </br>
    `Computing hash for: "my message"` </br>
    `Secret: "terces os"`

- `%q` adds quotes around string.

- By default, `go build` builds for the current system OS/CPU architecture. Can use Go cross-compliation to compile binary for a different system, for most common OS: </br>
*Intel*:</br>
`GOOS=[CPU architecure] go build -o [filename, plus extension if required]`</br>
e.g. `GOOS=windows go build -o first-go.exe`</br>
*ARM* (additions to start and end of command):</br>
e.g. `GOARCH=arm GOOS=linux go build -o first-go-arm64`</br>
*To debug:*</br>
Run `file [filename]`

### Come back to: 

- How to use `hmac.Validate()` (`--generate` or `--validate` flag modes)

- Errors vs panics (something that unexpectedly goes wrong)! Want to dive into this more: </br>
 https://gobyexample.com/errors </br>
 https://gobyexample.com/panic 










