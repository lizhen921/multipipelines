/*

 */
package multipipes

import (
	"log"
	"time"
)

func pip1(arg interface{}) interface{} {
	log.Print("P1 get : ", arg)
	s := "I'm pip1 return " + arg.(string)
	return s
}

func pip2(arg interface{}) interface{} {
	log.Println("P2 get : ", arg)
	s := "I'm pip2 return " + arg.(string)
	time.Sleep(time.Second * 3)
	return s
}

func pip3(arg interface{}) interface{} {
	log.Println("P3 get : ", arg)
	s := "I'm pip3 return " + arg.(string)
	return s
}

func createPip() (txPip Pipeline) {
	pipNodeSlice := make([]*Node, 0)
	pipNodeSlice = append(pipNodeSlice, &Node{Target: pip1, Name: "pip1", Capacity: 5})
	pipNodeSlice = append(pipNodeSlice, &Node{Target: pip2, Name: "pip2", RoutineNum: 2})
	pipNodeSlice = append(pipNodeSlice, &Node{Target: pip3, Name: "pip3", Timeout: 10})
	txPip = Pipeline{
		Nodes: pipNodeSlice,
	}
	return txPip
}

func startPipCase() {
	txPip := createPip()
	indata := startProduceData()
	outData := startProcessData()
	txPip.Setup(indata, outData) //if you don't neet the `indata`&`outdata`,just set them `nil`
	//txPip.setup(indata, nil)
	txPip.Start()

	//waitRoutine := sync.WaitGroup{}
	//waitRoutine.Add(1)
	//waitRoutine.Wait()
}
