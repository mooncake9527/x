package encryptUtil

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

// MD5
func MD5(data []byte) string {
	h := md5.Sum(data)
	return hex.EncodeToString(h[:])
}

// sha256
func SHA256(data []byte) string {
	h := sha256.Sum256(data)
	return hex.EncodeToString(h[:])
}

/*
获取文件的MD5
*/
func MD5File(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		//fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
		return ""
	}
	md5 := md5.New()
	io.Copy(md5, f)
	MD5Str := hex.EncodeToString(md5.Sum(nil))
	f.Close()
	return MD5Str
}
