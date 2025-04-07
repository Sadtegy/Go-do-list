package main

import "fmt"

func main() {
	var install = "install go language from the website"
	var setup = "setup go and code editor in your pc"
	var basics = "learn go basic fundamentals"
	var create = "create some small project by using go"
	var TaskLists = []string{install, setup, basics, create}

	printTask(TaskLists)
	fmt.Println("")
	addTask(TaskLists, "have a cup of coffee")
}

func printTask(TaskLists []string) {
	fmt.Println("Welcome to my GO TO LIST")
	fmt.Println("")
	fmt.Println("list of tasks")

	for index, task := range TaskLists {
		fmt.Printf("%d. %s \n", index+1, task)
	}
}

func addTask(TaskLists []string, newTask string) {
	var updatedTaskItem = append(TaskLists, newTask)
	printTask(updatedTaskItem)
}
