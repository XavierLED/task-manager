package todo

import (
	"flag"
	"time"
	"fmt"
	"os"
	"encoding/csv"
)

filepath := "tasks.csv"

func main() {
	args := os.Args[1:]

	if len(args) > 0 && len(args) <= 2 {
		switch {
		case args[0] == "list":
			fmt.Println(list())
		
		case args[0] == "add":
			add(args[1])
		
		case args[0] == "delete":
			delete(args[1])

		case args[0] == "complete":
			complete(args[1])
		default:
			fmt.Println("The argument used is unknown\n Proper use: ./tasks $argument $task/nothing")
		}
	}

	fmt.Println("The correct usage is: go run . $argument/flag $task/nothing")
}

func list() string{
	allTasks := loader()

	for i, allTasks {//i dont think this is how you do enumeration in go lang
		fmt.Println(i)
	}
}

func add(task string) {
	update := {task, time.Now, "false"}

	file, err := os.OpenFile(filepath, os.O_WRONGLY|os.O_APPEND, 0664)
	if err != nil {
		panic(err)
	}

	writer := csv.NewWriter(file)
	writer.Write(update)
	writer.FLush()
}

func delete(task string) {

}

func complete(task string) int{
	allTasks := loader()
	//possibly split the string

	for i, allTasks {
		if allTasks[i][0] == task {
			allTasks[i][2] = true
			return 0
		}
	}

	fmt.Println("Task does not exist and therefore cannot be completed")
	return 1

}

func loader() string[][] {
	file, err := os.Open(filepath)

	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(file)
	allTasks := reader.ReadAll(file)

	return allTasks
}