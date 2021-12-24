nc在实际攻防中经常使用，这里通过go来实现一个简单的netcat demo

### exec
接收到新的连接后，可以使用os/exec中的函数Command（name string,arg...string)创建新的cmd实列。该函数使用操作系统命令及其任何选项作为参数。在此实例中，将/bin/bash硬编码为命令并将-i作为参数，以使我们处于交互模式，这样就可以更可靠的操作stdin和stdout。
```go
cmd ：=exec.Commond("/bin/sh","-i")
```

### io.Pipe()
io.Pipe()会同时创建同步连接的一个reader和一个write----任何被写入write(wp)的数据都会被reader(rp)读取。因此需要将write分配给cmd.Stdout，然后使用io.Copy(conn,rp)将PipeReader连接到TCP连接。可使用goroutine防止代码被阻塞。命令的任何标准输出都将发送到write,然后通过冠道传送到reader并通过TCP链接输出。
```go
func handle(conn net.Conn)  {
	cmd :=exec.Command("/bin/sh","-i")

	rp,wp :=io.Pipe()
	cmd.Stdin=conn
	cmd.Stdout=wp
	go io.Copy(conn,rp)
	cmd.Run()
	conn.Close()
}
```

### io.copy 
io.Copy将PipeReader链接到TCP连接

### go os/exec 简明教程
https://colobu.com/2020/12/27/go-with-os-exec/