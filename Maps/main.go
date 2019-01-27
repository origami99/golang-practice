package main

import "fmt"

func main() {
	phonebook := map[string]int{
		"Nani": 1234,
		"Jo":   7777,
		"Bo":   3030,
	}

	printMap(phonebook)

	delete(phonebook, "Jo")

	fmt.Println("After deleting Jo:")
	printMap(phonebook)

}

func printMap(m map[string]int) {
	for key, value := range m {
		fmt.Printf("%s => %v\n\r", key, value)
	}
}
