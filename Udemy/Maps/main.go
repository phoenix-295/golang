package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	// m := make(map[string][]string)
	// str1 := []string{"Jon", "Doe"}
	// str2 := "Doe"
	// m["p1"] = str1

	map1 := make(map[string][]string)
	str1 := []string{"Jon", "Doe", "Captain", "America"}
	str2 := "Doe"
	map1["p1"] = str1

	idx := slices.Index(map1["p1"], str2)
	map1["p1"] = slices.Delete(map1["p1"], idx, idx+1)

	fmt.Println(map1)
}
