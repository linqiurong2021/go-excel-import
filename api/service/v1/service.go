package v1

// Service Service
type Service struct {
	// 需要操作的表名
	TableName string
}

// GetAllData 获取所有数据
func (s *Service) GetAllData() (err error) {
	// 返回数据
	return nil
}

// GetDataByPageSize 分页(需要有参数)
func (s *Service) GetDataByPageSize() (err error) {
	return nil
}

// GetDataByID 获取一条数据
func (s *Service) GetDataByID() (err error) {
	return nil
}

// AddData 新增一条数据
func (s *Service) AddData() (err error) {
	return nil
}

// UpdateDataByID 更新一条数据
func (s *Service) UpdateDataByID() (err error) {
	return nil
}

// ImportData 导入数据
func (s *Service) ImportData() (err error) {
	return nil
}

// ExportData 导出数据(需要有参数)
func (s *Service) ExportData() (err error) {
	return nil
}
