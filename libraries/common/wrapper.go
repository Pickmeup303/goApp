package common

import (
	"kaffein/libraries/cryptography/caesarCipher"
	"kaffein/libraries/cryptography/transpose"
)

func WrapperCaesarEncrypt(plainText, alphabet string, keyShifter, keyTransposition int) (string, error) {
	cipher, err := caesarCipher.NewCaesarCipher(plainText, alphabet, keyShifter).Encrypt()
	if err != nil {
		return "", err
	}
	enTranspose := transpose.NewTranspose(cipher, 6, keyTransposition).Encrypt()
	return enTranspose, nil
}

func WrapperCaesarDecrypt(cipherText, alphabet string, keyShifter, keyTransposition int) (string, error) {
	deTranspose := transpose.NewTranspose(cipherText, 6, keyTransposition).Decrypt()

	deCipher, err := caesarCipher.NewCaesarCipher(deTranspose, alphabet, keyShifter).Decrypt()
	if err != nil {
		return "", err
	}

	return deCipher, nil
}
