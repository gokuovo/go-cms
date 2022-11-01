package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/pelletier/go-toml"
	"log"
)

var CMDB *gorm.DB
var InDB *gorm.DB
var DW *gorm.DB

// InitDB 数据库链接
func InitDB(dbname, env string) *gorm.DB {
	if len(dbname) == 0 {
		dbname = "go_cms"
	}
	config, _ := toml.LoadFile("./conf/" + env + "/db.toml") //加载toml文件
	mysqlName := config.Get(dbname + ".db_name").(string)
	mysqlUser := config.Get(dbname + ".db_user").(string)
	mysqlPwd := config.Get(dbname + ".db_pwd").(string)
	mysqlHost := config.Get(dbname + ".db_host").(string)
	mysqlPort := config.Get(dbname + ".db_port").(string)
	mysqlCharset := config.Get(dbname + ".db_charset").(string) //读取key对应的值.括号为指定数据类型

	var dataSource string
	// parseTime=true 表示自动解析为时间
	dataSource = mysqlUser + ":" + mysqlPwd + "@tcp(" + mysqlHost + ":" + mysqlPort + ")/" + mysqlName + "?charset=" + mysqlCharset + "&parseTime=true"
	Db, err := gorm.Open("mysql", dataSource)
	if err != nil {
		Db.Close()
		log.Println("数据库加载失败", err)
	}

	// 设置数据库连接池空闲连接数
	Db.DB().SetConnMaxIdleTime(50)

	// 打开连接
	Db.DB().SetMaxOpenConns(100)

	// 表明禁用后缀加s
	Db.SingularTable(true)

	return Db
}
func CreateDB(config string) {
	CMDB = InitDB("go_cms", config)
	//InDB = InitDB("ad_novel", config)
	//DW = InitDB("adsense_win", config)
}
