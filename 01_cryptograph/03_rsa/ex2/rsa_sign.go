package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
)

func sign(msg string, pvKey *rsa.PrivateKey) ([]byte, error) {
	digest := sha256.Sum256([]byte(msg))

	signature, err := rsa.SignPKCS1v15(
		rand.Reader,
		pvKey,
		crypto.SHA256,
		digest[:],
	)

	return signature, err
}

func main() {
	pvKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic("Failed to Generate RSA.")
	}

	msg := "Hello World."
	sig, err := sign(msg, pvKey)
	if err != nil {
		panic("Failed to Sign.")
	}

	fmt.Printf("Signature: %v\n", sig)

	sigBase58 := base58.Encode(sig)
	fmt.Printf("Signature(base58): %v\n", sigBase58)
}

// Signature: [2 251 192 55 142 54 103 194 59 131 25 19 65 131 97 30 252 172 219 28 84 163 89 182 56 221 191 146 93 140 234 21 179 248 30 151 227 56 173 9 248 158 238 187 220 157 147 50 95 183 28 36 155 126 149 5 234 201 207 18 244 41 12 21 88 201 75 118 74 232 131 84 38 41 54 235 182 135 36 255 178 236 45 8 123 249 62 88 172 239 71 196 232 105 248 23 87 192 248 54 252 125 75 78 118 220 126 230 217 38 124 134 99 246 153 169 76 24 120 113 98 51 205 147 15 222 222 225 22 125 118 165 93 40 107 116 135 174 133 129 208 176 32 239 158 173 207 145 217 50 159 114 246 39 147 230 67 189 173 74 19 77 202 141 109 61 71 174 80 6 242 214 18 96 152 169 225 89 119 187 90 97 161 77 153 235 29 81 89 10 133 139 42 129 210 111 128 93 213 205 236 34 197 168 64 24 163 165 110 189 147 134 106 208 40 58 73 182 144 26 53 174 94 229 69 140 111 245 247 142 129 4 187 166 121 203 25 2 0 6 213 65 119 168 108 32 204 1 219 227 69 47 114 184 170 64]
// Signature(base58): 8zs1xjDscsMoXxJ4BM9tT4BwrkGkK5w4Tg4cK3tTBiMTB1YHBYyYyeWNSigHxXxVSL92ycwuYKPMVxzaoaeFgVYTPasrBWZ6Qt9w1xTUtLXSqMbPY1Ybd9s23dtXKtXnPnGbunvZSDnA4WnmbopXZHU4WVoJx7Wd76JkQvER3UAVJfYvfp3n8k7i94YGqPfe5fnVFrtD5wcRg27tz7fAm476888q4zEuQKEm2yTzG65R7NBtmSJBMUiLs98m1LB2RHaaJRHu7C5ig6FpnQBQ6uzo9XGx2odf3tZwSdmiw4ampgUrReEgvjVRofXM5AWT9CCut76MsPwHAc8eENqewot4Bu8gF
