package main

import "fmt"

func main() {
	var instal = "Install Go and set up your workspace"
	var basics = "learn Go basics variables, data type, function"
	var practice = "Practice control flow: if, switch, loop"
	var understand = "Understand arrays, slice and map"
	var cli = "make small ClI project like to-do-list"
	var todoIteam = []string{instal, basics, practice, understand, cli}

	printTask(todoIteam)
}

func printTask(todoIteam []string) {

	fmt.Println("Welcome to the Go-Do-List")
	fmt.Println("")

	fmt.Println("Things you need to do for lean go")

	for index, task := range todoIteam {
		fmt.Printf("%d. %s \n", index+1, task)
	}

}
