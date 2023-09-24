package main

type Inner interface {
	Fn()
}

type Outer interface {
	Inner
}

type Impl struct{}

func (f *Impl) Fn() {}

func Const() *Impl {
	return &Impl{}
}

func test() {
	var inst Outer = Const()
	var inst2 Inner = Const()
	inst.Fn()
	inst2.Fn()
}
