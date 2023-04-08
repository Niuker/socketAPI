package common

import (
	"strconv"
	"strings"
	"time"
)

func StringStrip(input string) string {
	if input == "" {
		return ""
	}
	return strings.Join(strings.Fields(input), "")
}

func AddUserStartRecord(machine_code string, uid int) {

	var userRecord UserRecord
	userRecord.UserId = uid
	userRecord.MachineCode = machine_code
	userRecord.Time = int(time.Now().Unix())
	userRecord.Types = 1

	_, err := Db.NamedExec(`INSERT INTO userrecord (user_id, machine_code, time, types)
VALUES (:user_id, :machine_code, :time, :types)`, userRecord)
	if err != nil {
		Log("AddUserStartRecord error", err)
	}

}

func AddUserEndRecord(uid string, types int) {

	id, err := strconv.Atoi(uid)
	if err != nil {
		Log("AddUserEndRecord uid error", err, uid)
	}

	var userRecord UserRecord
	userRecord.UserId = id
	userRecord.Time = int(time.Now().Unix())
	userRecord.Types = types

	_, err = Db.NamedExec(`INSERT INTO userrecord (user_id, machine_code, time, types)
VALUES (:user_id, :machine_code, :time, :types)`, userRecord)
	if err != nil {
		Log("AddUserEndRecord error", err)
	}
}
