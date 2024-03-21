package lab2

import (
	"fmt"
	"strings"
)

// cheks if passed string is math operation
func isOperator(c string) bool {
	return strings.ContainsAny(c, "+-*/^")
}

// PrefixToPostfix converts
func PrefixToPostfix(input string) (string, error) {
	if input == "" {
		return "", fmt.Errorf("invalid expression")
	}

	stack := []string{}
	expression := strings.Split(input, " ")
	flag := false

	for i := len(expression) - 1; i >= 0; i-- {
		if isOperator(expression[i]) {
			if len(stack) < 2 {
				return "", fmt.Errorf("invalid expression")
			}
			operator1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			operator2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if len(operator1) == 1 && flag {
				operator1, operator2 = operator2, operator1
			}
			stack = append(stack, operator1+" "+operator2+" "+expression[i])
			flag = true
		} else {
			stack = append(stack, expression[i])
		}
	}

	if len(stack) != 1 {
		return "", fmt.Errorf("invalid expression")
	}

	return stack[len(stack)-1], nil
}
