package todo

import (
	"flag"
	"time"
	"fmt"
	"os"
)

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

		default:
			fmt.Println("argument is not recoginzed")
		}
	}

	fmt.Println("The correct usage is: go run . $argument/flag $task/nothing")
}

func list() string{

}

func add(task string) {

}

func delete(task string) {

}

func complete(task string) {

}

func loader(filePath string) {

}

func writer(filePath string) {

}