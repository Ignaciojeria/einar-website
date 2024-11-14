package usecase

import (
	"context"

	ioc "github.com/Ignaciojeria/einar-ioc/v2"
)

type ProcessChatMessage func(ctx context.Context, input interface{}) (interface{}, error)

func init() {
	ioc.Registry(NewProcessChatMessage)
}

func NewProcessChatMessage() ProcessChatMessage {
	return func(ctx context.Context, input interface{}) (interface{}, error) {
		return input, nil
	}
}
