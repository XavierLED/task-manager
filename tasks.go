package todo

import (
	"flag"
	"time"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	wholeList := flag.Bool("a", false, "a flag that will list all tasks completed and uncompleted")

	if len(args) > 0 && len(args) <= 2 {
		switch {
		case args[0] == "list":
			fmt.Println(list())
		
		case args[0] == "add":
			add(args[1])
		
		case args[0] == "delete":
			delete(args[1])

		default:
			fmt.Println("argument is not recoginzed")
		}
	}

	else if *wholeList {
		fmt.Println(wholeList())
	}

	fmt.Println("The correct usage is: go run . $argument/flag $task/nothing")
}

func wholeList() {

}

func list() string{

}

func add(task string) {
	currentTime := time.Now()
}

func delete(task string) {

}

func complete(task string) {

}

func loader(filePath string) {

}

func writer(filePath string) {

}