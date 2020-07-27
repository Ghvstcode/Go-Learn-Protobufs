package main

import (
	"fmt"

	"github.com/GhvstCode/ProtoBuf/src/simple"
)

func main() {
	doSimple()
}
 func doSimple() {
 	sm := example_simple.SimpleMessage{
		Id:         12345,
		IsSimple:   false,
		Name:       "This is a simple message!!",
		SampleList: []int32{1,8,6,4},
	}

	fmt.Print("This is sm :", sm.GetId())
 }