package tasks

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"regexp"
	"socketAPI/common"
	"strconv"
	"strings"
	"time"
)

func AutoGift() {
	common.Log("AutoGift start")

	url := "https://wiki.biligame.com/nmd3/%E5%85%91%E6%8D%A2%E7%A0%81"
	res, err := http.Get(url)
	if err != nil {
		common.Log("crontab error", err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		common.Log("crontab error", err)
		common.Log("status code error: %d %s", res.StatusCode, res.Status)
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	tmpym := "2006年1月"
	tmpymd := "2006年1月2日"
	tmpymdhi := "2006年1月2日15:04"
	if err != nil {
		common.Log("crontab error", err)
		return
	}
	doc.Find(".main-line-wrap .resp-tabs .resp-tabs-container .resp-tab-content").Eq(0).Each(func(i int, s *goquery.Selection) {
		s.Find(".BOX-zt").Each(func(i int, s *goquery.Selection) {
			//c := s.Parent().Parent().Parent().Find("span").Text()
			//fmt.Printf("Review2 %d: %s\n", i, c)

			timeurl := "https://wiki.biligame.com/nmd3/" + s.Text()
			timeRes, err := http.Get(timeurl)
			if err != nil {
				common.Log("crontab error", err)
				return
			}
			defer timeRes.Body.Close()
			if timeRes.StatusCode != 200 {
				common.Log("crontab error", err)
				common.Log("status code error: %d %s", timeRes.StatusCode, timeRes.Status)
				return
			}
			resDoc, err := goquery.NewDocumentFromReader(timeRes.Body)
			if err != nil {
				common.Log("crontab NewDocumentFromReader error", err)
				return
			}
			t := resDoc.Find("#codeContent").Parent().Next().Text()
			if t == "" {
				timeurl := "https://wiki.biligame.com/nmd3/" + s.Text() + "（兑换码）"
				timeRes, err := http.Get(timeurl)
				if err != nil {
					common.Log("crontab 2 http.Get error", err)
					return
				}
				defer timeRes.Body.Close()
				if timeRes.StatusCode != 200 {
					common.Log("crontab2 error", err)
					common.Log("status code error: %d %s", timeRes.StatusCode, timeRes.Status)
					return
				}
				resDoc, err = goquery.NewDocumentFromReader(timeRes.Body)
				if err != nil {
					common.Log("crontab2 NewDocumentFromReader error", err)
					return
				}
				t = resDoc.Find("#codeContent").Parent().Next().Text()

			}
			tt := strings.Split(t, "-")

			for k, v := range tt {
				tt[k] = common.StringStrip(v)
				if strings.Contains(tt[k], "过期") {
					return
				}

				if strings.Contains(tt[k], "长期") {
					tt[k] = "9999年1月1日"
				}
				if strings.HasSuffix(tt[k], "24:00") {
					tt[k] = tt[k][:len(tt[k])-5]

				}

				if (!strings.Contains(tt[k], "年")) && k == 1 && tt[k] != "" {
					formatT, err := time.ParseInLocation(tmpymdhi, tt[0], time.Local)
					if err != nil {
						common.Log("crontab ParseInLocation 1 error", err)
						return
					}
					tt[1] = formatT.Format(tmpymd) + tt[1]
				}

				if strings.HasSuffix(tt[k], "月") {
					formatT, err := time.ParseInLocation(tmpym, tt[k], time.Local)
					if err != nil {
						common.Log("crontab ParseInLocation 2 error", err)
						return
					}
					tt[k] = formatT.AddDate(0, 1, -1).Format(tmpymd)
				}

			}

			if len(tt) == 0 {
				tt[0] = "9999年12月31日"
				tt[1] = "9999年12月31日"
			}
			if len(tt) == 1 {
				tt = append(tt, tt[0])
			}
			if len(tt) == 2 && tt[1] == "" {
				tt[1] = tt[0]
			}

			if strings.HasSuffix(tt[1], "日") {
				tt[1] = tt[1] + "23:59:59"
			}

			var autoCronGiftCodes []common.AutoCronGiftcode
			regQ, err := regexp.Compile("[年月]")
			if err != nil {
				common.Log("crontab error", err)
				return
			}
			regQr, err := regexp.Compile("[日]")
			if err != nil {
				common.Log("crontab error", err)
				return
			}
			regQt, err := regexp.Compile("[:]")
			if err != nil {
				common.Log("crontab error", err)
				return
			}
			tt[0] = regQ.ReplaceAllString(tt[0], "-")
			tt[1] = regQ.ReplaceAllString(tt[1], "-")
			tt[0] = regQr.ReplaceAllString(tt[0], " ")
			tt[1] = regQr.ReplaceAllString(tt[1], " ")
			tt[0] = regQt.ReplaceAllString(tt[0], ":")
			tt[1] = regQt.ReplaceAllString(tt[1], ":")

			err = common.Db.Select(&autoCronGiftCodes, "select * from autocrongiftcode where code = ?", s.Text())
			if err != nil {
				common.Log("crontab error", err)
				return
			}
			if len(autoCronGiftCodes) == 0 {
				var autoCronGiftCode common.AutoCronGiftcode
				autoCronGiftCode.Code = s.Text()
				autoCronGiftCode.Start = tt[0]
				autoCronGiftCode.End = tt[1]
				in, err := common.Db.NamedExec(`INSERT INTO autocrongiftcode (code, start, end) 
VALUES (:code, :start, :end)`, autoCronGiftCode)
				if err != nil {
					common.Log("crontab error", err)
					common.Log(in)
					return
				}
			} else {
				_, err = common.Db.Exec("update autocrongiftcode set `start`=? , `end`=? where  code=?",
					tt[0], tt[1], s.Text())
				if err != nil {
					common.Log("crontab update error", err)
					return
				}
			}

		})
	})

	common.Log("AutoGift end")

}

func GiftStart() {
	common.Log("GiftStart start")
	AutoGift()
	var cronuids []common.Cronuid
	var machines []common.Machines

	err := common.Db.Select(&cronuids, "select * from cronuid where exp_time > ? and del=0", time.Now().Unix())
	if err != nil {
		common.Log("crontab error", err)
		return
	}
	t := time.Now()
	newTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	err = common.Db.Select(&machines, "select * from machines where update_time > ?", newTime.Unix())
	if err != nil {
		common.Log("crontab error", err)
		return
	}

	for _, u := range cronuids {
		getGift(u.UserId)
	}
	for _, u := range machines {
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

	var crongifts []common.AutoCronGiftcode
	err = common.Db.Select(&crongifts, "select * from autocrongiftcode")

	if err != nil {
		common.Log("crontab error", err)
		return
	}

	for _, crongift := range crongifts {
		formatTs, err := time.ParseInLocation("2006-1-2 15:04:05", crongift.Start, time.Local)
		formatTe, err := time.ParseInLocation("2006-1-2 15:04:05", crongift.End, time.Local)
		if time.Now().Before(formatTs) || time.Now().After(formatTe) {
			continue
		}
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
		//if resData["code"] == float64(424) {
		//	_, err = common.Db.Exec(`update crongiftcode set del = 1 where id = ? and create_time < ?`, crongift.Id, time.Now().AddDate(0, 0, -7).Unix())
		//	if err != nil {
		//		common.Log("crontab 424 error", err)
		//		return
		//	}
		//}

		if resData["code"] == float64(425) || resData["code"] == float64(0) || resData["code"] == float64(417) {
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

		common.Log("gift get", uidString, resData, resData["code"], crongift.Code)
	}

}
