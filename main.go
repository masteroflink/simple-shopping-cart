package main

import (
	"fmt"
)

func listItems(items []string) {
	for index, item := range items {
		fmt.Print(index+1, ".", item, "\n")
	}
}

func main() {
	items := []string{}

	for {
		fmt.Println("Enter command (help for a list of commands):")
		var command string
		fmt.Scanln(&command)
		switch command {
		case "add":
			fmt.Println("Enter item to add:")
			var item string
			fmt.Scanln(&item)
			items = append(items, item)
			fmt.Println(item, "added Successfully!")
		case "remove":
			fmt.Println("Enter item number to remove:")
			listItems(items)
			var selection int
			fmt.Scanln(&selection)
			selection--
			deleted_item := items[selection]
			items = append(items[:selection], items[selection+1:]...)
			fmt.Println("Successfully Deleted", deleted_item)
		case "list":
			listItems(items)
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println(command, " is an invalid command.")
		}
	}
}
