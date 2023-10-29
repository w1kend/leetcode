package employee

type Kind uint8

const (
	KindDirector Kind = iota + 1
	KindManager
	KindRespondent
)

type Employee struct {
	ID          int64
	isAvailable bool
	kind        Kind
}

func (e *Employee) MarkAvailable() {
	e.isAvailable = true
}

func (e *Employee) MarkUnavailable() {
	e.isAvailable = false
}

func (e *Employee) IsAvailable() bool {
	return e.isAvailable
}

func (e *Employee) Kind() Kind {
	return e.kind
}
