package log

import "server/system"

func Log(args ...any) {
	system.Log("CHARGED", args...)
}
