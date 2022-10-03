package myCron

import (
	"github.com/robfig/cron"
	"socketAPI/app/myCron/tasks"
)

func Start() {
	c := cron.New()
	c.AddFunc("0 0 5 * * ?", func() {
		tasks.MissionsDay()
		tasks.MissionsWeek()
		tasks.TimersWeek()
	})
	c.AddFunc("0 0 */1 * * ?", func() {
		//c.AddFunc("*/5 * * * * ?", func() {
		tasks.GiftStart()
	})
	c.Start()

}
