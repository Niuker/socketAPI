package common

import (
	"WebsocketDemo/app/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Machines struct {
	Id          int    `db:"id" json:"id"`
	MachineCode string `db:"machine_code" json:"machine_code"`
	UserId      int    `db:"user_id" json:"user_id"`
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
	Id             int `db:"id" json:"id"`
	UserId         int `db:"user_id" json:"userId"`
	MissionFieldId int `db:"mission_field_id" json:"mission_field_id"`
	Value          int `db:"value" json:"value"`
	UpdateTime     int `db:"update_time" json:"update_time"`
	Date           int `db:"date" json:"date"`
}

type TimerField struct {
	Id      int    `db:"id" json:"id"`
	Name    string `db:"name" json:"name"`
	Default int    `db:"default" json:"default"`
}
type Timers struct {
	Id           int `db:"id" json:"id"`
	UserId       int `db:"user_id" json:"user_id"`
	TimerFieldId int `db:"timer_field_id" json:"timer_field_id"`
	Value        int `db:"value" json:"value"`
	UpdateTime   int `db:"update_time" json:"update_time"`
}

type MissionsANDMissionField struct {
	Missions
	MissionField
}
type TimersANDTimerField struct {
	Timers
	TimerField
}

var Db *sqlx.DB

func InitDB() {
	database, err := sqlx.Open("mysql", config.MyConfig.DB.User+":"+config.MyConfig.DB.Password+
		"@tcp(127.0.0.1:"+config.MyConfig.DB.Port+")/script")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
	Db.SetMaxIdleConns(1)
	Db.SetMaxOpenConns(1)
	//defer Db.Close()
}
