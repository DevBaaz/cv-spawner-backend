package log

import (
	"cvgo/conn"
	"cvgo/cv"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/lib/pq"
)

func LogInFcv(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
	DB := conn.Connect()

	rbody, rbodyerr := ioutil.ReadAll(r.Body)
	if rbodyerr != nil {
		w.Write([]byte("Bad Request"))
	}

	var body cv.Fcv
	if len(rbody) > 0 {
		if err := json.Unmarshal(rbody, &body); err != nil {
			w.Write([]byte("Data Processing Error"))
		}
	}
	defer r.Body.Close()

	user := cv.UserLog{Username: body.Username}
	if err := DB.Model(&user).Update("fcv", pq.StringArray(body.Fcv)).Error; err != nil {

		w.Write([]byte("Can't Update Your Data, Check Your Internet Connection"))
	}

	var userfcv cv.UserLog
	if err := DB.Where("username = ?", body.Username).Find(&userfcv).Error; err != nil {
		w.Write([]byte("Can't Access Your Data, Check Your Internet Connection"))
	}

	if userfcv.Username == body.Username {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("FCV Data Updated"))
	} else {
		w.Write([]byte("Unable to Upload FCV Data"))
	}
}
