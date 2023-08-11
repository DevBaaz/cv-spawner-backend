package cv

import "github.com/lib/pq"

const MySecret = "abc&1*~#^2^#s0^=)^^7%b34"

type EntryReq struct {
	Username string
	Password string
}

type Query struct {
	Username string
}

type Fcv struct {
	Username string
	Fcv      []string
}
type Tcv struct {
	Username string
	Tcv      []string
}

type UserLog struct {
	Username string `gorm:"primaryKey"`
	Password string
	Fcv      pq.StringArray `gorm:"type:string[]"`
	Tcv      pq.StringArray `gorm:"type:string[]"`
}
