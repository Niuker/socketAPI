package tasks

import (
	"github.com/jmoiron/sqlx"
	"socketAPI/common"
	"time"
)

func MissionsWeek() {
	delMissions(time.Now().Add(-7 * time.Hour * 24 * 4).Unix())

}
func MissionsDay() {
	delMissions(time.Now().Add(-7 * time.Hour * 24).Unix())
}

func delMissions(t int64) {
	var missionField []common.MissionField

	err := common.Db.Select(&missionField, "select id from mission_field")

	if err != nil {
		common.Log("crontab error", err)
	}
	var mfids []int
	for _, mf := range missionField {
		mfids = append(mfids, mf.Id)
	}

	for i := 0; i < 1000; i++ {
		sql, inIds, err := sqlx.In("delete from missions where `date` < ? and mission_field_id IN (?) LIMIT 1000",
			t, mfids)
		if err != nil {
			common.Log("crontab error", err)
		}

		res, err := common.Db.Exec(sql, inIds...)
		if err != nil {
			common.Log("crontab error", err)
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
