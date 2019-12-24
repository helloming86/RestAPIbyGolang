package errno

import (
	"fmt"
)

type Errno struct {
	Code int
	Message string
}

func (err Errno) Error() string {
	return err.Message
}

// Err represents an error
type Err struct {
	Code int
	Message string
	Err error
}

func (err *Err) Error() string  {
	// 返回字符串
	// 包含 自定义的 返回码code、自定义的返回信息message、系统的错误类型error
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s",
		err.Code, err.Message, err.Err)
}

// 对外展示更多的信息可以调用此函数
func (err *Err) Add(message string) error  {
	err.Message += " " + message
	return err
}

// 对外展示更多的信息可以调用此函数
func (err *Err) Addf(format string, args ...interface{}) error  {
	// fmt.Sprintf() 返回为 格式化后的字符串
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}

// 新建定制的错误
// 为系统error定制对应的自定义返回码code和返回信息message
func New(errno *Errno, err error) *Err  {
	return &Err{Code: errno.Code, Message: errno.Message, Err: err}
}

// 解析定制的错误
// 要根据系统error解析得到对应的自定义错误返回
// 函数DecodeErr的入参是一个error类型 根据go语言的鸭子类型特性，结构体类型Errno和Err也算是error类型
func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *Errno:
		return typed.Code, typed.Message
	default:

	}
	return InternalServerError.Code, err.Error()
}

func IsErrUserNotFound(err error) bool {
	code, _ := DecodeErr(err)
	return code == ErrUserNotFound.Code
}






