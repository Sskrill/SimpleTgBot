package events

type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}
type Proccesor interface {
	Procces(e Event) error
}
type Event struct {
	Type Type
	Text string
	Meta interface{}
}
type Type int

const (
	Unknouwn Type = iota
	Message
)
