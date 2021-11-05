package lab2

import (
	"fmt"
	"strings"
)

// PrefixToInfix converts string with an expression in prefix form to a string with an expression in infix (regular) form
func PrefixToInfix(input string) (string, error) {
	var stack []string
	const operators = "+-*/"

	if len(input) == 0 {
		return "", fmt.Errorf("Empty string")
	}

	// fmt.Println("input string: ", input)

	var inputs = strings.Split(input, " ")

	var literal = "" // Current character (number or operator)
	var e1 = ""      // first top value from stack
	var e2 = ""      // second

	// Go from right to left
	for i := len(inputs) - 1; i >= 0; i-- {
		literal = inputs[i]

		if !strings.Contains(operators, literal) {
			stack = append(stack, literal)
		} else {

			// Если в стеке не хватает элементов, значит опратор не в нужном месте (или лишний)
			if len(stack) < 2 {
				return "", fmt.Errorf("Wrong Expression: " + input)
			}

			e1, stack = Pop(stack)
			e2, stack = Pop(stack)

			// Поставить скобочки, если надо (если не надо, не ставить)
			if strings.Contains(e1, " ") && strings.Contains("*/", literal) {
				e1 = "(" + e1 + ")"
			}
			if strings.Contains(e2, " ") && strings.Contains("*/", literal) {
				e2 = "(" + e2 + ")"
			}

			var new = e1 + " " + literal + " " + e2
			stack = append(stack, new)
		}

		// fmt.Println(stack)
	}

	var res = ""
	res, stack = Pop(stack)

	if len(stack) > 0 {
		return "", fmt.Errorf("Wrong Expression: " + input)
	}

	return res, nil
}

func Pop(s []string) (string, []string) {
	if len(s) == 0 {
		return "", []string{}
	}

	return s[len(s)-1], s[:len(s)-1]
}
