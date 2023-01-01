package services

import (
	"errors"
	"fmt"
	"socketAPI/app/encr"
	"socketAPI/common"
	"socketAPI/config"
	"strconv"
	"time"
)

func initTimer(id int) error {
	var timers []common.Timers
	var timerField []common.TimerField
	var insertTimers []common.Timers

	err := common.Db.Select(&timers, "select * from timers where user_id=?", id)
	if err != nil {
		common.Log("exec failed, ", err)
		return errors.New("get timers error")
	}

	err = common.Db.Select(&timerField, "select * from timer_field")
	if err != nil {
		common.Log("exec failed, ", err)
		return errors.New("get timer field error")
	}
	if len(timers) == len(timerField) {
		return nil
	}

outside:
	for _, mf := range timerField {
		for _, m := range timers {
			if m.TimerFieldId == mf.Id {
				continue outside
			}
		}
		var tmpTimer common.Timers
		tmpTimer.TimerFieldId = mf.Id
		tmpTimer.UserId = id
		tmpTimer.Value = mf.Default
		tmpTimer.UpdateTime = int(time.Now().Unix())
		insertTimers = append(insertTimers, tmpTimer)
	}
	_, err = common.Db.NamedExec(`INSERT INTO timers (user_id, timer_field_id, value,update_time) 
VALUES (:user_id, :timer_field_id, :value, :update_time)`, insertTimers)
	if err != nil {
		common.Log("exec failed, ", err)
		return errors.New("insert timers field error")
	}
	return nil
}

func getTimers(id int) ([]common.TimersANDTimerField, error) {
	var timers []common.TimersANDTimerField
	err := common.Db.Select(&timers, `select user_id, value, name, mf.default 
from timers as m left join timer_field as mf ON m.timer_field_id = mf.id
where m.user_id=?`, id)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return nil, errors.New("getTimers field error")
	}

	return timers, nil
}

func GetTimers(req map[string]string) (interface{}, error) {
	if _, ok := req["user_id"]; !ok {
		return nil, errors.New("user不能为空")
	}
	mid, err := encr.ECBDecrypter(config.MyConfig.ENCR.Desckey, req["user_id"])
	if mid == "" || err != nil {
		return nil, errors.New("本次user解密失败")
	}
	id, err := strconv.Atoi(mid)
	if err != nil {
		return nil, errors.New("id错误")
	}

	err = initTimer(id)
	if err != nil {
		return nil, err
	}

	m, err := getTimers(id)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func SetTimersWithMachine(req map[string]string) (interface{}, error) {
	err := VerifyMachine(req)
	if err != nil {
		return nil, err
	}

	return SetTimers(req)
}

func GetTimersWithMachine(req map[string]string) (interface{}, error) {
	err := VerifyMachine(req)
	if err != nil {
		return nil, err
	}

	return GetTimers(req)
}

func SetTimers(req map[string]string) (interface{}, error) {
	if _, ok := req["user_id"]; !ok {
		return nil, errors.New("user不能为空")
	}
	mid, err := encr.ECBDecrypter(config.MyConfig.ENCR.Desckey, req["user_id"])
	if mid == "" || err != nil {
		return nil, errors.New("本次user解密失败")
	}
	id, err := strconv.Atoi(mid)
	if err != nil {
		return nil, errors.New("id错误")
	}

	err = initTimer(id)
	if err != nil {
		return nil, err
	}

	var timerField []common.TimerField
	err = common.Db.Select(&timerField, "select * from timer_field")
	if err != nil {
		common.Log("exec failed, ", err)
		return nil, errors.New("get timer field error")
	}

	for _, mf := range timerField {
		if _, ok := req[mf.Name]; ok {
			if req[mf.Name] == "add" {
				_, err = common.Db.Exec("update timers set `value`=value+1 where  user_id=? and timer_field_id=?",
					id, mf.Id)
			} else if req[mf.Name] == "default" {
				_, err = common.Db.Exec("update timers set value=? where user_id=? and timer_field_id=? ",
					mf.Default, id, mf.Id)
			} else {
				_, err = common.Db.Exec("update timers set value=? where user_id=? and timer_field_id=? ",
					req[mf.Name], id, mf.Id)
			}
			if err != nil {
				common.Log("exec failed, ", err)
				return nil, errors.New("update timers field error")
			}
		}
	}
	m, err := getTimers(id)
	if err != nil {
		return nil, err
	}
	return m, nil
}
