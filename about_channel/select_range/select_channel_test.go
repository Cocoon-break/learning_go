package select_range

import "testing"

// go test -v . -test.run TestNoBlockChannel
func TestNoBlockChannel(t *testing.T) {
	NoBlockChannel()
}

// go test -v . -test.run TestBlockChannel
func TestBlockChannel(t *testing.T) {
	BlockChannel()
}

// go test -v . -test.run TestSelectClosedChannel
func TestSelectClosedChannel(t *testing.T) {
	selectClosedChannel()
}

// go test -v . -test.run TestSelectNilChannel
func TestSelectNilChannel(t *testing.T) {
	selectNilChannel()
}
