package services

import (
	"errors"
	"socketAPI/common"
)

func GetVersions(req map[string]string) (interface{}, error) {

	if _, ok := req["name"]; !ok {
		return nil, errors.New("name不能为空")
	}

	var cronVersion []common.CronVersion

	err := common.Db.Select(&cronVersion, "select * from cronversion  where name=?", req["name"])
	if err != nil {
		return nil, err
	}

	if len(cronVersion) == 0 {
		return []struct{}{}, nil
	}

	return cronVersion, nil
}
