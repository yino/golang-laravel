# [Golang-Laravel](https://github.com/sk532359025/golang-laravel)


模仿PHP-Laravel，开箱即用

#### package
* gin
* gorm
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

2020-11-09 更新
1. mysql连接池验证 
2. 增加自定义log模块 (在config中自定义log路径，package path /extend/log，按小时分割日志)

