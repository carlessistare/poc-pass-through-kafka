package main

import (
	"github.com/golang/protobuf/proto"
	"net/http"
	"fmt"
	"bytes"
	"ws-test/proto-files/example"
)

func main() {
	myClient := example.Test{Label: "hola"}

	data, err := proto.Marshal(&myClient)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = http.Post("http://localhost:8001/proto", "", bytes.NewBuffer(data))

	if err != nil {
		fmt.Println(err)
		return
	}
}
