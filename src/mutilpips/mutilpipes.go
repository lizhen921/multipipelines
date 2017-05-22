package pipelines

import (
	"errors"
	"log"
)

type Node struct {
	target     func(interface{}) interface{}
	input      chan interface{}
	output     chan interface{}
	routineNum int //the number of goroutine
	capacity   int //channel capacity
	name       string
}

//Start the Node(goroutines) based on the routineNum
func (n *Node) start() {
	for i := 0; i < n.routineNum; i++ {
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
	//TODO add timeout
	//timeout := make(chan bool, 1)
	//go func() {
	//	time.Sleep(time.Second * 1) //等待1秒钟
	//	timeout <- true
	//}()
	//
	////然后，我们把timeout这个channel利用起来
	//select {
	//	case <-ch:
	//	//从ch中读到数据
	//	case <-timeout:
	//	//一直没有从ch中读取到数据，但从timeout中读取到数据
	//}
	x, ok := <-n.input
	if !ok {
		log.Fatal(errors.New("read data from inputchannel error"))
		return nil
	}
	//TODO not good enough, how to support multi params and returns
	out := n.target(x)
	if n.output == nil {
		return nil
	}
	n.output <- out
	return nil
}

type Pipeline struct {
	nodes []*Node
}

/*
setup pip: Combine all nodes
actually the indata Node and outdata Node doesn't belong to the pipline, I just use their's output or input.
Args:
	indata (Node): the mothod produce data which will come in to the pipline
	outdata (Node): data processing method when the pipeline handler is finished
Returns:
*/
func (p *Pipeline) setup(indata *Node, outdata *Node) {
	var nodesAll []*Node = p.nodes
	if indata != nil {
		inNode := []*Node{indata}
		nodesAll = append(inNode, nodesAll...)
	}
	if outdata != nil {
		nodesAll = append(nodesAll, outdata)
	}
	p.connect(nodesAll)
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
	if head.capacity == 0 {
		head.capacity = 50
	}
	head.input = make(chan interface{}, head.capacity)
	head.output = make(chan interface{}, head.capacity)
	tail := nodes[1:]
	head.output = p.connect(tail)
	return head.input
}

// for..range start each Node
func (p *Pipeline) start() {
	for index, _ := range p.nodes {
		p.nodes[index].start()
	}
}
