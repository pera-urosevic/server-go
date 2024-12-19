package check

import (
	"testing"
)

func TestCheckRunning(t *testing.T) {
	running := Running("explorer.exe", "explorer")
	if !running {
		t.Errorf("Check running lowercase failed")
	}

	running = Running("explorer.exe", "EXPLORER")
	if !running {
		t.Errorf("Check running uppercase failed")
	}

	running = Running("NotReal.exe", "NotReal")
	if running {
		t.Errorf("Check running non existing failed")
	}
}
