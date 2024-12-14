package utils

type Stack[T any] []T;

func StackOf[T any](elems... T) Stack[T] {
	stack := make(Stack[T], 0, len(elems))

	for _, elem := range elems {
		stack = stack.Push(elem)
	}

	return stack
}

func (stack Stack[T]) Push(elem T) Stack[T] {
	return append(stack, elem)
}

func (stack Stack[T]) Pop() (T, Stack[T]) {
	if len(stack) == 0 {
		panic("Popped empty stack")
	}

	return stack[0], stack[1:]
}

func (stack Stack[T]) Peek() (T, Stack[T]) {
	if len(stack) == 0 {
		panic("Popped empty stack")
	}

	return stack[0], stack
}

func (stack Stack[T]) IsEmpty() bool {
	return len(stack) == 0
}