// map containing what we want
// stack containing the string
package validParenthesis

func validParenthesis(s string) bool {
	m := map[string]string{
		'(':')',
		'{':'}',
		'[':']',
	}
	stack := make([]rune, 0)

	for _, i := range s {
		if curri, exists := m[i]; exists {
			if len(stack) > 0 && stack[len(stack)-1] == curri {
				stack = stack[len(stack)-1]
			} else {
				stack = append(stack, curri)
			}
		}
	}
    return len(stack) == 0
}