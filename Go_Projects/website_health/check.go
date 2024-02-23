package main

import (
	"fmt"
	"net"
	"time"
)

func Check(destination string, port string) string {
	address := destination + ":" + port
	timeOut := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp", address, timeOut)

	var status string

	if err != nil {
		status = fmt.Sprintf("The domain %s is down,\n Error: %v", destination, err)
	} else {
		status = fmt.Sprintf("The domain %s is up\nLocal Address: %s\nRemote Address: %s", destination, conn.LocalAddr(), conn.RemoteAddr())
	}
	return status
}
