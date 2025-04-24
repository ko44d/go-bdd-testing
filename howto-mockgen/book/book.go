//go:generate mockgen -source=book.go -destination=mock/mock_book.go -package=mocks
package book

type book struct {
	Title  string
	Author string
	Pages  int64
}

type Book interface {
	GetTitle() string
	GetAuthor() string
	GetPages() int64
}

func NewBook(title, author string, pages int64) Book {
	return &book{
		Title:  title,
		Author: author,
		Pages:  pages,
	}
}

func (b *book) GetTitle() string {
	return b.Title
}

func (b *book) GetAuthor() string {
	return b.Author
}

func (b *book) GetPages() int64 {
	return b.Pages
}
