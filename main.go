package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"

	"github.com/GhvstCode/ProtoBuf/src/simple"
)

func main() {
	sm := doSimple()

	sm2 := &example_simple.SimpleMessage{}
	writeToFile("simple.bin", sm)
	readFromFile("simple.bin", sm2)
	fmt.Println(sm2)
}

func writeToFile(fname string, pb proto.Message) error{
	out, err := proto.Marshal(pb)
	if err !=  nil {
		log.Fatal("Cant serialize data to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatal("Cant Write to file", err)
		return err
	}
	return nil
}
func readFromFile(fname string, pb proto.Message) error{
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Print("Something went wrong when reading the file", err)
		return err
	}

	err = proto.Unmarshal(in, pb)
	if err != nil {
		fmt.Print("couldnt unmarshal Protobuff", err)
		return err
	}
	return nil
}
 func doSimple() *example_simple.SimpleMessage{
 	sm := example_simple.SimpleMessage{
		Id:         12345,
		IsSimple:   false,
		Name:       "This is a simple message!!",
		SampleList: []int32{1,8,6,4},
	}

	fmt.Print("This is sm :", sm.GetId())

 	return &sm
 }