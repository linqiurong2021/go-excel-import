package v1

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/linqiurong2021/go-excel-import/api/db"
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

	var field string
	for rows.Next() {
		//
		if err = rows.Scan(&field); err != nil {
			return
		}
		fields = append(fields, field)
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
func (s *Service) getParamsFieldsAndValues(params map[string]string) string {
	// 需要判断是什么类型的 如果是下拉则直接使用等于 其它使用 like %%
	return ""
}

// GetAllData 获取所有数据
func (s *Service) GetAllData(tableName string, params map[string]string) (rows *sql.Rows, err error) {
	// 获取字段
	s.getFields(tableName)
	fields := "*"
	//
	var getAllSQL string
	// 获取参数中的条件
	where := s.getParamsFieldsAndValues(params)
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
func (s *Service) GetDataByPageSize(tableName string, params map[string]string) (rows *sql.Rows, err error) {
	// 获取字段
	s.getFields(tableName)
	fields := "*"
	//
	var getAllSQL string
	// pageSize
	pageSize, err := strconv.Atoi(params["page_size"])
	if err != nil {
		return nil, err
	}
	pageNum, err := strconv.Atoi(params["page"])
	if err != nil {
		return nil, err
	}
	offset := (pageNum - 1) * pageSize

	// 获取参数中的条件
	where := s.getParamsFieldsAndValues(params)
	if len(where) == 0 {
		// 获取所有的数据
		getAllSQL = fmt.Sprintf("SELECT %s FROM %s LIMIT %d, %d;", tableName, fields, offset, pageNum)
	} else {
		// 获取所有的数据 分页
		getAllSQL = fmt.Sprintf("SELECT %s FROM %s WHRE %s LIMIT %d,%d;", tableName, fields, where, offset, pageNum)
	}
	rows, err = s.db.ExecuteSQLRows(getAllSQL)
	// 返回数据
	return rows, nil

}

// GetDataByID 获取一条数据
func (s *Service) GetDataByID(tableName string, sysID int64) (rows *sql.Rows, err error) {
	// 获取字段
	s.getFields(tableName)
	fields := "*"
	getSQL := fmt.Sprintf("SELECT %s FROM %s WHERE SYS_ID = %d", fields, tableName, sysID)
	rows, err = s.db.ExecuteSQLRows(getSQL)
	return
}

// AddData 新增一条数据
func (s *Service) AddData(tableName string, params map[string]string) (result sql.Result, err error) {
	// 获取字段
	s.getFields(tableName)
	fields := "*"
	// 数据处理
	inserData := ""
	getSQL := fmt.Sprintf("INSERT INTO %s(%s) VALUES %s", fields, tableName, inserData)
	result, err = s.db.ExecuteSQLResult(getSQL)
	return
}

// UpdateDataByID 更新一条数据
func (s *Service) UpdateDataByID(tableName string, params map[string]string) (result sql.Result, err error) {
	// 获取字段
	s.getFields(tableName)
	// 数据处理
	updateData := ""
	updateSQL := fmt.Sprintf("UPDATE TABLE %s SET %s", tableName, updateData)
	result, err = s.db.ExecuteSQLResult(updateSQL)
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
