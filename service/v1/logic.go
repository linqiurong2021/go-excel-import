package v1

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/fatih/set"
	"github.com/linqiurong2021/go-excel-import/conf"
	"github.com/linqiurong2021/go-excel-import/db"
	"github.com/linqiurong2021/go-excel-import/utils"
)

// Logic 逻辑
type Logic struct {
	// 数据库
	db *db.DB
	// 模型
	template *Template
	//
	service *Service
}

// NewLogic 创建逻辑
func NewLogic(db *db.DB) *Logic {
	return &Logic{
		db:       db,
		template: &Template{},
		service: &Service{
			db: db,
		},
	}
}

// backupTable 备份数据表
func (l *Logic) backupTable(tx *sql.Tx) error {
	// 需要备份的数据表名
	tableEnName := l.template.TableEnName
	backupEnName := fmt.Sprintf("%s_%d", tableEnName, time.Now().Unix())
	// fmt.Println("数据备份表名:", tableEnName)
	// 备份操作 https://blog.csdn.net/moshowgame/article/details/82952992
	backupSQL := fmt.Sprintf("CREATE TABLE %s SELECT * FROM %s ", backupEnName, tableEnName)
	fmt.Printf("备份数据%s表\n新表名为:%s\n 执行语句为:%s\n", tableEnName, backupEnName, backupSQL)

	result, err := tx.Exec(backupSQL)
	if err != nil {
		return err
	}
	fmt.Printf("backup result %#v\n", result)
	return nil
}

// ExistTable 判断表名是否存在
func (l *Logic) existTable(tableName string) (exists bool, err error) {
	// 判断表是否存在
	existTableSQL := fmt.Sprintf("SELECT * FROM information_schema.tables WHERE TABLE_NAME ='%s' and TABLE_SCHEMA = '%s';", tableName, conf.Conf.DBConfig.Sechma)
	rows, err := l.db.ExecuteSQLRows(existTableSQL)
	if err != nil {
		return false, err
	}
	// 如果存在
	if rows.Next() {
		return true, nil
	}
	return false, nil
}

// 判断原表字段是否包括新表字(包括则说明是插入数据)
func (l *Logic) oldFieldHasNotNewFields(fields []string, newFields []string) (diff []string) {
	// 判断 newFields是否是fields的子集
	oldFieldSet := set.New(set.NonThreadSafe)
	oldFieldSet.Add(strings.Join(fields, ","))
	//
	// fmt.Printf("oldFieldSet %#v\n", oldFieldSet.List())

	newFieldSet := set.New(set.NonThreadSafe)
	newFieldSet.Add(strings.Join(newFields, ","))
	// fmt.Printf("newFieldSet %#v\n", newFieldSet.List())

	defer oldFieldSet.Clear()
	defer newFieldSet.Clear()

	return utils.Difference(newFields, fields)
}

// CreateTable 创建数据表并添加数据
func (l *Logic) CreateTable(templatePath string, needBackup bool) error {
	// 如果不存在则直接创建
	// createTableSQL, insertSQL, err := l.template.ReadTemplate(templatePath)
	table, err := l.template.ReadTemplate(templatePath)
	fmt.Printf("## %#v ##\n", table)
	// 判断表是否存在
	exists, err := l.existTable(l.template.TableEnName)
	if err != nil {
		return err
	}
	//
	oldFields, err := l.service.getFields(l.template.TableEnName)
	if err != nil {
		return err
	}
	fmt.Printf("exists %v, Table Fields  %v\n", exists, oldFields)
	// 原表中没有的字段 如果长度为0则说明是子集
	diff := l.oldFieldHasNotNewFields(oldFields, l.template.FieldsEnName)
	// 如果存在且不是原字段的子集 则提示表名已存在
	if exists && len(diff) > 0 {
		// 提示数据库中的表,无新字段%s
		errText := fmt.Sprintf("\nTemplateName %s, 且无新字段: %s\n", l.template.TableEnName, strings.Join(diff, ","))
		return errors.New(errText)
	} else if exists && len(diff) <= 0 { // 如果表名相同且原来字段包括新字段则直接插入数据
		// 直接添加数据
		result, err := l.db.ExecuteSQLResult(table.insertTableSQL)
		//
		if err != nil {
			return err
		}
		fmt.Printf("直接新增数据: %#v\n", result)
		return nil
	} else { // 不存在情况下
		// 开启事务
		tx, err := l.db.DB.Begin()
		if err != nil {
			return err
		}
		// 如果表存在且需要备份 开启备份数据
		if needBackup && exists {
			// 先备份再进行操作
			if err := l.backupTable(tx); err != nil {
				tx.Rollback()
				return err
			}
		}
		fmt.Printf("TableName %s , Exists TableName %v \n", l.template.TableEnName, exists)

		result, err := tx.Exec(table.tableSQL)
		if err != nil {
			tx.Rollback()
			return err
		}
		// 创建配置表
		result, err = tx.Exec(table.configTableSQL)
		if err != nil {
			tx.Rollback()
			return err
		}
		// 配置数据
		result, err = tx.Exec(table.insertTableSQL)
		if err != nil {
			tx.Rollback()
			return err
		}
		// 创建配置数据表
		result, err = tx.Exec(table.insertConfigSQL)
		if err != nil {
			tx.Rollback()
			return err
		}

		// 事务提交
		fmt.Printf("tableData %#v\n, Result %#v\n", table, result)
		return tx.Commit()
	}
}