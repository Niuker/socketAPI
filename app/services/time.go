package services

import (
	"errors"
	"strconv"
	"time"
)

func GetTimeStamp(req map[string]string) (interface{}, error) {
	currentTime := time.Now()
	milliseconds := currentTime.UnixNano() / int64(time.Millisecond)
	return milliseconds, nil
}

func GetSystemTimers(req map[string]string) (interface{}, error) {

	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return nil, err
	}
	if _, ok := req["timestamp"]; !ok {
		return nil, errors.New("timestamp不能为空")
	}
	timestamp, err := strconv.ParseInt(req["timestamp"], 10, 64)
	if err != nil {
		return nil, errors.New("timestamp  ParseInt error")
	}
	//now := time.Now().In(location)
	now := time.Unix(timestamp, 0).In(location)

	previous5AM := time.Date(now.Year(), now.Month(), now.Day(), 5, 0, 0, 0, location).Unix()
	if now.Hour() < 5 {
		previous5AM = previous5AM - 86400
	}

	var previousSunday time.Time
	if now.Weekday() == 0 && now.Hour() < 5 {
		previousSunday = now.AddDate(0, 0, -7).Add(-time.Duration(now.Hour()-5) * time.Hour)
	} else {
		previousSunday = now.AddDate(0, 0, -int((now.Weekday())+7)%7).Add(-time.Duration(now.Hour()-5) * time.Hour)
	}
	previousSunday5AM := time.Date(previousSunday.Year(), previousSunday.Month(), previousSunday.Day(), 5, 0, 0, 0, location).Unix()

	//return map[string]string{
	//	"previous5AM":       time.Unix(previous5AM, 0).Format("2006-01-02 15:04:05"),
	//	"previousSunday5AM": time.Unix(previousSunday5AM, 0).Format("2006-01-02 15:04:05"),
	//}, nil
	return map[string]int64{
		"previous5AM":       previous5AM,
		"previousSunday5AM": previousSunday5AM,
	}, nil
}
