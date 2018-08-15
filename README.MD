# onenet_device_export
OneNET设备批量下载

## 快速使用
下载程序

编辑conf.ini中的api-key改为你的key

运行程序

## conf.ini中查询的一些参数解释
key_words=my, //匹配关键字（可选，从id和title字段中左匹配）

auth_info=203x1 //鉴权信息（可选，对应注册时的sn参数，唯一设备编号）

tag=mobile //标签（可选）

online=true| false //在线状态（可选）

private=true| false //私密性（可选）

begin=2016-06-20 //起始时间，包括当天（可选）

end=2016-06-20 //结束时间，包括当天（可选）

## 源码编译
golang 1.8+

go get https://github.com/limao777/onenet_device_export

再进入该目录 go build即可

## 更新历史
2018-08-15
变更url为从配置文件获取，兼容各个省平台

2018-08-10
增加查询参数

2018-08-09
基础全量版研发完毕