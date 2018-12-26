package concurrency

import (
	"testing"
)

func TestChanOverChan(t *testing.T) {
	result := AckChannels()
	t.Logf("Did %v work", result)
}
