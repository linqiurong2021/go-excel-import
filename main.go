package main

import v1 "github.com/linqiurong2021/go-excel-import/api/service/v1"

func main() {
	execl := v1.NewExcel()
	execl.ReadExcel()
}
