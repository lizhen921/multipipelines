package pipelines

import (
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	startPipCase()
	time.Sleep(time.Second * 100)
}
