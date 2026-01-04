package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Todo struct {
	Text string
	Done bool
}

func loadTodos(fileName string) ([]Todo, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []Todo{}, nil
		}
		return nil, err
	}

	var todos []Todo
	err = json.Unmarshal(data, &todos)
	return todos, err
}

func saveTodos(fileName string, todos []Todo) error {
	data, err := json.MarshalIndent(todos, "", "	")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

func printTodos(todos []Todo) {
	fmt.Println("Listing tasks...")
	for i, t := range todos {
		status := ""
		if t.Done {
			status = "x"
		}

		fmt.Printf("%d: [%s] %s\n", i+1, status, t.Text)
	}
}

func main() {
	fileName := "todos.json"
	if len(os.Args) < 2 {
		fmt.Println("expcted 'add', 'list', 'check' or 'delete' subcommand")
		os.Exit(1)
	}

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	checkCmd := flag.NewFlagSet("check", flag.ExitOnError)
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])

		if addCmd.NArg() < 1 {
			fmt.Println("usage: todo add \"task text\"")
			os.Exit(1)
		}

		task := strings.Join(addCmd.Args(), " ")

		todos, err := loadTodos(fileName)
		if err != nil {
			panic(err)
		}

		todos = append(todos, Todo{Text: task, Done: false})

		err = saveTodos(fileName, todos)
		if err != nil {
			panic(err)
		}

	case "list":
		listCmd.Parse(os.Args[2:])

		todos, err := loadTodos(fileName)
		if err != nil {
			panic(err)
		}
		printTodos(todos)

	case "delete":
		deleteCmd.Parse(os.Args[2:])
		deleteArgs := deleteCmd.Args()
		if len(deleteArgs) < 1 {
			fmt.Println("usage: todo delete 1,2 ..")
			os.Exit(1)
		}

		// Parse and validate indices
		deleteArgFields := strings.Split(strings.Join(deleteArgs, ""), ",")
		var indices []int
		for _, v := range deleteArgFields {
			idx, err := strconv.Atoi(strings.TrimSpace(v))
			if err != nil {
				fmt.Printf("Invalid arg: %v\n", v)
				os.Exit(1)
			}
			indices = append(indices, idx-1) // Convert to 0-based
		}

		todos, err := loadTodos(fileName)
		if err != nil {
			panic(err)
		}

		// Validate all indices before deleting
		for _, idx := range indices {
			if idx < 0 || idx >= len(todos) {
				fmt.Printf("Not a valid delete number (%d-%d are valid)\n", 1, len(todos))
				os.Exit(1)
			}
		}

		// Sort indices in descending order to avoid shifting
		slices.Sort(indices)
		slices.Reverse(indices)

		// Delete todos at specified indices
		for _, idx := range indices {
			todos = slices.Delete(todos, idx, idx+1)
		}

		err = saveTodos(fileName, todos)
		if err != nil {
			panic(err)
		}
		fmt.Println("Updated todo list:")
		printTodos(todos)

	case "check":
		checkCmd.Parse(os.Args[2:])
		taskArg := checkCmd.Args()
		if len(taskArg) < 1 {
			fmt.Println("usage: todo check 1,2 ..")
			os.Exit(1)
		}

		taskArgFields := strings.Split(strings.Join(taskArg, ""), ",")

		todos, err := loadTodos(fileName)
		if err != nil {
			panic(err)
		}
		for _, v := range taskArgFields {
			taskNbr, err := strconv.Atoi(v)
			if err != nil {
				fmt.Printf("Invalid arg: %v \n", v)
			}

			if taskNbr-1 < 0 || taskNbr-1 > len(todos) {
				fmt.Printf("Not a valid task number (%d-%d are valid) \n", 1, len(todos))
				os.Exit(1)
			}

			checkTask := &todos[taskNbr-1]
			checkTask.Done = !checkTask.Done
		}

		saveTodos(fileName, todos)
		fmt.Println("Updated todo list:")
		printTodos(todos)

	default:
		fmt.Println("expected 'add' or 'list' subcommand")
		os.Exit(1)
	}
}
