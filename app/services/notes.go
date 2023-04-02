package services

import (
	"errors"
	"socketAPI/app/encr"
	"socketAPI/common"
	"socketAPI/config"
	"strconv"
	"time"
)

func GetNotes(req map[string]string) (interface{}, error) {
	var notes []common.Notes

	err := common.Db.Select(&notes, "select * from notes  where time>?", int(time.Now().Unix()))
	if err != nil {
		return nil, err
	}

	if len(notes) == 0 {
		return []struct{}{}, nil
	}

	return notes, nil
}

func AddNotes(req map[string]string) (interface{}, error) {
	if _, ok := req["notes"]; !ok {
		return nil, errors.New("notes不能为空")
	}
	if _, ok := req["exp"]; !ok {
		return nil, errors.New("exp不能为空")
	}
	exp, err := strconv.Atoi(req["exp"])
	if err != nil {
		return nil, errors.New("exp错误")
	}

	if _, ok := req["machine_code"]; !ok {
		return nil, errors.New("machine_code不能为空")
	}
	machineCode, err := encr.ECBDecrypter(config.MyConfig.ENCR.Desckey, req["machine_code"])

	if machineCode == "" || err != nil {
		return nil, errors.New("本次machine_code解密失败")
	}

	var fontmanagers []common.Fontmanager
	err = common.Db.Select(&fontmanagers, "select * from fontmanager  where machine_code=? ", machineCode)
	if err != nil {
		return nil, err
	}

	if len(fontmanagers) == 0 {
		return nil, errors.New("非管理员无法掉用此接口")
	}

	var notes []common.Notes
	var note common.Notes

	err = common.Db.Select(&notes, "select * from notes  where notes = ?", req["notes"])
	if err != nil {
		return nil, err
	}

	if len(notes) == 0 {
		note.Notes = req["notes"]
		note.Time = exp
		_, err = common.Db.NamedExec(`INSERT INTO notes (notes, time) 
VALUES (:notes, :time)`, note)
		if err != nil {
			return nil, err
		}
		return true, nil
	}

	if len(notes) == 1 {
		_, err = common.Db.Exec("update notes set `time`=? where  notes=?",
			exp, req["notes"])
		return true, nil
	}

	return nil, errors.New("note数据异常")

}
