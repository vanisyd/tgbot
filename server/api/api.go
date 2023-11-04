package api

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"strings"
)

func Init() {
	go RunService()
}

func RunService() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%s", TCPHost, TCPPort))
	if err != nil {
		log.Fatal(err)
	}

	defer func(listen net.Listener) {
		err := listen.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(listen)
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleServiceRequest(conn)
	}
}

func handleServiceRequest(conn net.Conn) {
	var command string
	var arguments []string
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		log.Println("[TCP.Error] " + err.Error())
		return
	}
	cmdSize := bytes.IndexByte(buffer[:], 0)
	cmdText := strings.TrimSpace(string(buffer[:cmdSize]))

	if strings.Contains(cmdText, "#") {
		cmdArguments := strings.Split(cmdText, "#")[1]
		arguments = strings.Split(cmdArguments, ",")
		command = strings.Split(cmdText, "#")[0]
	} else {
		command = cmdText
		arguments = []string{}
	}

	for _, cmd := range Commands {
		if cmd.Command == command {
			cmd.Handler(arguments)
			err := conn.Close()
			if err != nil {
				log.Println("[TCP.Error] " + err.Error())
				return
			}
		}
	}
}
