package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func EncryptPassword(text string, MySecret string) (string, error) {
	block, blockerr := aes.NewCipher([]byte(MySecret))
	if blockerr != nil {
		return "", blockerr
	}

	plainText := []byte(text)
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cipherText[aes.BlockSize:], plainText)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}
