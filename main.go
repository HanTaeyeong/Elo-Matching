package main

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	matchMaking "go-streaming/matchMaking"
	"io"
	"log"
	"net"
)

type FileServer struct {
}

func main() {

	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Println(sendFile(40000))
	// }()
	// server := &FileServer{}
	// server.start()
	matchMaking.InitMatchMaking()
}

func (fs *FileServer) start() {
	ln, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go fs.readLoop((conn))
	}
}

func (fs *FileServer) readLoop(conn net.Conn) {
	buf := new(bytes.Buffer)
	for {
		var size int64
		binary.Read(conn, binary.LittleEndian, &size)
		n, err := io.CopyN(buf, conn, size)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(buf.Bytes())
		fmt.Println("recieved %d bytes", n)
	}
}

func sendFile(size int) error {
	file := make([]byte, size)
	_, err := io.ReadFull(rand.Reader, file)
	if err != nil {
		return err
	}
	conn, err := net.Dial("tcp", ":5000")
	if err != nil {
		return err
	}
	binary.Write(conn, binary.LittleEndian, int64(size))
	n, err := io.CopyN(conn, bytes.NewReader(file), int64(size))
	if err != nil {
		return err
	}
	fmt.Printf("written%d bytes over network\n", n)
	return nil
}
