package feed

import (
	"log"

	"Cloudflare_rss/robot"

	"github.com/lunny/html2md"

	"github.com/mmcdole/gofeed"
)

// Tg is cool
type Tg struct {
	Name     string
	Url      string
	Interval int
}

//FeedContext is cool
type FeedContext struct {
	Title   string
	Service string
	Version string
	Url     string
	Time    string
	Content string
}

//Canparse check the url is ok
func (fc *FeedContext) Canparse(t Tg) bool {
	parse := gofeed.NewParser()
	_, err := parse.ParseURL(t.Url)
	if err != nil {
		log.Print("New parse of atom error", err)
		return false
	}
	return true
}

// Update check新的atom
func (fc *FeedContext) Update(t Tg) bool {
	parse := gofeed.NewParser()
	f, err := parse.ParseURL(t.Url)
	if err != nil {
		log.Print("New parse of atom error", err)
		return false
	}
	if len(f.Items) == 0 {
		return false
	}
	it := f.Items[0]
	md := html2md.Convert(it.Content)
	init := true
	if fc.Version != "" && fc.Title != "" {
		init = false
	}
	// 以Version 定义是否更新
	if fc.Version == it.Title {
		return false
	}
	fc.Title = fc.Service + "New Issue"
	fc.Version = it.Title
	fc.Url = it.Link
	fc.Time = it.Updated
	fc.Content = md
	return !init
}

//GetContent is a func to send msg
func (fc *FeedContext) GetContent() (c robot.Text) {
	c.Title = fc.Title
	content := "# " + fc.Title + "\n\n" +
		"----------------------------\n\n" +
		// "# Overview \n\n" +
		"### Issue: \n\n" + fc.Version + "\n\n" +
		"### Url: \n\n" + fc.Url + "\n\n" +
		"###Time: \n\n" + fc.Time + "\n\n" +
		"---------------------------\n\n" +
		/*"# Detail \n\n"*/ fc.Content
	c.Text = content
	return c
}
