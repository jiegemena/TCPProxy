package main

import (
	"fmt"
	"net"
	"os"

	"github.com/jiegemena/gotools/configs"
)

func Println(a ...interface{}) {
	fmt.Println(a...)
}

func tcpProxy(paddr, taddr string) {
	// proxyaddr := fmt.Sprintf(":%d", fromport)
	plistener, err := net.Listen("tcp", paddr)
	if err != nil {
		Println("监听: ", paddr, ", 错误: ", err.Error())
		os.Exit(1)
	}
	defer plistener.Close()

	for {
		pconn, err := plistener.Accept()
		if err != nil {
			Println("连接错误：", err.Error())
			continue
		}

		buffer := make([]byte, 1024)
		n, err := pconn.Read(buffer)
		if err != nil {
			Println("读取缓冲区失败", err.Error())
			continue
		}

		tconn, err := net.Dial("tcp", taddr)
		if err != nil {
			Println("连接错误", taddr, ", error: ", err.Error())
			pconn.Close()
			continue
		}

		n, err = tconn.Write(buffer[:n])
		if err != nil {
			Println("中转错误", err.Error())
			pconn.Close()
			tconn.Close()
			continue
		}

		go pRequestSwap(pconn, tconn)
		go pRequestSwap(tconn, pconn)
	}
}

func pRequestSwap(r net.Conn, w net.Conn) {
	defer r.Close()
	defer w.Close()

	var buffer = make([]byte, 4096000)
	for {
		n, err := r.Read(buffer)
		if err != nil {
			Println("读取失败", err.Error())
			break
		}

		n, err = w.Write(buffer[:n])
		if err != nil {
			Println("中转失败", err.Error())
			break
		}
	}
}

func main() {
	configs.Init("")

	bindtcp := configs.GetConfig("bindtcp").(string)
	totcp := configs.GetConfig("totcp").(string)
	Println("启动server", "监听：", bindtcp, "转发到:", totcp)
	tcpProxy(bindtcp, totcp)
}
