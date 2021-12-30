package howto_ginkgo

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

		Context("when the JSON parses successfully", func() {
			BeforeEach(func() {
				book, err = NewBookFromJSON(`{
            "title":"Les Miserables",
            "author":"Victor Hugo",
            "pages":2783
			}`)
			})

			It("should populate the fields correctly", func() {
				Expect(book.GetTitle()).To(Equal("Les Miserables"))
				Expect(book.GetAuthor()).To(Equal("Victor Hugo"))
				Expect(book.GetPages()).To(Equal(int64(2783)))
			})

			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when the JSON fails to parse", func() {
			BeforeEach(func() {
				book, err = NewBookFromJSON(`{
                    "title":"Les Miserables",
                    "author":"Victor Hugo",
                    "pages":2783oops
                }`)
			})
			It("should return the zero-value for the book", func() {
				Expect(book.GetTitle()).To(Equal(""))
				Expect(book.GetAuthor()).To(Equal(""))
				Expect(book.GetPages()).To(Equal(int64(0)))
			})
			It("should error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("invalid character 'o' after object key:value pair"))
			})
		})
	})
})
