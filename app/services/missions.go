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

func initMission(id int, date int, mcode string, repeat bool) error {
	var missions []common.Missions
	var missionField []common.MissionField
	var insertMissions []common.Missions

	err := common.Db.Select(&missions, "select * from missions where user_id=? and date=? and machine_code=?", id, date, mcode)

	if err != nil {
		common.Log("exec failed, ", err)
		return errors.New("get missions error")
	}

	if len(missions) == 0 && mcode != "" && !repeat {
		_, err = common.Db.Exec("update missions set machine_code=? where user_id=? and date=? and machine_code=''", mcode, id, date)
		if err != nil {
			common.Log("exec update missions failed, ", err)
			return errors.New("get missions 2 error")
		}
		return initMission(id, date, mcode, true)
	}

	err = common.Db.Select(&missionField, "select * from mission_field")
	if err != nil {
		common.Log("exec failed, ", err)
		return errors.New("get mission field error1")
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
		var tmpMission common.Missions
		tmpMission.MissionFieldId = mf.Id
		tmpMission.UserId = id
		tmpMission.Value = mf.Default
		tmpMission.UpdateTime = int(time.Now().Unix())
		tmpMission.Date = date
		tmpMission.MachineCode = mcode
		insertMissions = append(insertMissions, tmpMission)
	}
	if len(insertMissions) == 0 {
		return nil
	}
	_, err = common.Db.NamedExec(`INSERT INTO missions (user_id, mission_field_id, value,update_time,date,machine_code) 
VALUES (:user_id, :mission_field_id, :value, :update_time, :date,:machine_code)`, insertMissions)
	if err != nil {
		common.Log("exec failed, ", err)
		return errors.New("insert missions field error")
	}
	return nil
}

func getMissions(id int, isday string, date int, mcode string) ([]common.MissionsANDMissionField, error) {
	var missions []common.MissionsANDMissionField
	err := common.Db.Select(&missions, `select user_id, value, name, mf.default, isday 
from missions as m left join mission_field as mf ON m.mission_field_id = mf.id
where m.user_id=? and  m.date=? and mf.isday=? and m.machine_code=?`, id, date, isday, mcode)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return nil, errors.New("getMissions field error")
	}

	var missionField []common.MissionField
	err = common.Db.Select(&missionField, "select * from mission_field")
	if err != nil {
		common.Log("exec failed, ", err)
		return nil, errors.New("get mission field error2")
	}

outside:
	for _, mf := range missionField {
		for _, m := range missions {
			if m.MissionFieldId == mf.Id {
				continue outside
			}
		}
		var tmpMission common.MissionsANDMissionField
		tmpMission.UserId = id
		tmpMission.Value = mf.Default
		tmpMission.Name = mf.Name
		tmpMission.Default = mf.Default
		tmpMission.Isday = mf.Isday
		missions = append(missions, tmpMission)
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
	if req["machine_code"] != "" {
		req["machine_code"], err = encr.ECBDecrypter(config.MyConfig.ENCR.Desckey, req["machine_code"])
		if err != nil {
			return nil, errors.New("本次machine_code解密失败")
		}
	}
	date, err := strconv.Atoi(req["date"])
	if err != nil {
		return nil, errors.New("date错误")
	}

	isday := "1"
	if _, ok := req["isday"]; ok {
		isday = req["isday"]
	}

	m, err := getMissions(id, isday, date, req["machine_code"])
	common.Log(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func SetMissionsWithMachine(req map[string]string) (interface{}, error) {
	err := VerifyMachine(req, false)
	if err != nil {
		return nil, err
	}

	return SetMissions(req)
}

func GetMissionsWithMachine(req map[string]string) (interface{}, error) {
	err := VerifyMachine(req, false)

	if err != nil {
		return nil, err
	}

	return GetMissions(req)
}

func SetMissionsWithMachineStrict(req map[string]string) (interface{}, error) {
	err := VerifyMachine(req, true)
	if err != nil {
		return nil, err
	}

	return SetMissions(req)
}

func GetMissionsWithMachineStrict(req map[string]string) (interface{}, error) {
	err := VerifyMachine(req, true)
	if err != nil {
		return nil, err
	}

	return GetMissions(req)
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
	if req["machine_code"] != "" {
		req["machine_code"], err = encr.ECBDecrypter(config.MyConfig.ENCR.Desckey, req["machine_code"])
		if err != nil {
			return nil, errors.New("本次machine_code解密失败")
		}
	}

	date, err := strconv.Atoi(req["date"])
	if err != nil {
		return nil, errors.New("date错误")
	}

	isday := "1"
	if _, ok := req["isday"]; ok {
		isday = req["isday"]
	}

	var missionField []common.MissionField
	err = common.Db.Select(&missionField, "select * from mission_field")
	if err != nil {
		common.Log("exec failed, ", err)
		return nil, errors.New("get mission field error3")
	}

	for _, mf := range missionField {
		if _, ok := req[mf.Name]; ok {
			var missions []common.Missions
			err = common.Db.Select(&missions, "select * from missions where date=? and user_id=? and mission_field_id=? and (machine_code=? or machine_code='')",
				date, id, mf.Id, req["machine_code"])
			if err != nil {
				common.Log("exec failed, ", err)
				return nil, errors.New("get mission field error4")
			}
			if len(missions) == 0 {
				var insertMission common.Missions
				insertMission.MissionFieldId = mf.Id
				insertMission.UserId = id
				insertMission.MachineCode = req["machine_code"]
				insertMission.Date = date
				insertMission.UpdateTime = int(time.Now().Unix())
				insertMission.Value = mf.Default
				_, err = common.Db.NamedExec(`INSERT INTO missions (user_id, mission_field_id, value,update_time,date,machine_code) 
VALUES (:user_id, :mission_field_id, :value, :update_time, :date,:machine_code)`, insertMission)

				if err != nil {
					common.Log("exec failed, ", err)
					return nil, errors.New("get mission field error5")
				}
			}

			if req[mf.Name] == "add" {
				_, err = common.Db.Exec("update missions set `value`=value+1 where date=? and user_id=? and mission_field_id=? and (machine_code=? or machine_code='')",
					date, id, mf.Id, req["machine_code"])
			} else if req[mf.Name] == "default" {
				_, err = common.Db.Exec("update missions set value=? where date=? and user_id=? and mission_field_id=? and (machine_code=? or machine_code='')",
					mf.Default, date, id, mf.Id, req["machine_code"])
			} else {
				_, err = common.Db.Exec("update missions set value=? where date=? and user_id=? and mission_field_id=? and (machine_code=? or machine_code='')",
					req[mf.Name], date, id, mf.Id, req["machine_code"])
			}
			if err != nil {
				common.Log("exec failed, ", err)
				return nil, errors.New("update missions field error")
			}
		}
	}
	m, err := getMissions(id, isday, date, req["machine_code"])
	if err != nil {
		return nil, err
	}
	return m, nil
}
