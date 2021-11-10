package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

func main() {
	marshal()
	unMarshall()
}

func marshal() {
	reds := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	byteVal, err := json.Marshal(reds)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("\n%+v", reds)
	fmt.Println(reds.Colors)
	fmt.Println(reds.Colors[1])
	os.Stdout.Write(byteVal)
}

func unMarshall() {
	var jsonBlob = []byte(`
	{
		"ID":2, 
		"Name":"Whites",
		"Colors": ["Beige","Gray","Ivory"]
	}
		`)

	var whites ColorGroup
	err := json.Unmarshal(jsonBlob, &whites)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("\n%+v", whites)
	fmt.Println(whites.Colors)
	fmt.Println(whites.Colors[1])
}
