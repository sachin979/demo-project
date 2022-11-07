package database

import (
"fmt"
"log"
"todo/config"
"gorm.io/driver/mysql"
"gorm.io/gorm"
gosql "github.com/go-sql-driver/mysql"
)

func New(conf *config.Conf) (*config.Dbs, error) {  
	cfg := &gosql.Config{
	  Net:                  "tcp",
	  Addr:                 fmt.Sprintf("%v:%v", conf.Db.Host, conf.Db.Port),
	  DBName:               conf.Db.DbName,
	  User:                 conf.Db.Username,
	  Passwd:               conf.Db.Password,
	  AllowNativePasswords: true,
	  ParseTime:            true,
	}
  
	db,err:= gorm.Open(mysql.Open(cfg.FormatDSN()), &gorm.Config{})
	

	if err != nil {
		log.Println("Connection failed")
		panic(err)
	}

	return &config.Dbs{
		DB:    db,
	}, nil
	
  }
 