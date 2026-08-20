package main

type User struct {
	Name string
	Age int
}

type Iterator interface {
	HasNext() bool
	Next() *User
}

type UserIterator struct {
	index int
	users []*User
}

func (u *UserIterator) HasNext() bool {
	return u.index < len(u.users)
}

func (u *UserIterator) Next() *User {
	user := u.users[u.index]
	u.index++
	return user
}
