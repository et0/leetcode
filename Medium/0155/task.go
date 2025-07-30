package lm0155 // https://leetcode.com/problems/min-stack/

type element struct {
	element int
	min     int
}

type MinStack struct {
	data []*element
}

func Constructor() MinStack {
	return MinStack{
		data: make([]*element, 0),
	}
}

func (s *MinStack) Push(val int) {
	newElement := &element{element: val, min: val}

	size := len(s.data)
	if size == 0 {
		s.data = append(s.data, newElement)
		return
	}

	if val > s.data[size-1].min {
		newElement.min = s.data[size-1].min
	}

	s.data = append(s.data, newElement)
}

func (s *MinStack) Pop() {
	size := len(s.data)
	if size == 0 {
		return
	}

	s.data = s.data[:size-1]
}

func (s *MinStack) Top() int {
	return s.data[len(s.data)-1].element
}

func (s *MinStack) GetMin() int {
	return s.data[len(s.data)-1].min
}
