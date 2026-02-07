package main

import (
	"fmt"
)

func main() {
	var percent, part, whole float64
	fmt.Println("find a whole figure using a percentage and it's part.")
	fmt.Print("part: ")
	fmt.Scan(&part)
	fmt.Print("percent: ")
	fmt.Scan(&percent)
	whole = (part*100)/percent
	fmt.Printf("%v is %v percent of %v\n", part, percent, whole)
}

