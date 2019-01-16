package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gogo/protobuf/proto"
)

// Message ...
type Message struct {
	X int64  `protobuf:"varint,1,opt,name=x" json:"x,omitempty"`
	Y string `protobuf:"bytes,2,opt,name=y" json:"y,omitempty"`
}

// Reset ...
func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }

// ProtoMessage ...
func (*Message) ProtoMessage() {}

func main() {

	loopTimes := 1000

	m := &Message{X: 1, Y: "2"}

	jsonBeginTime := time.Now()
	for i := 0; i < loopTimes; i++ {
		_, err := json.Marshal(m)
		if err != nil {
			panic(err)
		}
	}
	jsonEndTime := time.Now()

	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	gobBeginTime := time.Now()
	for i := 0; i < loopTimes; i++ {
		err := enc.Encode(m)
		if err != nil {
			panic(err)
		}
	}
	gobEndTime := time.Now()

	protoBeginTime := time.Now()
	for i := 0; i < loopTimes; i++ {
		_, err := proto.Marshal(m)
		if err != nil {
			panic(err)
		}
	}
	protoEndTime := time.Now()

	fmt.Println("    json:", jsonEndTime.UnixNano()-jsonBeginTime.UnixNano())
	fmt.Println("     gob:", gobEndTime.UnixNano()-gobBeginTime.UnixNano())
	fmt.Println("protobuf:", protoEndTime.UnixNano()-protoBeginTime.UnixNano())
}
