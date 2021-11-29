package vo

type UserVO struct {
	UserId   int64  "json:'userid'"
	UserName string "json:'username'"
	Age      int8   "json:'age'"
	Gender   int8   "json:'gender'"
}
