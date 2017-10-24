package testmultipips

import (
	"log"
	"multipipelines/multipipes"
	"strconv"
	"time"
)

type preNode struct {
	node multipipes.Node
	note string
}

func (p *preNode) produceData() {
	//note you can init some datas before start produce
	for i := 0; i < 20; i++ {
		s := "produce data : " + strconv.Itoa(i)
		log.Println(s)
		p.node.Output <- s
		time.Sleep(time.Second)
	}
}
func startProduceData() *multipipes.Node {
	pre := &preNode{
		note: "just do nothing",
	}
	go pre.produceData()
	return &pre.node
}

type afterNode struct {
	node   multipipes.Node
	messge string
}

func (a *afterNode) processResult() {
	for {
		s := <-a.node.Input
		log.Println("get final data : ", s)
		time.Sleep(time.Second)
	}
}
func startProcessData() *multipipes.Node {
	after := afterNode{
		messge: "just do nothing",
	}
	go after.processResult()
	return &after.node
}
