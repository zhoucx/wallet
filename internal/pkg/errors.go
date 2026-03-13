// package pkg It contains some public functions, such as error handling, logging, timing and so on
package pkg

import (
	"encoding/json"
	"fmt"
)

const (
	// LibErrCode Error code of public function
	LibErrCode = 5000
	// WallentNotExistCode Error code of wallet not exit
	WallentNotExistCode = 100000 // 钱包不存在
	WallentNotExistMsg  = "wallent not exist"
	// DestWallentNotExistCode Error code of transaction wallet does not exist
	DestWallentNotExistCode = 100001 // 目的钱包不存在
	DestWallentNotExist     = "dest wallent not exist"
	// AmountNotEnoughCode Error code of amount not enough
	AmountNotEnoughCode = 100100
	AmountNotEnoughMsg  = "Amount not Enough"
)

// ErrorCode Error message returned by the interface
type ErrorCode struct {
	Code    int64
	Message string
}

func (e *ErrorCode) Error() string {
	if e == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%+v", *e)
}

// ToJson ErrorCode to json format
func (e *ErrorCode) ToJson() string {
	if e == nil {
		return "{}"
	}
	buf, err := json.Marshal(e)
	if err != nil {
		fmt.Println("")
		return "{}"
	}
	return string(buf)
}

// NewErrCode new a errorCode by code and message
func NewErrCode(code int64, msg string) *ErrorCode {
	return &ErrorCode{
		Code:    code,
		Message: msg,
	}
}

// NewWallentNotExistErr new wallent not exist error code
func NewWallentNotExistErr() *ErrorCode {
	return &ErrorCode{
		Code:    WallentNotExistCode,
		Message: WallentNotExistMsg,
	}
}

// NewDestWallentNotExistErr new destination wallent not exist error code
func NewDestWallentNotExistErr() *ErrorCode {
	return &ErrorCode{
		Code:    DestWallentNotExistCode,
		Message: DestWallentNotExist,
	}
}

// NewAmountNotEnoughErr new amount not enough error code
func NewAmountNotEnoughErr() *ErrorCode {
	return &ErrorCode{
		Code:    AmountNotEnoughCode,
		Message: AmountNotEnoughMsg,
	}
}
