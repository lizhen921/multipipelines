package multipipes

import (
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	startPipCase()
	//waitRoutine := sync.WaitGroup{}
	//waitRoutine.Add(1)
	//waitRoutine.Wait()
	time.Sleep(time.Second * 100)
}

func pip1(arg map[string]interface{}) map[string]interface{} {
	log.Println("P1 get : ", arg)
	arg["message"] = "I'm pip1 return " + arg["message"].(string)

	return arg
}

func pip2(arg map[string]interface{}) map[string]interface{} {
	log.Println("P2 get : ", arg)
	arg["message"] = "I'm pip2 return " + arg["message"].(string)
	//time.Sleep(time.Second * 10)

	return arg
}

func pip3(arg map[string]interface{}) map[string]interface{} {
	log.Println("P3 get : ", arg)
	//判断超时等待
	if arg["timeout"] == true {
		arg["message"] = "P3 is timeout"
		log.Println("P3 timeout, do some 执行超时动作，然后进入下次等待")
	}
	arg["message"] = "I'm pip3 return " + arg["message"].(string)
	return arg
}

func pip4(arg map[string]interface{}) map[string]interface{} {
	log.Println("P4 get : ", arg)
	if arg["timeout"] == true {
		//log.Println("P3 get : ", timeout)
	}

	arg["message"] = "I'm pip4 return " + arg["message"].(string)
	return arg
}
func createPip() (txPip Pipeline) {
	pipNodeSlice := make([]*Node, 0)
	pipNodeSlice = append(pipNodeSlice, &Node{Target: pip1, Name: "pip1", Capacity: 5})
	pipNodeSlice = append(pipNodeSlice, &Node{Target: pip2, Name: "pip2", RoutineNum: 2})
	pipNodeSlice = append(pipNodeSlice, &Node{Target: pip3, Name: "pip3", Timeout: 9})
	pipNodeSlice = append(pipNodeSlice, &Node{Target: pip4, Name: "pip4", RoutineNum: 2})
	txPip = Pipeline{
		Nodes: pipNodeSlice,
	}
	return txPip
}

func startPipCase() {
	txPip := createPip()
	indata := startProduceData()
	outData := startProcessData()
	txPip.Setup(indata, outData)
	//txPip.Setup(indata, nil)
	txPip.Start()

	//waitRoutine := sync.WaitGroup{}
	//waitRoutine.Add(1)
	//waitRoutine.Wait()
}

func produceData(node *Node) {
	//note you can init some datas before start produce
	for i := 0; i < 20; i++ {
		s := "produce data : " + strconv.Itoa(i)
		m := map[string]interface{}{"message": s}
		log.Println(s)
		node.Output <- m
		time.Sleep(time.Second)
	}
}
func startProduceData() *Node {
	preNode := &Node{
		Name: "just do nothing",
	}
	go produceData(preNode)
	return preNode
}

func processResult(node *Node) {
	for {
		s := <-node.Input
		log.Println("get final data : ", s["message"].(string))
		time.Sleep(time.Second)
	}
}
func startProcessData() *Node {
	afterNode := &Node{
		Name: "just do nothing",
	}
	go processResult(afterNode)
	return afterNode
}

func TestSlice(t *testing.T) {
	rcmResultList := []int{1, 2, 3, 4}
	fmt.Printf("%v\n", rcmResultList[0:3])
	position := 2
	rcmResultList = append(rcmResultList[:position+1], rcmResultList[position:]...)
	fmt.Printf("%v", rcmResultList)

}
