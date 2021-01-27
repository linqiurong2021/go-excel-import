package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/linqiurong2021/go-excel-import/conf"
)

// DB DB
type DB struct {
	DB *sql.DB
}

// NewDB NewDB
func NewDB() *DB {
	db := initDB()
	return &DB{
		DB: db,
	}
}

// initDB 初始化数据库连接
func initDB() *sql.DB {
	params := "charset=utf8&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", conf.Conf.DBConfig.User, conf.Conf.DBConfig.Password, conf.Conf.DBConfig.Host, conf.Conf.DBConfig.Sechma, params)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("connect database error ", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("ping error ", err)
	}
	return db
}

// ExecuteSQLResult 获取执行结果(新增或更新)
func (db *DB) ExecuteSQLResult(SQL string) (result sql.Result, err error) {
	result, err = db.DB.ExecContext(context.Background(), SQL)
	return
}

// ExecuteSQLRows 获取执行行
func (db *DB) ExecuteSQLRows(SQL string) (rows *sql.Rows, err error) {
	rows, err = db.DB.QueryContext(context.Background(), SQL)
	return
}

// ExecuteSQLResultTx 获取执行结果(新增或更新)
func (db *DB) ExecuteSQLResultTx(tx *sql.Tx, SQL string) (result sql.Result, err error) {
	result, err = tx.ExecContext(context.Background(), SQL)
	return
}

// ExecuteSQLRowsTx 获取执行行
func (db *DB) ExecuteSQLRowsTx(tx *sql.Tx, SQL string) (rows *sql.Rows, err error) {
	rows, err = tx.QueryContext(context.Background(), SQL)
	return
}
