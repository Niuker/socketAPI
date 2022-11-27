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

func GetMessages(req map[string]string) (interface{}, error) {
	var messages []common.Message
	page := "1"
	if _, ok := req["page"]; ok {
		page = req["page"]
	}

	p, err := strconv.Atoi(page)
	if err != nil {
		return nil, err
	}

	if title, ok := req["title"]; ok {
		err = common.Db.Select(&messages, "select * from message where title=? order by time limit ?,200", title, (p-1)*200)
	} else {
		err = common.Db.Select(&messages, "select * from message  order by time limit ?,200", (p-1)*200)
	}
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func AddMessages(req map[string]string) (interface{}, error) {
	if _, ok := req["title"]; !ok {
		return nil, errors.New("title不能为空")
	}

	imei := ""
	if _, ok := req["imei"]; ok {
		imei = req["imei"]
	}

	content := ""
	if _, ok := req["content"]; ok {
		content = req["content"]
	}

	var messages []common.Message
	var message common.Message
	err := common.Db.Select(&messages, "select * from message  where imei=? and time>?",
		imei, time.Now().Add(-5*time.Minute).Format("2012-01-02 15:04:05"))

	if err != nil {
		return nil, err
	}
	if len(messages) > 0 {
		return nil, errors.New("can not repeat on 5 minutes")
	}

	err = common.Db.Select(&messages, "select * from message  where content=?", content)
	if err != nil {
		return nil, err
	}
	if len(messages) > 0 {
		return nil, errors.New("content repeat")
	}

	message.Title = req["title"]
	message.Content = content
	message.Time = time.Now().Format("2006-01-02 15:04:05")
	message.Imei = imei
	_, err = common.Db.NamedExec(`INSERT INTO message (title, content, time,imei) 
VALUES (:title, :content, :time, :imei)`, message)
	if err != nil {
		return nil, err
	}

	return message, nil
}

func DelMessages(req map[string]string) (interface{}, error) {
	if _, ok := req["id"]; !ok {
		return nil, errors.New("id不能为空")
	}
	fmt.Printf("%T,,,,%s", req["id"], req["id"])
	ids := strings.Split(req["id"], ",")

	if len(ids) > 200 || len(ids) < 1 {
		return nil, errors.New("id不能超过200")
	}

	for _, id := range ids {
		_, err := strconv.Atoi(id)
		if err != nil {
			return nil, errors.New("id非法")
		}
	}

	sql, inIds, err := sqlx.In("delete from message where id IN (?)", ids)
	if err != nil {
		return nil, err
	}

	_, err = common.Db.Exec(sql, inIds...)
	if err != nil {
		return nil, err
	}
	return ids, nil
}
