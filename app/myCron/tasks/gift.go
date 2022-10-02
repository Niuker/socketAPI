package tasks

import (
	"encoding/json"
	"socketAPI/common"
	"strconv"
	"time"
)

func GiftStart() {
	common.Log("GiftStart start")

	var cronuids []common.Cronuid

	err := common.Db.Select(&cronuids, "select * from cronuid where exp_time > ? and del=0", time.Now().Unix())

	if err != nil {
		common.Log("crontab error", err)
		return
	}

	for _, u := range cronuids {
		getGift(u.UserId)
	}
	common.Log("GiftStart end")

}

func getGift(uid int) {
	uidString := strconv.Itoa(uid)
	sourceUrl := "https://cn.yhsvc.pandadastudio.com/captcha/apis/v1/apps/ninja3/versions/v1/captchas"
	sourceParams := map[string]string{}
	sourceBody, err := common.HttpGet(sourceUrl, sourceParams)
	if err != nil {
		common.Log("crontab error", err)
		return
	}
	var data map[string]interface{}
	err = json.Unmarshal(sourceBody, &data)
	if err != nil {
		common.Log("crontab error", err)
		return
	}

	if _, ok := data["data"]; !ok {
		common.Log("crontab error data", err)
		return
	}

	resdatadata := data["data"].(map[string]interface{})
	if _, ok := resdatadata["token"]; !ok {
		common.Log("crontab error token", err)
		return
	}

	token := resdatadata["token"].(string)

	codeUrl := "https://statistics.pandadastudio.com/player/giftCode"

	var crongifts []common.CronGiftcode
	err = common.Db.Select(&crongifts, "select * from crongiftcode where  del=0")

	if err != nil {
		common.Log("crontab error", err)
		return
	}

	for _, crongift := range crongifts {
		codeParams := map[string]string{"uid": uidString, "code": crongift.Code, "token": token}
		codeBody, err := common.HttpGet(codeUrl, codeParams)
		if err != nil {
			common.Log("crontab error", err)
			return
		}
		var data map[string]interface{}
		err = json.Unmarshal(codeBody, &data)
		if err != nil {
			common.Log("crontab error", err)
			return
		}
		if _, ok := data["code"]; !ok {
			common.Log("crontab error data", err)
			return
		}

		common.Log("gift get", uidString, data)
	}

}
