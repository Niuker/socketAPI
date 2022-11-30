package services

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"socketAPI/common"
)

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func Account(req map[string]string) (interface{}, error) {
	if _, ok := req["user"]; !ok {
		return nil, errors.New("user不能为空")
	}
	if _, ok := req["pass"]; !ok {
		return nil, errors.New("pass不能为空")
	}

	pass := GetMd5String(req["pass"])

	var accounts []common.UserConfigAccount
	var account common.UserConfigAccount
	var configs []common.UserConfig

	err := common.Db.Select(&accounts, "select id,pass from user_config_account where user = ?", req["user"])
	if err != nil {
		return nil, err
	}

	if len(accounts) == 0 {
		account.User = req["user"]
		account.Pass = pass
		_, err = common.Db.NamedExec(`INSERT INTO user_config_account (user, pass) 
VALUES (:user, :pass)`, account)
		if err != nil {
			return nil, err
		}

		return configs, nil
	}

	if len(accounts) == 1 {
		if accounts[0].Pass != pass {
			return nil, errors.New("密码错误")
		}

		return getConfigsByUserID(accounts[0].Id)
	}

	return nil, errors.New("account异常")
}

func getConfigsByUserID(userID int) ([]common.UserConfig, error) {
	var configs []common.UserConfig
	err := common.Db.Select(&configs, "select id,name from user_config where user_id = ? and del=0", userID)
	if err != nil {
		return nil, err
	}

	return configs, nil
}

func verifyUser(user string, pass string) ([]common.UserConfigAccount, error) {
	pass = GetMd5String(pass)

	var accounts []common.UserConfigAccount
	err := common.Db.Select(&accounts, "select id from user_config_account where user = ? and pass = ? ", user, pass)
	if err != nil {
		return accounts, err
	}
	if len(accounts) != 1 {
		return accounts, errors.New("account异常")
	}
	return accounts, nil
}

func AddUserConfig(req map[string]string) (interface{}, error) {
	if _, ok := req["user"]; !ok {
		return nil, errors.New("user不能为空")
	}
	if _, ok := req["pass"]; !ok {
		return nil, errors.New("pass不能为空")
	}

	if _, ok := req["name"]; !ok {
		return nil, errors.New("name不能为空")
	}

	if _, ok := req["config"]; !ok {
		return nil, errors.New("config不能为空")
	}

	var config common.UserConfig
	var configs []common.UserConfig

	accounts, err := verifyUser(req["user"], req["pass"])

	if err != nil {
		return nil, err
	}

	err = common.Db.Select(&configs, "select `id`,`name` from user_config where user_id = ? and del = 0", accounts[0].Id)
	if err != nil {
		return nil, err
	}

	for _, v := range configs {
		if v.Name == req["name"] {
			return nil, errors.New("name已存在")
		}
	}
	if len(configs) >= 1000 {
		return nil, errors.New("configs超过1000个")
	}
	config.Config = req["config"]
	config.Name = req["name"]
	config.UserId = accounts[0].Id
	_, err = common.Db.NamedExec(`INSERT INTO user_config (name, config, user_id) 
VALUES (:name, :config, :user_id)`, config)

	if err != nil {
		return nil, err
	}

	allConfigs, err := getConfigsByUserID(accounts[0].Id)
	if err != nil {
		return nil, err
	}
	return allConfigs, nil
}

func GetConfig(req map[string]string) (interface{}, error) {
	if _, ok := req["id"]; !ok {
		return nil, errors.New("id不能为空")
	}
	if _, ok := req["user"]; !ok {
		return nil, errors.New("user不能为空")
	}
	if _, ok := req["pass"]; !ok {
		return nil, errors.New("pass不能为空")
	}
	accounts, err := verifyUser(req["user"], req["pass"])
	if err != nil {
		return nil, err
	}
	var configs []common.UserConfig
	err = common.Db.Select(&configs, "select * from user_config where id = ? and del=0 and user_id = ?", req["id"], accounts[0].Id)
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
	if _, ok := req["user"]; !ok {
		return nil, errors.New("user不能为空")
	}
	if _, ok := req["pass"]; !ok {
		return nil, errors.New("pass不能为空")
	}
	accounts, err := verifyUser(req["user"], req["pass"])
	if err != nil {
		return nil, err
	}
	_, err = common.Db.Exec("update user_config set `del`=1 where id=? and user_id = ?", req["id"], accounts[0].Id)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"status": "success",
	}, nil
}
