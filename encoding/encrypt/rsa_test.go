package encryptUtil

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"sort"
	"testing"
)

func TestRSAEncrypt(t *testing.T) {
	message := "hello world"
	encryptMessage, err := RSAEncrypt(message, pubKey)
	if err != nil {
		t.Error(err)
	}
	decryptMessage, err := RSADecrypt(encryptMessage, priKey)
	if err != nil {
		t.Error(err)
	}
	if decryptMessage != message {
		t.Errorf("decrypt message is %s, expect %s", decryptMessage, message)
	}
}

const priKey = `
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA06YDe1Qc7EIZUffYiiKN7Nud+EcLomjE4NEPSj/Yv/vcmfLU
u6GKJwK2tCp7XYQvOvBSa3StjlwiBtulSG+yrCD+8hUDW8Hpv+MoNppLnM3f1HVj
NWY1ajwIAQW4OO44VYYaFuebq8Fq7pU/BD6B1qqzLY36LD9DL3pk/kNKnoMjTFxr
yc6UKA2MDjkhaIQIQsdJAePhJ2jeZl3FNzL2Sg+/u7nkYkUVVxk1zan1CnVJoHwu
V0zpnhPJA9LeOqJz6lqicd/HDqnHJpiNT9LcT5FeNbIWXeg6phyqvDp082415CbK
X6Z5nBrBXVvFJ9tGKFSr/5xONUBP5DwTjPAqiwIDAQABAoIBAGKyznja1soOYQRq
kh6I0qqcF+TKLNDGDtnQZFL5xzhwWYWpSi9du7FJgK4wCWuo0uWnYKgftzfmGzAb
ic5n2GiQ0pNN3U0S9tC3O++KsKwlDbZkL6mdSleKOylO44QZA0hOyyfXRi8oeZdd
Hg/6nk3fOMOLrBiqP8iKSUKRWuDj6LE5yrZBmqafvvOXyAhtKVJ7G9iFRBOvxRSC
vrKmZDc4Iz/LtwyLRm/ii8NmQ/Vm1vtLJR+XtIdYG+WrAyzLG8KKts3oSz/NB1NO
i/lwwZotRksZtjTD0aeUKMX0j0XqTub9mrtj/efAWf9SNHBJDq48/uhhNGnRPCFj
V8GocxECgYEA1DrjGRkbDOMjIAS7NsEjnKZjUHbgLiGmH9vIiFhPp/+f4wlYtKC3
LLpl8TFqgJ/HTm9bfCRPMZF3VI5Sx3SsT1+bxOJJrAT9upyExLQnq1zsHOQtjO3D
khF0xUi9A59Itf+Hcc4kHIkJADCQ+t8e8fMZ2j+abCyATX6thmMX5DcCgYEA/0xs
Tnwg1UBwe1I5/K2ZQ8rb9/n+HZiVdHmHsP1V3wWQKFpnBrWfpESBUGNWmIWYP+Tm
ulFw1+eD2fMSavmB8MMMQJTnkKAGwZFQyWVEIsRiNNWdHnvnShVAi43yxOvdFkW3
tPf68Ij5P09NIROQ/36o/jvpDHGb5ATDbjZsqk0CgYB6Cu8DRMuoaomNZQsfnotT
Dt+3qtSZ0qHMWkAEH/yWmEoibgKDxJPxdbMfsxISq08ajLDoP50G3SbpCfsSVcas
0kcqPhKtiCU8hbtXvl29jm784j5Ld4LqYX1r4btH9PYEKtCBolBj1G3HnSYSDfKm
oexw8/hiUmjpp3oz+JIJmwKBgCw5t/VsqV9n11R1rRfplshYpvxxMSU9Xn6b4va1
HCATXaKv7nMKGYqiV2hunPy0/+fpplKWcx7ju0KRShp/+JOVplS1ttul7SWxH7aT
tVb0gDK44ov6WNnLjq/eOjUEyvrlvuo5nx32DH98JFdbhV3NOkc4Z6nBMIkyjgxU
n0RtAoGBAMMvuUjs0z4r0q485btbTSuT0XlhtGhuR/FxCYEZYRXdiG3NCoCwpAv3
8aIXgZ4Vw+W2HdnRMAklSdxxRfi2OHhXvONLJwyzLKrLWrmHlrw0uxJdWfByGqpy
cy7I4YSVe3P2r0zAQ7SUdppCD1uP+F3GHT+FHjSUQJn/2NRCas1h
-----END RSA PRIVATE KEY-----
`

const pubKey = `
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA06YDe1Qc7EIZUffYiiKN
7Nud+EcLomjE4NEPSj/Yv/vcmfLUu6GKJwK2tCp7XYQvOvBSa3StjlwiBtulSG+y
rCD+8hUDW8Hpv+MoNppLnM3f1HVjNWY1ajwIAQW4OO44VYYaFuebq8Fq7pU/BD6B
1qqzLY36LD9DL3pk/kNKnoMjTFxryc6UKA2MDjkhaIQIQsdJAePhJ2jeZl3FNzL2
Sg+/u7nkYkUVVxk1zan1CnVJoHwuV0zpnhPJA9LeOqJz6lqicd/HDqnHJpiNT9Lc
T5FeNbIWXeg6phyqvDp082415CbKX6Z5nBrBXVvFJ9tGKFSr/5xONUBP5DwTjPAq
iwIDAQAB
-----END PUBLIC KEY-----
`

func TestGenerateRsaKey(t *testing.T) {
	pubKey, priKey, err := GenerateRsaKey(2048)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(pubKey)
	fmt.Println(priKey)
}

func TestEncryptStruct(t *testing.T) {
	appId := "tenant1001"
	appSecret := "768eb8b80c5c432c8cdf668f4ea4b5cb"
	timestamp := "1687980647"
	appSecret = SHA256([]byte(appSecret + timestamp))
	aesKey := "0794c045b93b42aeb6dcd1a0c0cfc40a"
	envelop, err := RSAEncrypt(aesKey, pubKey) // envlop
	if err != nil {
		t.Error(err)
	}
	envelopSend := base64.StdEncoding.EncodeToString([]byte(envelop))
	params := map[string]string{
		"publicIp":  "10.10.10.102",
		"system":    "Windows 7",
		"uAversion": "122",
		"kernel":    "chrome",
		"ipChannel": "ipdata",
	}
	var keys []string
	for k, _ := range params {
		keys = append(keys, k)
	}
	data, err := json.Marshal(params)
	if err != nil {
		t.Error(err)
	}
	params2 := map[string]string{
		"appId":     appId,
		"appSecret": appSecret,
		"data":      string(data),
	}
	data2, err := json.Marshal(params2)
	if err != nil {
		t.Error(err)
	}
	signSend := MD5([]byte(appSecret + string(data2))) // sign
	encrypts, err := AESEncrypt(data2, []byte(aesKey))
	if err != nil {
		t.Error(err)
	}
	contentSend := base64.StdEncoding.EncodeToString(encrypts) // request body

	fmt.Println("envelop:", envelopSend)
	fmt.Println("sign:", signSend)
	fmt.Println("content:", contentSend)

	//服务端解开
	envelopBytes, err := base64.StdEncoding.DecodeString(envelopSend)
	if err != nil {
		t.Error(err)
	}
	envelopDecode := string(envelopBytes)
	aesKeyDecode, err := RSADecrypt(envelopDecode, priKey)
	if err != nil {
		t.Error(err)
	}
	contentBytes, err := base64.StdEncoding.DecodeString(contentSend)
	if err != nil {
		t.Error(err)
	}
	contentDec, err := AESDecrypt(contentBytes, []byte(aesKeyDecode))
	if err != nil {
		t.Error(err)
	}
	result := make(map[string]string)
	err = json.Unmarshal(contentDec, &result)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("aesKeyDecode:", aesKeyDecode)
	fmt.Println("result:", result)

	var rkeys []string
	for k, _ := range result {
		rkeys = append(rkeys, k)
	}
	sort.Strings(rkeys)

	resultInner := make(map[string]string)
	err = json.Unmarshal([]byte(result["data"]), &resultInner)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("resultInner:", resultInner)

}
