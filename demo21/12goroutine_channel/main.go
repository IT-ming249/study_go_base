package main

import (
	"sync"
)

var wg sync.WaitGroup

//向 intChan放入 1-120000个数

// 从 intChan取出数据，并判断是否为素数,如果是，就把得到的素数放在primeChan

//要关闭 primeChan
// close(primeChan) //如果一个channel关闭了就没法给这个channel发送数据了
//什么时候关闭primeChan?

//给exitChan里面放入一条数据

//printPrime打印素数的方法

func main() {

}
