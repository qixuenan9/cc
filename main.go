package settings

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File
	Username,Password,Host,DataBase,HttpPort,TablePrefix,Type  string
	PageSize int
	Bb interface{}
	ReadTimeOut time.Duration
)
func init(){

	var err error
	if Cfg,err = ini.Load("conf/app.ini");err != nil {
		if err != nil{
			log.Fatalf("Fail to parse 'conf/app.ini':%v",err)
		}
	}
	Type = Cfg.Section("database").Key("TYPE" ).String()
	Username = Cfg.Section("database").Key("USER" ).String()
	Password = Cfg.Section("database").Key("PASSWORD" ).String()
	Host = Cfg.Section("database").Key("HOST" ).String()
	DataBase = Cfg.Section("database").Key("NAME" ).String()
	TablePrefix = Cfg.Section("database").Key("TABLE_PREFIX").String()

	HttpPort = Cfg.Section("server").Key("HTTP_PORT").String()
	ReadTimeOut = time.Duration(Cfg.Section("server").Key("READ_TIMEOUT" ).MustInt(0))  *time.Second


	//PageSize = Cfg.Section("app").Key("PAGE_SIZE" ).Int()



}
