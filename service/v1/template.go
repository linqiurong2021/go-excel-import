package v1

import (
	"fmt"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/linqiurong2021/go-excel-import/conf"
)

// Template Template
type Template struct {
	// 模型名称
	TableName string
	// 表名称
	TableEnName string
	// 备注列表
	FieldsName []string
	// 字段列表
	FieldsEnName []string
	// 类型列表
	FieldsType []string
}

// Table 数据表
type Table struct {
	tableSQL        string
	insertTableSQL  string
	configTableSQL  string
	insertConfigSQL string
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
func (e *Template) ReadTemplate(excelPath string) (table *Table, err error) {
	// 运行时的目录为主
	f, err := e.openFile(excelPath) // "./template/template_test.xlsx"
	if err != nil {
		fmt.Printf("open file error: %s\n", err)
		return nil, err
	}
	// 获取第一个sheet名称
	firstSheetName := f.GetSheetName(0)
	// 获取行数
	rows, err := f.GetRows(firstSheetName)
	table = e.rowsHandle(rows)

	return
}

// getTemplateInfo 获取模板信息
func (e *Template) getTemplateInfo(rows [][]string) {
	// 表单名称
	e.getTableName(rows)
	// 表单英文名称
	e.getTableEnName(rows)
	// 字段备注
	e.getFieldName(rows)
	// 字段名称
	e.getFieldEnName(rows)
	// 字段类型
	e.getFieldType(rows)
}

// RowsHandle 行数据处理
func (e *Template) rowsHandle(rows [][]string) (table *Table) {
	// 获取基本信息
	e.getTemplateInfo(rows)
	// 获取新建表的SQL 语句
	tableSQL := e.getCreateTableSQL()
	fmt.Printf("创建表SQL:\n %s \n", tableSQL)
	// 获取表单数据
	tableData := e.getInsertData(rows)
	// 获取插入数据的SQL 语句
	insertTableSQL := e.getInsertSQL(e.TableEnName, e.FieldsEnName, tableData)
	configTableSQL := e.getCreateConfigTableSQL()
	configTableData := e.getInsertConfigData(rows)
	configTableName := fmt.Sprintf("%s%s", e.TableEnName, conf.Conf.DBConfig.ConfigTableSuffix)
	insertConfigSQL := e.getInsertConfigSQL(configTableName, e.FieldsEnName, configTableData)
	// 配置
	var importTable = &Table{
		tableSQL:        tableSQL,
		configTableSQL:  configTableSQL,
		insertTableSQL:  insertTableSQL,
		insertConfigSQL: insertConfigSQL,
	}

	fmt.Printf("\n%#v\n", table)
	return importTable
}

// GetFieldEnName 获取字段名称
func (e *Template) getFieldEnName(rows [][]string) []string {
	e.FieldsEnName = rows[3]
	return e.FieldsEnName
}

// GetFieldName 获取字段名称
func (e *Template) getFieldName(rows [][]string) []string {
	e.FieldsName = rows[2]
	return e.FieldsName
}

// GetFieldType 获取字段类型
func (e *Template) getFieldType(rows [][]string) []string {
	e.FieldsType = rows[4]
	return e.FieldsType
}

// GetTableName 获取模版名称
func (e *Template) getTableName(rows [][]string) string {
	e.TableName = rows[0][0]
	return e.TableName
}

// GetTableEnName 获取模版名称
func (e *Template) getTableEnName(rows [][]string) string {
	e.TableEnName = rows[1][0]
	return e.TableEnName
}

// getCreateConfigTableSQL 获取配置数据表
func (e *Template) getCreateConfigTableSQL() string {
	//
	var fieldsSQL string
	// 系统默认字段
	systemDefaultFields := "`SYS_ID` int(1) NOT NULL AUTO_INCREMENT,\n 	PRIMARY KEY (`SYS_ID`)"

	//
	for i, field := range e.FieldsName {
		fieldTemplate := fmt.Sprintf("	`%s` varchar(50) COMMENT '%s',\n", e.FieldsEnName[i], field) // 固定使用varchar 来存储一些数据
		fieldsSQL += fieldTemplate
	}
	// 系统默认主键

	fieldsSQL += systemDefaultFields // "`SYS_ID` bigint(15) NOT NULL AUTO_INCREMENT,\n PRIMARY KEY (`SYS_ID`)"
	createTableSQL := fmt.Sprintf("CREATE TABLE `%s%s` (\n %s ) \n ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT = '%s';", e.TableEnName, conf.Conf.DBConfig.ConfigTableSuffix, fieldsSQL, e.TableName)

	return createTableSQL
}

// createConfigTable 创建配置表
func (e *Template) createConfigTable() {
	// 获取创建表数据
	e.getCreateConfigTableSQL()
	// 获取
}

// getInsertData 获取入库的数据
func (e *Template) getInsertConfigData(rows [][]string) (insertData string) {
	// 插入数据
	dataRows := rows[5:8]
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

// CreateTable 创建数据库
func (e *Template) getCreateTableSQL() string {

	var fieldsSQL string
	// 系统默认字段
	systemDefaultFields := "	`SYS_CREATE_AT` timestamp NULL DEFAULT CURRENT_TIMESTAMP ,\n	`SYS_UPDATE_AT` timestamp NULL ON UPDATE CURRENT_TIMESTAMP,\n	`SYS_NOTE` varchar(255) NULL,\n	`SYS_ID` bigint(15) NOT NULL AUTO_INCREMENT,\n 	PRIMARY KEY (`SYS_ID`)"

	//
	for i, field := range e.FieldsName {
		fieldTemplate := fmt.Sprintf("	`%s` %s COMMENT '%s',\n", e.FieldsEnName[i], e.FieldsType[i], field)
		fieldsSQL += fieldTemplate
	}
	// 系统默认主键

	fieldsSQL += systemDefaultFields // "`SYS_ID` bigint(15) NOT NULL AUTO_INCREMENT,\n PRIMARY KEY (`SYS_ID`)"
	createTableSQL := fmt.Sprintf("CREATE TABLE `%s` (\n %s ) \n ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT = '%s';", e.TableEnName, fieldsSQL, e.TableName)

	return createTableSQL
}

// getInsertData 获取入库的数据
func (e *Template) getInsertData(rows [][]string) (insertData string) {
	// 插入数据
	dataRows := rows[8:]
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

// getInsertSQL 数据入表
func (e *Template) getInsertSQL(tableEnName string, fieldsEnName []string, data string) (insertSQL string) {
	//
	fields := strings.Join(fieldsEnName, ",")
	// 插入数据
	insertSQL = fmt.Sprintf("INSERT INTO `%s` (%s) VALUES \n %s ", tableEnName, fields, data)

	return
}

// getInsertSQL 数据配置入表
func (e *Template) getInsertConfigSQL(tableEnName string, fieldsEnName []string, data string) (insertSQL string) {
	//
	fields := strings.Join(fieldsEnName, ",")
	// 插入数据
	insertSQL = fmt.Sprintf("INSERT INTO `%s` (%s) VALUES \n %s ", tableEnName, fields, data)

	return
}
