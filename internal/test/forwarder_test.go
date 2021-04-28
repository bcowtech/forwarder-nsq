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

func TestForwarderWrite(t *testing.T) {

	forwarder := nsq.NewForwarder(&nsq.Option{
		NsqAddr: NsqServers,
		Config:  nsq.NewConfig(),
	})

	defer forwarder.Close()

	for _, word := range []string{`{
		"round": "2101250000021350358",
		"provider": "10547",
		"player": "test03",
		"totalBet": 10.00000,
		"totalWin": 5.00000,
		"totalRefund": 0,
		"startAt": 1611539874,
		"finishAt": 1611539877,
		"game": "100101101",
		"account": "10218261",
		"accountType": "1",
		"dealer": "PP",
		"dealerPlayer": "",
		"currency": "CNY",
		"wallet": "CNY",
		"betCount": 10,
		"casino": "",
		"room": "2",
		"roundType": "NORMAL",
		"playerOs": "Windows",
		"playerDevice": "other",
		"timezone": "UTC"
	}`} {
		forwarder.Write("myTopic", []byte(word))
		time.Sleep(100 * time.Millisecond)
		fmt.Println(word)
	}
}

func TestForwarderDeferredWrite(t *testing.T) {

	forwarder := nsq.NewForwarder(&nsq.Option{
		NsqAddr: NsqServers,
		Config:  nsq.NewConfig(),
	})

	defer forwarder.Close()

	for _, word := range []string{`{
		"round": "2101250000021350358",
		"provider": "10547",
		"player": "test03",
		"totalBet": 10.00000,
		"totalWin": 5.00000,
		"totalRefund": 0,
		"startAt": 1611539874,
		"finishAt": 1611539877,
		"game": "100101101",
		"account": "10218261",
		"accountType": "1",
		"dealer": "PP",
		"dealerPlayer": "",
		"currency": "CNY",
		"wallet": "CNY",
		"betCount": 10,
		"casino": "",
		"room": "2",
		"roundType": "NORMAL",
		"playerOs": "Windows",
		"playerDevice": "other",
		"timezone": "UTC"
	}`} {
		forwarder.DeferredWrite("myTopic", time.Second*10, []byte(word))
		fmt.Println(word)
	}
}
