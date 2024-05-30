package main

import (
	"flag"
	"fmt"
	"os"
	"todo-cli/todo"
)

const todofile = "todo.json"

func main() {
	list := &todo.List{}
	if err := list.Load(todofile); err != nil {
		fmt.Fprintf(os.Stderr, "Error loading to-do list: %v\n", err)
		os.Exit(1)
	}
	add := flag.String("add", "", "Task to add to the to-do list")
	remove := flag.Int("remove", 0, "ID of the task to remove from the to-do list")
	complete := flag.Int("complete", 0, "ID of the task to mark as completed")
	listFlag := flag.Bool("list", false, "List all tasks")
	flag.Parse()

	switch {
	case *add != "":
		list.Add(*add)
		if err := list.Save(todofile); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving to-do list: %v\n", err)
			os.Exit(1)
		}
	case *remove > 0:
		list.Remove(*remove)
		if err := list.Save(todofile); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving to-do list: %v\n", err)
			os.Exit(1)
		}
	case *complete > 0:
		list.Completed(*complete)
		if err := list.Save(todofile); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving to-do list: %v\n", err)
			os.Exit(1)
		}
	case *listFlag:
		list.List()
	default:
		fmt.Println("Invalid command. Use -add, -remove, -complete, or -list.")
		os.Exit(1)
	}
}

/*go build -o todo-cli
./todo-cli -add "Buy milk"
./todo-cli -add "Learn Go"
./todo-cli -list
./todo-cli -complete 1
./todo-cli -list
./todo-cli -remove 1
./todo-cli -list
*/
