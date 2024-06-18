package main

import (
	"fmt"
	"time"
	"os/exec"
)

func say(s string) {
	for j := 0; j < 3; j++{
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func printTime() {
	// os exec
	dateCmd := exec.Command("date")
	dateOut, err := dateCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> date")
    fmt.Println(string(dateOut))
}

func main() {
	// goroutines
	go say("Async")
	say("Sync")
	
	printTime()
}
