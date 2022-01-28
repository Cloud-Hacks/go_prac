package main

import "fmt"

type Stack struct {
	sk   [10]string
	temp []string
}

var i int

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.sk) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	// fmt.Println(str, i)
	if i == 0 {
		s.sk[i] = str
	}
	s.sk[i] = str // Simply append the new value to the end of the stack
	i = i + 1
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := i // Get the index of the top most element.
		i = i - 1
		element := (s.sk)[index-1] // Index into the slice and obtain the element.
		s.temp = (s.sk)[:index]    // Remove it from the stack by slicing it off.
		return element, true
	}
}

func main() {
	var stack Stack // create a stack variable of type Stack

	stack.Push("this")
	stack.Push("is")
	stack.Push("sparta!!")
	z := 3

	fmt.Println(len(stack.temp))
	for z > 0 {
		x, y := stack.Pop()
		if y == true {
			fmt.Println(x)
		}
		z = z - 1
	}
}
