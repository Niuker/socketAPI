package services

import (
	"regexp"
	"socketAPI/common"
)

func Doit() {
	q2()
}

func q2() {
	var questions []common.Questions
	var questionsmd5 []common.QuestionMd5
	_ = common.Db.Select(&questions, "select * from questions")
	_ = common.Db.Select(&questionsmd5, "select * from question_md5")
	regAll, _ := regexp.Compile("[^0-9a-zA-Z\u4e00-\u9fa5]+")
	common.Log(222)

	for _, v := range questionsmd5 {
		common.Log(v.Question)
		common.Log(regAll.ReplaceAllString(v.Question, ""))
		_, err := common.Db.Exec("update question_md5 set question=? where question = ? ", regAll.ReplaceAllString(v.Question, ""), v.Question)
		if err != nil {
			common.Log(err)
		}
	}

	for _, vv := range questions {
		_, err := common.Db.Exec("update questions set question=?,select1=?,select2=?,select3=?  where question = ? ", regAll.ReplaceAllString(vv.Question, ""), regAll.ReplaceAllString(vv.Select1, ""), regAll.ReplaceAllString(vv.Select2, ""), regAll.ReplaceAllString(vv.Select3, ""), vv.Question)
		if err != nil {
			common.Log(err)
		}
	}

}

func q1() {
	var questions []common.Questions
	var questionsmd5 []common.QuestionMd5
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
