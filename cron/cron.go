package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	c := cron.New()
	c.AddFunc("30 * * * *", func() { fmt.Println("Every hour on the half hour") })
	c.AddFunc("30 3-6,20-23 * * *", func() { fmt.Println(".. in the range 3-6am, 8-11pm") })
	c.AddFunc("CRON_TZ=Asia/Tokyo 30 04 * * *", func() { fmt.Println("Runs at 04:30 Tokyo time every day") })
	c.AddFunc("@hourly", func() { fmt.Println("Every hour, starting an hour from now") })
	c.AddFunc("@every 2s", func() { fmt.Println("Every 2 seconds, starting an 2 seconds from now") })
	go c.Start()
	defer c.Stop()
	select {
	case <-time.After(time.Second * 10):
		return
	}
}
