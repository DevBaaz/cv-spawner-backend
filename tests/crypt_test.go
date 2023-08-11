package tests

import (
	"cvgo/decrypt"
	"cvgo/encrypt"
	"testing"
)

func TestCrypt(t *testing.T) {
	const Key = "abc&1*~#^2^#s0^=)^^7%b34"
	text := "test"
	encrypt, _ := encrypt.EncryptPassword(text, Key)
	decrypt, _ := decrypt.DecryptPassword(encrypt, Key)

	if decrypt != text {
		t.Error("Expected", text, "found", decrypt)
	}
}

/*
PS C:\Users\USER\Desktop\myprojects\cv-spawner\backend\test> go test -v
=== RUN   TestCrypt
--- PASS: TestCrypt (0.00s)
PASS
ok      cvgo/test       0.934s
*/
