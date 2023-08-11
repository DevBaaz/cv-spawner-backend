package tests

import (
	"cvgo/auth"
	"github.com/steinfletcher/apitest"
	"testing"
)

func TestAdd(t *testing.T) {
	apitest.New().HandlerFunc(auth.LogIn).Post("/loginuser").JSON(`{"Username":"bola", "Password":"enter"}`).Expect(t).Body("Log In Success").End()
}
