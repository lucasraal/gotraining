package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main()  {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()
	
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		io.WriteString(conn, "\nHello from heaven\n")
		fmt.Fprintln(conn, "I'm not on Earth anymore, but I'm fine here")
		fmt.Fprintf(conn, "%v", "I hope you're doing great!")

		conn.Close()
	}
}
// use telnet localhost 8080 to interact with the server