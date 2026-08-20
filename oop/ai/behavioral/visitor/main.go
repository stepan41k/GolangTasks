package main

type Shape interface {
	Accept(Visitor)
}

type Visitor interface {
	VisitForSquare(*Square)
	VisitForCircle(*Circle)
}

type Square struct {
	side int
}

type Circle struct {
	radius int
}

func (s *Square) Accept(v Visitor) {
	v.VisitForSquare(s)
}

type AreaCalculator struct {
	totalArea float64
}

func (a *AreaCalculator) VisitForSquare(s *Square) {
	a.totalArea += float64(s.side * s.side)
}