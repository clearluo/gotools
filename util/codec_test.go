package util

import (
	"encoding/base64"
	"fmt"
)

var (
	key3Des = []byte("didong01didong02didong03")
	key     = []byte("ddbs-i8xLiN7jaWDb5")
	// 私钥生成
	// openssl genrsa -out rsa_private_key.pem 2048
	privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAtCp4g/aTGdTqjU9n3g/tJEYGo4m+JwJb4WAMLc+48xMxMl4a
RE1eMcJo1e+e4UZeMgLDiHHAMy93mZS+pSNxyscHRbgAYLCZck3fr3kj6KE4LeDB
8e3UT928/kJJC+k1sNRTOCTgaUlr6EnF5m/+sBAur6bsc+Se+V/eDHzixdptRatx
Sd12+C6Mz3ISGA3tIzJ4knKFH9adNsENJOZ4yR3fUC1jJlIrrmXSmfKKvbGtSEQE
QaprGhQGCHAr3egZKo6+at7OZk7UbUl+D0tNV/xsjYXa1YewbPXoy/oxqD63YEPU
xLfNmVyDXfE2kfsaOKAyDGT/e+6yIfSkXOYO+wIDAQABAoIBABHdwFO9yPqfzZ6z
yCYSSD73nsLq7utpZXklEDAe0EVymsdW453wWi3vDHxipcvLMZ4d/gvy089/nomh
YYJ40Fj+ga8dPxRXju/x3wPErJXuHucVycXXAt3krFA+fIXs24EsnRANQB/Qwx2m
6zQd2RDWmfqbQAmCG225Xwi3/bkG9IQUZH8JELxxh8xf55B6BkCcfEV31Fssdj0G
zEPS23lp5iwyKuKatAyZkfLUvZb7h9RLL5KFC7uTnfkd++qitbVmVkFsdcZEa3Ir
sg+eaTNEzyKxnu8nIt3j8woiiTEquoaK+pTeqS17GD+zAUAW7WkE8h0VxXgf4iTQ
+H0BSwECgYEA2XXsoVRdRhaJ6WUWHBpGsWo41WCZ6xnOMvcSyNWKtBr/Us7mJ9wT
QRBVPtSFoGEF51SiHl3GczV86gCWfNBD5cfYpdsYl95ORKBf4sRd6bGAGujMn9oi
P4u+Mr1iuMrEmlmuPkjmVzRUKpYfJt6PxMt7nb/BYqyiLFG5wfQZdoECgYEA1BiA
6CraXX2yzzi9GchNSlnRTllZ2XFC7u/LhcrN1PorAzmI0FpzmG/BneD4xnLX29wd
lwqqfMXF4kteHSpLC8sscGmJgRG0UAsGqvUx0WfGecNCzBsuh5FXL+OYAr6E4Q/6
lKNlc42ODKWAU+HkrTrx9QCUiyCgdCOJatwQn3sCgYBXW0b2vCRIHo/CQYhzO58A
cFJqbUcHqbMqyBQ0t4vjtCCzTEgq5P6bGYuVFNylQ7SSbG4/0p4A9BC9FAVgGG3e
Jb0DS5OClpxMdzxtpUKwuUxkAvcIlCFD88gxK+E3qMT32GTlwnU9vNi+ztWu0KNi
g/ehtEFkeUMgmKgNoRiFgQKBgE5zVdMKbsTgBrCxavjLZxNWT54sXJiaVUit99jg
H+xkMF67/Egc/N7oj3RHT52Pwxo2u9cvgcovGTfP6trc1u9g0mouD0dndguZWHkJ
wsiTGw3U1LNMZpSMhPRYudRiBiJ9V5F9MrxgIqe429OrHXuZ7v9RnKAtjEwJDP4y
sg9bAoGADS/LMA5KQFrRlcnMaNv1FuiuYTJkqgiXyAtSVTlbxAHYL9Y1sOYuwIHa
pZS/Y1cPf7p03UNO0vYneaBU6I8nKCduxSEFIt4kHj1og//YgAkLgmmdNn8QP3rP
MI7sEzda19M4E3Ns9EX8FgBCh/jmeOugJ2AQV6DBMzu0KTECYhs=
-----END RSA PRIVATE KEY-----
`)
	// 公钥：根据私钥生成
	// openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
	publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtCp4g/aTGdTqjU9n3g/t
JEYGo4m+JwJb4WAMLc+48xMxMl4aRE1eMcJo1e+e4UZeMgLDiHHAMy93mZS+pSNx
yscHRbgAYLCZck3fr3kj6KE4LeDB8e3UT928/kJJC+k1sNRTOCTgaUlr6EnF5m/+
sBAur6bsc+Se+V/eDHzixdptRatxSd12+C6Mz3ISGA3tIzJ4knKFH9adNsENJOZ4
yR3fUC1jJlIrrmXSmfKKvbGtSEQEQaprGhQGCHAr3egZKo6+at7OZk7UbUl+D0tN
V/xsjYXa1YewbPXoy/oxqD63YEPUxLfNmVyDXfE2kfsaOKAyDGT/e+6yIfSkXOYO
+wIDAQAB
-----END PUBLIC KEY-----

`)
)

func estDes(key []byte) {

	result, err := DesEncrypt([]byte("longlongtextlongtesttagshjasjas asaasldksd"), key)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	origData, err := DesDecrypt(result, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))
}
func test3Des(key3Des []byte) {

	result, err := TripleDesEncrypt([]byte("longlongtextlongtesttagshjasjas asaasldksd2"), key3Des)
	if err != nil {
		panic(err)
	}
	passwd := Base64Encode(string(result))
	fmt.Println(passwd)
	c, err := Base64Decode(passwd)
	origData, err := TripleDesDecrypt([]byte(c), key3Des)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))
}
