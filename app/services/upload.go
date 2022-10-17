package services

import (
	"errors"
	"socketAPI/common"
)

func UploadPic1(req map[string]string) (interface{}, error) {
	if _, ok := req["image"]; !ok {
		return nil, errors.New("image can not be empty")
	}
	if _, ok := req["prestr"]; !ok {
		return nil, errors.New("prestr can not be empty")
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
