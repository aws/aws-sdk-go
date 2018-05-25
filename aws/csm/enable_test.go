package csm

import (
	"testing"
)

func TestStartStop(t *testing.T) {
	Start("clientID", ":0")
	if sender == nil {
		t.Errorf("expected sender to be a non-nil value")
	}

	Start("clientID", ":0")
	if sender == nil {
		t.Errorf("expected sender to be a non-nil value")
	}

	Stop()

	if !sender.metricsCh.IsPaused() {
		t.Errorf("expected metrics channel to be paused but wasn't")
	}
}
