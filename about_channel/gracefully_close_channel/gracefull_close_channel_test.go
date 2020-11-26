package gracefully_close_channel

import "testing"

// go test -v . -test.run TestOneSenderNReceviers
func TestOneSenderNReceivers(t *testing.T) {
	OneSenderNReceivers()
}

// go test -v . -test.run TestNSendersOneReceiver
func TestNSendersOneReceiver(t *testing.T) {
	NSendersOneReceiver()
}
