package reversedpolishnotation

import (
	"leetcode/go/pkg/stack"
	"strconv"
)

const (
	division = 0
	multiply = 1
	sum      = 2
	diff     = 3
)

func evalRPN(tokens []string) int {
	operators := map[string]int{
		"/": division,
		"*": multiply,
		"+": sum,
		"-": diff,
	}

	st := stack.Stack{}

	for _, t := range tokens {
		operator, ok := operators[t]
		if !ok {
			intVal, err := strconv.Atoi(t)
			if err != nil {
				panic(err)
			}
			st.Push(intVal)
			continue
		}

		r := st.Pop()
		l := st.Pop()
		res := 0

		switch operator {
		case division:
			res = l / r
		case multiply:
			res = l * r
		case sum:
			res = l + r
		case diff:
			res = l - r
		}

		st.Push(res)
	}

	return st.Pop()
}
