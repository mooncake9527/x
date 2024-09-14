package encryptUtil

import (
	"encoding/base64"
	"testing"
)

func TestAESEncrypt(t *testing.T) {
	message := []byte("hello world")
	key := []byte("myverystrongpasswordo32bitlength")

	encryptedMessage, err := AESEncrypt(message, key)
	if err != nil {
		t.Error(err)
	}
	encryptedMessageString := base64.StdEncoding.EncodeToString(encryptedMessage)

	//////////

	encryptedMessageBytes, err := base64.StdEncoding.DecodeString(encryptedMessageString)
	if err != nil {
		t.Error(err)
	}
	decryptedMessage, err := AESDecrypt(encryptedMessageBytes, key)
	if err != nil {
		t.Error(err)
	}
	if string(decryptedMessage) != string(message) {
		t.Error("decrypted message is not equal to original message")
	}
}
