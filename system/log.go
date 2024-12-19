package system

import (
	"fmt"
	"strings"
	"time"
)

func TimeFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func Log(module string, args ...any) {
	timestamp := fmt.Sprintf("[%s] [%s]", TimeFormat(time.Now()), strings.ToUpper(module))
	line := append([]any{timestamp}, args...)
	fmt.Println(line...)
}
