package cipher

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
)

func Encrypt(data []byte, key []byte) ([]byte, []byte, error) {
	iv, err := generateIV()
	if err != nil {
		return nil, nil, fmt.Errorf("%w", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, fmt.Errorf("%w", err)
	}

	padded := pkcs7Pad(data)

	encrypted := make([]byte, len(padded))

	cbcEncrypter := cipher.NewCBCEncrypter(block, iv)

	cbcEncrypter.CryptBlocks(encrypted, padded)

	return iv, encrypted, nil
}

func Decrypt(data []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	decrypted := make([]byte, len(data))

	cbcDecrypter := cipher.NewCBCDecrypter(block, iv)

	cbcDecrypter.CryptBlocks(decrypted, data)

	return pkcs7Unpad(decrypted), nil
}

func generateIV() ([]byte, error) {
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return iv, nil
}

func pkcs7Pad(data []byte) []byte {
	length := aes.BlockSize - (len(data) % aes.BlockSize)

	trailing := bytes.Repeat([]byte{byte(length)}, length)

	return append(data, trailing...)
}

func pkcs7Unpad(data []byte) []byte {
	dataLength := len(data)

	padLength := int(data[dataLength-1])

	return data[:dataLength-padLength]
}
