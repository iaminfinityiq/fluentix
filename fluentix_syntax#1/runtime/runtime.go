package runtime

import (
	"fmt"
)

type Error interface {
	ErrorType() string
	Reason_() string
}

func DisplayError(error_ Error) {
	fmt.Println(error_.ErrorType() + ": " + error_.Reason_())
}

type RuntimeResult struct {
	Result any
	Error  *Error
}

func Success(result any) RuntimeResult {
	return RuntimeResult{
		Result: result,
		Error:  nil,
	}
}

func Failure(error_ Error) RuntimeResult {
	return RuntimeResult{
		Result: nil,
		Error: &error_,
	}
}