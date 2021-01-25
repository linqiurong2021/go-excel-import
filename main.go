package main

import (
	"fmt"
	"log"

	"github.com/linqiurong2021/go-excel-import/api/db"
	v1 "github.com/linqiurong2021/go-excel-import/api/service/v1"
	"github.com/linqiurong2021/go-excel-import/conf"
)

func main() {

	err := conf.InitConfig("./conf/config.ini")
	if err != nil {
		log.Fatal("init config error: ", err)
	}
	// 数据库
	db := db.NewDB()
	// 创建模板
	// createTable(db)
	createLogic(db)
	//
	// v1.NewService(db)
	// service := v1.NewService(db)
	// service.GetAllData("", )

}

func createLogic(db *db.DB) {
	logic := v1.NewLogic(db)
	if err := logic.CreateTable("./template/template_test_2.xlsx", false); err != nil {
		fmt.Printf("CreateTable error %s\n", err)
	}
}

func createTable(db *db.DB) {

	execl := v1.NewTemplate()
	//
	createTableSQL, insertSQL, err := execl.ReadTemplate("./template/template_test_2.xlsx")
	if err != nil {
		fmt.Printf("ReadTemplate error: %s\n", err)
	}
	result, err := db.ExecuteSQLResult(createTableSQL)
	if err != nil {
		fmt.Printf("create table error: %s\n", err)
	}
	fmt.Printf("create table result :%#v\n", result)

	result, err = db.ExecuteSQLResult(insertSQL)
	if err != nil {
		fmt.Printf("insert data error: %s\n", err)
	}
	fmt.Printf("insert data %#v\n ", result)
}
