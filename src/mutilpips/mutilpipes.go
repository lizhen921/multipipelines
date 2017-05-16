package pipelines

import (
	"errors"
	"log"
)

type Node struct {
	target     func(interface{}) interface{}
	input      chan interface{}
	output     chan interface{}
	routineNum int
	cache      int
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
		//logs.Info(n.name, ",in run forever")
		err := n.run()
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}

//execute the Node method,and save the result in to the channel
func (n *Node) run() error {
	x, ok := <-n.input
	if !ok {
		log.Fatal(errors.New("read data from inputchannel error"))
		return nil
	}
	//TODO  not good enough, how to support multi params and returns
	if n.output == nil {
		return nil
	}
	n.output <- n.target(x)
	return nil
}

type Pipeline struct {
	nodes []*Node
}

/*
setup pip: Combine all nodes
Args:
	indata (Node): the mothod produce data which will come in to the pipline
	outdata (Node): data processing method when the pipeline handler is finished
Returns:
*/
func (p *Pipeline) setup(indata *Node, outdata *Node) {
	inNode := []*Node{indata}
	nodes_all := append(inNode, p.nodes...)
	nodes_all = append(nodes_all, outdata)
	p.connect(nodes_all)
}

//connect all nodes's output and input.
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
	if head.cache == 0 {
		head.cache = 10
	}
	head.input = make(chan interface{}, head.cache)
	head.output = make(chan interface{}, head.cache)
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
