package log

import (
	"cvgo/conn"
	"cvgo/cv"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/lib/pq"
)

func LogInTcv(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
	DB := conn.Connect()

	rbody, rbodyerr := ioutil.ReadAll(r.Body)
	if rbodyerr != nil {
		w.Write([]byte("Bad Request"))
	}

	var body cv.Tcv
	if len(rbody) > 0 {
		if err := json.Unmarshal(rbody, &body); err != nil {
			w.Write([]byte("Data Processing Error"))
		}
	}
	defer r.Body.Close()

	user := cv.UserLog{Username: body.Username}
	if err := DB.Model(&user).Update("tcv", pq.StringArray(body.Tcv)).Error; err != nil {
		w.Write([]byte("Can't Update Your Data, Check Your Internet Connection"))
	}

	var usertcv cv.UserLog
	if err := DB.Where("username = ?", body.Username).Find(&usertcv).Error; err != nil {
		w.Write([]byte("Can't Access Your Data, Check Your Internet Connection"))
	}

	if usertcv.Username == body.Username {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("TCV Data Updated"))
	} else {
		w.Write([]byte("Unable to Upload TCV Data"))
	}
}
