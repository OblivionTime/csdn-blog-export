/*
 * @Description:
 * @Version: 1.0
 * @Autor: solid
 * @Date: 2022-09-27 14:58:04
 * @LastEditors: solid
 * @LastEditTime: 2022-09-27 16:23:58
 */
package conf

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Config CSDNConfig
var GVA_VP *viper.Viper

type CSDNConfig struct {
	Username  string `yaml:"username" json:"username"`   //用户名
	Cookie    string `yaml:"cookie" json:"cookie"`       //Cookie
	SaveExcel bool   `yaml:"saveExcel" json:"saveExcel"` //将数据保存成excel
	FileType  string `yaml:"fileType" json:"fileType"`   //生成文件的类型(all(html和md),html,md)
	TotalPage int    `yaml:"totalPage" json:"totalPage"` //总页数

}

// 初始化Viper配置文件
func InitViper(path string) *viper.Viper {
	v := viper.New()
	//设置配置文件
	v.SetConfigFile(path)
	//设置配置文件类型
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	//监听配置文件的修改
	v.WatchConfig()
	//当配置文件发送改变
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&Config); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&Config); err != nil {
		fmt.Println("***************", err)
	}
	return v
}
