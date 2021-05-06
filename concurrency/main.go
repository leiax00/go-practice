package main

import (
	"concurrency/task"
	"fmt"
	"os"
)

func main() {
	err := task.HttpStart()
	if err != nil {
		fmt.Printf("Error Happen: %+v\n", err)
		os.Exit(0)
	}
	fmt.Println("Success to stop serve!!!")
	os.Exit(1)
}
