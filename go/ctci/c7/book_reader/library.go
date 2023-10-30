package bookreader

type Library struct {
	books map[int64]Book
}

func (l *Library) FindBook(id int64) *Book {
	book, ok := l.books[id]
	if !ok {
		return nil
	}

	return &book
}

func (l *Library) AddBook(book Book) {
	if _, ok := l.books[book.ID]; ok {
		// already exists
		return
	}
	l.books[book.ID] = book
}

func (l *Library) RemoveBook(id int64) {
	delete(l.books, id)
}
