package bookreader

type User struct {
	ID       int64
	Username string
	Email    string

	Membership *Membership
}
