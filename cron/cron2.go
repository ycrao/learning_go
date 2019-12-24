package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()
	spec := "@every 2s"
	i := 0
	c.AddFunc(spec, func() {
		i++
		fmt.Println("cron running:", i)
	})
	c.Start()
	defer c.Stop()
	select{}
}