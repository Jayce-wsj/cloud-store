package db

import (
	"database/sql"
	"fmt"
	"log"
	"workspace/cloud-pan/config"

	_ "github.com/go-sql-driver/mysql"
)

// dbConfig 数据库配置结构
type dbConfig struct {
	Host          string `yaml:"host"`           //数据库地址
	Port          int    `yaml:"port"`           //数据库端口
	Dbname        string `yaml:"dbname"`         //数据库库名
	User          string `yaml:"user"`           //数据库用户名
	Pass          string `yaml:"pass" spew:"-"`  //数据库密码
	Charset       string `yaml:"charset"`        //数据库字符集
	MaxIdle       int    `yaml:"max_idle"`       //最大闲置连接数
	MaxConnection int    `yaml:"max_conncetion"` //数据库最大连接数
}

// connMySQL 连接MySQL数据库
func connMySQL(conf *dbConfig) error {
	var err error
	DefaultDB = &Database{}
	DefaultDB.DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", conf.User, conf.Pass, conf.Host, conf.Port, conf.Dbname, conf.Charset))
	if err != nil {
		return err
	}

	if conf.MaxIdle > 0 {
		DefaultDB.DB.SetMaxIdleConns(conf.MaxIdle)
		DefaultDB.DB.SetConnMaxLifetime(-1)
	}
	if conf.MaxConnection > 0 {
		DefaultDB.DB.SetMaxOpenConns(conf.MaxConnection)
	}

	err = DefaultDB.DB.Ping()
	if err != nil {
		return err
	}
	return nil
}

func init() {
	err := connMySQL(&dbConfig{
		Host:          config.MysqlHost,
		Port:          config.MysqlPort,
		Dbname:        config.MysqlDB,
		User:          config.MysqlUser,
		Pass:          config.MysqlPwd,
		Charset:       "utf8",
		MaxIdle:       10,
		MaxConnection: 10,
	})
	if err != nil {
		panic(err)
	} else {
		log.Print("数据库连接成功")
	}
}

/*func main() {
	_, err := DefaultDB.Exec(`
CREATE TABLE IF NOT EXISTS users (
  id int unsigned NOT NULL AUTO_INCREMENT,
  name varchar(36) NOT NULL,
  old int unsigned NOT NULL,
  pwd binary(16) NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
`)
	if err != nil {
		panic(err)
	}

	// 插入记录
	passHex, err := hex.DecodeString("68F18A7E6B7E9645F2E32CE1346BF9C5")
	if err != nil {
		panic(passHex)
	}
	ret := Insert().Table("users").Value(SBValues{"name": "小明", "old": 11, "pwd": passHex}).Exec()
	if !ret.Success {
		panic(err)
	}
	fmt.Printf("插入新记录 => ID为: %d", ret.LastID)

	_ = Insert().Table("users").Value(SBValues{"name": "小黑", "old": 12, "pwd": passHex}).Exec()
	_ = Insert().Table("users").Value(SBValues{"name": "小白", "old": 12, "pwd": passHex}).Exec()
	_ = Insert().Table("users").Value(SBValues{"name": "小二", "old": 14, "pwd": passHex}).Exec()
	_ = Insert().Table("users").Value(SBValues{"name": "小三", "old": 15, "pwd": passHex}).Exec()

	// 查表全部记录
	rows, err := Select().From("users").Query()
	for _, row := range rows {
		fmt.Println("多条记录 =>", row)
	}

	// 查表部分记录
	rows, err = Select().From("users").Limit(2, 1).Order("id DESC").Query()
	for _, row := range rows {
		fmt.Println("指定数量多条记录 =>", row)
	}

	// 查单行记录
	row, err := Select().From("users").Where("id=?").QueryOne(3)
	if err != nil {
		panic(err)
	}
	fmt.Printf("查单条记录 => ID: %d, Name: %s, Old: %d\n", row.GetInt("id"), row.Get("name"), row.GetInt("old"))

	// 获取单行二进制字段数据
	var pass []byte
	err = Select("pwd").From("users").Where("id=?").QueryRow(3).Scan(&pass)
	if err != nil {
		panic(err)
	}
	fmt.Printf("查单条二进制 => Pass: %X\n", pass)

	// 带Group字句
	rows, err = Select("old, SUM(1) AS total").From("users").Group("old").Query()
	if err != nil {
		panic(err)
	}
	for _, row := range rows {
		fmt.Printf("GROUP => old: %d, count: %d\n", row.GetInt("old"), row.GetInt("total"))
	}

	// 更新
	ret = Update().Table("users").AddValue("name", "小花").AddValue("old", 10).Where("id=?").Exec(3)
	if !ret.Success {
		panic(ret.Err)
	}
	fmt.Printf("更新 => %d 条记录", ret.Affected)

	ret = Update().Table("users").Value(SBValues{"name": "小花222", "old": 11}).Where("id=?").Exec(3)
	if !ret.Success {
		panic(ret.Err)
	}
	fmt.Printf("更新 => %d 条记录", ret.Affected)

	// 基于唯一索引, 存在则更新, 不存在则插入
	ret = InsertUpdate().Table("users").Value(SBValues{"id": 4, "name": "二狗", "old": 21}).Value2(SBValues{"old": 21}).Exec()
	if !ret.Success {
		panic(ret.Err)
	}

	// 删除
	ret = Delete().Table("users").Where("id=?").Exec(5)
	if !ret.Success {
		panic(ret.Err)
	}
	fmt.Printf("删除 => %d 条记录", ret.Affected)

	// 也可通过 db.DefaultDB 直接传入完整对 SQL 语句查询
	// db.DefaultDB.Exec(sql[, args...])
	// db.DefaultDB.Insert(sql[, args...])
	// db.DefaultDB.Update(sql[, args...])
	// db.DefaultDB.Delete(sql[, args...])
	// db.DefaultDB.Select()
	// db.DefaultDB.SelectOne()
	// db.DefaultDB.Query()
	// db.DefaultDB.QueryRow()

	// 或通过 db.DefaultDB.DB 访问 database/sql 调用标准库自带方法
}*/
