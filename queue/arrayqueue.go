package queue

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// 判断ArrayQueue是否实现了Config接口，如果 *ArrayQueue 与 Queue 的接口不匹配,
// 那么语句 var _ Queue = (*ArrayQueue)(nil) 将无法编译通过。
// 将实例是否实现接口提前到编译期就表现出来
var _ Queue = (*ArrayQueue)(nil)

type ArrayQueue struct {
	elements []interface{}
	size     int
}

func NewArrayQueue() *ArrayQueue {
	return &ArrayQueue{}
}

func (queue *ArrayQueue) Top() (value interface{}, err error) {
	if queue.size == 0 {
		return nil, errors.New("queue is empty")
	}
	return queue.elements[0], nil
}

func (queue *ArrayQueue) Push(value interface{}) (err error) {
	if queue.size != 0 {
		if reflect.TypeOf(value) != reflect.TypeOf(queue.elements[0]) {
			return errors.New("inconsistent type")
		}
	}
	queue.elements = append(queue.elements, value)
	queue.size++
	return nil
}

func (queue *ArrayQueue) Pop() (err error) {
	if queue.size == 0 {
		return errors.New("queue is empty")
	}

	queue.elements[0] = nil
	copy(queue.elements[0:], queue.elements[1:queue.size])
	queue.size--
	return nil
}

func (queue *ArrayQueue) Empty() (is bool) {
	if queue.size == 0 {
		return true
	}
	return false
}

func (queue *ArrayQueue) Size() (size int) {
	return queue.size
}

func (queue *ArrayQueue) Clear() {
	queue.elements = []interface{}{}
	queue.size = 0
}

func (queue *ArrayQueue) Values() (values []interface{}) {
	return queue.elements
}

func (queue *ArrayQueue) String() string {
	var values []string
	for _, value := range queue.Values() {
		values = append(values, fmt.Sprintf("%v", value))
	}
	return strings.Join(values, ",")
}
