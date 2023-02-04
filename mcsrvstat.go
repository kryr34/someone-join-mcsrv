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
	Status         ServerStatus
	Description    string
	CurrecntPlayer int
	MaxPlayer      int
}

func getStatus(address string) (ServerInfo, error) {
	con, err := net.Dial("tcp", address)
	if err != nil {
		log.Println(err)
		info := ServerInfo{}
		info.Status = Offline
		return info, nil
	}

	_, err = con.Write([]byte("\xfe"))
	if err != nil {
		return ServerInfo{}, err
	}

	reply := make([]byte, 1024)
	_, err = con.Read(reply)
	if err != nil {
		return ServerInfo{}, err
	}
	con.Close()

	s := string(reply)
	s = strings.Replace(s, "\x00", "", -1)
	data := strings.Split(s, "\xa7")

	cur, _ := strconv.Atoi(data[1])
	max, _ := strconv.Atoi(data[2])

	return ServerInfo{
		Online,
		data[0],
		cur, max,
	}, nil
}
