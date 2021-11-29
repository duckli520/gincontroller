package constant

type ResponseCode int
type ResponseMsg string

const (
	SelectSuccessCode ResponseCode = 2005
	SelectSuccessMsg  ResponseMsg  = "查询成功"
	DataIsNilCode     ResponseCode = 9999
	DataIsNilMsg      ResponseMsg  = "访问错误"
)
