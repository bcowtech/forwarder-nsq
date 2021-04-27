package nsq

import "github.com/nsqio/go-nsq"

type (
	Producer = nsq.Producer
	Config   = nsq.Config
)

func NewConfig() *Config {
	return nsq.NewConfig()
}

type Forwarder struct {
	producer *Producer
}

func NewForwarder(option *Option) *Forwarder {
	p, err := nsq.NewProducer(option.NsqAddr, option.Config)
	if err != nil {
		panic(err)
	}

	return &Forwarder{
		producer: p,
	}
}

func (f *Forwarder) Write(topic string, body []byte) error {
	p := f.producer

	return p.Publish(topic, body)
}

func (f *Forwarder) Close() {
	p := f.producer

	defer p.Stop()
}
