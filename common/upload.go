package common

import (
	"encoding/base64"
	"os"
	"strings"
)

func UploadByJson(json string, dir string, name string) error {
	json = strings.Replace(json, "_JH_", "+", -1)
	path := "./upload/" + dir
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	pic, err := base64.StdEncoding.DecodeString(json)
	if err != nil {
		return err
	}
	file := path + "/" + name

	err = os.WriteFile(file, pic, 0644)
	if err != nil {
		return err
	}
	return nil
}
