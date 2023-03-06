package services

import (
	"errors"
	"socketAPI/app/encr"
	"socketAPI/common"
	"socketAPI/config"
	"strconv"
	"time"
)

func SetMachines(req map[string]string) (interface{}, error) {
	if _, ok := req["user_id"]; !ok {
		return nil, errors.New("user不能为空")
	}
	if _, ok := req["machine_code"]; !ok {
		return nil, errors.New("machine_code不能为空")
	}
	machineCode, err := encr.ECBDecrypter(config.MyConfig.ENCR.Desckey, req["machine_code"])
	if machineCode == "" || err != nil {
		return nil, errors.New("本次machine_code解密失败")
	}
	mid, err := encr.ECBDecrypter(config.MyConfig.ENCR.Desckey, req["user_id"])
	if mid == "" || err != nil {
		return nil, errors.New("本次user解密失败")
	}
	id, err := strconv.Atoi(mid)
	if err != nil {
		return nil, errors.New("id错误")
	}

	var machines []common.Machines
	var machine common.Machines
	err = common.Db.Select(&machines, "select id,machine_code from machines where  user_id = ? and machine_code=?", id, machineCode)
	if err != nil {
		return nil, err
	}

	if len(machines) > 1 {
		return nil, errors.New("machine数据异常")
	}

	machine.MachineCode = machineCode
	machine.UserId = id
	machine.UpdateTime = int(time.Now().Unix())

	if len(machines) == 1 {
		//if machines[0].MachineCode == machineCode {
		_, err = common.Db.Exec("update machines set update_time= ?  where  user_id=? and machine_code=?",
			int(time.Now().Unix()), id, machineCode)
		if err != nil {
			common.Log("update machines error", err)
			return nil, err
		}
		return machine, nil
		//}
		//		else {
		//			// del machines missions timers
		//			_, err = common.Db.Exec("delete from machines where `user_id` = ?", id)
		//			if err != nil {
		//				common.Log("del machines error", err)
		//				return nil, err
		//			}
		//			//_, err = common.Db.Exec("delete from missions where `user_id` = ?", id)
		//			//if err != nil {
		//			//	common.Log("del missions error", err)
		//			//	return nil, err
		//			//}
		//			//_, err = common.Db.Exec("delete from timers where `user_id` = ?", id)
		//			//if err != nil {
		//			//	common.Log("del timers error", err)
		//			//	return nil, err
		//			//}
		//
		//			_, err = common.Db.NamedExec(`INSERT INTO machines (machine_code, user_id,update_time)
		//VALUES (:machine_code, :user_id, :update_time)`, machine)
		//			if err != nil {
		//				return nil, err
		//			}
		//			return machine, nil
		//
		//		}
	}

	if len(machines) < 1 {
		_, err = common.Db.NamedExec(`INSERT INTO machines (machine_code, user_id,update_time) 
VALUES (:machine_code, :user_id, :update_time)`, machine)
		if err != nil {
			return nil, err
		}
	}

	return machine, nil
}

func GetMachines(req map[string]string) (interface{}, error) {
	if _, ok := req["machine_code"]; !ok {
		return nil, errors.New("machine_code不能为空")
	}
	machineCode, err := encr.ECBDecrypter(config.MyConfig.ENCR.Desckey, req["machine_code"])
	if machineCode == "" || err != nil {
		return nil, errors.New("本次machine_code解密失败")
	}

	var machines []common.Machines

	err = common.Db.Select(&machines, "select * from machines where machine_code = ? ", machineCode)
	if err != nil {
		return nil, err
	}

	for k, m := range machines {
		machines[k].Mid, err = encr.ECBEncrypt(config.MyConfig.ENCR.Desckey, strconv.Itoa(m.UserId))
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}
	return machines, nil
}

func VerifyMachine(req map[string]string, strict bool) error {
	if _, ok := req["user_id"]; !ok {
		return errors.New("user不能为空")
	}
	if _, ok := req["machine_code"]; !ok {
		req["machine_code"] = ""
	}

	mcode := ""
	if req["machine_code"] != "" {
		var err error
		mcode, err = encr.ECBDecrypter(config.MyConfig.ENCR.Desckey, req["machine_code"])
		if err != nil {
			return errors.New("本次macine_code解密失败")
		}
	}

	mid, err := encr.ECBDecrypter(config.MyConfig.ENCR.Desckey, req["user_id"])
	if mid == "" || err != nil {
		return errors.New("本次user解密失败")
	}
	id, err := strconv.Atoi(mid)
	if err != nil {
		return errors.New("id错误")
	}

	var machines []common.Machines

	if mcode != "" {
		err = common.Db.Select(&machines, "select * from machines where machine_code = ? and user_id = ? ", mcode, id)

		if len(machines) < 1 {
			return errors.New("machine不存在")
		}
	}

	if _, ok := req["date"]; !ok {
		return nil
	}

	if strict {
		var missions []common.Missions
		err = common.Db.Select(&missions, "select * from missions where machine_code = '' and user_id = ? and date=? ", id, req["date"])

		if len(missions) > 0 {
			return errors.New("machine_code can not be empty")
		}
	}

	return nil
}
