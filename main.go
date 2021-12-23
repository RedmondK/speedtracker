package main

import (
	"fmt"
	"speedtracker/speedtrackertypes"
)

func main() {  
    fmt.Println("Speed Tracker Library")
	var test speedtrackertypes.Session
	test.Protocol = "Level 1"
	fmt.Println(test.Protocol)
}