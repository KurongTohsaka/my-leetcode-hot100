package main

func main() {

}

func isValid(s string) bool {
	stack := make([]rune, 0)
	pair := map[rune]rune{'(': ')', '{': '}', '[': ']'}

	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
		}
		if ch == ')' || ch == ']' || ch == '}' {
			if len(stack) == 0 {
				return false
			}
			right := pair[stack[len(stack)-1]]
			if right != ch {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	return len(stack) == 0
}
