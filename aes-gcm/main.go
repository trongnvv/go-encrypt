package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

var key []byte
var nonce []byte

func aesGcm() cipher.AEAD {
	//key, err := hex.DecodeString(aesKey)
	//if err != nil {
	//	panic(err.Error())
	//}

	//fmt.Println(string(key))
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err.Error())
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	return gcm
}

func encrypt(plaintext []byte) []byte {
	//nonce := make([]byte, aesGcm().NonceSize())
	//dst := make([]byte, 30)
	//if _, err := io.ReadFull(rand.Reader, dst); err != nil {
	//	panic(err.Error())
	//}
	cipherText := aesGcm().Seal(nil, nonce, plaintext, nil)
	return cipherText
}

func decrypt(textEncrypted []byte) string {
	//nonceSize := aesGcm().NonceSize()
	//if len(textEncrypted) <= nonceSize {
	//	panic("cipher text wrong")
	//}
	//_, ciphertext := textEncrypted[:30], textEncrypted[30:]
	//fmt.Println(nonce, ciphertext)
	plaintext, err := aesGcm().Open(nil, nonce, textEncrypted, nil)
	if err != nil {
		panic("open fail")
	}
	return string(plaintext)
}

func main() {
	key = make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		panic(err.Error())
	}
	nonce = make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	encrypted := encrypt([]byte("safdfdas"))
	fmt.Println("encrypted", encrypted)
	decrypted := decrypt(encrypted)
	fmt.Println("decrypted", decrypted)
}
