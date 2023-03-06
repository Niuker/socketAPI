package common

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"socketAPI/config"
)

type Machines struct {
	Id          int    `db:"id" json:"id"`
	MachineCode string `db:"machine_code" json:"machine_code"`
	UserId      int    `db:"user_id" json:"user_id"`
	UpdateTime  int    `db:"update_time" json:"update_time"`
	Mid         string `db:"mid" json:"mid"`
}

type Message struct {
	Id      int    `db:"id" json:"id"`
	Title   string `db:"title" json:"title"`
	Content string `db:"content" json:"content"`
	Time    string `db:"time" json:"time"`
	Imei    string `db:"imei" json:"imei"`
}

type MissionField struct {
	Id      int    `db:"id" json:"id"`
	Name    string `db:"name" json:"name"`
	Default int    `db:"default" json:"default"`
	Isday   int    `db:"isday" json:"isday"`
}
type Missions struct {
	Id             int    `db:"id" json:"id"`
	UserId         int    `db:"user_id" json:"userId"`
	MissionFieldId int    `db:"mission_field_id" json:"mission_field_id"`
	Value          int    `db:"value" json:"value"`
	UpdateTime     int    `db:"update_time" json:"update_time"`
	Date           int    `db:"date" json:"date"`
	MachineCode    string `db:"machine_code" json:"machine_code"`
}

type TimerField struct {
	Id      int    `db:"id" json:"id"`
	Name    string `db:"name" json:"name"`
	Default int    `db:"default" json:"default"`
}
type Timers struct {
	Id           int    `db:"id" json:"id"`
	UserId       int    `db:"user_id" json:"user_id"`
	TimerFieldId int    `db:"timer_field_id" json:"timer_field_id"`
	Value        int    `db:"value" json:"value"`
	UpdateTime   int    `db:"update_time" json:"update_time"`
	MachineCode  string `db:"machine_code" json:"machine_code"`
}
type Questions struct {
	Id       int    `db:"id" json:"id"`
	Question string `db:"question" json:"question"`
	Select1  string `db:"select1" json:"select1"`
	Select2  string `db:"select2" json:"select2"`
	Select3  string `db:"select3" json:"select3"`
	Answer   string `db:"answer" json:"answer"`
}

type QuestionMd5 struct {
	Id         int    `db:"id" json:"id"`
	Question   string `db:"question" json:"question"`
	Md5        string `db:"md5" json:"md5"`
	UpdateTime int    `db:"update_time" json:"update_time"`
}

type Cronuid struct {
	Id      int    `db:"id" json:"id"`
	UserId  int    `db:"user_id" json:"user_id"`
	ExpTime int    `db:"exp_time" json:"exp_time"`
	Source  int    `db:"source" json:"source"`
	Del     int    `db:"del" json:"del"`
	Name    string `db:"name" json:"name"`
}

type CronGiftcode struct {
	Id         int    `db:"id" json:"id"`
	Code       string `db:"code" json:"code"`
	Del        int    `db:"del" json:"del"`
	CreateTime int    `db:"create_time" json:"create_time"`
}

type AutoCronGiftcode struct {
	Id    int    `db:"id" json:"id"`
	Code  string `db:"code" json:"code"`
	Start string `db:"start" json:"start"`
	End   string `db:"end" json:"end"`
}

type CronUidgift struct {
	Id     int `db:"id" json:"id"`
	CodeId int `db:"code_id" json:"code_id"`
	UserId int `db:"user_id" json:"user_id"`
}

type UserConfigAccount struct {
	Id   int    `db:"id" json:"id"`
	User string `db:"user" json:"user"`
	Pass string `db:"pass" json:"pass"`
}

type UserConfig struct {
	Id     int    `db:"id" json:"id"`
	UserId int    `db:"user_id" json:"user_id"`
	Config string `db:"config" json:"config"`
	Name   string `db:"name" json:"name"`
	Del    string `db:"del" json:"del"`
}

type Version struct {
	Id      int    `db:"id" json:"id"`
	Name    string `db:"name" json:"name"`
	Version string `db:"version" json:"version"`
}

type MissionsANDMissionField struct {
	Missions
	MissionField
}
type TimersANDTimerField struct {
	Timers
	TimerField
}

type QuestionsANDMd5 struct {
	Questions
	Md5        string `db:"md5" json:"md5"`
	UpdateTime int    `db:"update_time" json:"update_time"`
}

type MachinesANDMid struct {
	Machines
	mid string
}

var Db *sqlx.DB

func InitDB() {
	database, err := sqlx.Open("mysql", config.MyConfig.DB.User+":"+config.MyConfig.DB.Password+
		"@tcp(127.0.0.1:"+config.MyConfig.DB.Port+")/script2")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
	Db.SetMaxIdleConns(30)
	Db.SetMaxOpenConns(30)
	//defer Db.Close()
}
