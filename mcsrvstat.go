package main

import (
	"fmt"
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

func (e ServerStatus) String() string {
	switch e {
	case Online:
		return "Online"
	case Offline:
		return "Offline"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

type ServerInfo struct {
	Status         ServerStatus
	Description    string
	CurrecntPlayer int
	MaxPlayer      int
}

func getStatus(address string) (ServerInfo, error) {
	info := ServerInfo{}

	con, err := net.Dial("tcp", address)
	if err != nil {
		log.Println(err)
		info.Status = Offline
		return info, nil
	}

	_, err = con.Write([]byte("\xfe"))
	if err != nil {
		return info, err
	}

	reply := make([]byte, 1024)
	_, err = con.Read(reply)
	if err != nil {
		return info, err
	}

	con.Close()

	s := string(reply)
	s = strings.Replace(s, "\x00", "", -1)
	data := strings.Split(s, "\xa7")

	info.Status = Online
	info.Description = data[0]
	info.CurrecntPlayer, _ = strconv.Atoi(data[1])
	info.MaxPlayer, _ = strconv.Atoi(data[2])

	return info, nil
}
