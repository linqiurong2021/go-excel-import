package v1

import (
	"fmt"
	"strings"
)

// Config 配置文件
type Config struct {
	Service *Service
}

// GetConfigFields 获取配置文件
func (c *Config) GetConfigFields(tableName string) (fields []string, err error) {
	//
	return c.Service.getFields(tableName)
}

// GetConfig 获取配置信息
func (c *Config) GetConfig(tableName string) (getConfigSQL string, err error) {
	// 字段
	fields, err := c.GetConfigFields(tableName)
	if err != nil {
		return "", err
	}
	// 获取配置数据
	getConfigSQL = fmt.Sprintf("SELECT %s FROM %s_config ORDER BY SYS_ID", strings.Join(fields, ","), tableName)
	//
	return getConfigSQL, nil
}
