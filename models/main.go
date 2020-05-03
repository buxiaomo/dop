package models

import (
	"database/sql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

var (
	o orm.Ormer
)

type DBInit struct {
	Db *sql.DB
}

func DBInitNew() *DBInit {
	db, _ := orm.GetDB()
	dbInit := DBInit{Db: db}
	return &dbInit
}

func (dbInit *DBInit) InitDatabase() {
	if dbInit.Db == nil {
		dbType := beego.AppConfig.String("dbType")
		sqlConn := beego.AppConfig.String("sqlconn")
		maxIdle, _ := beego.AppConfig.Int("maxIdle")
		maxConn, _ := beego.AppConfig.Int("maxConn")

		switch dbType {
			case "mysql":
				orm.RegisterDriver("mysql", orm.DRMySQL)
				orm.RegisterDataBase("default", dbType, sqlConn, maxIdle, maxConn)
			case "sqlite3":
				orm.RegisterDataBase("default", dbType, sqlConn)
		}
		orm.Debug = beego.AppConfig.DefaultBool("db_debug", true)
	}
}

func initTable(){
	if o.QueryTable(new(User)).Filter("username", "admin").Exist() == false {
		user := User {
			Username: "admin",
			Password: "admin",
			Email: "95112082@qq.com",
			Lock: false,
		}
		o.Insert(&user)
		o.Commit()
	}
	if o.QueryTable(new(Role)).Filter("role", "admin").Exist() == false {
		roles := []Role {
			{Role: "admin"},
			{Role: "dev"},
			{Role: "ops"},
		}
		o.InsertMulti(len(roles),roles)
		o.Commit()
	}
}

func init() {
	DBInitNew().InitDatabase()
	orm.RegisterModel(new(User), new(Role), new(Project), new(Cluster))
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		logs.Error(err.Error())
	}
	o = orm.NewOrm()
	o.Using("default")

	initTable()
}