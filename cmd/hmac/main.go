package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fail("No secret key provided\n")
	}

	secret := os.Args[1]

	h := hmac.New(sha256.New, []byte(secret))

	_, err := io.Copy(h, os.Stdin)
	if err != nil {
		fail("Failed to write to hash: %v\n", err)
	}

	fmt.Printf("sha256=%x\n", h.Sum(nil))
}

func fail(m string, args ...any) {
	fmt.Fprintf(os.Stderr, m, args...)
	os.Exit(1)
}
