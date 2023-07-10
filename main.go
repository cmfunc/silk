package main

import (
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	// 处理连接逻辑
	log.Printf("New connection from: %s\n", conn.RemoteAddr().String())

	// 循环读取客户端发送的数据
	for {
		data := make([]byte, 1024)
		_, err := conn.Read(data)
		if err != nil {
			log.Printf("Error reading data: %s\n", err.Error())
			break
		}

		// 处理接收到的数据
		// 这里可以根据需要编写相应的逻辑
		log.Printf("Received: %s",string(data))

		// 回复客户端
		response := []byte("Received: " + string(data))
		_, err = conn.Write(response)
		if err != nil {
			log.Printf("Error sending response: %s\n", err.Error())
			break
		}
	}

	// 关闭连接
	conn.Close()
	log.Printf("Connection closed: %s\n", conn.RemoteAddr().String())
}

func main() {
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Failed to create listener:", err)
	}

	defer listen.Close()

	log.Println("Server started. Listening on localhost:8080")

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal("Failed to accept connection:", err)
		}

		go handleConnection(conn)
	}
}
