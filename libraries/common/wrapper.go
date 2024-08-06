package common

import (
	"kaffein/libraries/cryptography/caesarCipher"
	"kaffein/libraries/cryptography/transpose"
)

func WrapperCaesarEncrypt(plainText, alphabet string, key int) (string, error) {
	cipher, err := caesarCipher.NewCaesarCipher(plainText, alphabet, key).Encrypt()
	if err != nil {
		return "", err
	}
	enTranspose := transpose.NewTranspose(cipher, 6).Encrypt()
	return enTranspose, nil
}

func WrapperCaesarDecrypt(cipherText, alphabet string, key int) (string, error) {
	deTranspose := transpose.NewTranspose(cipherText, 6).Decrypt()

	deCipher, err := caesarCipher.NewCaesarCipher(deTranspose, alphabet, key).Decrypt()
	if err != nil {
		return "", err
	}

	return deCipher, nil
}
