package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports , results chan int)  {
	for p:=range ports{
		address :=fmt.Sprintf("scanme.nmap.org:%d",p)
		conn ,err :=net.Dial("tcp",address)
		if err !=nil{
			results <-0
			continue
		}
		conn.Close()
		// 将开放端口发送给results通道
		results<-p
	}
}

func main()  {
	ports :=make(chan int,100)
	results :=make(chan int)
	var openports []int
	/*
	同时启动100个线程进行扫描
	*/
	for i:=0;i<cap(ports);i++{
		go worker(ports,results)
	}

	/*
	采用channel来传输端口
	所有goroutine同时从ports获取扫描资源。保证了在并发扫描时不会重复扫描同一个端口
	*/
	go func() {
		for i:=1;i<=1024;i++{
			ports <-i
		}
	}()

	for i:=0;i<1024;i++{
	//将接受的结果添加到openports切片
		port :=<-results
		if port!=0{
			openports=append(openports,port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	for _,port :=range openports{
		fmt.Printf("%d open\n",port)
	}

}