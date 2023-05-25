package config

//Tg is a type
type Tg struct {
	Name     string
	Url      string
	Interval int
}

//Config is a type
type Config struct {
	Target    []Tg   `json:"Target" `
	Dingtalk  string `json:"Dingtalk"`
	Heartbeat string `json:"Heartbeat"`
}
