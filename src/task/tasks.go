package main

import (
	"flag"
	"time"
	"fmt"
	"os"
	"encoding/csv"
)

var filepath string = "tasks.csv"
var allTasks := []string[]{}

func main() {
	loader()
	args := os.Args[1:]
	listWhole := flag.Bool("a", false, "a flag that will list all tasks completed and uncompleted")

	if len(args) > 0 && len(args) <= 2 {
		switch {
		case args[0] == "list":
			list()
		
		case args[0] == "add":
			add(args[1])
		
		case args[0] == "delete":
			delete(args[1])

		case args[0] == "complete":
			complete(args[1])
		default:
			fmt.Println("The argument used is unknown\n Proper use: ./tasks $argument $task/nothing")
		}
	} else if (*listWhole) {
		wholeList()
	}

	fmt.Println("The correct usage is: go run . $argument/flag $task/nothing")
}




func wholeList() {

}




func list() {
	for _, i := range allTasks {//i dont think this is how you do enumeration in go lang
		fmt.Println(i)
	}
}




func add(task string) {
	currentTime := time.Now()
	update := []string {task, currentTime.String(), "false"}

	allTasks = append(allTasks, update)
}




func delete(task string) {

	writer()
}




func complete(task string) int{
	//possibly split the string

	for i := 0; i < len(allTasks); i++ {
		if (allTasks[i][0] == task) {
			allTasks[i][2] = "true"
			return 0
		}
	}

	writer()

	fmt.Println("Task does not exist and therefore cannot be completed")
	return 1
}




func loader() {
	file, err := os.Open(filepath)

	if err != nil {
		panic(err)
		file.Close()
	}

	reader := csv.NewReader(file)
	tempTasks, err := reader.ReadAll()
	file.Close()
	if err != nil {
		panic(err)
	}

	allTasks = append(allTasks, tempTasks)
}



func writer() {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	
	err = writer.Write(update)
	if err != nil {
		panic(err)
	}

	writer.Flush()

	if err := writer.Error; err != nil {
		panic(err)
	}
}