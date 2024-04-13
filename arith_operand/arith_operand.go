package main

import (
	"fmt"
)

// 给定一个字符串描述的算术表达式，计算出结果值。
//
// 输入字符串长度不超过 100 ，合法的字符包括 ”+, -, *, /, (, )” ， ”0-9” 。
// 1+2*3
// 1+3*(5/2-2)+2*5
// 参考：https://writings.sh/post/arithmetic-expression
type stack []interface{}

func (s *stack) push(v interface{}) {
	*s = append(*s, v)
}

func (s *stack) pop() interface{} {
	l := len(*s)
	v := (*s)[l-1]
	*s = (*s)[:l-1]
	return v
}
func (s *stack) top() interface{} {
	return (*s)[len(*s)-1]
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func Evaluate(input string) int {
	num_stack := new(stack)
	op_stack := new(stack)
	//op_stack.push('(')
	i := 0
	for i < len(input) {
		c := input[i]
		i++
		if isOperator(c) {
			for !op_stack.isEmpty() && isOperator(c) && precedence(c) <= precedence(op_stack.top().(byte)) {
				cal(op_stack, num_stack)
			}
			op_stack.push(c)
		} else if c == '(' {
			op_stack.push(c)
		} else if c == ')' {
			for !op_stack.isEmpty() && op_stack.top().(byte) != '(' {
				cal(op_stack, num_stack)
			}
			op_stack.pop()
		} else if isDigit(c) {
			v := int(c - '0')
			for i < len(input) && isDigit(input[i]) {
				v = v*10 + int(input[i]-'0')
				i++
			}
			num_stack.push(v)
		}
	}
	for !op_stack.isEmpty() {
		cal(op_stack, num_stack)
	}
	return num_stack.pop().(int)
}

func isOperator(c byte) bool {
	return c == '+' || c == '-' || c == '*' || c == '/'
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func precedence(c byte) int {
	switch c {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	default:
		return 0
	}
}

func cal(op_stack *stack, num_stack *stack) {
	op := op_stack.pop().(byte)
	num2 := num_stack.pop().(int)
	num1 := num_stack.pop().(int)
	switch op {
	case '+':
		num_stack.push(num1 + num2)
	case '-':
		num_stack.push(num1 - num2)
	case '*':
		num_stack.push(num1 * num2)
	case '/':
		num_stack.push(num1 / num2)
	}
}

func main() {
	var input string
	fmt.Printf("请输入算术表达式：")
	_, err := fmt.Scan(&input)
	if err != nil {
		println("输入错误")
		return
	}
	val := Evaluate(input)
	fmt.Printf("计算结果：%d\n", val)
}
