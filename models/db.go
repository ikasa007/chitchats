package models

import (
	"database/sql"
	"fmt"
	log "github.com/tim/chitchat/common/log"
	. "github.com/tim/chitchat/config"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

var Db *sql.DB

func init(){
	var err error
	// todo move the mysql connection into config file later
	//config := LoadConfig()
	driver := ViperConfig.Db.Driver
	// source := user:password@address/database?charset&parseTime
	source := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=%t",
		ViperConfig.Db.User,ViperConfig.Db.Passwd,ViperConfig.Db.Addr,strconv.Itoa(ViperConfig.Db.Port),ViperConfig.Db.Database,
		ViperConfig.Db.Charset,ViperConfig.Db.ParseTime)
	//log.Info(driver+source)
	Db, err = sql.Open(driver,source)
	if err != nil {
		log.Crash(err)
	}
	return

}


