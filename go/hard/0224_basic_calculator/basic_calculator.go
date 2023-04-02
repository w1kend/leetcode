package basiccalculator

import (
	"fmt"
	"leetcode/go/pkg/stack"
	"strconv"
	"strings"
)

const (
	lparenthesis = "("
	rparenthesis = ")"
	plus         = "+"
	minus        = "-"
)

func calculate(s string) int {
	fmt.Println("case -", s)
	st := stack.Stack[string]{}
	s = strings.ReplaceAll(s, " ", "")

	for i := 0; i < len(s); i++ {
		char := s[i]
		if isDigit(char) {
			number := ""
			for i < len(s) && isDigit(s[i]) {
				number += string(s[i])
				i++
			}
			st.Push(number)
			i--
			continue
		}
		op := string(char)
		if op != rparenthesis {
			st.Push(op)
			continue
		}

		// calculate all inside parentheses
		operation := make([]string, 0, 3)
		for !st.IsEmpty() {
			v := st.Pop()
			if v == lparenthesis {
				break
			}
			operation = append(operation, v)
		}
		st.Push(calcPlain(operation))
	}

	// calculate all that left in a stack
	operation := make([]string, 0, st.Size())
	for !st.IsEmpty() {
		operation = append(operation, st.Pop())
	}

	res, err := strconv.Atoi(calcPlain(operation))
	if err != nil {
		panic(err)
	}
	return res
}

func calcPlain(operation []string) string {
	fmt.Println("calc plain -", operation)
	j := len(operation) - 1
	sum := operation[j]
	if operation[j] == minus {
		num := operation[j-1]
		if num[0] == '-' {
			sum = num[1:]
		} else {
			sum += num
		}

		j--
	}
	for j = j - 1; j > 0; j = j - 2 {
		operator := operation[j]
		rval := operation[j-1]
		sum = calculateOperation(sum, rval, operator)
	}

	return sum
}

func calculateOperation(l, r, operator string) string {
	fmt.Printf("calculating: %s %s %s\n", l, operator, r)
	lInt, err := strconv.Atoi(l)
	if err != nil {
		panic(err)
	}
	rInt, err := strconv.Atoi(r)
	if err != nil {
		panic(err)
	}

	result := 0
	switch operator {
	case plus:
		result = lInt + rInt
	case minus:
		result = lInt - rInt
	}

	return strconv.Itoa(result)
}

func isDigit(v byte) bool {
	if v >= '0' && v <= '9' {
		return true
	}

	return false
}
