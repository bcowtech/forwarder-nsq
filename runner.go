package nsq

type Runner struct {
	forwarder *Forwarder
}

func (r *Runner) Start() {
	r.forwarder.logger.Printf("[bcowtech/forwarder-nsq] Started\n")
}

func (r *Runner) Stop() {
	logger := r.forwarder.logger
	logger.Printf("[bcowtech/forwarder-nsq] Closing\n")
	r.forwarder.Close()
	logger.Printf("[bcowtech/forwarder-nsq] Closed\n")
}
