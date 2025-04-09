package main

import (
	"fmt"
	"github.com/learning-go-book-2e/proto_generate/data"
	"google.golang.org/protobuf/proto"
)

//go:generate protoc -I=. --go_out=. --go_opt=module=github.com/learning-go-book-2e/proto_generate --go_opt=Mperson.proto=github.com/learning-go-book-2e/proto_generate/data person.proto
func main() {
	p := &data.Person{
		Name:  "Bob Bobson",
		Id:    20,
		Email: "bob@bobson.com",
	}
	fmt.Println(p)
	protoBytes, _ := proto.Marshal(p)
	fmt.Println(protoBytes)
	var p2 data.Person
	proto.Unmarshal(protoBytes, &p2)
	fmt.Println(&p2)
}
