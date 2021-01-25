package test

import v1 "github.com/linqiurong2021/go-excel-import/api/service/v1"

func TestReadExecl() {
	execl := v1.NewTemplate()
	execl.ReadTemplate("./template/template_test.xlsx")
}
