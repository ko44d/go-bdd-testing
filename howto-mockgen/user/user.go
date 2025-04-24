package user

import "ko44d/go-bdd-testing/howto-mockgen/book"

type user struct {
	book book.Book
}

type User interface {
	UseBookTitle() string
}

func NewUser(b book.Book) User {
	return &user{
		book: b,
	}
}

func (u *user) UseBookTitle() string {
	return u.book.GetTitle()
}
