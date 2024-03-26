package ginx

var (
	RequestFailed        = NewError(-1, "请求失败")
	ParamsError          = NewError(-2, "参数错误")
	InternalServerError  = NewError(-3, "服务器内部错误")
	UnAuthorizationError = NewError(-401, "用户认证失败")
)

type Error struct {
	cause error
	code  int
	msg   string
}

func (e Error) Error() string {
	if e.cause != nil {
		return e.cause.Error()
	}

	return e.msg
}

func NewError(code int, msg string) error {
	return &Error{
		code: code,
		msg:  msg,
	}
}
