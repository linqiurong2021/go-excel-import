package v1

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/fatih/set"
	"github.com/linqiurong2021/go-excel-import/conf"
	"github.com/linqiurong2021/go-excel-import/db"
	"github.com/linqiurong2021/go-excel-import/utils"
)

// Service Service
type Service struct {
	// 数据库
	db *db.DB
}

// NewService 创建服务
func NewService(db *db.DB) *Service {
	return &Service{
		db: db,
	}
}

// getFields 获取表字段
func (s *Service) getFields(tableName string) (fields []string, err error) {
	// 需要加 DISTINCT 否则有可能重复
	getFieldsSQL := fmt.Sprintf("SELECT DISTINCT(COLUMN_NAME) FROM information_schema.COLUMNS where table_name = '%s';", tableName)
	rows, err := s.db.ExecuteSQLRows(getFieldsSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// fmt.Printf("Rows:%#v\n", rows)
	var field string
	for rows.Next() {
		//
		if err = rows.Scan(&field); err != nil {
			return
		}
		// 取消
		if field != "SYS_CREATE_AT" && field != "SYS_UPDATE_AT" && field != "SYS_NOTE" {
			fields = append(fields, field)
		}

	}

	return fields, nil
}

// getDistinctValueByField 获取不重复数据(下拉或单选或多选)
func (s *Service) getDistinctValueByField(fieldName string, tableName string) (rows *sql.Rows, err error) {
	distinctSQL := fmt.Sprintf("SELECT DISTINCT %s FROM %s ", fieldName, tableName)
	rows, err = s.db.ExecuteSQLRows(distinctSQL)
	if err != nil {
		return nil, err
	}
	return
}

// getParamsFieldsAndValues 获取参数中的字段
func (s *Service) getParamsFieldsAndValues(searchFields []string, searchValues []string, searchFieldTypes []string) string {
	// 需要判断是什么类型的 如果是下拉则直接使用等于 其它使用 like %%
	// 如果无搜索条件则直接返回
	if len(searchFields) <= 0 || len(searchValues) <= 0 || len(searchFieldTypes) <= 0 {
		fmt.Printf("search empty: %#v\n %#v\n %#v\n", searchFields, searchValues, searchFieldTypes)
		return ""
	}
	var tmpWhere string
	for index, item := range searchFieldTypes {
		if item == "下拉框" {
			tmpWhere = tmpWhere + fmt.Sprintf(" %s = '%s' and ", searchFields[index], searchValues[index])
		} else {
			likeStr := "'%" + searchValues[index] + "%'"
			tmpWhere = tmpWhere + fmt.Sprintf(" %s like %s and ", searchFields[index], likeStr)
		}
	}
	fmt.Printf("#where# %v\n", tmpWhere)
	// 删除最后一个and
	return strings.Trim(tmpWhere, " and ")
}

// GetAllData 获取所有数据
func (s *Service) GetAllData(tableName string, params map[string]string) (rows *sql.Rows, err error) {
	// 获取字段
	s.getFields(tableName)
	fields := "*"
	//
	var getAllSQL string
	// 获取参数中的条件
	where := "" // s.getParamsFieldsAndValues([],[]])
	if len(where) == 0 {
		// 获取所有的数据
		getAllSQL = fmt.Sprintf("SELECT %s FROM %s ", tableName, fields)
	} else {
		// 获取所有的数据
		getAllSQL = fmt.Sprintf("SELECT %s FROM %s WHRE %s ", tableName, fields, where)
	}
	rows, err = s.db.ExecuteSQLRows(getAllSQL)
	// 返回数据
	return rows, nil
}

// GetDataByPageSize 分页(需要有参数)
func (s *Service) GetDataByPageSize(tableName string, page int, pageSize int, searchFields []string, searchValues []string, searchFieldsType []string) (list []map[string]interface{}, total int64, err error) {

	// 获取字段
	fields, err := s.getFields(tableName)
	if err != nil {
		return nil, 0, err
	}
	// 判断搜索条件是否是字集 如果不是字集则提示
	diff := utils.Difference(searchFields, fields)
	fmt.Printf("# diff%#v #\n", diff)
	if len(diff) > 0 {
		errMsg := fmt.Sprintf("%s params invalidate", strings.Join(diff, ","))
		return nil, 0, errors.New(errMsg)
	}
	var getAllSQL string
	var totalSQL string
	// pageSize
	offset := (page - 1) * pageSize

	fmt.Printf("searchValues: %#v\n searchFieldsType: %#v\n", searchValues, searchFieldsType)
	// 获取参数中的条件
	where := s.getParamsFieldsAndValues(searchFields, searchValues, searchFieldsType)
	if len(where) <= 0 {
		totalSQL = fmt.Sprintf("SELECT COUNT(SYS_ID) FROM `%s`;", tableName)
		// 获取所有的数据
		getAllSQL = fmt.Sprintf("SELECT %s FROM `%s` ORDER BY SYS_ID LIMIT %d, %d ;", strings.Join(fields, ","), tableName, offset, pageSize)
	} else {
		totalSQL = fmt.Sprintf("SELECT COUNT(SYS_ID) FROM `%s` WHERE %s;", tableName, where)
		// 获取所有的数据 分页
		getAllSQL = fmt.Sprintf("SELECT %s FROM `%s` WHERE %s ORDER BY SYS_ID LIMIT %d,%d;", strings.Join(fields, ","), tableName, where, offset, pageSize)
	}
	// fmt.Printf("#total#: %v\n#getAllSQL#: %v", totalSQL, getAllSQL)
	// 获取所有数据
	totalRows, err := s.db.ExecuteSQLRows(totalSQL)
	if err != nil {
		return nil, 0, err
	}
	// 总行数
	if totalRows.Next() {
		err = totalRows.Scan(&total)
		if err != nil {
			return nil, 0, err
		}
	}

	rows, err := s.db.ExecuteSQLRows(getAllSQL)
	if err != nil {
		return nil, 0, err
	}
	//
	scanArgs := make([]interface{}, len(fields))
	for i := range fields {
		scanArgs[i] = &fields[i]
	}

	columns, _ := rows.Columns()
	for rows.Next() {
		if err := rows.Scan(scanArgs...); err != nil {
			return nil, 0, err
		}
		item := make(map[string]interface{})
		for i, data := range scanArgs {
			item[columns[i]] = *data.(*string) //取实际类型
		}
		list = append(list, item)
	}
	fmt.Printf("##list:%+v##, total: %d\n", list, total)
	// 返回数据
	return list, total, nil

}

// GetDataByID 获取一条数据
func (s *Service) GetDataByID(tableName string, sysID int64) (list map[string]interface{}, err error) {
	// 获取字段
	fields, err := s.getFields(tableName)
	if err != nil {
		return nil, err
	}

	getSQL := fmt.Sprintf("SELECT %s FROM %s WHERE SYS_ID = %d", strings.Join(fields, ","), tableName, sysID)
	fmt.Printf("GET DATA SQL:%#v\n", getSQL)
	rows, err := s.db.ExecuteSQLRows(getSQL)
	//
	scanArgs := make([]interface{}, len(fields))
	for i := range fields {
		scanArgs[i] = &fields[i]
	}

	item := make(map[string]interface{}, len(fields))

	if rows.Next() {
		columns, _ := rows.Columns()
		if err := rows.Scan(scanArgs...); err != nil {
			return nil, err
		}

		for i, data := range scanArgs {
			item[columns[i]] = *data.(*string) //取实际类型
		}
	}
	return item, nil
}

// CreateData 新增一条数据
func (s *Service) CreateData(tableName string, params map[string]interface{}) (result sql.Result, err error) {
	// 获取字段
	fields, err := s.getFields(tableName)
	// 判断更新的字段 是否在数据表中
	var addFields []string
	for k := range params {
		addFields = append(addFields, k)
	}
	diff := utils.Difference(addFields, fields)
	if len(diff) > 0 {
		errMsg := fmt.Sprintf("%s fields not exists", strings.Join(diff, ","))
		return nil, errors.New(errMsg)
	}
	// 删除SYS_ID
	noSysIDFields := fields[0 : len(fields)-1]
	//
	inserData := s.getCreateKVSQL(params, noSysIDFields)
	if inserData == "" {
		return nil, errors.New("Create failure")
	}

	// 数据处理
	insertSQL := fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s)", tableName, strings.Join(noSysIDFields, ","), inserData)
	fmt.Printf("#insertSQL# %v\n", insertSQL)
	result, err = s.db.ExecuteSQLResult(insertSQL)
	return
}

// getCreateKVSQL 获取新增SQL
func (s *Service) getCreateKVSQL(params map[string]interface{}, noSysIDFields []string) string {
	var tmpSQL string
	for _, v := range noSysIDFields {
		tmpSQL += fmt.Sprintf("'%s'%s", params[v], ",")
	}
	return strings.Trim(tmpSQL, ",")
}

// getUpdateSQL 获取更新语句
func (s *Service) getUpdateKVSQL(params map[string]interface{}) string {
	var tmpSQL string
	for k, v := range params {
		// 不修改SYS_ID
		if k != "SYS_ID" {
			tmpSQL += fmt.Sprintf("%s = '%s'%s", k, v, ",")
		}
	}
	return strings.Trim(tmpSQL, ",")
}

// UpdateByID 更新一条数据
func (s *Service) UpdateByID(tableName string, params map[string]interface{}) (result sql.Result, err error) {
	fields, err := s.getFields(tableName)
	// 判断更新的字段 是否在数据表中
	var updateFields []string
	for k := range params {
		updateFields = append(updateFields, k)
	}
	if params["SYS_ID"] == "" {
		return nil, errors.New("SYS_ID must")
	}
	diff := utils.Difference(updateFields, fields)
	if len(diff) > 0 {
		errMsg := fmt.Sprintf("%s fields not exists", strings.Join(diff, ","))
		return nil, errors.New(errMsg)
	}
	updateData := s.getUpdateKVSQL(params)
	if updateData == "" {
		return nil, errors.New("Update failure")
	}
	fmt.Printf("#updateData#: %#v", updateData)
	// 数据处理
	updateSQL := fmt.Sprintf("UPDATE `%s` SET %s WHERE SYS_ID = %s ;", tableName, updateData, params["SYS_ID"])
	fmt.Printf("#updateSQL#: %#v", updateSQL)
	result, err = s.db.ExecuteSQLResult(updateSQL)
	return
}

// DeleteBySysIDs 更新一条数据
func (s *Service) DeleteBySysIDs(tableName string, sysIDs string) (result sql.Result, err error) {
	//
	if tableName == "" {
		return nil, errors.New("table invlidate")
	}
	if len(sysIDs) <= 0 {
		return nil, errors.New("sys_ids invlidate")
	}
	// 数据处理
	deleteSQL := fmt.Sprintf("DELETE FROM `%s` WHERE SYS_ID in (%s) ;", tableName, sysIDs)
	fmt.Printf("#deleteSQL#: %#v", deleteSQL)
	result, err = s.db.ExecuteSQLResult(deleteSQL)
	return
}

// batchAddData 批量添加数据
func (s *Service) batchAddData(tableName string, data map[string]string) (result sql.Result, err error) {
	//
	s.getFields(tableName)
	// 数据处理
	insertData := ""
	fields := "*"
	batchInsertSQL := fmt.Sprintf("INSERT INTO %s(%s) VALUES %s", tableName, fields, insertData)
	result, err = s.db.ExecuteSQLResult(batchInsertSQL)
	return
}

// GetConfigFields 获取配置文件
func (s *Service) getConfigFields(tableName string) (fields []string, err error) {
	//
	configTableName := tableName + conf.Conf.DBConfig.ConfigTableSuffix
	return s.getFields(configTableName)
}

// GetConfig 获取配置信息
func (s *Service) GetConfig(tableName string) (configs []map[string]string, err error) {
	// 字段
	fields, err := s.getConfigFields(tableName)
	if err != nil {
		return nil, err
	}
	// 获取配置数据
	getConfigSQL := fmt.Sprintf("SELECT %s FROM %s%s ORDER BY SYS_ID", strings.Join(fields, ","), tableName, conf.Conf.DBConfig.ConfigTableSuffix)
	// fmt.Printf("getConfigSQL:%s\n", getConfigSQL)
	// 执行
	rows, err := s.db.ExecuteSQLRows(getConfigSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// 获取类数据
	columns, _ := rows.Columns()

	// 以下代码来自官方 https://github.com/go-sql-driver/mysql/wiki/Examples
	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	var list []map[string]string
	// Fetch rows
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		var value string
		var rowMap = make(map[string]string)
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			rowMap[columns[i]] = value
			// fmt.Println(columns[i], ": ", value)
		}
		list = append(list, rowMap)
	}
	return list, nil
}

// GetSelectOptions 获取下拉配置信息
func (s *Service) GetSelectOptions(tableName string, keys []string) (options map[string][]string, err error) {
	// 字段
	fields, err := s.getConfigFields(tableName)
	if err != nil {
		return nil, err
	}
	// 需要判断字段是否在字段里
	fieldsSet := set.New(set.NonThreadSafe)
	fieldsSet.Add(strings.Join(fields, ","))
	keysSet := set.New(set.NonThreadSafe)
	keysSet.Add(strings.Join(fields, ","))

	defer fieldsSet.Clear()
	defer keysSet.Clear()

	// 如果字段不存在直接返回不存在的keys
	if !keysSet.IsSubset(fieldsSet) {
		diff := utils.Difference(keys, fields)
		invalidateFields := strings.Join(diff, ",")
		return nil, errors.New(invalidateFields + "keys invalidate")
	}
	// 已选择的数据
	var selectedOptions = make(map[string][]string, len(keys))
	// 如果有则 获取唯一值
	for _, key := range keys {

		optionsSQL := fmt.Sprintf("SELECT DISTINCT(%s) FROM %s", key, tableName)
		rows, err := s.db.ExecuteSQLRows(optionsSQL)
		if err != nil {
			return nil, err
		}
		//
		var value string
		var values []string
		for rows.Next() {
			//
			if err = rows.Scan(&value); err != nil {
				return nil, err
			}
			values = append(values, value)
		}
		selectedOptions[key] = values
	}
	return selectedOptions, nil
}

// GetFieldTypeConfig 获取类型配置
func (s *Service) GetFieldTypeConfig(configs []map[string]string) (jsonData string, err error) {
	json, err := json.Marshal(configs[0])
	if err != nil {
		return "", err
	}
	return string(json), nil
}

// GetSearchFieldConfig 获取搜索字段配置
func (s *Service) GetSearchFieldConfig(configs []map[string]string) (jsonData string, err error) {
	json, err := json.Marshal(configs[1])
	if err != nil {
		return "", err
	}
	return string(json), nil
}

// GetListFieldConfig 获取列表显示字段配置
func (s *Service) GetListFieldConfig(configs []map[string]string) (jsonData string, err error) {
	json, err := json.Marshal(configs[2])
	if err != nil {
		return "", err
	}
	return string(json), nil
}

// GetSearchFieldsType 获取搜索列表类型
func (s *Service) GetSearchFieldsType(tableName string, searchFieldsCopy []string) (fieldsType []string, err error) {
	fmt.Printf("fields: %#v\n", strings.Join(searchFieldsCopy, ","))
	// 获取配置数据
	getConfigSQL := fmt.Sprintf("SELECT %s FROM `%s%s` WHERE SYS_ID = 1; ", strings.Join(searchFieldsCopy, ","), tableName, conf.Conf.DBConfig.ConfigTableSuffix)
	// 执行
	rows, err := s.db.ExecuteSQLRows(getConfigSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	scanArgs := make([]interface{}, len(searchFieldsCopy))
	for i := range searchFieldsCopy {
		scanArgs[i] = &searchFieldsCopy[i]
	}
	if rows.Next() {
		//
		rows.Scan(scanArgs...)
	}
	for _, item := range scanArgs {
		fieldsType = append(fieldsType, *item.(*string)) //
	}
	fmt.Printf("fieldsType: %#v\n", fieldsType)
	return
}

// GetFieldsType 获取字段类型
func (s *Service) GetFieldsType(tableName string) (fieldsType map[string]string, err error) {
	fields, err := s.getFields(tableName)
	if err != nil {
		return nil, err
	}
	fmt.Printf("fields: %#v\n", strings.Join(fields, ","))
	// 获取配置数据
	getConfigSQL := fmt.Sprintf("SELECT %s FROM `%s%s` WHERE SYS_ID = 1; ", strings.Join(fields, ","), tableName, conf.Conf.DBConfig.ConfigTableSuffix)
	// 执行
	rows, err := s.db.ExecuteSQLRows(getConfigSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 获取类数据
	columns, _ := rows.Columns()

	// 以下代码来自官方 https://github.com/go-sql-driver/mysql/wiki/Examples
	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	var rowMap = make(map[string]string)
	// Fetch rows
	if rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		var value string

		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			rowMap[columns[i]] = value
			// fmt.Println(columns[i], ": ", value)
		}

	}
	return rowMap, nil
}

// GetFields 获取字段类型
func (s *Service) GetFields(tableName string) (fieldsType map[string]string, err error) {
	fields, err := s.getFields(tableName)
	if err != nil {
		return nil, err
	}
	var rowMap = make(map[string]string)
	for _, item := range fields {
		rowMap[item] = ""
	}
	//
	return rowMap, nil
}

// GetFieldsName 获取字段类型
func (s *Service) GetFieldsName(tableName string) (fieldsType map[string]string, err error) {
	fields, err := s.getFields(tableName)
	if err != nil {
		return nil, err
	}

	// 获取配置数据
	getConfigSQL := fmt.Sprintf("SELECT %s FROM `%s%s` WHERE SYS_ID = 4; ", strings.Join(fields, ","), tableName, conf.Conf.DBConfig.ConfigTableSuffix)
	// 执行
	fmt.Printf("\n\n\ngetConfigSQL: %#v\n\n\n", getConfigSQL)
	rows, err := s.db.ExecuteSQLRows(getConfigSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 获取类数据
	columns, _ := rows.Columns()

	// 以下代码来自官方 https://github.com/go-sql-driver/mysql/wiki/Examples
	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	var rowMap = make(map[string]string)
	// Fetch rows
	if rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		var value string

		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			rowMap[columns[i]] = value
			// fmt.Println(columns[i], ": ", value)
		}

	}
	return rowMap, nil

}
