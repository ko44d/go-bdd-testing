package go_bdd_testing

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Book", func() {
	var (
		longBook Book
		shortBook Book
	)

	BeforeEach(func() {
		longBook = NewBook("title1", "author1", 1000)
		shortBook = NewBook("title2", "author2", 10)
	},
	)

	Describe("Categorizing book length", func() {
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
})
