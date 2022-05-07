package crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

type Crypt struct {
	key []byte
	iv  []byte
}

func NewCrypt(key []byte, iv []byte) Crypt {
	return Crypt{key, iv}
}

func (c *Crypt) Encrypt(input []byte) (encrypted []byte, err error) {
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return
	}
	padded := c.padding(input)

	encrypted = make([]byte, len(padded))
	mode := cipher.NewCBCEncrypter(block, c.iv)
	mode.CryptBlocks(encrypted, padded)
	return
}

func (c *Crypt) Decrypt(input []byte) (decrypted []byte, err error) {
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return
	}

	decrypted = make([]byte, len(input))
	mode := cipher.NewCBCDecrypter(block, c.iv)
	mode.CryptBlocks(decrypted, input)
	decrypted = c.trimming(decrypted)
	return
}

func (c *Crypt) padding(input []byte) (padded []byte) {
	blockSize := aes.BlockSize
	padding := (blockSize - len(input)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	padded = append(input, padtext...)
	return
}

func (c *Crypt) trimming(input []byte) (trimmed []byte) {
	padding := input[len(input)-1]
	return input[:len(input)-int(padding)]
}
