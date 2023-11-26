package utils

type Stack []int

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Len() int {
	tmp := *s
	return len(tmp)
}

func (s *Stack) Push(data int) {
	*s = append(*s, data)
}

func (s *Stack) Pop() {
	tmp := *s
	if len(tmp) > 0 {
		tmp = tmp[:len(tmp)-1]
	}
	*s = tmp
}
