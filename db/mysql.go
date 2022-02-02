package db

import (
	"log"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/citymall/geo-loc/types"
	"github.com/citymall/geo-loc/util"
	_ "github.com/go-sql-driver/mysql"
)

type sqlConnectionType struct {
	isconnected bool
}

var SqlConnection = &sqlConnectionType{}

func (db *sqlConnectionType) ConnectMySql() {
	// if not conected already then connect
	if !db.isconnected {
		mysql := util.GetConfig().Mysql
		mysqlConf := mysql["user"].(string) + ":" + mysql["password"].(string) + "@tcp(" + mysql["host"].(string) + ")/" + mysql["database"].(string)
		log.Println("conf", mysqlConf)
		orm.RegisterDataBase("default", "mysql", mysqlConf)
		orm.DefaultTimeLoc, _ = time.LoadLocation("Asia/Kolkata")
		orm.Debug = true
		db.isconnected = true
		log.Println("SQL Connected!")
	} else {
		// do nothing
	}

	return
}

func init() {
	orm.RegisterModel(new(types.SqlDbOutput))
}

//GetAllInstagramPost retrieves all InstagramPost matches certain condition. Returns empty list if
// no records exist
func (db *sqlConnectionType) GetAllDataFromSql() (v []types.SqlDbOutput, err error) {
	if !db.isconnected {
		db.ConnectMySql()
	}
	o := orm.NewOrm()
	v = []types.SqlDbOutput{}
	_, err = o.QueryTable(new(types.SqlDbOutput)).RelatedSel().OrderBy("id").All(&v)
	return v, err
}
