package v1

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// Logic 逻辑
type Logic struct {
	// 数据库
	db *sql.DB
	// 模型
	template *Template
}

// NewLogic 创建逻辑
func NewLogic(db *sql.DB) *Logic {
	return &Logic{
		db:       db,
		template: &Template{},
	}
}

// backupTable 备份数据表
func (l *Logic) backupTable() {
	// 需要备份的数据表名
	tableEnName := l.template.TableEnName
	backupEnName := fmt.Sprintf("%s_%d", tableEnName, time.Now().Unix())
	fmt.Println("数据备份表名:", tableEnName)
	// 备份操作 https://blog.csdn.net/moshowgame/article/details/82952992
	backupSQL := fmt.Sprintf("CREATE TABLE %s SELECT * FROM %s ", backupEnName, tableEnName)
	fmt.Printf("备份数据%s表\n新表名为:%s\n 执行语句为:%s\n", tableEnName, backupEnName, backupSQL)
}

// ExistTable 判断表名是否存在
func (l *Logic) existTable(tableName string) (exists bool, err error) {
	return false, nil
}

// 创建数据表
func (l *Logic) createTable(createTableSQL string) error {
	// 判断表是否存在
	exists, err := l.existTable(l.template.TableEnName)
	if err != nil {
		return err
	}
	// 需要开启事务
	if exists {
		// 先备份再进行操作
		l.backupTable()
	}
	// 如果不存在则直接创建
	createTableSQL, insertSQL, err := l.template.ReadTemplate("./template/template_test.xlsx")
	// 如果存在则备份
	result, err := l.executeSQL(createTableSQL)
	result, err = l.executeSQL(insertSQL)
	// 事务提交
	fmt.Printf("%#v", result)
	return nil
}

// 新增数据
func (l *Logic) insertData(insertSQL string) error {
	return nil
}

// 执行SQL语句
func (l *Logic) executeSQL(SQL string) (result sql.Result, err error) {
	result, err = l.db.ExecContext(context.Background(), SQL)
	return
}
