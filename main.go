package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	viper "github.com/spf13/viper"
)

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

	interval := viper.GetDuration("interval")

	var lastStat ServerStatus
	var lastAP int
	for range time.Tick(time.Second * interval) {
		srvInfo, _ := getStatus(mcsrvAddress)
		if srvInfo.Status == Offline && lastStat == Online {
			err = webhook.sentMessage("OOF! the server apparently broken")
		} else {
			if srvInfo.Status == Online && lastStat == Offline {
				err = webhook.sentMessage("POG! the server is working again")
			}
			if srvInfo.CurrecntPlayer == lastAP {
				//pass
			} else if srvInfo.CurrecntPlayer > lastAP {
				err = webhook.sentMessage("Someone join the server")
			} else {
				err = webhook.sentMessage("Someone leave the server")
			}
		}
		log.Printf("Webhook error: %v\n", err)
		log.Println(fmt.Sprintf("%+v\n", srvInfo))
		lastStat = srvInfo.Status
		lastAP = srvInfo.CurrecntPlayer
	}
}
