package auth

import (
	"cvgo/conn"
	"cvgo/cv"
	"cvgo/decrypt"
	"cvgo/encrypt"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func LogIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
	DB := conn.Connect()

	rbody, rbodyerr := ioutil.ReadAll(r.Body)
	if rbodyerr != nil {
		w.Write([]byte("Bad Request"))
	}

	var body cv.EntryReq
	if len(rbody) > 0 {
		if err := json.Unmarshal(rbody, &body); err != nil {
			w.Write([]byte("Couldn't Proccess Request"))
		}
	}
	defer r.Body.Close()

	encryptedpassword, _ := encrypt.EncryptPassword(body.Password, cv.MySecret)

	retrieveduserslog := cv.UserLog{}
	if err := DB.Where(&cv.UserLog{Username: body.Username, Password: encryptedpassword}).Find(&retrieveduserslog).Error; err != nil {

		w.Write([]byte("Can't Verify Your Data, Check Your Internet Connection"))
	}

	decryptedpassword, _ := decrypt.DecryptPassword(retrieveduserslog.Password, cv.MySecret)

	if retrieveduserslog.Username == body.Username && decryptedpassword == body.Password {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Log In Success"))
	} else {
		w.Write([]byte("Incorrect Username Or Password"))
	}

}
