package services

import (
	"errors"
	"socketAPI/app/encr"
	"socketAPI/common"
	"socketAPI/config"
)

func PutAccount(req map[string]string) (interface{}, error) {
	if _, ok := req["machine_code"]; !ok {
		return nil, errors.New("machine_code不能为空")
	}
	if _, ok := req["account"]; !ok {
		return nil, errors.New("account不能为空")
	}
	machineCode, err := encr.ECBDecrypter(config.MyConfig.ENCR.Desckey, req["machine_code"])
	if machineCode == "" || err != nil {
		return nil, errors.New("本次machine_code解密失败")
	}
	var account common.UserConfigAccount
	var accounts []common.UserConfigAccount
	var aleradyAccounts []common.UserConfigAccount

	err = common.Db.Select(&aleradyAccounts, "select id from user_config_account where account = ? ", req["account"])
	if err != nil {
		return nil, err
	}
	if len(aleradyAccounts) >= 1 {
		return nil, errors.New("account已存在")
	}

	err = common.Db.Select(&accounts, "select id from user_config_account where machine_code = ? ", machineCode)
	if err != nil {
		return nil, err
	}
	if len(accounts) >= 1 {
		_, err = common.Db.Exec("DELETE FROM user_config_account WHERE id = ?", accounts[0].Id)

		if err != nil {
			return nil, err
		}

		_, err = common.Db.Exec("DELETE FROM user_config WHERE account_id = ?", accounts[0].Id)

		if err != nil {
			return nil, err
		}
	}

	account.Account = req["account"]
	account.MachineCode = machineCode
	_, err = common.Db.NamedExec(`INSERT INTO user_config_account (machine_code, account) 
VALUES (:machine_code, :account)`, account)

	if err != nil {
		return nil, err
	}

	return account, nil
}

func AddUserConfig(req map[string]string) (interface{}, error) {
	if _, ok := req["account"]; !ok {
		return nil, errors.New("account不能为空")
	}

	if _, ok := req["name"]; !ok {
		return nil, errors.New("name不能为空")
	}

	if _, ok := req["config"]; !ok {
		return nil, errors.New("config不能为空")
	}

	var accounts []common.UserConfigAccount
	err := common.Db.Select(&accounts, "select id from user_config_account where account = ? ", req["account"])
	if err != nil {
		return nil, err
	}
	if len(accounts) != 1 {
		return nil, errors.New("account异常")
	}

	var config common.UserConfig
	var configs []common.UserConfig

	err = common.Db.Select(&configs, "select `id`,`name` from user_config where account_id = ? and del = 0", accounts[0].Id)
	if err != nil {
		return nil, err
	}

	for _, v := range configs {
		if v.Name == req["name"] {
			return nil, errors.New("name已存在")
		}
	}
	if len(configs) >= 100 {
		return nil, errors.New("configs超过100个")
	}
	config.Config = req["config"]
	config.Name = req["name"]
	config.AccountId = accounts[0].Id
	_, err = common.Db.NamedExec(`INSERT INTO user_config (name, config, account_id) 
VALUES (:name, :config, :account_id)`, config)

	if err != nil {
		return nil, err
	}

	allConfigs, err := getConfigsByAccount(req["account"])
	if err != nil {
		return nil, err
	}
	return allConfigs, nil
}

func getConfigsByAccount(account string) ([]common.UserConfig, error) {
	var accounts []common.UserConfigAccount

	err := common.Db.Select(&accounts, "select id from user_config_account where account = ? ", account)
	if err != nil {
		return nil, err
	}
	if len(accounts) != 1 {
		return nil, errors.New("account异常")
	}

	var configs []common.UserConfig

	err = common.Db.Select(&configs, "select id,name from user_config where account_id = ? and del=0", accounts[0].Id)
	if err != nil {
		return nil, err
	}

	return configs, nil
}

func GetConfigList(req map[string]string) (interface{}, error) {
	if _, ok := req["account"]; !ok {
		return nil, errors.New("account不能为空")
	}
	allConfigs, err := getConfigsByAccount(req["account"])
	if err != nil {
		return nil, err
	}
	return allConfigs, nil
}

func GetConfig(req map[string]string) (interface{}, error) {
	if _, ok := req["id"]; !ok {
		return nil, errors.New("id不能为空")
	}
	var configs []common.UserConfig
	err := common.Db.Select(&configs, "select * from user_config where id = ? and del=0", req["id"])
	if err != nil {
		return nil, err
	}
	if len(configs) != 1 {
		return nil, errors.New("configs数据异常")
	}

	return configs[0], nil
}

func DelConfig(req map[string]string) (interface{}, error) {
	if _, ok := req["id"]; !ok {
		return nil, errors.New("id不能为空")
	}
	_, err := common.Db.Exec("update user_config set `del`=1 where id=?", req["id"])
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"status": "success",
	}, nil
}
