package xzip

import (
	"archive/zip"
	"bytes"
	"compress/flate"
	"hash/crc32"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type ZipCrypto struct {
	password []byte
	Keys     [3]uint32
}

func NewZipCrypto(passphrase []byte) *ZipCrypto {
	z := &ZipCrypto{}
	z.password = passphrase
	z.init()
	return z
}

func (z *ZipCrypto) init() {
	z.Keys[0] = 0x12345678
	z.Keys[1] = 0x23456789
	z.Keys[2] = 0x34567890

	for i := 0; i < len(z.password); i++ {
		z.updateKeys(z.password[i])
	}
}

func (z *ZipCrypto) updateKeys(byteValue byte) {
	z.Keys[0] = crc32update(z.Keys[0], byteValue)
	z.Keys[1] += z.Keys[0] & 0xff
	z.Keys[1] = z.Keys[1]*134775813 + 1
	z.Keys[2] = crc32update(z.Keys[2], (byte)(z.Keys[1]>>24))
}

func (z *ZipCrypto) magicByte() byte {
	var t uint32 = z.Keys[2] | 2
	return byte((t * (t ^ 1)) >> 8)
}

func (z *ZipCrypto) Encrypt(data []byte) []byte {
	length := len(data)
	chiper := make([]byte, length)
	for i := 0; i < length; i++ {
		v := data[i]
		chiper[i] = v ^ z.magicByte()
		z.updateKeys(v)
	}
	return chiper
}

func (z *ZipCrypto) Decrypt(chiper []byte) []byte {
	length := len(chiper)
	plain := make([]byte, length)
	for i, c := range chiper {
		v := c ^ z.magicByte()
		z.updateKeys(v)
		plain[i] = v
	}
	return plain
}

func crc32update(pCrc32 uint32, bval byte) uint32 {
	return crc32.IEEETable[(pCrc32^uint32(bval))&0xff] ^ (pCrc32 >> 8)
}

func ZipCryptoDecryptor(r *io.SectionReader, password []byte) (*io.SectionReader, error) {
	z := NewZipCrypto(password)
	b := make([]byte, r.Size())

	r.Read(b)

	m := z.Decrypt(b)
	return io.NewSectionReader(bytes.NewReader(m), 12, int64(len(m))), nil
}

type unzip struct {
	offset int64
	fp     *os.File
	name   string
}

func (uz *unzip) init() (err error) {
	uz.fp, err = os.Open(uz.name)
	return err
}

func (uz *unzip) close() {
	if uz.fp != nil {
		uz.fp.Close()
	}
}

func (uz *unzip) Size() int64 {
	if uz.fp == nil {
		if err := uz.init(); err != nil {
			return -1
		}
	}

	fi, err := uz.fp.Stat()
	if err != nil {
		return -1
	}

	return fi.Size() - uz.offset
}

func (uz *unzip) ReadAt(p []byte, off int64) (int, error) {
	if uz.fp == nil {
		if err := uz.init(); err != nil {
			return 0, err
		}
	}

	return uz.fp.ReadAt(p, off+uz.offset)
}

// DeCompressZip 解压zip包
func DeCompressZip(zipFile, dest, passwd string, offset int64) error {
	uz := &unzip{offset: offset, name: zipFile}
	defer uz.close()

	zr, err := zip.NewReader(uz, uz.Size())
	if err != nil {
		return err
	}

	if passwd != "" {
		// Register a custom Deflate compressor.
		zr.RegisterDecompressor(zip.Deflate, func(r io.Reader) io.ReadCloser {
			rs := r.(*io.SectionReader)
			r, _ = ZipCryptoDecryptor(rs, []byte(passwd))
			return flate.NewReader(r)
		})

		zr.RegisterDecompressor(zip.Store, func(r io.Reader) io.ReadCloser {
			rs := r.(*io.SectionReader)
			r, _ = ZipCryptoDecryptor(rs, []byte(passwd))
			return ioutil.NopCloser(r)
		})
	}

	for _, f := range zr.File {
		fpath := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		inFile, err := f.Open()
		if err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			inFile.Close()
			return err
		}

		_, err = io.Copy(outFile, inFile)
		inFile.Close()
		outFile.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

//
//func main() {
//	err := DeCompressZip("Go.zip", "./tmp", "123456", 0)
//	if err != nil {
//		fmt.Println(err)
//	}
//	return
//}
