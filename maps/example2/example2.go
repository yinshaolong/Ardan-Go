package main

import "fmt"

type player struct {
	name  string
	score int
}

func main() {
	players := map[string]player{
		"anna": {"anna", 42},
		"john": {"john", 20},
	}
	//to access the attributes of the class players, must assign to a variable
	//to make accessible in memory
	player := players["anna"]
	player.score++
	fmt.Printf("Name: %v\nScore: %d", player.name, player.score)
}
