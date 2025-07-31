package main

import (
	"flag"
	"time"
	"fmt"
	"os"
	"encoding/csv"
)

var filepath string = "tasks.csv"
var allTasks [][]string

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
	} else{
		fmt.Println("The correct usage is: go run . $argument/flag $task/nothing")
	}
}




func wholeList() {

}




func list() {
	for _, i := range allTasks {//i dont think this is how you do enumeration in go lang
		fmt.Printf("%s\t%s\t%s\n",i[0],i[1],i[2])
	}
}




func add(task string) {
	currentTime := time.Now()
	update := []string {task, currentTime.String(), "false"}

	allTasks = append(allTasks, update)
	writer()
}




func delete(delTask string) {
	newlength := 0
	for _, task := range allTasks {
		if task[0] != delTask {
			allTasks[newlength] = task 
			newlength++
		}
	}
	allTasks = allTasks[:newlength]


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
	}

	reader := csv.NewReader(file)
	tempTasks, err := reader.ReadAll()
	file.Close()
	if err != nil {
		panic(err)
	}

	allTasks = tempTasks
}



func writer() {
	file, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	
	for _, task := range allTasks{
		err = writer.Write(task)
	}

	if err != nil {
		panic(err)
	}

	writer.Flush()
}
