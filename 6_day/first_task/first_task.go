package main

import "fmt"

// Node представляет один элемент (узел) стека
type Node struct {
	data int
	next *Node // Указатель на следующий узел, который лежит ниже
}

// Stack управляет всей стопкой. Он хранит только указатель на самую верхнюю тарелку.
type Stack struct {
	top *Node
}

// Push добавляет элемент на вершину стека
func (s *Stack) Push(value int) {
	// ТВОЙ КОД ЗДЕСЬ
	// 1. Создай новый узел Node, используя литерал структуры и взятие адреса ( &Node{...} )
	// 2. Сделай так, чтобы его поле next указывало на текущую вершину s.top
	// 3. Обнови s.top, сделав новый узел текущей вершиной
	next := s.top
	new := Node{value, next}
	s.top = &new
}

// Pop удаляет элемент с вершины стека и возвращает его значение.
// Вторым аргументом возвращает bool (флаг ok), показывающий, не пустой ли стек.
func (s *Stack) Pop() (int, bool) {
	// ТВОЙ КОД ЗДЕСЬ
	// 1. Проверь, не пустой ли стек (если s.top == nil, значит элементов нет). Верни (0, false)
	// 2. Сохрани значение из текущей вершины (s.top.data) в отдельную переменную
	// 3. Сдвинь s.top на следующий элемент (s.top = s.top.next)
	// 4. Верни сохраненное значение и true

	top := s.top
	if top == nil {
		return 0, false
	}
	value := (*top).data
	s.top = (*top).next
	return value, true
}

func main() {
	myStack := &Stack{} // Создаем пустой стек

	myStack.Push(10)
	myStack.Push(20)
	myStack.Push(30)

	// Мы ожидаем, что элементы будут извлекаться в обратном порядке: 30, 20, 10
	val, ok := myStack.Pop()
	fmt.Printf("Забрали: %d, Успешно: %t\n", val, ok) // Должно быть: 30, true

	val, ok = myStack.Pop()
	fmt.Printf("Забрали: %d, Успешно: %t\n", val, ok) // Должно быть: 20, true

	val, ok = myStack.Pop()
	fmt.Printf("Забрали: %d, Успешно: %t\n", val, ok) // Должно быть: 10, true

	// Проверяем извлечение из пустого стека
	val, ok = myStack.Pop()
	fmt.Printf("Забрали из пустого: %d, Успешно: %t\n", val, ok) // Должно быть: 0, false
}
