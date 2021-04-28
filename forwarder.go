package nsq

import (
	"log"
	"os"
	"time"

	"github.com/nsqio/go-nsq"
)

var (
	defaultLogger = log.New(os.Stdout, "[bcow-go/forwarder-nsq]", log.LstdFlags|log.Lmicroseconds|log.Llongfile|log.Lmsgprefix)
)

type (
	Producer = nsq.Producer
	Config   = nsq.Config
)

func NewConfig() *Config {
	return nsq.NewConfig()
}

type Forwarder struct {
	producer *Producer
	logger   *log.Logger
}

func NewForwarder(option *Option) *Forwarder {
	p, err := nsq.NewProducer(option.NsqAddr, option.Config)
	if err != nil {
		panic(err)
	}

	return &Forwarder{
		producer: p,
		logger:   defaultLogger,
	}
}

func (f *Forwarder) Write(topic string, body []byte) error {
	p := f.producer

	return p.Publish(topic, body)
}

func (f *Forwarder) DeferredWrite(topic string, delay time.Duration, body []byte) error {
	p := f.producer

	return p.DeferredPublish(topic, delay, body)
}

func (f *Forwarder) Close() {
	p := f.producer

	defer p.Stop()
}

func (f *Forwarder) Runner() *Runner {
	return &Runner{
		forwarder: f,
	}
}
