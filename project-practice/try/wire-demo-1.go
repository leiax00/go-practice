// IDE忽略冲突
//+build wireinject
// 调用wire生成 wire_gen.go
//go:generate wire

package try

import "github.com/google/wire"

type Fooer interface {
	Foo() string
}

type MyFooer string

func (b *MyFooer) Foo() string {
	return string(*b)
}

func provideMyFooer() *MyFooer {
	b := new(MyFooer)
	*b = "Hello, World!"
	return b
}

type Bar string

func provideBar(f Fooer) string {
	// f will be a *MyFooer.
	return f.Foo()
}

// WireSet 函数调用链
var WireSet = wire.NewSet(
	provideMyFooer,
	wire.Bind(new(Fooer), new(*MyFooer)),
	provideBar,
)

func InitializeWire1() (string, error) {
	wire.Build(WireSet)
	return "", nil
}
