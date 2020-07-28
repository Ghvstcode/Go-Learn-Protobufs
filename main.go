package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"

	complex_pb_go "github.com/GhvstCode/ProtoBuf/src/complex"
	enumpb "github.com/GhvstCode/ProtoBuf/src/enum_example"
	"github.com/GhvstCode/ProtoBuf/src/simple"
)

func main() {
	//sm := doSimple()

	//sm2 := &example_simple.SimpleMessage{}
	//sm3 := &example_simple.SimpleMessage{}
	//writeToFile("simple.bin", sm)
	//readFromFile("simple.bin", sm2)
	//fmt.Println(sm2)
	//
	//jsonstring := toJson(sm)
	//fromJson(jsonstring, sm3)
	//fmt.Println(jsonstring)
	//fmt.Println("Succesfully created proto struct :", sm3)
	doEnum()
	doComplex()
}
func doComplex(){
	cm := complex_pb_go.ComplexMessage{
		OneDummy:     &complex_pb_go.DummyMessage{
			Id:   1,
			Name: "firstMessage",
		},
		MultipleDummy: []*complex_pb_go.DummyMessage{
			&complex_pb_go.DummyMessage{
				Id:   2,
				Name: "secondMessage",
			},
			&complex_pb_go.DummyMessage{
				Id:   3,
				Name: "thirdMessage",
			},
		},
	}
	fmt.Println(cm)
}

func doEnum(){
	em := enumpb.EnumMessage{
		Id: 42,
		DayOfTheWeek: enumpb.DayOfTheWeek_FRIDAY,
	}
	fmt.Println(em)
}
func toJson(pb proto.Message) string{
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Unable to marshalthe stuff", err)
		return ""
	}

	return out
}

func fromJson(in string, pb proto.Message) error{
	if err := jsonpb.UnmarshalString(in, pb); err != nil {
		log.Fatal("Cant serialize data to bytes", err)
		return err
	}
	return nil
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