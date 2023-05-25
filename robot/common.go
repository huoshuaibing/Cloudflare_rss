package robot

import (
	"net/http"
)

type Text struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}
type At struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}
type DTText struct {
	MsgType  string `json:"msgtype"`
	Markdown Text   `json:"markdown"`
	At       At     `json:"at"`
}

//DingRobot is a type
type DingRobot struct {
	Url    string       `json:"Url"`
	client *http.Client `json:"client"`
	Data   chan Text    `json:"Data"`
}
