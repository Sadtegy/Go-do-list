package main

import "fmt"

func main() {
	var install = "Install go language from the website"
	var setup = "Setup go and code editor in your pc"
	var basics = "Learn go basic fundamentals"
	var create = "Create some small project by using go"
	var taskLists = []string{install, setup, basics, create}

	taskLists = addTask(taskLists, "Have a cup of coffee")
	fmt.Println("")

	printTask(taskLists)
}

func printTask(taskLists []string) {
	fmt.Println("Welcome to the GO TO LIST")
	fmt.Println("")
	fmt.Println("List of Todos")

	for index, task := range taskLists {
		fmt.Printf("%d. %s \n", index, task)
	}
}

func addTask(taskLists []string, newList string) []string {
	var updatedTaskItem = append(taskLists, newList)
	return updatedTaskItem
}
