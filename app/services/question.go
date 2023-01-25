package services

import (
	"errors"
	"socketAPI/app/tesseract"
	"socketAPI/common"
	"strings"
)

func getQuestionUseMD5(md5 string) (string, error) {
	var questions []common.Questions
	err := common.Db.Select(&questions, "select * from questions where md5=?", md5)
	if err != nil {
		return "", errors.New("get question MD5 error")
	}
	if len(questions) > 0 {
		return questions[0].Answer, nil
	}
	return "", nil
}

func getQuestion(ques string) (string, error) {
	var questions []common.Questions
	err := common.Db.Select(&questions, "select * from questions where question=?", ques)
	if err != nil {
		return "", errors.New("get question MD5 error")
	}
	if len(questions) > 0 {
		return questions[0].Answer, nil
	}
	return "", nil
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
		return nil, err
	}
	if answer != "" {
		res["answer"] = answer
		return res, nil
	}

	question, err := tesseract.GetWordOfPic(strings.Replace(req["question"], "_JH_", "+", -1))
	if err != nil {
		return nil, err
	}
	answer, err = getQuestion(question)
	if err != nil {
		return nil, err
	}
	if answer != "" {
		res["answer"] = answer
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
	insertQuestion.Question = question
	insertQuestion.Select1 = select1
	insertQuestion.Select2 = select2
	insertQuestion.Select3 = select3
	insertQuestion.Md5 = req["md5"]
	insertQuestion.Answer = res["answer"]

	_, err = common.Db.NamedExec(`INSERT INTO questions (question, select1, select2,select3,md5)
VALUES (:question, :select1, :select2, :select3, :md5)`, insertQuestion)
	if err != nil {
		return nil, err
	}
	return res, nil
}
