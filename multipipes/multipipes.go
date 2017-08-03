package multipipes

import (
	"errors"
	"log"
	"time"
)

type Node struct {
	Target     func(interface{}) interface{}
	Input      chan interface{}
	Output     chan interface{}
	RoutineNum int //the number of goroutine
	Capacity   int //channel capacity
	Name       string
	Timeout    int64
}

//Start the Node(goroutines) based on the routineNum
func (n *Node) start() {
	if n.RoutineNum == 0 {
		n.RoutineNum = 1
	}
	for i := 0; i < n.RoutineNum; i++ {
		go n.runForever()
	}
}

//each Node goroutine should run forver
func (n *Node) runForever() {
	for {
		err := n.run()
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}

//execute the Node method,and save the result in to the channel
func (n *Node) run() error {
	isTimeout := make(chan bool, 1)
	go func() {
		time.Sleep(time.Second * time.Duration(n.Timeout)) //等待
		if n.Timeout != 0 {
			isTimeout <- true
		}
	}()
	select {
	case x, ok := <-n.Input:
		//从ch中读到数据
		if !ok {
			log.Println(errors.New("read data from inputchannel error"))
			return nil
		}
		//TODO  not good enough, how to support multi params and returns
		out := n.Target(x)
		if n.Output == nil || out == nil {
			return nil
		}
		n.Output <- out
	case <-isTimeout:
		//一直没有从ch中读取到数据，但从timeout中读取到数据
		log.Println("read data timeout")
		return nil
	}
	return nil
}

type Pipeline struct {
	Nodes []*Node
}

//connect all nodes's output and input after .
/*
		indata			 node1			  node2			  outdata
	* * * * * * *	 * * * * * * *	  * * * * * * *	   * * * * * * *
	*	   out<-*----*-in 	out<-*----*-in	 out<-*----*-in		   *
	* * * * * * *	 * * * * * * *	  * * * * * * *	   * * * * * * *
*/
func (p *Pipeline) connect(nodes []*Node) (ch chan interface{}) {
	if len(nodes) == 0 {
		return nil
	}
	head := nodes[0]
	if head.Capacity == 0 {
		head.Capacity = 50
	}
	head.Input = make(chan interface{}, head.Capacity)
	head.Output = make(chan interface{}, head.Capacity)
	tail := nodes[1:]
	head.Output = p.connect(tail)
	return head.Input
}

/*
setup pip: Combine all nodes
actually the indata Node and outdata Node doesn't belong to the pipline, I just use their's output or input.
Args:
	indata (Node): the mothod produce data which will come in to the pipline
	outdata (Node): data processing method when the pipeline handler is finished
Returns:
*/
func (p *Pipeline) Setup(indata *Node, outdata *Node) {
	var nodesAll []*Node = p.Nodes
	if indata != nil {
		inNode := []*Node{indata}
		nodesAll = append(inNode, nodesAll...)
	}
	if outdata != nil {
		nodesAll = append(nodesAll, outdata)
	}
	p.connect(nodesAll)
}

// for..range start each Node
func (p *Pipeline) Start() {
	for index, _ := range p.Nodes {
		p.Nodes[index].start()
	}
}
