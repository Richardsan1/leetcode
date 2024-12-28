package main

type Stack struct {
	items []byte
	size  int
}

func (s *Stack) Push(data byte) {
	s.items = append(s.items, data)
	s.size += 1
}

func (s *Stack) Pop() byte {
	if s.size == 0 {
		return 0
	}
	item := s.items[s.size-1]
	s.items = s.items[:s.size-1]
	s.size -= 1
	return item
}

func (s *Stack) Top() byte {
	if s.size == 0 {
		return 0
	}
	return s.items[s.size-1]
}

func isValid(s string) bool {
	var pilha Stack

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			pilha.Push(s[i])
		} else {
			if s[i] == ')' && pilha.Top() != '(' {
				return false
			} else if s[i] == ']' && pilha.Top() != '[' {
				return false
			} else if s[i] == '}' && pilha.Top() != '{' {
				return false
			} else {
				pilha.Pop()
			}
		}
	}

	return pilha.size == 0
}

// func main() {
// 	fmt.Println(isValid("()"))
// }
