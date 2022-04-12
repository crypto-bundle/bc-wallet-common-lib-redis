package nats

// ConsumerDirective ....
type ConsumerDirective uint8

const (
	QueueStreamNameTag  = "queue_stream"
	QueueSubjectNameTag = "queue_subject"

	QueuePubAckStreamNameTag = "queue_pub_ack_stream"
	QueuePubAckSequenceTag   = "queue_pub_ack_sequence"

	WorkerUnitNumberTag = "worker_unit_num"
)

const (
	DirectiveForRejectName  = "rejected"
	DirectiveForPassName    = "passed"
	DirectiveForReQueueName = "requeue"
)

const (
	DirectiveForReject ConsumerDirective = iota
	DirectiveForPass
	DirectiveForReQueue
)

func (d ConsumerDirective) String() string {
	return [...]string{DirectiveForRejectName,
		DirectiveForPassName,
		DirectiveForReQueueName,
	}[d]
}
