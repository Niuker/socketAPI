package myCron

import (
	"github.com/robfig/cron"
	"socketAPI/app/myCron/tasks"
)

func Start() {
	c := cron.New()
	c.AddFunc("0 0 2 * * ?", func() {
		tasks.MissionsWeek()
		tasks.TimersWeek()
	})
	//c.AddFunc("0 0 0-4,6-23 * * ?", func() {
	//	//c.AddFunc("*/5 * * * * ?", func() {
	//	tasks.GiftStart()
	//})
	c.Start()

}
