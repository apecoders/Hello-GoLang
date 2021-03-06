package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned map of
	// messages to the console.
	fmt.Println(messages)

	email := "test@test.com"

	result := greetings.IsEmailValid(email)
	fmt.Println(result)

	number := 1
	printNumbers(number)

}

func printNumbers(num int) {
	if num <= 100 {
		fmt.Print(num, "\t")
		printNumbers(num + 1)
	}
}
