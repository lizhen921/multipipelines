package multipipes

import (
	"sync"
	"testing"
)

func TestStart(t *testing.T) {
	startPipCase()
	waitRoutine := sync.WaitGroup{}
	waitRoutine.Add(1)
	waitRoutine.Wait()
	//time.Sleep(time.Second * 100)
}
