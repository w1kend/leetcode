package bookreader

import "time"

type UserManager struct {
	users map[int64]User
}

func (um *UserManager) FindUser(id int64) *User {
	user, ok := um.users[id]
	if !ok {
		return nil
	}

	return &user
}

func (um *UserManager) CreateSubscription(userID int64) bool {
	user, ok := um.users[userID]
	if !ok {
		return false
	}

	if user.Membership != nil && user.Membership.IsActive() {
		return false
	}

	if user.Membership != nil {
		user.Membership.Prolongate()
		return true
	}

	user.Membership = um.createMembership(&user)

	return true
}

func (um *UserManager) createMembership(user *User) *Membership {
	m := Membership{
		ID:      1,
		StartAt: time.Now(),
		EndAt:   time.Now().Add(time.Hour * 24 * 30),
		User:    user,
	}

	return &m
}
