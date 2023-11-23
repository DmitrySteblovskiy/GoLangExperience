package main

import "fmt"

type Stack struct {
	data []rune
}

func (s *Stack) Push(item rune) {
	s.data = append(s.data, item)
}

func (s *Stack) Pop() rune {
	if s.IsEmpty() {
		return 0 // Возвращаем нулевое значение для пустого стека
	}
	item := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return item
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func isBalanced(expression string) bool {
	stack := &Stack{}

	for _, char := range expression {
		if char == '(' || char == '[' || char == '{' {
			stack.Push(char)
		} else if char == ')' || char == ']' || char == '}' {
			if stack.IsEmpty() {
				return false // Скобочная последовательность не сбалансирована
			}
			popped := stack.Pop()
			if (char == ')' && popped != '(') ||
				(char == ']' && popped != '[') ||
				(char == '}' && popped != '{') {
				return false // Несоответствие закрывающей и открывающей скобок
			}
		}
	}

	return stack.IsEmpty() // Сбалансированность достигается, если стек пуст
}

func main() {
	expression1 := "{[()()]}"
	expression2 := "{[()]}["

	if isBalanced(expression1) {
		fmt.Println("Выражение 1 сбалансировано.")
	} else {
		fmt.Println("Выражение 1 не сбалансировано.")
	}

	if isBalanced(expression2) {
		fmt.Println("Выражение 2 сбалансировано.")
	} else {
		fmt.Println("Выражение 2 не сбалансировано.")
	}
}
