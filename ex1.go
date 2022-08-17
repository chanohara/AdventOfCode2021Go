package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input_ex1.txt")
	if err != nil {
		log.Fatal(err)
	}
	string := string(data)
	string2 := strings.Split(string, "\n")
	prev := 0
	count := 0
	for i, s := range string2 {
		currInt, _ := strconv.Atoi(s)
		if i > 0 {
			if prev < currInt {
				count += 1
			}
		}
		prev = currInt
	}
	fmt.Print(count)
}
