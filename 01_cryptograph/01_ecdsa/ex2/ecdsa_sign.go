package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"
)

type Signature struct {
	R *big.Int
	S *big.Int
}

func (s Signature) String() string { // 익명함수(Anonymous Function)
	return s.R.String() + s.S.String()
}

func sign(digest []byte, pvKey *ecdsa.PrivateKey) (*Signature, error) {
	r := big.NewInt(0)
	s := big.NewInt(0)

	r, s, err := ecdsa.Sign(rand.Reader, pvKey, digest)
	if err != nil {
		return nil, err //errors.New("failed to sign to msg.")
	}

	// prepare a signature structure to marshal into json
	signature := &Signature{
		R: r,
		S: s,
	}
	/*
		signature := r.Bytes()
		signature = append(signature, s.Bytes()...)
	*/
	return signature, nil
}

func SignASN1(digest []byte, pvKey *ecdsa.PrivateKey) ([]byte, error) {
	signature, err := ecdsa.SignASN1(rand.Reader, pvKey, digest)
	if err != nil {
		return nil, err //errors.New("failed to sign to msg.")
	}

	return signature, nil
}

func SignToString(digest []byte, pvKey *ecdsa.PrivateKey) (string, error) {
	signature, err := sign(digest, pvKey)
	if err != nil {
		return "", err
	}

	return signature.String(), nil
}

func main() {
	pvKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader) // elliptic.p224, elliptic.P384(), elliptic.P521()

	if err != nil {
		log.Println("ECDSA Keypair generation was Fail!")
	}

	//pbKey := &pvKey.PublicKey

	msg := "Hello SSI-KOREA."
	digest := sha256.Sum256([]byte(msg))
	signature, err := sign(digest[:], pvKey)
	if err != nil {
		log.Printf("Failed to sign msg.")
	}

	signatureASN1, err := SignASN1(digest[:], pvKey)
	if err != nil {
		log.Printf("Failed to sign msg.")
	}

	fmt.Printf("########## Sign ##########\n")
	fmt.Printf("===== Message =====\n")
	fmt.Printf("Msg: %s\n", msg)
	fmt.Printf("Digest: %x\n", digest)
	fmt.Printf("R: %s, S: %s\n", signature.R, signature.S)
	fmt.Printf("Signature: %+v\n", signature.String())
	fmt.Printf("SignatureASN1: %+v\n", signatureASN1)

}

// ########## Sign ##########
// ===== Message =====
// Msg: Hello SSI-KOREA.
// Digest: 4cc2b6b47522042263b8af953e3b766117a5095c376103b6a0db345ea47a27b3
// R: 77746183928329358379702647101136469287265183586970860829752690144804467377587, S: 2385433142901116708206832246275459130076555275705766322796242087935945570809
// Signature: 777461839283293583797026471011364692872651835869708608297526901448044673775872385433142901116708206832246275459130076555275705766322796242087935945570809
// SignatureASN1: [48 69 2 32 96 70 181 235 164 125 161 59 186 6 68 240 12 228 141 26 16 237 154 233 91 176 22 38 134 42 142 110 74 163 253 150 2 33 0 247 91 59 140 11 121 17 186 127 238 114 64 13 47 178 187 3 142 67 252 44 199 101 70 135 255 130 34 57 110 12 243]
