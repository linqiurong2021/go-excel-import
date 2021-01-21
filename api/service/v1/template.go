package v1

import (
	"fmt"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

// Template Template
type Template struct {
	// 模型名称
	TableEnName string
}

// NewTemplate NewTemplate
func NewTemplate() *Template {
	return &Template{}
}

// OpenFile OpenFile
func (e *Template) openFile(filePath string) (f *excelize.File, err error) {
	f, err = excelize.OpenFile(filePath)
	return
}

// ReadTemplate ReadTemplate
func (e *Template) ReadTemplate(excelPath string) (createTableSQL string, insertSQL string, err error) {
	// 运行时的目录为主
	f, err := e.openFile(excelPath) // "./template/template_test.xlsx"
	if err != nil {
		fmt.Printf("open file error: %s\n", err)
		return "", "", err
	}
	// 获取第一个sheet名称
	firstSheetName := f.GetSheetName(0)
	// 获取行数
	rows, err := f.GetRows(firstSheetName)
	createTableSQL, insertSQL = e.rowsHandle(rows)

	return
}

// RowsHandle 行数据处理
func (e *Template) rowsHandle(rows [][]string) (createTableSQL string, insertSQL string) {
	// 表单名称
	tableName := e.getTableName(rows[0])
	// 表单英文名称
	tableEnName := e.getTableEnName(rows[1])
	// 字段备注
	fieldsName := rows[2]
	// 字段名称
	fieldsEnName := rows[3]
	// 字段类型
	fieldsType := rows[4]
	// 设置表名
	e.TableEnName = tableEnName

	fmt.Printf("模版名称:%s , 数据表名称: %s\n", tableName, tableEnName)
	fmt.Printf("字段名称:%#v , 字段英文名称: %#v\n", fieldsName, fieldsEnName)
	fmt.Printf("字段类型:%#v \n", fieldsType)
	//
	createTableSQL = e.createTable(tableName, tableEnName, fieldsName, fieldsEnName, fieldsType)
	fmt.Printf("创建表SQL:\n %s \n", createTableSQL)
	// 执行sql
	insertData := e.getInsertData(rows)
	fmt.Printf("插入的数据:\n %s \n", insertData)

	insertSQL = e.insertToTable(tableEnName, fieldsEnName, insertData)
	fmt.Printf("\n%s\n", insertSQL)

	return
}

// CreateTable 创建数据库
func (e *Template) createTable(tableName string, tableEnName string, fieldsName []string, fieldsEnName []string, fieldsType []string) string {

	var fieldsSQL string
	//
	for i, field := range fieldsName {
		fieldTemplate := fmt.Sprintf("	`%s` %s COMMENT '%s',\n", fieldsEnName[i], fieldsType[i], field)
		fieldsSQL += fieldTemplate
	}
	// 系统默认主键
	fieldsSQL += "`SYS_ID` bigint(15) NOT NULL AUTO_INCREMENT,\n PRIMARY KEY (`SYS_ID`)"
	createTableSQL := fmt.Sprintf("CREATE TABLE `%s` (\n %s ) \n ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT = '%s';", tableEnName, fieldsSQL, tableName)

	return createTableSQL
}

// GetInsertData 获取入库的数据
func (e *Template) getInsertData(rows [][]string) (insertData string) {
	// 插入数据
	dataRows := rows[5:]
	for _, row := range dataRows {
		rowData := ""
		for _, cel := range row {
			// 需要都转为字符串 否则在插入时会报错
			rowData += fmt.Sprintf("'%s',", cel)
		}
		insertData += fmt.Sprintf("(%s),\n", strings.TrimRight(rowData, ","))
	}
	// 删除最后一个,
	insertData = strings.TrimRight(insertData, ",\n")
	return
}

// InsertToTable 数据入表
func (e *Template) insertToTable(tableEnName string, fieldsEnName []string, data string) (insertSQL string) {
	//
	fields := strings.Join(fieldsEnName, ",")
	// 插入数据
	insertSQL = fmt.Sprintf("INSERT INTO `%s` (%s) VALUES \n %s ", tableEnName, fields, data)

	return
}

// GetTableName 获取模版名称
func (e *Template) getTableName(rows []string) string {
	return rows[0]
}

// GetTableEnName 获取模版名称
func (e *Template) getTableEnName(rows []string) string {
	return rows[0]
}
