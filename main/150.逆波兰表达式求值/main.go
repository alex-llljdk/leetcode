package main

import (
	"strconv"

	"github.com/emirpasic/gods/stacks/arraystack"
)

// 逆波兰计算，通常使用栈实现
func evalRPN(tokens []string) int {
	st := arraystack.New() //arraystack 官方包
	for _, val := range tokens {
		if len(val) > 1 || val[0] >= '0' && val[0] <= '9' {
			num, _ := strconv.Atoi(val)
			st.Push(num)
		} else {
			y := popInt(st)
			x := popInt(st)
			switch val {
			case "+":
				st.Push(x + y)
			case "-":
				st.Push(x - y)
			case "*":
				st.Push(x * y)
			default:
				st.Push(x / y)
			}
		}
	}
	ans := popInt(st)
	return ans
}

func popInt(stack *arraystack.Stack) int {
	v, _ := stack.Pop()
	return v.(int) //接口断言用法
}

func main() {

}
