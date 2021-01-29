package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/linqiurong2021/go-excel-import/conf"
	"github.com/linqiurong2021/go-excel-import/db"
	"github.com/linqiurong2021/go-excel-import/server"
	v1 "github.com/linqiurong2021/go-excel-import/service/v1"
)

func main() {

	err := conf.InitConfig("./conf/config.ini")
	if err != nil {
		log.Fatal("init config error: ", err)
	}
	// 数据库
	db := db.NewDB()

	server.NewRestful(db).StartServer()
	// 创建模板
	// createTable(db)
	// createLogic(db)

	// createService(db)
	//
	// v1.NewService(db)
	// service := v1.NewService(db)
	// service.GetAllData("", )

}

func createService(db *db.DB) {
	service := v1.NewService(db)
	data, err := service.GetConfig("TemplateTest")
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	row, err := json.Marshal(data[0])
	if err != nil {
		fmt.Printf("err %\n", err)
	}
	fieldType := string(row)

	row, err = json.Marshal(data[1])
	if err != nil {
		fmt.Printf("err %\n", err)
	}
	searchField := string(row)

	row, err = json.Marshal(data[2])
	if err != nil {
		fmt.Printf("err %\n", err)
	}
	listField := string(row)
	fmt.Printf("%s\n%s\n%s\n", fieldType, searchField, listField)
	// for _, item := range data {
	// 	fmt.Println(item)
	// }
}

func createLogic(db *db.DB) {
	logic := v1.NewLogic(db)
	if err := logic.CreateTable("./template/template_test.xlsx", false); err != nil {
		fmt.Printf("CreateTable error %s\n", err)
	}

}

func createTable(db *db.DB) {

	// execl := v1.NewTemplate()
	// //
	// table, err := execl.ReadTemplate("./template/template_test_2.xlsx")
	// if err != nil {
	// 	fmt.Printf("ReadTemplate error: %s\n", err)
	// }
	// result, err := db.ExecuteSQLResult(table.tableSQL)
	// if err != nil {
	// 	fmt.Printf("create table error: %s\n", err)
	// }
	// fmt.Printf("create table result :%#v\n", result)

	// result, err = db.ExecuteSQLResult(table.insertTableSQL)
	// if err != nil {
	// 	fmt.Printf("insert data error: %s\n", err)
	// }
	// fmt.Printf("insert data %#v\n ", result)
}
