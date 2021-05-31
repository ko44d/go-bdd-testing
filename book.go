package go_bdd_testing

type book struct {
	Title  string
	Author string
	Pages  int64
}

type Book interface {
	CategoryByLength() string
}

func NewBook(title, author string, pages int64) Book {
	return &book{
		Title: title,
		Author: author,
		Pages: pages,
	}
}

func(b *book) CategoryByLength() string {
	if b.Pages >= 300 {
		return "Novel"
	}
	return "Short story"
}