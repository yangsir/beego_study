package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	_DB_NAME      = "root:123456@/test?charset=utf8"
	_MYSQL_DRIVER = "mysql"
)

// 分类
type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

// 文章
type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegisterDB() {
	// 注册模型
	orm.RegisterModel(new(Category), new(Topic))
	orm.RegisterDriver(_MYSQL_DRIVER, orm.DRMySQL)

	// 注册默认数据库
	orm.RegisterDataBase("default", _MYSQL_DRIVER, _DB_NAME, 10)
}
