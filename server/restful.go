package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/linqiurong2021/go-excel-import/db"
	v1 "github.com/linqiurong2021/go-excel-import/service/v1"
)

// Restful Restful
type Restful struct {
	// 数据库
	db    *db.DB
	logic *v1.Logic
}

// NewRestful NewRestful
func NewRestful(db *db.DB) *Restful {
	return &Restful{
		db:    db,
		logic: v1.NewLogic(db),
	}
}

// GetConfig GetConfig
func (rest *Restful) GetConfig(w http.ResponseWriter, r *http.Request) {
	tableName := r.URL.Query().Get("table")
	if tableName == "" {
		fmt.Fprintln(w, "table params must")
	}
	data, err := rest.logic.GetConfig(tableName)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, string(jsonData))
	return
}

// SelectOptionsRequest SelectOptionsRequest
type SelectOptionsRequest struct {
	Table string `json:"table"` // 需要大写
	Keys  string `json:"keys"`
}

// GetSelectOptions 获取下拉选项数据
func (rest *Restful) GetSelectOptions(w http.ResponseWriter, r *http.Request) {
	// 获取body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return
	}
	//
	fmt.Printf("json:%s\n", string(body))

	var selectOption SelectOptionsRequest
	if err := json.Unmarshal(body, &selectOption); err != nil {
		fmt.Fprintln(w, "params must")
		return
	}
	fmt.Printf("selectOption: %#v\n", selectOption)
	// tableName := r.PostFormValue("table")
	// keys := r.PostFormValue("keys")

	if selectOption.Table == "" {
		fmt.Fprintln(w, "table params must")
		return
	}
	// 获取需要下拉的字段
	if selectOption.Keys == "" {
		fmt.Fprintln(w, "keys params must")
		return
	}

	data, err := rest.logic.GetSelectOptions(selectOption.Table, selectOption.Keys)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, string(jsonData))
	return
}

// TableListRequest 列表请求数据
type TableListRequest struct {
	Table    string `json:"table"`
	Params   string `json:"params"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}

// TableListRequest1 列表请求数据
type TableListRequest1 struct {
	Table    string            `json:"table"`
	Params   map[string]string `json:"params"`
	Page     int               `json:"page"`
	PageSize int               `json:"page_size"`
}

// 获取搜索的字段与数据
func (rest *Restful) getSearchKeyAndValues(params map[string]string) (searchFields []string, searchValues []string) {
	//
	fmt.Printf("len params %d\n", len(params))
	if len(params) <= 0 {
		return
	}
	// 地址分割
	for k, v := range params {
		searchFields = append(searchFields, k)
		searchValues = append(searchValues, v)
	}
	return
}

// GetListByPage 获取分页数据
func (rest *Restful) GetListByPage(w http.ResponseWriter, r *http.Request) {
	// 获取body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return
	}
	// 列表请求数据
	var tableList TableListRequest1
	if err := json.Unmarshal(body, &tableList); err != nil {
		fmt.Fprintln(w, "json invalidate")
		return
	}
	if tableList.Table == "" {
		fmt.Fprintln(w, "table params must")
		return
	}
	var params = make(map[string]string, 3)
	//
	params["table"] = tableList.Table
	params["page"] = fmt.Sprintf("%d", tableList.Page)
	params["page_size"] = fmt.Sprintf("%d", tableList.PageSize)

	// 搜索字段与对应值
	searchFields, searchValues := rest.getSearchKeyAndValues(tableList.Params)
	//
	data, err := rest.logic.GetTableListByPage(params, searchFields, searchValues)

	if err != nil {
		fmt.Printf("error: %s\n", err)
		fmt.Fprintln(w, err)
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("conver to json error: %s\n", err)
		fmt.Fprintln(w, err)
	}

	fmt.Fprintln(w, string(jsonData))
	return
	// rest.logic.GetTableListByPage(params)
}

// GetDataByID GetDataByID
func (rest *Restful) GetDataByID(w http.ResponseWriter, r *http.Request) {
	tableName := r.URL.Query().Get("table")
	if tableName == "" {
		fmt.Fprintln(w, "table params must")
	}
	strID := r.URL.Query().Get("sysID")
	// 转64位
	sysID, err := strconv.ParseInt(strID, 10, 64)
	fmt.Printf("#SYS_ID:%#v\n", sysID)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	//
	data, err := rest.logic.GetDataByID(tableName, sysID)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprintln(w, string(jsonData))
	return
}

// UpdateRequest 更新请求数据
type UpdateRequest struct {
	Table  string `json:"table"`
	Params string `json:"params"`
}

// UpdateBySysID UpdateBySysID
func (rest *Restful) UpdateBySysID(w http.ResponseWriter, r *http.Request) {
	// 获取body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return
	}
	// 列表请求数据
	var updateRequest UpdateRequest
	if err := json.Unmarshal(body, &updateRequest); err != nil {
		fmt.Fprintln(w, "json invalidate")
		return
	}
	if updateRequest.Table == "" {
		fmt.Fprintln(w, "table params must")
		return
	}
	var paramsMap map[string]interface{}
	err = json.Unmarshal([]byte(updateRequest.Params), &paramsMap)
	if err != nil {
		fmt.Println("params formate invalidate err: ", err)
	}
	result, err := rest.logic.UpdateBySysID(updateRequest.Table, paramsMap)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	affected, err := result.RowsAffected()
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, affected)
	return
}

// DeleteRequest 删除请求数据
type DeleteRequest struct {
	Table  string `json:"table"`
	SysIDs string `json:"sys_ids"`
}

// DeleteBySysIDs DeleteBySysIDs
func (rest *Restful) DeleteBySysIDs(w http.ResponseWriter, r *http.Request) {
	// 批量或
	// 获取body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return
	}
	// 列表请求数据
	var deleteRequest DeleteRequest
	if err := json.Unmarshal(body, &deleteRequest); err != nil {
		fmt.Fprintln(w, "json invalidate")
		return
	}
	if deleteRequest.Table == "" {
		fmt.Fprintln(w, "table params must")
		return
	}
	if len(deleteRequest.SysIDs) <= 0 {
		fmt.Fprintln(w, "sys_id params must")
		return
	}
	result, err := rest.logic.DeleteBySysIDs(deleteRequest.Table, deleteRequest.SysIDs)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	affected, err := result.RowsAffected()
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, affected)
	return
}

// GetFieldsType 获取字段类型
func (rest *Restful) GetFieldsType(w http.ResponseWriter, r *http.Request) {
	tableName := r.URL.Query().Get("table")
	if tableName == "" {
		fmt.Fprintln(w, "table params must")
	}
	data, err := rest.logic.GetFieldsType(tableName)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprintln(w, string(jsonData))
	return
}

// GetFieldsName 获取字段类型
func (rest *Restful) GetFieldsName(w http.ResponseWriter, r *http.Request) {
	tableName := r.URL.Query().Get("table")
	if tableName == "" {
		fmt.Fprintln(w, "table params must")
	}
	data, err := rest.logic.GetFieldsName(tableName)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	jsonData, err := json.Marshal(data)
	fmt.Printf("#JSONDATA#%#v\n", jsonData)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprintln(w, string(jsonData))
	return
}

// GetFields 获取字段
func (rest *Restful) GetFields(w http.ResponseWriter, r *http.Request) {
	tableName := r.URL.Query().Get("table")
	if tableName == "" {
		fmt.Fprintln(w, "table params must")
	}
	data, err := rest.logic.GetFields(tableName)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	jsonData, err := json.Marshal(data)
	fmt.Printf("#JSONDATA#%#v\n", jsonData)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprintln(w, string(jsonData))
	return
}

// CreateData CreateData
func (rest *Restful) CreateData(w http.ResponseWriter, r *http.Request) {
	// 获取body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return
	}
	// 列表请求数据
	var updateRequest UpdateRequest
	if err := json.Unmarshal(body, &updateRequest); err != nil {
		fmt.Fprintln(w, "json invalidate")
		return
	}
	if updateRequest.Table == "" {
		fmt.Fprintln(w, "table params must")
		return
	}
	var paramsMap map[string]interface{}
	err = json.Unmarshal([]byte(updateRequest.Params), &paramsMap)
	if err != nil {
		fmt.Println("params formate invalidate err: ", err)
	}
	result, err := rest.logic.CreateData(updateRequest.Table, paramsMap)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	affected, err := result.RowsAffected()
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, affected)
	return
}

// Import 导入
func (rest *Restful) Import(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20) // limit your max input length!
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	var buf bytes.Buffer
	// in your case file would be fileupload
	file, header, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	name := strings.Split(header.Filename, ".")
	// Copy the file data to my buffer
	io.Copy(&buf, file)

	fmt.Printf("content: %v\n", buf.Bytes())

	// do something else
	path, err := os.Getwd()
	if err != nil {
		//
		fmt.Fprintln(w, err)
		return
	}
	suffix := name[len(name)-1]
	if suffix != "xls" && suffix != "xlsx" {
		fmt.Fprintln(w, "只支持xlsx 或 xls 文件格式")
		return
	}
	filename := fmt.Sprintf("%s/uploads/%s_%d.%s", path, name[0], time.Now().Unix(), suffix)
	fmt.Printf("path:%s\n", filename)

	err = ioutil.WriteFile(filename, buf.Bytes(), 0644)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	// 重置
	buf.Reset()
	err = rest.logic.CreateTable(filename, false)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprint(w, "success")
	return
}

// ExportRequest 导出请求参数
type ExportRequest struct {
	Table  string `json:"table"`
	SysIDs string `json:"sys_ids"`
}

// Export 导出
func (rest *Restful) Export(w http.ResponseWriter, r *http.Request) {
	// 获取body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return
	}
	// 列表请求数据
	var exportRequest ExportRequest
	if err := json.Unmarshal(body, &exportRequest); err != nil {
		fmt.Fprintln(w, "json invalidate")
		return
	}
	if exportRequest.Table == "" {
		fmt.Fprintln(w, "table params must")
		return
	}
	fileName, content, err := rest.logic.ExportData(exportRequest.Table, exportRequest.SysIDs)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	// fmt.Printf("fileName: %s\n", fileName)
	// 添加头部
	w.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fileName))
	w.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	// 文件
	http.ServeContent(w, r, fileName, time.Now(), content)

	return
}

// GetTemplates 获取已有的模板
func (rest *Restful) GetTemplates(w http.ResponseWriter, r *http.Request) {
	list, err := rest.logic.GetSysTemplateList()
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	jsonData, err := json.Marshal(list)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, string(jsonData))
	return
}

// StartServer 启用Server
func (rest *Restful) StartServer() {
	// 获取下拉或单选或多选项
	http.HandleFunc("/getSelectOptions", rest.GetSelectOptions)
	// 获取配置信息(字段填写类型 字段名称等)
	http.HandleFunc("/getConfig", rest.GetConfig)
	// 通过ID 获取数据
	http.HandleFunc("/getDataByID", rest.GetDataByID)
	// 获取字段填写类型
	http.HandleFunc("/getFieldsType", rest.GetFieldsType)
	// 获取字段名称
	http.HandleFunc("/getFieldsName", rest.GetFieldsName)
	// 获取字段
	http.HandleFunc("/getFields", rest.GetFields)
	// 列表分页
	http.HandleFunc("/getListByPage", rest.GetListByPage)
	// 列表分页
	http.HandleFunc("/getListByPage1", rest.GetListByPage)
	// 数据更新
	http.HandleFunc("/updateBySysID", rest.UpdateBySysID)
	// 删除
	http.HandleFunc("/deleteBySysIDs", rest.DeleteBySysIDs)
	// 新增
	http.HandleFunc("/createData", rest.CreateData)
	// 导入
	http.HandleFunc("/import", rest.Import)
	// 导出
	http.HandleFunc("/export", rest.Export)
	// 获取已有的模板
	http.HandleFunc("/getTemplates", rest.GetTemplates)
	fmt.Println("\n StartServer \n ")
	//
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf("http listen failed %s\n", err)
	}
}
