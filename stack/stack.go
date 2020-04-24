package stack

type (
	Stack struct {
		top *node
		length int
	}
	node struct {
		val interface{}
		prev *node
	}
)

func New() *Stack {
	return &Stack{nil,0}
}

func (s *Stack) Len() int {
	return s.length
}

func (s *Stack) Top() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.top.val
}

func (s *Stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}
	n := s.top
	s.top = n.prev
	s.length--
	return n.val
}

func (s *Stack) Push(value interface{}) {
	n := &node{
		val: value,
		prev:  s.top,
	}
	s.top = n
	s.length++
}
