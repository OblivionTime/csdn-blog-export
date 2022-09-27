/*
 * @Description:
 * @Version: 1.0
 * @Autor: solid
 * @Date: 2022-09-22 17:00:57
 * @LastEditors: solid
 * @LastEditTime: 2022-09-27 16:32:28
 */
package csdn

import (
	"CSDN/conf"
	"fmt"

	"github.com/xuri/excelize/v2"
)

func SaveExcel(ArticleList []Article) {
	f := excelize.NewFile()
	// 创建一个工作表
	index := f.GetSheetIndex("Sheet1")

	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	f.SetCellValue("Sheet1", "A1", "文章id")
	f.SetCellValue("Sheet1", "B1", "文章标题")
	f.SetCellValue("Sheet1", "C1", "描述")
	f.SetCellValue("Sheet1", "D1", "链接")
	f.SetCellValue("Sheet1", "E1", "阅读量")
	f.SetCellValue("Sheet1", "F1", "评论数")
	f.SetCellValue("Sheet1", "G1", "编辑URL")
	f.SetCellValue("Sheet1", "H1", "文章发布时间")
	f.SetCellValue("Sheet1", "I1", "文章发布日期")
	for index, art := range ArticleList {
		// 设置单元格的值
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", index+2), art.ArticleId)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", index+2), art.Title)
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", index+2), art.Description)
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", index+2), art.URL)
		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", index+2), art.ViewCount)
		f.SetCellValue("Sheet1", fmt.Sprintf("F%d", index+2), art.CommentCount)
		f.SetCellValue("Sheet1", fmt.Sprintf("G%d", index+2), art.EditUrl)
		f.SetCellValue("Sheet1", fmt.Sprintf("H%d", index+2), art.PostTime)
		f.SetCellValue("Sheet1", fmt.Sprintf("I%d", index+2), art.FormatTime)
	}

	fileName := fmt.Sprintf("./exportData/%v.xlsx", conf.Config.Username+"博客数据表")
	// 根据指定路径保存文件
	if err := f.SaveAs(fileName); err != nil {
		fmt.Println(err)
	}
	fmt.Println("保存成功!!!!!!")
}
