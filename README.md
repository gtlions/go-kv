# go-kv
一个玩具而已，```Golang```实现的KV存储系统。
整理文档时候发现的老代码，是早前学习```Gin```框架的练手项目，最近给重新整理了下。

> ⚠️代码很菜

But... **正所谓**

> 生命不息 造轮不止


# 使用方法
使用```http```方式操作KV数据，支持查询、创建、更新和删除操作。

- Ping Server
```
GET http://127.0.0.1:8080/ping
```

- Get Key
```
GET http://127.0.0.1:8080?gtlions
```

- Add/Update Key-Value
```
POST http://127.0.0.1:8080
Content-Type: application/json

{"key":"gtlions",
"value":"gtlions.l@qq.com"}
```

- Delete Key
```
DELETE http://127.0.0.1:8080?gtlions
```

# TODO
- 事务
- 持久化存储
- ...