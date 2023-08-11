package main

import (
	"cvgo/auth"
	"cvgo/conn"
	"cvgo/generate"
	"cvgo/log"

	"fmt"
	"net/http"
)

func main() {
	conn.Connect()

	http.HandleFunc("/", conn.Nil)
	http.HandleFunc("/signupuser", auth.SignUp)
	http.HandleFunc("/loginuser", auth.LogIn)
	http.HandleFunc("/loginusertcv", log.LogInTcv)
	http.HandleFunc("/loginuserfcv", log.LogInFcv)
	http.HandleFunc("/generatetcv", generate.GenerateTcv)
	http.HandleFunc("/generatefcv", generate.GenerateFcv)

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
