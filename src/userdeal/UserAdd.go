package userdeal

import "time"

type AddUserInfo struct {
	UserId     int64     "json:'UserId'"
	UserName   string    "json:'UserName'"
	Age        int8      "json:'Age'"
	Gender     int8      "json:'Gender'"
	CreateTime time.Time "json:'create_time'"
}

func (AddUserInfo) TableName() string {
	return "t_user"
}
