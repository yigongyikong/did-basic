package main

import (
	"fmt"

	"github.com/hyperledger/aries-framework-go/pkg/doc/verifiable"
)

func main() {
	credential := verifiable.Credential{
		Context: []string{"https://www.w3.org/2018/credentials/v1"},
		ID:      "12345",
		Types:   []string{"VerifiableCredential"},
		Issuer:  verifiable.Issuer{ID: "did:test:12343543"},
	}

	fmt.Println(credential)

}

// {[https://www.w3.org/2018/credentials/v1] [] 12345 [VerifiableCredential] <nil> {did:test:12343543 map[]} <nil> <nil> [] <nil> [] <nil> [] []   []  map[]}
