package services

import (
	"errors"
	"socketAPI/app/encr"
	"socketAPI/common"
	"socketAPI/config"
	"strconv"
)

func SetMachines(req map[string]string) (interface{}, error) {
	if _, ok := req["user_id"]; !ok {
		return nil, errors.New("user不能为空")
	}
	if _, ok := req["machine_code"]; !ok {
		return nil, errors.New("machine_code不能为空")
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
	err = common.Db.Select(&machines, "select id from machines where machine_code = ? and user_id = ?", req["machine_code"], id)
	if err != nil {
		return nil, err
	}
	machine.MachineCode = req["machine_code"]
	machine.UserId = id
	if len(machines) < 1 {
		_, err = common.Db.NamedExec(`INSERT INTO machines (machine_code, user_id) 
VALUES (:machine_code, :user_id)`, machine)
	}

	if err != nil {
		return nil, err
	}

	return machine, nil
}

func GetMachines(req map[string]string) (interface{}, error) {
	if _, ok := req["machine_code"]; !ok {
		return nil, errors.New("machine_code不能为空")
	}

	var machines []common.Machines

	err := common.Db.Select(&machines, "select * from machines where machine_code = ? ", req["machine_code"])
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
