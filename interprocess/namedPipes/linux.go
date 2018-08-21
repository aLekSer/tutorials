package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"syscall"
	"time"
)

const pipeFile string = "tmpPipe"

func main() {
	//to create pipe: does not work in windows
	os.Remove(pipeFile)
	err3 := syscall.Mkfifo(pipeFile, 0666)
	if err3 != nil {
		fmt.Println(err3)
	}

	go scheduleWrite()

	file2, err := os.OpenFile(pipeFile, os.O_CREATE, os.ModeNamedPipe)
	if err != nil {
		log.Fatal("Open named pipe file error:", err)
	}
	//to open pipe to read
	reader := bufio.NewReader(file2)

	defer os.Remove(pipeFile)
	for {
		line, err := reader.ReadBytes('\n')
		if err == nil {
			fmt.Print("load string:" + string(line))
		}
	}
	/*
		file2, err := os.OpenFile("tmpPipe", os.O_RDONLY, os.ModeNamedPipe)
		var b []byte
		n, err := file2.Read(b)
		if n > 0 && err != nil {
			fmt.Println(b)
		} else {
			fmt.Println(err)
		}
	*/
}

func scheduleWrite() {
	fmt.Println("start schedule writing.")
	f, err := os.OpenFile(pipeFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	i := 0
	for {
		fmt.Println("write string to named pipe file.")
		f.WriteString(fmt.Sprintf("test write times:%d\n", i))
		i++
		time.Sleep(time.Second)
	}
}
