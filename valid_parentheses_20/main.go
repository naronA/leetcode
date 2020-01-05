package main

import "fmt"

func isValid(s string) bool {
	stack := make([]rune, 0)

	for _, c := range s {
		switch c {
		case '(':
			stack = append(stack, ')')
		case '{':
			stack = append(stack, '}')
		case '[':
			stack = append(stack, ']')
		default:
			if len(stack) != 0 && c == stack[len(stack)-1] {
				stack = stack[:len(stack)-1] // 末尾削除
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}

func main() {
	r := isValid("[")
	fmt.Println(r)
}
