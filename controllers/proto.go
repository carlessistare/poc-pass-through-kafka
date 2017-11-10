package controllers

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"github.com/golang/protobuf/proto"
	"ws-test/proto-files/example"
	"fmt"
	"github.com/linkedin/goavro"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
)

type Request struct {
	Label	string	`json:"label"`
}

func SendToKafka(obj *example.Test) {
	fmt.Print("coucoucoucou")

	codec, err := goavro.NewCodec(`
        {
			"type": "record",
			"name": "Testlabel",
			"fields" : [
				{
				"name": "label",
				"type": "string"
				}
	 		 ]
        }`)

	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(obj)
	json.Unmarshal(inrec, &inInterface)

	fmt.Print("coucoucoucou")

	binary, err := codec.BinaryFromNative(nil, inInterface)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(binary)

	fmt.Print(kafka.ErrBadMsg)

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}
	p.Close()

	deliveryChan := make(chan kafka.Event)
	topic := "test"
	err = p.Produce(&kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny}, Value: binary}, deliveryChan)
	close(deliveryChan)

}

func Proto(c *gin.Context) {

	buf, err := ioutil.ReadAll(c.Request.Body)
	fmt.Print(buf)

	if err != nil {
		c.Err()
		return
	}

	obj := example.Test{}
	if err = proto.Unmarshal(buf, &obj); err != nil {

	}

	SendToKafka(&obj)

	c.String(200, "OK!!")
}
