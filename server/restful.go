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

// UpdateData UpdateData
func (rest *Restful) UpdateData(w http.ResponseWriter, r *http.Request) {
	tableName := r.URL.Query().Get("table")
	if tableName == "" {
		fmt.Fprintln(w, "table params must")
	}
	//
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

// StartServer 启用Server
func (rest *Restful) StartServer() {

	http.HandleFunc("/getSelectOptions", rest.GetSelectOptions)
	//
	http.HandleFunc("/getConfig", rest.GetConfig)
	http.HandleFunc("/getDataByID", rest.GetDataByID)
	http.HandleFunc("/getFieldsType", rest.GetFieldsType)
	http.HandleFunc("/getFieldsName", rest.GetFieldsName)
	//
	http.HandleFunc("/getListByPage", rest.GetListByPage)

	fmt.Println("\n StartServer \n ")
	//
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf("http listen failed %s\n", err)
	}
}
