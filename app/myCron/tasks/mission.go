package tasks

import (
	"socketAPI/common"
	"time"
)

func MissionsWeek() {
	common.Log("MissionsWeek start")
	for i := 0; i < 500; i++ {

		res, err := common.Db.Exec("delete from missions where `date` < ? LIMIT 1000", time.Now().Add(-7*time.Hour*24*4).Unix())
		if err != nil {
			common.Log("crontab error", err)
			return
		}

		num, err := res.RowsAffected()
		common.Log("del mission 1000 success", num)
		time.Sleep(1 * time.Second)

		if num == 0 {
			break
		}

		if err != nil {
			common.Log("crontab error", err)
			return
		}
	}
}
