package main

import (
	"flag"
	"time"
	"fmt"
	"os"
	"text/tabwriter"
	"encoding/csv"
	"github.com/mergestat/timediff"
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
			case args[0] == "-a":
				wholeList()
			default:
				fmt.Println("The argument used is unknown\n Proper use: ./tasks $argument $task/nothing")
		}
	}else if (*listWhole) {
		wholeList()
	} else{
		fmt.Println("The correct usage is: go run . $argument/flag $task/nothing")
	}
}




func wholeList() {
		w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)

	fmt.Fprintln(w, "ID\tDescription\tTime\tCompletion")
	for j, i := range allTasks {//i dont think this is how you do enumeration in go lang
		date, _ := time.Parse(time.RFC1123, i[1])
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\n",j,i[0],timediff.TimeDiff(date),i[2])
		
	}
	w.Flush()
}




func list() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)

	fmt.Fprintln(w, "ID\tDescription\tTime\tCompletion")
	for j, i := range allTasks {//i dont think this is how you do enumeration in go lang
		if i[2] == "false" {
			date, _ := time.Parse(time.RFC1123, i[1])
			fmt.Fprintf(w, "%d\t%s\t%s\t%s\n",j,i[0],timediff.TimeDiff(date),i[2])
		
		}
	}
	w.Flush()
}




func add(task string) {
	currentTime := time.Now()
	update := []string {task, currentTime.Format(time.RFC1123), "false"}

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
			writer()
			return 0
		}
	}

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
