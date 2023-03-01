package encr

import (
	"bytes"
	"crypto/des"
	"encoding/base64"
	"errors"
)

//wcL+kZel294= 666

func desECBEncrypt(data, key []byte) []byte {
	//NewCipher创建一个新的加密块
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	bs := block.BlockSize()
	// pkcs5填充
	data = zeroPadding(data, bs)
	if len(data)%bs != 0 {
		return nil
	}

	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		//Encrypt加密第一个块，将其结果保存到dst
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return out
}

func desECBDecrypter(data, key []byte) []byte {
	//NewCipher创建一个新的加密块
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	bs := block.BlockSize()
	if len(data)%bs != 0 {
		return nil
	}

	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		//Encrypt加密第一个块，将其结果保存到dst
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}

	// pkcs5填充
	out = zeroUnPadding(out)

	return out
}

// zero补码算法
func zeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

// zero减码算法
func zeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}

func ECBEncrypt(key string, str string) (string, error) {
	ebckey := []byte(key)
	var lstr []byte

	lstr = desECBEncrypt([]byte(str), ebckey)
	return base64.StdEncoding.EncodeToString(lstr), nil
}

func ECBDecrypter(key string, str string) (string, error) {
	ebckey := []byte(key)
	var lstr []byte
	bt, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	lstr = desECBDecrypter(bt, ebckey)
	if string(lstr) == "" {
		return "", errors.New("des ECBDecrypter  error")
	}
	return string(lstr), nil
}
