package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/toothsy/cpi-ll/cll"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func createIndividualCLL(cllCollection map[byte]*cll.CLLNode, word string) {
	if _, ok := cllCollection[word[0]]; ok {
		cllCollection[word[0]].Insert(word)
		return
	}
	cllCollection[word[0]] = &cll.CLLNode{}
	cllCollection[word[0]].Insert(word)
}

func main() {
	cllCollection := cll.Init()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nEnter 'i' to insert, 's' to search, 'd' to delete, '-d' to display, 'q' to quit:")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "i":
			fmt.Println("Enter word(s) to insert separated by comma:")
			data, _ := reader.ReadString('\n')
			data = strings.TrimSpace(data)

			for _, word := range strings.Split(data, ",") {
				if len(word) == 0 {
					continue
				}
				capitalizedWord := cases.Title(language.English).String(word)
				createIndividualCLL(cllCollection, capitalizedWord)
			}
			fmt.Printf("Inserted node(s) with following data '%s' into the circular linked list.\n", data)

		case "s":
			fmt.Println("Enter word to search:")
			data, _ := reader.ReadString('\n')
			data = strings.TrimSpace(data)
			capitalizedWord := cases.Title(language.English).String(data)
			wordCLL := cllCollection[capitalizedWord[0]]
			if wordCLL == nil {
				fmt.Printf("Node with data '%s' not found in the circular linked list.\n", data)
				continue
			}
			found := wordCLL.Search(capitalizedWord)
			if found {
				fmt.Printf("Node with data '%s' found in the circular linked list.\n", data)
			} else {
				fmt.Printf("Node with data '%s' not found in the circular linked list.\n", data)
			}

		case "d":
			fmt.Println("Enter word to delete:")
			data, _ := reader.ReadString('\n')
			data = strings.TrimSpace(data)

			capitalizedWord := cases.Title(language.English).String(data)
			wordCLL := cllCollection[capitalizedWord[0]]
			if wordCLL == nil {
				fmt.Printf("Node with data '%s' not found in the circular linked list.\n", data)
				continue
			}
			deleted := wordCLL.Delete(capitalizedWord)
			if deleted {
				fmt.Printf("Node with data '%s' deleted from the circular linked list.\n", data)
			} else {
				fmt.Printf("Node with data '%s' not found in the circular linked list.\n", data)
			}
		case "-d":
			printCollection(cllCollection)

		case "q":
			fmt.Println("Exiting the program...")
			return

		default:
			fmt.Println("Invalid input. Please try again.")

		}

	}

}

func printCollection(cllCollection map[byte]*cll.CLLNode) {
	fmt.Println("Circular linked list:")
	for key, wordCLL := range cllCollection {
		fmt.Println(string(key))
		wordCLL.Display()
	}
	fmt.Println("")
}
