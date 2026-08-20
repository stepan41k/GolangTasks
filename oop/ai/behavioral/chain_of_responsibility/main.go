package main

import "fmt"

type Patient struct {
	Name string
	Age int
	RegistrationDone bool
}

type Department interface {
	Execute(*Patient)
	SetNext(Department)
}

type Reception struct {
	next Department
}

func (r *Reception) Execute(p *Patient) {
	if p.RegistrationDone {
		r.next.Execute(p)
	} else {
		fmt.Println("Регистрация пациента..")
		p.RegistrationDone = true
		r.next.Execute(p)
	}
}

func (r *Reception) SetNext(next Department) {
	r.next = next
}