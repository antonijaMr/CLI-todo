package main

import (
	// todo "../../todo"
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	todo "toDo.com/todos"
)

const (
	todoFile = ".todos.json"
)

func main() {
	add := flag.Bool("add", false, "add a new todo")
	complete := flag.Int("complete", 0, "mark todo as completed")
	//complete has index of task that is completed
	del := flag.Int("del", 0, "delete todo")
	list := flag.Bool("list", false, "list all todos")
	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:

		task, err := getInput(os.Stdin, flag.Args()...)

		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
		todos.Add(task)

		err = todos.Store(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}

	case *complete > 0:
		err := todos.Complete(*complete)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
		err = todos.Store(todoFile)

		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	case *del > 0:
		err := todos.Delete(*del)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}

		err = todos.Store(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	case *list:
		todos.Print()
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(0)
	}
}

func getInput(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return "", err
	}

	text := scanner.Text()

	if len(text) == 0 {
		return "", errors.New("empty todo is not allowed")
	}

	return text, nil

}
