# [Golang-Laravel](https://github.com/sk532359025/golang-laravel)


模仿PHP-Laravel，开箱即用、go语言版的laravel，phper转Go开发 可以立马上手~

#### package
* [gin](https://github.com/gin-gonic/gin)
* [gorm](http://gorm.book.jasperxu.com/)
* [go-redis](https://github.com/go-redis/redis)
* [go lavael](https://github.com/sk532359025/golang-laravel)
---
#### 支持功能
* 路由
* 中间件
* 路由中间件
* 控制器
* 模型
* 验证器（*默认是英文，将其翻译为中文）
---
#### 目录结构
与laravel几乎一模一样

    app					        应用目录
    |---- exception				        返回http状态及异常接收
	|----|---- baseException.go 	                panic异常接收文件
    |---- http					web核心目录
	|----|---- controller 		 	        控制器存放目录
	|----|---- middleware 		 	        路由中间件存放目录
    |---- models				        模型存放目录
    |---- service					业务编写目录
    |---- validate					验证器目录
    |---- helpers.go 				常用函数文件
    config
    |----|----app.go  				应用配置 如：ip、port等
    |----|----databases.go  			数据库配置
    extend						自定义扩展包
    |----|---- log		                        自定义日志package(与系统log相同，增加按小时分割日志)
    routes						路由文件夹
	|----|---- route.go				路由控制目录
	|----|---- .....				更多路由
    sql						测试sql文件存放目录
    test 						压力测试文件目录
    |----|---- abtest.sh                            ab压力测试文件
    |----|---- .....				配置 Or Log文件          
    vendor						扩展包目录
	|----|---- .....				包存放


* 还在完善中
* 目前发现不合理地方 在返回success状态码 目前是用panic抛出异常返回，有为常理。
* 注意路由文件编写，必须 在router.go文件中调用
* 将sql/databases.go 文件 数据连接信息一改，运行 go run main.go 即可使用

---
#### 项目注意事项
1. 重新定义 log(`/extend/log`)模块 将日志分割 storage/logs/`Year`/`Month`/`Day`/`Hour.txt`
2. gorm 更改了 软删除 默认是 is null(`/vendor/github.com/jinzhu/gorm/scope.go line 719`) 本项目更改为 deleted_at=0 根据自己项目需求更改
3. `config/database.go` mysql 连接配置 需
4. `config/redis.go` redis 连接配置 需修改
5. 路由 在 `/routes` 中定义 要在 route.go 引入自己创建的 目前并不会自动引入
6. 运行命令  go run main.go


---
#### go语言常用简介
`自己总结有误则提醒，修改`
1. defer 延迟执行 defer标识的函数 在当前函数末尾执行 多个defer时以“压栈”的方式执行（先进后出）
    例：main.go 中的 mysql、redis关闭连接
2. struct 结构体 go 语言没有class的概念 但是可以通过结构体来实现 class
3. 输出到控制台 fmt.Println("string") 本项目中 可以使用log.Println("....") 输出到控制台同时记录log
4. 常用数据结构 map int string float64 struct interface等
5. 在数据类型是interface时 引用时 要进行‘断言’ `示例：/app/models/baseModel.go line 29`
---
#### 更新log

##### 2020-11-09 
1. mysql连接池验证 压力测试
2. 增加自定义log模块 (在config中自定义log路径，package path /extend/log，按小时分割日志)
3. 增加压力测试模块
##### 2020-11-06
1. 增加mysql连接池
##### 2020-11-10
1. 增加redis 连接池
##### 2020-11-11
1. 修改log模块 获取当天日期
2. 增加autoload 将方法封装 使入口文件简单清晰
