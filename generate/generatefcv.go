package generate

import (
	"cvgo/conn"
	"cvgo/cv"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GenerateFcv(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
	DB := conn.Connect()

	rbody, rbodyerr := ioutil.ReadAll(r.Body)
	if rbodyerr != nil {
		w.Write([]byte("Bad Request"))
	}

	var body cv.Query
	if len(rbody) > 0 {
		if err := json.Unmarshal(rbody, &body); err != nil {
			w.Write([]byte("Data Processing Error 1"))
		}
	}
	defer r.Body.Close()

	var userfcv cv.UserLog
	if err := DB.Where("username = ?", body.Username).Find(&userfcv).Error; err != nil {
		w.Write([]byte("Can't Access Your Data, Check Your Internet Connection"))
	}

	if userfcv.Username == body.Username {
		jsonretrieveduserslog, err := json.Marshal(userfcv)
		if err != nil {
			w.Write([]byte("Data Processing Error 2"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonretrieveduserslog))
	} else {
		w.Write([]byte("Unable to Genarate FCV Data"))
	}
}
