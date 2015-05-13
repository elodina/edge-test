package sarama

import (
	"testing"
)

var (
	produceRequestEmpty = []byte{
		0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00}

	produceRequestHeader = []byte{
		0x01, 0x23,
		0x00, 0x00, 0x04, 0x44,
		0x00, 0x00, 0x00, 0x00}

	produceRequestOneMessage = []byte{
		0x01, 0x23,
		0x00, 0x00, 0x04, 0x44,
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x05, 't', 'o', 'p', 'i', 'c',
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0xAD,
		0x00, 0x00, 0x00, 0x1C,
		// messageSet
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x10,
		// message
		0x23, 0x96, 0x4a, 0xf7, // CRC
		0x00,
		0x00,
		0xFF, 0xFF, 0xFF, 0xFF,
		0x00, 0x00, 0x00, 0x02, 0x00, 0xEE}
)

func TestProduceRequest(t *testing.T) {
	request := new(ProduceRequest)
	testEncodable(t, "empty", request, produceRequestEmpty)

	request.RequiredAcks = 0x123
	request.Timeout = 0x444
	testEncodable(t, "header", request, produceRequestHeader)

	request.AddMessage("topic", 0xAD, &Message{Codec: CompressionNone, Key: nil, Value: []byte{0x00, 0xEE}})
	testEncodable(t, "one message", request, produceRequestOneMessage)
}
