package auth

import (
	"cvgo/conn"
	"cvgo/cv"
	"cvgo/encrypt"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
	DB := conn.Connect()

	rbody, rbodyerr := ioutil.ReadAll(r.Body)
	if rbodyerr != nil {
		fmt.Println("Sign Up decoding err")
	}

	var body cv.EntryReq
	if len(rbody) > 0 {
		if err := json.Unmarshal(rbody, &body); err != nil {
			fmt.Println("Sign Up unmarshall error")
		}
	}

	defer r.Body.Close()

	var retrieveduser cv.UserLog
	if err := DB.Where(&cv.UserLog{Username: body.Username}).Find(&retrieveduser).Error; err != nil {
		w.Write([]byte("Can't Access the Database, Check Your Internet Connection"))
	}

	if retrieveduser.Username == "" {
		encryptedpassword, _ := encrypt.EncryptPassword(body.Password, cv.MySecret)

		user := &cv.UserLog{Username: body.Username, Password: encryptedpassword}

		if err := DB.Create(&user).Error; err != nil {
			w.Write([]byte("Can't Access the Database, Check Your Internet Connection"))
		}

		var usernamelog cv.UserLog
		if err := DB.Where("username = ?", body.Username).Find(&usernamelog).Error; err != nil {
			w.Write([]byte("Can't Access the Database, Check Your Internet Connection"))
		}

		if usernamelog.Username == body.Username {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Sign Up Success"))
		} else {
			w.Write([]byte("Unable to Create Account"))
		}
	} else {
		w.Write([]byte("Invalid Username, choose another"))
	}
}
