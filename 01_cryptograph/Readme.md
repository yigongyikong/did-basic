[[root@fc3b49fc03e3:/workspaces/did-basic/01_cryptograph/01_ecdsa#]]
http://cryptostudy.xyz/crypto/article/4-ECDSA-%EC%84%9C%EB%AA%85
- ECDSA(Elliptic Curve Digital Signature Algorithm) 란?
    - DSA는 디지털 서명 알고리즘을 말하고, ECDSA는 Elliptic Curve(타원 곡선) 기술을 사용하여 
    서명을 수행하는 알고리즘입니다.
    - ?? 타원 곡선에 위치한 점의 곱셈 연산이 일방향 함수의 기능을 한다는 점을 이용한다.
    한 시작점(base point)에서 N번 곱셈 연산을 수행했을 때, 종료점은 쉽게 구할 수 있지만
    시작점에서 얼마만큼의 곱셈 연산을 수행해야 종료점이 되는지 알기 어렵다??

[root@fc3b49fc03e3:/workspaces/did-basic/01_cryptograph/01_ecdsa/ex1# go run ecdsa_gen.go]
- ECDSA 방식을 통해 지갑을 생성합니다.

[root@fc3b49fc03e3:/workspaces/did-basic/01_cryptograph/01_ecdsa/ex2# go run ecdsa_sign.go]
- ECDSA 방식을 통해 서명을 진행합니다.
    - sign : privKey를 사용해 서명하며, integers 값으로 서명값이 나온다.
    Sign signs a hash (which should be the result of hashing a larger message) using the private key, priv. If the hash is longer than the bit-length of the private key's curve order, the hash will be truncated to that length. It returns the signature as a pair of integers. Most applications should use [SignASN1] instead of dealing directly with r, s.
    - SignASN1 : privKey를 사용해 서명하며, ASN.1으로 인코딩된 서명값을 도출됩니다.
    SignASN1 signs a hash (which should be the result of hashing a larger message) using the private key, priv. If the hash is longer than the bit-length of the private key's curve order, the hash will be truncated to that length. It returns the ASN.1 encoded signature.

[root@fc3b49fc03e3:/workspaces/did-basic/01_cryptograph/01_ecdsa/ex3# go run ecdsa_verify.go]
- ECDSA 방식을 통해 서명된 값을 검증합니다.
    - verify : sign 값 검증
    - verifyASN1 : SignASN1 값 검증


[[root@fc3b49fc03e3:/workspaces/did-basic/01_cryptograph/02_eddsa#]]
https://m.blog.naver.com/aepkoreanet/222088849086
- 
[root@fc3b49fc03e3:/workspaces/did-basic/01_cryptograph/02_eddsa/ex1# go run ed25519.go]
- ED25519 방식을 통해 지갑 생성, 서명, 검증을 합니다.

[[root@fc3b49fc03e3:/workspaces/did-basic/01_cryptograph/03_rsa#]]
https://gngsn.tistory.com/96
[root@fc3b49fc03e3:/workspaces/did-basic/01_cryptograph/03_rsa/ex1# go run rsa_gen.go]
- rsa 방식을 통해 지갑을 생성합니다.

[root@fc3b49fc03e3:/workspaces/did-basic/01_cryptograph/03_rsa/ex2# go run rsa_sign.go]
- rsa 방식을 통해 서명을 진행합니다.
- import할 package를 찾을 수 없다.
    rsa_sign.go:10:2: no required module provides package github.com/btcsuite/btcutil/base58: go.mod file not found in current directory or any parent directory; see 'go help modules'
- root@fc3b49fc03e3:/workspaces/did-basic# go mod init github.com/btcsuite/btcutil/base58
    go: creating new go.mod: module github.com/btcsuite/btcutil/base58
    go: to add module requirements and sums:
        go mod tidy
- root@fc3b49fc03e3:/workspaces/did-basic/01_cryptograph/03_rsa/ex2# go run rsa_sign.go 
    rsa_sign.go:10:2: no required module provides package github.com/btcsuite/btcutil/base58; to add it:
        go get github.com/btcsuite/btcutil/base58
- root@fc3b49fc03e3:/workspaces/did-basic/01_cryptograph/03_rsa/ex2# go get github.com/btcsuite/btcutil/base58
    go: downloading github.com/btcsuite/btcutil v1.0.2
    go: added github.com/btcsuite/btcutil v1.0.2
- root@fc3b49fc03e3:/workspaces/did-basic/01_cryptograph/03_rsa/ex2# go run rsa_sign.go

[root@fc3b49fc03e3:/workspaces/did-basic/01_cryptograph/03_rsa/ex3# go run rsa_verify.go]
- rsa 방식을 통해 서명된 값을 검증합니다.

[root@fc3b49fc03e3:/workspaces/did-basic/01_cryptograph/04_aries# go run aries.go]
- hyperledger의 aries는 decentralized identity solutions and digital trust를 위한 toolkit이다.
    https://www.hyperledger.org/projects/aries
- VC를 보여주고 있다.
- root@fc3b49fc03e3:/workspaces/did-basic# go get github.com/hyperledger/aries-framework-go/pkg/doc/verifiable
- root@fc3b49fc03e3:/workspaces/did-basic# cd 01_cryptograph/04_aries/
- root@fc3b49fc03e3:/workspaces/did-basic/01_cryptograph/04_aries# go run aries.go