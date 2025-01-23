package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	if len(os.Args) != 1 {
		log.Fatalln("No secret key provided")
	}

	secret := os.Args[0]

	h := hmac.New(sha256.New, []byte(secret))

	_, err := io.Copy(h, os.Stdin)
	if err != nil {
		log.Fatalf("Failed to write to hash: %v", err)
	}

	fmt.Printf("sha256=%x\n", h.Sum(nil))
}
