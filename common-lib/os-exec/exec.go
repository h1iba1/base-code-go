package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main()  {

	cmd := exec.Command("ls", "-la")
	cmd.Stdout = os.Stdout

	b :=bufio.NewReader(os.Stdout)
	err :=cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	f, _ :=os.Create("./text.txt")

	b.WriteTo(f)

}