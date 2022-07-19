package queue

type Queue interface {
	Top() (interface{}, error)
	Push(interface{}) error
	Pop() error
	Empty() bool
	Size() int
	Clear()
	Values() []interface{}
	String() string
}
