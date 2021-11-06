package main

import (
	"flag"

	"fmt"

	"github.com/alexellis/hmac"
)

func main() {

	var inputVar string
	var secretVar string

	// Bind flags to variables, then parse
	flag.StringVar(&inputVar, "message", "", "message to create digest from")
	flag.StringVar(&secretVar, "secret", "", "secret for the digest")
	flag.Parse()

	if len(secretVar) == 0 {
		panic("--secret is required")
	}

	fmt.Printf("Computing hash for: %q\nSecret: %q\n", inputVar, secretVar)

	digest := hmac.Sign([]byte(inputVar), []byte(secretVar))
	fmt.Printf("Digest: %x\n", digest)
}
