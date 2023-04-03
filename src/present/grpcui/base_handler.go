package grpcui

import (
	"base/src/common"
	"fmt"
)

type baseHandler struct {
}

func NewBaseHandler() *baseHandler {
	return &baseHandler{}
}

func (b *baseHandler) IErrorToGRPCError(err *common.Error) error {
	return fmt.Errorf(err.ToJSon())
}
