package tasks

import (
	"socketAPI/common"
	"time"
)

func TimersWeek() {
	for i := 0; i < 1000; i++ {
		_, err := common.Db.Exec("delete from timers where `value` < ? LIMIT 1000", time.Now().Add(-7*time.Hour*24*4).Unix())
		if err != nil {
			common.Log("crontab error", err)
		}
	}
}
