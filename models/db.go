package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"net/url"
	"os"
	"time"
	"github.com/qor/media/asset_manager"
)

var DB *gorm.DB

func SyncDB() {
	createDB()
	Connect()
	DB.
		Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
		&User{},
		&Article{},
		&Category{},
		&Menu{},
		&asset_manager.AssetManager{},
	)
}

/**
数据库链接
*/
func Connect() {

	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=%s&parseTime=true",
		db_user,
		db_pass,
		db_host,
		db_port,
		db_name,
		url.QueryEscape("Asia/Shanghai"))

	var err error

	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Print("master detabase connect error:", err)
		os.Exit(0)
	}

	DB.SingularTable(true)
	DB.DB().SetMaxOpenConns(2000)
	DB.DB().SetMaxIdleConns(200)
	DB.DB().SetConnMaxLifetime(1 * time.Second)
	if beego.AppConfig.String("runmode") == "dev" {
		DB = DB.Debug()
	}
}

func createDB() {

	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&loc=%s&parseTime=true", db_user, db_pass, db_host, db_port, url.QueryEscape("Asia/Shanghai"))
	sqlstring := fmt.Sprintf("CREATE DATABASE  if not exists `%s` CHARSET utf8mb4 COLLATE utf8mb4_general_ci", db_name)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	r, err := db.Exec(sqlstring)
	if err != nil {
		log.Println(err)
		log.Println(r)
	} else {
		log.Println("Database ", db_name, " created")
	}
	defer db.Close()

}
