package leetcode

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	if len(tokens) == 1 {
		i, _ := strconv.Atoi(tokens[0])
		return i
	}
	stack, top := []int{}, 0
	for _, v := range tokens {
		switch v {
		case "+":
			{
				sum := stack[top-2] + stack[top-1]
				stack = stack[:top-2]
				stack = append(stack, sum)
				top -= 1
			}
		case "-":
			{
				sub := stack[top-2] - stack[top-1]
				stack = stack[:top-2]
				stack = append(stack, sub)
				top -= 1
			}
		case "*":
			{
				mul := stack[top-2] * stack[top-1]
				stack = stack[:top-2]
				stack = append(stack, mul)
				top -= 1
			}
		case "/":
			{
				div := stack[top-2] / stack[top-1]
				stack = stack[:top-2]
				stack = append(stack, div)
				top -= 1
			}
		default:
			{
				i, _ := strconv.Atoi(v)
				stack = append(stack, i)
				top++
			}
		}
	}
	return stack[0]
}
