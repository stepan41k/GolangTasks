package main

type Train struct {
	Number string
}

type Mediator interface {
	CanLand(Train) bool
	NotifyFree()
}

type StationManager struct {
	isPlatformFree bool
}

func (s *StationManager) CanLand(t Train) bool {
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	
	return false
}