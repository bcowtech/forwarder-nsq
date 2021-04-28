package nsq

type Runner struct {
	forwarder *Forwarder
}

func (r *Runner) Start() {
	r.forwarder.logger.Printf("[bcow-go/forwarder-nsq] Started\n")
}

func (r *Runner) Stop() {
	logger := r.forwarder.logger
	logger.Printf("[bcow-go/forwarder-nsq] Closing\n")
	r.forwarder.Close()
	logger.Printf("[bcow-go/forwarder-nsq] Closed\n")
}
