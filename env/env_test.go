package env

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	Test()
	port := os.Getenv("API_PORT")
	if port != "55557" {
		t.Errorf("load test env failed")
	}
}
