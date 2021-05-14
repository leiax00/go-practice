//+build wireinject
// 调用wire生成 wire_gen.go
//go:generate wire

package try

import (
	"github.com/google/wire"
)

func InitializeEvent(msg string, code int) (*Event, error) {
	wire.Build(EventSet)
	return &Event{}, nil
}
