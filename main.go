package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"github.com/tealeg/xlsx"
	"github.com/tidwall/gjson"
	"strings"
)

func main() {
	fileptr, err := os.Open("./data.json")
	if err != nil {
		fmt.Println("open file failed")
		return
	}
	defer fileptr.Close()

	json, err := ioutil.ReadAll(fileptr)

	all := gjson.Get(string(json), "data.pageStats")
	file, err := xlsx.OpenFile("template.xlsx")
	if err != nil {
		panic(err)
	}
	first := file.Sheets[0]
	first_row := first.AddRow()
	frst_cell := first_row.AddCell()
	frst_cell.Value = string("时间")
	frst_cell = first_row.AddCell()
	frst_cell.Value = string("分类")
	frst_cell = first_row.AddCell()
	frst_cell.Value = string("类型")
	frst_cell = first_row.AddCell()
	frst_cell.Value = string("金额")
	frst_cell = first_row.AddCell()
	frst_cell.Value = string("账户1")
	frst_cell = first_row.AddCell()
	frst_cell.Value = string("账户2")
	frst_cell = first_row.AddCell()
	frst_cell.Value = string("备注")
	for _, allPage := range all.Array() {
		page := gjson.Get(allPage.Raw, "pageExpenses")
		for _, pageJson := range page.Array() {
			details := pageJson.Get("expenses")
			for _, detail := range details.Array() {
				account := detail.Get("accountName").Str
				cost := detail.Get("cost").Num
				costz := fmt.Sprintf("%.2f", cost)
				time := strings.Replace(detail.Get("cts").Str, "-", "/", -1)
				type_num := detail.Get("type").Num
				var consume_type string
				if type_num == 0 {
					consume_type = "支出"
				} else {
					consume_type = "收入"
				}
				category1 := detail.Get("categoryName").Str
				remark := strings.Replace(detail.Get("remark").Str, "#", "", -1)
				account2 := ""

				// fmt.Println("时间:" + time + ",分类:" + category1 + ",类型:" + consume_type + ",金额:" + costz + "账户1:" + account + ",账户2:" + account2 + ",备注:" + remark)

				row := first.AddRow()
				cell := row.AddCell()
				cell.Value = time
				cell = row.AddCell()
				cell.Value = category1
				cell = row.AddCell()
				cell.Value = consume_type
				cell = row.AddCell()
				cell.Value = costz
				cell = row.AddCell()
				cell.Value = account
				cell = row.AddCell()
				cell.Value = account2
				cell = row.AddCell()
				cell.Value = remark
			}
		}
	}
	err = file.Save("qianji.xlsx")
	if err != nil {
		panic(err)
	}

}
