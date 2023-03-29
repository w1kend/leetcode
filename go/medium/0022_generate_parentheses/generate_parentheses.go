package generateparentheses

import (
	"fmt"
	"leetcode/go/pkg/stack"
	"strings"
)

func generateParenthesis(n int) []string {
	s := stack.Stack[string]{}
	res := make(chan string)

	go calc(0, 0, n, &s, res)
	result := make([]string, 0, n*2)
	for v := range res {
		result = append(result, v)
	}

	return result
}

func calc(openN, closeN, n int, stack *stack.Stack[string], res chan<- string) {
	if openN == closeN && openN == n {
		s := strings.Join(stack.Source(), "")
		fmt.Printf("variant - %s\n", s)
		res <- s
		return
	}

	if openN < n {
		stack.Push("(")
		calc(openN+1, closeN, n, stack, res)
		stack.Pop()
	}

	if closeN < openN {
		stack.Push(")")
		calc(openN, closeN+1, n, stack, res)
		stack.Pop()
	}

	if closeN == openN && closeN == 0 {
		close(res)
	}
}
