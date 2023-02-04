package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"time"

	viper "github.com/spf13/viper"
)

func FatalIfErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	b, err := ioutil.ReadFile("config.toml")
	if err != nil {
		log.Fatalln("setting.toml no found")
	}
	viper.SetConfigType("toml")
	viper.ReadConfig(bytes.NewBuffer(b))
	webhookUrl := viper.GetString("webhookUrl")
	webhook := Webhook{webhookUrl}
	mcsrvAddress := viper.GetString("mcsrvAddress")

	var lastStat ServerStatus
	lastAP := 0
	for range time.Tick(time.Minute * 15) {
		srvInfo, _ := getStatus(mcsrvAddress)
		if srvInfo.Status == Offline && lastStat == Online {
			webhook.sentMessage("OOF! the server apparently broken")
		} else {
			if srvInfo.Status == Online && lastStat == Offline {
				webhook.sentMessage("POG! the server is working again")
			}
			if srvInfo.CurrecntPlayer == lastAP {
				//pass
			} else if srvInfo.CurrecntPlayer > lastAP {
				webhook.sentMessage("Someone join the server")
			} else {
				webhook.sentMessage("Someone leave the server")
			}
		}
		log.Println(srvInfo)
		lastStat = srvInfo.Status
		lastAP = srvInfo.CurrecntPlayer
	}
}
