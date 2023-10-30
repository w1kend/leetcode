package bookreader

import "time"

type Membership struct {
	ID      int64
	StartAt time.Time
	EndAt   time.Time
	User    *User
}

func (m *Membership) IsActive() bool {
	now := time.Now()

	return now.After(m.StartAt) && now.Before(m.EndAt)
}

func (m *Membership) Prolongate() {
	m.EndAt.Add(time.Hour * 24 * 2) // for 2 days
}
