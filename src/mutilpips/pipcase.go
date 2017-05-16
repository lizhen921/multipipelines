/*

 */
package pipelines

import (
	"log"
	"sync"
	"time"
)

func pip1(arg interface{}) interface{} {
	log.Println("P1 get : ", arg)
	s := "I'm pip1 return"
	log.Println("P1 return:", s)
	return s
}

func pip2(arg interface{}) interface{} {
	log.Println("P2 get : ", arg)
	s := "I'm pip2 return"
	time.Sleep(time.Second * 2)
	log.Println("P2 return:", s)
	return s
}

func pip3(arg interface{}) interface{} {
	log.Println("P3 get : ", arg)
	s := "I'm pip3 return"
	log.Println("P3 return:", s)
	return s
}

func createPip() (txPip Pipeline) {
	pipNodeSlice := make([]*Node, 0)
	pipNodeSlice = append(pipNodeSlice, &Node{target: pip1, routineNum: 1, name: "pip1"})
	pipNodeSlice = append(pipNodeSlice, &Node{target: pip2, routineNum: 1, name: "pip2"})
	pipNodeSlice = append(pipNodeSlice, &Node{target: pip3, routineNum: 1, name: "pip3"})
	txPip = Pipeline{
		nodes: pipNodeSlice,
	}
	return txPip
}

func startPipCase() {
	txPip := createPip()
	//changefeed := getChangefeed()
	txPip.setup(&changefeed.node)
	txPip.start()

	waitRoutine := sync.WaitGroup{}
	waitRoutine.Add(1)
	waitRoutine.Wait()
}
