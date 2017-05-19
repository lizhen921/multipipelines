/*

 */
package pipelines

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
	pipNodeSlice = append(pipNodeSlice, &Node{target: pip1, routineNum: 1, name: "pip1"})
	pipNodeSlice = append(pipNodeSlice, &Node{target: pip2, routineNum: 2, name: "pip2"})
	pipNodeSlice = append(pipNodeSlice, &Node{target: pip3, routineNum: 1, name: "pip3"})
	txPip = Pipeline{
		nodes: pipNodeSlice,
	}
	return txPip
}

func startPipCase() {
	txPip := createPip()
	indata := startProduceData()
	outData := startProcessData()
	txPip.setup(indata, outData) //if you don't neet the `indata`&`outdata`,just set them `nil`
	//txPip.setup(indata, nil)
	txPip.start()

	//waitRoutine := sync.WaitGroup{}
	//waitRoutine.Add(1)
	//waitRoutine.Wait()
}
