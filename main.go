/*
 * @Description:
 * @Version: 1.0
 * @Autor: solid
 * @Date: 2022-09-27 14:56:10
 * @LastEditors: solid
 * @LastEditTime: 2022-09-27 15:50:44
 */
package main

import (
	csdn "CSDN/CSDN"
	"CSDN/conf"
)

func main() {
	conf.InitViper("config.yaml")
	csdn.GetCSDNArticle()
}
