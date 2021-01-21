package interfacer

// Servicer Servicer
type Servicer interface {
	AddData() error
	GetAllData() error
	GetDataByPageSize() error
	GetDataByID() error
	UpdateDataByID() error
	ImportData() error
	ExportData() error
}
