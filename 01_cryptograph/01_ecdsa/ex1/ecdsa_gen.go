package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"log"
)

func main() {
	pvKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	if err != nil {
		log.Println("ECDSA Keypair generation was Fail!")
	}

	pbKey := &pvKey.PublicKey

	fmt.Printf("########## Key pair ##########\n")
	fmt.Printf("===== Private Key =====\n")
	fmt.Printf("Private Key: %x\n", pvKey.D)
	fmt.Printf("===== Public Key(X, Y) =====\n")
	fmt.Printf("X=%s Y=%s\n", pbKey.X, pbKey.Y)
	fmt.Printf("  Hex: X=%x Y=%x\n\n", pbKey.X.Bytes(), pbKey.Y.Bytes())
}

// root@fc3b49fc03e3:/workspaces/did-basic/01_cryptograph/01_ecdsa/ex1# go run ecdsa_gen.go
// ########## Key pair ##########
// ===== Private Key =====
// Private Key: f87b1c4b01f7b62466758bdb4c7844005a35154827f61f538821d09570b4d83c
// ===== Public Key(X, Y) =====
// X=64295562928159964977808443422744468796977481578217909140148472127059693908866 Y=35742744471439766032210439715299802192444756308958688345729360339191681268247
//   Hex: X=8e25ffbec01aeea349afc38cdd1bab6574c1d095f0686c299e6063458c37b382 Y=4f05ad2c5bf01438370e6f60dd6adcc33b1adb89b20b70f39982458080343a17

// root@fc3b49fc03e3:/workspaces/did-basic/01_cryptograph/01_ecdsa/ex1#
