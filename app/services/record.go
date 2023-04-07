package services

import (
	"errors"
	"socketAPI/common"
)

func GetUserRecord(req map[string]string) (interface{}, error) {
	if _, ok := req["start"]; !ok {
		return nil, errors.New("start不能为空")
	}
	if _, ok := req["end"]; !ok {
		return nil, errors.New("end不能为空")
	}

	var userRecord []common.UserRecord

	if _, ok := req["types"]; !ok {
		err := common.Db.Select(&userRecord, "select * from userrecord  where time>? and time <? limit 10000", req["start"], req["end"])
		if err != nil {
			return nil, err
		}

		if len(userRecord) == 0 {
			return []struct{}{}, nil
		}
	} else {
		err := common.Db.Select(&userRecord, "select * from userrecord  where time>? and time <? and types = ? limit 10000", req["start"], req["end"], req["types"])
		if err != nil {
			return nil, err
		}

		if len(userRecord) == 0 {
			return []struct{}{}, nil
		}
	}

	return userRecord, nil
}
