package v1

import (
	"fmt"
	"strings"
	"time"

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

// 获取最后的行数
func (e *Template) getColumn(column string, row int) (endColumn string) {
	endColumn = fmt.Sprintf("%s%d", column, row)
	return
}

// GetRowEndColunn 获取行的最后列
func (e *Template) GetRowEndColunn(column int, row int) (columnChar string) {
	// 可使用字符串长度
	const CharCount int = 26
	// 定义A
	const A int = 65 // 如果从64开始则不需要减1
	// 计算第几列
	offset := (column - 1) / CharCount // column-1
	//
	column = (column - 1) % CharCount
	//
	if offset == 0 { // 从A开始
		columnChar = fmt.Sprintf("%s%d", string(A+column), row)
	} else { // 从AA开始
		for i := 0; i < offset; i++ {
			columnChar = fmt.Sprintf("%s%s%d", string(A+i), string(A+column), row)
		}
	}
	return
}

// WriteTemplate 写模板
func (e *Template) WriteTemplate(path string, templateCnName string, templateName string, fields []string, configList []map[int]interface{}, dataList []map[int]interface{}) (downloadPath string, err error) {
	// 创建excel
	f := excelize.NewFile()
	f.SetActiveSheet(0)
	sheetName := "Sheet1"
	// 定义A
	const A int = 97
	const CharCount int = 26
	// 开始行数
	var startRow int = 2 //
	// 配置项行数
	var configRow int
	// 文字垂直居中 水平居中
	style, err := f.NewStyle(`{"alignment":{"horizontal":"center","vertical":"ceneter"}}`)
	if err != nil {
		return "", err
	}
	// 设置表头(名称与模板名称)
	A1 := e.getColumn("A", 1)
	A2 := e.getColumn("A", 2)
	// len(fields)
	A1EndColumn := e.GetRowEndColunn(len(fields), 1)
	A2EndColumn := e.GetRowEndColunn(len(fields), 2)
	// 合并单元格
	f.MergeCell(sheetName, A1, A1EndColumn)
	f.MergeCell(sheetName, A2, A2EndColumn)
	// // 设置居中
	f.SetCellStyle(sheetName, A1, A1EndColumn, style)
	f.SetCellStyle(sheetName, A2, A2EndColumn, style)
	// 设置表头(名称与模板名称)
	f.SetCellValue(sheetName, A1, templateCnName) // A1
	f.SetCellValue(sheetName, A2, templateName)   // A2
	// 字段开始行
	startRow = startRow + 1
	// 设置配置
	for row, item := range configList {
		// col 从0开始
		for col, val := range item {
			// 计算第几列
			column := e.GetRowEndColunn(col+1, row+startRow)
			f.SetCellValue(sheetName, column, val)
		}
		configRow = configRow + 1
	}
	// 数据开始行
	startRow = startRow + configRow
	// 数据开始
	// fmt.Printf("#########\n %d #######\n", configRow)
	for row, item := range dataList {
		for col, val := range item {
			// 计算第几列
			column := e.GetRowEndColunn(col+1, row+startRow)
			f.SetCellValue(sheetName, column, val)
		}
	}
	downloadPath = fmt.Sprintf("/downloads/%s_%d.xlsx", templateCnName, time.Now().Unix())
	// Save spreadsheet by the given path.
	filename := fmt.Sprintf("%s%s", path, downloadPath)
	if err := f.SaveAs(filename); err != nil {
		fmt.Println(err)
	}

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
	dataRows := rows[2:8]
	//
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
