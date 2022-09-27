/*
 * @Description:
 * @Version: 1.0
 * @Autor: solid
 * @Date: 2022-09-27 15:01:22
 * @LastEditors: solid
 * @LastEditTime: 2022-09-27 16:32:53
 */
package csdn

import (
	"CSDN/conf"
	"CSDN/utils"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Article struct { //用来解析json
	ArticleId    int    `json:"articleId"`    //文章id
	Title        string `json:"title"`        //文章标题
	Description  string `json:"description"`  //描述
	URL          string `json:"url"`          //链接
	ViewCount    int    `json:"viewCount"`    //阅读量
	CommentCount int    `json:"commentCount"` //评论数
	EditUrl      string `json:"editUrl"`      //编辑URL
	PostTime     string `json:"postTime"`     //文章发布时间
	FormatTime   string `json:"formatTime"`   //文章发布日期
}
type GetBusinessList struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    CSDNData `json:"data"`
}
type CSDNData struct {
	List  []Article `json:"list"`
	Total int       `json:"total"`
}
type GetArticle struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    ArticleDetail `json:"data"`
}
type ArticleDetail struct {
	ArticleId       string `json:"article_id"`      //文章id
	Content         string `json:"content"`         //html内容
	Markdowncontent string `json:"markdowncontent"` //MD内容
	Title           string `json:"title"`           //标题
}

func GetCSDNArticle() {
	var ArticleList []Article

	for i := 1; i <= conf.Config.TotalPage; i++ {
		time.Sleep(100) //设置延时
		url := fmt.Sprintf("https://blog.csdn.net/community/home-api/v1/get-business-list?page=%d&size=100&businessType=blog&orderby=&noMore=false&year=&month=&username=%s", i, conf.Config.Username)
		resp := utils.SendMessageServer(url, nil)
		var getBusinessList GetBusinessList
		err := json.Unmarshal(resp, &getBusinessList)
		if err != nil {
			fmt.Println(err)
			return
		}
		if getBusinessList.Code == 200 {
			artList := getBusinessList.Data.List
			if len(artList) == 0 {
				break
			}
			ArticleList = append(ArticleList, artList...)
		}
	}
	if !utils.IsExist("./exportData") {
		os.MkdirAll("./exportData", 777)
	}
	SaveExcel(ArticleList)
	for _, art := range ArticleList {
		detailURL := fmt.Sprintf("https://blog-console-api.csdn.net/v1/editor/getArticle?id=%d", art.ArticleId)
		resp := utils.SendMessageServer(detailURL, nil)
		var getArticle GetArticle
		err := json.Unmarshal(resp, &getArticle)
		if err != nil {
			fmt.Println(err)
			return
		}
		if getArticle.Code == 200 {
			t := conf.Config.FileType
			if t == "html" {
				utils.WriteWithIoutil("./exportData/"+getArticle.Data.Title+".html", getArticle.Data.Content)
			} else if t == "md" {
				utils.WriteWithIoutil("./exportData/"+getArticle.Data.Title+".md", getArticle.Data.Markdowncontent)
			} else {
				utils.WriteWithIoutil("./exportData/"+getArticle.Data.Title+".html", getArticle.Data.Content)
				utils.WriteWithIoutil("./exportData/"+getArticle.Data.Title+".md", getArticle.Data.Markdowncontent)
			}
		}
	}
}
