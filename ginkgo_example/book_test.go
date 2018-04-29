package ginkgo_example_test

import (
	"fmt"
	"testing"

	. "github.com/huangping40/fabric_learning/ginkgo_example"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBook(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Book one Suite")
}

var _ = Describe("Book", func() {
	var (
		longBook  Book
		shortBook Book
	)

	BeforeEach(func() {
		fmt.Println("\nBeforeEach")
		longBook = Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  1488,
		}

		shortBook = Book{
			Title:  "Fox In Socks",
			Author: "Dr. Seuss",
			Pages:  24,
		}
	})

	Describe("Categorizing book length", func() {
		Context("With more than 300 pages", func() {
			It("should be a novel", func() {
				By("Documenting Complex It s")

				Expect(longBook.CategoryByLength()).To(Equal("NOVEL"))
			})
		})

		Context("With fewer than 300 pages", func() {
			It("should be a short story", func() {
				Expect(shortBook.CategoryByLength()).To(Equal("SHORT STORY"))
			})
		})

		Describe("panic test", func() {
			XIt("panics in a goroutine", func(done Done) {
				go func() {
					defer GinkgoRecover()
					Ω(shortBook.DoNothing(true)).Should(BeTrue())
					close(done)
				}()
			})
		})
	})

})
