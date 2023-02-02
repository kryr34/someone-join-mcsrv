package main

import (
	"log"
	"net"
	"strconv"
	"strings"
)

type ServerStatus int8

const (
	Online ServerStatus = iota
	Offline
)

type ServerInfo struct {
	description    string
	currecntPlayer int
	maxPlayer      int
}

func getStatus(address string) (ServerInfo, ServerStatus) {
	con, err := net.Dial("tcp", address)
	if err != nil {
		log.Println(err)
		return ServerInfo{}, Offline
	}

	_, err = con.Write([]byte("\xfe"))
	FatalIfErr(err)

	reply := make([]byte, 1024)
	_, err = con.Read(reply)
	FatalIfErr(err)

	s := string(reply)
	s = strings.Replace(s, "\x00", "", -1)
	data := strings.Split(s, "\xa7")

	cur, _ := strconv.Atoi(data[1])
	max, _ := strconv.Atoi(data[2])
	return ServerInfo{data[0], cur, max}, Online
}
