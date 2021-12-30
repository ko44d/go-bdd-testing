package howto_ginkgo

import "encoding/json"

type book struct {
	Title  string
	Author string
	Pages  int64
}

type Book interface {
	CategoryByLength() string
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

func NewBookFromJSON(object string) (*book, error) {
	var book book
	if err := json.Unmarshal([]byte(object), &book); err != nil {
		return &book, err
	}
	return &book, nil
}

func (b *book) CategoryByLength() string {
	if b.Pages >= 300 {
		return "Novel"
	}
	return "Short story"
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
