package services

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"socketAPI/common"
	"strconv"
	"strings"
)

func GetUids(req map[string]string) (interface{}, error) {
	var cronuids []common.Cronuid
	page := "1"
	if _, ok := req["page"]; ok {
		page = req["page"]
	}

	p, err := strconv.Atoi(page)
	if err != nil {
		return nil, err
	}

	err = common.Db.Select(&cronuids, "select user_id,exp_time,source from cronuid  order by exp_time limit ?,200", (p-1)*10)
	if err != nil {
		return nil, err
	}
	return cronuids, nil
}

func AddUids(req map[string]string) (interface{}, error) {
	if _, ok := req["user_id"]; !ok {
		return nil, errors.New("user_id不能为空")
	}

	if _, ok := req["exp_time"]; !ok {
		return nil, errors.New("exp_time不能为空")
	}

	if _, ok := req["source"]; !ok {
		return nil, errors.New("source不能为空")
	}

	var cronuid common.Cronuid
	var cronuids []common.Cronuid

	user_id, err := strconv.Atoi(req["user_id"])
	if err != nil {
		return nil, err
	}
	exp_time, err := strconv.Atoi(req["exp_time"])
	if err != nil {
		return nil, err
	}
	source, err := strconv.Atoi(req["source"])
	if err != nil {
		return nil, err
	}

	err = common.Db.Select(&cronuids, "select id  from cronuid where user_id = ?", user_id)
	if err != nil {
		return nil, err
	}
	if len(cronuids) > 0 {
		return nil, errors.New("uid exisit")

	}

	cronuid.UserId = user_id
	cronuid.Source = source
	cronuid.ExpTime = exp_time
	_, err = common.Db.NamedExec(`INSERT INTO cronuid (user_id, source, exp_time) 
VALUES (:user_id, :source, :exp_time)`, cronuid)
	if err != nil {
		return nil, err
	}

	return cronuid, nil
}

func DelUids(req map[string]string) (interface{}, error) {
	if _, ok := req["user_id"]; !ok {
		return nil, errors.New("id不能为空")
	}
	ids := strings.Split(req["user_id"], ",")

	if len(ids) > 200 || len(ids) < 1 {
		return nil, errors.New("user_id不能超过200")
	}

	for _, id := range ids {
		_, err := strconv.Atoi(id)
		if err != nil {
			return nil, errors.New("id非法")
		}
	}

	sql, inIds, err := sqlx.In("delete from cronuid where user_id IN (?)", ids)
	if err != nil {
		return nil, err
	}

	_, err = common.Db.Exec(sql, inIds...)
	if err != nil {
		return nil, err
	}
	return ids, nil
}
