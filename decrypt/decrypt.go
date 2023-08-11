package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func DecryptPassword(text string, MySecret string) (string, error) {
	cipherText, decodingerr := base64.StdEncoding.DecodeString(text)
	if decodingerr != nil {
		return "", decodingerr
	}

	block, blockerr := aes.NewCipher([]byte(MySecret))
	if blockerr != nil {
		return "", blockerr
	}

	if len(cipherText) < aes.BlockSize {
		return "", blockerr
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}
