package services

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"socketAPI/common"
	"strconv"
	"strings"
	"time"
)

func GetUids(req map[string]string) (interface{}, error) {

	if _, ok := req["source"]; !ok {
		return nil, errors.New("source不能为空")
	}

	var cronuids []common.Cronuid
	page := "1"
	if _, ok := req["page"]; ok {
		page = req["page"]
	}

	p, err := strconv.Atoi(page)
	if err != nil {
		return nil, err
	}
	fmt.Println(time.Now().Unix())

	err = common.Db.Select(&cronuids, "select user_id,exp_time,source,name from cronuid where del=0 and source=? and exp_time>? order by exp_time limit ?,200", req["source"], time.Now().Unix(), (p-1)*10)
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
	if _, ok := req["name"]; !ok {
		return nil, errors.New("name不能为空")
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

	err = common.Db.Select(&cronuids, "select *  from cronuid where user_id = ? and del=0", user_id)
	if err != nil {
		return nil, err
	}
	if len(cronuids) > 0 {
		_, err = common.Db.Exec("update cronuid set exp_time=?,source=?,name=? where user_id = ? and del=0", exp_time, source, req["name"], user_id)
		if err != nil {
			return nil, err
		}

		return cronuids[0], nil
	}

	cronuid.UserId = user_id
	cronuid.Source = source
	cronuid.ExpTime = exp_time
	cronuid.Name = req["name"]
	cronuid.Del = 0
	_, err = common.Db.NamedExec(`INSERT INTO cronuid (user_id, source, exp_time,del,name) 
VALUES (:user_id, :source, :exp_time, :del, :name)`, cronuid)
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

	sql, inIds, err := sqlx.In("update cronuid set del=1 where user_id IN (?)", ids)
	if err != nil {
		return nil, err
	}

	_, err = common.Db.Exec(sql, inIds...)
	if err != nil {
		return nil, err
	}
	return ids, nil
}

func GetGifts(req map[string]string) (interface{}, error) {
	var CronGiftcodes []common.CronGiftcode

	err := common.Db.Select(&CronGiftcodes, "select id,code from crongiftcode where del=0 order by id")
	if err != nil {
		return nil, err
	}
	return CronGiftcodes, nil
}

func AddGifts(req map[string]string) (interface{}, error) {
	if _, ok := req["code"]; !ok {
		return nil, errors.New("code不能为空")
	}

	var CronGiftcode common.CronGiftcode
	var CronGiftcodes []common.CronGiftcode

	err := common.Db.Select(&CronGiftcodes, "select id  from crongiftcode where code = ? and del=0", req["code"])
	if err != nil {
		return nil, err
	}
	if len(CronGiftcodes) > 0 {
		return nil, errors.New("code exisit")
	}

	CronGiftcode.Code = req["code"]
	CronGiftcode.Del = 0
	_, err = common.Db.NamedExec(`INSERT INTO crongiftcode (code, del) 
VALUES (:code, :del)`, CronGiftcode)
	if err != nil {
		return nil, err
	}

	return CronGiftcode, nil
}

func DelGifts(req map[string]string) (interface{}, error) {
	if _, ok := req["id"]; !ok {
		return nil, errors.New("id不能为空")
	}
	ids := strings.Split(req["id"], ",")

	if len(ids) > 200 || len(ids) < 1 {
		return nil, errors.New("id不能超过200")
	}

	sql, inIds, err := sqlx.In("update crongiftcode set del=1 where id IN (?)", ids)
	if err != nil {
		return nil, err
	}

	_, err = common.Db.Exec(sql, inIds...)
	if err != nil {
		return nil, err
	}
	return ids, nil
}
