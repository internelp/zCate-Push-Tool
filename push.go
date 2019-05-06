// push.go
// By https://www.appgao.com
// Do not delete this line

package main

import (
	"encoding/json"
	"flag"
	"fmt"
)

type ZabbixPush struct {
	Platform     string `json:"platform"`
	Token        string `json:"token"`
	Notification struct {
		Title    string `json:"title"`
		Subtitle string `json:"subtitle"`
		Body     string `json:"body"`
		Badge    int    `json:"badge"`
	} `json:"notification"`
}

var zabbixPush ZabbixPush

func init() {
	flag.StringVar(&zabbixPush.Platform, "platform", "", "iOS or Android (require)")
	flag.StringVar(&zabbixPush.Token, "token", "", "Your token (require)")
	flag.StringVar(&zabbixPush.Notification.Title, "title", "", "Message title")
	flag.StringVar(&zabbixPush.Notification.Subtitle, "subtitle", "", "Message subtitle")
	flag.StringVar(&zabbixPush.Notification.Body, "body", "", "Message body (require)")
	flag.IntVar(&zabbixPush.Notification.Badge, "badge", 1, "Message notification badge")
	flag.Parse()
}

func main() {
	data, _ := json.MarshalIndent(zabbixPush, "", "  ")
	fmt.Println(string(data))
}
