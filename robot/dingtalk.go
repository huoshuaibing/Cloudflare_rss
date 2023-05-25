package robot

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

//Init 初始化
func (d *DingRobot) Init(Url string) bool {
	if len(Url) == 0 {
		panic("url should not be empty")
		return false
	}
	d.Url = Url

	d.client = new(http.Client)
	// todo : add http time out
	d.client.Timeout = 50 * time.Second
	d.Data = make(chan Text, 100)
	return true
}

// Run 运行dingtalk
func (d *DingRobot) Run() {
	for newMsg := range d.Data {
		msg := DTText{}
		msg.MsgType = "markdown"
		msg.Markdown.Title = newMsg.Title
		msg.Markdown.Text = newMsg.Text
		msg.At.AtMobiles = []string{"xxxx", "xxx", "xxxxx"}
		msg.At.IsAtAll = false
		jsonByte, err := json.Marshal(msg)
		if err != nil {
			log.Println("marshal json error.", msg)
			continue
		}
		jsonString := string(jsonByte)
		log.Printf("%s", jsonString)
		req, err := http.NewRequest("POST", d.Url, strings.NewReader(jsonString))
		if err != nil {
			log.Println("new request error.", err)
			continue
		}
		req.Header.Add("Accept-Charset", "utf-8")
		req.Header.Add("Content-Type", "application/json")
		_, err = d.client.Do(req)
		if err != nil {
			log.Println(err)
		}
	}
}
