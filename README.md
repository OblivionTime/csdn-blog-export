
# csdn-blog-export
基于golang实现CSDN一键导出所有文章,支持导成html和md
支持生成博客数据表

# 基本使用
先修改config.yaml文件,修改成你自己的信息
```yaml
# 用户名
username: test
# cookie
cookie: test
# 总页数(每一页100条)
totalPage: 2
#将数据保存成excel
saveExcel: true
# 生成文件的类型(all(html和md),html,md)
fileType: md
```
将config.yaml文件和csdn.exe文件放在同一级目录下,执行以下命令
```shell
csdn.exe 
```
