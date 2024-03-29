package services

import (
	"errors"
	"socketAPI/common"
)

func UploadPic1(req map[string]string) (interface{}, error) {
	if _, ok := req["version"]; !ok {
		return nil, nil
	}
	if _, ok := req["image"]; !ok {
		return nil, errors.New("image can not be empty")
	}
	if _, ok := req["prestr"]; !ok {
		return nil, errors.New("prestr can not be empty")
	}
	var versions []common.CronVersion
	err := common.Db.Select(&versions, "select * from cronversion where name = ?", "upload")
	if err != nil {
		return nil, err
	}
	if len(versions) == 0 {
		return nil, errors.New("upload version can not be empty")
	}

	v := false
	for _, vv := range versions {
		if vv.Version == req["version"] {
			v = true
			break
		}
	}

	if !v {
		return nil, nil
	}

	return nil, common.UploadByJson(req["image"], "pic1", req["prestr"]+".png")
}

func UploadPic2(req map[string]string) (interface{}, error) {
	if _, ok := req["image"]; !ok {
		return nil, errors.New("image can not be empty")
	}
	if _, ok := req["prestr"]; !ok {
		return nil, errors.New("prestr can not be empty")
	}
	return nil, common.UploadByJson(req["image"], "pic2", req["prestr"]+".png")
}
