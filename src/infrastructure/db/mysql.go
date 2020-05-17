package db

import (
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Config struct {
	DB struct {
		Host     string
		Username string
		Password string
		DBName   string
	}
}

func NewConfig() *Config {

	c := new(Config)

	c.DB.Host = os.Getenv("MYSQL_HOST")
	c.DB.Username = os.Getenv("MYSQL_USER")
	c.DB.Password = os.Getenv("MYSQL_PASSWORD")
	c.DB.DBName = os.Getenv("MYSQL_DB")

	return c
}

type DB struct {
	Host     string
	Username string
	Password string
	DBName   string
	Connect  *gorm.DB
}

func NewDB() *DB {
	c := NewConfig()
	return newDB(&DB{
		Host:     c.DB.Host,
		Username: c.DB.Username,
		Password: c.DB.Password,
		DBName:   c.DB.DBName,
	})
}

func newDB(d *DB) *DB {
	log.Println(d)
	log.Println("connect to mysql server")

	var db *gorm.DB
	var err error
	// DBにつながらない時はリトライする
	for i := 1; i < 5; i++ {
		db, err = gorm.Open("mysql", d.Username+":"+d.Password+"@tcp("+d.Host+")/"+d.DBName+"?charset=utf8&parseTime=True&loc=Local")
		db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8 auto_increment=1")
		if err == nil {
			break
		}
		time.Sleep(time.Second * 5)
	}
	if err != nil {
		panic(err.Error())
	}

	d.Connect = db
	return d
}
