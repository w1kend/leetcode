package bookreader

type Reader struct {
	book    Book
	pageNum int64
}

func (r *Reader) NextPage() {
	if r.pageNum+1 > r.book.PagesNum {
		// do nothing
		return
	}

	r.pageNum++
	r.render()
}

func (r *Reader) PrevPage() {
	if r.pageNum == 0 {
		// do nothing
		return
	}
	r.pageNum--
	r.render()
}

func (r *Reader) render() { /* render current page */ }
