package services

import (
	"errors"
	"regexp"
	"socketAPI/app/tesseract"
	"socketAPI/common"
	"strings"
	"time"
)

func getQuestionUseMD5(md5 string) (string, error) {
	var questions []common.QuestionsANDMd5
	err := common.Db.Select(&questions, "select * from question_md5 as qm LEFT JOIN questions as q ON qm.question = q.question where qm.md5=?", md5)
	if err != nil {
		return "", errors.New("get question MD5 error")
	}
	if len(questions) > 0 {
		return questions[0].Answer, nil
	}
	return "", errors.New("no md5")
}

func getQuestion(ques string) (string, error) {
	var questions []common.Questions
	err := common.Db.Select(&questions, "select * from questions where question=?", ques)
	if err != nil {
		return "", errors.New("get question error")
	}
	if len(questions) > 0 {
		return questions[0].Answer, nil
	}
	return "", errors.New("no question")
}

func UploadQuestion(req map[string]string) (interface{}, error) {
	if _, ok := req["question"]; !ok {
		return nil, errors.New("question can not be empty")
	}
	if _, ok := req["select1"]; !ok {
		return nil, errors.New("select1 can not be empty")
	}
	if _, ok := req["select2"]; !ok {
		return nil, errors.New("select2 can not be empty")
	}
	if _, ok := req["select3"]; !ok {
		return nil, errors.New("select3 can not be empty")
	}

	if _, ok := req["md5"]; !ok {
		return nil, errors.New("md5 can not be empty")
	}

	res := map[string]string{"answer": ""}

	answer, err := getQuestionUseMD5(req["md5"])
	if err != nil {
		if err.Error() != "no md5" {
			return nil, err
		}
	} else {
		res["answer"] = answer
		return res, nil
	}

	question, err := tesseract.GetWordOfPic(strings.Replace(req["question"], "_JH_", "+", -1))
	if err != nil {
		return nil, err
	}
	re, err := regexp.Compile("^第[\\s\\S]{1,5}问\\s?\\S?")
	ques := re.ReplaceAllString(question, "")

	if err != nil {
		return nil, err
	}
	answer, err = getQuestion(ques)
	if err != nil {
		if err.Error() != "no question" {
			return nil, err
		}
	} else {
		var insertQuestionMd5 common.QuestionMd5
		insertQuestionMd5.Question = ques
		insertQuestionMd5.Md5 = req["md5"]
		insertQuestionMd5.UpdateTime = int(time.Now().Unix())
		_, err = common.Db.NamedExec(`INSERT INTO question_md5 (question,md5,update_time)
VALUES (:question, :md5, :update_time)`, insertQuestionMd5)
		if err != nil {
			return nil, err
		}
		res["answer"] = answer
		return res, nil
	}

	select1, err := tesseract.GetWordOfPic(strings.Replace(req["select1"], "_JH_", "+", -1))
	if err != nil {
		return nil, err
	}
	select2, err := tesseract.GetWordOfPic(strings.Replace(req["select2"], "_JH_", "+", -1))
	if err != nil {
		return nil, err
	}
	select3, err := tesseract.GetWordOfPic(strings.Replace(req["select3"], "_JH_", "+", -1))
	if err != nil {
		return nil, err
	}

	var insertQuestion common.Questions
	var insertQuestionMd5 common.QuestionMd5

	insertQuestion.Question = ques
	insertQuestion.Select1 = select1
	insertQuestion.Select2 = select2
	insertQuestion.Select3 = select3
	insertQuestion.Answer = res["answer"]

	insertQuestionMd5.Question = ques
	insertQuestionMd5.Md5 = req["md5"]
	insertQuestionMd5.UpdateTime = int(time.Now().Unix())

	_, err = common.Db.NamedExec(`INSERT INTO questions (question, select1, select2,select3)
VALUES (:question, :select1, :select2, :select3)`, insertQuestion)
	if err != nil {
		return nil, err
	}
	_, err = common.Db.NamedExec(`INSERT INTO question_md5 (question,md5,update_time)
VALUES (:question, :md5, :update_time)`, insertQuestionMd5)
	if err != nil {
		return nil, err
	}
	return res, nil
}
