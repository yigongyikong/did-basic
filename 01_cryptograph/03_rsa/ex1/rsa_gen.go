package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

func main() {
	// Generate RSA Key
	pvKey, err := rsa.GenerateKey(rand.Reader, 1024) // 2048, 4096
	if err != nil {
		log.Printf("Failed to Generate RSA.")
		os.Exit(0)
	}

	fmt.Printf("Private Key: %+v\n", pvKey) //%+v 필드명과 함께 출력.

	pvKeyByte := x509.MarshalPKCS1PrivateKey(pvKey)
	pvKeyStr := base64.StdEncoding.EncodeToString(pvKeyByte)
	fmt.Printf("Private Key(str): %+v\n", pvKeyStr)

	pbKey := pvKey.PublicKey
	fmt.Printf("Public Key: %+v\n", pbKey)

	pbKeyByte := x509.MarshalPKCS1PublicKey(&pbKey)
	fmt.Printf("Public Key(Byte): %+v\n", pbKeyByte)

	pbKeyStr := base64.StdEncoding.EncodeToString(pbKeyByte)
	fmt.Printf("Public Key(str): %+v\n", pbKeyStr)

}

// Private Key: &{PublicKey:{N:+133948575203947659963751538838871879769784325307216761966270167485187408443674387657316454848677842735180017712717308656751486192169318887352707342099919939182401754440534164267680196832425411592947013974604420014801498500082899175968756275731518789486346012037233187637023013191642063462946963889852194731973 E:65537} D:+95483094859435476609038873307224584695744607250518721772101918220325925206572415096614932346089732197985945152890954857827779429018326601163563008695575914627907918411496576729164167925272150496759269381209979670049642226199212330900528524825831705433984829744276602536684641677401649045878656094660869324161 Primes:[+11245997939124979982818031129752321307234774918119253247673237994786317696771070232180041114980092540960058045823232512557060899011677660698465450763243109 +11910777143034922909126956426856661654562463633466148988176569103333209563589687189079925039327962483182275425228555746593690691852574725880886100825865697] Precomputed:{Dp:+11211678400729493524839124462893135438771968335827969688011754441999994497061681428200060215289901984355513870637236105908733813856691385279393886758899781 Dq:+3556131290076816707552789361800270055004405540626396344838044580083325625688685921368492504150788743891657891655815040261207344972753248719216600315237409 Qinv:+1934104625302408632373101496226149555419007601854042480989448071633584218640885136133020644136143301277398373421572802996466193137863883832972291411595273 CRTValues:[] n:0x4000074160 p:0x4000074180 q:0x40000741a0}}
// Private Key(str): MIICXQIBAAKBgQC+v8Uy5cfTms7LHqkQx7qm2q4ZVrbW3hc2lZCus85Qg67EGJNf5r1YJYLeh9S6q4OIUapOJ1i6zdyHjpKTotX1kT1Qcet3k7DuliATnWj09ddkQgxjLsBnes5ql5HmGjgHmHxi/L21puZgjkX9JSpOIas4K2BJIgSq9gcfNfO3xQIDAQABAoGBAIf48dyHWbuopfK7B29zwMUCK4raenmOVWPOmBVBDTfttp4Imr2JIL39910j3Gu4qYl6FXmiqKh46NQkuK3PLM6mLWfj3cd9uqYIgD7ma33t1QMs1otjXMZ5NWY7cXdUCAUeE3bmLJc853ME/hEi1m6hzrQiTyL8HBmeW6g8PX2BAkEA1rlK+ftV0fU4mlsyHY3u/wOMT9KSltcKDJkExhyKTAj+lO6VSB3XMIR911ROhVQCGEgStecYxM+N8aHTZpBOZQJBAONqqnr4XE5BrrMy0IImQlPyNms4sLIrfwLFserG2wLG6Duw+oFD3F81wdVYdet899lge8QGX9yXLQmgnJxB/eECQQDWEYrfKCRKGrjgolSXRF5l3Oqw1dURDrfVEWTCAcgav8jXP/iSGjXkB+LRnxpdMACappgIw6lc5TO1F0mD8hBFAkBD5gH2axasKsItMJu+cAvXonaK9scSuxfVzQ68yRh0sMx/nex9EbTHHHa96wZyA8LAPSVWtoyCgQRqXnDD2pQhAkAk7bTb8PvmECGXFGc6LJGgan1drrrhfpdehVhLeiEa7AgL5n+MX2/5ccqp5S7MYBGHrkpEMyvwOon0lh0toNAJ
// Public Key: {N:+133948575203947659963751538838871879769784325307216761966270167485187408443674387657316454848677842735180017712717308656751486192169318887352707342099919939182401754440534164267680196832425411592947013974604420014801498500082899175968756275731518789486346012037233187637023013191642063462946963889852194731973 E:65537}
// Public Key(Byte): [48 129 137 2 129 129 0 190 191 197 50 229 199 211 154 206 203 30 169 16 199 186 166 218 174 25 86 182 214 222 23 54 149 144 174 179 206 80 131 174 196 24 147 95 230 189 88 37 130 222 135 212 186 171 131 136 81 170 78 39 88 186 205 220 135 142 146 147 162 213 245 145 61 80 113 235 119 147 176 238 150 32 19 157 104 244 245 215 100 66 12 99 46 192 103 122 206 106 151 145 230 26 56 7 152 124 98 252 189 181 166 230 96 142 69 253 37 42 78 33 171 56 43 96 73 34 4 170 246 7 31 53 243 183 197 2 3 1 0 1]
// Public Key(str): MIGJAoGBAL6/xTLlx9OazsseqRDHuqbarhlWttbeFzaVkK6zzlCDrsQYk1/mvVglgt6H1Lqrg4hRqk4nWLrN3IeOkpOi1fWRPVBx63eTsO6WIBOdaPT112RCDGMuwGd6zmqXkeYaOAeYfGL8vbWm5mCORf0lKk4hqzgrYEkiBKr2Bx8187fFAgMBAAE=
