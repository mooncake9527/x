package encryptUtil

import (
	"crypto/aes"
	"crypto/cipher"
	"github.com/mooncake9527/x/xerrors/xerror"
)

//func main() {
//	key := []byte("myverystrongpasswordo32bitlength") // 32 bytes key for AES-256
//	plaintext := []byte("Hello, World!")
//
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		panic(err)
//	}
//
//	ciphertext := make([]byte, len(plaintext))
//	stream := cipher.NewCTR(block, make([]byte, block.BlockSize()))
//	stream.XORKeyStream(ciphertext, plaintext)
//
//	fmt.Printf("Encrypted: %s\n", base64.StdEncoding.EncodeToString(ciphertext))
//
//	decrypted := make([]byte, len(ciphertext))
//	stream = cipher.NewCTR(block, make([]byte, block.BlockSize()))
//	stream.XORKeyStream(decrypted, ciphertext)
//
//	fmt.Printf("Decrypted: %s\n", decrypted)
//}

func AESEncrypt(plaintext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, xerror.New(err.Error())
	}
	ciphertext := make([]byte, len(plaintext))
	stream := cipher.NewCTR(block, make([]byte, block.BlockSize()))
	stream.XORKeyStream(ciphertext, plaintext)
	return ciphertext, nil
}

func AESDecrypt(encrypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, xerror.New(err.Error())
	}
	decrypted := make([]byte, len(encrypted))
	stream := cipher.NewCTR(block, make([]byte, block.BlockSize()))
	stream.XORKeyStream(decrypted, encrypted)
	return decrypted, nil
}

//
//// ECB PKCS5 加密
//func AESEncrypt(src, key []byte) []byte {
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		fmt.Println("txn put fail: ", err)
//		return nil
//	}
//	ecb := NewECBEncrypter(block)
//	content := []byte(src)
//	content = PKCS5Padding(content, block.BlockSize())
//	des := make([]byte, len(content))
//	ecb.CryptBlocks(des, content)
//	return des
//}
//
//func AESDecrypt(crypted, key []byte) []byte {
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		fmt.Println("err is:", err)
//	}
//	ecb := NewECBDecrypter(block)
//	origData := make([]byte, len(crypted))
//	ecb.CryptBlocks(origData, crypted)
//	origData = PKCS5UnPadding(origData)
//	fmt.Println("source is :", origData, string(origData))
//	return origData
//}
//
//func AESSHA1PRNG(keyBytes []byte, encryptLength int) []byte {
//	hashs := SHA1(SHA1(keyBytes))
//	maxLen := len(hashs)
//	realLen := encryptLength / 8
//	if realLen > maxLen {
//		return nil
//	}
//
//	return hashs[0:realLen]
//}
//
//func SHA1(data []byte) []byte {
//	h := sha1.New()
//	h.Write(data)
//	return h.Sum(nil)
//}
//
//func EncryptAES(data, aesKey string) string {
//	// w4sB8ExNU7wD1xczU+Y/vg==
//	keyAsArray := getKeyByStr(aesKey)
//	keyAsArray = AESSHA1PRNG(keyAsArray, 128)
//	//encryptedByteArray:=AESEncrypt([]byte("ssss"),keyAsArray[0:len(keyAsArray)/2])
//	encryptedByteArray := AESEncrypt([]byte(data), keyAsArray)
//	encryptedByteArrayAsStr := base64.StdEncoding.EncodeToString(encryptedByteArray)
//	//encryptedByteArrayAsStr=encryptedByteArrayAsStr[0:len(encryptedByteArrayAsStr)-3]
//	fmt.Println(encryptedByteArrayAsStr)
//	return encryptedByteArrayAsStr
//}
//
//func DecryptAES(data, aesKey string) string {
//	keyAsArray := getKeyByStr(aesKey)
//	keyAsArray = AESSHA1PRNG(keyAsArray, 128)
//	encryptedByteArray, _ := base64.StdEncoding.DecodeString(data)
//	fmt.Println(encryptedByteArray)
//	//fmt.Println(string(encryptedByteArray))
//	content := AESDecrypt(encryptedByteArray, keyAsArray)
//	encryptedByteArrayAsStr := string(content)
//	fmt.Println(encryptedByteArrayAsStr)
//	return encryptedByteArrayAsStr
//}
