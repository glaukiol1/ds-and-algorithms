package main

import "fmt"

var people []string

func main() {
	addPerson("kevin")
	addPerson("devin")
	addPerson("gaga")
	printPeople()
}

func addPerson(name string) {
	people = append(people, name)
}

func printPeople() {
	for i, person := range people {
		fmt.Println(fmt.Sprint(i+1) + ": " + person)
	}
}
