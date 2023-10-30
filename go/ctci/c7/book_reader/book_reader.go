package bookreader

type BookReaderSystem struct {
	library     *Library
	userManager *UserManager
	reader      *Reader

	activeUser *User
	activeBook *Book
}

func (b *BookReaderSystem) Login(id int64) bool {
	// only one user at a time
	if b.activeUser != nil && b.activeUser.ID != id {
		return false
	}

	user := b.userManager.FindUser(id)
	if user == nil {
		return false
	}

	b.activeUser = user
	return true
}

func (b *BookReaderSystem) FindBook(id int64) *Book {
	return b.library.FindBook(id)
}

func (b *BookReaderSystem) ReadBook(id int64) bool {
	if b.activeUser == nil || b.activeBook != nil {
		return false
	}
	book := b.library.FindBook(id)
	if book == nil {
		return false
	}

	b.activeBook = book

	if b.activeUser.Membership == nil || !b.activeUser.Membership.IsActive() {
		b.userManager.CreateSubscription(b.activeUser.ID)
	}

	return true
}

func (b *BookReaderSystem) FinishReading(id int64, user *User) {
	b.activeBook = nil
}

func (b *BookReaderSystem) Logout() {
	b.activeUser = nil
}
