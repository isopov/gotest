package main


func main() {
	p1 := PointerFoo{}
	p1.Foo()

	var pi InterfaceFoo
	//pi = p1 //does not compile
	//pi.Foo()

	p2 := &PointerFoo{}
	p2.Foo()

	pi = p2
	pi.Foo()

	v1 := ValueBar{}
	v1.Bar()

	var vi InterfaceBar
	vi = v1
	vi.Bar()

	v2 := &ValueBar{}
	v2.Bar()

	vi = v2
	vi.Bar()
}

type PointerFoo struct{}

func (p *PointerFoo) Foo() {
	println("foo")
}

type InterfaceFoo interface {
	Foo()
}

type ValueBar struct{}

func (v ValueBar) Bar() {
	println("bar")
}

type InterfaceBar interface {
	Bar()
}