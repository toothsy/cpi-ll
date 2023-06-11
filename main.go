package main

import (
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
	fmt.Println(cllCollection)
}

func main() {
	cllCollection := cll.Init()
	if len(os.Args) < 2 {
		fmt.Println("Invalid arguments. Usage: go run main.go [option] [data]")
		fmt.Println("Options:")
		fmt.Println("  -insert [data0,data1,data2...]    Insert a node with the given data separated by comma")
		fmt.Println("  -search [data]    			    Search for a node with the given data")
		fmt.Println("  -delete [data]    			    Delete a node with the given data")
		fmt.Println("  -display          			    Display the circular linked list")
		return
	}

	option := os.Args[1]

	switch option {
	case "-insert":
		if len(os.Args) < 3 {
			fmt.Println("Missing data argument. Usage: go run main.go -insert [data0,data1,data2...] separated by comma")
			return
		}
		csdata := os.Args[2]
		for _, word := range strings.Split(csdata, ",") {
			word = cases.Title(language.English).String(word)
			createIndividualCLL(cllCollection, word)

		}
		fmt.Printf("Inserted node with following data '%s' into the circular linked list.\n", csdata)

	case "-search":
		if len(os.Args) < 3 {
			fmt.Println("Missing data argument. Usage: go run main.go -search [data]")
			return
		}
		data := os.Args[2]
		wordCLL := cllCollection[data[0]]
		found := wordCLL.Search(data)
		if found {
			fmt.Printf("Node with data '%s' found in the circular linked list.\n", data)
		} else {
			fmt.Printf("Node with data '%s' not found in the circular linked list.\n", data)
		}

	case "-delete":
		if len(os.Args) < 3 {
			fmt.Println("Missing data argument. Usage: go run main.go -delete [data]")
			return
		}
		data := os.Args[2]
		wordCLL := cllCollection[data[0]]
		deleted := wordCLL.Delete(data)
		if deleted {
			fmt.Printf("Node with data '%s' deleted from the circular linked list.\n", data)
		} else {
			fmt.Printf("Node with data '%s' not found in the circular linked list.\n", data)
		}

	case "-display":
		printCollection(cllCollection)
	case "-help":
		fmt.Println("Invalid option. Usage: go run main.go [option] [data]")
		fmt.Println("Options:")
		fmt.Println("  -insert [data0,data1,data2...]    Insert a node with the given data separated by comma")
		fmt.Println("  -search [data]    			    Search for a node with the given data")
		fmt.Println("  -delete [data]    			    Delete a node with the given data")
		fmt.Println("  -display          			    Display the circular linked list")

	default:
		fmt.Println("Invalid option. Usage: go run main.go [option] [data]")
		fmt.Println("Options:")
		fmt.Println("  -insert [data0,data1,data2...]    Insert a node with the given data")
		fmt.Println("  -search [data]    			    Search for a node with the given data")
		fmt.Println("  -delete [data]    			    Delete a node with the given data")
		fmt.Println("  -display          			    Display the circular linked list")

	}
	printCollection(cllCollection)

}

func printCollection(cllCollection map[byte]*cll.CLLNode) {
	fmt.Println("Circular linked list:")
	for key, wordCLL := range cllCollection {
		fmt.Println(string(key))
		wordCLL.Display()
	}
}
