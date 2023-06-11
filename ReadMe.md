# Installation

-    for unix systems, please build locally after cloning it
     `git clone https://github.com/toothsy/cpi-ll`
-    for building in unix systems, please use
     `go build`
-    for windows you can directly run the binary `cpi-ll.exe` once cloned.

# CLL

-    there will be 26 circular linked lists in the event where all words of different alphabets are inserted.
-    each alphabet CLL is stored in a map for faster lookup.

-    nodes are always inserted in ascending order of their word size, this somewhat reduces the lookup time
-    Node structure is this
-    ```go
     type Node struct {
     	Word string
     	Next *Node
     }
     ```
-    cll Head structure is this
-    ```go
     type CLLNode struct {
      	Head *Node
      	Tail *Node
     }
     ```
-    Tail is used for easier insertion for longer words
-    words stored are case-insensitive, so `Hello` and `HeLLo` are same and wont be inserted
