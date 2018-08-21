package main

import (
	"fmt"
	"interprocess/server"
	"log"
	"net/rpc"
)

func main() {
	args := &server.Args{7, 8}

	serverAddress := "[127.0.0.1]"
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Asynchronous call
	quotient := new(server.Quotient)
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	replyCall := <-divCall.Done // will be equal to divCall
	fmt.Println(replyCall)
	fmt.Println(quotient)
	args2 := &server.Args{7, 2}
	divCall = client.Go("Arith.Divide", args2, quotient, nil)
	replyCall = <-divCall.Done // will be equal to divCall
	fmt.Println(quotient)
	v := int(1)
	divCall = client.Go("Arith.Multiply", args2, &v, nil)
	replyCall = <-divCall.Done // will be equal to divCall
	fmt.Println(v)
	args3 := &server.Args{7, 0}
	divCall = client.Go("Arith.Divide", args3, quotient, nil)
	replyCall = <-divCall.Done // will be equal to divCall
	//&{Arith.Divide 0xc4201200d0 0xc420010220 divide by zero 0xc420148000}
	fmt.Println(replyCall)
	// check errors, print, etc.

}
