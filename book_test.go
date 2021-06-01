package go_bdd_testing

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Book", func() {

	Describe("Categorizing book length", func() {
		var (
			longBook  Book
			shortBook Book
		)

		BeforeEach(func() {
			longBook = NewBook("title1", "author1", 1000)
			shortBook = NewBook("title2", "author2", 10)
		},
		)
		Context("With more than 300 pages", func() {
			It("should be a novel", func() {
				actual := longBook.CategoryByLength()
				Expect(actual).To(Equal("Novel"))
			})
		})
		Context("With fewer than 300 pages", func() {
			It("should be a short story", func() {
				actual := shortBook.CategoryByLength()
				Expect(actual).To(Equal("Short story"))
			})
		})
	},
	)

	Describe("loading from JSON", func() {
		var (
			book Book
			err  error
		)
		BeforeEach(func() {
			book, err = NewBookFromJSON(`{
            "title":"Les Miserables",
            "author":"Victor Hugo",
            "pages":2783
			}`)
		})

		Context("when the JSON parses successfully", func() {
			It("should populate the fields correctly", func() {
				Expect(book.GetTitle()).To(Equal("Les Miserables"))
				Expect(book.GetAuthor()).To(Equal("Victor Hugo"))
				Expect(book.GetPages()).To(Equal(int64(2783)))
			})
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})
})
