package main

import (
	"exc9/mapred"
	"fmt"
	"log"
	"os"
	"strings"
)

// Main function
func main() {
	// todo read file
	data, err := os.ReadFile("solution/res/meditations.txt")
	if err != nil {
		log.Fatal(err)
	}

	// todo run your mapreduce algorithm
	lines := strings.Split(string(data), "\n")
	var mr mapred.MapReduce
	results := mr.Run(lines)

	// todo print your result to stdout
	for k, v := range results {
		fmt.Println("%s: %d\n", k, v)
	}
}
