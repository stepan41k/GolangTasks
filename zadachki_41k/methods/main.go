package main

type T struct {
	
}

func (T) X() {
}

func (*T) Z() {
}

func main() {
	var t T
	t.X()
	t.Z()
	var p = &t
	p.X()
	p.Z()
	T{}.X()
	// T{}.Z() // Error
	(&T{}).X()
	(&T{}).Z()
}

