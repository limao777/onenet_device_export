# onenet_device_export
OneNET设备批量下载

最新代码已迁移到
https://github.com/cm-heclouds/onenet_device_export

最新代码已迁移到
https://github.com/cm-heclouds/onenet_device_export

最新代码已迁移到
https://github.com/cm-heclouds/onenet_device_export

从2018年12月开始不再更新到本项目，访问https://github.com/cm-heclouds/onenet_device_export
获取最新程序

## 快速使用(适用windows 7及以上环境)
下载程序，进入exe目录

编辑app.conf中的api-key改为你的key，注意“(your api-key)”的括号不应保留

运行程序

## app.conf中查询的一些参数解释
key_words=my, //匹配关键字（可选，从id和title字段中左匹配）

auth_info=203x1 //鉴权信息（可选，对应注册时的sn参数，唯一设备编号）

tag=mobile //标签（可选）

online=true| false //在线状态（可选）

private=true| false //私密性（可选）

begin=2016-06-20 //起始时间，包括当天（可选）

end=2016-06-20 //结束时间，包括当天（可选）

## 从源码编译(windows以外系统或windows高级用户)
golang 1.8+

go get https://github.com/limao777/onenet_device_export

再进入该目录 go build即可

## 更新历史
2018-12-04

增加了windows 32位可执行文件的支持

更新了说明文档


2018-08-27

变更配置文件名字，尽量不让windows文本编辑器直接打开编辑，使配置文件使用UTF-8编码

url的key_word、tag进行URL编码以支持中文搜索


2018-08-15

变更url为从配置文件获取，兼容各个省平台


2018-08-10

增加查询参数


2018-08-09

基础全量版研发完毕