package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

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

// 获取搜索的字段与数据
func (rest *Restful) getSearchKeyAndValues(params string) (searchFields []string, searchValues []string) {
	//
	fmt.Printf("len params %d\n", len(params))
	if len(params) <= 0 {
		return
	}
	// 地址分割
	searchParams := strings.Split(params, "&")
	for _, item := range searchParams {
		tmp := strings.Split(item, "=")
		// 如果为空则不需要搜索
		if tmp[1] != "" {
			searchFields = append(searchFields, tmp[0])
			searchValues = append(searchValues, tmp[1])
		}
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
	var tableList TableListRequest
	if err := json.Unmarshal(body, &tableList); err != nil {
		fmt.Fprintln(w, "json invalidate")
		return
	}
	if tableList.Table == "" {
		fmt.Fprintln(w, "table params must")
		return
	}
	var params = make(map[string]string, 4)
	//
	params["table"] = tableList.Table
	params["page"] = fmt.Sprintf("%d", tableList.Page)
	params["page_size"] = fmt.Sprintf("%d", tableList.PageSize)
	params["params"] = tableList.Params
	// 搜索字段与对应值
	searchFields, searchValues := rest.getSearchKeyAndValues(tableList.Params)
	//
	fmt.Printf("searchFields: %#v\n", searchFields)
	fmt.Printf("searchValues: %#v\n", searchValues)
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
	// 数据更新
	http.HandleFunc("/updateBySysID", rest.UpdateBySysID)
	// 删除
	http.HandleFunc("/deleteBySysIDs", rest.DeleteBySysIDs)
	// 新增
	http.HandleFunc("/createData", rest.CreateData)
	fmt.Println("\n StartServer \n ")
	//
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf("http listen failed %s\n", err)
	}
}
