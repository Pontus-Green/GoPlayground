package main

import "fmt"

type Todo struct {
	Text string
	Done bool
}

func (t Todo) PrintTodo() string {
	return fmt.Sprintf("Todo: %s, Completed: %t", t.Text, t.Done)
}

func main() {
	todo := Todo{
		Text: "Learn Go",
		Done: false,
	}
	fmt.Println(todo.PrintTodo())
}
