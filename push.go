// push.go
// By https://www.appgao.com
// Do not delete this lines

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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

// {
//   "platform": "ios",
//   "token": "token",
//   "notification": {
//     "title": "title",
//     "subtitle": "subtitle",
//     "body": "body",
//     "badge": 1
//   }
// }

var zabbixPush ZabbixPush

const zCateServer = "https://zcate.appgao.com/push"

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

	// fmt.Println(string(data))  // show json string

	req, err := http.NewRequest("POST", zCateServer, bytes.NewBuffer(data))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.StatusCode)
	// fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	if resp.StatusCode == 200 {
		os.Exit(0)
	}
	os.Exit(1)
}
