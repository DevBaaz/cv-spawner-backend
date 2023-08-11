package conn

import (
	"cvgo/cv"
	"fmt"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "postgres://fcwears_user:cpFfKGaLhVbfs5jmYPqmnkm81gPcm6kQ@dpg-cj3lcgd9aq0e0q7sccm0-a.frankfurt-postgres.render.com/fcwears"
	userslog, userslogerr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if userslogerr != nil {
		fmt.Println("Failed to connect to database")
	}

	userslogerrAM := userslog.AutoMigrate(&cv.UserLog{})
	if userslogerrAM != nil {
		fmt.Println("Failed to migrate to database")
	}

	return userslog
}

func Nil(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")

	w.Write([]byte("It seems you're lost"))
}
