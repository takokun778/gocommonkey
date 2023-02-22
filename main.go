package main

import (
	"encoding/base64"
	"encoding/hex"
	"gocommonkey/cipher"
	"log"

	"github.com/google/uuid"
)

func main() {
	text := uuid.NewString()

	log.Println("Input:", text)

	plain := []byte(text)

	keyString := "645E739A7F9F162725C1533DC2C5E827"

	key, _ := hex.DecodeString(keyString)

	log.Println("Key:", keyString)

	iv, encrypted, _ := cipher.Encrypt(plain, key)

	log.Println("IV:", hex.EncodeToString(iv))

	log.Println("Encrypted:", base64.StdEncoding.EncodeToString(encrypted))

	decrypted, _ := cipher.Decrypt(encrypted, key, iv)

	log.Printf("Decrypted: %s", decrypted)
}
