package ginkgo_example

type Book struct {
	Title  string
	Author string
	Pages  int64
}

func (b *Book) CategoryByLength() string {
	if b.Pages >= 300 {
		return "NOVEL"
	}
	return "SHORT STORY"
}

func (b *Book) DoNothing(ok bool) bool {
	return ok
}
