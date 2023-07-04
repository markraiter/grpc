package main

import (
	"encoding/json"
	"fmt"

	"github.com/markraiter/grpc/api/user"
	"google.golang.org/protobuf/proto"
)

func main() {
	user1 := &user.User{
		Name:    "Bob",
		Age:     21,
		IsAdmin: false,
	}

	dataJSON, _ := json.Marshal(user1)
	fmt.Printf("\n\ndataJSON len %d byte: \n%v\n", len(dataJSON), dataJSON)

	dataPB, _ := proto.Marshal(user1)
	fmt.Printf("dataPB len %d byte: \n%v\n", len(dataPB), dataPB)

}

type User struct {
	Name    string
	Age     int
	IsAdmin bool
}
