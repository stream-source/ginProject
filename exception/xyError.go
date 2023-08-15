package exception

import (
	"fmt"
	"time"
)

// 私有化，外部不能直接构造该错误类型
type BaseError struct {
	Code      string
	Request   string
	ErrorTime time.Time
	Msg       string
}

func (err BaseError) Error() string {
	return fmt.Sprintf("code:%s,request：%s, msg:%s", err.Code, err.Request, err.Msg)
}

func DefineBaseError(code string, request string, msg string) error {
	return BaseError{
		Code:      code,
		Request:   request,
		ErrorTime: time.Now(),
		Msg:       msg,
	}
}

func (err BaseError) GetBaseError() BaseError {
	return err
}
