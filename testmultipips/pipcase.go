package testmultipips

import (
	"log"
	"multipipelines/multipipes"
	"time"
)

func pip1(arg interface{}) interface{} {
	log.Println("P1 get : ", arg)
	s := "I'm pip1 return " + arg.(string)
	return s
}

func pip2(arg interface{}) interface{} {
	log.Println("P2 get : ", arg)
	s := "I'm pip2 return " + arg.(string)
	time.Sleep(time.Second * 10)
	return s
}

func pip3(arg interface{}) interface{} {
	log.Println("P3 get : ", arg)
	if arg == "timeout" {
		arg = "P3 is timeout"
		log.Println("P3 timeout, do some ")
	}

	s := "I'm pip3 return " + arg.(string)
	return s
}

func pip4(arg interface{}) interface{} {
	log.Println("P4 get : ", arg)
	if arg == "timeout" {
		//log.Println("P3 get : ", timeout)
	}

	s := "I'm pip4 return " + arg.(string)
	return s
}
func createPip() (txPip multipipes.Pipeline) {
	pipNodeSlice := make([]*multipipes.Node, 0)
	pipNodeSlice = append(pipNodeSlice, &multipipes.Node{Target: pip1, Name: "pip1", Capacity: 5})
	pipNodeSlice = append(pipNodeSlice, &multipipes.Node{Target: pip2, Name: "pip2", RoutineNum: 2})
	pipNodeSlice = append(pipNodeSlice, &multipipes.Node{Target: pip3, Name: "pip3", Timeout: 1})
	pipNodeSlice = append(pipNodeSlice, &multipipes.Node{Target: pip4, Name: "pip4", RoutineNum: 2})
	txPip = multipipes.Pipeline{
		Nodes: pipNodeSlice,
	}
	return txPip
}

func startPipCase() {
	txPip := createPip()
	indata := startProduceData()
	outData := startProcessData()
	txPip.Setup(indata, outData)
	//txPip.setup(indata, nil)
	txPip.Start()

	//waitRoutine := sync.WaitGroup{}
	//waitRoutine.Add(1)
	//waitRoutine.Wait()
}
