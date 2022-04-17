package main

import (
	"encoding/json"
	"os"
)

type Omg struct {
	ID   int
	Name string
	Age  int
}

func main() {
	grp := Omg{
		ID:   100,
		Name: "Nik",
		Age:  100,
	}
	var1, _ := json.Marshal(grp)
	os.Stdout.Write(var1)
}
