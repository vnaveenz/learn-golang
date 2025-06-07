package main

import "fmt"

func main() {
	m := map[string]string{
		"one": "1",
		"two": "2",
		"three": "3",
	}
	fmt.Println(m)
	// other way of declaring map
	var newmap map[string]string // this creates nil map, this is not initialized, so you won't be able to update
	// newmap["movie"] = "thug life"
	fmt.Println(newmap) // above line will fail
	// third way of declaring map
	m2 := make(map[string]string)
	m2["movie 1"] = "thug life"
	m2["movie 2"] = "aalavandhan"
	fmt.Println(m2) // above line wouldn't fail
	delete(m2, "movie")
	fmt.Println(m2)
	printMap(m2)
}

func printMap(m map[string]string) {
	for movie, moviename := range m {
		fmt.Println("The", movie, "is",  moviename, ", that we will be watching")
	}
}
