package tasks

import (
	"socketAPI/common"
	"time"
)

func TimersWeek() {
	common.Log("TimersWeek start")

	for i := 0; i < 1000; i++ {
		res, err := common.Db.Exec("delete from timers where `value` < ? LIMIT 1000", time.Now().Add(-7*time.Hour*24*4).Unix())
		common.Log("timer crontab", "delete from timers where `value` < ", time.Now().Add(-7*time.Hour*24*4).Unix())
		if err != nil {
			common.Log("crontab error", err)
			return
		}

		lastrId, err := res.LastInsertId()
		if lastrId == 0 {
			break
		}
		if err != nil {
			common.Log("crontab error", err)
		}
	}
}
