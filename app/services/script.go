package services

import (
	"regexp"
	"socketAPI/common"
)

func Doit() {
	var questions []common.Questions
	var questionsmd5 []common.QuestionsANDMd5
	_ = common.Db.Select(&questions, "select * from questions")
	_ = common.Db.Select(&questionsmd5, "select * from question_md5")
	regQL, _ := regexp.Compile(" [\u4E00-\u9FA5]$")

	for _, v := range questionsmd5 {
		newq := ""
		if !regQL.Match([]byte(v.Question)) {
			if last := len(v.Question) - 1; last >= 0 {
				newq = v.Question[:last]
			}
		}
		//_, _ = common.Db.Exec("update question_md5 set question=? where question = ? ", newq, v.Question)

		common.Log(newq)
	}
	for _, vv := range questions {
		newqq := ""
		if !regQL.Match([]byte(vv.Question)) {
			if last := len(vv.Question) - 1; last >= 0 {
				newqq = vv.Question[:last]
			}
		}
		common.Log(newqq)
		_, _ = common.Db.Exec("update questions set question=? where question = ? ", newqq, vv.Question)

	}
}
