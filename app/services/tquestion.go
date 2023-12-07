package services

//func UploadTQuestion(req map[string]string) (interface{}, error) {
//	return nil, nil
//}

import (
	"context"
	"errors"
	"fmt"
	t_common "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	t_errors "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	ocr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ocr/v20181119"
	"regexp"
	"socketAPI/common"
	"strconv"
	"strings"
	"time"
)

func getQuestionAndAnswer(ques string, s1 string, s2 string, s3 string) (string, error) {
	var questions []common.Questions
	err := common.Db.Select(&questions, "select * from questions where question=? and select1=? and select2=? and select3=?", ques, s1, s2, s3)
	if err != nil {
		return "", errors.New("get t question error")
	}
	if len(questions) > 0 {
		return questions[0].Answer, nil
	}
	return "", errors.New("no question")
}

func UploadTQuestion(req map[string]string) (interface{}, error) {
	if _, ok := req["pic"]; !ok {
		return nil, errors.New("pic can not be empty")
	}
	if _, ok := req["id"]; !ok {
		return nil, errors.New("id can not be empty")
	}
	if _, ok := req["key"]; !ok {
		return nil, errors.New("key can not be empty")
	}
	if _, ok := req["md5"]; !ok {
		return nil, errors.New("md5 can not be empty")
	}

	timestampNano := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)

	pic := strings.Replace(req["pic"], "_JH_", "+", -1)
	credential := t_common.NewCredential(req["id"], req["key"])

	client, err := ocr.NewClient(credential, regions.Beijing, profile.NewClientProfile())
	if err != nil {
		return nil, err
	}

	request := ocr.NewGeneralBasicOCRRequest()
	request.ImageBase64 = &pic
	var ctx context.Context
	response, err := client.GeneralBasicOCRWithContext(ctx, request)

	if _, ok := err.(*t_errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	l := len(response.Response.TextDetections)
	if l < 4 {
		return nil, errors.New("pic res text < 4")
	}
	text := make([]string, 4)
	regAll, err := regexp.Compile("[^0-9a-zA-Z\u4e00-\u9fa5]+")
	if err != nil {
		return nil, err
	}
	regQ1, err := regexp.Compile("^第[\\s\\S]{0,5}问")
	if err != nil {
		return nil, err
	}
	regQL, err := regexp.Compile("[^\u4e00-\u9fa5]+$")
	if err != nil {
		return nil, err
	}
	regS1, err := regexp.Compile("^A")
	if err != nil {
		return nil, err
	}
	regS2, err := regexp.Compile("^(B|8)")
	if err != nil {
		return nil, err
	}
	regS3, err := regexp.Compile("^C")
	if err != nil {
		return nil, err
	}

	var haveA bool = false
	for _, v := range response.Response.TextDetections {
		t := regAll.ReplaceAllString(common.StringStrip(*v.DetectedText), "")
		common.Log(t)
		if regS1.Match([]byte(t)) {
			text[1] = t
			haveA = true
			continue
		}
		if !haveA {
			text[0] = text[0] + t
		}
		if regS2.Match([]byte(t)) {
			text[2] = t
			continue
		}
		if regS3.Match([]byte(t)) {
			text[3] = t
			break
		}
	}

	for _, vv := range text {
		if vv == "" {
			return nil, errors.New("reg error")
		}
	}

	text[0] = regQ1.ReplaceAllString(text[0], "")
	text[0] = regQL.ReplaceAllString(text[0], "")

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

	answer, err = getQuestionAndAnswer(text[0], text[1], text[2], text[3])
	if err != nil {
		if err.Error() != "no question" {
			return nil, err
		}
	} else {
		var insertQuestionMd5 common.QuestionMd5
		insertQuestionMd5.Question = text[0]
		insertQuestionMd5.Md5 = req["md5"]
		insertQuestionMd5.UpdateTime = int(time.Now().Unix())

		_, err = common.Db.NamedExec(`INSERT INTO question_md5 (question,md5,update_time)
VALUES (:question, :md5,:update_time)`, insertQuestionMd5)
		if err != nil {
			return nil, err
		}
		res["answer"] = answer
		return res, nil
	}

	common.UploadByJson(pic, "question", time.Now().Format("2006-01-02 15:04:05")+"nano"+timestampNano+".png")

	var insertQuestion common.Questions
	var insertQuestionMd5 common.QuestionMd5

	insertQuestion.Question = text[0]
	insertQuestion.Select1 = text[1]
	insertQuestion.Select2 = text[2]
	insertQuestion.Select3 = text[3]
	insertQuestion.Answer = ""

	insertQuestionMd5.Question = text[0]
	insertQuestionMd5.Md5 = req["md5"]
	insertQuestionMd5.UpdateTime = int(time.Now().Unix())

	_, err = common.Db.NamedExec(`INSERT INTO questions (question, select1, select2,select3)
VALUES (:question, :select1, :select2, :select3)`, insertQuestion)
	if err != nil {
		return nil, err
	}
	_, err = common.Db.NamedExec(`INSERT INTO question_md5 (question,md5,update_time)
VALUES (:question, :md5,:update_time)`, insertQuestionMd5)
	if err != nil {
		return nil, err
	}
	return res, nil

}
