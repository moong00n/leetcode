package main

func solution(s string) string {
    stack := make([]byte, 0, len(s))

    for i := 0; i < len(s); i++ {
        if s[i] != '*' {
            stack = append(stack, s[i])
        } else if len(stack) > 0 {
            stack = stack[:len(stack)-1]
        }
    }

    return string(stack)
}
