package check

import (
	"os"
	"server/env"
	"testing"
)

func TestCheckNet(t *testing.T) {
	env.Test()

	res1 := Net()
	if !res1 {
		t.Errorf("check net should success")
	}

	err := os.Setenv("MICRON_CHECK_NET_URL", "http://error.internal")
	if err != nil {
		t.Errorf("err: %v", err)
	}
	res2 := Net()
	if res2 {
		t.Errorf("check net should failed")
	}
}
