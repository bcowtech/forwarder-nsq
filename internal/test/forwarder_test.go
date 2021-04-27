package test

import (
	"fmt"
	"os"
	"testing"
	"time"

	nsq "gitlab.bcowtech.de/bcow-go/forwarder-nsq"
)

var (
	NsqServers = os.Getenv("NSQ_SERVERS")
)

func TestForwarder(t *testing.T) {

	forwarder := nsq.NewForwarder(&nsq.Option{
		NsqAddr: NsqServers,
		Config:  nsq.NewConfig(),
	})

	defer forwarder.Close()

	for _, word := range []string{"Welcome", "to", "the", "nsqio", "go-nsq", "Golang", "client"} {
		forwarder.Write("myTopic", []byte(word))
		time.Sleep(100 * time.Millisecond)
		fmt.Println(word)
	}
}
