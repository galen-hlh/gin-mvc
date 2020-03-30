package errors

import "fmt"

type BusinessLib interface {
	String() string
	Error() string
	Code() int32
}

// 业务异常基类
type Business struct {
	Num int32
	Msg string
}

func (e *Business) Error() string {
	return e.Msg
}

func (e *Business) Code() int32 {
	return e.Num
}

func (e *Business) String() string {
	return fmt.Sprintf("Code:%d,Msg:%s", e.Num, e.Msg)
}

func NewError(code int32, msg string) BusinessLib {
	return &Business{
		Num: code,
		Msg: msg,
	}
}
