package main

import (
	stack "leetcode/stack"
)

func isValid(s string) bool {
	stack := stack.New()

	for _, char := range s {
		// if char
		if char == '(' || char == '{' || char == '[' {
			stack.Push(char)
		} else {
			stack.Pop()
		}
	}
	return stack.Len() == 0
}
