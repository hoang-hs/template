package message_queue

type AbstractMessageQueue interface {
	Topic() *string
	Payload() ([]byte, error)
}
