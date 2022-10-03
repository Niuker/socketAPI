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
		var cronuidgifts []common.CronUidgift
		err = common.Db.Select(&cronuidgifts, "select * from cronuidgift where  code_id=? and user_id=?", crongift.Id, uidString)
		if err != nil {
			common.Log("crontab error", err)
			return
		}
		if len(cronuidgifts) > 0 {
			continue
		}

		codeParams := map[string]string{"uid": uidString, "code": crongift.Code, "token": token}
		codeBody, err := common.HttpGet(codeUrl, codeParams)
		if err != nil {
			common.Log("crontab error", err)
			return
		}
		var resData map[string]interface{}
		err = json.Unmarshal(codeBody, &resData)
		if err != nil {
			common.Log("crontab error", err)
			return
		}
		if _, ok := resData["code"]; !ok {
			common.Log("crontab error data", err)
			return
		}

		if resData["code"] == float64(425) || resData["code"] == float64(0) {
			var CronUidgift common.CronUidgift

			CronUidgift.CodeId = crongift.Id
			CronUidgift.UserId = uid
			_, err = common.Db.NamedExec(`INSERT INTO cronuidgift (code_id, user_id) 
VALUES (:code_id, :user_id)`, CronUidgift)
			if err != nil {
				common.Log("crontab error", err)
				return
			}
		}

		common.Log("gift get", uidString, resData, resData["code"])
	}

}
