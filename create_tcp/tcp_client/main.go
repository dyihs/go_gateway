package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//   1. 连接服务器
	conn, err := net.Dial("tcp", "localhost:9090")
	defer conn.Close()
	if err != nil {
		fmt.Printf("connect failed, err: %v\n", err.Error())
		return
	}

	//   step2: 读取命令行输入
	inputReader := bufio.NewReader(os.Stdin)
	for {
		//   step3: 一直读取直到读到\n
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("read from consle failed, err: %v\n", err)
			break
		}
		//   step4: 读取Q时停止
		trimmedInput := strings.TrimSpace(input)
		if trimmedInput == "Q" {
			break
		}
		//   step5: 回复服务器信息
		_, err = conn.Write([]byte(trimmedInput))
		if err != nil {
			fmt.Printf("write failed, err: %v\n", err)
			break
		}
	}

}
