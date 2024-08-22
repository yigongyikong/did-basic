package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func main() {
	pubKey, pvKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Println("Fail to generate key.")
	}

	fmt.Println("Private Key: ", pvKey)
	fmt.Println("Public Key: ", pubKey)

	msg := "Hello SSI-KOREA."
	digest := sha256.Sum256([]byte(msg))

	signature := ed25519.Sign(pvKey, digest[:])

	isVerify := ed25519.Verify(pubKey, digest[:], signature)

	if isVerify {
		fmt.Println("Verified.")
	} else {
		fmt.Println("Not Verified.")
	}
}

// Private Key:  [104 109 52 26 45 129 64 237 68 57 81 189 84 172 31 136 120 31 102 195 131 113 157 96 77 216 50 88 65 38 65 107 244 24 129 106 28 141 143 137 151 138 192 230 227 240 165 201 195 244 226 65 83 140 124 68 160 243 33 149 53 66 66 8]
// Public Key:  [244 24 129 106 28 141 143 137 151 138 192 230 227 240 165 201 195 244 226 65 83 140 124 68 160 243 33 149 53 66 66 8]
// Verified.
