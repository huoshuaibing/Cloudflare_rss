package main

import (
	"log"
	"net/http"
	"sync"
	"time"

	"Cloudflare_rss/config"
	"Cloudflare_rss/feed"
	"Cloudflare_rss/robot"
)

// todo : 使用raii技术
// var cfg = config.GetConfig()

// todo : 删除全局变量
var d = robot.DingRobot{}

func init() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)
	// todo : delete this var init func
	// d = robot.DingRobot{}
	d.Init(config.GetConfig().Dingtalk)
}

//Subscription is a tool for check xxxx.atom
func Subscription(target config.Tg, dr robot.DingRobot) bool {
	feeder := new(feed.FeedContext)
	feeder.Service = target.Name
	// cfg := config.GetConfig()
	for {
		if feeder.Canparse(feed.Tg{Name: target.Name, Url: target.Url, Interval: target.Interval}) {
			http.Get(config.GetConfig().Heartbeat)
		}
		if feeder.Update(feed.Tg{Name: target.Name, Url: target.Url, Interval: target.Interval}) {
			dr.Data <- feeder.GetContent()
		}
		// todo : comment this
		// http.Get(config.GetConfig().Heartbeat)
		time.Sleep(time.Second * time.Duration(target.Interval))
	}
}

func main() {
	var wg sync.WaitGroup
	// cfg := config.GetConfig()
	// d := robot.DingRobot{}
	// d.Init(cfg.Dingtalk)
	go d.Run()
	wg.Add(1)
	for _, item := range config.GetConfig().Target {
		go Subscription(item, d)
	}
	wg.Wait()
}
