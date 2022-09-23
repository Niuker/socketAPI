package services

import (
	"errors"
	"fmt"
	"socketAPI/app/encr"
	common2 "socketAPI/common"
	"socketAPI/config"
	"strconv"
	"time"
)

func initMission(id int, date int) error {
	var missions []common2.Missions
	var missionField []common2.MissionField
	var insertMissions []common2.Missions

	err := common2.Db.Select(&missions, "select * from missions where user_id=? and date=?", id, date)
	if err != nil {
		common2.Log("exec failed, ", err)
		return errors.New("get missions error")
	}
	err = common2.Db.Select(&missionField, "select * from mission_field")
	if err != nil {
		common2.Log("exec failed, ", err)
		return errors.New("get mission field error")
	}
	if len(missions) == len(missionField) {
		return nil
	}

outside:
	for _, mf := range missionField {
		for _, m := range missions {
			if m.MissionFieldId == mf.Id {
				continue outside
			}
		}
		var tmpMission common2.Missions
		tmpMission.MissionFieldId = mf.Id
		tmpMission.UserId = id
		tmpMission.Value = mf.Default
		tmpMission.UpdateTime = int(time.Now().Unix())
		tmpMission.Date = date
		insertMissions = append(insertMissions, tmpMission)
	}
	fmt.Println(insertMissions[0].UserId)
	_, err = common2.Db.NamedExec(`INSERT INTO missions (user_id, mission_field_id, value,update_time,date) 
VALUES (:user_id, :mission_field_id, :value, :update_time, :date)`, insertMissions)
	if err != nil {
		common2.Log("exec failed, ", err)
		return errors.New("insert missions field error")
	}
	return nil
}

func getMissions(id int, isday string, date int) ([]common2.MissionsANDMissionField, error) {
	var missions []common2.MissionsANDMissionField
	err := common2.Db.Select(&missions, `select user_id, value, name, mf.default, isday 
from missions as m left join mission_field as mf ON m.mission_field_id = mf.id
where m.user_id=? and  m.date=? and mf.isday=?`, id, date, isday)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return nil, errors.New("getMissions field error")
	}

	return missions, nil
}

func GetMissions(req map[string]string) (interface{}, error) {
	if _, ok := req["user_id"]; !ok {
		return nil, errors.New("user不能为空")
	}
	if _, ok := req["date"]; !ok {
		return nil, errors.New("date不能为空")
	}
	mid, err := encr.ECBDecrypter(config.MyConfig.ENCR.Desckey, req["user_id"])
	if mid == "" || err != nil {
		return nil, errors.New("本次user解密失败")
	}
	id, err := strconv.Atoi(mid)
	if err != nil {
		return nil, errors.New("id错误")
	}
	date, err := strconv.Atoi(req["date"])
	if err != nil {
		return nil, errors.New("date错误")
	}

	isday := "1"
	if _, ok := req["isday"]; ok {
		isday = req["isday"]
	}

	err = initMission(id, date)
	if err != nil {
		return nil, err
	}

	m, err := getMissions(id, isday, date)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func SetMissions(req map[string]string) (interface{}, error) {
	if _, ok := req["user_id"]; !ok {
		return nil, errors.New("user不能为空")
	}
	if _, ok := req["date"]; !ok {
		return nil, errors.New("date不能为空")
	}
	mid, err := encr.ECBDecrypter(config.MyConfig.ENCR.Desckey, req["user_id"])
	if mid == "" || err != nil {
		return nil, errors.New("本次user解密失败")
	}
	id, err := strconv.Atoi(mid)
	if err != nil {
		return nil, errors.New("id错误")
	}
	date, err := strconv.Atoi(req["date"])
	if err != nil {
		return nil, errors.New("date错误")
	}

	isday := "1"
	if _, ok := req["isday"]; ok {
		isday = req["isday"]
	}

	err = initMission(id, date)
	if err != nil {
		return nil, err
	}

	var missionField []common2.MissionField
	err = common2.Db.Select(&missionField, "select * from mission_field")
	if err != nil {
		common2.Log("exec failed, ", err)
		return nil, errors.New("get mission field error")
	}

	for _, mf := range missionField {
		if _, ok := req[mf.Name]; ok {
			if req[mf.Name] == "add" {
				_, err = common2.Db.Exec("update missions set `value`=value+1 where date=? and user_id=? and mission_field_id=?",
					date, id, mf.Id)
			} else if req[mf.Name] == "default" {
				_, err = common2.Db.Exec("update missions set value=? where date=? and user_id=? and mission_field_id=? ",
					mf.Default, date, id, mf.Id)
			} else {
				_, err = common2.Db.Exec("update missions set value=? where date=? and user_id=? and mission_field_id=? ",
					req[mf.Name], date, id, mf.Id)
			}
			if err != nil {
				common2.Log("exec failed, ", err)
				return nil, errors.New("update missions field error")
			}
		}
	}
	m, err := getMissions(id, isday, date)
	if err != nil {
		return nil, err
	}
	return m, nil
}
